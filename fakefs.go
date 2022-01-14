package fakefs

import (
	"github.com/spf13/afero"
)

type FakeFs struct {
	afero.Fs
}

func New() *FakeFs {
	return &FakeFs{afero.NewMemMapFs()}
}

func (fs *FakeFs) SymlinkIfPossible(oldname, newname string) error {
	return afero.WriteFile(fs, newname, []byte(oldname), 0755)
}
