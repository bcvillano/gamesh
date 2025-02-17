package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Opening bash...")
	cmd := exec.Command("/bin/bash")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
