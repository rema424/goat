# goat

## reference

- [spf13/cobra](https://github.com/spf13/cobra)

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