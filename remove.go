package main

import (
	"fmt"
	"os"
	"os/exec"
)

var homer string
var pkgr string

func removal() {
	cmd := exec.Command("rm", "-rf", pkgr)
	if err := cmd.Run(); err != nil {
		fmt.Println("Could not delete package:", err)
		return
	}
}

func remove() {
	home = os.Getenv("home")
	fmt.Println("This tool removes packages from $home/packages. Enter the path of the package.", home)
	fmt.Println("All packages are stored at $home/packages. Run '$echo home' to find your home path.")
	fmt.Println("Parts of the package may remain as this tool removes only the package folder, not its data,")
	fmt.Println("")
	fmt.Print("Package Path:")

	fmt.Scanln(&pkgr)
	
	fmt.Println("")
	fmt.Printf("You chose '%v', is this correct? [y/n] ", pkgr)
	var confirm string
	fmt.Scanln(&confirm)
	
	if confirm == "Y" || confirm == "y" || confirm == "yes" || confirm == "Yes" || confirm == "YES" {
		removal()
		fmt.Println("Operation Completed.")
		return
	} else {
		fmt.Println("Operation cancelled.")
		return
	}
}

