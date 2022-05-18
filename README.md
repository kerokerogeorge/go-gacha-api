# Gacha API

## Dockerのインストール
このアプリケーションではDockerを使用するので、事前にインストールしてください。
docker hub の公式ページでインストーラーがダウンロードできます。

## SETUP

```
$ cd go-gacha-api/   # ディレクトリへ移動
$ docker-compose up  # 移動したディレクトリでコンテナの起動
```
[注意]
初回起動時にdbコンテナよりサーバーコンテナが立ち上がってしまうので、エラーが起こる場合があります。立ち上げ直すと正常に起動します。


データベースのコンテナとAPI用のコンテナが起動したら準備完了です。

### ユーザーの作成
```
$ curl -X POST http://localhost:8000/user/create -H "Content-Type: application/json" -d '{"name": "任意の名前"}'
###-> {"token":"b6af5c29-ad7e-4230-b08d-010911a5f68f"}
```

### ユーザー情報の取得
ユーザー作成時に取得したTOKENをHEADERにセット
```
$ curl -X GET http://localhost:8000/user -H "x-token: b6af5c29-ad7e-4230-b08d-010911a5f68f"
###-> {"name":"shuji-curl"}
```

### ユーザー情報のアップデート
ユーザー作成時に取得したTOKENをHEADERにセットし、名前をBODYにいれてリクエストを送信
```
$ curl -X PATCH http://localhost:8000/user/update -H "x-token: 01dd4bd0-b72a-4885-94cb-a338aa6f7466" -d '{"name": "shuji-curl-edited"}'
###-> {"name":"shuji-curl-edited"}
```

### ガチャの実行
ユーザー作成時に取得したTOKENをHEADERにセット、ガチャを引きたい回数をBODYにセットし、リクエストを送信
```
$ curl -X POST http://localhost:8000/gacha/draw -H "x-token: 01dd4bd0-b72a-4885-94cb-a338aa6f7466" -d '{"times": 3 }'
###-> {"results":[{"characterId":"2","name":"kimori"},{"characterId":"4","name":"mizugoro"},{"characterId":"2","name":"kimori"}]}
```

### ユーザーが保有しているキャラクターの一覧取得
```
$ curl -X GET http://localhost:8000/character/list -H "x-token: 01dd4bd0-b72a-4885-94cb-a338aa6f7466"
###-> {"characters":[{"userCharacterId":"3","characterId":"2","name":"kimori"}]}
```


