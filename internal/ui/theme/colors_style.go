package theme //yeh theme package ka hissa hai

import "github.com/charmbracelet/lipgloss" //importing lib

var (
	Primary   = lipgloss.Color("#8839ef")
	Secondary = lipgloss.Color("#209fb5")
	Error     = lipgloss.Color("#d20f39")
) // color scheme design kar rhe

var CalendarColors = []lipgloss.Color{
	lipgloss.Color("#11111b"), // 0 commits (very dark/barely visible bg)
	lipgloss.Color("#0e4429"), // 1 commit (dark green)
	lipgloss.Color("#006d32"), // 2 commits
	lipgloss.Color("#26a641"), // 3 commits
	lipgloss.Color("#39d353"), // 4+ commits (bright green)
}

var TitleStyle = lipgloss.NewStyle().
	Foreground(Primary).
	Bold(true).
	MarginLeft(2) 