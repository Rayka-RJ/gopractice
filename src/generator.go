package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"sync"
	"time"
)

type TimeData struct {
	Timestamp time.Time
	Value float64
}

type Generator struct {
	DataAmount int
	Min float64
	Max float64
	Interval time.Duration 
	Threshold float64
}

func (g *Generator) dataCollect(buf1, buf2 chan <- TimeData, wg *sync.WaitGroup)  {
	defer wg.Done()

	for i := 0; i < g.DataAmount; i++ {
		if g.Min > g.Max {
			log.Fatalf("Min (%f) cannot be larger than Max (%f)", g.Min, g.Max)
		}
		data := TimeData {
			Timestamp: time.Now(),
			Value: g.Min + rand.Float64() * (g.Max - g.Min),
		}
		buf1 <- data // For clickhouse
		buf2 <- data // For prometheus

		fmt.Printf("Generated: %+v \n", data)
		time.Sleep(g.Interval)
	}

	close(buf1)
	close(buf2)
}