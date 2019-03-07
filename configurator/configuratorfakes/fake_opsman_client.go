// Code generated by counterfeiter. DO NOT EDIT.
package configuratorfakes

import (
	"sync"

	"github.com/starkandwayne/om-configurator/configurator"
)

type FakeOpsmanClient struct {
	ApplyChangesStub        func() error
	applyChangesMutex       sync.RWMutex
	applyChangesArgsForCall []struct {
	}
	applyChangesReturns struct {
		result1 error
	}
	applyChangesReturnsOnCall map[int]struct {
		result1 error
	}
	ConfigureAuthenticationStub        func() error
	configureAuthenticationMutex       sync.RWMutex
	configureAuthenticationArgsForCall []struct {
	}
	configureAuthenticationReturns struct {
		result1 error
	}
	configureAuthenticationReturnsOnCall map[int]struct {
		result1 error
	}
	ConfigureProductStub        func([]byte) error
	configureProductMutex       sync.RWMutex
	configureProductArgsForCall []struct {
		arg1 []byte
	}
	configureProductReturns struct {
		result1 error
	}
	configureProductReturnsOnCall map[int]struct {
		result1 error
	}
	DownloadProductStub        func(configurator.DownloadProductArgs) error
	downloadProductMutex       sync.RWMutex
	downloadProductArgsForCall []struct {
		arg1 configurator.DownloadProductArgs
	}
	downloadProductReturns struct {
		result1 error
	}
	downloadProductReturnsOnCall map[int]struct {
		result1 error
	}
	StageProductStub        func(configurator.StageProductArgs) error
	stageProductMutex       sync.RWMutex
	stageProductArgsForCall []struct {
		arg1 configurator.StageProductArgs
	}
	stageProductReturns struct {
		result1 error
	}
	stageProductReturnsOnCall map[int]struct {
		result1 error
	}
	UploadProductStub        func(string) error
	uploadProductMutex       sync.RWMutex
	uploadProductArgsForCall []struct {
		arg1 string
	}
	uploadProductReturns struct {
		result1 error
	}
	uploadProductReturnsOnCall map[int]struct {
		result1 error
	}
	UploadStemcellStub        func(string) error
	uploadStemcellMutex       sync.RWMutex
	uploadStemcellArgsForCall []struct {
		arg1 string
	}
	uploadStemcellReturns struct {
		result1 error
	}
	uploadStemcellReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOpsmanClient) ApplyChanges() error {
	fake.applyChangesMutex.Lock()
	ret, specificReturn := fake.applyChangesReturnsOnCall[len(fake.applyChangesArgsForCall)]
	fake.applyChangesArgsForCall = append(fake.applyChangesArgsForCall, struct {
	}{})
	fake.recordInvocation("ApplyChanges", []interface{}{})
	fake.applyChangesMutex.Unlock()
	if fake.ApplyChangesStub != nil {
		return fake.ApplyChangesStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.applyChangesReturns
	return fakeReturns.result1
}

func (fake *FakeOpsmanClient) ApplyChangesCallCount() int {
	fake.applyChangesMutex.RLock()
	defer fake.applyChangesMutex.RUnlock()
	return len(fake.applyChangesArgsForCall)
}

func (fake *FakeOpsmanClient) ApplyChangesCalls(stub func() error) {
	fake.applyChangesMutex.Lock()
	defer fake.applyChangesMutex.Unlock()
	fake.ApplyChangesStub = stub
}

func (fake *FakeOpsmanClient) ApplyChangesReturns(result1 error) {
	fake.applyChangesMutex.Lock()
	defer fake.applyChangesMutex.Unlock()
	fake.ApplyChangesStub = nil
	fake.applyChangesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOpsmanClient) ApplyChangesReturnsOnCall(i int, result1 error) {
	fake.applyChangesMutex.Lock()
	defer fake.applyChangesMutex.Unlock()
	fake.ApplyChangesStub = nil
	if fake.applyChangesReturnsOnCall == nil {
		fake.applyChangesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.applyChangesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOpsmanClient) ConfigureAuthentication() error {
	fake.configureAuthenticationMutex.Lock()
	ret, specificReturn := fake.configureAuthenticationReturnsOnCall[len(fake.configureAuthenticationArgsForCall)]
	fake.configureAuthenticationArgsForCall = append(fake.configureAuthenticationArgsForCall, struct {
	}{})
	fake.recordInvocation("ConfigureAuthentication", []interface{}{})
	fake.configureAuthenticationMutex.Unlock()
	if fake.ConfigureAuthenticationStub != nil {
		return fake.ConfigureAuthenticationStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.configureAuthenticationReturns
	return fakeReturns.result1
}

func (fake *FakeOpsmanClient) ConfigureAuthenticationCallCount() int {
	fake.configureAuthenticationMutex.RLock()
	defer fake.configureAuthenticationMutex.RUnlock()
	return len(fake.configureAuthenticationArgsForCall)
}

func (fake *FakeOpsmanClient) ConfigureAuthenticationCalls(stub func() error) {
	fake.configureAuthenticationMutex.Lock()
	defer fake.configureAuthenticationMutex.Unlock()
	fake.ConfigureAuthenticationStub = stub
}

func (fake *FakeOpsmanClient) ConfigureAuthenticationReturns(result1 error) {
	fake.configureAuthenticationMutex.Lock()
	defer fake.configureAuthenticationMutex.Unlock()
	fake.ConfigureAuthenticationStub = nil
	fake.configureAuthenticationReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOpsmanClient) ConfigureAuthenticationReturnsOnCall(i int, result1 error) {
	fake.configureAuthenticationMutex.Lock()
	defer fake.configureAuthenticationMutex.Unlock()
	fake.ConfigureAuthenticationStub = nil
	if fake.configureAuthenticationReturnsOnCall == nil {
		fake.configureAuthenticationReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.configureAuthenticationReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOpsmanClient) ConfigureProduct(arg1 []byte) error {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.configureProductMutex.Lock()
	ret, specificReturn := fake.configureProductReturnsOnCall[len(fake.configureProductArgsForCall)]
	fake.configureProductArgsForCall = append(fake.configureProductArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	fake.recordInvocation("ConfigureProduct", []interface{}{arg1Copy})
	fake.configureProductMutex.Unlock()
	if fake.ConfigureProductStub != nil {
		return fake.ConfigureProductStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.configureProductReturns
	return fakeReturns.result1
}

func (fake *FakeOpsmanClient) ConfigureProductCallCount() int {
	fake.configureProductMutex.RLock()
	defer fake.configureProductMutex.RUnlock()
	return len(fake.configureProductArgsForCall)
}

func (fake *FakeOpsmanClient) ConfigureProductCalls(stub func([]byte) error) {
	fake.configureProductMutex.Lock()
	defer fake.configureProductMutex.Unlock()
	fake.ConfigureProductStub = stub
}

func (fake *FakeOpsmanClient) ConfigureProductArgsForCall(i int) []byte {
	fake.configureProductMutex.RLock()
	defer fake.configureProductMutex.RUnlock()
	argsForCall := fake.configureProductArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeOpsmanClient) ConfigureProductReturns(result1 error) {
	fake.configureProductMutex.Lock()
	defer fake.configureProductMutex.Unlock()
	fake.ConfigureProductStub = nil
	fake.configureProductReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOpsmanClient) ConfigureProductReturnsOnCall(i int, result1 error) {
	fake.configureProductMutex.Lock()
	defer fake.configureProductMutex.Unlock()
	fake.ConfigureProductStub = nil
	if fake.configureProductReturnsOnCall == nil {
		fake.configureProductReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.configureProductReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOpsmanClient) DownloadProduct(arg1 configurator.DownloadProductArgs) error {
	fake.downloadProductMutex.Lock()
	ret, specificReturn := fake.downloadProductReturnsOnCall[len(fake.downloadProductArgsForCall)]
	fake.downloadProductArgsForCall = append(fake.downloadProductArgsForCall, struct {
		arg1 configurator.DownloadProductArgs
	}{arg1})
	fake.recordInvocation("DownloadProduct", []interface{}{arg1})
	fake.downloadProductMutex.Unlock()
	if fake.DownloadProductStub != nil {
		return fake.DownloadProductStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.downloadProductReturns
	return fakeReturns.result1
}

func (fake *FakeOpsmanClient) DownloadProductCallCount() int {
	fake.downloadProductMutex.RLock()
	defer fake.downloadProductMutex.RUnlock()
	return len(fake.downloadProductArgsForCall)
}

func (fake *FakeOpsmanClient) DownloadProductCalls(stub func(configurator.DownloadProductArgs) error) {
	fake.downloadProductMutex.Lock()
	defer fake.downloadProductMutex.Unlock()
	fake.DownloadProductStub = stub
}

func (fake *FakeOpsmanClient) DownloadProductArgsForCall(i int) configurator.DownloadProductArgs {
	fake.downloadProductMutex.RLock()
	defer fake.downloadProductMutex.RUnlock()
	argsForCall := fake.downloadProductArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeOpsmanClient) DownloadProductReturns(result1 error) {
	fake.downloadProductMutex.Lock()
	defer fake.downloadProductMutex.Unlock()
	fake.DownloadProductStub = nil
	fake.downloadProductReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOpsmanClient) DownloadProductReturnsOnCall(i int, result1 error) {
	fake.downloadProductMutex.Lock()
	defer fake.downloadProductMutex.Unlock()
	fake.DownloadProductStub = nil
	if fake.downloadProductReturnsOnCall == nil {
		fake.downloadProductReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.downloadProductReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOpsmanClient) StageProduct(arg1 configurator.StageProductArgs) error {
	fake.stageProductMutex.Lock()
	ret, specificReturn := fake.stageProductReturnsOnCall[len(fake.stageProductArgsForCall)]
	fake.stageProductArgsForCall = append(fake.stageProductArgsForCall, struct {
		arg1 configurator.StageProductArgs
	}{arg1})
	fake.recordInvocation("StageProduct", []interface{}{arg1})
	fake.stageProductMutex.Unlock()
	if fake.StageProductStub != nil {
		return fake.StageProductStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.stageProductReturns
	return fakeReturns.result1
}

func (fake *FakeOpsmanClient) StageProductCallCount() int {
	fake.stageProductMutex.RLock()
	defer fake.stageProductMutex.RUnlock()
	return len(fake.stageProductArgsForCall)
}

func (fake *FakeOpsmanClient) StageProductCalls(stub func(configurator.StageProductArgs) error) {
	fake.stageProductMutex.Lock()
	defer fake.stageProductMutex.Unlock()
	fake.StageProductStub = stub
}

func (fake *FakeOpsmanClient) StageProductArgsForCall(i int) configurator.StageProductArgs {
	fake.stageProductMutex.RLock()
	defer fake.stageProductMutex.RUnlock()
	argsForCall := fake.stageProductArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeOpsmanClient) StageProductReturns(result1 error) {
	fake.stageProductMutex.Lock()
	defer fake.stageProductMutex.Unlock()
	fake.StageProductStub = nil
	fake.stageProductReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOpsmanClient) StageProductReturnsOnCall(i int, result1 error) {
	fake.stageProductMutex.Lock()
	defer fake.stageProductMutex.Unlock()
	fake.StageProductStub = nil
	if fake.stageProductReturnsOnCall == nil {
		fake.stageProductReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.stageProductReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOpsmanClient) UploadProduct(arg1 string) error {
	fake.uploadProductMutex.Lock()
	ret, specificReturn := fake.uploadProductReturnsOnCall[len(fake.uploadProductArgsForCall)]
	fake.uploadProductArgsForCall = append(fake.uploadProductArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("UploadProduct", []interface{}{arg1})
	fake.uploadProductMutex.Unlock()
	if fake.UploadProductStub != nil {
		return fake.UploadProductStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.uploadProductReturns
	return fakeReturns.result1
}

func (fake *FakeOpsmanClient) UploadProductCallCount() int {
	fake.uploadProductMutex.RLock()
	defer fake.uploadProductMutex.RUnlock()
	return len(fake.uploadProductArgsForCall)
}

func (fake *FakeOpsmanClient) UploadProductCalls(stub func(string) error) {
	fake.uploadProductMutex.Lock()
	defer fake.uploadProductMutex.Unlock()
	fake.UploadProductStub = stub
}

func (fake *FakeOpsmanClient) UploadProductArgsForCall(i int) string {
	fake.uploadProductMutex.RLock()
	defer fake.uploadProductMutex.RUnlock()
	argsForCall := fake.uploadProductArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeOpsmanClient) UploadProductReturns(result1 error) {
	fake.uploadProductMutex.Lock()
	defer fake.uploadProductMutex.Unlock()
	fake.UploadProductStub = nil
	fake.uploadProductReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOpsmanClient) UploadProductReturnsOnCall(i int, result1 error) {
	fake.uploadProductMutex.Lock()
	defer fake.uploadProductMutex.Unlock()
	fake.UploadProductStub = nil
	if fake.uploadProductReturnsOnCall == nil {
		fake.uploadProductReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.uploadProductReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOpsmanClient) UploadStemcell(arg1 string) error {
	fake.uploadStemcellMutex.Lock()
	ret, specificReturn := fake.uploadStemcellReturnsOnCall[len(fake.uploadStemcellArgsForCall)]
	fake.uploadStemcellArgsForCall = append(fake.uploadStemcellArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("UploadStemcell", []interface{}{arg1})
	fake.uploadStemcellMutex.Unlock()
	if fake.UploadStemcellStub != nil {
		return fake.UploadStemcellStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.uploadStemcellReturns
	return fakeReturns.result1
}

func (fake *FakeOpsmanClient) UploadStemcellCallCount() int {
	fake.uploadStemcellMutex.RLock()
	defer fake.uploadStemcellMutex.RUnlock()
	return len(fake.uploadStemcellArgsForCall)
}

func (fake *FakeOpsmanClient) UploadStemcellCalls(stub func(string) error) {
	fake.uploadStemcellMutex.Lock()
	defer fake.uploadStemcellMutex.Unlock()
	fake.UploadStemcellStub = stub
}

func (fake *FakeOpsmanClient) UploadStemcellArgsForCall(i int) string {
	fake.uploadStemcellMutex.RLock()
	defer fake.uploadStemcellMutex.RUnlock()
	argsForCall := fake.uploadStemcellArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeOpsmanClient) UploadStemcellReturns(result1 error) {
	fake.uploadStemcellMutex.Lock()
	defer fake.uploadStemcellMutex.Unlock()
	fake.UploadStemcellStub = nil
	fake.uploadStemcellReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOpsmanClient) UploadStemcellReturnsOnCall(i int, result1 error) {
	fake.uploadStemcellMutex.Lock()
	defer fake.uploadStemcellMutex.Unlock()
	fake.UploadStemcellStub = nil
	if fake.uploadStemcellReturnsOnCall == nil {
		fake.uploadStemcellReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.uploadStemcellReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOpsmanClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.applyChangesMutex.RLock()
	defer fake.applyChangesMutex.RUnlock()
	fake.configureAuthenticationMutex.RLock()
	defer fake.configureAuthenticationMutex.RUnlock()
	fake.configureProductMutex.RLock()
	defer fake.configureProductMutex.RUnlock()
	fake.downloadProductMutex.RLock()
	defer fake.downloadProductMutex.RUnlock()
	fake.stageProductMutex.RLock()
	defer fake.stageProductMutex.RUnlock()
	fake.uploadProductMutex.RLock()
	defer fake.uploadProductMutex.RUnlock()
	fake.uploadStemcellMutex.RLock()
	defer fake.uploadStemcellMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOpsmanClient) recordInvocation(key string, args []interface{}) {
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

var _ configurator.OpsmanClient = new(FakeOpsmanClient)
