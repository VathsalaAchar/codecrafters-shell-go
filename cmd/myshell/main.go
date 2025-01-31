package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
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
	builtin_cmds := []string{"echo", "exit", "type", "pwd", "cd"}
	switch {
	case cmd == "exit 0":
		os.Exit(0)
	case cmd == "pwd":
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(dir)
	case strings.HasPrefix(cmd, "cd"):
		// change directory here
		dir_to_ch := strings.TrimSpace(strings.TrimLeft(cmd, "cd"))
		if dir_to_ch == "~" {
			dir_to_ch = os.Getenv("HOME")
		}
		err := os.Chdir(dir_to_ch)
		if err != nil {
			fmt.Printf("cd: %s: No such file or directory\n", dir_to_ch)
		}
	case strings.HasPrefix(cmd, "echo"):
		echo_msg := strings.TrimSpace(strings.TrimLeft(cmd, "echo"))
		if strings.Contains(echo_msg, "'") {
			// remove the single quotes and join
			echo_msg = strings.Join(remove_single_quotes(echo_msg), "")
		} else {
			echo_msg = remove_space((echo_msg))
		}
		fmt.Println(echo_msg)
	case strings.HasPrefix(cmd, "type"):
		var msg string
		cmd_to_type := strings.TrimSpace(strings.TrimLeft(cmd, "type"))

		if slices.Contains(builtin_cmds, cmd_to_type) {
			msg = fmt.Sprintf("%s is a shell builtin", cmd_to_type)
		} else {
			msg = get_type(cmd_to_type)
		}
		fmt.Println(msg)
	default:
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
		run_exe(exe_name, args)

	}
}

func remove_space(msg string) string {
	space := regexp.MustCompile(`\s+`)
	msg = space.ReplaceAllString(msg, " ")
	return msg
}

func remove_single_quotes(msg string) []string {
	msg_no_quotes := make([]string, 1)
	msg_split := strings.SplitAfter(msg, "'")
	for _, elem := range msg_split {
		if elem != "'" {
			msg_no_quotes = append(msg_no_quotes, strings.TrimSuffix(elem, "'"))
		}
	}
	return msg_no_quotes
}

func get_type(cmd_to_type string) (msg string) {
	// get the path from environment variable
	paths_to_check := strings.Split(os.Getenv("PATH"), ":")
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
	return
}

func run_exe(exe_name, args string) {
	// get the path from environment variable
	paths_to_check := strings.Split(os.Getenv("PATH"), ":")
	// set of arguments
	var arguments []string
	if strings.Contains(args, "'") {
		arguments = remove_single_quotes(args)
	} else {
		arguments = strings.Split(args, " ")
	}

	// if executable exists then run it
	for _, cpath := range paths_to_check {
		exec_path := filepath.Join(cpath, exe_name)
		_, err := os.Stat(exec_path)
		// if no error i.e., exe file exists
		if err == nil {
			// run the command
			c := exec.Command(exe_name, arguments...)
			stdout, _ := c.Output()
			// get the output message
			fmt.Print(string(stdout))
			return
		}
	}
}
