package main

import "fmt"

func testDim2() {
	slice := [][]int{{10}, {100, 200}}
	//fmt.Println(slice[0][1])
	//run error
	slice[0] = append(slice[0], 20)
	fmt.Println(slice)
}

//func main() {
//	testDim2()
//}
