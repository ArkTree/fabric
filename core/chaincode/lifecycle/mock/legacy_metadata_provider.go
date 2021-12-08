// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/ArkTree/fabric/common/chaincode"
)

type LegacyMetadataProvider struct {
	MetadataStub        func(string, string, ...string) *chaincode.Metadata
	metadataMutex       sync.RWMutex
	metadataArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 []string
	}
	metadataReturns struct {
		result1 *chaincode.Metadata
	}
	metadataReturnsOnCall map[int]struct {
		result1 *chaincode.Metadata
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *LegacyMetadataProvider) Metadata(arg1 string, arg2 string, arg3 ...string) *chaincode.Metadata {
	fake.metadataMutex.Lock()
	ret, specificReturn := fake.metadataReturnsOnCall[len(fake.metadataArgsForCall)]
	fake.metadataArgsForCall = append(fake.metadataArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 []string
	}{arg1, arg2, arg3})
	fake.recordInvocation("Metadata", []interface{}{arg1, arg2, arg3})
	fake.metadataMutex.Unlock()
	if fake.MetadataStub != nil {
		return fake.MetadataStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.metadataReturns
	return fakeReturns.result1
}

func (fake *LegacyMetadataProvider) MetadataCallCount() int {
	fake.metadataMutex.RLock()
	defer fake.metadataMutex.RUnlock()
	return len(fake.metadataArgsForCall)
}

func (fake *LegacyMetadataProvider) MetadataCalls(stub func(string, string, ...string) *chaincode.Metadata) {
	fake.metadataMutex.Lock()
	defer fake.metadataMutex.Unlock()
	fake.MetadataStub = stub
}

func (fake *LegacyMetadataProvider) MetadataArgsForCall(i int) (string, string, []string) {
	fake.metadataMutex.RLock()
	defer fake.metadataMutex.RUnlock()
	argsForCall := fake.metadataArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *LegacyMetadataProvider) MetadataReturns(result1 *chaincode.Metadata) {
	fake.metadataMutex.Lock()
	defer fake.metadataMutex.Unlock()
	fake.MetadataStub = nil
	fake.metadataReturns = struct {
		result1 *chaincode.Metadata
	}{result1}
}

func (fake *LegacyMetadataProvider) MetadataReturnsOnCall(i int, result1 *chaincode.Metadata) {
	fake.metadataMutex.Lock()
	defer fake.metadataMutex.Unlock()
	fake.MetadataStub = nil
	if fake.metadataReturnsOnCall == nil {
		fake.metadataReturnsOnCall = make(map[int]struct {
			result1 *chaincode.Metadata
		})
	}
	fake.metadataReturnsOnCall[i] = struct {
		result1 *chaincode.Metadata
	}{result1}
}

func (fake *LegacyMetadataProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.metadataMutex.RLock()
	defer fake.metadataMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *LegacyMetadataProvider) recordInvocation(key string, args []interface{}) {
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
