// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cq-provider-aws/client (interfaces: EmrClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	emr "github.com/aws/aws-sdk-go-v2/service/emr"
	gomock "github.com/golang/mock/gomock"
)

// MockEmrClient is a mock of EmrClient interface.
type MockEmrClient struct {
	ctrl     *gomock.Controller
	recorder *MockEmrClientMockRecorder
}

// MockEmrClientMockRecorder is the mock recorder for MockEmrClient.
type MockEmrClientMockRecorder struct {
	mock *MockEmrClient
}

// NewMockEmrClient creates a new mock instance.
func NewMockEmrClient(ctrl *gomock.Controller) *MockEmrClient {
	mock := &MockEmrClient{ctrl: ctrl}
	mock.recorder = &MockEmrClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmrClient) EXPECT() *MockEmrClientMockRecorder {
	return m.recorder
}

// DescribeCluster mocks base method.
func (m *MockEmrClient) DescribeCluster(arg0 context.Context, arg1 *emr.DescribeClusterInput, arg2 ...func(*emr.Options)) (*emr.DescribeClusterOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCluster", varargs...)
	ret0, _ := ret[0].(*emr.DescribeClusterOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCluster indicates an expected call of DescribeCluster.
func (mr *MockEmrClientMockRecorder) DescribeCluster(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCluster", reflect.TypeOf((*MockEmrClient)(nil).DescribeCluster), varargs...)
}

// ListClusters mocks base method.
func (m *MockEmrClient) ListClusters(arg0 context.Context, arg1 *emr.ListClustersInput, arg2 ...func(*emr.Options)) (*emr.ListClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListClusters", varargs...)
	ret0, _ := ret[0].(*emr.ListClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClusters indicates an expected call of ListClusters.
func (mr *MockEmrClientMockRecorder) ListClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusters", reflect.TypeOf((*MockEmrClient)(nil).ListClusters), varargs...)
}
