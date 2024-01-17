// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Layr-Labs/eigensdk-go-master/chainio/clients (interfaces: ERC20ContractClient)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/erc20ContractClient.go -package=mocks github.com/Layr-Labs/eigensdk-go-master/chainio/clients ERC20ContractClient
//
// Package mocks is a generated GoMock package.
package mocks

import (
	big "math/big"
	reflect "reflect"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	gomock "go.uber.org/mock/gomock"
)

// MockERC20ContractClient is a mock of ERC20ContractClient interface.
type MockERC20ContractClient struct {
	ctrl     *gomock.Controller
	recorder *MockERC20ContractClientMockRecorder
}

// MockERC20ContractClientMockRecorder is the mock recorder for MockERC20ContractClient.
type MockERC20ContractClientMockRecorder struct {
	mock *MockERC20ContractClient
}

// NewMockERC20ContractClient creates a new mock instance.
func NewMockERC20ContractClient(ctrl *gomock.Controller) *MockERC20ContractClient {
	mock := &MockERC20ContractClient{ctrl: ctrl}
	mock.recorder = &MockERC20ContractClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockERC20ContractClient) EXPECT() *MockERC20ContractClientMockRecorder {
	return m.recorder
}

// Approve mocks base method.
func (m *MockERC20ContractClient) Approve(arg0 *bind.TransactOpts, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Approve", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Approve indicates an expected call of Approve.
func (mr *MockERC20ContractClientMockRecorder) Approve(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Approve", reflect.TypeOf((*MockERC20ContractClient)(nil).Approve), arg0, arg1, arg2)
}
