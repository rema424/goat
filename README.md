# goat

## reference

- [spf13/cobra](https://github.com/spf13/cobra)
- [Go初心者がcobraを使ってコマンドラインツールを作ってみた話](https://blog.engineer.adways.net/entry/advent_calendar_2018/18)
- [Go Modulesの概要とGo1.12に含まれるModulesに関する変更](https://budougumi0617.github.io/2019/02/15/go-modules-on-go112/)
- [Cobra の使い方とテスト](https://text.baldanders.info/golang/using-and-testing-cobra/)
- [go get コマンドでパッケージを管理する](https://text.baldanders.info/golang/go-get-package/)
- [strings パッケージ](http://golang.jp/pkg/strings)
- [CLI作成支援パッケージ Cobra を使い、Go 言語でコマンドラインツールを作ってみる](https://qiita.com/kent_ocean/items/eb518c0816addd69f353)
- [GolangでwebサービスのAPIを叩くCLIツールを作ろう](https://qiita.com/minamijoyo/items/cfd22e9e6d3581c5d81f)
- [Golangのコマンドライブラリcobraを使って少しうまく実装する](https://qiita.com/tkit/items/3cdeafcde2bd98612428)
- [VSCodeでGoのModules設定](https://qiita.com/msmsny/items/a8d4573d03774815a198)
- [Create a Golang CLI application with Cobra and goxc](https://sbstjn.com/create-golang-cli-application-with-cobra-and-goxc.html)
- [mattn/echo-scaffold](https://github.com/mattn/echo-scaffold/blob/master/command/model_command.go)
- [go-openapi/inflect](https://github.com/go-openapi/inflect)
- [GoでCLIツール作るのに便利そうなパッケージを集めてみました](https://qiita.com/isaoshimizu/items/71dd2ca2a08ddb607e31)
- [manifoldco/promptui](https://github.com/manifoldco/promptui)
- [Go 1.11 の Modules (vgo) を CircleCI で使う](https://blog.tsub.me/post/go111-modules-in-circleci/)
- [Go moduleを使ったCircleCIのconfig.yml](https://qiita.com/tanden/items/2134c2650f18406b1a1d)
- [AE/Go アプリケーションを Go 1.11 に移行するためにやったこと](https://blog.a-know.me/entry/2018/10/28/215508)
- []()
- []()
- []()
- []()

## memo

```bash
# cobra-cli は GOPATH 配下でしか使えないので GOPATH 配下で開発する
cd $GOPATH/src/github.com/rema424

# リモートリポジトリをクローンする
git clone https://github.com/rema424/goat.git

# ディレクトリ移動
cd goat

# cobra-cli でプロジェクトを作成するためにはディレクトリが空である必要がある
rm -rf *

# cobra-cli をインストール
go get -u github.com/spf13/cobra/cobra

# ヘルプを表示してみる
cobra -h

# プロジェクトをカレントディレクトリに初期化する
cobra init .

# GOPATH 配下でモジュール対応モードを使うために環境変数を設定する
echo 'GO111MODULE=on' >> ~/.bash_profile

# 環境変数を読み込ませるためにシェルを再起動
exec $SHELL -l

# モジュールファイルを作成する
go mod init github.com/rema424/goat
# dep init

# cobra でサブコマンドを追加する
cobra add hello

# バイナリとしてインストールするためには一度リモートに上げてから go get するか、ローカルで go install する
go get github.com/rema424/goat
# go install

# 実行してみる
goat
```