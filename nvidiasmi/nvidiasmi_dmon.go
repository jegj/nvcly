package nvidiasmi

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

var ALLOWED_DMON_QUERIES = map[string]bool{
	"t": true,
}

func GetNvidiaSmiDmonQueryGpu(selectQuery string, count int) (string, error) {
	if _, exists := ALLOWED_DMON_QUERIES[selectQuery]; !exists {
		return "", errors.New("dmon query not allowed")
	}

	if count < 1 {
		return "", errors.New("dmon count invalid")
	}

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
	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(stdout)), nil
}
