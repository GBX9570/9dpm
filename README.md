# 9dpm
A git frontend for 9Front that makes installing programs and managing them easy!

# What is it?
9DPM stands for 9Front Decentralised Package Manager. It functions with dpmi.rc, dpmd.rc and dpmu.rc. These are easy to implement in programs, requiring no customisation to codebases.
It is an easy to use frontend for git that downloads from any git server that the user specifies in install.go. For developers, it is as easy as putting build, update and dependencie management in .rc files.

# Why should I use it over another package manager, i.e. Sports?
Sports (https://github.com/henesy/sports) has a very different design to 9dpm. Sports is intended to be a fully fledges, ready to use **package manager**. That is **not** what 9dpm is.
9dpm, at its core, is nothing more than git and rc. It is an easy to use git frontend that makes the lives of developers and users easier, but is **not** intended to replace something like Sports.
Sports and 9dpm both handle package management differently, and both have their pros and cons. Go for Sports if you want a **package manager**, go for 9dpm if you want a *pacman-like*, *easy-to-use* **git frontend**.

# How does it work for users?
Users should first install Go (1.22.12 or newer, preferably). 9dpm is written in Go for its versatility. Secondly, they should pull the git repo and run `go build`. This will produce a 9dpm binary.
Running 9dpm on its own will return 'No argument or invalid argument given. Run '9dpm -h' for a list of valid arguments.'. Here is a list of arguments:

9dpm -s
Installs a package

9dpm -u
Upgrades a package.

9dpm -r
Removes a package.

9dpm -h
Shows this message.

9dpm -v
Shows the version of 9DPM.

9dpm -l
Lists all packages in $home/packages. Currently broken, will be fixed in next update.

All these commands provide clear information on how to use them while running them, but **installation** will be explained as it may be a bit daunting.
In a git URL, you have 3 parts - the domain (eg, github.com), the user (eg, /user/) and the project (eg, /project/). These come together to make a link (eg, github.com/user/project/)
in 9dpm, you want to take the **'user'** part and the **'project'** part, to create something that looks like **'user/project'**. This will be called your package address.
Firstly, create a folder in $home/packages (/usr/yourusername/packages).
Then, run 9dpm -s.
You then enter your previously created package address as the package to download, go through the prompts, and wait for the installation to finish.
Then, the package *should* be in $home/packages. You can generally ignore errors that git throws, unless all git servers fail to download.

# How does it work for developers?
For developers, it is just development as usual. Upload to your favourite Git server, do normal development, and your done. However, for proper integration with 9dpm, you need at least 2 of these files:

**necessary files**
dpmd.rc - This is a regular rc file. All of the following will be regular .rc files. Here, you simply put the commands to install dependencies - which could be a git/clone, another 9dpm command, literally anything. This is ran by 9dpm BEFORE dpmi.rc.
dpmi.rc - This rc file is much the same, except its for installation commands. This can be anything from mk install, git/clone, etc - an rc file its just running regular shell commands, just like a .sh file on Linux.
dpmu.rc - This rc file is slightly different in that it isn't ran during installation, its ran during 9dpm -u. This file just includes commands to update your package.

None of this is any different to your usual installation/updates/dependency management processes - its just in a .rc file now instead. Please note that all of these files must be at the root of your project folder, or 9dpm will not detect them.

# How do I add my own git servers?
This is the more technical part. You must go into install.go with acme and Go 1.22.12 or later installed. Copy this code block ABOVE the github() func but below the variables:

```go
func website() {
	url := fmt.Sprintf("https://website/%s", pkg)
	fmt.Println("Attempting fetch from website...")
	cmd := exec.Command("/bin/git/clone", url, fmt.Sprintf("%s/packages/%s", home, packageName))
	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to fetch package (website):", err)
		return
	}
}
```

and change the name of 'website' to your preferred git host. Now, go to the fetch() func and find this part:
```go
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
```

Here, below gitlab, enter the name of the function you just edited with () after it. Run `go build` again, and then your done! There's no limit to how many git hosts you have, but the more you have, the longer installations will take.
In an upcoming version, this will be replaced with a config file

# Final Notes
Thanks for looking at my project! Despite some of my previous projects, I am very much an amateur programmer, so any suggestions, improvments or ideas are welcome in the issues tab!
