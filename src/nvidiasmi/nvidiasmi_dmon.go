package nvidiasmi

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func GetNvidiaSmiDmonQueryGpu(selectQuery string, count int) (string, error) {
	stat, err := runNvidiaSmiDmonQueryGpuCmd(selectQuery, count)
	if err != nil {
		return "", err
	}

	return stat, nil
}

func runNvidiaSmiDmonQueryGpuCmd(selectQuery string, count int) (string, error) {
	app := "nvidia-smi"
	arg0 := "dmon"
	arg1 := fmt.Sprintf("-s=%s", selectQuery)
	arg2 := fmt.Sprintf("-c=%d", count)
	arg3 := "--format=csv,noheader,nounit"

	cmd := exec.Command(app, arg0, arg1, arg2, arg3)
	log.Printf("===>%s", cmd.String())
	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(stdout)), nil
}
