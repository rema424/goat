package lib

import (
	"os"
)

// GetImportPath は go get でダウンロードされたビルド前のソースコードの存在するパスを文字列で返却します。
func GetImportPath() string {
	return os.ExpandEnv("${GOPATH}/src/github.com/rema424/goat")
}

// GetTemplateDirPath は template ディレクトリのパスを文字列で返却します。
func GetTemplateDirPath() string {
	return GetImportPath() + "/template"
}

// GetTemplatePath はテンプレートファイルの存在する絶対パスを文字列で返却します。
func GetTemplatePath(tmplName string) string {
	return GetTemplateDirPath() + "/" + tmplName
}
