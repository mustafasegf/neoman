package components

import tea "github.com/charmbracelet/bubbletea"

type state int

const (
	Typing state = iota
	Focus
	Blur
	Close
	Open
)

type UpdateSize tea.WindowSizeMsg

type UpdateFocus struct {
	Name string
}

type HttpRequestCmd struct {}
