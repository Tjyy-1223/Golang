package main

import "fmt"

func testPanic() {
	defer fmt.Println("panic thing 1")
	defer fmt.Println("panic thing 2")
	panic(" my panic ")
}
