// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/evt/rest-api-example/model"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// FileMetaService is an autogenerated mock type for the FileMetaService type
type FileMetaService struct {
	mock.Mock
}

// CreateFileMeta provides a mock function with given fields: _a0, _a1
func (_m *FileMetaService) CreateFileMeta(_a0 context.Context, _a1 *model.File) (*model.File, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.File
	if rf, ok := ret.Get(0).(func(context.Context, *model.File) *model.File); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.File) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteFileMeta provides a mock function with given fields: _a0, _a1
func (_m *FileMetaService) DeleteFileMeta(_a0 context.Context, _a1 uuid.UUID) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetFileMeta provides a mock function with given fields: _a0, _a1
func (_m *FileMetaService) GetFileMeta(_a0 context.Context, _a1 uuid.UUID) (*model.File, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.File
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *model.File); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateFileMeta provides a mock function with given fields: _a0, _a1
func (_m *FileMetaService) UpdateFileMeta(_a0 context.Context, _a1 *model.File) (*model.File, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.File
	if rf, ok := ret.Get(0).(func(context.Context, *model.File) *model.File); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.File) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}