// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/ArkTree/fabric/common/ledger"
)

type LedgerProvider struct {
	LedgerStub        func(string) (ledger.Ledger, error)
	ledgerMutex       sync.RWMutex
	ledgerArgsForCall []struct {
		arg1 string
	}
	ledgerReturns struct {
		result1 ledger.Ledger
		result2 error
	}
	ledgerReturnsOnCall map[int]struct {
		result1 ledger.Ledger
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *LedgerProvider) Ledger(arg1 string) (ledger.Ledger, error) {
	fake.ledgerMutex.Lock()
	ret, specificReturn := fake.ledgerReturnsOnCall[len(fake.ledgerArgsForCall)]
	fake.ledgerArgsForCall = append(fake.ledgerArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.LedgerStub
	fakeReturns := fake.ledgerReturns
	fake.recordInvocation("Ledger", []interface{}{arg1})
	fake.ledgerMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *LedgerProvider) LedgerCallCount() int {
	fake.ledgerMutex.RLock()
	defer fake.ledgerMutex.RUnlock()
	return len(fake.ledgerArgsForCall)
}

func (fake *LedgerProvider) LedgerCalls(stub func(string) (ledger.Ledger, error)) {
	fake.ledgerMutex.Lock()
	defer fake.ledgerMutex.Unlock()
	fake.LedgerStub = stub
}

func (fake *LedgerProvider) LedgerArgsForCall(i int) string {
	fake.ledgerMutex.RLock()
	defer fake.ledgerMutex.RUnlock()
	argsForCall := fake.ledgerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *LedgerProvider) LedgerReturns(result1 ledger.Ledger, result2 error) {
	fake.ledgerMutex.Lock()
	defer fake.ledgerMutex.Unlock()
	fake.LedgerStub = nil
	fake.ledgerReturns = struct {
		result1 ledger.Ledger
		result2 error
	}{result1, result2}
}

func (fake *LedgerProvider) LedgerReturnsOnCall(i int, result1 ledger.Ledger, result2 error) {
	fake.ledgerMutex.Lock()
	defer fake.ledgerMutex.Unlock()
	fake.LedgerStub = nil
	if fake.ledgerReturnsOnCall == nil {
		fake.ledgerReturnsOnCall = make(map[int]struct {
			result1 ledger.Ledger
			result2 error
		})
	}
	fake.ledgerReturnsOnCall[i] = struct {
		result1 ledger.Ledger
		result2 error
	}{result1, result2}
}

func (fake *LedgerProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.ledgerMutex.RLock()
	defer fake.ledgerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *LedgerProvider) recordInvocation(key string, args []interface{}) {
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
