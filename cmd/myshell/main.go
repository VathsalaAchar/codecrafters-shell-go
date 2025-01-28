package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	builtin_cmds := []string{"echo", "exit", "type"}
	switch {
	case cmd == "exit 0":
		os.Exit(0)
	case strings.HasPrefix(cmd, "echo"):
		echo_msg := strings.TrimSpace(strings.TrimLeft(cmd, "echo"))
		fmt.Println(echo_msg)
	case strings.HasPrefix(cmd, "type"):
		var msg string
		cmd_to_type := strings.TrimSpace(strings.TrimLeft(cmd, "type"))
		if slices.Contains(builtin_cmds, cmd_to_type) {
			msg = fmt.Sprintf("%s is a shell builtin\n", cmd_to_type)
		} else {
			msg = fmt.Sprintf("%s: not found\n", cmd_to_type)
		}
		fmt.Print(msg)
	default:
		msg := fmt.Sprintf("%s: command not found\n", cmd)
		fmt.Print(msg)
	}
}
