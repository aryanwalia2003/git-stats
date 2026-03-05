package git

import (
	"os/exec" // os/exec lets us run terminal commands like "git log"
	"strings" // strings package has helpers to cut/trim/split text
)

// runGit is a small helper that runs any git command in our target directory.
// It returns the trimmed output as a string, or an error if the command fails.
// This keeps other files clean — they just call r.runGit("log", "--oneline")
func (r *Reader) runGit(args ...string) (string, error) {
	cmd := exec.Command("git", args...) // create the command: git <args>
	cmd.Dir = r.targetDir               // run it inside the repo folder

	out, err := cmd.Output() // execute and capture stdout
	if err != nil {
		return "", err // return empty string + error if command failed
	}

	return strings.TrimSpace(string(out)), nil // trim newlines and return
}
