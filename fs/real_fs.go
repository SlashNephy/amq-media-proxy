package fs

import "os"

// RealFileSystem は実際のファイルシステムにアクセスする FileSystem 実装。
type RealFileSystem struct {
	existsFn func(path string) (bool, error)
}

func NewRealFileSystem() *RealFileSystem {
	return &RealFileSystem{
		existsFn: func(path string) (bool, error) {
			_, err := os.Stat(path)
			return os.IsExist(err), err
		},
	}
}

func (fs *RealFileSystem) Exists(path string) (bool, error) {
	return fs.existsFn(path)
}

var _ FileSystem = (*RealFileSystem)(nil)
