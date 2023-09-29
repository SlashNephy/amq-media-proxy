package fs

// FakeFileSystem はテスト用の FileSystem 実装。
type FakeFileSystem struct {
	exists bool
}

func NewFakeFileSystem(exists bool) *FakeFileSystem {
	return &FakeFileSystem{
		exists: exists,
	}
}

func (fs *FakeFileSystem) Exists(_ string) (bool, error) {
	return fs.exists, nil
}

var _ FileSystem = (*FakeFileSystem)(nil)
