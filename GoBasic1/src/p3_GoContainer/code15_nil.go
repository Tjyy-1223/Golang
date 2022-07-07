package main

import "fmt"

func testNil() {
	var arr []int
	var num *int
	fmt.Printf("%p\n", arr)
	fmt.Printf("%p\n", num)
}

//func main() {
//	testNil()
//}
