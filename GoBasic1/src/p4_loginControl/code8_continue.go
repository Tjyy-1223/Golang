package main

import "fmt"

func testContinue() {

Outer2Loop:
	for i := 0; i < 2; i++ {
		//fmt.Println(i)
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				continue Outer2Loop
			}
		}
	}

}
