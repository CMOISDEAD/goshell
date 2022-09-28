package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	if m.code {
		m.textInput.PromptStyle.Foreground(lipgloss.Color("#DE8F78"))
	} else {
		m.textInput.PromptStyle.Foreground(lipgloss.Color("#EA5294"))
	}
	return fmt.Sprintf("%s", m.textInput.View()+"\n")
}
