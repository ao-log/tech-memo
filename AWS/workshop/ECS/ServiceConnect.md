
[ECS Service Connect ワークショップ](https://catalog.workshops.aws/ecs-service-connect/ja-JP)


## [シナリオ 1 : Service Connect を設定した ECS アプリケーションをデプロイする](https://catalog.workshops.aws/ecs-service-connect/ja-JP/scenario-1)

* CreateCluster で `[serviceConnectDefaults](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_CreateCluster.html#ECS-CreateCluster-request-serviceConnectDefaults)` を設定。これがデフォルトの名前空間になる。Cloud Map にも namespace ができる
* ECS サービスでは nmaespace を指定しつつ、clientAliases.dnsName でそのサービスに接続する際の FQDN を指定する
* ECS タスクは envoy コンテナがサイドカーとして稼働した状態になっている
* CloudWatch Metrics では ClusterName、DiscoveryName、ServiceName、TargetDiscoveryName のディメンションがある。RequestCount, NewConnectionCount などのメトリクスがある


## [シナリオ 2 : Service Connect を使用して、クラスターをまたいだサービス間通信を実現する](https://catalog.workshops.aws/ecs-service-connect/ja-JP/scenario-2)

* 2 つの ECS クラスターがある構成
* 2 つの ECS クラスターで同じ名前空間を設定する。これがポイント


## [シナリオ 3 : Service Connect の耐障害性を活用する](https://catalog.workshops.aws/ecs-service-connect/ja-JP/scenario-3)

* 検証用のアプリでは /api/responsecode/503 にアクセスすると 503 が返却されるようになっている
* Service Connect の envoy 側でリトライを行い正常なタスクの結果を返却するため、クライアントに対しては 200 しか返却されない


## [シナリオ 4 : 既存の ECS サービスを Service Connect を設定したサービスに移行する](https://catalog.workshops.aws/ecs-service-connect/ja-JP/scenario-4)

* Internal ロードバランサを使用する構成から Service Connect に移行する
* ECS サービスの設定で ServiceConnectConfiguration を設定し更新すればよい。clientAliases.dnsName を Internal ロードバランサと同じ内容にする




