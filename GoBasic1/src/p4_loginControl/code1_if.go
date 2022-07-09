package main

import "fmt"

func TestIf() {
	if err := 8; err != 8 {
		fmt.Println("err != 8")
	} else {
		fmt.Println("err == 8")
	}
}
