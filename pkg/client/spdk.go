package client

import "encoding/json"

var (
	_ SPDKClientIface = &SPDK{}
)

type SPDKClientIface interface {
	GetRawClient() JsonRpcClientIface
	// RpcGetMethods rpc_get_methods
	RpcGetMethods() ([]string, error)

	// nvmf_create_transport
	NVMFCreateTransport(req NVMFCreateTransportReq) (result bool, err error)

	// BdevAioCreate bdev_aio_create, return the name of bdev
	BdevAioCreate(req BdevAioCreateReq) (name string, err error)
	// bdev_lvol_create_lvstore
	BdevLVolCreateLVStore(req BdevLVolCreateLVStoreReq) (uuid string, err error)
	// bdev_lvol_create
	BdevLVolCreate(req BdevLVolCreateReq) (uuid string, err error)

	// nvmf_create_subsystem
	NVMFCreateSubsystem(req NVMFCreateSubsystemReq) (result bool, err error)
	// nvmf_subsystem_add_ns
	NVMFSubsystemAddNS(req NVMFSubsystemAddNSReq) (nsID int, err error)
	// nvmf_subsystem_add_listener
	NVMFSubsystemAddListener(req NVMFSubsystemAddListenerReq) (result bool, err error)
}

type SPDK struct {
	rawCli JsonRpcClientIface
}

func NewSPDK(rawCli JsonRpcClientIface) *SPDK {
	return &SPDK{
		rawCli: rawCli,
	}
}

func (s *SPDK) GetRawClient() JsonRpcClientIface {
	return s.rawCli
}

func (s *SPDK) RpcGetMethods() (methods []string, err error) {
	result, err := s.rawCli.Call("rpc_get_methods", nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(result, &methods)
	return
}

func (s *SPDK) BdevAioCreate(req BdevAioCreateReq) (name string, err error) {
	result, err := s.rawCli.Call("bdev_aio_create", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(result, &name)
	return
}

func (s *SPDK) NVMFCreateTransport(req NVMFCreateTransportReq) (res bool, err error) {
	result, err := s.rawCli.Call("nvmf_create_transport", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(result, &res)
	return
}

func (s *SPDK) BdevLVolCreateLVStore(req BdevLVolCreateLVStoreReq) (uuid string, err error) {
	result, err := s.rawCli.Call("bdev_lvol_create_lvstore", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(result, &uuid)
	return
}

// bdev_lvol_create
func (s *SPDK) BdevLVolCreate(req BdevLVolCreateReq) (uuid string, err error) {
	result, err := s.rawCli.Call("bdev_lvol_create", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(result, &uuid)
	return
}

func (s *SPDK) NVMFCreateSubsystem(req NVMFCreateSubsystemReq) (res bool, err error) {
	result, err := s.rawCli.Call("nvmf_create_subsystem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(result, &res)
	return
}

// nvmf_subsystem_add_ns
func (s *SPDK) NVMFSubsystemAddNS(req NVMFSubsystemAddNSReq) (nsID int, err error) {
	result, err := s.rawCli.Call("nvmf_create_transport", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(result, &nsID)
	return
}

// nvmf_subsystem_add_listener
func (s *SPDK) NVMFSubsystemAddListener(req NVMFSubsystemAddListenerReq) (res bool, err error) {
	result, err := s.rawCli.Call("nvmf_create_transport", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(result, &res)
	return
}
