package slack

import (
	"fmt"
	"os"
	"path/filepath"
)

func NewFromPath(path string) (*MessageExport, error) {
	res := NewMessageExport()
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
		if !file.IsDir() {
			fname := filepath.Join(path, file.Name())
			content, err := os.ReadFile(fname)
			if err != nil {
				return nil, err
			}
			_ = res.AddRawMessages(content)
		}
	}
	return res, nil
}
