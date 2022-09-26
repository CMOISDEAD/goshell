package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/CMOISDEAD/goshell/internals"
)

func main() {
	// Shell init
	fmt.Println("Goshell, version 0.01")
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print(">> ")
		scanner.Scan()
		userinput := scanner.Text()
		if len(userinput) > 0 {
			internals.Execute(userinput)
		}
	}
}
