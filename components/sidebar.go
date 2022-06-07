package components

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	SelectedItemsStyle lipgloss.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("201"))
)

const (
	Root = iota
	Folder
	Req
)

type Items struct {
	Name     string
	Typ      int
	Children []*Items
	Prev     *Items
	Next     *Items
	Parent   *Items
	Selected bool
}

func NewItems(name string, typ int) *Items {
	return &Items{
		Name:     name,
		Typ:      typ,
		Children: []*Items{},
		Selected: false,
	}
}

func (p *Items) Add(parent, n *Items) *Items {
	n.Parent = parent
	parent.Children = append(parent.Children, n)

	if p.Next == nil {
		p.Next = n
		n.Prev = p
	} else {
		n.Prev = p
		n.Next = p.Next

		n.Next.Prev = n
		n.Prev.Next = n
	}

	for n.Next != nil {
		n = n.Next
	}

	return n
}

func (i *Items) SelectNext() {
	if i.Next == nil {
		return
	}
	i.Selected = false
	i.Next.Selected = true
}

func (i *Items) SelectPrev() {
	if i.Prev == nil {
		return
	}
	if i.Prev.Typ == Root {
		return
	}

	i.Prev.Selected = true
	i.Selected = false
}
func (m Items) Init() tea.Cmd {
	var cmds []tea.Cmd
	return tea.Batch(cmds...)
}

func (m *Items) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		s := msg.String()
		switch s {
		case "down":
			if m.Next.Selected {
				m.Next.SelectNext()
			} else {
				m.Next.Update(msg)
			}

		case "up":
			if m.Next.Selected {
				m.Next.SelectPrev()
			} else {
				m.Next.Update(msg)
			}
		}
	}
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (i Items) PrintTree(d int) string {
	if i.Typ == Root {
		d--
	}
	var str string
	for _, child := range i.Children {
		str += child.PrintTree(d + 1)
	}
	if i.Typ == Root {
		return str
	}

	if i.Selected {
		return SelectedItemsStyle.Render(strings.Repeat("-", d)+i.Name) + "\n" + str
	}
	return strings.Repeat("-", d) + i.Name + "\n" + str
}

func (m Items) View() string {
	return m.PrintTree(0)
}

type SideBar struct {
	Viewport viewport.Model
	Parent   tea.Model
	Style    lipgloss.Style
	State    state
	Items    *Items
}

func MakeSideBar(size tea.WindowSizeMsg, updateSize UpdateSize, parent tea.Model) SideBar {
	r := NewItems("", Root)
	m := SideBar{
		Parent: parent,
		Style:  lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Width(updateSize.Width - 2).Height(size.Height - 2),
		Items:  r,
		State:  Focus,
	}

	f := NewItems("folder", Folder)
	f2 := NewItems("folder 2", Folder)
	f2.Add(f2, NewItems("item 6", Req))
	f.Add(f, NewItems("item 4", Req)).Add(f, NewItems("item 5", Req)).Add(f, f2)
	r.Add(r, NewItems("item 1", Req)).Add(r, NewItems("item 2", Req)).Add(r, NewItems("item 3", Req)).Add(r, f).Add(r, NewItems("item 7", Req))
	r.Next.Selected = true

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
	case UpdateFocus:
		if msg.Name == "sidebar" {
			m.State = Focus
		} else {
			m.State = Blur
		}
	case tea.KeyMsg:
		s := msg.String()
		if m.State == Blur {
			return m, tea.Batch(cmds...)
		}
		switch s {
		case "down", "up":
			m.Items.Update(msg)
		}
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m SideBar) View() string {
	return m.Style.Render(m.Items.View())
}
