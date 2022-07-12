package main

import (
	"fmt"
	"runtime"
)

type panicContext struct {
	function string
}

func ProtectRun(entry func()) {
	defer func() {
		err := recover()

		switch err.(type) {
		case runtime.Error:
			fmt.Println("run time error:", err)
		default:
			fmt.Println("other errors", err)
		}
	}()
	entry()
}

func testProtectRun() {
	fmt.Println("before.....")

	ProtectRun(func() {
		fmt.Println("before panic...")

		panic(&panicContext{
			"hand for panic",
		})

		fmt.Println("after panic")

	})

	fmt.Println("continue....")
}
