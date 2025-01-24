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

		run_command(cmd)
	}
}

func run_command(cmd string) {
	// built-in commands
	// builtInCommands :=

	switch {
	case cmd == "exit 0":
		os.Exit(0)
	case strings.HasPrefix(cmd, "echo"):
		echo_msg := strings.TrimSpace(strings.TrimLeft(cmd, "echo"))
		fmt.Println(echo_msg)
	default:
		msg := fmt.Sprintf("%s: command not found\n", cmd)
		fmt.Print(msg)
	}
}
