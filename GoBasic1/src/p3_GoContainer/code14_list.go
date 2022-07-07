package main

import (
	"container/list"
	"fmt"
)

func testList() {
	l := list.New()

	l.PushBack("first")
	l.PushFront(67)
	fmt.Println(l)

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

func main() {
	testList()
}
