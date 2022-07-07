package main

import (
	"fmt"
)

//
//  TestDelete1
//  @Description: 测试删除函数
//

func TestDelete1() {
	a := []int{1, 2, 3}
	b := a[:1]
	b[0] = 99
	fmt.Println(a, b)
}

func TestDelete2() {
	a := []int{1, 2, 3}
	b := make([]int, 5)
	copy(b, a)
	b[0] = 99
	fmt.Println(a, b)
}

func TestDelete3() {
	seq := []string{"a", "b", "c", "d", "e"}

	index := 2

	fmt.Println(seq[:index], seq[index:])
	seq = append(seq[:index], seq[index+1:]...)
	fmt.Println(seq)
}

//func main() {
//	//TestDelete1()
//	//TestDelete2()
//	TestDelete3()
//}
