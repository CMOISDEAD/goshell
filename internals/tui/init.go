package tui

import (
	"time"

	"github.com/CMOISDEAD/goshell/internals"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AliasStruct struct {
	alias   string
	command string
}

func InitialModel(prompt string) model {
	ti := textinput.New()
	ti.Placeholder = "One shell to rule them all..."
	ti.Focus()
	ti.CharLimit = 99999
	ti.Width = 100
	ti.BlinkSpeed = 500 * time.Millisecond
	ti.Prompt = prompt
	ti.PromptStyle = lipgloss.NewStyle().Bold(true)

	return model{
		textInput:    ti,
		err:          nil,
		code:         true,
		path:         internals.GetPwd(),
		history:      []string{"none"},
		historyIndex: 0,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}
