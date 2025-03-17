package main

import (
	"fmt"
	"os/exec"
	"os"
)

var env string

func lister() {
	env = os.Getenv("home")
	path := fmt.Sprintf("%s/packages", env)
	fmt.Println(path)
	cmd := exec.Command("/bin/ls", path)
	if err := cmd.Run(); err != nil {
		fmt.Println("Could not list packages ($home/packages probably doesn't exist)", err)
		return
	}
}

func list() {
	fmt.Println("Installed Packages")
	lister()
}