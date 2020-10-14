// Code generated by mockery v1.0.0. DO NOT EDIT.

package vulnsrc

import (
	db "github.com/aquasecurity/trivy-db/pkg/db"
	mock "github.com/stretchr/testify/mock"
)

// MockOperation is an autogenerated mock type for the Operation type
type MockOperation struct {
	mock.Mock
}

type SetMetadataArgs struct {
	Metadata         db.Metadata
	MetadataAnything bool
}

type SetMetadataReturns struct {
	Err error
}

type SetMetadataExpectation struct {
	Args    SetMetadataArgs
	Returns SetMetadataReturns
}

type StoreMetadataArgs struct {
	Metadata         db.Metadata
	MetadataAnything bool
	Dir              string
	DirAnything      bool
}

type StoreMetadataReturns struct {
	Err error
}

type StoreMetadataExpectation struct {
	Args    StoreMetadataArgs
	Returns StoreMetadataReturns
}

func (_m *MockOperation) ApplySetMetadataExpectation(e SetMetadataExpectation) {
	var args []interface{}
	if e.Args.MetadataAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Metadata)
	}
	_m.On("SetMetadata", args...).Return(e.Returns.Err)
}

func (_m *MockOperation) ApplySetMetadataExpectations(expectations []SetMetadataExpectation) {
	for _, e := range expectations {
		_m.ApplySetMetadataExpectation(e)
	}
}

// SetMetadata provides a mock function with given fields: metadata
func (_m *MockOperation) SetMetadata(metadata db.Metadata) error {
	ret := _m.Called(metadata)

	var r0 error
	if rf, ok := ret.Get(0).(func(db.Metadata) error); ok {
		r0 = rf(metadata)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *MockOperation) ApplyStoreMetadataExpectation(e StoreMetadataExpectation) {
	var args []interface{}
	if e.Args.MetadataAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Metadata)
	}

	if e.Args.DirAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Dir)
	}

	_m.On("StoreMetadata", args...).Return(e.Returns.Err)
}

func (_m *MockOperation) ApplyStoreMetadataExpectations(expectations []StoreMetadataExpectation) {
	for _, e := range expectations {
		_m.ApplyStoreMetadataExpectation(e)
	}
}

// StoreMetadata provides a mock function with given fields: metadata, dir
func (_m *MockOperation) StoreMetadata(metadata db.Metadata, dir string) error {
	ret := _m.Called(metadata, dir)

	var r0 error
	if rf, ok := ret.Get(0).(func(db.Metadata, string) error); ok {
		r0 = rf(metadata, dir)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}