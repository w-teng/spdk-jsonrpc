package nvme

import (
	"encoding/json"
	"os/exec"
)

const (
	DefaultCmdName = "nvme"
)

type CmdClient struct {
	NvmeCmdPath string
}

// NewCmdClient finds the full file paht of the command nvme
func NewCmdClient() (cli *CmdClient, err error) {
	path, err := exec.LookPath(DefaultCmdName)
	if err != nil {
		return
	}

	cli = &CmdClient{
		NvmeCmdPath: path,
	}
	return
}

// NewCmdClientWithFilePath use file as the full path of command nvme
func NewCmdClientWithFilePath(file string) (cli *CmdClient) {
	return &CmdClient{
		NvmeCmdPath: file,
	}
}

func (cli *CmdClient) ListNvmeDisk() (list []NvmeDevice, err error) {
	bs, err := exec.Command(cli.NvmeCmdPath, "list", "-o", "json").Output()
	if err != nil {
		return
	}

	// no nvme devices
	if len(bs) == 0 {
		return
	}

	var resp = make(map[string][]NvmeDevice)
	err = json.Unmarshal(bs, &resp)
	if err != nil {
		return
	}

	list = resp["Devices"]
	return
}

func (cli *CmdClient) ConnectTarget(ip, svcID, nqn string, opts ...string) (output []byte, err error) {
	// nvme connect -t tcp -a 11.71.55.18 -s 4450 -n nqn.2021-03.com.alipay.ob:test-aio2
	// opts: --reconnect-delay 2 --ctrl-loss-tmo 10
	args := []string{
		"connect", "-t", "tcp", "-a", ip, "-s", svcID, "-n", nqn,
	}
	if len(opts) > 0 {
		args = append(args, opts...)
	}
	output, err = exec.Command(cli.NvmeCmdPath, args...).CombinedOutput()
	return
}

func (cli *CmdClient) DisconnectTarget(nqn string) (output []byte, err error) {
	// nvme disconnect -n nqn.2021-03.com.alipay.ob:test-aio2
	output, err = exec.Command(cli.NvmeCmdPath, "disconnect", "-n", nqn).CombinedOutput()
	return
}
