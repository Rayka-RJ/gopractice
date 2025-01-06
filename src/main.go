package main

import "time"

func main() { 
	// Data Preparation
	// Currently saved in txt
	gene := NewGenerator("timedata_100k.txt", 100_000, -1000, 1000, 1*time.Millisecond)
	gene.Run()

}
