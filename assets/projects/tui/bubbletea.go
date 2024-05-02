package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// START OMIT

type todoModel struct {
	todos  []string
	cursor int
	done   map[int]struct{}
}

func newTodoModel() todoModel {
	return todoModel{
		todos: []string{"Vacuum", "Do the dishes", "Mop the floor"},
		done:  make(map[int]struct{}),
	}
}

func (m todoModel) Init() tea.Cmd {
	return nil
}

// UPDATE OMIT

func (m todoModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.todos)-1 {
				m.cursor++
			}
		case "enter":
			_, ok := m.done[m.cursor]
			if ok {
				delete(m.done, m.cursor)
			} else {
				m.done[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

// VIEW OMIT

func (m todoModel) View() string {
	s := "Todos\n\n"
	for i, todo := range m.todos {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		done := " "
		if _, ok := m.done[i]; ok {
			done = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, done, todo)
	}

	s += "\nPress q to quit.\n"
	return s
}

func main() {
	p := tea.NewProgram(newTodoModel())
	p.Run()
}

// END OMIT
