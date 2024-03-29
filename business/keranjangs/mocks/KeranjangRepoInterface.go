// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	keranjangs "kemejaku/business/keranjangs"

	mock "github.com/stretchr/testify/mock"
)

// KeranjangRepoInterface is an autogenerated mock type for the KeranjangRepoInterface type
type KeranjangRepoInterface struct {
	mock.Mock
}

// DeleteKeranjang provides a mock function with given fields: id, ctx
func (_m *KeranjangRepoInterface) DeleteKeranjang(id int, ctx context.Context) (keranjangs.Keranjang, error) {
	ret := _m.Called(id, ctx)

	var r0 keranjangs.Keranjang
	if rf, ok := ret.Get(0).(func(int, context.Context) keranjangs.Keranjang); ok {
		r0 = rf(id, ctx)
	} else {
		r0 = ret.Get(0).(keranjangs.Keranjang)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, context.Context) error); ok {
		r1 = rf(id, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EditKeranjang provides a mock function with given fields: keranjang, id, ctx
func (_m *KeranjangRepoInterface) EditKeranjang(keranjang keranjangs.Keranjang, id int, ctx context.Context) (keranjangs.Keranjang, error) {
	ret := _m.Called(keranjang, id, ctx)

	var r0 keranjangs.Keranjang
	if rf, ok := ret.Get(0).(func(keranjangs.Keranjang, int, context.Context) keranjangs.Keranjang); ok {
		r0 = rf(keranjang, id, ctx)
	} else {
		r0 = ret.Get(0).(keranjangs.Keranjang)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(keranjangs.Keranjang, int, context.Context) error); ok {
		r1 = rf(keranjang, id, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllKeranjang provides a mock function with given fields: ctx
func (_m *KeranjangRepoInterface) GetAllKeranjang(ctx context.Context) ([]keranjangs.Keranjang, error) {
	ret := _m.Called(ctx)

	var r0 []keranjangs.Keranjang
	if rf, ok := ret.Get(0).(func(context.Context) []keranjangs.Keranjang); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]keranjangs.Keranjang)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetKeranjangDetail provides a mock function with given fields: id, ctx
func (_m *KeranjangRepoInterface) GetKeranjangDetail(id int, ctx context.Context) (keranjangs.Keranjang, error) {
	ret := _m.Called(id, ctx)

	var r0 keranjangs.Keranjang
	if rf, ok := ret.Get(0).(func(int, context.Context) keranjangs.Keranjang); ok {
		r0 = rf(id, ctx)
	} else {
		r0 = ret.Get(0).(keranjangs.Keranjang)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, context.Context) error); ok {
		r1 = rf(id, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertKeranjang provides a mock function with given fields: keranjang, ctx
func (_m *KeranjangRepoInterface) InsertKeranjang(keranjang keranjangs.Keranjang, ctx context.Context) (keranjangs.Keranjang, error) {
	ret := _m.Called(keranjang, ctx)

	var r0 keranjangs.Keranjang
	if rf, ok := ret.Get(0).(func(keranjangs.Keranjang, context.Context) keranjangs.Keranjang); ok {
		r0 = rf(keranjang, ctx)
	} else {
		r0 = ret.Get(0).(keranjangs.Keranjang)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(keranjangs.Keranjang, context.Context) error); ok {
		r1 = rf(keranjang, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
