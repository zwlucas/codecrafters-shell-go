package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, "$ ")

		in, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
		}

		trim := strings.TrimSpace(in)

		commands := strings.Split(trim, " ")

		switch commands[0] {
		case "exit":
			code, err := strconv.Atoi(commands[1])
			if err != nil {
				fmt.Println(err.Error())
			}
			os.Exit(code)
			return
		case "echo":
			fmt.Printf("%s\n", strings.Join(commands[1:], " "))
			return
		default:
			fmt.Printf("%s: command not found\n", commands[0])
			return
		}
	}
}
