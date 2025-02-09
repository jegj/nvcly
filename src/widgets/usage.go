package widgets

import (
	"time"

	"github.com/gizak/termui/v3/widgets"
)

type UsageWidget struct {
	*widgets.SparklineGroup
	updateInterval time.Duration
}

func NewUsageWidget(title string, updateInterval time.Duration) *UsageWidget {
	return nil
}
