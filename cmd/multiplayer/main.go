package main

import (
	"fmt"
	"log"

	"github.com/Broderick-Westrope/tetrigo/internal/multiplayer/config"
	"github.com/Broderick-Westrope/tetrigo/internal/multiplayer/page"
	"github.com/Broderick-Westrope/tetrigo/internal/multiplayer/starter"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if err := launchStarter(); err != nil {
		log.Fatal(err)
	}
}

func launchStarter() error {
	cfg, err := config.NewConfig()
	if err != nil {
		return fmt.Errorf("getting config: %w", err)
	}

	model, err := starter.NewModel(
		starter.NewInput(cfg, page.PageMain),
	)
	if err != nil {
		return fmt.Errorf("creating starter model: %w", err)
	}

	exitModel, err := tea.NewProgram(model, tea.WithAltScreen()).Run()
	if err != nil {
		return fmt.Errorf("failed to run program: %w", err)
	}

	typedExitModel, ok := exitModel.(*starter.Model)
	if !ok {
		return fmt.Errorf("failed to assert exit model type: %w", err)
	}

	if err = typedExitModel.ExitError; err != nil {
		return fmt.Errorf("starter model exited with an error: %w", err)
	}

	return nil
}
