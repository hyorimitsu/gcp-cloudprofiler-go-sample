# GAE で動かす場合

## GCPプロジェクトを作成
https://console.cloud.google.com/

## gcloud設定
`$ gcloud init`

## GAE へデブロイ
`& gcloud app deploy src/app.yaml`

# ローカル で動かす場合

## GCPプロジェクトを作成
https://console.cloud.google.com/

## GAE の設定
- Profiler 用のサービスアカウントを作成＆キーを生成
- docker-composeの環境設定を設定
　`GOOGLE_APPLICATION_CREDENTIALS=Docker内でのファイルパス`

## サービス起動
`./local.sh start`