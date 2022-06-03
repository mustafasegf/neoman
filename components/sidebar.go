package components

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SideBar struct {
	Url      textinput.Model
	Viewport viewport.Model
	Parent   tea.Model
	Style    lipgloss.Style
	State    state
}

func MakeSideBar(size tea.WindowSizeMsg, updateSize UpdateSize, parent tea.Model) SideBar {
	m := SideBar{
		Url:    textinput.New(),
		Parent: parent,
		Style:  lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Width(updateSize.Width - 2).Height(size.Height - 2),
	}

	return m
}

func (m SideBar) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, textinput.Blink)
	return tea.Batch(cmds...)
}

func (m SideBar) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case UpdateSize:
		m.Style = m.Style.MarginLeft(msg.Width).Width(m.Style.GetWidth() - msg.Width)

	default:
		m.Url, cmd = m.Url.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m SideBar) View() string {
	return m.Style.Render(m.Url.View())
}
