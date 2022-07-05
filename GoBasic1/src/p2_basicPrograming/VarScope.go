package main

import "fmt"

//global variable
var a float32 = 3.14

func varScope() {
	// variable
	var a int = 3

	fmt.Printf("a = %d\n", a)
}
