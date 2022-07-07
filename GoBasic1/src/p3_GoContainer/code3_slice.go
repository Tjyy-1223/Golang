package main

import "fmt"

func testSlice() {
	var a = [3]int{1, 2, 3}
	fmt.Println(a, a[1:2])
}

func testSlice2() {
	var highRiseBuilding [30]int

	for i := range highRiseBuilding {
		highRiseBuilding[i] = i
	}

	fmt.Println(highRiseBuilding[10:15])
	fmt.Println(highRiseBuilding[20:])
	fmt.Println(highRiseBuilding[:2])
}

func testSlice3() {
	var strList []string
	var numList []int
	var numListEmpty = []int{}

	fmt.Println(strList, numList, numListEmpty)
	fmt.Println(len(strList))
	fmt.Println(numListEmpty == nil)
}

func testSlice4() {
	a := make([]int, 2)
	b := make([]int, 2, 10)

	//a[3] = 10 //panic runtime error
	//b[2] = 10 // runtime error

	fmt.Println(a, b)
	fmt.Println(len(a), len(b))
}

//func main() {
//	//testSlice()
//
//	//testSlice2()
//
//	//testSlice3()
//	testSlice4()
//}
