package main

import "fmt"

func main() {

	_, hour, minute := resolveTime(18000)
	fmt.Println(hour, minute)

}
