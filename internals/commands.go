package internals

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

// env paths?
var path string = "/usr/bin/"

// commandNotFound error
type commandNotFound struct{}

func (m *commandNotFound) Error() string {
	return "Command not found"
}

// Main function
func Execute(input string) error {
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
			cmd := exec.Command(path+command, args[1:]...)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				return err
			}
		} else {
			fmt.Println("goshell: ", command, ": command not found...")
			return &commandNotFound{}
		}
	}
	return nil
}

// echo command like
func echo(words []string) {
	var res string
	for _, word := range words {
		res += word + " "
	}
	fmt.Println(res)
}

// ls command like
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

// check if the command exist in the paths
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
