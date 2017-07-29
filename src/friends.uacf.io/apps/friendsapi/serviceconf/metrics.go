package serviceconf

import (
	"fmt"
	"time"

	"go.uacf.io/env"
	"go.uacf.io/metrics"
)

type MetricsConfig struct {
	Url           env.NetworkLink
	FlushInterval time.Duration
	FlushBytes    int
	DefaultRate   float32
}

func InitMetrics(conf MetricsConfig, environment, appName string) error {
	if conf.Url.Host != "" {
		metricsPrefix := fmt.Sprintf("%s.%s", environment, appName)
		return metrics.Init(conf.Url.Host, conf.Url.Port, metricsPrefix, metrics.Params{
			MaxFlushInterval: conf.FlushInterval,
			MaxFlushBytes:    conf.FlushBytes,
			SampleRate:       conf.DefaultRate,
		})
	}
	return nil
}
