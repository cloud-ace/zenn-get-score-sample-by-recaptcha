# reCAPTCHA Enterprise のスコアベースのキーの保護を使ってみた

（cloud ace の zenn 投稿目的のリポジトリです。）

# 本リポジトリについて

reCAPTCHA Enterprise のスコアベースのキーの保護の検証を行うために実装したアプリケーションです。

# アプリケーション構成

- フロント
  - nginx
- バックエンド
  - golang

# 前提

Google Cloud プロジェクトに reCAPTCHA Enterprise のキーを作成していること

# アプリケーションの起動手順

1. 環境変数の設定
   下記の環境変数を動作環境に合わせて設定する。

- bakend/Dockerfile

  - PROJECT_ID: プロジェクト ID
  - SITE_KEY: reCAPTCHA Enterprise のキー ID

- frontend/init.html

  - 4,27 行目に reCAPTCHA Enterprise のキー ID を設定

- docker-copose.yml
  - volumes に[アプリケーションのデフォルト認証情報](https://cloud.google.com/docs/authentication/application-default-credentials?hl=ja)のパスを設定

2. ビルド

```bash
docker-copose build
```

3. 起動

```bash
docker-compose up -d
```
