// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/policy-server/store"
)

type PolicyStore struct {
	ByGuidsStub        func([]string, []string, bool) ([]store.Policy, error)
	byGuidsMutex       sync.RWMutex
	byGuidsArgsForCall []struct {
		arg1 []string
		arg2 []string
		arg3 bool
	}
	byGuidsReturns struct {
		result1 []store.Policy
		result2 error
	}
	byGuidsReturnsOnCall map[int]struct {
		result1 []store.Policy
		result2 error
	}
	CreateStub        func([]store.Policy) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 []store.Policy
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func([]store.Policy) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 []store.Policy
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PolicyStore) ByGuids(arg1 []string, arg2 []string, arg3 bool) ([]store.Policy, error) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.byGuidsMutex.Lock()
	ret, specificReturn := fake.byGuidsReturnsOnCall[len(fake.byGuidsArgsForCall)]
	fake.byGuidsArgsForCall = append(fake.byGuidsArgsForCall, struct {
		arg1 []string
		arg2 []string
		arg3 bool
	}{arg1Copy, arg2Copy, arg3})
	stub := fake.ByGuidsStub
	fakeReturns := fake.byGuidsReturns
	fake.recordInvocation("ByGuids", []interface{}{arg1Copy, arg2Copy, arg3})
	fake.byGuidsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *PolicyStore) ByGuidsCallCount() int {
	fake.byGuidsMutex.RLock()
	defer fake.byGuidsMutex.RUnlock()
	return len(fake.byGuidsArgsForCall)
}

func (fake *PolicyStore) ByGuidsCalls(stub func([]string, []string, bool) ([]store.Policy, error)) {
	fake.byGuidsMutex.Lock()
	defer fake.byGuidsMutex.Unlock()
	fake.ByGuidsStub = stub
}

func (fake *PolicyStore) ByGuidsArgsForCall(i int) ([]string, []string, bool) {
	fake.byGuidsMutex.RLock()
	defer fake.byGuidsMutex.RUnlock()
	argsForCall := fake.byGuidsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *PolicyStore) ByGuidsReturns(result1 []store.Policy, result2 error) {
	fake.byGuidsMutex.Lock()
	defer fake.byGuidsMutex.Unlock()
	fake.ByGuidsStub = nil
	fake.byGuidsReturns = struct {
		result1 []store.Policy
		result2 error
	}{result1, result2}
}

func (fake *PolicyStore) ByGuidsReturnsOnCall(i int, result1 []store.Policy, result2 error) {
	fake.byGuidsMutex.Lock()
	defer fake.byGuidsMutex.Unlock()
	fake.ByGuidsStub = nil
	if fake.byGuidsReturnsOnCall == nil {
		fake.byGuidsReturnsOnCall = make(map[int]struct {
			result1 []store.Policy
			result2 error
		})
	}
	fake.byGuidsReturnsOnCall[i] = struct {
		result1 []store.Policy
		result2 error
	}{result1, result2}
}

func (fake *PolicyStore) Create(arg1 []store.Policy) error {
	var arg1Copy []store.Policy
	if arg1 != nil {
		arg1Copy = make([]store.Policy, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 []store.Policy
	}{arg1Copy})
	stub := fake.CreateStub
	fakeReturns := fake.createReturns
	fake.recordInvocation("Create", []interface{}{arg1Copy})
	fake.createMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *PolicyStore) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *PolicyStore) CreateCalls(stub func([]store.Policy) error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *PolicyStore) CreateArgsForCall(i int) []store.Policy {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1
}

func (fake *PolicyStore) CreateReturns(result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *PolicyStore) CreateReturnsOnCall(i int, result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *PolicyStore) Delete(arg1 []store.Policy) error {
	var arg1Copy []store.Policy
	if arg1 != nil {
		arg1Copy = make([]store.Policy, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 []store.Policy
	}{arg1Copy})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1Copy})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *PolicyStore) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *PolicyStore) DeleteCalls(stub func([]store.Policy) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *PolicyStore) DeleteArgsForCall(i int) []store.Policy {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *PolicyStore) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *PolicyStore) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *PolicyStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.byGuidsMutex.RLock()
	defer fake.byGuidsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *PolicyStore) recordInvocation(key string, args []interface{}) {
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