

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


[詳細解説「AWS Cloud Map」とは #reinvent](https://dev.classmethod.jp/articles/cloud-map-perfect/)

* AWS Cloud Mapへの問い合わせ方式
  * Cloud Mapへの問い合わせ方式は、2種類。
    * DNSクエリ
    * APIコール
      * AWSのARN（Amazon Resource Name）などでアクセスするリソースに対して付与
* 設定が必要なリソース
  * 名前空間
  * サービス
  * サービスインスタンス
* 名前空間
  * 次のいずれかを選択。
    * API呼び出し
    * API呼び出しとVPCのDNSクエリ: VPC の指定も必要
    * API呼び出しと公開DNSクエリ：インターネットからの名前解決が可能。Route 53 のヘルスチェックが利用可能。
* サービス
  * サブドメインのようなイメージ
  * ルーティングポリシーを設定
  * ヘルスチェック方法を設定
* サービスインスタンス
  * 3 つ指定方法がある
    * IP アドレス
    * CNAME
    * リソースを特定するための情報(ARN など)



[CloudWatch Logs Insights でコンテナ単位のCPU・メモリ使用量などを確認する](https://dev.classmethod.jp/articles/ways-to-check-fargate-cpu-usage/)

* Container Insight の [View performance logs] から Logs Insights に遷移する。既にクエリが入力済みの状態になっている。



[Container Insights でコンテナ単位のCPU・メモリ使用率を表示させる](https://dev.classmethod.jp/articles/how-to-check-container-cpu-usage-by-container-insights/)

* コンテナ単位で表示させるにはタスク定義でコンテナの CPU、メモリを設定する必要がある。




