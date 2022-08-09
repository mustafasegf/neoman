package components

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ResponseBody struct {
	Viewport viewport.Model
	State    state
	Parent   tea.Model
	Style    lipgloss.Style
}

func MakeResponseBody(body string, size tea.WindowSizeMsg, updateSize UpdateSize) ResponseBody {
	m := ResponseBody{
		State:    Blur,
		Style:    lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Width(size.Width - updateSize.Width - 2).Height(size.Height - 8),
		Viewport: viewport.New(size.Width-updateSize.Width-2, size.Height-8),
	}

	return m
}

func (m ResponseBody) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, textinput.Blink)
	return tea.Batch(cmds...)
}

func (m ResponseBody) Update(msg tea.Msg) (ResponseBody, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case UpdateSize:
		m.Style = m.Style.Width(m.Style.GetWidth() + msg.Width)

	case tea.WindowSizeMsg:
		m.Style = m.Style.Width(m.Style.GetWidth()).Height(msg.Height - 8)

	case UpdateFocus:
		if msg.Name == "responsebody" {
			m.State = Focus
			m.Style = m.Style.BorderForeground(lipgloss.Color("201"))
		} else {
			m.State = Blur
			m.Style = m.Style.BorderForeground(lipgloss.Color("255"))
		}

	}

	if m.State == Focus {
		m.Viewport, cmd = m.Viewport.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m ResponseBody) View() string {
	return m.Style.Render(m.Viewport.View())
}
