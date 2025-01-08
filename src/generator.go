package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"os"
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
}

func (g *Generator) dataCollect(buf chan TimeData, wg *sync.WaitGroup)  {
	defer wg.Done()

	for i := 0; i < g.DataAmount; i++ {
		if g.Min > g.Max {
			log.Fatalf("Min (%f) cannot be larger than Max (%f)", g.Min, g.Max)
		}
		data := TimeData {
			Timestamp: time.Now(),
			Value: g.Min + rand.Float64() * (g.Max - g.Min),
		}
		buf <- data
		fmt.Printf("Generated: %+v \n", data)
		time.Sleep(g.Interval)
	}

	close(buf)
}


// -+-----------------------------+-
//		Auxillary 
// -+-----------------------------+-
// Save to txt. Used for the early generator test
func (g *Generator)_saveToFile(data []TimeData, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	for _,d := range data {
		line := fmt.Sprintf("Timestamp: %s, Value: %f \n", d.Timestamp, d.Value)
		_, err := file.WriteString(line)
		if err != nil {
			return err
		}
	}
	return nil
}
