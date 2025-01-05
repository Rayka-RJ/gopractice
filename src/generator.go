package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"time"
)

type TimeData struct {
	Timestamp time.Time
	Value float64
}

type Generator struct {
	FileName string 
	DataAmount int
	Min float64
	Max float64
	Interval time.Duration
}

func NewGenerator(filename string, dataAmount int, min float64, max float64, interval time.Duration) *Generator {
	return &Generator{
		FileName: filename,
		DataAmount: dataAmount,
		Min: min,
		Max: max,
		Interval: interval,
	}
}

func (g *Generator) Run() {
	data := g.dataCollect()
	err := g.saveToFile(data)
	if err != nil {
		fmt.Printf("Fail to save the data: %s \n", err)
	} else {
		fmt.Printf("File successfully saved to %s \n", g.FileName)
	}
}

// -+---------------+-
//		Auxillary 
// -+---------------+-

func (g *Generator) dataCollect() []TimeData {
	if g.Min > g.Max {
		panic("Min cannot be larger than Max")
	}
	res := make([]TimeData, g.DataAmount)
	for i := range res {
		res[i] = TimeData{
			Timestamp: time.Now(),
			Value: g.Min + rand.Float64() * (g.Max - g.Min),
		}

		time.Sleep(g.Interval)
	}
	return res
}

func (g *Generator)saveToFile(data []TimeData) error {
	file, err := os.Create(g.FileName)
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
