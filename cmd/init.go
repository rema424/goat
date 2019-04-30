// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"unsafe"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"github.com/rema424/goat/lib"
)

type (
	// BaseInfo ...
	BaseInfo struct {
		Host      string
		User      string
		Project   string
		GoVersion string
	}
)

var (
	goatBasePash     = lib.GetImportPath()
	templateBasePath = lib.GetTemplateDirPath()
	baseInfo         = &BaseInfo{}
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := args[0]

		// 基本情報を取得
		getBaseInfo(arg)

		// アプリケーション名が入力されたら、プロジェクトディレクトリを作成する。
		if arg != "." {
			os.Mkdir(arg, 0755)
			os.Chdir(arg)
		}

		// ディレクトリ作成
		makeDirs()

		// ファイル作成
		makeFiles(templateBasePath, 0)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func makeDirs() {
	os.Mkdir(".circleci", 0755)
	os.Mkdir(".vscode", 0755)
	os.Mkdir("key", 0755)
	os.Mkdir("lib", 0755)
	os.Mkdir("model", 0755)
	os.Mkdir("service", 0755)
	os.Mkdir("repository", 0755)
	os.Mkdir("module", 0755)
	os.Mkdir("module/default", 0755)
	os.Mkdir("module/default/controller", 0755)
	os.Mkdir("module/default/main", 0755)
	os.Mkdir("module/default/main/public", 0755)
	os.Mkdir("module/default/main/static", 0755)
	os.Mkdir("module/default/main/static/css", 0755)
	os.Mkdir("module/default/main/static/css/src", 0755)
	os.Mkdir("module/default/main/static/css/src/user", 0755)
	os.Mkdir("module/default/main/static/css/src/vendor", 0755)
	os.Mkdir("module/default/main/static/js", 0755)
	os.Mkdir("module/default/main/static/js/src", 0755)
	os.Mkdir("module/default/main/static/js/src/user", 0755)
	os.Mkdir("module/default/main/static/js/src/vendor", 0755)
	os.Mkdir("module/default/main/template", 0755)
	os.Mkdir("module/default/main/template/user", 0755)
	os.Mkdir("module/default/main/viewmodel", 0755)
}

func makeFiles(dirPath string, indent int) {
	// ディレクトリを開く
	dir, err := os.Open(dirPath)
	if err != nil {
		panic(err)
	}

	// ディレクトリ内のファイル情報を一覧で取得する
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		panic(err)
	}

	// ファイル情報の一覧でループ処理
	for _, fileInfo := range fileInfos {
		fileName := fileInfo.Name()
		fileNameTrim := strings.Replace(fileName, ".tmpl", "", 1)
		indentStr := strings.Repeat("  ", indent)
		nextPath := filepath.Join(dirPath, fileName)

		if fileInfo.IsDir() {
			fmt.Printf("%s[Dir] %s\n", indentStr, fileNameTrim)
			// ディレクトリの場合は再帰的に処理
			makeFiles(nextPath, indent+1)
		} else {
			fmt.Printf("%s[File] %s\n", indentStr, fileNameTrim)
			rel, err := filepath.Rel(templateBasePath, nextPath)
			if err != nil {
				panic(err)
			}
			makeFile(rel, baseInfo)
		}
	}
}

func makeFile(relPath string, baseInfo *BaseInfo) {
	// 作成ファイル名
	n := strings.Replace(relPath, ".tmpl", "", 1)

	// ファイル作成
	f, err := os.Create(n)
	if err != nil {
		panic(err)
	}

	// テンプレートファイルの内容を作成ファイルに書き込み
	t := template.Must(template.ParseFiles(lib.GetTemplatePath(relPath)))
	if err := t.Execute(f, baseInfo); err != nil {
		panic(err)
	}
}

func getBaseInfo(project string) {
	host := getHost()
	user := getUser()
	pjt := getProject(project)

	if confirm(host, user, pjt) {
		baseInfo.Host = host
		baseInfo.User = user
		baseInfo.Project = pjt
		baseInfo.GoVersion = getGoVersion()
	} else {
		getBaseInfo(project)
	}
}

// go version go1.12.1 darwin/amd64 -> 1.12
func getGoVersion() string {
	out, err := exec.Command("go", "version").Output()
	if err != nil {
		panic(err)
	}

	// バイト列から文字列に変換する。これが一番速いらしい。
	outStr := *(*string)(unsafe.Pointer(&out)) // go version go1.12.1 darwin/amd64
	// 正規表現はリソース効率が悪いのでゴリゴリ整形する
	version := strings.Replace(strings.Split(outStr, " ")[2], "go", "", 1)
	return strings.Join(strings.Split(version, ".")[:2], ".")
}

func getHost() string {
	prompt := promptui.Select{
		Label: "Git Host",
		Items: []string{"github.com", "gitlab.com", "bitbucket.org", "other"},
	}

	_, host, err := prompt.Run()
	if err != nil {
		panic(err)
	}

	if host == "other" {
		s := "Git Host > "
		fmt.Print(s)

		scanner := bufio.NewScanner(os.Stdin)
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		for scanner.Scan() {
			host = strings.TrimSpace(scanner.Text())
			if host != "" {
				break
			} else {
				fmt.Print(s)
			}
		}
	}
	return host
}

func getUser() string {
	s := "User Name > "
	fmt.Print(s)

	var user string
	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	for scanner.Scan() {
		user = strings.TrimSpace(scanner.Text())
		if user != "" {
			break
		} else {
			fmt.Print(s)
		}
	}
	return user
}

func getProject(project string) string {
	var pjt string
	if project == "." {
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		pjt = filepath.Base(dir)
	} else {
		pjt = project
	}
	return pjt
}

func confirm(host string, user string, app string) bool {
	fmt.Println("【Confirm】")
	fmt.Println("host: ", host)
	fmt.Println("user: ", user)
	fmt.Println("project: ", app)
	fmt.Print("OK? [Y/n] ")

	var res bool

	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	for scanner.Scan() {
		input := scanner.Text()

		if input == "Y" || input == "y" || input == "" {
			res = true
			break
		} else if input == "N" || input == "n" {
			res = false
			break
		} else {
			fmt.Println("")
		}
	}

	return res
}
