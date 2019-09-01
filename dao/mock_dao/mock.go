// Code generated by MockGen. DO NOT EDIT.
// Source: dao.go

// Package mock_dao is a generated GoMock package.
package mock_dao

import (
	gomock "github.com/golang/mock/gomock"
	model "metrics-collector/model"
	reflect "reflect"
)

// MockDao is a mock of Dao interface
type MockDao struct {
	ctrl     *gomock.Controller
	recorder *MockDaoMockRecorder
}

// MockDaoMockRecorder is the mock recorder for MockDao
type MockDaoMockRecorder struct {
	mock *MockDao
}

// NewMockDao creates a new mock instance
func NewMockDao(ctrl *gomock.Controller) *MockDao {
	mock := &MockDao{ctrl: ctrl}
	mock.recorder = &MockDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDao) EXPECT() *MockDaoMockRecorder {
	return m.recorder
}

// StoreMetric mocks base method
func (m *MockDao) StoreMetric(id string, metric model.Metric) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StoreMetric", id, metric)
}

// StoreMetric indicates an expected call of StoreMetric
func (mr *MockDaoMockRecorder) StoreMetric(id, metric interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreMetric", reflect.TypeOf((*MockDao)(nil).StoreMetric), id, metric)
}