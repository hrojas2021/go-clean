// Code generated by MockGen. DO NOT EDIT.
// Source: response.go

// Package mock is a generated GoMock package.
package mock

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockResponse is a mock of Response interface.
type MockResponse struct {
	ctrl     *gomock.Controller
	recorder *MockResponseMockRecorder
}

// MockResponseMockRecorder is the mock recorder for MockResponse.
type MockResponseMockRecorder struct {
	mock *MockResponse
}

// NewMockResponse creates a new mock instance.
func NewMockResponse(ctrl *gomock.Controller) *MockResponse {
	mock := &MockResponse{ctrl: ctrl}
	mock.recorder = &MockResponseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResponse) EXPECT() *MockResponseMockRecorder {
	return m.recorder
}

// Fail mocks base method.
func (m *MockResponse) Fail(w http.ResponseWriter, r *http.Request, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Fail", w, r, err)
}

// Fail indicates an expected call of Fail.
func (mr *MockResponseMockRecorder) Fail(w, r, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fail", reflect.TypeOf((*MockResponse)(nil).Fail), w, r, err)
}

// Failf mocks base method.
func (m *MockResponse) Failf(w http.ResponseWriter, r *http.Request, format string, a ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{w, r, format}
	for _, a_2 := range a {
		varargs = append(varargs, a_2)
	}
	m.ctrl.Call(m, "Failf", varargs...)
}

// Failf indicates an expected call of Failf.
func (mr *MockResponseMockRecorder) Failf(w, r, format interface{}, a ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{w, r, format}, a...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Failf", reflect.TypeOf((*MockResponse)(nil).Failf), varargs...)
}

// JSON mocks base method.
func (m *MockResponse) JSON(w http.ResponseWriter, r *http.Request, data interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "JSON", w, r, data)
}

// JSON indicates an expected call of JSON.
func (mr *MockResponseMockRecorder) JSON(w, r, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JSON", reflect.TypeOf((*MockResponse)(nil).JSON), w, r, data)
}
