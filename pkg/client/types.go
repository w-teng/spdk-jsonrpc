package client

import (
	"encoding/json"
	"fmt"
)

const (
	JSONRPCVersion = "2.0"

	ClearMethodNone        ClearMethod = "none"
	ClearMethodUnmap       ClearMethod = "unmap"
	ClearMethodWriteZeroes ClearMethod = "write_zeroes"

	// "IPv4", "IPv6", "IB", or "FC"
	AddrFamilyIPv4 AddressFamily = "IPv4"
	AddrFamilyIPv6 AddressFamily = "IPv6"
	AddrFamilyIB   AddressFamily = "IB"
	AddrFamilyFC   AddressFamily = "FC"

	TransportTypeRDMA = "RDMA"
)

type ClearMethod string
type AddressFamily string

type RPCRequest struct {
	RPCVersion string      `json:"jsonrpc"`
	Method     string      `json:"method"`
	ID         uint64      `json:"id"`
	Params     interface{} `json:"params,omitempty"`
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

// BdevAioCreateReq bdev_aio_create
type BdevAioCreateReq struct {
	// required
	BdevName string `json:"name"`
	FileName string `json:"filename"`
	// optional
	BlockSize int `json:"block_size,omitempty"`
}

type BdevLVolCreateLVStoreReq struct {
	// required
	BdevName string `json:"bdev_name"`
	LvsName  string `json:"lvs_name"`
	// optional
	// Cluster size of the logical volume store in bytes
	ClusterSize int         `json:"cluster_sz,omitempty"`
	ClearMethod ClearMethod `json:"clear_method,omitempty"`
}

type BdevLVolCreateReq struct {
	// required
	LVolName string `json:"lvol_name"`
	// Desired size of logical volume in bytes; Size will be rounded up to a multiple of cluster size. Either uuid or lvs_name must be specified, but not both. lvol_name will be used in the alias of the created logical volume.
	Size int `json:"size"`

	// optional
	ThinProvision bool        `json:"thin_provision,omitempty"`
	UUID          string      `json:"uuid,omitempty"`
	LvsName       string      `json:"lvs_name,omitempty"`
	ClearMethod   ClearMethod `json:"clear_method,omitempty"`
}

type NVMFCreateSubsystemReq struct {
	// required
	NQN string `json:"nqn"`

	// optional
	// Parent NVMe-oF target name.
	TargetName string `json:"tgt_name,omitempty"`
	// Serial number of virtual controller
	SerialNumber string `json:"serial_number,omitempty"`
	// Model number of virtual controller
	ModelNumber   string `json:"model_number,omitempty"`
	MaxNamespaces int    `json:"max_namespaces,omitempty"`
	AllowAnyHost  bool   `json:"allow_any_host,omitempty"`
	ANAReporting  bool   `json:"ana_reporting,omitempty"`
}

type NVMFSubsystemAddNSReq struct {
	// required
	NQN       string            `json:"nqn"`
	Namespace NamespaceForAddNS `json:"namespace"`

	// optional
	TargetName string `json:"tgt_name,omitempty"`
}

type NamespaceForAddNS struct {
	// req
	BdevName string `json:"bdev_name"`

	// opt
	NSID     int    `json:"nsid,omitempty"`
	NGUID    string `json:"nguid,omitempty"`
	EUI64    string `json:"eui64,omitempty"`
	UUID     string `json:"uuid,omitempty"`
	PtplFile string `json:"ptpl_file,omitempty"`
}

type ListenAddress struct {
	// required
	TrType string        `json:"trtype"` // RDMA
	AdrFam AddressFamily `json:"adrfam"`
	TrAddr string        `json:"traddr"`
	// opt
	TrSvcID string `json:"trsvcid,omitempty"`
}

type NVMFSubsystemAddListenerReq struct {
	// required
	NQN string `json:"nqn"`
	// opt
	TargetName    string        `json:"tgt_name,omitempty"`
	ListenAddress ListenAddress `json:"listen_address,omitempty"`
}

type NVMFCreateTransportReq struct {
	// required
	TrType string `json:"trtype"`

	// optional
	TargetName          string `json:"tgt_name,omitempty"`
	MaxQueueDepth       int    `json:"max_queue_depth,omitempty"`
	MaxQPairsPerCtrlr   int    `json:"max_qpairs_per_ctrlr,omitempty"`
	MaxIOQPairsPerCtrlr int    `json:"max_io_qpairs_per_ctrlr,omitempty"`
	InCapsuleDataSize   int    `json:"in_capsule_data_size,omitempty"`
	MaxIOSize           int    `json:"max_io_size,omitempty"`
	IOUnitSize          int    `json:"io_unit_size,omitempty"`
	MaxAQDepth          int    `json:"max_aq_depth,omitempty"`
	NumSharedBuffers    int    `json:"num_shared_buffers,omitempty"`
	BufCacheSize        int    `json:"buf_cache_size,omitempty"`
	NumCQE              int    `json:"num_cqe,omitempty"`
	MaxSRQDepth         int    `json:"max_srq_depth,omitempty"`
	SockPriority        int    `json:"sock_priority,omitempty"`
	AcceptorBacklog     int    `json:"acceptor_backlog,omitempty"`
	AbortTimeoutSec     int    `json:"abort_timeout_sec,omitempty"`

	NoWrBatching     bool `json:"no_wr_batching,omitempty"`
	DifInsertOrStrip bool `json:"dif_insert_or_strip,omitempty"`
	C2hSuccess       bool `json:"c2h_success,omitempty"`
	NoSRQ            bool `json:"no_srq,omitempty"`
}

type Transport struct {
	// TCP or RDMA
	TransType     string `json:"trtype"`
	MaxQueueDepth int    `json:"max_queue_depth"`
}

type Bdev struct {
	Name        string `json:"name"`
	ProductName string `json:"product_name"`
	UUID        string `json:"uuid"`
	BlockSize   int    `json:"block_size"`
	NumBlocks   int    `json:"num_blocks"`
}

type Subsystem struct {
	NQN             string          `json:""`
	ListenAddresses []ListenAddress `json:"listen_addresses"`
	AllowAnyHost    bool            `json:"allow_any_host"`
	SerialNum       string          `json:"serial_number"`
	ModelNum        string          `json:"model_number"`
	Namespaces      []Namespace     `json:"namespaces"`
}

type Namespace struct {
	NsID     int    `json:"nsid"`
	BdevName string `json:"bdev_Name"`
	Name     string `json:"name"`
	UUID     string `json:"uuid"`
}
