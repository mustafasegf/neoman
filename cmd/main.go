package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mustafasegf/neoman/components"
	"github.com/mustafasegf/neoman/internal/logger"
)

func main() {
	logger.SetLogger()
	if err := tea.NewProgram(components.InitialModel(), tea.WithAltScreen()).Start(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}
