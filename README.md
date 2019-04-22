# サポーターズ勉強会資料

## 必要な環境
- Golang
- Docker


## Golangのインストール
``` bash
$ brew install go
$ vi ~/.bash_profile
```

~/.bash_profile
```
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:/usr/local/go/bin
```

``` bash
$ cd $HOME
$ mkdir -p go/src/github.com/RikiyaFujii
$ cd go/src/github.com/RikiyaFujii
$ git clone https://github.com/RikiyaFujii/supporterz.git
```


## Dockerのインストール
docker for mac のインストール (https://docs.docker.com/docker-for-mac/install/)

## Docker
```
supporterz/
          ├ docker/
          |       └ mysql/
          |              ├ initdb.d/
          |              |         ├ schema.sql  // テーブル定義
          |              |         └ testdata.sql // 初期データのinsert
          |              └ Dockerfile
          └ docker-compose.yml
```

### 実行方法
``` bash
$ cd $HOME/go/src/github.com/RikiyaFujii/supporterz
$ docker-compose up -d
```

### mysqlへの接続
``` bash
# パスワードは"password"
$ mysql --host 127.0.0.1 --port 3306 -u user -p
```

## Golangの実行
``` bash
$ pwd
  $HOME/go/src/github.com/RikiyaFujii
$ cd supporterz/src
$ go run server.go
```
