package main

import (
	"fmt"
	"os"

	"github.com/alfiehiscox/spqr-api/scraper"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(scraper.NewModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Oops something went wrong")
		os.Exit(1)
	}
}
