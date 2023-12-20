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
// Source: github.com/aws/amazon-ecs-agent/ecs-agent/netlib/platform (interfaces: API)

// Package mock_platform is a generated GoMock package.
package mock_platform

import (
	context "context"
	reflect "reflect"

	ecsacs "github.com/aws/amazon-ecs-agent/ecs-agent/acs/model/ecsacs"
	data "github.com/aws/amazon-ecs-agent/ecs-agent/netlib/data"
	appmesh "github.com/aws/amazon-ecs-agent/ecs-agent/netlib/model/appmesh"
	networkinterface "github.com/aws/amazon-ecs-agent/ecs-agent/netlib/model/networkinterface"
	serviceconnect "github.com/aws/amazon-ecs-agent/ecs-agent/netlib/model/serviceconnect"
	tasknetworkconfig "github.com/aws/amazon-ecs-agent/ecs-agent/netlib/model/tasknetworkconfig"
	gomock "github.com/golang/mock/gomock"
)

// MockAPI is a mock of API interface.
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI.
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance.
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// BuildTaskNetworkConfiguration mocks base method.
func (m *MockAPI) BuildTaskNetworkConfiguration(arg0 string, arg1 *ecsacs.Task) (*tasknetworkconfig.TaskNetworkConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildTaskNetworkConfiguration", arg0, arg1)
	ret0, _ := ret[0].(*tasknetworkconfig.TaskNetworkConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildTaskNetworkConfiguration indicates an expected call of BuildTaskNetworkConfiguration.
func (mr *MockAPIMockRecorder) BuildTaskNetworkConfiguration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildTaskNetworkConfiguration", reflect.TypeOf((*MockAPI)(nil).BuildTaskNetworkConfiguration), arg0, arg1)
}

// ConfigureAppMesh mocks base method.
func (m *MockAPI) ConfigureAppMesh(arg0 context.Context, arg1 string, arg2 *appmesh.AppMesh) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigureAppMesh", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConfigureAppMesh indicates an expected call of ConfigureAppMesh.
func (mr *MockAPIMockRecorder) ConfigureAppMesh(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigureAppMesh", reflect.TypeOf((*MockAPI)(nil).ConfigureAppMesh), arg0, arg1, arg2)
}

// ConfigureInterface mocks base method.
func (m *MockAPI) ConfigureInterface(arg0 context.Context, arg1 string, arg2 *networkinterface.NetworkInterface, arg3 data.NetworkDataClient) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigureInterface", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConfigureInterface indicates an expected call of ConfigureInterface.
func (mr *MockAPIMockRecorder) ConfigureInterface(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigureInterface", reflect.TypeOf((*MockAPI)(nil).ConfigureInterface), arg0, arg1, arg2, arg3)
}

// ConfigureServiceConnect mocks base method.
func (m *MockAPI) ConfigureServiceConnect(arg0 context.Context, arg1 string, arg2 *networkinterface.NetworkInterface, arg3 *serviceconnect.ServiceConnectConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigureServiceConnect", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConfigureServiceConnect indicates an expected call of ConfigureServiceConnect.
func (mr *MockAPIMockRecorder) ConfigureServiceConnect(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigureServiceConnect", reflect.TypeOf((*MockAPI)(nil).ConfigureServiceConnect), arg0, arg1, arg2, arg3)
}

// CreateDNSConfig mocks base method.
func (m *MockAPI) CreateDNSConfig(arg0 string, arg1 *tasknetworkconfig.NetworkNamespace) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDNSConfig", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDNSConfig indicates an expected call of CreateDNSConfig.
func (mr *MockAPIMockRecorder) CreateDNSConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDNSConfig", reflect.TypeOf((*MockAPI)(nil).CreateDNSConfig), arg0, arg1)
}

// CreateNetNS mocks base method.
func (m *MockAPI) CreateNetNS(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNetNS", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNetNS indicates an expected call of CreateNetNS.
func (mr *MockAPIMockRecorder) CreateNetNS(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNetNS", reflect.TypeOf((*MockAPI)(nil).CreateNetNS), arg0)
}

// DeleteDNSConfig mocks base method.
func (m *MockAPI) DeleteDNSConfig(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDNSConfig", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDNSConfig indicates an expected call of DeleteDNSConfig.
func (mr *MockAPIMockRecorder) DeleteDNSConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDNSConfig", reflect.TypeOf((*MockAPI)(nil).DeleteDNSConfig), arg0)
}

// DeleteNetNS mocks base method.
func (m *MockAPI) DeleteNetNS(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNetNS", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNetNS indicates an expected call of DeleteNetNS.
func (mr *MockAPIMockRecorder) DeleteNetNS(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNetNS", reflect.TypeOf((*MockAPI)(nil).DeleteNetNS), arg0)
}

// GetNetNSPath mocks base method.
func (m *MockAPI) GetNetNSPath(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetNSPath", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetNetNSPath indicates an expected call of GetNetNSPath.
func (mr *MockAPIMockRecorder) GetNetNSPath(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetNSPath", reflect.TypeOf((*MockAPI)(nil).GetNetNSPath), arg0)
}
