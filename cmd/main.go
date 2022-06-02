package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mustafasegf/neoman/components"
)

type MainModel struct {
	urlbar components.Urlbar
}

func initialModel() MainModel {
	m := MainModel{
		urlbar: components.MakeUrlbar("ini url", ""),
	}

	return m
}

func (m MainModel) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, m.urlbar.Init())
	return tea.Batch(cmds...)
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		s := msg.String()
		switch s {
		case "ctrl+c":
			return m, tea.Quit

		}
	}

	urlbarmodel, cmd := m.urlbar.Update(msg)
	cmds = append(cmds, cmd)
	m.urlbar = urlbarmodel.(components.Urlbar)
	return m, tea.Batch(cmds...)
}

func (m MainModel) View() string {
	return m.urlbar.View()
}

func main() {
	if err := tea.NewProgram(initialModel(), tea.WithAltScreen()).Start(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}
