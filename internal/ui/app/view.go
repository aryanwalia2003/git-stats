package app

import (
	"fmt"
	"github.com/aryanwalia2003/git-stats/internal/ui/theme"
)

// View renders the root application UI state.
func (m Model) View() string {
	if m.ErrMsg != "" {
		return fmt.Sprintf("\n  ❌ Error: %s\n", m.ErrMsg)
	}
	if m.Loading {
		return fmt.Sprintf("\n  %s Fetching stats for %s...\n",
			m.Spinner.View(),
			theme.TitleStyle.Render(m.RepoName))
	}
	if !m.Ready {
		return "\n  Initializing viewport...\n"
	}

	// Render the scrollable viewport!
	return m.Viewport.View()
}
