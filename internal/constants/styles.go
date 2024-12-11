package constants

import (
	"github.com/charmbracelet/lipgloss"
)

/*STYLING*/

var (
	DocStyle      = lipgloss.NewStyle().Padding(1).Margin(10, 1, 0, 1).BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#C06E52")).Align(lipgloss.Center, lipgloss.Center)
	SenderStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("5"))
	ErrStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#bd534b")).Render
	LLMNameStyle  = lipgloss.NewStyle().Background(lipgloss.Color("#B084CC")).Foreground(lipgloss.Color("#4D243D"))
	MsgInputStyle = lipgloss.NewStyle().Padding(1).BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#FFF275"))
	MsgViewStyle  = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#92D5E6"))

	UnfocusedModelStyle = lipgloss.NewStyle().Padding(1)

	FocusedModelStyle = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#92D5E6"))
)
