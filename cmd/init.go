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
	"path/filepath"
	"strings"
	"text/template"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"github.com/rema424/goat/lib"
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
		fmt.Println("init called")
		fmt.Println("Arts Count: ", len(args))
		arg := args[0]

		// var host, user, app string
		// for {
		// 	host = getHost()
		// 	user = getUser()
		// 	app = getApp(arg)
		// 	ok := confirm(host, user, app)
		// 	if ok {
		// 		break
		// 	}
		// }

		// アプリケーション名が入力されたら、プロジェクトディレクトリを作成する。
		if arg != "." {
			os.Mkdir(arg, 0755)
			os.Chdir(arg)
		}

		// ディレクトリ作成
		makeDirs()

		// ファイル作成
		makeFiles("github.com", "rema424", "example")
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

func makeFiles(host string, user string, project string) {
	m := map[string]string{
		"host":    host,
		"user":    user,
		"project": project,
	}
	fmt.Println(m)
	listFiles(lib.GetTemplateDirPath(), 0)
}

func listFiles(absDirPath string, indent int) {
	dir, err := os.Open(absDirPath)
	if err != nil {
		panic(err)
	}
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		panic(err)
	}
	tmplBasePath := lib.GetTemplateDirPath()
	baseInfo := &BaseInfo{
		Host:    "github.com",
		User:    "rema424",
		Project: "goat-test",
	}
	for _, fileInfo := range fileInfos {
		fileName := fileInfo.Name()
		indentStr := strings.Repeat("  ", indent)
		nextAbsPath := filepath.Join(absDirPath, fileName)
		if fileInfo.IsDir() {
			fmt.Printf("%s[Dir] %s\n", indentStr, fileName)
			listFiles(nextAbsPath, indent+1)
		} else {
			fmt.Printf("%s[File] %s\n", indentStr, fileInfo.Name())
			tmplRelPath := strings.Replace(nextAbsPath, tmplBasePath, "", 1)
			makeDefaultFile(baseInfo, tmplRelPath)
		}
	}
}

type (
	// BaseInfo ...
	BaseInfo struct {
		Host    string
		User    string
		Project string
	}
)

func makeDefaultFile(baseInfo *BaseInfo, relPath string) {
	// fmt.Println(relPath)
	moveTo := "." + strings.Replace(relPath, ".tmpl", "", 1)
	moveFrom := filepath.Base(moveTo)
	f, err := os.Create(moveFrom)
	if err != nil {
		panic(err)
	}

	t := template.Must(template.ParseFiles(lib.GetTemplatePath(relPath)))
	err = t.Execute(f, baseInfo)
	if err != nil {
		panic(err)
	}

	err = os.Rename(moveFrom, moveTo)
	if err != nil {
		panic(err)
	}
}

func makeGitignore() {
	tmplPath := lib.GetTemplatePath("gitignore.go.tmpl")

	// テンプレートを読み込む
	t := template.Must(template.ParseFiles(tmplPath))

	// ファイルをコマンド実行時のカレントディレクトリに作成する
	f, err := os.Create(".gitignore")
	if err != nil {
		panic(err)
	}

	// ファイルにテンプレートの内容を書き込む
	err = t.Execute(f, nil)
	if err != nil {
		panic(err)
	}
}

func makeCircleCiConfig() {

}

func getHost() string {
	prompt := promptui.Select{
		Label: "リモートリポジトリのホスティング先を選択してください。",
		Items: []string{"github.com", "gitlab.com", "bitbucket.org", "other"},
	}

	_, host, err := prompt.Run()
	if err != nil {
		panic(err)
	}

	if host == "other" {
		fmt.Print("リモートリポジトリのホスティング先を入力してください > ")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			host = scanner.Text()
			if host != "" {
				break
			} else {
				fmt.Print("リモートリポジトリのホスティング先を入力してください > ")
			}
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}

	fmt.Println("")
	return host
}

func getUser() string {
	fmt.Println("リモートリポジトリのユーザ名または組織名を入力してください。")
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	var user string
	for scanner.Scan() {
		user = scanner.Text()
		if user != "" {
			break
		} else {
			fmt.Println("ユーザ名または組織名を入力してください。")
			fmt.Print("> ")
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("")
	return user
}

func getApp(arg string) string {
	var appName string
	if arg == "." {
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		appName = filepath.Base(dir)
	} else {
		appName = arg
	}
	return appName
}

func confirm(host string, user string, app string) bool {

	var res bool

	fmt.Println("【入力確認】")
	fmt.Println("ホスト先: ", host)
	fmt.Println("ユーザ名: ", user)
	fmt.Println("プロジェクト名: ", app)
	fmt.Println("パス: ")
	fmt.Print("入力内容はこちらで間違いないですか? [Y/n] ")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		input := scanner.Text()

		if input == "Y" || input == "y" || input == "" {
			res = true
			break
		} else if input == "N" || input == "n" {
			res = false
			break
		} else {
			fmt.Println("yかnで答えてください。")
			fmt.Println("")
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return res
}
