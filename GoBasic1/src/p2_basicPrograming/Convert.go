package main

import (
	"fmt"
	"strconv"
)

func TestParseInt() {
	str := "-11"
	num, err := strconv.ParseInt(str, 10, 0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}
}

func main() {
	TestParseInt()
}