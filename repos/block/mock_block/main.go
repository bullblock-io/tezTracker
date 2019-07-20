// Code generated by MockGen. DO NOT EDIT.
// Source: ./block.go

// Package mock_block is a generated GoMock package.
package mock_block

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

// Last mocks base method
func (m *MockRepo) Last() (models.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Last")
	ret0, _ := ret[0].(models.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Last indicates an expected call of Last
func (mr *MockRepoMockRecorder) Last() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Last", reflect.TypeOf((*MockRepo)(nil).Last))
}

// List mocks base method
func (m *MockRepo) List(limit, since uint64) ([]models.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", limit, since)
	ret0, _ := ret[0].([]models.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockRepoMockRecorder) List(limit, since interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepo)(nil).List), limit, since)
}

// Find mocks base method
func (m *MockRepo) Find(filter models.Block) (bool, models.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", filter)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(models.Block)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Find indicates an expected call of Find
func (mr *MockRepoMockRecorder) Find(filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockRepo)(nil).Find), filter)
}
