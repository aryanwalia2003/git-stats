package app

import tea "github.com/charmbracelet/bubbletea"

// Update handles user input and system messages.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) { // current state of repo ko leta hai as a param , which brings reponame, loading and spinner with it, returns an updated new state , and command , for eg quit program, start action A, call api etc
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.Spinner, cmd = m.Spinner.Update(msg)
	return m, cmd
}

func (m Model) Init() tea.Cmd {
	return m.Spinner.Tick
}

// NOTE TO SELF : this Update function is a function of Model 

//func (receiver) name(parameters) returns

//func (s Server) Start(port int) error ==> Server ka start method hai

//class Model:
//	def update(msg tea.msg):
//		pass


//above is the python equivalent
