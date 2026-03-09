package app

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/aryanwalia2003/git-stats/internal/domain"
)

// Model holds the entire application state.
type Model struct {
	RepoName     string
	Loading      bool
	Ready        bool // true when viewport is sized
	Spinner      spinner.Model
	Viewport     viewport.Model
	Reader       domain.LocalGitReader
	Contributors []domain.Stat // top committers
	Commits      []domain.Stat // recent commits
	Churn        []domain.Stat // lines added/deleted
	History      []domain.Stat // full history for timeline
	ErrMsg       string        // error message if stats fail
}

// New creates a default application state with the reader wired in.
func New(name string, reader domain.LocalGitReader) Model {
	s := spinner.New()
	s.Spinner = spinner.Dot // animated dot spinner style
	return Model{
		RepoName: name,
		Loading:  true,
		Spinner:  s,
		Reader:   reader,
	}
}
