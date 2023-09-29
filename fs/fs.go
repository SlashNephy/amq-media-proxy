package fs

// FileSystem はファイルシステムの操作を抽象化するインターフェイス。
type FileSystem interface {
	Exists(path string) (bool, error)
}
