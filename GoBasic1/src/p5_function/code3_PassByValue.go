package main

import "fmt"

// 测试 值传递效果的 结构体

type Data struct {
	complax  []int
	instance InnerData
	ptr      *InnerData
}

type InnerData struct {
	a int
}

func passByValue(inFunc Data) Data {
	fmt.Printf("inFunc value: %+v\n", inFunc)

	fmt.Printf("inFunc ptr: %p\n", &inFunc)
	return inFunc
}

func testPassByValue() {

	in := Data{
		complax:  []int{1, 2, 3},
		instance: InnerData{5},
		ptr:      &InnerData{1},
	}

	fmt.Printf("in value: %+v\n", in)
	fmt.Printf("in ptr: %p\n", &in)

	out := passByValue(in)

	fmt.Printf("out value: %+v\n", out)
	fmt.Printf("out ptr: %p\n", &out)
}
