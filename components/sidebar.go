package components

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	Root = iota
	Folder
	Req
)

type Items struct {
	// for up and down navigation
	Name     string
	Typ      int
	Children []*Items
	Prev     *Items
	Next     *Items
	Parent   *Items
}

func NewItems(name string, typ int) *Items {
	return &Items{
		Name:     name,
		Typ:      typ,
		Children: []*Items{},
	}
}

// i = previous item
func (i *Items) Add(parent, child *Items) *Items {
	child.Parent = parent
	parent.Children = append(parent.Children, child)

	if parent.Next == nil {
		parent.Next = child
		child.Prev = parent
	} else {
		child.Prev = parent
		child.Next = parent.Next

		child.Next.Prev = child
		child.Prev.Next = child
	}

	return child
}

func (m Items) Init() tea.Cmd {
	var cmds []tea.Cmd
	return tea.Batch(cmds...)
}

func (m Items) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

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
	}

	f := NewItems("folder", Folder)
	f2 := NewItems("folder 2", Folder)
	f2.Add(f2, NewItems("item 6", Req))
	f.Add(f, NewItems("item 4", Req)).Add(f, NewItems("item 5", Req)).Add(f, f2)
	r.Add(r, NewItems("item 1", Req)).Add(r, NewItems("item 2", Req)).Add(r, NewItems("item 3", Req)).Add(r, f).Add(r, NewItems("item 7", Req))
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
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m SideBar) View() string {
	return m.Style.Render(m.Items.View())
}
