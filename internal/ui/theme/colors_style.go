package theme //yeh theme package ka hissa hai

import "github.com/charmbracelet/lipgloss" //importing lib

var (
	Primary   = lipgloss.Color("#8839ef")
	Secondary = lipgloss.Color("#209fb5")
	Error     = lipgloss.Color("#d20f39")
) // color scheme design kar rhe

var TitleStyle = lipgloss.NewStyle().
	Foreground(Primary).
	Bold(true).
	MarginLeft(2) 