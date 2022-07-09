package main

import "fmt"

func TestFor() {
	for j := 0; j < 5; j++ {
		for i, k := 0, 0; i < 10; i++ {
			k++
			if i > 5 {
				goto JLoop
			}
			fmt.Println(i)
		}
	}

JLoop:
	fmt.Println("end for")
}
