package lib

import (
	"os"
)

// GetImportPath は go get の結果ソースコードが配置されたディレクトリのパスを返却します。
func GetImportPath() string {
	return os.ExpandEnv("${GOPATH}/src/github.com/rema424/goat")
}

// GetTemplatePath は引数で与えられてたテンプレートファイルの存在するパスを返却します。
func GetTemplatePath(tmplName string) string {
	return GetImportPath() + "/template/" + tmplName
}
