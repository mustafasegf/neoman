package components

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var SelectedItemsStyle lipgloss.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("201"))

const (
	Head = iota
	Tail
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

func MakeItems(name string, typ int) Items {
	return Items{
		Name:     name,
		Typ:      typ,
		Children: make([]*Items, 0),
		Selected: false,
	}
}

func NewItems(name string, typ int) *Items {
	return &Items{
		Name:     name,
		Typ:      typ,
		Children: make([]*Items, 0),
		Selected: false,
	}
}

func (i *Items) Last() *Items {
	next := i
	for next.Next != nil && next.Next.Typ != Tail {
		next = next.Next
	}
	return next
}

func (p *Items) Add(parent, n *Items) (next *Items) {
	n.Parent = parent
	parent.Children = append(parent.Children, n)

	if p.Next == nil {
		p.Next = n
		n.Prev = p
	} else {
		// n.Prev = p
		// n.Next = p.Next
		//
		// n.Next.Prev = n
		// p.Next = n

		p.Next.Prev = n.Last()
		n.Last().Next = p.Next

		p.Next = n
		n.Prev = p
	}

	return n.Last()
}

func (i *Items) SelectNext() {
	if i.Next == nil {
		return
	}

	if i.Next.Typ == Tail {
		return
	}

	i.Selected = false
	i.Next.Selected = true
	return
}

func (i *Items) SelectPrev() {
	if i.Prev == nil {
		return
	}

	if i.Prev.Typ == Head {
		return
	}

	i.Selected = false
	i.Prev.Selected = true
	return
}

func (m *Items) Init() tea.Cmd {
	var cmds []tea.Cmd
	return tea.Batch(cmds...)
}

func (m *Items) Update(msg tea.Msg) tea.Cmd {
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
				cmd := m.Next.Update(msg)
				cmds = append(cmds, cmd)
			}

		case "up":
			if m.Prev.Selected {
				m.Prev.SelectPrev()
			} else {
				cmd := m.Prev.Update(msg)
				cmds = append(cmds, cmd)
			}
		}
	}
	cmds = append(cmds, cmd)
	return tea.Batch(cmds...)
}

func (i Items) PrintTree(d int) string {
	if i.Typ == Tail {
		return ""
	}
	if i.Typ == Head {
		d--
	}

	var str string
	for _, child := range i.Children {
		str += child.PrintTree(d + 1)
	}

	if i.Typ == Head {
		return str
	}

	if i.Selected {
		return SelectedItemsStyle.Render(strings.Repeat("-", d)+i.Name) + "\n" + str
	}

	return strings.Repeat("-", d) + i.Name + "\n" + str
}

func (i Items) PrintHead() string {
	if i.Typ == Head {
		return i.Next.PrintHead()
	}

	if i.Typ == Tail {
		return ""
	}
	if i.Selected {
		return SelectedItemsStyle.Render(i.Name) + "\n" + i.Next.PrintHead()
	}
	return i.Name + "\n" + i.Next.PrintHead()
}

func (i Items) PrintTail() string {
	if i.Typ == Tail {
		return i.Prev.PrintTail()
	}

	if i.Typ == Head {
		return ""
	}

	if i.Selected {
		return SelectedItemsStyle.Render(i.Name) + "\n" + i.Prev.PrintTail()
	}
	return i.Name + "\n" + i.Prev.PrintTail()
}

type SideBar struct {
	Viewport viewport.Model
	Style    lipgloss.Style
	State    state
	Head     Items
	Tail     Items
}

func MakeSideBar(size tea.WindowSizeMsg, updateSize UpdateSize) SideBar {
	head := NewItems("", Head)

	tail := NewItems("", Tail)
	head.Add(head, tail)

	head.Add(head, NewItems("item 1", Req)).
		Add(head, NewItems("item 2", Req)).
		Add(head, NewItems("item 3", Req))

	f1 := NewItems("folder 1", Folder)
	f1.Add(f1, NewItems("item 4", Req)).
		Add(f1, NewItems("item 5", Req))

	tail.Prev.Add(head, f1).
		Add(head, NewItems("item 6", Req)).
		Add(head, NewItems("item 6", Req)).
		Add(head, NewItems("item 7", Req))

	head.Next.Selected = true

	m := SideBar{
		Style:  lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Width(updateSize.Width - 2).Height(size.Height - 2),
		Head:   *head,
		Tail:   *tail,
		State:  Blur,
	}
	return m
}

func (m SideBar) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, textinput.Blink)
	return tea.Batch(cmds...)
}

func (m SideBar) Update(msg tea.Msg) (SideBar, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case UpdateSize:
		m.Style = m.Style.MarginLeft(msg.Width).Width(m.Style.GetWidth() - msg.Width)
	case UpdateFocus:
		if msg.Name == "sidebar" {
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
		case "down":
			cmd := m.Head.Update(msg)
			cmds = append(cmds, cmd)

		case "up":
			cmd := m.Tail.Update(msg)
			cmds = append(cmds, cmd)
		}
	}
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m SideBar) View() string {
	return m.Style.Render(m.Head.PrintTree(0))
	// return m.Style.Render(m.Head.PrintHead() + "\n\n" + m.Tail.PrintTail())
}
