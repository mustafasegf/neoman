package components

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MainModel struct {
	urlbar       Urlbar
	sidebar      SideBar
	textbody     TextBody
	Viewport     viewport.Model
	Ready        bool
	focus        string
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

	m.sidebar = MakeSideBar(size, updateSize)
	m.urlbar = MakeUrlbar("https://jsonplaceholder.typicode.com/todos/1", "", size, updateSize)
	m.textbody = MakeTextBody("", size, updateSize)
	m.focus = "urlbar"
	return m
}

func (m MainModel) HandleFocus() MainModel {
	var msg UpdateFocus
	if m.sidebar.State == Focus {
		msg = UpdateFocus{Name: "urlbar"}
		m.focus = "urlbar"
	} else if m.urlbar.State == Focus {
		msg = UpdateFocus{Name: "textbody"}
		m.focus = "textbody"
	} else if m.textbody.State == Focus {
		msg = UpdateFocus{Name: "sidebar"}
		m.focus = "sidebar"
	}
	var cmd tea.Cmd
	var cmds []tea.Cmd

	m.urlbar, cmd = m.urlbar.Update(msg)
	cmds = append(cmds, cmd)

	m.sidebar, cmd = m.sidebar.Update(msg)
	cmds = append(cmds, cmd)

	m.textbody, cmd = m.textbody.Update(msg)
	cmds = append(cmds, cmd)

	return m
}

func (m MainModel) Init() tea.Cmd {
	var cmds []tea.Cmd

	cmds = append(cmds, m.urlbar.Init())
	cmds = append(cmds, m.sidebar.Init())
	cmds = append(cmds, m.textbody.Init())
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
			m.sidebar, cmd = m.sidebar.Update(msg)
			cmds = append(cmds, cmd)
		case "tab":
			m = m.HandleFocus()

			cmds = append(cmds, cmd)
		}

	case HttpRequestCmd:
		m.HttpRequest()
	}

	m.urlbar, cmd = m.urlbar.Update(msg)
	cmds = append(cmds, cmd)

	m.sidebar, cmd = m.sidebar.Update(msg)
	cmds = append(cmds, cmd)

	m.textbody, cmd = m.textbody.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m *MainModel) HttpRequest() {
	res, err := http.Get(m.urlbar.Url.Value())
	if err != nil {
		log.Println("http req", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("http req", err)
	}
	m.textbody.Body.SetValue(string(body))
}

func (m MainModel) View() string {
	if !m.Ready {
		return "initializing"
	}
	if m.SideBarState == Close {
		return m.urlbar.View()
	}
	return lipgloss.JoinHorizontal(lipgloss.Left, m.sidebar.View(), lipgloss.JoinVertical(lipgloss.Bottom, m.urlbar.View(), m.textbody.View()))
}
