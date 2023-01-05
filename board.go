package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type note []string

type keyboard struct {
	board   []note
	pressed int
	unselectCanceled bool
}

func InitialModel() keyboard {
	var notes []note

	notes = append(notes, first())
	notes = append(notes, second())
	notes = append(notes, third())
	notes = append(notes, fourth())
	notes = append(notes, fifth())
	notes = append(notes, sixth())
	notes = append(notes, seventh())

	return keyboard{
		board: notes,
		pressed: -1,
		unselectCanceled: false,
	}
}

func (model keyboard) Init() tea.Cmd {
	return nil
}

type unselectKeysEvent struct{}
func unselectKeys() tea.Msg {
	time.Sleep(time.Millisecond * 200)
	return unselectKeysEvent{}
}

func (model keyboard) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch message := message.(type) {
	case tea.KeyMsg:
		if !model.unselectCanceled {
			cmd = unselectKeys
		}

		model.unselectCanceled = true

		switch message.String() {
		case "ctrl+c":
			return model, tea.Quit

		case "d":
			model.pressed = 0

		case "f":
			model.pressed = 1

		case "g":
			model.pressed = 2

		case "h":
			model.pressed = 3

		case "j":
			model.pressed = 4

		case "k":
			model.pressed = 5

		case "l":
			model.pressed = 6
		}

	case unselectKeysEvent:
		if model.unselectCanceled {
			model.unselectCanceled = false
			return model, unselectKeys
		}
		model.pressed = -1
	}

	return model, cmd
}

func (model keyboard) View() string {
	str := ""

	for i := 0; i < 8; i++ {
		for _, note := range model.board {
			char := note[i]
			str += char
		}

		str += "\n"
	}

	for index := range model.board {
		selected := model.pressed == index

		if selected {
			str += "  "
			str += "↑"
			str += " "
		} else {
			str += "    "
		}
	}

	return str
}

