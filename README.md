# サポーターズ勉強会 サンプルコード

## 必要な環境
- Golang
- Docker
- Swagger


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

### Dockerの実行
``` bash
$ cd $HOME/go/src/github.com/RikiyaFujii/supporterz
$ docker-compose up -d
```

### mysqlへの接続
``` bash
# パスワードは"password"

# mysql環境がない人
docker exec -it supporterz_mysql_1 bash

# mysql環境がある人
$ mysql --host 127.0.0.1 --port 3306 -u user -p
```

## Golangの実行
``` bash
$ pwd
  $HOME/go/src/github.com/RikiyaFujii
$ cd supporterz/src
$ go run server.go
```

## Swaggerのインストール
$ brew install swagger-codegen

## SwaggerUIの起動
``` bash
$ pwd
  $HOME/go/src/githug.com/RikiyaFujii
$ cd swagger
$ swagger-codegen generate -i api_document.yaml -l nodejs-server -o doc/
$ cd doc
$ npm install
$ npm start
```
