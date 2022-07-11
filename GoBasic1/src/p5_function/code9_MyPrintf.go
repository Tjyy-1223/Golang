package main

import "fmt"

func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value")
		case string:
			fmt.Println(arg, "is a string value")
		case int64:
			fmt.Println(arg, "is an int64 value")
		default:
			fmt.Println(arg, "is an unknown type")
		}
	}
}

func testMyPrintf() {
	var v1 int = 1
	var v4 float32 = 1.234
	MyPrintf(v1, v4)
}

func rawPrint(rawList ...interface{}) {
	for _, a := range rawList {
		fmt.Println(a)
	}
}

func print(slist ...interface{}) {
	rawPrint(slist...)
}
