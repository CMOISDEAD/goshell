package internals

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

var path string = "/usr/bin/"

func Execute(input string) {
	args := strings.Fields(input)
	command := args[0]
	if command == "list" {
		list(args[1:])
	} else if command == "echo" {
		echo(args[1:])
	} else if command == "clear" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		if check(command) {
			cmd := exec.Command(path + command)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(command, " command dont exists")
		}
	}
}

func echo(words []string) {
	var res string
	for _, word := range words {
		res += word + " "
	}
	fmt.Print(res)
}

func list(args []string) {
	var res, path string

	if len(args) > 0 {
		path = args[0]
	} else {
		path = "."
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		res += file.Name() + "\n"
	}

	fmt.Print(res)
}
func check(command string) bool {
	_, err := os.Stat(path + command)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
