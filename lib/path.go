package lib

import (
	// "fmt"
	"os"
	"path/filepath"
)

// GetImportPath は go get の結果ソースコードが配置されたディレクトリのパスを返却します。
func GetImportPath() string {
	path := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "rema424", "goat")
	// fmt.Println("GetImportPath(): ", path)
	return path
}

// GetTemplatePath は引数で与えられてたテンプレートファイルの存在するパスを返却します。
func GetTemplatePath(tmplName string) string {
	path := filepath.Join(GetImportPath(), "template", tmplName)
	// fmt.Println("GetTemplatePath(): ", path)
	return path
}

// GetTemplateDirPath は引数で与えられてたテンプレートファイルの存在するパスを返却します。
func GetTemplateDirPath() string {
	path := filepath.Join(GetImportPath(), "template")
	// fmt.Println("GetTemplateDirPath(): ", path)
	return path
}
