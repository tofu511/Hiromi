# Hiromi
Go言語で実装されたシンプルなHTTPサーバーです。Go言語なので名前は`Hiromi`です。  
言語学習のために作成しました。

## 仕様
- `localhost:5163`でHTTPリクエストを受け取り、HTTPレスポンスを返す
- 対応するHTTPメソッドは`GET`のみ
- リクエストはブロックしない（マルチスレッド）
- Keep-Aliveしない
- HTTP Cacheしない
- `Accept-Language`が`ja`の場合、ステータスコード`200 OK`が`240 Exotic Japan!`になる

## 起動方法
```sh
$ go build server.go
$ ./server
```

or 

```sh
$ go run server.go
```

## Special Thanks
[simple-http-server](https://github.com/todokr/simple-http-server)