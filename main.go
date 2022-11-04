package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type model struct {
	nameInput string
	listInput string
	event     string
	count     int //added count field
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	//if message us a key pressed
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "up":
			m.count++
		case "down":
			m.count++

		}
	}
	return m, nil
}
func (m model) View() string {
	return fmt.Sprint("count: ", m.count, "\n\n INCREASE      DECREASE")
}

func (m model) Init() tea.Cmd {
	return nil
}

func main() {
	err := tea.NewProgram(&model{}, tea.WithAltScreen()).Start()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	//p := tea.NewProgram(initialModel())
	//if _, err := p.Run(); err != nil {
	//	fmt.Printf("Alas, there's been an error: %v", err)
	//	os.Exit(1)
	//}
}
