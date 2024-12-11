package constants

import (
	"github.com/charmbracelet/bubbles/key"
)

type keymap struct {
	Back  key.Binding
	Enter key.Binding
	Quit  key.Binding
	Focus key.Binding
}

var Keymap = keymap{
	Back: key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "back")),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit program"),
	),
	Focus: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "switch focus between chat input and chat view"),
	),
}
