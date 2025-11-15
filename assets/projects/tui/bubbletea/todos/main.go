package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// START OMIT

type TodoModel struct {
	todos  []string
	cursor int
	done   map[int]bool
}

func NewTodoModel() TodoModel {
	return TodoModel{
		todos: []string{"Vacuum", "Do the dishes", "Mop the floor"},
		done:  make(map[int]bool),
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
			done := tm.done[tm.cursor]
			if done {
				delete(tm.done, tm.cursor)
			} else {
				tm.done[tm.cursor] = true
			}
		}
	}
	return tm, nil
}

// VIEW OMIT

func (tm TodoModel) View() string {
	var b strings.Builder
	b.WriteString("Todos\n\n")
	for i, todo := range tm.todos {
		cursor := " "
		if tm.cursor == i {
			cursor = ">"
		}
		done := " "
		if ok := tm.done[i]; ok {
			done = "x"
		}
		b.WriteString(fmt.Sprintf("%s [%s] %s\n", cursor, done, todo))
	}

	b.WriteString("\nPress q to quit.\n")
	return b.String()
}

func main() {
	p := tea.NewProgram(NewTodoModel())
	_, _ = p.Run()
}

// END OMIT
