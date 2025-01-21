package nvidiasmi

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

var ALLOWED_QUERIES = map[string]bool{
	"utilization.gpu":     true,
	"utilization.memory":  true,
	"utilization.encoder": true,
	"utilization.decoder": true,
	"fan.speed":           true,
}

func GetNvidiaSmiQueryGpu(query string) (string, error) {
	if _, exists := ALLOWED_QUERIES[query]; !exists {
		return "", errors.New("query not allowed")
	}
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
	return strings.TrimSpace(string(stdout)), nil
}
