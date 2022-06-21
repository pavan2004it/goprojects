package main

import (
	"os/exec"
	"shell"
	"strings"
)

func main() {
	//shell.RunCLI()
	shell.RunCliWithInteractive(GeneratePrompt)
}

func GeneratePrompt() string {
	output, _ := exec.Command("pwd").Output()
	return string(strings.TrimSuffix(string(output), "\n")) + " # "
}
