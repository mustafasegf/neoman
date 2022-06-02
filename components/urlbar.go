package components

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type state int

type Urlbar struct {
	Url   textinput.Model
	Title string
	State state
}

const (
	Typing state = iota
	Finished
)

func MakeUrlbar(url string, title string) Urlbar {
	m := Urlbar{
		Url:   textinput.New(),
		Title: title,
		State: Typing,
	}
	m.Url.Focus()
	m.Url.SetValue(url)
	return m
}

func (m Urlbar) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, textinput.Blink)
	return tea.Batch(cmds...)
}

func (m Urlbar) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		s := msg.String()
		switch s {
		case "i":
			if m.State == Typing {
				m.Url, cmd = m.Url.Update(msg)
				cmds = append(cmds, cmd)
			} else {
				cmds = append(cmds, m.Url.Focus())
				m.State = Typing
			}
		case "esc":
			m.Url.Blur()
			m.State = Finished
		default:
			m.Url, cmd = m.Url.Update(msg)
			cmds = append(cmds, cmd)
		}

	default:
		m.Url, cmd = m.Url.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Urlbar) View() string {
	return m.Url.View()
}
