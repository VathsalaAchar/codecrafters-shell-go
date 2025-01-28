package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
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
		paths_to_check := strings.Split(os.Getenv("PATH"), ":")
		cmd_to_type := strings.TrimSpace(strings.TrimLeft(cmd, "type"))

		if slices.Contains(builtin_cmds, cmd_to_type) {
			msg = fmt.Sprintf("%s is a shell builtin", cmd_to_type)
		} else {
			// assign message for failure
			msg = fmt.Sprintf("%s: not found", cmd_to_type)
			// then check if the command to find type is in path
			for _, cpath := range paths_to_check {
				exec_path := filepath.Join(cpath, cmd_to_type)
				_, err := os.Stat(exec_path)
				if err == nil {
					msg = fmt.Sprintf("%s is %s", cmd_to_type, exec_path)
					break
				}
			}
		}
		fmt.Println(msg)
	default:
		msg := fmt.Sprintf("%s: command not found", cmd)
		fmt.Println(msg)
	}
}
