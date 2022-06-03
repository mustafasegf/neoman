package components

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Urlbar struct {
	Url      textinput.Model
	Viewport viewport.Model
	Title    string
	State    state
	Parent   tea.Model
	Style    lipgloss.Style
}

func MakeUrlbar(url string, title string, size tea.WindowSizeMsg, updateSize UpdateSize, parent tea.Model) Urlbar {
	m := Urlbar{
		Url:    textinput.New(),
		Title:  title,
		State:  Typing,
		Parent: parent,
		Style:  lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Width(size.Width - 2 - updateSize.Width),
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
	case UpdateSize:
		m.Style = m.Style.Width(m.Style.GetWidth() + msg.Width)
	case tea.KeyMsg:
		s := msg.String()
		switch s {
		case "i":
			if m.State == Typing {
				m.Url, cmd = m.Url.Update(msg)
				cmds = append(cmds, cmd)
			} else {
				m.State = Typing
				cmds = append(cmds, m.Url.Focus())
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
	return m.Style.Render(m.Url.View())
}
