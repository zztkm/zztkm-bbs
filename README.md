# zztkm bbs

フロント部分はパクらせてもらう
https://github.com/Neos21/legacy-of-bbs/tree/master

バックグラウンドを syumai/workers で実装する


## Development

開発サーバーの起動

```bash
make dev
```

app/hello.go 内の Response 用の構造体定義を変更したときは、以下のコマンドを実行する

```bash
easyjson app/hello.go
```