package git 

import "github.com/aryanwalia2003/git-stats/internal/domain"

type Reader struct {
	targetDir string
}

func NewReader(dir string) *Reader{
	return &Reader{targetDir: dir}
}

var _ domain.LocalGitReader = (*Reader)(nil) //this means that our reader struct implements the LocalGitReader interface
