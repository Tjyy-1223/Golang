package main

import "fmt"

type Invoke interface {
	Call(interface{})
}

type Struct struct {
}

func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

type FuncCaller func(interface{})

func (f FuncCaller) Call(p interface{}) {
	f(p)
}

func testInterface() {
	var invoker Invoke

	s := new(Struct)

	invoker = s

	invoker.Call("hello")

	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function", v)
	})

	invoker.Call("hello")
}