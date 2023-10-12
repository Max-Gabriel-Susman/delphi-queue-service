package main

import (
	"fmt"
	"os/exec"
	"strings"
)

type Command struct {
	Name string
	Args []string
}

func main() {
	cmdA := Command{
		Name: "eksctl",
		Args: []string{"create", "cluster", "--name", "delphi-cluster", "--region", "us-east-1"},
	}
	cmdB := Command{
		Name: "ls",
		Args: []string{"-l"},
	}

	out, err := ExecuteCommand(cmdB.Name, cmdB.Args...)
	fmt.Println("out: ", out)
	if err != nil {
		fmt.Println("error: ", err)
	}

	out, err = ExecuteCommand(cmdA.Name, cmdA.Args...)
	fmt.Println("out: ", out)
	if err != nil {
		fmt.Println("error: ", err)
	}
}

func ExecuteCommand(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println("command execution failed: ", err)
		// panic(err)
	}

	fmt.Print(strings.TrimSpace(string(stdout)))
	return strings.TrimSpace(string(stdout)), nil
}
