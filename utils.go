package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Unpack list to string
func unpack(list []string) string {
	result := ""
	for _, v := range list {
		result += fmt.Sprintf(" %s", v)
	}
	return strings.TrimSpace(result)
}

// Run command
func run(command string) *exec.Cmd {
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
	return cmd
}
