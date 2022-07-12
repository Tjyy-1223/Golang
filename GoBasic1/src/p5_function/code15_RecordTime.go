package main

import (
	"fmt"
	"time"
)

func recordTime() {
	start := time.Now()
	sum := 0

	for i := 0; i < 1000000; i++ {
		sum++
	}

	elapsed := time.Since(start)
	fmt.Println("time for a - b :", elapsed)
}
