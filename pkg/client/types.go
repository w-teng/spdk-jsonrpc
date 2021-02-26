package client

import (
	"encoding/json"
	"fmt"
)

const (
	JSONRPCVersion = "2.0"
)

type RPCRequest struct {
	RPCVersion string        `json:"jsonrpc"`
	Method     string        `json:"method"`
	ID         uint64        `json:"id"`
	Params     []interface{} `json:"params,omitempty"`
}

type RPCResponse struct {
	JsonRPCVersion string          `json:"jsonrpc"`
	ID             uint64          `json:"id"`
	Result         json.RawMessage `json:"result"`
	Error          RPCError        `json:"error"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e RPCError) Error() string {
	return fmt.Sprintf("Code=%d Msg=%s", e.Code, e.Message)
}
