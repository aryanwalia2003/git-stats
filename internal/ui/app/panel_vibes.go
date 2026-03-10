package app

import (
	"fmt"
	"time"

	"github.com/aryanwalia2003/git-stats/internal/ui/theme"
)

// renderVibesPanel shows the fun behavioral stats in a single panel.
func renderVibesPanel(m Model) string {
	title := theme.PanelTitleStyle.Render("🎭 Developer Vibes")

	nightOwl := computeNightOwlIndex(m.Commits)
	weekend := computeWeekendWarrior(m.Commits)
	bestHour, _ := computeBusiestHour(m.Commits)
	bestDay, _ := computeBusiestDay(m.Commits)
	streak := computeStreak(m.Commits)
	avgMsg := computeAvgMessageLength(m.Commits)
	topEmoji, emojiCount := computeTopEmoji(m.Commits)
	caffeineActive, caffeineBurst := computeCaffeineMode(m.Commits)

	// Pad labels to 18 characters to perfectly align the values column
	lines := fmt.Sprintf(" %s %s %.0f%%\n", theme.LabelStyle.Render(fmt.Sprintf("%-18s", "🌙 Night Owl:")), renderProgressBar(nightOwl, 10, theme.CalendarColors[3]), nightOwl)
	lines += fmt.Sprintf(" %s %s %.0f%%\n\n", theme.LabelStyle.Render(fmt.Sprintf("%-18s", "⚔️  Weekend Warrior:")), renderProgressBar(weekend, 10, theme.CalendarColors[4]), weekend)
	
	lines += fmt.Sprintf(" %s %s %s\n",
		theme.LabelStyle.Render(fmt.Sprintf("%-18s", "🗓️  Busiest:")),
		theme.ValueStyle.Render(bestDay.String()[:3]),
		theme.SubtleStyle.Render(fmt.Sprintf("at %d:00", bestHour)))
	
	lines += fmt.Sprintf(" %s %s\n", 
		theme.LabelStyle.Render(fmt.Sprintf("%-18s", "🔥 Streak:")), 
		theme.ValueStyle.Render(fmt.Sprintf("%d days", streak)))
	
	lines += fmt.Sprintf(" %s %s\n\n", 
		theme.LabelStyle.Render(fmt.Sprintf("%-18s", "🗯️  Avg Message:")), 
		theme.ValueStyle.Render(fmt.Sprintf("%d chars", avgMsg)))

	if emojiCount > 0 {
		lines += fmt.Sprintf(" %s %s %s\n", 
			theme.LabelStyle.Render(fmt.Sprintf("%-18s", "👑 Top Emoji:")), 
			topEmoji, theme.SubtleStyle.Render(fmt.Sprintf("x%d", emojiCount)))
	} else {
		lines += fmt.Sprintf(" %s %s\n", theme.LabelStyle.Render(fmt.Sprintf("%-18s", "👑 Top Emoji:")), theme.SubtleStyle.Render("No emojis found"))
	}

	if caffeineActive {
		lines += fmt.Sprintf(" %s %s\n",
			theme.LabelStyle.Render(fmt.Sprintf("%-18s", "☕ Caffeine Mode:")),
			theme.ValueStyle.Render(fmt.Sprintf("ON! %d commits in 1h", caffeineBurst)))
	} else {
		lines += fmt.Sprintf(" %s %s\n", theme.LabelStyle.Render(fmt.Sprintf("%-18s", "☕ Caffeine Mode:")), theme.SubtleStyle.Render("off"))
	}

	_ = time.Now // ensure time import is used
	return title + "\n" + lines
}
