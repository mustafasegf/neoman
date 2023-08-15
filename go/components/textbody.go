package components

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TextBody struct {
	Body     textinput.Model
	Viewport viewport.Model
	State    state
	Parent   tea.Model
	Style    lipgloss.Style
}

func MakeTextBody(body string, size tea.WindowSizeMsg, updateSize UpdateSize) TextBody {
	m := TextBody{
		Body:  textinput.New(),
		State: Blur,
		Style: lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Width(size.Width - 2 - updateSize.Width),
	}
	m.Body.SetValue(body)
	m.Body.Blur()

	return m
}

func (m TextBody) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, textinput.Blink)
	return tea.Batch(cmds...)
}

func (m TextBody) Update(msg tea.Msg) (TextBody, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case UpdateSize:
		m.Style = m.Style.Width(m.Style.GetWidth() + msg.Width)

	case tea.WindowSizeMsg:
		m.Style = m.Style.Width(m.Style.GetWidth())

	case UpdateFocus:
		if msg.Name == "textbody" {
			m.State = Focus
			m.Body.Focus()
			// cmds = append(cmds, cmd)
			m.Style = m.Style.BorderForeground(lipgloss.Color("201"))
		} else {
			m.State = Blur
			m.Body.Blur()
			m.Style = m.Style.BorderForeground(lipgloss.Color("255"))
		}

	case tea.KeyMsg:
		s := msg.String()
		switch s {
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
