package tui

import "github.com/charmbracelet/bubbles/textinput"

type (
	tickMsg struct{}
	errMsg  error
)

type model struct {
	textInput    textinput.Model
	err          error
	code         bool
	path         string
	history      []string
	historyIndex int
}
