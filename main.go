package main

import (
	"fmt"

	"github.com/CMOISDEAD/goshell/cmd"
	"github.com/CMOISDEAD/goshell/internals"
	"github.com/charmbracelet/lipgloss"
)

var welcome = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#25BE6A")).
	SetString("Goshell 0.1 version")

func main() {
	c := internals.Config(true)
	fmt.Println(welcome)
	cmd.Initialize(c.Char.(string), c.Scripts)
}
