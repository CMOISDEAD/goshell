package tui

import (
	"fmt"

	"github.com/CMOISDEAD/goshell/internals"
	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	if m.code {
		m.textInput.PromptStyle.Foreground(lipgloss.Color("#DE8F78"))
	} else {
		m.textInput.PromptStyle.Foreground(lipgloss.Color("#EA5294"))
	}
	path := lipgloss.
		NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7B679A")).
		SetString(internals.FormatPath(m.path))
	user := lipgloss.
		NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FF91C1")).
		SetString(internals.GetUser())
	date := lipgloss.
		NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#DE8F78")).
		SetString(internals.GetTime())
	return fmt.Sprintf("%s at %s in %s\n%s", date, user, path, m.textInput.View()+"\n")
}
