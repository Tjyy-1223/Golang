package main

import "fmt"

func testNew1() {
	var sum *int
	sum = new(int)
	*sum = 98
	fmt.Println(*sum)
}

func testNew2() {
	var sum *[5]int
	sum = new([5]int)
	sum[0] = 99
	fmt.Println(sum)
}

func main() {
	//testNew1()
	testNew2()
}
