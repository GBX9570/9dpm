package main

import (
	"os"
	"fmt"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "-h" {
		ShowHelp()
		return
	}

	if len(os.Args) > 1 && os.Args[1] == "-v" {
		ShowVersion()
		return
	}

	if len(os.Args) > 1 && os.Args[1] == "-s" {
		Fetch()
		return
	}


	if len(os.Args) > 1 && os.Args[1] == "-r" {
		remove()
		return
	}

	if len(os.Args) > 1 && os.Args[1] == "-u" {
		update()
		return
	}

	if len(os.Args) > 1 && os.Args[1] == "-l" {
		list()
		return
	}

	fmt.Println("No argument or invalid argument given. Run '9dpm -h' for a list of valid arguments.")
}