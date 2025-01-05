package main

import "time"

func main() { 
	gene := NewGenerator("timedata.txt", 500, -30, 120, 10*time.Millisecond)
	gene.Run()
}
