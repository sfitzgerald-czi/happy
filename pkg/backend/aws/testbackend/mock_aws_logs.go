// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs (interfaces: GetLogEventsAPIClient)

// Package testbackend is a generated GoMock package.
package testbackend

import (
	context "context"
	reflect "reflect"

	cloudwatchlogs "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	gomock "github.com/golang/mock/gomock"
)

// MockGetLogEventsAPIClient is a mock of GetLogEventsAPIClient interface.
type MockGetLogEventsAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockGetLogEventsAPIClientMockRecorder
}

// MockGetLogEventsAPIClientMockRecorder is the mock recorder for MockGetLogEventsAPIClient.
type MockGetLogEventsAPIClientMockRecorder struct {
	mock *MockGetLogEventsAPIClient
}

// NewMockGetLogEventsAPIClient creates a new mock instance.
func NewMockGetLogEventsAPIClient(ctrl *gomock.Controller) *MockGetLogEventsAPIClient {
	mock := &MockGetLogEventsAPIClient{ctrl: ctrl}
	mock.recorder = &MockGetLogEventsAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetLogEventsAPIClient) EXPECT() *MockGetLogEventsAPIClientMockRecorder {
	return m.recorder
}

// GetLogEvents mocks base method.
func (m *MockGetLogEventsAPIClient) GetLogEvents(arg0 context.Context, arg1 *cloudwatchlogs.GetLogEventsInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetLogEventsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLogEvents", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.GetLogEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogEvents indicates an expected call of GetLogEvents.
func (mr *MockGetLogEventsAPIClientMockRecorder) GetLogEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogEvents", reflect.TypeOf((*MockGetLogEventsAPIClient)(nil).GetLogEvents), varargs...)
}
