package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type MainModel struct{}

func initialModel() MainModel {
	m := MainModel{}

	return m
}

func (m MainModel) Init() tea.Cmd {
	var cmds []tea.Cmd
	return tea.Batch(cmds...)
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		s := msg.String()
		switch s {
		case "ctrl+c", "esc", "q":
			return m, tea.Quit

		}
	}
	return m, tea.Batch(cmds...)
}

func (m MainModel) View() string {
	return "hello world"
}

func main() {
	if err := tea.NewProgram(initialModel(), tea.WithAltScreen()).Start(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}
