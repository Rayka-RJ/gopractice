package main

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

var rdmGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "real_time_data",
		Help: "Randomly generated value (float64)",
	},
	[] string {"Type"},
)

func (g *Generator) getAlertData(buf <-chan TimeData, wg *sync.WaitGroup) {
	defer wg.Done()

	for data := range buf {
		rdmGauge.WithLabelValues("Value").Set(data.Value)

		if data.Value > g.Threshold {
			rdmGauge.WithLabelValues("Alert").Set(1)		
		} else {
			rdmGauge.WithLabelValues("Alert").Set(0)				
		}
		
	}
}
