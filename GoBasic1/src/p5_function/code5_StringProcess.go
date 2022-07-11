package main

import (
	"fmt"
	"strings"
)

func StringProcess(list []string, chain []func(string) string) {
	for index, str := range list {

		//	第一个需要处理的字符串
		result := str

		for _, proc := range chain {
			result = proc(result)
		}

		list[index] = result
	}
}

// remove prefix

func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

func testPassList(l []int) {
	fmt.Printf("%p\n", &l)
	l[0] = 555
	l[1] = 666
	l[2] = 888
}

func testStringProcess() {
	list := []string{
		"go run",
		"go parser",
		"go compiler",
	}

	chain := []func(string) string{
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}

	StringProcess(list, chain)

	for _, str := range list {
		fmt.Println(str)
	}

	l := []int{1, 2, 3}
	fmt.Printf("%p\n", &l)
	fmt.Println(l)

	testPassList(l)
	fmt.Println(l)
}
