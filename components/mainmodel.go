package components

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type MainModel struct {
	urlbar   Urlbar
	Viewport viewport.Model
	Ready    bool
}

func InitialModel() MainModel {
	m := MainModel{
		Ready: false,
	}

	return m
}

func (m MainModel) InitComponent() MainModel {
	m.urlbar = MakeUrlbar("ini url", "", &m)
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
	case tea.WindowSizeMsg:
		fmt.Printf("window size msg: %v\n", msg)
		time.Sleep(time.Second)
		if !m.Ready {
			m.Viewport = viewport.New(msg.Width, msg.Height)
			m.Ready = true
			m = m.InitComponent()
		} else {
			m.Viewport.Width = msg.Width
			m.Viewport.Height = msg.Height
		}
	case tea.KeyMsg:
		s := msg.String()
		switch s {
		case "ctrl+c":
			return m, tea.Quit

		}
	}

	urlbarmodel, cmd := m.urlbar.Update(msg)
	cmds = append(cmds, cmd)
	m.urlbar = urlbarmodel.(Urlbar)
	return m, tea.Batch(cmds...)
}

func (m MainModel) View() string {
	if !m.Ready {
		return "initializing"
	}
	return m.urlbar.View()
}
