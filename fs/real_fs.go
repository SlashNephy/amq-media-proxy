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
			if _, err := os.Stat(path); err != nil {
				if os.IsNotExist(err) {
					return false, nil
				}

				return false, errors.WithStack(err)
			}

			return true, nil
		},
	}
}

func (fs *RealFileSystem) Exists(path string) (bool, error) {
	return fs.existsFn(path)
}

var _ FileSystem = (*RealFileSystem)(nil)
