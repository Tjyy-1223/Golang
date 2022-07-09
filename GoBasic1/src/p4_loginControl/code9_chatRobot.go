package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ChatRobot() {
	// get the data
	inputReader := bufio.NewReader(os.Stdin)

	fmt.Println("please input your name:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("An error occurred:%s\n", err)
		os.Exit(1)
	} else {
		name := input[:len(input)-2]
		fmt.Printf("hello,%s!what can i do for you?\n", name)
	}

	for true {
		input, err = inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred:%s\n", err)
			continue
		}
		input = input[:len(input)-2]
		input = strings.ToLower(input)
		switch input {
		case "":
			continue
		case "nothing", "bye":
			fmt.Println("Bye")
			os.Exit(0)
		default:
			fmt.Println("Sorry,I didn't catch you.")
		}
	}
}
