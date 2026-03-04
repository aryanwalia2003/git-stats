package git 

import (
	"os/exec"
	"github.com/aryanwalia2003/git-stats/internal/domain"
)

func (r *Reader) GetLocalContributors() ([]domain.Stat, error) {
	cmd := exec.Command("git", "shortlog", "-sn", "--all")

	cmd.Dir = r.targetDir

	out, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	return parseContributorsStats(string(out)), nil
}
