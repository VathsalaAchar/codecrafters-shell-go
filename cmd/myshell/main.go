package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		// remove the newline from the end of cmd
		cmd := strings.TrimSuffix(input, "\n")
		msg := fmt.Sprintf("%s: command not found\n", cmd)
		fmt.Print(msg)
	}
}
