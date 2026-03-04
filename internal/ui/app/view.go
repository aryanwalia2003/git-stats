package app //yeh file app package ka part hai , go mein packages ko logically group karte hain

import (
	"fmt" //fmt printf , sprintf ke liye use  hota hai standard package hai
	"github.com/aryanwalia2003/gh-stats/internal/ui/theme"
)

// View renders the current state as a string.
func (m Model) View() string {
	
	if m.Loading { //checks if app abhi data load kar rha hai yaan nhi 
		return fmt.Sprintf("%s Fetching %s...", 
            m.Spinner.View(), 
            m.RepoName)
	}

	return theme.TitleStyle.Render("Stats for " + m.RepoName)
}


//basically chm ek state hai , check karo if m is still loading if it is so , then show a spinner icon and write a string saying fetching string for the repo blah , agar loading nhi hai toh do show the string stats for the rep 
