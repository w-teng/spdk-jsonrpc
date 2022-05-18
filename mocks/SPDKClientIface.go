// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	client "github.com/silentred/spdk-jsonrpc/pkg/client"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// SPDKClientIface is an autogenerated mock type for the SPDKClientIface type
type SPDKClientIface struct {
	mock.Mock
}

// BdevAioCreate provides a mock function with given fields: req
func (_m *SPDKClientIface) BdevAioCreate(req client.BdevAioCreateReq) (string, error) {
	ret := _m.Called(req)

	var r0 string
	if rf, ok := ret.Get(0).(func(client.BdevAioCreateReq) string); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(client.BdevAioCreateReq) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BdevAioDelete provides a mock function with given fields: req
func (_m *SPDKClientIface) BdevAioDelete(req client.BdevAioDeleteReq) (bool, error) {
	ret := _m.Called(req)

	var r0 bool
	if rf, ok := ret.Get(0).(func(client.BdevAioDeleteReq) bool); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(client.BdevAioDeleteReq) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BdevGetBdevs provides a mock function with given fields:
func (_m *SPDKClientIface) BdevGetBdevs() ([]client.Bdev, error) {
	ret := _m.Called()

	var r0 []client.Bdev
	if rf, ok := ret.Get(0).(func() []client.Bdev); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]client.Bdev)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BdevLVolCreate provides a mock function with given fields: req
func (_m *SPDKClientIface) BdevLVolCreate(req client.BdevLVolCreateReq) (string, error) {
	ret := _m.Called(req)

	var r0 string
	if rf, ok := ret.Get(0).(func(client.BdevLVolCreateReq) string); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(client.BdevLVolCreateReq) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BdevLVolCreateLVStore provides a mock function with given fields: req
func (_m *SPDKClientIface) BdevLVolCreateLVStore(req client.BdevLVolCreateLVStoreReq) (string, error) {
	ret := _m.Called(req)

	var r0 string
	if rf, ok := ret.Get(0).(func(client.BdevLVolCreateLVStoreReq) string); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(client.BdevLVolCreateLVStoreReq) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FrameworkGetConfig provides a mock function with given fields: req
func (_m *SPDKClientIface) FrameworkGetConfig(req client.FrameworkGetConfigReq) ([]client.FrameworkGetConfigItem, error) {
	ret := _m.Called(req)

	var r0 []client.FrameworkGetConfigItem
	if rf, ok := ret.Get(0).(func(client.FrameworkGetConfigReq) []client.FrameworkGetConfigItem); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]client.FrameworkGetConfigItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(client.FrameworkGetConfigReq) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FrameworkGetSubsystems provides a mock function with given fields:
func (_m *SPDKClientIface) FrameworkGetSubsystems() ([]client.FrameworkGetSubsystemsItem, error) {
	ret := _m.Called()

	var r0 []client.FrameworkGetSubsystemsItem
	if rf, ok := ret.Get(0).(func() []client.FrameworkGetSubsystemsItem); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]client.FrameworkGetSubsystemsItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRawClient provides a mock function with given fields:
func (_m *SPDKClientIface) GetRawClient() client.JsonRpcClientIface {
	ret := _m.Called()

	var r0 client.JsonRpcClientIface
	if rf, ok := ret.Get(0).(func() client.JsonRpcClientIface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.JsonRpcClientIface)
		}
	}

	return r0
}

// GetSpdkVersion provides a mock function with given fields:
func (_m *SPDKClientIface) GetSpdkVersion() (client.SpdkVersion, error) {
	ret := _m.Called()

	var r0 client.SpdkVersion
	if rf, ok := ret.Get(0).(func() client.SpdkVersion); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(client.SpdkVersion)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NVMFCreateSubsystem provides a mock function with given fields: req
func (_m *SPDKClientIface) NVMFCreateSubsystem(req client.NVMFCreateSubsystemReq) (bool, error) {
	ret := _m.Called(req)

	var r0 bool
	if rf, ok := ret.Get(0).(func(client.NVMFCreateSubsystemReq) bool); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(client.NVMFCreateSubsystemReq) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NVMFCreateTransport provides a mock function with given fields: req
func (_m *SPDKClientIface) NVMFCreateTransport(req client.NVMFCreateTransportReq) (bool, error) {
	ret := _m.Called(req)

	var r0 bool
	if rf, ok := ret.Get(0).(func(client.NVMFCreateTransportReq) bool); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(client.NVMFCreateTransportReq) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NVMFDeleteSubsystem provides a mock function with given fields: req
func (_m *SPDKClientIface) NVMFDeleteSubsystem(req client.NVMFDeleteSubsystemReq) (bool, error) {
	ret := _m.Called(req)

	var r0 bool
	if rf, ok := ret.Get(0).(func(client.NVMFDeleteSubsystemReq) bool); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(client.NVMFDeleteSubsystemReq) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NVMFGetSubsystems provides a mock function with given fields:
func (_m *SPDKClientIface) NVMFGetSubsystems() ([]client.Subsystem, error) {
	ret := _m.Called()

	var r0 []client.Subsystem
	if rf, ok := ret.Get(0).(func() []client.Subsystem); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]client.Subsystem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NVMFGetTransports provides a mock function with given fields:
func (_m *SPDKClientIface) NVMFGetTransports() ([]client.Transport, error) {
	ret := _m.Called()

	var r0 []client.Transport
	if rf, ok := ret.Get(0).(func() []client.Transport); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]client.Transport)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NVMFSubsystemAddListener provides a mock function with given fields: req
func (_m *SPDKClientIface) NVMFSubsystemAddListener(req client.NVMFSubsystemAddListenerReq) (bool, error) {
	ret := _m.Called(req)

	var r0 bool
	if rf, ok := ret.Get(0).(func(client.NVMFSubsystemAddListenerReq) bool); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(client.NVMFSubsystemAddListenerReq) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NVMFSubsystemAddNS provides a mock function with given fields: req
func (_m *SPDKClientIface) NVMFSubsystemAddNS(req client.NVMFSubsystemAddNSReq) (int, error) {
	ret := _m.Called(req)

	var r0 int
	if rf, ok := ret.Get(0).(func(client.NVMFSubsystemAddNSReq) int); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(client.NVMFSubsystemAddNSReq) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RpcGetMethods provides a mock function with given fields:
func (_m *SPDKClientIface) RpcGetMethods() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSPDKClientIface creates a new instance of SPDKClientIface. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewSPDKClientIface(t testing.TB) *SPDKClientIface {
	mock := &SPDKClientIface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
