package components

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	style = lipgloss.NewStyle().Border(lipgloss.NormalBorder())
)

type state int

type Urlbar struct {
	Url      textinput.Model
	Viewport viewport.Model
	Title    string
	State    state
	Parent   tea.Model
}

const (
	Typing state = iota
	Finished
)

func MakeUrlbar(url string, title string, parent tea.Model) Urlbar {
	m := Urlbar{
		Url:    textinput.New(),
		Title:  title,
		State:  Typing,
		Parent: parent,
	}

	p := parent.(*MainModel)

	style = style.Width(p.Viewport.Width - 2)
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
	case tea.WindowSizeMsg:

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
	return style.Render(m.Url.View())
}
