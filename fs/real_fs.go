package fs

import (
	"os"

	"github.com/pkg/errors"
)

// RealFileSystem は実際のファイルシステムにアクセスする FileSystem 実装。
type RealFileSystem struct {
	existsFn func(path string) (bool, error)
}

func NewRealFileSystem() *RealFileSystem {
	return &RealFileSystem{
		existsFn: func(path string) (bool, error) {
			_, err := os.Stat(path)
			return os.IsExist(err), errors.WithStack(err)
		},
	}
}

func (fs *RealFileSystem) Exists(path string) (bool, error) {
	return fs.existsFn(path)
}

var _ FileSystem = (*RealFileSystem)(nil)
