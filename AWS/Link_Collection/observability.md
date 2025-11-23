
[ベストプラクティス: AWS によるオブザーバビリティの実装](https://aws.amazon.com/jp/blogs/news/best-practices-implementing-observability-with-aws/)

[オブザーバビリティとは](https://aws-observability.github.io/observability-best-practices/ja/)

* オブザーバビリティとは、観測対象のシステムからのシグナルに基づいて、継続的にアクション可能な洞察を生成および発見する機能
* [ダッシュボード、アプリケーションパフォーマンスモニタリング、コンテナなどのソリューションごとのベストプラクティス](https://aws-observability.github.io/observability-best-practices/ja/guides/)
  * ビジネス、プロジェクト、ユーザーにとって何が重要かを検討し KPI を定める
  * 時系列データとして採取できる必要がある
* [ログやトレースなど、異なるデータタイプの使用に関するベストプラクティス](https://aws-observability.github.io/observability-best-practices/ja/signals/logs/)
  * 構造化ログにしておくと、処理しやすい
* 特定のAWSツールのベストプラクティス(ただし、これらは他のベンダー製品にも大部分適用できます)
* AWSによるオブザーバビリティのためのキュレーションされたレシピ


[Amazon Managed Service for Prometheus の一般提供開始](https://aws.amazon.com/jp/blogs/news/amazon-managed-service-for-prometheus-is-now-generally-available/)

[Amazon Managed Service for Prometheus を使用して EC2 環境を監視する](https://aws.amazon.com/jp/blogs/news/using-amazon-managed-service-for-prometheus-to-monitor-ec2-environments/)

[【レポート】Open-source observability at AWS − 可観測性を支える OSS と AWS の『いま』を知る #AWS-35 #AWSSummit](https://dev.classmethod.jp/articles/awssummit2021-aws-35/)

[「CloudWatch MCP サーバーと CloudWatch Application Signals MCP サーバーを使ってみた」というタイトルで登壇しました](https://dev.classmethod.jp/articles/cloudwatch-and-application-signals-mcp-server/)


# builders.flash

[OpenTelemetry Collector の中身と種類を知ろう](https://aws.amazon.com/jp/builders-flash/202503/opentelemetry-collector/)

* 3 種のコンポーネント
  * レシーバー : テレメトリーの受信を担当
  * プロセッサー : テレメトリーの加工やバッファリングなどの処理を担当
  * エクスポーター : テレメトリーのバックエンドへの送信を担当
  * レシーバー、プロセッサー、エクスポーターの方向。プロセッサーは複数を連結可能
* 設定ファイル例
```yaml
receivers:
  otlp:

processors:
  batch:
  memory_limiter:
    limit_mib: 1536
    spike_limit_mib: 512
    check_interval: 5s

exporters:
  otlp:

service:
  pipelines:
    metrics:
      receivers: [otlp]
      processors: [memory_limiter, batch]
      exporters: [otlp]
```

