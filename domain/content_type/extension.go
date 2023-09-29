package content_type

import (
	"fmt"
	"path/filepath"
	"strings"
)

// ErrUnexpectedFileExtension は未知の拡張子が渡ってきたときに発生する。
var ErrUnexpectedFileExtension = fmt.Errorf("unexpected file extension")

// DetectContentTypeByFilename はファイル名の拡張子から Content-Type を特定する。
func DetectContentTypeByFilename(filename string) (ContentType, error) {
	filename = filepath.Base(filename)
	fileExtension := filepath.Ext(filename)
	fileExtension = strings.TrimPrefix(fileExtension, ".")
	fileExtension = strings.ToLower(fileExtension)

	switch fileExtension {
	case "mp3":
		return ContentTypeMP3, nil
	case "webm":
		return ContentTypeWebM, nil
	default:
		return "", fmt.Errorf("%w: got %s", ErrUnexpectedFileExtension, fileExtension)
	}
}
