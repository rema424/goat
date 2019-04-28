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
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"

	"github.com/rema424/goat/lib"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")

		// テンプレートが配置されたパスを文字列で取得する（go get で配置された GOPATH 配下のパス）
		tmplPath := lib.GetTemplatePath("hello.go.tmpl")

		// テンプレートを読み込む
		t := template.Must(template.ParseFiles(tmplPath))

		// ファイルをコマンド実行時のカレントディレクトリに作成する
		f, err := os.Create("test.txt")
		if err != nil {
			panic(err)
		}

		// ファイルにテンプレートの内容を書き込む
		err = t.Execute(f, "gadddgaaddgaagdaa")
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	// helloCmd.AddCommand(addCmd)
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
