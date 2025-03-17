package main

import (
	"fmt"
	"time"
)

func ShowHelp() {
	fmt.Println("HELP FOR 9DPM")
	fmt.Println("-------------")
	fmt.Println("9DPM stands for 9Front Decentralised Package Manager.")
	fmt.Println("To add it to your project, look at the documentation at")
	fmt.Println("https://git.sr.ht/GBX/9dpm. Below is a list of commands")
	fmt.Println("that you can use with 9DPM. Ensure webfs is enabled ! ")
	fmt.Println("-------------------------------------------------------")
	fmt.Println("9dpm -s")
	fmt.Println("Installs a package.")
	fmt.Println("");
	fmt.Println("9dpm -u")
	fmt.Println("Upgrades a package.")
	fmt.Println("")
	fmt.Println("9dpm -r")
	fmt.Println("Removes a package.")
	fmt.Println("")
	fmt.Println("9dpm -h")
	fmt.Println("Shows this message.")
	fmt.Println("")
	fmt.Println("9dpm -v")
	fmt.Println("Shows the version of 9DPM.")
	fmt.Println("")
	fmt.Println("9dpm -l")
	fmt.Println("Lists all packages in $home/packages. Currently broken,")
	fmt.Println("will be fixed in next update.")
	fmt.Println("-------------------------------------------------------")
	now := time.Now()
	fmt.Println("Current Time:", now)
	fmt.Println("Enjoy 9DPM!")
}