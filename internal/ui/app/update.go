package app

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

// Update handles incoming messages and returns updated state + next command.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		if !m.Ready {
			m.Viewport = viewport.New(msg.Width, msg.Height)
			m.Ready = true
		} else {
			m.Viewport.Width = msg.Width
			m.Viewport.Height = msg.Height
		}
		if !m.Loading {
			m.Viewport.SetContent(m.renderDashboard())
		}
	case StatsLoadedMsg:
		m.Loading = false
		m.Contributors, m.Commits, m.Churn = msg.Contributors, msg.Commits, msg.Churn
		m.History, m.Merges = msg.History, msg.Merges
		m.Files = msg.Files
		if m.Ready {
			m.Viewport.SetContent(m.renderDashboard())
		}
	case StatsErrorMsg:
		m.Loading, m.ErrMsg = false, msg.Err.Error()
	}

	if m.Loading {
		var cmd tea.Cmd
		m.Spinner, cmd = m.Spinner.Update(msg)
		cmds = append(cmds, cmd)
	} else if m.Ready {
		var cmd tea.Cmd
		m.Viewport, cmd = m.Viewport.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
