package nvidiasmi

import (
	"fmt"
	"os/exec"
)

func GetNvidiaSmiStats() error {
	app := "nvidia-smi"
	arg0 := "-q"
	arg1 := "-x"

	cmd := exec.Command(app, arg0, arg1)
	stdout, err := cmd.Output()
	if err != nil {
		return err
	}
	fmt.Println(string(stdout))
	return nil
}
