package main

import (
	"fmt"
	"os"
	"os/exec"
)

var homeu string
var pkgu string

func updater() {
	cmd := exec.Command("rc", fmt.Sprintf("%s/packages/%s/dpmu.rc", homeu, pkgu))
	if err := cmd.Run(); err != nil {
		fmt.Println("Could not find dpmi.rc (probably does not exist):", err)
		return
	}
}

func update() {
	homeu = os.Getenv("home")
	fmt.Println("This tool upgrades packages from $home/packages. Enter the path of the package.", home)
	fmt.Println("All packages are stored at $home/packages. Run '$echo home' to find your home path.")
	fmt.Println("")
	fmt.Print("Package Path:")

	fmt.Scanln(&pkgu)
	
	fmt.Println("")
	fmt.Printf("You chose '%v', is this correct? [y/n] ", pkgu)
	var confirm string
	fmt.Scanln(&confirm)
	
	if confirm == "Y" || confirm == "y" || confirm == "yes" || confirm == "Yes" || confirm == "YES" {
		updater()
		fmt.Println("Operation Completed.")
		return
	} else {
		fmt.Println("Operation cancelled.")
		return
	}
}