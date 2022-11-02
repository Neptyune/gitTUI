package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

func NewModel() (*model, error) {
	return model{}, nil
}

type model struct {
	nameInput string
	listInput string
	event     string
}

func main() {
	fmt.Println("Hello world")

	var _ tea.Model = (*model)(nil)
}
