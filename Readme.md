# BackMe Server

おうちハッカソン

## Doc

[Swagger](https://app.swaggerhub.com/apis/nakao107107/ca-home-hackathon/1.0.0#/)  
[Client Repository](https://github.com/CA21engineer/Back-me-client)

## Setup

環境変数をコピーし、環境変数の内容を編集します。
使用するIAMは、S3にPUTする権限が必要です。  
詳細 → [#15](https://github.com/CA21engineer/Back-me-server/pull/15)

```
cp .env.example .env
```

コンテナを起動する。

```
# Start
docker-compose up

# ホットリロードが有効のため、コードに変更を加えると自動でビルドが走り、変更が反映されます。  
# http://localhost:8084/

# Stop
docker-compose down
```

## Deploy

Pull Request を Merge するなど、 `master` および `develop` ブランチに変更が発生すると、GitHub Actions を通じて Google App Engine にデプロイされます。

## Other

- On Production の Base URL は `https://ca-back-me-api.appspot.com/` です。
- 使用するIDEによっては、依存パッケージを解決するために `git clone` 後にプロジェクトのディレクトリ名を `ca-zoooom` に変更するか、IDE側の設定が別途必要になる可能性があります。（GitHubのリポジトリ名を変更したことに由来する）
- ゆるいクリーンアーキテクチャを採用しています。

## Libraries

- gorp (for connecting to MariaDB)
