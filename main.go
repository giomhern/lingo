package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

func main() {
	m := newModel()
	p := tea.NewProgram(m)

	_, err := p.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

// model: app state

type Model struct {
	title string
	input textinput.Model
}

// newModel: initial model

func newModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Enter the name of your term"
	ti.Focus()
	return Model{
		title: "Hello, World!",
		input: ti,
	}
}

// init: kick off event loop

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

// update: handle new messages

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			v := m.input.Value()
			return m, handleQuerySearch(v)
		}
	}
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	s := m.input.View()
	return s
}


func handleQuerySearch