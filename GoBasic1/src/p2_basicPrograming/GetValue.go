package main

import "fmt"

func swap1() {
	var a int = 100
	var b int = 200
	var t int

	t = a
	a = b
	b = t
	fmt.Println(a, b)
}

func swap2() {
	var a int = 100
	var b int = 200

	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, b)
}

// go 语言交换 多重赋值
func swap3() {
	var a int = 100
	var b int = 200

	b, a = a, b
	fmt.Println(a, b)
}

func swapValue(a, b int) {
	b, a = a, b
}

func swapPtrValue(a, b *int) {
	*b, *a = *a, *b
}

func swapPtrValue2(a, b *int) {
	b, a = a, b
}

func getValue() {
	a := 100
	b := 200
	fmt.Println(&a, &b)
	fmt.Println(a, b)

	swapPtrValue2(&a, &b)

	fmt.Println(&a, &b)
	fmt.Println(a, b)
}
