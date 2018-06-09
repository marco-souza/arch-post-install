package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var cmds string

// Unpack list to string
func unpack(list []string) string {
	result := ""
	for _, v := range list {
		result += fmt.Sprintf(" %s", v)
	}
	return strings.TrimSpace(result)
}

// Run command
func start() *exec.Cmd {
	cmd := exec.Command("bash", "-c", strings.TrimSpace(cmds))
	fmt.Println("TO RUN -> \n", cmds)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
	return cmd
}

func run(command string) {
	sep := " && \\\n"
	if cmds == "" {
		sep = ""
	}
	cmds += sep + command
}
