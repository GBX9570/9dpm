package main

import (
	"fmt"
	"time"
	"os/exec"
	"strings"
	"os"
)

var pkg string
var parts []string
var packageName string
var home string

func github() {
	url := fmt.Sprintf("https://github.com/%s", pkg)
	fmt.Println("Attempting fetch from Github...")
	cmd := exec.Command("/bin/git/clone", url, fmt.Sprintf("%s/packages/%s", home, packageName))
	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to fetch package (GitHub):", err)
		return
	}
}

func sourcehut() {
	url := fmt.Sprintf("https://git.sr.ht/%s", pkg)
	fmt.Println("Attempting fetch from SourceHut...")
	cmd := exec.Command("/bin/git/clone", url, fmt.Sprintf("%s/packages/%s", home, packageName))
	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to fetch package (SourceHut):", err)
		return
	}
}

func gitlab() {
	url := fmt.Sprintf("https://gitlab.com/%s", pkg)
	fmt.Println("Attempeting fetch from GitLab...")
	cmd := exec.Command("/bin/git/clone", url, fmt.Sprintf("%s/packages/%s", home, packageName))
	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to fetch package (GitLab):", err)
		return
	}
}

func install() {
	fmt.Println("Searching for and running dpmd.rc")
	cmd := exec.Command("rc", fmt.Sprintf("%s/packages/%s/dpmd.rc", home, packageName))
	if err := cmd.Run(); err != nil {
		fmt.Println("")
		fmt.Println("Failed to install dependencies (dpmd.rc probably doesn't exist, probably nothing to worry about):", err)
	}

	install2()
}

func install2() {
	fmt.Println("Searching for and running dmpi.rc")
	cmd := exec.Command("rc", fmt.Sprintf("%s/packages/%s/dpmi.rc", home, packageName))
	if err := cmd.Run(); err != nil {
		fmt.Println("")
		fmt.Println("Failed to run installation steps. This package likely does not have a dpmi.rc, therefore it has not")
		fmt.Println("been setup to use 9dpm or does not require installation steps. The package is in:", fmt.Sprintf("%s/packages/%s/", home, packageName))
		fmt.Println(", look there for further information or the repo it is from. Error information is below:")
		fmt.Println(err)
		os.Exit(0)
	}
}

func Fetch() {
	fmt.Println("Expect errors! This scans multiple git repos, therefore if a project does not exist in one")
	fmt.Println("it will error out. Unless this tool fails to download from all repos, this is nothing to worry about.")
	fmt.Println("")
	fmt.Println("Enter the name of package you would like to install below, in the format <git user>/<project>.")
	fmt.Println("For example, for 'https://github.com/example/project' you would enter example/project.")
	fmt.Println("")
	fmt.Print("Package Name:")
	fmt.Scanln(&pkg)
	fmt.Printf("You entered '%v', is this correct? [y/n]", pkg)
	
	var confirm string
	fmt.Scanln(&confirm)

	parts = strings.Split(pkg, "/")
	
	packageName = parts[1]

	home = os.Getenv("home")

	if confirm == "Y" || confirm == "y" || confirm == "yes" || confirm == "Yes" || confirm == "YES" {
		github()
		sourcehut()
		gitlab()
		time.Sleep(5 * time.Second)
		install()
		fmt.Println("Installation has completed!")
		return
	} else {
		fmt.Println("Operation cancelled.")
		return
	}
}