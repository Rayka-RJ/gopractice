package main

import (
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var rdmGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "real_time_data",
		Help: "Randomly generated value (float64)",
	},
	[] string {"timestamp"},
)

func getAlertData(buf <-chan TimeData, wg *sync.WaitGroup) {
	defer wg.Done()

	for data := range buf {
		rdmGauge.WithLabelValues(data.Timestamp.Format(time.RFC3339Nano)).Set(data.Value)		
	}
}



