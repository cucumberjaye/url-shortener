// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cucumberjaye/url-shortener/internal/app/service (interfaces: URLRepository)

// Package mock_service is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/cucumberjaye/url-shortener/models"
	gomock "github.com/golang/mock/gomock"
)

// MockURLRepository is a mock of URLRepository interface.
type MockURLRepository struct {
	ctrl     *gomock.Controller
	recorder *MockURLRepositoryMockRecorder
}

// MockURLRepositoryMockRecorder is the mock recorder for MockURLRepository.
type MockURLRepositoryMockRecorder struct {
	mock *MockURLRepository
}

// NewMockURLRepository creates a new mock instance.
func NewMockURLRepository(ctrl *gomock.Controller) *MockURLRepository {
	mock := &MockURLRepository{ctrl: ctrl}
	mock.recorder = &MockURLRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockURLRepository) EXPECT() *MockURLRepositoryMockRecorder {
	return m.recorder
}

// BatchSetURL mocks base method.
func (m *MockURLRepository) BatchSetURL(arg0 []models.BatchInputJSON, arg1 []string, arg2 string) ([]models.BatchInputJSON, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchSetURL", arg0, arg1, arg2)
	ret0, _ := ret[0].([]models.BatchInputJSON)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchSetURL indicates an expected call of BatchSetURL.
func (mr *MockURLRepositoryMockRecorder) BatchSetURL(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchSetURL", reflect.TypeOf((*MockURLRepository)(nil).BatchSetURL), arg0, arg1, arg2)
}

// CheckStorage mocks base method.
func (m *MockURLRepository) CheckStorage() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckStorage")
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckStorage indicates an expected call of CheckStorage.
func (mr *MockURLRepositoryMockRecorder) CheckStorage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckStorage", reflect.TypeOf((*MockURLRepository)(nil).CheckStorage))
}

// GetAllUserURL mocks base method.
func (m *MockURLRepository) GetAllUserURL(arg0 string) ([]models.URLs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUserURL", arg0)
	ret0, _ := ret[0].([]models.URLs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUserURL indicates an expected call of GetAllUserURL.
func (mr *MockURLRepositoryMockRecorder) GetAllUserURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUserURL", reflect.TypeOf((*MockURLRepository)(nil).GetAllUserURL), arg0)
}

// GetURL mocks base method.
func (m *MockURLRepository) GetURL(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetURL", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetURL indicates an expected call of GetURL.
func (mr *MockURLRepositoryMockRecorder) GetURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetURL", reflect.TypeOf((*MockURLRepository)(nil).GetURL), arg0)
}

// GetURLCount mocks base method.
func (m *MockURLRepository) GetURLCount() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetURLCount")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetURLCount indicates an expected call of GetURLCount.
func (mr *MockURLRepositoryMockRecorder) GetURLCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetURLCount", reflect.TypeOf((*MockURLRepository)(nil).GetURLCount))
}

// SetURL mocks base method.
func (m *MockURLRepository) SetURL(arg0, arg1, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetURL", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetURL indicates an expected call of SetURL.
func (mr *MockURLRepositoryMockRecorder) SetURL(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetURL", reflect.TypeOf((*MockURLRepository)(nil).SetURL), arg0, arg1, arg2)
}