package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
		paths_to_check := strings.Split(os.Getenv("PATH"), ":")
		// set default message
		msg := fmt.Sprintf("%s: command not found", cmd)
		// get the executable name and arguments
		split_args := strings.SplitN(cmd, " ", 2)
		// fail if there is only 1 arg
		if len(split_args) < 2 {
			fmt.Println(msg)
			return
		}
		// if there are two or more arguments split into exe and arguments
		exe_name := split_args[0]
		args := split_args[1]

		for _, cpath := range paths_to_check {
			exec_path := filepath.Join(cpath, exe_name)
			_, err := os.Stat(exec_path)
			// if no error i.e., exe file exists
			if err == nil {
				// run the command
				c := exec.Command(exe_name, args)
				stdout, err := c.Output()
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				// get the output message
				fmt.Print(string(stdout))
				return
			}
		}
	}
}
