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

func MakeUrlbar(url string, title string, size tea.WindowSizeMsg, updateSize UpdateSize) Urlbar {
	m := Urlbar{
		Url:   textinput.New(),
		Title: title,
		State: Focus,
		Style: lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("201")).Width(size.Width - 2 - updateSize.Width),
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

func (m Urlbar) Update(msg tea.Msg) (Urlbar, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case UpdateSize:
		m.Style = m.Style.Width(m.Style.GetWidth() + msg.Width)

	case tea.WindowSizeMsg:
		m.Style = m.Style.Width(m.Style.GetWidth())
	// case tea.WindowSizeMsg:
	// 	m.Style = m.Style.Width(m.Style.GetWidth() + msg.Width).Height(msg.Height)

	case UpdateFocus:
		if msg.Name == "urlbar" {
			m.State = Focus
			m.Url.Focus()
			m.Style = m.Style.BorderForeground(lipgloss.Color("201"))
		} else {
			m.State = Blur
			m.Url.Blur()
			m.Style = m.Style.BorderForeground(lipgloss.Color("255"))
		}

	case tea.KeyMsg:
		s := msg.String()
		switch s {
		case "enter":
			cmd = func() tea.Msg {
				return HttpRequestCmd{}
			}
			cmds = append(cmds, cmd)
		}
	}
	m.Url, cmd = m.Url.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Urlbar) View() string {
	return m.Style.Render(m.Url.View())
}
