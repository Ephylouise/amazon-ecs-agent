// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-agent/ecs-agent/netlib/data (interfaces: NetworkDataClient)

// Package mock_data is a generated GoMock package.
package mock_data

import (
	reflect "reflect"

	tasknetworkconfig "github.com/aws/amazon-ecs-agent/ecs-agent/netlib/model/tasknetworkconfig"
	gomock "github.com/golang/mock/gomock"
)

// MockNetworkDataClient is a mock of NetworkDataClient interface.
type MockNetworkDataClient struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkDataClientMockRecorder
}

// MockNetworkDataClientMockRecorder is the mock recorder for MockNetworkDataClient.
type MockNetworkDataClientMockRecorder struct {
	mock *MockNetworkDataClient
}

// NewMockNetworkDataClient creates a new mock instance.
func NewMockNetworkDataClient(ctrl *gomock.Controller) *MockNetworkDataClient {
	mock := &MockNetworkDataClient{ctrl: ctrl}
	mock.recorder = &MockNetworkDataClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkDataClient) EXPECT() *MockNetworkDataClientMockRecorder {
	return m.recorder
}

// AssignGeneveDstPort mocks base method.
func (m *MockNetworkDataClient) AssignGeneveDstPort(arg0 string) (uint16, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignGeneveDstPort", arg0)
	ret0, _ := ret[0].(uint16)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignGeneveDstPort indicates an expected call of AssignGeneveDstPort.
func (mr *MockNetworkDataClientMockRecorder) AssignGeneveDstPort(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignGeneveDstPort", reflect.TypeOf((*MockNetworkDataClient)(nil).AssignGeneveDstPort), arg0)
}

// GetNetworkNamespace mocks base method.
func (m *MockNetworkDataClient) GetNetworkNamespace(arg0 string) (*tasknetworkconfig.NetworkNamespace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkNamespace", arg0)
	ret0, _ := ret[0].(*tasknetworkconfig.NetworkNamespace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetworkNamespace indicates an expected call of GetNetworkNamespace.
func (mr *MockNetworkDataClientMockRecorder) GetNetworkNamespace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkNamespace", reflect.TypeOf((*MockNetworkDataClient)(nil).GetNetworkNamespace), arg0)
}

// GetNetworkNamespacesByTaskID mocks base method.
func (m *MockNetworkDataClient) GetNetworkNamespacesByTaskID(arg0 string) ([]*tasknetworkconfig.NetworkNamespace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkNamespacesByTaskID", arg0)
	ret0, _ := ret[0].([]*tasknetworkconfig.NetworkNamespace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetworkNamespacesByTaskID indicates an expected call of GetNetworkNamespacesByTaskID.
func (mr *MockNetworkDataClientMockRecorder) GetNetworkNamespacesByTaskID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkNamespacesByTaskID", reflect.TypeOf((*MockNetworkDataClient)(nil).GetNetworkNamespacesByTaskID), arg0)
}

// ReleaseGeneveDstPort mocks base method.
func (m *MockNetworkDataClient) ReleaseGeneveDstPort(arg0 uint16, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReleaseGeneveDstPort", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReleaseGeneveDstPort indicates an expected call of ReleaseGeneveDstPort.
func (mr *MockNetworkDataClientMockRecorder) ReleaseGeneveDstPort(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseGeneveDstPort", reflect.TypeOf((*MockNetworkDataClient)(nil).ReleaseGeneveDstPort), arg0, arg1)
}

// SaveNetworkNamespace mocks base method.
func (m *MockNetworkDataClient) SaveNetworkNamespace(arg0 *tasknetworkconfig.NetworkNamespace) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveNetworkNamespace", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveNetworkNamespace indicates an expected call of SaveNetworkNamespace.
func (mr *MockNetworkDataClientMockRecorder) SaveNetworkNamespace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveNetworkNamespace", reflect.TypeOf((*MockNetworkDataClient)(nil).SaveNetworkNamespace), arg0)
}
