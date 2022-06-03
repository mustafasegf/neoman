package components

import tea "github.com/charmbracelet/bubbletea"

type state int

const (
	Typing state = iota
	Finished
)

const (
	Close state = iota
	Open
)

type UpdateSize tea.WindowSizeMsg
