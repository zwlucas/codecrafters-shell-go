package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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

		command := strings.Split(trim, " ")

		switch command[0] {
		case "exit":
			Exit(command[1])
		case "echo":
			Echo(command[1:])
		case "type":
			Type(command[1])
		default:
			executable := exec.Command(command[0], command[1:]...)
			executable.Stderr = os.Stderr
			executable.Stdout = os.Stdout

			err := executable.Run()
			if err != nil {
				fmt.Printf("%s: command not found\n", command[0])
			}
		}
	}
}

func Echo(message []string) {
	fmt.Println(strings.Join(message, " "))
}

func Exit(code string) {
	exitCode, err := strconv.Atoi(code)

	if err != nil {
		fmt.Println(err.Error())
	}

	os.Exit(exitCode)
}

func Type(command string) {
	switch command {
	case "exit", "echo", "type":
		fmt.Printf("%s is a shell builtin\n", command)
	default:
		env := os.Getenv("PATH")
		paths := strings.Split(env, ":")

		for _, path := range paths {
			exec := path + "/" + command

			if _, err := os.Stat(exec); err == nil {
				fmt.Printf("%s is %s\n", command, exec)
				return
			}
		}
		fmt.Printf("%s not found\n", command)
	}
}
