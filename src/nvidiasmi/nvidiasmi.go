package nvidiasmi

import (
	"encoding/xml"
	"os/exec"
)

type NvidiaSmiLog struct {
	DriverVersion string         `xml:"driver_version"`
	CudaVersion   string         `xml:"cuda_version"`
	GPU           []NvidiaSmiGpu `xml:"gpu"`
	AttachedGpus  int            `xml:"attached_gpus"`
}

type NvidiaSmiGpu struct {
	Id                  string                 `xml:"id,attr"`
	ProductName         string                 `xml:"product_name"`
	ProductBrand        string                 `xml:"product_brand"`
	PraductArchitecture string                 `xml:"product_architecture"`
	Uuid                string                 `xml:"uuid"`
	VbiosVersion        string                 `xml:"vbios_version"`
	MemoryUsage         NvidiaSmiFbMemoryUsage `xml:"fb_memory_usage"`
	PerformanceState    string                 `xml:"performance_state"`
	FanSpeed            string                 `xml:"fan_speed"`
	Utilization         NvidiaSmiUtilization   `xml:"utilization"`
	Processes           NvidiaSmiProcesses     `xml:"processes"`
}

type NvidiaSmiFbMemoryUsage struct {
	Total    string `xml:"total"`
	Reserved string `xml:"reserved"`
	Used     string `xml:"used"`
	Free     string `xml:"free"`
}

type NvidiaSmiUtilization struct {
	GpuUtil     string `xml:"gpu_util"`
	MemoryUtil  string `xml:"memory_util"`
	EncoderUtil string `xml:"encoder_util"`
	DecoderUtil string `xml:"decoder_util"`
	JpegUtil    string `xml:"jpeg_util"`
	OfaUtil     string `xml:"ofa_util"`
}

type NvidiaSmiProcesses struct {
	ProcessInfo []NvidiaSmiProcessInfo `xml:"process_info"`
}

type NvidiaSmiProcessInfo struct {
	GpuInstanceId     string `xml:"gpu_instance_id"`
	ComputeInstanceId string `xml:"compute_instance_id"`
	Pid               string `xml:"pid"`
	Type              string `xml:"type"`
	ProcessName       string `xml:"process_name"`
	UsedMemory        string `xml:"used_memory"`
}

func GetNvidiaSmiStats() (*NvidiaSmiLog, error) {
	out, err := RunNvidiaSmiCmd()
	if err != nil {
		return nil, err
	}
	stat, err := ParseNvidiaSmiCmdOutput(out)
	if err != nil {
		return nil, err
	}
	return stat, nil
}

func RunNvidiaSmiCmd() ([]byte, error) {
	app := "nvidia-smi"
	arg0 := "-q"
	arg1 := "-x"

	cmd := exec.Command(app, arg0, arg1)
	stdout, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return stdout, nil
}

func ParseNvidiaSmiCmdOutput(out []byte) (
	*NvidiaSmiLog,
	error,
) {
	var stat NvidiaSmiLog
	err := xml.Unmarshal(out, &stat)
	if err != nil {
		return nil, err
	}
	return &stat, err
}
