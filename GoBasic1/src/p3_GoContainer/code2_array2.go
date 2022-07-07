package main

import "fmt"

func testDim2Array() {
	var array [4][2]int = [4][2]int{3: {1: 10}}
	array3 := [4][2]int{3: {1: 10}}

	for i := range array3 {
		fmt.Println(array3[i][0])
	}

	var array2 [2]int = array[3]

	value := array2[1]
	fmt.Println(value)
}

//func main() {
//	testDim2Array()
//}
