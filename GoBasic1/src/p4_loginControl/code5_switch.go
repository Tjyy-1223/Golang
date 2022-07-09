package main

import "fmt"

func testSwitch() {
	var a string = "mum"

	switch a {
	case "mum", "daddy":
		fmt.Println("family")
	default:
		fmt.Println("stranger")
	}
}
