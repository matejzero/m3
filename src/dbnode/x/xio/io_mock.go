// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3db/src/dbnode/x/xio (interfaces: SegmentReader,SegmentReaderPool)

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package xio is a generated GoMock package.
package xio

import (
	"reflect"

	"github.com/m3db/m3db/src/dbnode/ts"

	"github.com/golang/mock/gomock"
)

// MockSegmentReader is a mock of SegmentReader interface
type MockSegmentReader struct {
	ctrl     *gomock.Controller
	recorder *MockSegmentReaderMockRecorder
}

// MockSegmentReaderMockRecorder is the mock recorder for MockSegmentReader
type MockSegmentReaderMockRecorder struct {
	mock *MockSegmentReader
}

// NewMockSegmentReader creates a new mock instance
func NewMockSegmentReader(ctrl *gomock.Controller) *MockSegmentReader {
	mock := &MockSegmentReader{ctrl: ctrl}
	mock.recorder = &MockSegmentReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSegmentReader) EXPECT() *MockSegmentReaderMockRecorder {
	return m.recorder
}

// Clone mocks base method
func (m *MockSegmentReader) Clone() (SegmentReader, error) {
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(SegmentReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Clone indicates an expected call of Clone
func (mr *MockSegmentReaderMockRecorder) Clone() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockSegmentReader)(nil).Clone))
}

// Finalize mocks base method
func (m *MockSegmentReader) Finalize() {
	m.ctrl.Call(m, "Finalize")
}

// Finalize indicates an expected call of Finalize
func (mr *MockSegmentReaderMockRecorder) Finalize() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finalize", reflect.TypeOf((*MockSegmentReader)(nil).Finalize))
}

// Read mocks base method
func (m *MockSegmentReader) Read(arg0 []byte) (int, error) {
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockSegmentReaderMockRecorder) Read(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockSegmentReader)(nil).Read), arg0)
}

// Reset mocks base method
func (m *MockSegmentReader) Reset(arg0 ts.Segment) {
	m.ctrl.Call(m, "Reset", arg0)
}

// Reset indicates an expected call of Reset
func (mr *MockSegmentReaderMockRecorder) Reset(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockSegmentReader)(nil).Reset), arg0)
}

// Segment mocks base method
func (m *MockSegmentReader) Segment() (ts.Segment, error) {
	ret := m.ctrl.Call(m, "Segment")
	ret0, _ := ret[0].(ts.Segment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Segment indicates an expected call of Segment
func (mr *MockSegmentReaderMockRecorder) Segment() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Segment", reflect.TypeOf((*MockSegmentReader)(nil).Segment))
}

// MockSegmentReaderPool is a mock of SegmentReaderPool interface
type MockSegmentReaderPool struct {
	ctrl     *gomock.Controller
	recorder *MockSegmentReaderPoolMockRecorder
}

// MockSegmentReaderPoolMockRecorder is the mock recorder for MockSegmentReaderPool
type MockSegmentReaderPoolMockRecorder struct {
	mock *MockSegmentReaderPool
}

// NewMockSegmentReaderPool creates a new mock instance
func NewMockSegmentReaderPool(ctrl *gomock.Controller) *MockSegmentReaderPool {
	mock := &MockSegmentReaderPool{ctrl: ctrl}
	mock.recorder = &MockSegmentReaderPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSegmentReaderPool) EXPECT() *MockSegmentReaderPoolMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockSegmentReaderPool) Get() SegmentReader {
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(SegmentReader)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockSegmentReaderPoolMockRecorder) Get() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSegmentReaderPool)(nil).Get))
}

// Init mocks base method
func (m *MockSegmentReaderPool) Init() {
	m.ctrl.Call(m, "Init")
}

// Init indicates an expected call of Init
func (mr *MockSegmentReaderPoolMockRecorder) Init() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockSegmentReaderPool)(nil).Init))
}

// Put mocks base method
func (m *MockSegmentReaderPool) Put(arg0 SegmentReader) {
	m.ctrl.Call(m, "Put", arg0)
}

// Put indicates an expected call of Put
func (mr *MockSegmentReaderPoolMockRecorder) Put(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockSegmentReaderPool)(nil).Put), arg0)
}
