# Document

[Amazon Managed Service for Prometheus とは](https://docs.aws.amazon.com/ja_jp/prometheus/latest/userguide/what-is-Amazon-Managed-Service-Prometheus.html)


[AWS Distro for OpenTelemetry をコレクターとして使用する](https://docs.aws.amazon.com/ja_jp/prometheus/latest/userguide/AMP-ingest-with-adot.html)



# BlackBelt

[Amazon Managed Service for Prometheus(AMP)](https://pages.awscloud.com/rs/112-TZM-766/images/AWS-Black-Belt_2023_AmazonManagedServiceForPrometheus_0131_v1.pdf)

* Prometheus には Remote Write の機能で外部ストレージに書き出す機能がある。この外部ストレージ側がマネージドになっているサービス
* 機能
  * AMP ではリモートストレージのエンドポイントを作ることができる。エンドポイントへの送信には IAM の許可が必要
  * Grafana にてデータソースに AMP ワークスペースのクエリエンドポイントを指定することで連携できる
  * 記録ルール、アラートルールをサポート。Amazon SNS に通知可能
* ECS のメトリクスを ADOT Collector にて AMP に送信できる
  * 自前でタスク定義を書いてもよいが、コンソールから設定する場合は自動でインジェクト可能。Prometheus ライブラリ、OpenTelemetry SDK から選択
  * [ecs-metrics Receiver](https://aws-otel.github.io/docs/components/ecs-metrics-receiver) によりメトリクスを収集
  * 独自の設定を行いたい場合は SSM パラメータストアに設定ファイルの内容を格納



# 参考

* Document
  * [Amazon Managed Service for Prometheus](https://aws.amazon.com/jp/prometheus/)
  * [Amazon Managed Service for Prometheus とは](https://docs.aws.amazon.com/ja_jp/prometheus/latest/userguide/what-is-Amazon-Managed-Service-Prometheus.html)
* Black Belt
  * [Amazon Managed Service for Prometheus(AMP)](https://pages.awscloud.com/rs/112-TZM-766/images/AWS-Black-Belt_2023_AmazonManagedServiceForPrometheus_0131_v1.pdf)

