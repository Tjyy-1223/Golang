package main

import "fmt"

func Accumulate(value int) func() int {
	// 返回一个函数
	return func() int {
		value++
		return value
	}
}

func testAccumulator() {
	accumulator := Accumulate(1)

	fmt.Println(accumulator())
	fmt.Println(accumulator())

	fmt.Printf("%p\n", &accumulator)

	accumulator2 := Accumulate(10)
	fmt.Println(accumulator2())
	fmt.Printf("%p\n", &accumulator2)
}

func PlayerGen(name string) func() (string, int) {
	hp := 150

	return func() (string, int) {
		return name, hp
	}
}

func testPlayerGen() {
	generator := PlayerGen("high mom")
	name, hp := generator()
	fmt.Println(name, hp)
}