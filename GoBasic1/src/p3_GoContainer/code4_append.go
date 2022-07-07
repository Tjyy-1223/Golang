package main

import "fmt"

// 使用append为切片动态添加元素
func testAppend1() {
	var a []int

	a = append(a, 1)                 //追加一个元素
	a = append(a, 1, 2, 3)           //追加多个元素
	a = append(a, []int{1, 2, 3}...) //追加多个元素
	fmt.Println(a)
}

// 空间不足以容纳元素，切片会自动进行扩容
func testAppend2() {
	var number []int
	for i := 0; i < 10; i++ {
		number = append(number, i)
		fmt.Printf("len : %d cap : %d pointer : %p\n", len(number), cap(number), number)
	}
}

func testAppend3() {
	var a = []int{1, 2, 3}
	a = append(a[:1], append([]int{5, 6, 7}, a[1:]...)...)
	fmt.Println(a)
}

func main() {
	//testAppend1()

	//testAppend2()
	testAppend3()
}
