package nvidiasmi

import (
	"os/exec"
)

func GetNvidiaSmiQueryGpu() (string, error) {
	stat, err := RunNvidiaSmiQueryGpuCmd()
	if err != nil {
		return "", err
	}

	return stat, nil
}

func RunNvidiaSmiQueryGpuCmd() (string, error) {
	app := "nvidia-smi"
	arg0 := "--query-gpu=utilization.gpu"
	arg1 := "--format=csv,noheader,nounits"

	cmd := exec.Command(app, arg0, arg1)
	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(stdout), nil
}
