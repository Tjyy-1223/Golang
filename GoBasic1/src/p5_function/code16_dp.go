package main

import (
	"fmt"
	"time"
)

const LIM = 41

var fibs [LIM]uint64

func fibonacci2(n int) (res uint64) {
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}

	if n <= 2 {
		res = 1
	} else {
		res = fibonacci2(n-1) + fibonacci2(n-2)
	}
	fibs[n] = res
	return
}

func originalTime() {
	result := 0
	start := time.Now()

	for i := 0; i < 40; i++ {
		result = int(fibonacci2(i))
		fmt.Printf("数列第 %d 位：%d\n", i, result)
	}

	elapsed := time.Since(start)
	fmt.Println("original time", elapsed)
}
