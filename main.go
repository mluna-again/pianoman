package main

// Alexander Craxton

// ------------------------------------------------
// Thank you for visiting https://asciiart.website/
// This ASCII pic can be found at
// https://asciiart.website/index.php?art=music/pianos

import (
	"log"
	"os"
	tea "github.com/charmbracelet/bubbletea"
)


func main() {
	program := tea.NewProgram(InitialModel())

	if _, err := program.Run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
