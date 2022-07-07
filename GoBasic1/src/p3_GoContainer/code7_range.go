package main

import "fmt"

func testRange1() {
	slice := []int{10, 20, 30, 40}

	for index, value := range slice {
		fmt.Printf("index : %d , value %d \n", index, value)
	}
}

// range 返回的是每个元素的副本，而不是直接返回对该元素的引用
func testRange2() {
	slice := []int{10, 20, 30, 40}

	for index, value := range slice {
		fmt.Printf("index : %d , value %d \n", index, value)
		//value = 10
		slice[index] = 10
	}
	fmt.Println(slice)
}

func testRange3() {
	slice := []int{10, 20, 30, 40}

	for i := range slice {
		fmt.Println(i)
	}
}

//func main() {
//	testRange3()
//}
