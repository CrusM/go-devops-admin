package file

import (
	"os"
	"path/filepath"
	"strings"
)

// 写入文件
func Put(data []byte, to string) error {
	err := os.WriteFile(to, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// 判断文件是否存在
func Exists(fileToCheck string) bool {
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		return false
	}
	return true
}

func FileNameWithoutExtension(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}
