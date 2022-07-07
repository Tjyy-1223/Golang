package main

import "fmt"

func testCopy1() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}

	//copy(slice2, slice1) // 只会复制slice1的前三个元素到slice2中
	copy(slice1, slice2) // 	只会复制slice2的3个元素到slice1中

	fmt.Println(slice1)
	fmt.Println(slice2)
}

func testCopy2() {
	// set number 1000
	const elementCount = 1000

	// 预分配足够多的元素切片
	srcData := make([]int, elementCount)

	// 将切片赋值
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}

	// 引用切片元素
	refData := srcData

	// 复制元素
	copyData := make([]int, elementCount)
	copy(copyData, srcData)

	srcData[0] = 999
	fmt.Println(refData[0])
	fmt.Println(copyData[0])

	refData[0] = 888
	fmt.Println(srcData[0])

	copyData[0] = 666
	fmt.Println(srcData[0])
}

//func main() {
//	//testCopy1()
//	testCopy2()
//}
