package nvidiasmi

import (
	"testing"
)

func TestGetNvidiaSmiDmonQuery(t *testing.T) {
	tests := []struct {
		inputSelectedQuery string
		testName           string
		inputCount         int
	}{
		{
			testName:           "Must return an error for utilization.gpu when the query is empty",
			inputSelectedQuery: "",
			inputCount:         1,
		},
		{
			testName:           "Must return an error for utilization.gpu when the query is invalid",
			inputSelectedQuery: "rand_input_query",
			inputCount:         1,
		},
		{
			testName:           "Must return an error when the input is less or equal than 0",
			inputSelectedQuery: "t",
			inputCount:         0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			_, err := GetNvidiaSmiDmonQueryGpu(tt.inputSelectedQuery, tt.inputCount)
			if err == nil {
				t.Errorf("GetNvidiaSmiDmonQueryGpu(%q, %q) must return an error", tt.inputSelectedQuery, tt.inputCount)
			}
		})
	}
}
