package components

import (
	"log"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MainModel struct {
	urlbar       Urlbar
	sidebar      SideBar
	Viewport     viewport.Model
	Ready        bool
	SideBarState state
	Style        lipgloss.Style
}

func InitialModel() MainModel {
	m := MainModel{
		Ready:        false,
		SideBarState: Open,
		Style:        lipgloss.NewStyle(),
	}

	return m
}

func (m MainModel) InitComponent(size tea.WindowSizeMsg) MainModel {
	updateSize := UpdateSize{Width: 50, Height: 0}

	m.urlbar = MakeUrlbar("ini url", "", size, updateSize, &m)
	m.sidebar = MakeSideBar(size, updateSize, &m)
	return m
}

func (m MainModel) HandleFocus() MainModel {
	var msg UpdateFocus
	if m.urlbar.State == Focus {
		msg = UpdateFocus{Name: "sidebar"}
	} else if m.sidebar.State == Focus {
		msg = UpdateFocus{Name: "urlbar"}
	}
	cmds := []tea.Cmd{}
	var cmd tea.Cmd
	mdl, cmd := m.sidebar.Update(msg)
	cmds = append(cmds, cmd)
	m.sidebar = mdl.(SideBar)
	mdl, cmd = m.urlbar.Update(msg)
	cmds = append(cmds, cmd)
	m.urlbar = mdl.(Urlbar)
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
		if !m.Ready {
			m.Viewport = viewport.New(msg.Width, msg.Height)
			m.Ready = true
			m = m.InitComponent(msg)
		} else {
			m.Viewport.Width = msg.Width
			m.Viewport.Height = msg.Height
		}
	case tea.KeyMsg:
		s := msg.String()
		switch s {
		case "ctrl+c":
			log.Println("----------")
			return m, tea.Quit
		case "ctrl+b":
			var width int
			if m.SideBarState == Open {
				m.SideBarState = Close
				width = 50
			} else {
				m.SideBarState = Open
				width = -50
			}
			m.urlbar.Update(UpdateSize{Width: width})
		case "up", "down":
			mdl, cmd := m.sidebar.Update(msg)
			cmds = append(cmds, cmd)
			m.sidebar = mdl.(SideBar)
		case "tab":
			m = m.HandleFocus()

			cmds = append(cmds, cmd)
		default:
			urlbarmodel, cmd := m.urlbar.Update(msg)
			cmds = append(cmds, cmd)
			m.urlbar = urlbarmodel.(Urlbar)

		}
	}

	return m, tea.Batch(cmds...)
}

func (m MainModel) View() string {
	if !m.Ready {
		return "initializing"
	}
	if m.SideBarState == Close {
		return m.urlbar.View()
	}
	return lipgloss.JoinHorizontal(lipgloss.Left, m.sidebar.View(), m.urlbar.View())
}
