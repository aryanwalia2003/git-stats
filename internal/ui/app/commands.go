package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/aryanwalia2003/git-stats/internal/domain"
)

// fetchStats returns a tea.Cmd that loads all git stats in the background.
// The TUI stays responsive (spinner keeps spinning) while this runs.
func fetchStats(reader domain.LocalGitReader) tea.Cmd {
	return func() tea.Msg {
		contributors, err := reader.GetLocalContributors()
		if err != nil {
			return StatsErrorMsg{Err: err}
		}

		commits, err := reader.GetRecentCommits(50)
		if err != nil {
			return StatsErrorMsg{Err: err}
		}

		churn, err := reader.GetCodeChurn()
		if err != nil {
			return StatsErrorMsg{Err: err}
		}

		history, err := reader.GetCommitHistory()
		if err != nil {
			return StatsErrorMsg{Err: err}
		}

		merges, err := reader.GetMergeHistory()
		if err != nil {
			return StatsErrorMsg{Err: err}
		}

		files, err := reader.GetFileFrequencies()
		if err != nil {
			return StatsErrorMsg{Err: err}
		}

		return StatsLoadedMsg{
			Contributors: contributors,
			Commits:      commits,
			Churn:        churn,
			History:      history,
			Merges:       merges,
			Files:        files,
		}
	}
}
