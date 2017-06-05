// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"lib/serial"
	"sync"
)

type OverwriteableFile struct {
	ReadStub        func(p []byte) (n int, err error)
	readMutex       sync.RWMutex
	readArgsForCall []struct {
		p []byte
	}
	readReturns struct {
		result1 int
		result2 error
	}
	readReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	WriteStub        func(p []byte) (n int, err error)
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		p []byte
	}
	writeReturns struct {
		result1 int
		result2 error
	}
	writeReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	SeekStub        func(offset int64, whence int) (int64, error)
	seekMutex       sync.RWMutex
	seekArgsForCall []struct {
		offset int64
		whence int
	}
	seekReturns struct {
		result1 int64
		result2 error
	}
	seekReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	TruncateStub        func(size int64) error
	truncateMutex       sync.RWMutex
	truncateArgsForCall []struct {
		size int64
	}
	truncateReturns struct {
		result1 error
	}
	truncateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *OverwriteableFile) Read(p []byte) (n int, err error) {
	var pCopy []byte
	if p != nil {
		pCopy = make([]byte, len(p))
		copy(pCopy, p)
	}
	fake.readMutex.Lock()
	ret, specificReturn := fake.readReturnsOnCall[len(fake.readArgsForCall)]
	fake.readArgsForCall = append(fake.readArgsForCall, struct {
		p []byte
	}{pCopy})
	fake.recordInvocation("Read", []interface{}{pCopy})
	fake.readMutex.Unlock()
	if fake.ReadStub != nil {
		return fake.ReadStub(p)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.readReturns.result1, fake.readReturns.result2
}

func (fake *OverwriteableFile) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *OverwriteableFile) ReadArgsForCall(i int) []byte {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return fake.readArgsForCall[i].p
}

func (fake *OverwriteableFile) ReadReturns(result1 int, result2 error) {
	fake.ReadStub = nil
	fake.readReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *OverwriteableFile) ReadReturnsOnCall(i int, result1 int, result2 error) {
	fake.ReadStub = nil
	if fake.readReturnsOnCall == nil {
		fake.readReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.readReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *OverwriteableFile) Write(p []byte) (n int, err error) {
	var pCopy []byte
	if p != nil {
		pCopy = make([]byte, len(p))
		copy(pCopy, p)
	}
	fake.writeMutex.Lock()
	ret, specificReturn := fake.writeReturnsOnCall[len(fake.writeArgsForCall)]
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		p []byte
	}{pCopy})
	fake.recordInvocation("Write", []interface{}{pCopy})
	fake.writeMutex.Unlock()
	if fake.WriteStub != nil {
		return fake.WriteStub(p)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.writeReturns.result1, fake.writeReturns.result2
}

func (fake *OverwriteableFile) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *OverwriteableFile) WriteArgsForCall(i int) []byte {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return fake.writeArgsForCall[i].p
}

func (fake *OverwriteableFile) WriteReturns(result1 int, result2 error) {
	fake.WriteStub = nil
	fake.writeReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *OverwriteableFile) WriteReturnsOnCall(i int, result1 int, result2 error) {
	fake.WriteStub = nil
	if fake.writeReturnsOnCall == nil {
		fake.writeReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.writeReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *OverwriteableFile) Seek(offset int64, whence int) (int64, error) {
	fake.seekMutex.Lock()
	ret, specificReturn := fake.seekReturnsOnCall[len(fake.seekArgsForCall)]
	fake.seekArgsForCall = append(fake.seekArgsForCall, struct {
		offset int64
		whence int
	}{offset, whence})
	fake.recordInvocation("Seek", []interface{}{offset, whence})
	fake.seekMutex.Unlock()
	if fake.SeekStub != nil {
		return fake.SeekStub(offset, whence)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.seekReturns.result1, fake.seekReturns.result2
}

func (fake *OverwriteableFile) SeekCallCount() int {
	fake.seekMutex.RLock()
	defer fake.seekMutex.RUnlock()
	return len(fake.seekArgsForCall)
}

func (fake *OverwriteableFile) SeekArgsForCall(i int) (int64, int) {
	fake.seekMutex.RLock()
	defer fake.seekMutex.RUnlock()
	return fake.seekArgsForCall[i].offset, fake.seekArgsForCall[i].whence
}

func (fake *OverwriteableFile) SeekReturns(result1 int64, result2 error) {
	fake.SeekStub = nil
	fake.seekReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *OverwriteableFile) SeekReturnsOnCall(i int, result1 int64, result2 error) {
	fake.SeekStub = nil
	if fake.seekReturnsOnCall == nil {
		fake.seekReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.seekReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *OverwriteableFile) Truncate(size int64) error {
	fake.truncateMutex.Lock()
	ret, specificReturn := fake.truncateReturnsOnCall[len(fake.truncateArgsForCall)]
	fake.truncateArgsForCall = append(fake.truncateArgsForCall, struct {
		size int64
	}{size})
	fake.recordInvocation("Truncate", []interface{}{size})
	fake.truncateMutex.Unlock()
	if fake.TruncateStub != nil {
		return fake.TruncateStub(size)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.truncateReturns.result1
}

func (fake *OverwriteableFile) TruncateCallCount() int {
	fake.truncateMutex.RLock()
	defer fake.truncateMutex.RUnlock()
	return len(fake.truncateArgsForCall)
}

func (fake *OverwriteableFile) TruncateArgsForCall(i int) int64 {
	fake.truncateMutex.RLock()
	defer fake.truncateMutex.RUnlock()
	return fake.truncateArgsForCall[i].size
}

func (fake *OverwriteableFile) TruncateReturns(result1 error) {
	fake.TruncateStub = nil
	fake.truncateReturns = struct {
		result1 error
	}{result1}
}

func (fake *OverwriteableFile) TruncateReturnsOnCall(i int, result1 error) {
	fake.TruncateStub = nil
	if fake.truncateReturnsOnCall == nil {
		fake.truncateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.truncateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *OverwriteableFile) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	fake.seekMutex.RLock()
	defer fake.seekMutex.RUnlock()
	fake.truncateMutex.RLock()
	defer fake.truncateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *OverwriteableFile) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ serial.OverwriteableFile = new(OverwriteableFile)
