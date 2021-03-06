package main

import "fmt"

func TestRange() {
	for key, value := range []int{1, 2, 3, 4} {
		fmt.Printf("key:%d	value:%d\n", key, value)
	}

	for key, value := range "hello 你好" {
		fmt.Printf("key:%d	%c	value:0x%x\n", key, value, value)
	}

	m := map[string]int{
		"hello": 100,
		"world": 200,
	}

	for key, value := range m {
		fmt.Println(key, value)
	}

	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}
