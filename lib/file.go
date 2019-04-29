package lib

import (
	"os"
)

// FileExist は与えられたパスのファイル・ディレクトリが存在するかどうかを判定します。
func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err == os.ErrNotExist {
		return false
	} else if err != nil {
		panic(err)
	} else {
		return true
	}
}

// IsDir は与えられたパスがディレクトリかどうかを判定します。
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	return info.IsDir()
}
