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

## API仕様書
http://localhost:8000/swagger/index.html

## ローカルでのAPI確認方法
```
$ cd go-gacha-api/   # ディレクトリへ移動
$ docker-compose up -d # 移動したディレクトリでコンテナの起動
$ cd tools/frontend
$ npm run dev

http://localhost:8080　にアクセス
```
