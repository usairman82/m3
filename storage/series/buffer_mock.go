// Copyright (c) 2017 Uber Technologies, Inc.
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

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/m3db/m3db/storage/series/buffer.go

package series

import (
	gomock "github.com/golang/mock/gomock"
	context "github.com/m3db/m3db/context"
	block "github.com/m3db/m3db/storage/block"
	io "github.com/m3db/m3db/x/io"
	time "github.com/m3db/m3x/time"
	time0 "time"
)

// Mock of databaseBuffer interface
type MockdatabaseBuffer struct {
	ctrl     *gomock.Controller
	recorder *_MockdatabaseBufferRecorder
}

// Recorder for MockdatabaseBuffer (not exported)
type _MockdatabaseBufferRecorder struct {
	mock *MockdatabaseBuffer
}

func NewMockdatabaseBuffer(ctrl *gomock.Controller) *MockdatabaseBuffer {
	mock := &MockdatabaseBuffer{ctrl: ctrl}
	mock.recorder = &_MockdatabaseBufferRecorder{mock}
	return mock
}

func (_m *MockdatabaseBuffer) EXPECT() *_MockdatabaseBufferRecorder {
	return _m.recorder
}

func (_m *MockdatabaseBuffer) Write(ctx context.Context, timestamp time0.Time, value float64, unit time.Unit, annotation []byte) error {
	ret := _m.ctrl.Call(_m, "Write", ctx, timestamp, value, unit, annotation)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseBufferRecorder) Write(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockdatabaseBuffer) ReadEncoded(ctx context.Context, start time0.Time, end time0.Time) [][]io.SegmentReader {
	ret := _m.ctrl.Call(_m, "ReadEncoded", ctx, start, end)
	ret0, _ := ret[0].([][]io.SegmentReader)
	return ret0
}

func (_mr *_MockdatabaseBufferRecorder) ReadEncoded(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadEncoded", arg0, arg1, arg2)
}

func (_m *MockdatabaseBuffer) FetchBlocks(ctx context.Context, starts []time0.Time) []block.FetchBlockResult {
	ret := _m.ctrl.Call(_m, "FetchBlocks", ctx, starts)
	ret0, _ := ret[0].([]block.FetchBlockResult)
	return ret0
}

func (_mr *_MockdatabaseBufferRecorder) FetchBlocks(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchBlocks", arg0, arg1)
}

func (_m *MockdatabaseBuffer) FetchBlocksMetadata(ctx context.Context, start time0.Time, end time0.Time, includeSizes bool, includeChecksums bool) block.FetchBlockMetadataResults {
	ret := _m.ctrl.Call(_m, "FetchBlocksMetadata", ctx, start, end, includeSizes, includeChecksums)
	ret0, _ := ret[0].(block.FetchBlockMetadataResults)
	return ret0
}

func (_mr *_MockdatabaseBufferRecorder) FetchBlocksMetadata(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchBlocksMetadata", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockdatabaseBuffer) IsEmpty() bool {
	ret := _m.ctrl.Call(_m, "IsEmpty")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockdatabaseBufferRecorder) IsEmpty() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsEmpty")
}

func (_m *MockdatabaseBuffer) MinMax() (time0.Time, time0.Time) {
	ret := _m.ctrl.Call(_m, "MinMax")
	ret0, _ := ret[0].(time0.Time)
	ret1, _ := ret[1].(time0.Time)
	return ret0, ret1
}

func (_mr *_MockdatabaseBufferRecorder) MinMax() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MinMax")
}

func (_m *MockdatabaseBuffer) NeedsDrain() bool {
	ret := _m.ctrl.Call(_m, "NeedsDrain")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockdatabaseBufferRecorder) NeedsDrain() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NeedsDrain")
}

func (_m *MockdatabaseBuffer) DrainAndReset() {
	_m.ctrl.Call(_m, "DrainAndReset")
}

func (_mr *_MockdatabaseBufferRecorder) DrainAndReset() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DrainAndReset")
}

func (_m *MockdatabaseBuffer) Bootstrap(bl block.DatabaseBlock) error {
	ret := _m.ctrl.Call(_m, "Bootstrap", bl)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseBufferRecorder) Bootstrap(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Bootstrap", arg0)
}

func (_m *MockdatabaseBuffer) Reset() {
	_m.ctrl.Call(_m, "Reset")
}

func (_mr *_MockdatabaseBufferRecorder) Reset() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reset")
}
