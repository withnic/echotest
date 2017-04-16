# echoの評価用

# とりあえずのtwitter clone

* echo-session が動かなかったので、session周りは実装しない。
* というか、微妙なのでversionあがるまで特に echo使わない。
* beego か一度、素で書く。


# つかいかた

```
cd ./web/app/

goose up

cd ..

go run server.go
```

http://localhost:8080

# やってないこと

* sessionのミドルウェアが動かかなったのでセッション周り。

* formは探せば、csrfのライブラリあるだろうけど取り敢えずなし。

* test動かなくなった(echoのバージョンアップ)ので削除した。
