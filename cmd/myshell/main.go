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

		switch in {
			case "exit\n":
				os.Exit(0)
		}

		fmt.Printf("%s: command not found\n", strings.TrimRight(in, "\n"))
	}
}
