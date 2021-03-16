package nvme

type NvmeDevice struct {
	DevicePath   string `json:"DevicePath"`
	ModelNumber  string `json:"ModelNumber"`
	SerialNumber string `json:"SerialNumber"`
	UsedBytes    uint64 `json:"UsedBytes"`
	PhysicalSize uint64 `json:"PhysicalSize"`
	SectorSize   uint64 `json:"SectorSize"`
}
