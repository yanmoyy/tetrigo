package views

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/huh"
)

type mainKeyMap struct {
	Exit  key.Binding
	Up    key.Binding
	Down  key.Binding
	Enter key.Binding
}

func defaultMainKeyMap() *mainKeyMap {
	keys := &mainKeyMap{
		Exit:  key.NewBinding(key.WithKeys("esc"), key.WithHelp("escape", "exit")),
		Up:    key.NewBinding(key.WithKeys("up", "k", "ctrl+k", "ctrl+p"), key.WithHelp("↑", "up")),
		Down:  key.NewBinding(key.WithKeys("down", "j", "ctrl+j", "ctrl+n"), key.WithHelp("↓", "down")),
		Enter: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "enter")),
	}
	return keys
}

type lobbyKeyMap struct {
	Exit     key.Binding
	formKeys *huh.KeyMap
}

func defaultLobbyKeyMap() *lobbyKeyMap {
	keys := &lobbyKeyMap{
		Exit:     key.NewBinding(key.WithKeys("esc"), key.WithHelp("escape", "exit")),
		formKeys: huh.NewDefaultKeyMap(),
	}
	keys.formKeys.Quit.SetEnabled(false)
	return keys
}
