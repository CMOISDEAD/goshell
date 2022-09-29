package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/CMOISDEAD/goshell/internals/tui"
	tea "github.com/charmbracelet/bubbletea"
)

type AliasStruct struct {
	alias   string
	command string
}

type Configuration struct {
	Char    interface{}
	Scripts []string
	Alias   []AliasStruct
}

func initScripts(scripts []string) {
	for _, script := range scripts {
		cmd := exec.Command("sh", script)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func Initialize(prompt string, scripts []string) {
	p := tea.NewProgram(tui.InitialModel(prompt))
	initScripts(scripts)
	for true {
		if err := p.Start(); err != nil {
			log.Fatal(err)
			break
		}
	}
}
