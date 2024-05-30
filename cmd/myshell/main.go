package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, "$ ")

		in, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
		}

		commands := strings.Split(in, " ")

		switch in {
			case commands[0] == "exit":
				n, err := strconv.Atoi(commands[1])
				os.Exit(n)
		}

		fmt.Printf("%s: command not found\n", strings.TrimRight(in, "\n"))
	}
}
