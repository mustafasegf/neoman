package components

import (
	"encoding/json"
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
	responsebody ResponseBody
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
	m.urlbar = MakeUrlbar("https://randomuser.me/api/", "", size, updateSize)
	m.textbody = MakeTextBody("", size, updateSize)
	m.responsebody = MakeResponseBody("", size, updateSize)
	m.focus = "urlbar"
	return m
}

func (m MainModel) HandleFocus() (MainModel, tea.Cmd) {
	var msg UpdateFocus
	if m.sidebar.State == Focus {
		msg = UpdateFocus{Name: "urlbar"}
		m.focus = "urlbar"
	} else if m.urlbar.State == Focus {
		msg = UpdateFocus{Name: "textbody"}
		m.focus = "textbody"
	} else if m.textbody.State == Focus {
		msg = UpdateFocus{Name: "responsebody"}
		m.focus = "responsebody"
	} else if m.responsebody.State == Focus {
		msg = UpdateFocus{Name: "sidebar"}
		m.focus = "sidebar"
	}

	cmd := m.UpdateAll(msg)

	return m, cmd
}

func (m *MainModel) UpdateAll(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	m.urlbar, cmd = m.urlbar.Update(msg)
	cmds = append(cmds, cmd)

	m.sidebar, cmd = m.sidebar.Update(msg)
	cmds = append(cmds, cmd)

	m.textbody, cmd = m.textbody.Update(msg)
	cmds = append(cmds, cmd)

	m.responsebody, cmd = m.responsebody.Update(msg)
	cmds = append(cmds, cmd)

	return tea.Batch(cmds...)
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
			cmd = m.UpdateAll(msg)
			cmds = append(cmds, cmd)
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
			m.textbody.Update(UpdateSize{Width: width})
		case "tab":
			m, cmd = m.HandleFocus()
			cmds = append(cmds, cmd)
		default:
			cmd = m.UpdateAll(msg)
			cmds = append(cmds, cmd)
		}

	case HttpRequestCmd:
		m.HttpRequest()

	default:
		cmd = m.UpdateAll(msg)
		cmds = append(cmds, cmd)

	}

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
	var resMap map[string]interface{}
	json.Unmarshal(body, &resMap)

	resByte, err := json.MarshalIndent(resMap, "", "  ")
	if err != nil {
		log.Println("http req", err)
	}
	m.responsebody.Viewport.SetContent(string(resByte))
}

func (m MainModel) View() string {
	if !m.Ready {
		return "initializing"
	}
	if m.SideBarState == Close {
		return m.urlbar.View()
	}
	return lipgloss.JoinHorizontal(lipgloss.Left, m.sidebar.View(),
		lipgloss.JoinVertical(lipgloss.Bottom,
			m.urlbar.View(),
			m.textbody.View(),
			m.responsebody.View(),
		),
	)
}
