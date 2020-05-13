// Code generated by MockGen. DO NOT EDIT.
// Source: UtsuruConcept/models (interfaces: UtsuruDataStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	models "UtsuruConcept/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUtsuruDataStore is a mock of UtsuruDataStore interface
type MockUtsuruDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockUtsuruDataStoreMockRecorder
}

// MockUtsuruDataStoreMockRecorder is the mock recorder for MockUtsuruDataStore
type MockUtsuruDataStoreMockRecorder struct {
	mock *MockUtsuruDataStore
}

// NewMockUtsuruDataStore creates a new mock instance
func NewMockUtsuruDataStore(ctrl *gomock.Controller) *MockUtsuruDataStore {
	mock := &MockUtsuruDataStore{ctrl: ctrl}
	mock.recorder = &MockUtsuruDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUtsuruDataStore) EXPECT() *MockUtsuruDataStoreMockRecorder {
	return m.recorder
}

// GetSimilarImages mocks base method
func (m *MockUtsuruDataStore) GetSimilarImages(arg0 uint64) (*[]models.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSimilarImages", arg0)
	ret0, _ := ret[0].(*[]models.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSimilarImages indicates an expected call of GetSimilarImages
func (mr *MockUtsuruDataStoreMockRecorder) GetSimilarImages(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSimilarImages", reflect.TypeOf((*MockUtsuruDataStore)(nil).GetSimilarImages), arg0)
}

// GetUserByEmail mocks base method
func (m *MockUtsuruDataStore) GetUserByEmail(arg0 string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail
func (mr *MockUtsuruDataStoreMockRecorder) GetUserByEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockUtsuruDataStore)(nil).GetUserByEmail), arg0)
}

// GetUserById mocks base method
func (m *MockUtsuruDataStore) GetUserById(arg0 string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", arg0)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById
func (mr *MockUtsuruDataStoreMockRecorder) GetUserById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUtsuruDataStore)(nil).GetUserById), arg0)
}

// InsertOrUpdateUser mocks base method
func (m *MockUtsuruDataStore) InsertOrUpdateUser(arg0 *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOrUpdateUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertOrUpdateUser indicates an expected call of InsertOrUpdateUser
func (mr *MockUtsuruDataStoreMockRecorder) InsertOrUpdateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOrUpdateUser", reflect.TypeOf((*MockUtsuruDataStore)(nil).InsertOrUpdateUser), arg0)
}

// IsUserEmailExists mocks base method
func (m *MockUtsuruDataStore) IsUserEmailExists(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUserEmailExists", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUserEmailExists indicates an expected call of IsUserEmailExists
func (mr *MockUtsuruDataStoreMockRecorder) IsUserEmailExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUserEmailExists", reflect.TypeOf((*MockUtsuruDataStore)(nil).IsUserEmailExists), arg0)
}

// IsUserIdExists mocks base method
func (m *MockUtsuruDataStore) IsUserIdExists(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUserIdExists", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUserIdExists indicates an expected call of IsUserIdExists
func (mr *MockUtsuruDataStoreMockRecorder) IsUserIdExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUserIdExists", reflect.TypeOf((*MockUtsuruDataStore)(nil).IsUserIdExists), arg0)
}

// IsUserImageExists mocks base method
func (m *MockUtsuruDataStore) IsUserImageExists(arg0, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUserImageExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUserImageExists indicates an expected call of IsUserImageExists
func (mr *MockUtsuruDataStoreMockRecorder) IsUserImageExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUserImageExists", reflect.TypeOf((*MockUtsuruDataStore)(nil).IsUserImageExists), arg0, arg1)
}
