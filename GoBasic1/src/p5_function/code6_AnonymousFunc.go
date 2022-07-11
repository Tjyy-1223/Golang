package main

import "fmt"

func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

func testVisit() {
	visit([]int{1, 2, 3, 4}, func(v int) {
		fmt.Println(v)
	})
}
