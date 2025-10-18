package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// START OMIT

type TodoModel struct {
	todos  []string
	cursor int
	done   map[int]struct{}
}

func NewTodoModel() TodoModel {
	return TodoModel{
		todos: []string{"Vacuum", "Do the dishes", "Mop the floor"},
		done:  make(map[int]struct{}),
	}
}

func (tm TodoModel) Init() tea.Cmd {
	return nil
}

// UPDATE OMIT

func (tm TodoModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return tm, tea.Quit
		case "up":
			if tm.cursor > 0 {
				tm.cursor--
			}
		case "down":
			if tm.cursor < len(tm.todos)-1 {
				tm.cursor++
			}
		case "enter":
			_, ok := tm.done[tm.cursor]
			if ok {
				delete(tm.done, tm.cursor)
			} else {
				tm.done[tm.cursor] = struct{}{}
			}
		}
	}
	return tm, nil
}

// VIEW OMIT

func (tm TodoModel) View() string {
	s := "Todos\n\n"
	for i, todo := range tm.todos {
		cursor := " "
		if tm.cursor == i {
			cursor = ">"
		}

		done := " "
		if _, ok := tm.done[i]; ok {
			done = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, done, todo)
	}

	s += "\nPress q to quit.\n"
	return s
}

func main() {
	p := tea.NewProgram(NewTodoModel())
	_, _ = p.Run()
}

// END OMIT
