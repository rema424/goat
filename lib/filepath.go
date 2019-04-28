package lib

import (
	// "go/build"
	"os"
)

// GetExecDir は実行バイナリの存在するパスを文字列で返却します。
func GetExecDir() string {
	edir, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return edir
}

// GetTemplateDir は template ディレクトリのパスを文字列で返却します。
func GetTemplateDir() string {
	return GetImportDir() + "/template"
}

// GetCurrDir はコマンド実行時のカレントディレクトリのパスを文字列で返却します。
func GetCurrDir() string {
	wdir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return wdir
}

// GetTemplate はテンプレートファイルの存在する絶対パスを文字列で返却します。
func GetTemplate(tmplName string) string {
	return GetTemplateDir() + "/" + tmplName
}

// GetImportDir は go get でダウンロードされたビルド前のソースコードの存在するパスを文字列で返却します。
func GetImportDir() string {
	// gopath := build.Default.GOPATH
	// gopath := os.ExpandEnv("${GOPATH}")
	// return gopath + "/src/github.com/rema424/goat"
	return os.ExpandEnv("${GOPATH}/src/github.com/rema424/goat")
}
