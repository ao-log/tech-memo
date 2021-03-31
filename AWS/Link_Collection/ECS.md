

## 記事

[[アップデート] 実行中のコンテナに乗り込んでコマンドを実行できる「ECS Exec」が公開されました](https://dev.classmethod.jp/articles/ecs-exec/)

使用するには以下の設定が必要。

* タスクロールに SSM 関連の権限を追加
* ECSサービスで「enableExecuteCommand」の設定を有効にする

以下のコマンドで接続。

```
$ aws ecs execute-command \
    --cluster クラスター名 \
    --task XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX \
    --container nginx \
    --interactive \
    --command "コマンド"
```

