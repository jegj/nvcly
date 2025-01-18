package nvidiasmi

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetNvidiaSmiQueryGpu(query string) (string, error) {
	stat, err := runNvidiaSmiQueryGpuCmd(query)
	if err != nil {
		return "", err
	}

	return stat, nil
}

func runNvidiaSmiQueryGpuCmd(query string) (string, error) {
	app := "nvidia-smi"
	arg0 := fmt.Sprintf("--query-gpu=%s", query)
	arg1 := "--format=csv,noheader,nounits"

	cmd := exec.Command(app, arg0, arg1)
	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(strings.TrimSpace(string(stdout))), nil
}
