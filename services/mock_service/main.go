// Code generated by MockGen. DO NOT EDIT.
// Source: ./main.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	account "github.com/bullblock-io/tezTracker/repos/account"
	block "github.com/bullblock-io/tezTracker/repos/block"
	operation "github.com/bullblock-io/tezTracker/repos/operation"
	operation_groups "github.com/bullblock-io/tezTracker/repos/operation_groups"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockProvider is a mock of Provider interface
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// GetBlock mocks base method
func (m *MockProvider) GetBlock() block.Repo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlock")
	ret0, _ := ret[0].(block.Repo)
	return ret0
}

// GetBlock indicates an expected call of GetBlock
func (mr *MockProviderMockRecorder) GetBlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*MockProvider)(nil).GetBlock))
}

// GetOperationGroup mocks base method
func (m *MockProvider) GetOperationGroup() operation_groups.Repo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperationGroup")
	ret0, _ := ret[0].(operation_groups.Repo)
	return ret0
}

// GetOperationGroup indicates an expected call of GetOperationGroup
func (mr *MockProviderMockRecorder) GetOperationGroup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperationGroup", reflect.TypeOf((*MockProvider)(nil).GetOperationGroup))
}

// GetOperation mocks base method
func (m *MockProvider) GetOperation() operation.Repo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperation")
	ret0, _ := ret[0].(operation.Repo)
	return ret0
}

// GetOperation indicates an expected call of GetOperation
func (mr *MockProviderMockRecorder) GetOperation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperation", reflect.TypeOf((*MockProvider)(nil).GetOperation))
}

// GetAccount mocks base method
func (m *MockProvider) GetAccount() account.Repo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount")
	ret0, _ := ret[0].(account.Repo)
	return ret0
}

// GetAccount indicates an expected call of GetAccount
func (mr *MockProviderMockRecorder) GetAccount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockProvider)(nil).GetAccount))
}
