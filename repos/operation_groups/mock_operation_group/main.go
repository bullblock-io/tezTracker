// Code generated by MockGen. DO NOT EDIT.
// Source: ./operation.go

// Package mock_operation_groups is a generated GoMock package.
package mock_operation_groups

import (
	models "github.com/bullblock-io/tezTracker/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepo is a mock of Repo interface
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// GetGroupsFor mocks base method
func (m *MockRepo) GetGroupsFor(block models.Block) ([]*models.OperationGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupsFor", block)
	ret0, _ := ret[0].([]*models.OperationGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupsFor indicates an expected call of GetGroupsFor
func (mr *MockRepoMockRecorder) GetGroupsFor(block interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupsFor", reflect.TypeOf((*MockRepo)(nil).GetGroupsFor), block)
}
