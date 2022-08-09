package components

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TextBody struct {
	Body      textinput.Model
	Viewport viewport.Model
	State    state
	Parent   tea.Model
	Style    lipgloss.Style
}

func MakeTextBody(body string, size tea.WindowSizeMsg, updateSize UpdateSize, parent tea.Model) TextBody {
	m := TextBody{
		Body:    textinput.New(),
		State:  Typing,
		Parent: parent,
		Style:  lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Width(size.Width - 2 - updateSize.Width),
	}
  m.Body.SetValue(body)

	return m
}

func (m TextBody) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, textinput.Blink)
	return tea.Batch(cmds...)
}

func (m TextBody) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case UpdateSize:
		m.Style = m.Style.Width(m.Style.GetWidth() + msg.Width)
	case UpdateFocus:
		if msg.Name == "textbody" {
			m.State = Focus
			m.Style = m.Style.BorderForeground(lipgloss.Color("201"))
		} else {
			m.State = Blur
			m.Style = m.Style.BorderForeground(lipgloss.Color("255"))
		}

	case tea.KeyMsg:
		if m.State == Blur {
			return m, tea.Batch(cmds...)
		}

		s := msg.String()
		switch s {
		case "i":
			if m.State == Typing {
				m.Body, cmd = m.Body.Update(msg)
				cmds = append(cmds, cmd)
			} else {
				m.State = Typing
				cmds = append(cmds, m.Body.Focus())
			}
		case "esc":
			m.Body.Blur()
			m.State = Blur
		default:
			m.Body, cmd = m.Body.Update(msg)
			cmds = append(cmds, cmd)
		}

	default:
		m.Body, cmd = m.Body.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m TextBody) View() string {
	return m.Style.Render(m.Body.View())
}
