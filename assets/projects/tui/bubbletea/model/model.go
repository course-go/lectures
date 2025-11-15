package model

import tea "github.com/charmbracelet/bubbletea"

// START OMIT

type Model interface {
	Init() tea.Cmd
	Update(tea.Msg) (Model, tea.Cmd)
	View() string
}

// END OMIT
