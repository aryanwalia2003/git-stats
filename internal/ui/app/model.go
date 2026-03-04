package app // this file sis part of the app package

import "github.com/charmbracelet/bubbles/spinner" 

// Model holds the application state.
type Model struct {
	RepoName string
	Loading  bool
	Spinner  spinner.Model
} // yeh state hai , where we have a reponame , loading which is a bool, and a spinner

// New creates a default application state.
func New(name string) Model {
	return Model{
		RepoName: name,
		Loading:  true,
		Spinner:  spinner.New(),
	}
} //new model ke liye sirf ek name pass karna hai just for the repo uske loading ki value default true hai and spinner toh external package hai 
