package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"

	"github.com/skratchdot/open-golang/open"
)

func checkVersion(exe string, arg string, minVersion string) bool {
	var out bytes.Buffer
	cmd := exec.Command(exe, arg)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		return false
	}
	version := getVersion(out.String())
	fmt.Printf("%s version: %s\n", exe, version)
	if version < minVersion {
		return false
	}
	return true
}

func getVersion(word string) string {
	var verToken string
	re := regexp.MustCompile("(\\d+.)+\\d")
	verToken = re.FindString(word)
	if verToken == "" {
		verToken = "0.0.0"
	}
	return verToken
}

func checkRequisites() bool {
	const (
		NODE_VERSION, NPM_VERSION string = "4.0.0", "2.0.0"
		GIT_VERSION               string = "2.0.0"
	)

	fmt.Println("[Prerequisites for webOS CLI]")
	fmt.Println(" - node", NODE_VERSION)
	fmt.Println(" - npm", NPM_VERSION)
	fmt.Println(" - git", GIT_VERSION)
	if checkVersion("node", "--version", NODE_VERSION) == false {
		fmt.Println("Please install the newer node.js")
		open.Start("https://nodejs.org/en/download/")
		return false
	}

	if checkVersion("npm", "--version", NPM_VERSION) == false {
		fmt.Println("Please install the newer npm")
		open.Start("https://nodejs.org/en/download/")
		return false
	}

	if checkVersion("git", "--version", GIT_VERSION) == false {
		fmt.Println("Please install the git client on the terminal")
		if runtime.GOOS == "windows" {
			fmt.Println(" (Note.) Please select 'Use Git from the Windows Command Prompt on the installation setup'")
			open.Start("https://git-for-windows.github.io/")
			open.Start("https://www.npmjs.com/package/bower#windows-users")
		} else if runtime.GOOS == "darwin" {
			fmt.Println(" (Note.) You can install git as describing in the web page or your own ways like port, brew and etc.")
			open.Start("https://git-scm.com/book/en/v2/Getting-Started-Installing-Git#Installing-on-Mac")
		} else if runtime.GOOS == "linux" {
			fmt.Println(" (Note.) You can install git as describing in the web page or your own ways like apt, yum and etc.")
			open.Start("https://git-scm.com/book/en/v2/Getting-Started-Installing-Git#Installing-on-Linux")
		} else {
			fmt.Println(" Unknown OS. Please install git manually then try again.")
		}
		return false
	}
	return true
}

func main() {
	if checkRequisites() == false {
		fmt.Println("Need to install prerequisites.")
		os.Exit(1)
	} else {
		fmt.Println("Passed checking all Prerequisites.")
		os.Exit(0)
	}
}
