package main

import "fmt"

func code1() {
	//var team [3]string = [3]string{"hammer", "solider", "mum"}
	team := [3]string{"hammer", "solider", "mum"}

	for k, v := range team {
		fmt.Println(k, v)
	}

	for i := range team {
		fmt.Println(i)
	}
}

//func main() {
//	code1()
//}
