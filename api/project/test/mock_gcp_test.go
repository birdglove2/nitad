// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/birdglove2/nitad-backend/gcp (interfaces: Uploader)

package project

import (
	context "context"
	multipart "mime/multipart"
	reflect "reflect"

	errors "github.com/birdglove2/nitad-backend/errors"
	gomock "github.com/golang/mock/gomock"
)

// MockUploader is a mock of Uploader interface.
type MockUploader struct {
	ctrl     *gomock.Controller
	recorder *MockUploaderMockRecorder
}

// MockUploaderMockRecorder is the mock recorder for MockUploader.
type MockUploaderMockRecorder struct {
	mock *MockUploader
}

// NewMockUploader creates a new mock instance.
func NewMockUploader(ctrl *gomock.Controller) *MockUploader {
	mock := &MockUploader{ctrl: ctrl}
	mock.recorder = &MockUploaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUploader) EXPECT() *MockUploaderMockRecorder {
	return m.recorder
}

// DeleteFile mocks base method.
func (m *MockUploader) DeleteFile(arg0 context.Context, arg1, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteFile", arg0, arg1, arg2)
}

// DeleteFile indicates an expected call of DeleteFile.
func (mr *MockUploaderMockRecorder) DeleteFile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFile", reflect.TypeOf((*MockUploader)(nil).DeleteFile), arg0, arg1, arg2)
}

// DeleteFiles mocks base method.
func (m *MockUploader) DeleteFiles(arg0 context.Context, arg1 []string, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteFiles", arg0, arg1, arg2)
}

// DeleteFiles indicates an expected call of DeleteFiles.
func (mr *MockUploaderMockRecorder) DeleteFiles(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFiles", reflect.TypeOf((*MockUploader)(nil).DeleteFiles), arg0, arg1, arg2)
}

// UploadFile mocks base method.
func (m *MockUploader) UploadFile(arg0 context.Context, arg1 *multipart.FileHeader, arg2 string) (string, errors.CustomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFile", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(errors.CustomError)
	return ret0, ret1
}

// UploadFile indicates an expected call of UploadFile.
func (mr *MockUploaderMockRecorder) UploadFile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockUploader)(nil).UploadFile), arg0, arg1, arg2)
}

// UploadFiles mocks base method.
func (m *MockUploader) UploadFiles(arg0 context.Context, arg1 []*multipart.FileHeader, arg2 string) ([]string, errors.CustomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFiles", arg0, arg1, arg2)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(errors.CustomError)
	return ret0, ret1
}

// UploadFiles indicates an expected call of UploadFiles.
func (mr *MockUploaderMockRecorder) UploadFiles(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFiles", reflect.TypeOf((*MockUploader)(nil).UploadFiles), arg0, arg1, arg2)
}
