package rotatablesink_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"iptables-logger/fakes"
	"iptables-logger/rotatablesink"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/lager/lagertest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rotatewatcher", func() {
	var (
		fileToWatch             *os.File
		fileToWatchName         string
		rotatableSink           *rotatablesink.RotatableSink
		fakeTestWriterFactory   *TestWriterFactory
		fakeLogger              *lagertest.TestLogger
		fakeDestinationFileInfo *fakes.DestinationFileInfo
	)

	BeforeEach(func() {
		var err error
		fileToWatch, err = ioutil.TempFile("", "")
		Expect(err).NotTo(HaveOccurred())
		fileToWatchName = fileToWatch.Name()

		fakeTestWriterFactory = NewTestWriterFactory(fileToWatch, nil)
		fakeDestinationFileInfo = &fakes.DestinationFileInfo{}
		fakeLogger = lagertest.NewTestLogger("test")
		rotatableSink, err = rotatablesink.NewRotatableSink(fileToWatchName, lager.DEBUG, fakeTestWriterFactory, fakeDestinationFileInfo, fakeLogger)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(os.Remove(fileToWatchName)).To(Succeed())
	})

	Describe("NewRotatableSink", func() {

		Context("when unable to open the destination file that was rotated", func() {
			BeforeEach(func() {
				fakeTestWriterFactory.SetReturnedError(errors.New("banana"))
			})

			It("returns an sensible error", func() {
				var err error
				rotatableSink, err = rotatablesink.NewRotatableSink(fileToWatchName, lager.DEBUG, fakeTestWriterFactory, fakeDestinationFileInfo, fakeLogger)
				Expect(err).To(MatchError("register file sink: rotate file sink: create file writer: banana"))
			})
		})

	})

	Describe("Log", func() {
		It("writes to output log file", func() {
			rotatableSink.Log(lager.LogFormat{Message: "hello"})

			Expect(fakeTestWriterFactory.InvocationCount()).To(Equal(1))

			Expect(ReadLines(fileToWatch.Name())).To(ContainElement(MatchJSON(`{"timestamp":"some-timestamp","source":"","message":"hello","log_level":0,"data":null}`)))
		})

		It("should only open the file when it has been rotated", func() {
			rotatableSink.Log(lager.LogFormat{Message: "hello"})
			Expect(fakeTestWriterFactory.InvocationCount()).To(Equal(1))

			rotatableSink.Log(lager.LogFormat{Message: "hello"})
			Expect(fakeTestWriterFactory.InvocationCount()).To(Equal(1))

			Expect(ReadLines(fileToWatch.Name())).To(ContainElement(MatchJSON(`{"timestamp":"some-timestamp","source":"","message":"hello","log_level":0,"data":null}`)))
		})

		Context("when the file is rotated", func() {
			It("writes to output log file", func() {
				By("rotating the file")
				Expect(os.Rename(fileToWatchName, filepath.Join(os.TempDir(), "backup_sink_log"))).To(Succeed())
				rotatedFile, err := os.Create(fileToWatchName)
				Expect(err).NotTo(HaveOccurred())

				fakeTestWriterFactory.SetReturnWriter(rotatedFile)

				time.Sleep(2 * time.Second)
				rotatableSink.Log(lager.LogFormat{Message: "hello2"})

				Expect(ReadLines(fileToWatch.Name())).To(ContainElement(MatchJSON(`{"timestamp":"some-timestamp","source":"","message":"hello2","log_level":0,"data":null}`)))
			})

			Context("when unable to open the destination file that was rotated", func() {
				It("returns a sensible error", func() {
					By("rotating the file")
					fakeDestinationFileInfo.FileExistsReturns(true, nil)
					fakeDestinationFileInfo.FileInodeReturns(uint64(1), nil)
					fakeTestWriterFactory.SetReturnedError(errors.New("apple"))

					time.Sleep(2 * time.Second)
					Expect(len(fakeLogger.Logs())).To(BeNumerically(">", 0))
					Expect(fakeLogger.Logs()[0]).To(SatisfyAll(
						LogsWith(lager.ERROR, "test.register-rotated-file-sink"),
						HaveLogData(HaveKeyWithValue("error", "rotate file sink: create file writer: apple")),
					))
				})
			})

			Context("when unable to get the file inode of the destination file that was rotated", func() {
				BeforeEach(func() {
					fakeDestinationFileInfo.FileExistsReturns(true, nil)
					fakeDestinationFileInfo.FileInodeReturnsOnCall(1, 1, nil)
					fakeDestinationFileInfo.FileInodeReturns(1, errors.New("get file inode: watermelon"))
					fakeTestWriterFactory = NewTestWriterFactory(fileToWatch, nil)
					var err error
					rotatableSink, err = rotatablesink.NewRotatableSink(fileToWatchName, lager.DEBUG, fakeTestWriterFactory, fakeDestinationFileInfo, fakeLogger)
					Expect(err).ToNot(HaveOccurred())
				})

				It("returns a sensible error and does not update the file sink", func() {
					time.Sleep(2 * time.Second)
					Expect(len(fakeLogger.Logs())).To(BeNumerically(">", 0))
					Expect(fakeLogger.Logs()[0]).To(SatisfyAll(
						LogsWith(lager.ERROR, "test.register-rotated-file-sink"),
						HaveLogData(HaveKeyWithValue("error", "get file inode: watermelon")),
					))
					Expect(fakeTestWriterFactory.InvocationCount()).To(Equal(1))
				})
			})

			Context("when the destination file is deleted", func() {
				BeforeEach(func() {
					By("deleting the file")
					fakeDestinationFileInfo.FileExistsReturns(false, nil)
				})

				It("returns a sensible error", func() {
					fakeTestWriterFactory.SetReturnedError(errors.New("apple"))

					Eventually(func() int { return len(fakeLogger.Logs()) }, "5s").Should(BeNumerically(">", 0))
					Eventually(func() lager.LogFormat {
						fakeLoggerLogs := fakeLogger.Logs()
						return fakeLoggerLogs[len(fakeLoggerLogs)-1]
					}).Should(SatisfyAll(
						LogsWith(lager.ERROR, "test.register-moved-file-sink"),
						HaveLogData(HaveKeyWithValue("error", "rotate file sink: create file writer: apple")),
					))
				})

				Context("when unable to get the destination file inode", func() {
					BeforeEach(func() {
						fakeDestinationFileInfo.FileInodeReturns(0, errors.New("banana"))
					})

					It("returns a sensible error", func() {
						Eventually(func() int { return len(fakeLogger.Logs()) }, "5s").Should(BeNumerically(">", 0))
						Eventually(func() lager.LogFormat {
							fakeLoggerLogs := fakeLogger.Logs()
							return fakeLoggerLogs[len(fakeLoggerLogs)-1]
						}).Should(SatisfyAll(
							LogsWith(lager.ERROR, "test.register-moved-file-sink"),
							HaveLogData(HaveKeyWithValue("error", "get file inode: banana")),
						))
					})
				})
			})

			Context("when unable to check if the file being watched exists", func() {
				It("returns a sensible error", func() {
					By("failing on checking if file exists")
					fakeDestinationFileInfo.FileExistsReturns(false, errors.New("pineapple"))

					time.Sleep(5 * time.Second)
					Eventually(func() int {
						return len(fakeLogger.Logs())
					}).Should(BeNumerically(">", 0))
					Eventually(func() lager.LogFormat {
						fakeLoggerLogs := fakeLogger.Logs()
						return fakeLoggerLogs[len(fakeLoggerLogs)-1]
					}).Should(SatisfyAll(
						LogsWith(lager.ERROR, "test.stat-file"),
						HaveLogData(HaveKeyWithValue("error", "stat file: pineapple")),
					))
				})
			})
		})

	})

	Describe("FileWriterFactory", func() {
		It("should return a writer that can write to a file", func() {
			writer, err := rotatablesink.DefaultFileWriter(fileToWatch.Name())
			Expect(err).NotTo(HaveOccurred())

			writer.Write([]byte("hello world"))

			contents, err := ioutil.ReadFile(fileToWatch.Name())
			Expect(err).NotTo(HaveOccurred())
			Expect(string(contents)).To(Equal("hello world"))
		})
	})

	Describe("DestinationFileInfo", func() {
		var (
			defaultDestinationFileInfo rotatablesink.DestinationFileInfo
		)
		BeforeEach(func() {
			defaultDestinationFileInfo = rotatablesink.DefaultDestinationFileInfo{}
		})

		Describe("FileExists", func() {

			It("should return true when file exists", func() {
				fileExists, err := defaultDestinationFileInfo.FileExists(fileToWatchName)
				Expect(err).ToNot(HaveOccurred())
				Expect(fileExists).To(BeTrue())
			})

			Context("when the file does not exist", func() {
				It("returns false", func() {
					fileExists, err := defaultDestinationFileInfo.FileExists(fmt.Sprintf("%s_does_not_exist", fileToWatchName))
					Expect(err).ToNot(HaveOccurred())
					Expect(fileExists).ToNot(BeTrue())
				})
			})

			Context("when an invalid file is provided", func() {
				It("should return a sensible error", func() {
					_, err := defaultDestinationFileInfo.FileExists(filepath.Join(fileToWatchName, fileToWatchName))
					Expect(err).To(MatchError(MatchRegexp("stat file: .* not a directory")))
				})
			})
		})

		Describe("FileInode", func() {
			It("should return the file to watch inode", func() {
				inode, err := defaultDestinationFileInfo.FileInode(fileToWatchName)
				Expect(err).ToNot(HaveOccurred())
				Expect(inode).To(BeNumerically(">", 0))
			})

			Context("when an invalid file is provided", func() {
				It("should return a sensible error", func() {
					_, err := defaultDestinationFileInfo.FileInode(filepath.Join(fileToWatchName, fileToWatchName))
					Expect(err).To(MatchError(MatchRegexp("stat file: .* not a directory")))
				})
			})
		})
	})
})

type TestWriterFactory struct {
	invocationCount int
	returnWriter    io.Writer
	returnedError   error
	mutex           *sync.Mutex
}

func NewTestWriterFactory(w io.Writer, e error) *TestWriterFactory {
	twf := &TestWriterFactory{
		returnWriter:  w,
		returnedError: e,
		mutex:         new(sync.Mutex),
	}
	return twf
}

func (twf *TestWriterFactory) SetReturnWriter(w io.Writer) {
	twf.mutex.Lock()
	defer twf.mutex.Unlock()

	twf.returnWriter = w
}

func (twf *TestWriterFactory) SetReturnedError(e error) {
	twf.mutex.Lock()
	defer twf.mutex.Unlock()

	twf.returnedError = e
}

func (twf *TestWriterFactory) InvocationCount() int {
	twf.mutex.Lock()
	defer twf.mutex.Unlock()

	return twf.invocationCount
}

func (twf *TestWriterFactory) NewWriter(_ string) (io.Writer, error) {
	twf.mutex.Lock()
	defer twf.mutex.Unlock()

	twf.invocationCount++
	return twf.returnWriter, twf.returnedError
}

func ReadLines(filename string) []string {
	output := strings.Split(ReadOutput(filename), "\n")
	output = output[:len(output)-1]

	var outputs []string
	for _, o := range output {
		var outputMap map[string]interface{}
		err := json.Unmarshal([]byte(o), &outputMap)
		Expect(err).NotTo(HaveOccurred())

		outputMap["timestamp"] = "some-timestamp"
		outputJson, err := json.Marshal(outputMap)
		Expect(err).NotTo(HaveOccurred())

		outputs = append(outputs, string(outputJson))
	}

	return outputs
}

func ReadOutput(outputFile string) string {
	bytes, err := ioutil.ReadFile(outputFile)
	Expect(err).NotTo(HaveOccurred())
	if string(bytes) == "" {
		return "{}"
	}
	return string(bytes)
}