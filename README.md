# goat

## reference

- [spf13/cobra](https://github.com/spf13/cobra)

## memo

```bash
# Cobra-cli は GOPATH 配下でしか使えないので
cd $GOPATH/src/github.com/rema424

git clone https://github.com/rema424/goat.git

cd goat

rm -rf *

go get -u github.com/spf13/cobra/cobra

cobra -h

cobra init .

echo 'GO111MODULE=on' >> ~/.bash_profile

exec $SHELL -l

go mod init github.com/rema424/goat
# dep init

cabra add hello

go get github.com/rema424/goat
go install
```