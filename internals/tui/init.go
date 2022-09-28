package tui

import (
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

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
		textInput: ti,
		err:       nil,
		code:      true,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}