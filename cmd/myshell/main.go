package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	cmd, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	// remove the newline from the end of cmd
	cmd = strings.TrimSuffix(cmd, "\n")
	fmt.Printf("%s: command not found", cmd)
}
