
[New – Amazon ECS Service Connect Enabling Easy Communication Between Microservices](https://aws.amazon.com/jp/blogs/aws/new-amazon-ecs-service-connect-enabling-easy-communication-between-microservices/)

* 名前空間内のサービスの名前解決ができるようになる
* ロードバランサを用意することなく、トラフィックをロードバランシングできる
* ヘルスチェック、503 エラー時の自動リトライ、コネクションドレイニング
* コンソールから簡易に監視できるダッシュボードが提供される


[Amazon ECS Service Connect によるサービス間通信の管理](https://pages.awscloud.com/rs/112-TZM-766/images/20230126_26th_ISV_DiveDeepSeminar_ECS_Service_Connect.pdf)

* 従来からある選択肢
  * ELB
    * ELB が必要
    * コンポーネントが多い分、レイテンシが増える
  * サービスディスカバリ
    * トレメトリデータの収集やリトライまでは対応していない
  * App Mesh
    * 柔軟性に伴う複雑性
* いいところどり
  * ELB のようにテレメトリデータが取得できる
  * ECS service discovery のようにシンプル
  * App Mesh のように信頼性のある通信を提供


[ECSの新ネットワーク機能「Service Connect」がリリースされました！ #reinvent](https://dev.classmethod.jp/articles/ecs-service-connet/)

* 特徴
  * 任意の名前をサービスに付与して接続
  * エンドポイントのヘルスチェック対応
  * コンソールやCloudWatchによる豊富なメトリクスの提供
  * 自動接続ドレインのサポート
  * 503 エラーの自動試行


[aws/amazon-ecs-service-connect-agent](https://github.com/aws/amazon-ecs-service-connect-agent)

* Envoy プロキシをモニタリングし、管理インターフェースを提供
* Envoy の設定に関すると思われるコード
  * [agent_config.go](https://github.com/aws/amazon-ecs-service-connect-agent/blob/main/agent/config/agent_config.go)
* HTTP クライアント側に関すると思われるコード
  * [agent_http_client.go](https://github.com/aws/amazon-ecs-service-connect-agent/blob/main/agent/client/agent_http_client.go)

