// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Layr-Labs/eigensdk-go-master/services/avsregistry (interfaces: AvsRegistryService)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/avsregistry.go -package=mocks github.com/Layr-Labs/eigensdk-go-master/services/avsregistry AvsRegistryService
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	contractBLSOperatorStateRetriever "github.com/Layr-Labs/eigensdk-go-master/contracts/bindings/BLSOperatorStateRetriever"
	types "github.com/Layr-Labs/eigensdk-go-master/types"
	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	gomock "go.uber.org/mock/gomock"
)

// MockAvsRegistryService is a mock of AvsRegistryService interface.
type MockAvsRegistryService struct {
	ctrl     *gomock.Controller
	recorder *MockAvsRegistryServiceMockRecorder
}

// MockAvsRegistryServiceMockRecorder is the mock recorder for MockAvsRegistryService.
type MockAvsRegistryServiceMockRecorder struct {
	mock *MockAvsRegistryService
}

// NewMockAvsRegistryService creates a new mock instance.
func NewMockAvsRegistryService(ctrl *gomock.Controller) *MockAvsRegistryService {
	mock := &MockAvsRegistryService{ctrl: ctrl}
	mock.recorder = &MockAvsRegistryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAvsRegistryService) EXPECT() *MockAvsRegistryServiceMockRecorder {
	return m.recorder
}

// GetCheckSignaturesIndices mocks base method.
func (m *MockAvsRegistryService) GetCheckSignaturesIndices(arg0 *bind.CallOpts, arg1 uint32, arg2 []byte, arg3 [][32]byte) (contractBLSOperatorStateRetriever.BLSOperatorStateRetrieverCheckSignaturesIndices, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCheckSignaturesIndices", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(contractBLSOperatorStateRetriever.BLSOperatorStateRetrieverCheckSignaturesIndices)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCheckSignaturesIndices indicates an expected call of GetCheckSignaturesIndices.
func (mr *MockAvsRegistryServiceMockRecorder) GetCheckSignaturesIndices(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCheckSignaturesIndices", reflect.TypeOf((*MockAvsRegistryService)(nil).GetCheckSignaturesIndices), arg0, arg1, arg2, arg3)
}

// GetOperatorsAvsStateAtBlock mocks base method.
func (m *MockAvsRegistryService) GetOperatorsAvsStateAtBlock(arg0 context.Context, arg1 []byte, arg2 uint32) (map[[32]byte]types.OperatorAvsState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorsAvsStateAtBlock", arg0, arg1, arg2)
	ret0, _ := ret[0].(map[[32]byte]types.OperatorAvsState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperatorsAvsStateAtBlock indicates an expected call of GetOperatorsAvsStateAtBlock.
func (mr *MockAvsRegistryServiceMockRecorder) GetOperatorsAvsStateAtBlock(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorsAvsStateAtBlock", reflect.TypeOf((*MockAvsRegistryService)(nil).GetOperatorsAvsStateAtBlock), arg0, arg1, arg2)
}

// GetQuorumsAvsStateAtBlock mocks base method.
func (m *MockAvsRegistryService) GetQuorumsAvsStateAtBlock(arg0 context.Context, arg1 []byte, arg2 uint32) (map[byte]types.QuorumAvsState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuorumsAvsStateAtBlock", arg0, arg1, arg2)
	ret0, _ := ret[0].(map[byte]types.QuorumAvsState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQuorumsAvsStateAtBlock indicates an expected call of GetQuorumsAvsStateAtBlock.
func (mr *MockAvsRegistryServiceMockRecorder) GetQuorumsAvsStateAtBlock(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuorumsAvsStateAtBlock", reflect.TypeOf((*MockAvsRegistryService)(nil).GetQuorumsAvsStateAtBlock), arg0, arg1, arg2)
}
