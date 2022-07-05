package main

import "fmt"

func Ptr() {
	var house = "Malibu Point 10880,90265"

	// & 代表取值操作符号
	ptr := &house

	// 打印ptr的类型和指针地址
	fmt.Printf("ptr type: %T\n", ptr)
	fmt.Printf("address: %p\n", ptr)

	// * 代表取值操作
	value := *ptr
	fmt.Printf("value type: %T\n", value)
	fmt.Printf("value: %s\n", value)

}
