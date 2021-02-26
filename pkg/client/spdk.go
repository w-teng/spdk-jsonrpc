package client

import "encoding/json"

type SPDKClientIface interface {
	GetRawClient() JsonRpcClientIface
	// RpcGetMethods rpc_get_methods
	RpcGetMethods() ([]string, error)
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
