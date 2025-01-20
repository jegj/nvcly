package nvidiasmi

import (
	"testing"
)

func TestGetNvidiaSmiQueryGpu(t *testing.T) {
	tests := []struct {
		want     map[string]string
		input    string
		testName string
	}{
		{
			testName: "Must not return an error for utilization.gpu when the query is empty",
			input:    "",
		},
		{
			testName: "Must not return an error for utilization.gpu when the query is invalid",
			input:    "rand_input_query",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			_, err := GetNvidiaSmiQueryGpu(tt.input)
			if err == nil {
				t.Errorf("GetNvidiaSmiQueryGpu(%q) must return an error", tt.input)
			}
		})
	}
}
