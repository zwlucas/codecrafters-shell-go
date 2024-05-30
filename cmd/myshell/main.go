package main

import (
	"bufio"
	"fmt"
	"os"
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

		commands := strings.Split(in, " ")

		switch in {
			case commnds[0] == "exit":
				os.Exit(commands[1].ToInt())
		}

		fmt.Printf("%s: command not found\n", strings.TrimRight(in, "\n"))
	}
}
