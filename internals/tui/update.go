package tui

import (
	"fmt"
	"os"

	"github.com/CMOISDEAD/goshell/internals"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var goodbye = lipgloss.NewStyle().Bold(true).Italic(true).Foreground(lipgloss.Color("#78A9FF")).SetString("Good bye!")

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			input := m.textInput.Value()
			if len(input) > 0 {
				for _, value := range internals.GetAlias(true) {
					if value.Alias == input {
						internals.Execute(value.Command)
						m.textInput.SetValue("")
						m.history = append(m.history, input)
						m.historyIndex = len(m.history)
						m.textInput.SetValue("")
						fmt.Print("\n\n")
						return m, nil
					}
				}
				err := internals.Execute(input)
				if err != nil {
					m.code = false
				} else {
					m.code = true
				}
				m.history = append(m.history, input)
				m.historyIndex = len(m.history)
				m.textInput.SetValue("")
				fmt.Print("\n\n")
			}
			return m, nil
		case tea.KeyCtrlD:
			fmt.Println(goodbye)
			os.Exit(0) // This need to be fixed
		case tea.KeyCtrlL:
			internals.Execute("clear")
		case tea.KeyUp:
			if m.historyIndex != 0 && m.historyIndex <= len(m.history) {
				m.historyIndex -= 1
			}
			m.textInput.SetValue(m.history[m.historyIndex])
		case tea.KeyDown:
			if m.historyIndex != len(m.history)-1 && m.historyIndex <= len(m.history) {
				m.historyIndex += 1
			}
			m.textInput.SetValue(m.history[m.historyIndex])
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}
