package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var KnowCommands = map[string]int{"exit": 0, "echo": 1, "type": 2}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, "$ ")

		in, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
		}

		trim := strings.TrimSpace(in)

		command := strings.Split(trim, " ")

		switch command[0] {
		case "exit":
			code, err := strconv.Atoi(command[1])
			if err != nil {
				fmt.Println(err.Error())
			}
			os.Exit(code)
		case "echo":
			fmt.Println(strings.Join(command[1:], " "))
		case "type":
			switch command[1] {
			case "exit", "echo", "type":
				fmt.Printf("%s is a shell builtin\n", command[1])
			default:
				env := os.Getenv("PATH")
				paths := strings.Split(env, ":")

				for _, path := range paths {
					exec := path + "/" + command[1]

					if _, err := os.Stat(exec); err == nil {
						fmt.Printf("%s is %s\n", command[1], exec)
					}
				}
				fmt.Printf("%s not found\n", command[1])
			}
		default:
			fmt.Printf("%s: command not found\n", command[0])
		}
	}
}
