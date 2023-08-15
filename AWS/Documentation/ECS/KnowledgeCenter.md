
# ナレッジセンター

[情報センター](https://repost.aws/ja/knowledge-center)


## アーキテクチャ

[Amazon ECS でホストされているサービスのブルー/グリーンデプロイはどのように実行できますか?](https://repost.aws/ja/knowledge-center/codedeploy-ecs-blue-green-deployment)

* CodeDeploy による Blue/Green デプロイ構成の作り方が書かれている。ELB は ALB


## キャパシティプランニング

[Amazon ECS の CPU 割り当てについて知っておくべきことは何ですか?](https://repost.aws/ja/knowledge-center/ecs-cpu-allocation)

* タスクレベルの CPU
  * Limit として機能する
  * EC2 では省略可能
  * Fargate では指定必須
* コンテナレベルの CPU
  * CpuShares にマッピングされる。よって 1 vCPU のインスタンス上では 512 の 1 タスクのみの場合はフルに CPU を使用できる。2 タスクの場合は、共にフルに使用している場合は 512 が Limit となる。ただし、余裕がある場合は 512 を超えて使用できる
  * Fargate では省略可能。合計値がタスクレベルの指定量を超えてはならない


## コンテナインスタンス

[Amazon ECS 最適化 AMI で ECS インスタンスを起動するにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/launch-ecs-optimized-ami)

* 最新の AMI ID は SSM パラメータから取得可能


[Amazon ECS でカスタム AMI を作成して使用する方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-create-custom-amis)

* AMI のレシピは公開されている https://github.com/aws/amazon-ecs-ami
* EC2 Image Builder でメンテナンスするのも手
* Docker, ECS Init のインストールが必要。ECS Init で ECS Agent がインストールされる
* 固有のデータを削除する
```
$ sudo rm -rf /var/log/ecs/*
$ sudo rm /var/lib/ecs/data/agent.db
```


[Amazon ECS でコンテナインスタンスタイプを変更するにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-change-container-instance-type)

* CloudFormation を使用している場合はインスタンスタイプのパラメータを変更しスタック更新する。ASG の台数を 2 倍に設定したあと、元のインスタンスをドレイニングする。その後 ASG の台数を元に戻す(ASG のデフォルトの終了ポリシーであることが前提)


[Amazon ECS または Amazon EC2 インスタンスがクラスターに参加できないのはなぜですか?](https://repost.aws/ja/knowledge-center/ecs-instance-unable-join-cluster)

* さまざまな原因が考えられる
  * ECS エンドポイントがインスタンスのドメインネームシステム (DNS) ホスト名にパブリックアクセスできない。
  * 公開サブネットの設定が正しくない
  * プライベートサブネットの設定が正しくない
  * VPC エンドポイントが正しく設定されていない
  * セキュリティグループがネットワークトラフィックを許可していない
  * EC2 インスタンスに必要な AWS Identity and Access Management (IAM) アクセス許可がない。または、ecs:RegisterContainerInstance 呼び出しが拒否されている。
  * ECS コンテナのインスタンスユーザーデータが正しく設定されていない
  * ECS エージェントが停止しているか、インスタンスで実行されていない
  * Auto Scaling グループの起動設定が正しくない (インスタンスが Auto Scaling グループの一部である場合)。
インスタンスに使用する Amazon マシンイメージ (AMI) が前提条件を満たしていない
* トラブルシューティング
  * AWSSupport-TroubleshootECSContainerInstance ランブックを使用することでトラブルシューティングの手順と推奨事項が記載される
  * ECS Agent のサービス稼働状況を確認
  * インスタンス内の構成を確認(カーネルバージョン 3.10 以上、Docker 1.9.0 以上)
  * ログファイルを確認


[コンテナが Amazon ECS の Amazon EC2 インスタンスメタデータにアクセスできないようにするにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-container-ec2-metadata)

* ホストネットワーキングモード:
  * 防ぐことができない
* awsvpc ネットワーキングモード:
  * `/etc/ecs/ecs.config ` に `ECS_AWSVPC_BLOCK_IMDS=true` を設定
* ブリッジネットワーキングモード
  * iptables により DROP する
  ```
  yum install iptables-services -y

  cat <<EOF > /etc/sysconfig/iptables 
  *filter
  :DOCKER-USER - [0:0]
  -A DOCKER-USER -d 169.254.169.254/32 -j DROP
  COMMIT
  EOF

  systemctl enable iptables && systemctl start iptables
  ```


[コンテナインスタンスが DRAINING に設定されているときに、Amazon ECS タスクが停止するまでに時間がかかる場合のトラブルシューティング方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-tasks-stop-delayed-draining)

* `minimumHealthyPercent` と `maximumPercent` を見直す
* 登録解除の遅延値が適切に設定されていることを確認する
* `ECS_CONTAINER_STOP_TIMEOUT` が適切に設定されていることを確認する


[終了した Amazon ECS コンテナインスタンスがまだそのクラスターに登録されているのはなぜですか?](https://repost.aws/ja/knowledge-center/deregister-ecs-instance)

* 通常はコンテナエージェントが自動的に登録解除する
* コンテナエージェントが切断されている場合にインスタンスが終了すると残ったままになるので、手動で登録解除する


## コンテナエージェント

[Amazon ECS で Docker のコンテナとイメージのクリーンアップを自動化するにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/container-image-cleanup-ecs)

* コンテナエージェントのパラメータで設定可能
  * ECS_DISABLE_IMAGE_CLEANUP: デフォルトでは false。true にすることでクリーンアップが無効化される
  * ECS_ENGINE_TASK_CLEANUP_WAIT_DURATION: コンテナのクリーンアップはオフにできない。間隔はこのパラメータで調整する


[Amazon Linux 2 の Docker と Amazon ECS コンテナエージェントに HTTP プロキシをセットアップする方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-http-proxy-docker-linux2)

* Docker の場合:
  * `/etc/systemd/system/docker.service.d/http-proxy.conf` にて設定。設定状況は docker info にて確認可能
  * `NO_PROXY` にて 169.254.169.254、169.254.170.2 の設定が必要
* コンテナエージェント側の設定
  * `/etc/ecs/ecs.config` にて設定
  * `NO_PROXY` にて 169.254.169.254、169.254.170.2、/var/run/docker.sock の設定が必要
* ecs-init 用の設定
  * `/etc/systemd/system/ecs.service.d` にて設定
  * `NO_PROXY` にて 169.254.169.254、169.254.170.2、/var/run/docker.sock の設定が必要


[Amazon ECS コンテナエージェントのプライベートリポジトリの認証情報を更新する方法を教えてください。](https://repost.aws/ja/knowledge-center/update-credentials-ecs)

* Secrets Manager を使用するのがベストプラクティス
* [タスクのプライベートレジストリの認証](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/private-auth.html) を参照のこと
* タスク定義のコンテナ定義にて `repositoryCredentials` を設定


[切断された Amazon ECS エージェントをトラブルシューティングするにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-agent-disconnected-linux2-ami)

* 正常なオペレーションの一環として切断と再接続を 1 時間に数回繰り返す場合がある
* 長期間続く場合は次の原因が考えられる
  * ネットワークの疎通性の問題
  * コンテナエージェントの IAM 権限不足
  * ホスト側の問題
  * Docker デーモンの問題
* トラブルシューティング
  * ECS Agent の稼働状況を確認
  ```
  $ sudo systemctl status ecs            
  $ sudo docker ps -f name=ecs-agent
  ```
  * Docker デーモンの稼働状況を確認
  ```
  $ sudo systemctl status docker
  ```
  * コンテナエージェント、ecs-init、Docker、cloud-init のログを確認
  * インスタンスロールの権限を確認
  * メモリなどのリソースが十分にあることを確認
  * ecs.config のクラスター名設定が正しいことを確認
  * エンドポイントとの疎通性を確認


[Amazon Linux 1 AMI の Amazon ECS コンテナインスタンスが切断されるのはなぜですか?](https://repost.aws/ja/knowledge-center/ecs-agent-disconnected)


## キャパシティープロバイダー

[Amazon ECS で Fargate Spot キャパシティープロバイダーを使用するにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-fargate-spot-capacity-provider)

* 既存クラスターに対しては `PutClusterCapacityProviders` で関連づけることができる
* `PutClusterCapacityProviders` コールから省略された既存のキャパシティープロバイダーは、クラスターとの関連付けが解除される
* Fargate Spot
  * SIGTERM シグナルを受信して、割り込みを処理できるように実装する
  * StopTimeout を 120 秒に設定するのがベストプラクティス
* キャパシティーを確保できない場合は `SERVICE_TASK_PLACEMENT_FAILURE` イベントがトリガーされる。確保できると `SERVICE_STEADY_STATE` イベントがトリガーされる
* FARGATE_SPOT キャパシティーがない場合、FARGATE にフェイルバックしない


[Amazon ECS の「キャパシティープロバイダーのマネージド終了保護設定が無効です」というエラーを解決する方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-termination-protection-error)

* キャパシティープロバイダーのマネージド終了保護を有効にする場合、Auto Scaling グループでスケールイン保護を有効にしておく必要がある


[Amazon ECS でキャパシティープロバイダーを削除する際、 DELETE_FAILED エラーを解決するにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-capacity-provider-error)

* キャパシティープロバイダーが使用中だと削除できない


## ECS サービス

[Fargate で Amazon ECS サービスの Auto Scaling を設定するにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-fargate-service-auto-scaling)

* Application Auto Scailng により対応可能
  * スケーラブルターゲット: 対象の ECS サービス、最小、最大タスク数を設定
  * スケーリングポリシー: ターゲット追跡ポリシーもしくはステップスケーリングポリシーを設定


[Amazon ECS でサービスのオートスケーリングの問題をトラブルシューティングするにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-service-auto-scaling-issues)

* ターゲット追跡ポリシー作成時に自動生成された CloudWatch アラームは編集、削除しないこと


[十分なディスク容量がない Amazon Linux 1 AMI を使用する Amazon ECS コンテナインスタンスにタスクを配置しないようにするにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-instance-task-disk-space)

* 残容量を定期的に監視し、属性値を PUT するスクリプトを稼働させる
```
aws ecs put-attributes \
  --cluster "$clusterName" \
  --attributes name="SpaceLeft",value="$SpaceLeftValue",targetType="container-instance",targetId="$instanceArn" \
  --region "$region"
```
* タスク配置制約を使用する
```
"placementConstraints": [
    {
        "expression": "attribute:SpaceLeft >= 0.1",
        "type": "memberOf"
    }
]
```


[AWS::ECS::Service リソースを UPDATE_IN_PROGRESS または UPDATE_ROLLBACK_IN_PROGRESS ステータスから解放する方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-service-stuck-update-status)

* タスクの起動に失敗するような場合にこの状態に陥る


[Amazon ECS サービスを削除または終了するにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-delete-terminate)

* 普通にコンソールなどから削除すれば OK


## タスク定義

[AWS Fargate の Amazon ECS コンテナのディスク容量を増やすにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/fargate-increase-disk-space)

* Fargate PV 1.4.0 ではデフォルト 20 GB のタスクストレージが割り当てられる
* 200 GB まで増やすことが可能。タスク定義にて ephemeralStorage を指定
* もしくは EFS を使用する


[AWS Fargate 上の Amazon ECS コンテナのディスク容量を増やす方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-fargate-increase-disk-space)

* ほぼ上記ナレッジと同じ内容


[Amazon ECS タスクに環境変数を渡す際の問題をトラブルシューティングするにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-task-environment-variables)

* Secrets Manager:
  * [Amazon ECS の AWS Secrets Manager シークレットに関連する問題のトラブルシューティングを行うにはどうすればよいですか？](https://repost.aws/ja/knowledge-center/ecs-secrets-manager-issues) を参照
  * タスク実行ロールの権限を確認
  * シークレットが存在すること、ARN 指定時の指定内容が正しいことを確認。 `myappsecret-xxxxxx` のようにランダムな文字が追加されている点に注意
  * Secrets Manager のエンドポイントとの疎通性を確認
* S3 上の環境変数ファイル:
  * タスク実行ロールの権限を確認
  * S3 のエンドポイントとの疎通性を確認
  * VPC にて DNS が有効化されていることを確認
* 既存タスクには変更が反映されない。反映させたい場合はタスクを置き換えること


## ネットワーキング

[Fargate の Amazon ECS タスクからデータベースに接続するにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-fargate-task-database-connection)

* このナレッジのサンプルでは `pymysql.connect()` により DB に接続している


[Application Load Balancer を作成し、Amazon ECS タスクを自動的に登録する方法にはどうすればよいですか?](https://repost.aws/ja/knowledge-center/create-alb-auto-register)

* ECS サービスの作成時に ELB と関連づける
* タスク定義にて `ContainerPort` の設定が必要


[CloudFormation で Amazon ECS サービス検出を使用するにはどうすればよいですか。](https://repost.aws/ja/knowledge-center/cloudformation-ecs-service-discovery)

* `AWS::ServiceDiscovery::PrivateDnsNamespace` で名前空間作成
* `AWS::ServiceDiscovery::Service` でサービス作成
* `AWS::ECS::Service` の `ServiceRegistries` にて `RegistryArn` を指定
* SRV レコードを確認する際は `dig srv awsExampleService.awsExampleNamespace. +short`


[AWS Cloud Map を使用して、ECS サービスのクロスアカウントサービス検出を設定する方法を教えてください。](https://repost.aws/ja/knowledge-center/fargate-service-discovery-cloud-map)

* クエリ元は別の AWS アカウント上にあるものとする
* ホストゾーンの所有アカウントから VPC 関連づけの承認リクエストを送信 `$ aws route53 create-vpc-association-authorization --hosted-zone-id <example-HoztedZoneId>  --vpc VPCRegion=<example_VPC_region>,VPCId=<example-source-vpc>`
* クエリ元の AWS アカウントでホストゾーンに VPC を関連づける `$ aws route53 associate-vpc-with-hosted-zone --hosted-zone-id <example-HoztedZoneId> --vpc VPCRegion=<example_VPC_region>,VPCId=<example-source-vpc>`
* ネットワークの疎通性がないので VPC ピアリングなどで疎通性を確保する


[Amazon EC2 の起動タイプの Amazon ECS タスクと Amazon RDS データベースとの間の接続に関する問題のトラブルシューティング方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-task-connect-rds-database)

* telnet により DB ポートへの接続が可能かを切り分け
* セキュリティグループを確認。bridge, host の場合はコンテナインスタンス側のセキュリティグループを確認
* ネットワーク ACL を確認
* パスワード類の環境変数を確認


## ヘルスチェック失敗のトラブルシューティング

[Amazon EC2 起動タイプを使用して Amazon ECS タスクを実行して Amazon ECS の Application Load Balancer のヘルスチェックに合格させるにはどうすればいいですか?](https://repost.aws/ja/knowledge-center/troubleshoot-unhealthy-checks-ecs)

* ロードバランサーと Amazon ECS タスク間の接続
  * セキュリティグループ、ネットワーク ACL 設定
  * ELB と ECS サービスが同じ AZ であること
* ターゲットグループのヘルスチェック
  * ポート番号、URL パス、タイムアウトの設定
* ECS コンテナ内のアプリケーションのステータスと設定
  * CPU 使用率が落ち着いていること(高いとアプリケーションがタイムアウト時間内に応答できないこともある)
  * ヘルスチェックの猶予期間を適切に設定
  * コンテナログを確認
  * 正しい HTTP ステータスコードを返却していることを確認
* コンテナインスタンスのステータス
  * コンテナインスタンス側で StatusCheckFailed メトリクスに問題がないことを確認

HTTP ステータスコードは以下のようなコマンドで確認可能
```
curl -I http://${IPADDR}:${PORT}/${PATH}
```


[Fargate での Amazon ECS タスクの Application Load Balancer のヘルスチェックエラーをトラブルシューティングする方法を教えてください。](https://repost.aws/ja/knowledge-center/fargate-alb-health-checks)

* 平均応答時間のチェック
```
time curl -Iv http://<example-task-pvt-ip>:<example-port>/<example_healthcheck_path>
```
* healthCheckGracePeriodSeconds を十分な長さに設定
* アクセスログにてヘルスチェックのアクセスを確認。CloudWatch Logs Insights にて以下のようなクエリで確認可能。
```
fields @timestamp, @message
  | sort @timestamp desc
  | filter @message like /ELB-HealthChecker/
```
* タスク内でポートの LISTEN 状況を確認
```
netstat -tulpn | grep LISTEN
```


[Fargate で Amazon ECS タスクを実行しているときの Network Load Balancer ヘルスチェックの失敗をトラブルシューティングするにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/fargate-nlb-health-checks)

* セキュリティグループ、ネットワーク ACL を確認
* TCP ヘルスチェックに応答できるかの確認
```
$ nc -z -v -w10 example-task-private-ip example-port
```
* healthCheckGracePeriodSeconds を十分な長さに設定
* アプリケーションログの確認


[Amazon ECS タスクのコンテナヘルスチェックの失敗をトラブルシューティングするにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-task-container-health-check-failures)

* ローカル環境でヘルスチェックに成功することを確認する
  * Dockerfile の HEALTHCHECK 設定でテスト可能
  * ※ ECS は Dockerfile 内の HEALTHCHECK 設定はモニタリングしない
* タスク定義中のコンテナヘルスチェックのコマンド指定が正しいことを確認する
* コンテナ起動に時間がかかる場合の対応
  * タスク定義の `startPeriod` にて十分な時間を指定
* コンテナログを確認


[Fargate での Amazon ECS タスクのヘルスチェックの失敗をトラブルシューティングするにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-fargate-health-check-failures)

* ネットワークの疎通性確認
* healthCheckGracePeriodSeconds を十分な長さに設定
* CPU, Memory 使用量を確認
* アプリケーションログを確認
* ヘルスチェックパスの確認
* 504 エラーのトラブルシューティング
  * 10 秒以内に接続を確立できない
* EC2 起動タイプだと正常に起動できるか


[ELB に登録されていて、正常に機能している Amazon ECS タスクが異常とマークされて置き換えられるのはなぜですか?](https://repost.aws/ja/knowledge-center/elb-ecs-tasks-improperly-replaced)

* ヘルスチェックの猶予時間


[Fargate での Amazon ECS タスクのロードバランサーのエラーのトラブルシューティング方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-fargate-load-balancer-errors)

* ecsServiceRole ロールなど


## タスク起動失敗のトラブルシューティング

[Amazon ECS タスクが保留状態のままになっているのはなぜですか?](https://repost.aws/ja/knowledge-center/ecs-tasks-stuck-pending-state)

* 以下の状況において保留状態が持続する場合がある
  * Docker デーモンが応答しない
    * CPU、メモリ使用率、ディスク IO を確認する
  * Docker イメージが大きい
    * イメージキャッシュに関するパラメータを調整する
  * Amazon ECS コンテナエージェントが、タスクの起動中に Amazon ECS サービスとの接続を失った
    * メタデータアクセスできるか確認する
    * コンテナエージェントのログを確認する
  * Amazon ECS コンテナエージェントで、既存のタスクを停止するのに長い時間がかかる
    * stopTimeout が長くなっていないかを確認する
  * Amazon Virtual Private Cloud (Amazon VPC) ルーティングが正しく設定されていない
    * 各エンドポイントへの疎通性があることを確認する
  * 必須のコンテナが、必須ではない異常なコンテナに依存している
    * 依存先のコンテナが正常に起動しているかを確認する


[Fargate の Amazon ECS タスクが [Pending] (保留中) 状態のまま停止している場合のトラブルシューティング方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-fargate-tasks-pending-state)

* ほぼ上記ナレッジと同じ内容


[Amazon ECS クラスターのタスクが開始されないのはなぜですか?](https://repost.aws/ja/knowledge-center/ecs-run-task-cluster)

* タスク停止理由を確認する
* タスク配置制約を確認する


[Amazon ECS のサービスで「the closest matching container-instance container-instance-id encountered error 'AGENT'」というエラーを解決する方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-container-instance-agent-error)

* コンテナエージェントのログを確認する
* コンテナエージェントを再起動する


[Amazon ECS で「[AWS service] was unable to place a task because no container instance met all of its requirements」(要件をすべて満たすコンテナインスタンスがないため、[AWS のサービス] はタスクを配置できませんでした) というエラーを解決するにはどうすればよいですか。](https://repost.aws/ja/knowledge-center/ecs-container-instance-requirement-error)

* 以下のような場合に発生する
  * クラスターにコンテナインスタンスがない
  * タスクに必要なポートがすでに使用されている
  * 十分な CPU, Memory, ENI がない
  * コンテナインスタンスに必要な属性がない


[Amazon ECS で「the closest matching container-instance container-instance-id has insufficient CPU units available」というエラーを解決する方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-container-instance-cpu-error)

* インスタンスタイプ、もしくはタスク定義の CPU 指定量を見直す
* コンテナインスタンスを追加する


[Fargate の Amazon ECS のネットワークインターフェイスプロビジョニングエラーを解決する方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-fargate-network-interface-errors)

* RunTask で起動している場合は StepFunction によりリトライを実装する


[Amazon ECS でスケジュールされたタスクに関連する問題をトラブルシューティングするにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-scheduled-task-issues)

* CloudWatch メトリクスで FailedInvocations が記録されていないかを確認する
* CloudTrail で RunTask を実行する
* タスクが成功している場合もあり、コンテナのログを確認する


[Amazon ECS リソースを起動したときに表示される「ロールを適用し、ロードバランサーに設定されたリスナーを検証できません」という AWS CloudFormation のエラーを解決する方法を教えてください。](https://repost.aws/ja/knowledge-center/assume-role-validate-listeners)

* サービスロールに ELB 関連のアクションの許可があるかを確認する
* DependsOn を指定し、IAM エンティティの作成後に ECS リソースを作成する


[Amazon ECS クラスターでタスクの起動に失敗する場合の「Image is exist」エラーを解決する方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-cluster-image-not-exist-error)

* イメージの指定が正しいか、リポジトリに格納されているかを確認する


[Amazon ECS for Fargate で「dockertimeouterror unable transition start timeout after wait 3m0s」というエラーを解決するにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-fargate-docker-timeout-error)

* 各エンドポイントへの疎通性があるかを確認する


[Amazon ECS の「ResourceInitializationError: failed to validate logger args (ResourceInitializationError: ロガー引数の検証に失敗しました)」エラーを解決するにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-resource-initialization-error)

* AWSSupport-TroubleshootECSTaskFailedToStart ランブックを使用して切り分ける
* CloudWatch Logs のロググループがあることを確認する。もしくはタスク定義で自動作成する設定にする


## タスク停止

**TODO**

[Amazon ECS タスクが停止するのはなぜですか?](https://repost.aws/ja/knowledge-center/ecs-task-stopped)

[Amazon ECS サービスで実行中のタスク数が変更されたのはなぜですか?](https://repost.aws/ja/knowledge-center/ecs-running-task-count-change)

[Amazon ECS における OutOfMemory エラーのトラブルシューティング方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-resolve-outofmemory-errors)


## ストレージ

[Auto Scaling グループを使用してクラスターを手動で起動した場合、Amazon ECS コンテナインスタンスで使用可能なディスク容量を増やすにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-container-storage-increase-autoscale)

[コンテナインスタンスをスタンドアロン Amazon EC2 インスタンスとして起動した場合、Amazon ECS コンテナインスタンスで使用可能なディスク容量を増やすにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-container-storage-increase-ec2)

[AWS マネジメントコンソールから ECS クラスターを起動した場合、Amazon ECS コンテナインスタンスで利用可能なディスクスペースを増やすにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-container-storage-increase-console)

[AWS Fargate 上の Amazon ECS コンテナのディスク容量を増やす方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-fargate-increase-disk-space)

[EC2 で実行されている ECS コンテナまたはタスクに EFS ファイルシステムをマウントする方法を教えてください。](https://repost.aws/ja/knowledge-center/efs-mount-on-ecs-container-or-task)

[Fargate で実行されている Amazon ECS コンテナまたはタスクに Amazon EFS ファイルシステムをマウントする方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-fargate-mount-efs-containers-tasks)

[AWS Fargate タスクで Amazon EFS ボリュームをマウントできないのはなぜですか?](https://repost.aws/ja/knowledge-center/fargate-unable-to-mount-efs)


## 認証、認可

[Fargate の Amazon ECS タスクから他の AWS サービスにアクセスする方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-fargate-access-aws-services)

* 例えば S3 にアクセスする際はタスクロールを設定
* タスクメタデータより認証情報を取得。`curl 169.254.170.2$AWS_CONTAINER_CREDENTIALS_RELATIVE_URI`
* 環境変数 AWS_CONTAINER_CREDENTIALS_RELATIVE_URI は PID 1 でのみ使用可能


[Amazon ECS で「Access Denied」(アクセス拒否) エラーを発生させないように IAM タスクロールを設定するにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-iam-task-roles-config-errors)

[Amazon ECS タスクの実行中の「ECS がロールを引き受けることができません」というエラーをトラブルシューティングするには、どうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-unable-to-assume-role)


## モニタリング

[ECS タスクとコンテナのデプロイをモニタリングするように CloudWatch Container Insights を設定するにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/cloudwatch-container-insights-ecs)

[Fargate で Amazon ECS タスクの高いメモリ使用率をモニタリングする方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-tasks-fargate-memory-utilization)

[Fargate での Amazon ECS タスクの高い CPU 使用率をトラブルシューティングするにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-fargate-high-cpu-utilization)


## ECR、コンテナイメージ

[Amazon ECS タスクに Amazon ECR イメージリポジトリからイメージを取得することを許可する方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-tasks-pull-images-ecr-repository)

* EC2 の場合: タスク実行ロールもしくはインスタンスロールにて指定。タスク実行ロールへの指定がベストプラクティス
* Fargate の場合: タスク実行ロールにて指定


[Amazon ECS EC2 起動タイプタスクの「CannotPullContainerError」エラーを解決するにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-pull-container-error)

原因としてありえるのは以下のもの。

* ネットワークの疎通性
  * インターネットへのアクセス経路
  * もしくは VPC エンドポイント
* IAM ロールに、イメージをプルまたはプッシュするための権限が許可されていない
* DockerHub のレート制限
* イメージ名、タグ名の指定ミス。もしくは存在していない


[Fargate での Amazon ECS タスクの「cannotpullcontainererror」エラーはどのように解決すればよいですか?](https://repost.aws/ja/knowledge-center/ecs-fargate-pull-container-error)

[Amazon ECR でエラー「CannotPullContainerError: API エラー」を解決する方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-pull-container-api-error-ecr)

* ECR エンドポイントへの疎通性がない
  * EC2 起動タイプでパブリックサブネットの場合、[パブリック IP の自動割り当て] が有効化されていること
* Amazon ECR リポジトリポリシーで制限されている
* イメージが見つからない
* S3 へのアクセスがエンドポイントポリシーによって拒否されている


[エラー "CannotPullContainerError を解決するには： Amazon ECS の「プルレート上限」に達しましたか?](https://repost.aws/ja/knowledge-center/ecs-pull-container-error-rate-limit)

[Amazon ECR から Docker イメージを取り出す際に、Amazon ECS の「error pulling image configuration: error parsing HTTP 403 response body」を解決する方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-ecr-docker-image-error)

[Amazon ECS で「unable to pull secrets or registry auth」(シークレットまたはレジストリ認証をプルできません) というエラーのトラブルシューティング方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-unable-to-pull-secrets)


## ログ関連

[AWS Fargate で Amazon ECS タスクのログドライバを設定するにはどうすればよいですか？](https://repost.aws/ja/knowledge-center/ecs-tasks-fargate-log-drivers)

[AWS Fargate 上の Amazon ECS の複数の送信先にコンテナログを送信するにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-container-log-destinations-fargate)

[Amazon ECS コンテナログが Amazon CloudWatch Logs に配信されないのはなぜですか?](https://repost.aws/ja/knowledge-center/ecs-container-logs-cloudwatch)

[欠落した Amazon ECS または Amazon EKS のコンテナログをトラブルシューティングするにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-eks-troubleshoot-container-logs)

一時的なログ送信失敗ではなく、全くログ送信できない状況のトラブルシューティング記事

* タスク定義の `logConfiguration` で正しく設定されていることを確認
  * コンテナ定義中で指定するので、対象コンテナに設定されているかを確認
* インスタンスロール、タスク実行ロールにて CloudWatch Logs の許可があることを確認

コンテナ側のトラブルシューティング

* STDOUT と STDERR にリンクされたアプリケーションログファイルでコンテナを構築、または /proc/1/fd/1 (stdout) と /proc/1/fd/2 (stderr) に直接ログを記録するようにアプリケーションを設定


## ECS Exec

[Fargate タスクで Amazon ECS Exec を実行したときに表示されるエラーをトラブルシューティングする方法を教えてください。](https://repost.aws/ja/knowledge-center/fargate-ecs-exec-errors)

* AWS CloudShell を使うのがベストプラクティス
* ExecuteCommand が true になっている必要がある
* タスクロール側に SSM に関するアクションの許可が必要
* 各エンドポイントとの疎通性が必要
* check-ecs-exec.sh による診断
* シェルの指定がイメージにあった内容であることを確認
* SSM Agent のログを確認


[Amazon ECS Exec がアクティブ化されている Fargate タスクの SSM エージェントログを取得する方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-exec-ssm-logs-fargate-tasks)

* EFS をマウントすることで実現する
* タスク定義のコンテナ定義におけるコンテナパスは `/var/log/amazon` とする。また起動する ECS タスクは 1 個にする
* EC2 に EFS をマウントしてログを確認する


## その他

[ECS タスクのタグ付けに関連する問題をトラブルシューティングするにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-troubleshoot-tagging-tasks)

* サービス定義パラメータの `PropagateTags` が適切に設定されているかを確認する
* 新しい ARN 形式を使用しているかを確認する
  * 旧形式: arn:aws:ecs:region:aws_account_id:service/service-name
  * 新形式: arn:aws:ecs:region:aws_account_id:service/cluster-name/service-name
* `ecs:TagResource` アクションが許可されていることを確認する
* [ECS マネージドタグ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-using-tags.html#managed-tags)
* Billing and Cost Management コンソールからコスト配分タグが有効化されているかを確認する


[Amazon ECS の API コールに関する一般的なエラーをトラブルシューティングするにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-api-common-errors)

* CloudTrail を確認する
* 各エラーへの対応
  * AccessDeniedException: ポリシーを見直す
  * ClientException: リソースが存在しない場合などに発生するので、コマンドラインの指定内容が正しいかなど確認
  * ClusterNotFoundException: クラスター名の指定ミスを疑う
  * InvalidParameterException: コマンドオプションの指定内容が正しいかなど確認
  * ServerException: 一時的なものである場合が多いが、続く場合は AWS サポートへ問い合わせる
  * ServiceNotActiveException: サービスがアクティブではない場合に発生
  * PlatformTaskDefinitionIncompatibilityException: プラットフォームバージョンでサポートされていない機能を使用している場合に発生
  * PlatformUnknownException: プラットフォームバージョンの指定ミス
  * ServiceNotFoundException: サービスが存在しない
  * UnsupportedFeatureException: 特定のリージョンで使用できない機能を使用しようとした時


[Amazon ECS のブルー/グリーンデプロイに関連する問題のトラブルシューティング方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-blue-green-deployment)

* CodeDeploy のサービスロールの権限に起因する場合は、権限を見直す。`AWSCodeDeployRoleForECS` ポリシーが用意されている
* ELB ヘルスチェックに失敗する場合は、コンテナイメージの作成内容や、タスク定義のポートマッピングがターゲットグループのポート番号とマッチしていることを確認
* 本番リスナー、テストリスナーが両方ともプライマリターゲットグループに設定されていることを確認
* コンテナインスタンスの容量が十分であることを確認










## 旧ナレッジセンター

[Amazon ECS のタスクにメモリを割り当てるにはどうすればよいですか ?](https://aws.amazon.com/jp/premiumsupport/knowledge-center/allocate-ecs-memory-tasks/)

タスク定義のパラメータで設定する。
* memoryReservation (ソフトリミット) 
* memory (ハードリミット)


[Amazon ECS の動的ポートマッピングのセットアップ方法を教えてください。](https://aws.amazon.com/jp/premiumsupport/knowledge-center/dynamic-port-mapping-ecs/)

タスク定義のホストポートを 0 に設定する。


[Amazon Linux で Docker と Amazon ECS コンテナエージェントに HTTP プロキシをセットアップするにはどうすればよいですか?](https://aws.amazon.com/jp/premiumsupport/knowledge-center/http-proxy-docker-ecs/)

```/etc/sysconfig/docker``` に環境変数を設定し、docker デーモンを再起動する。


[Amazon ECS の Application Load Balancer ヘルスチェックに合格するために Amazon EC2 インスタンスを取得するにはどうすればよいですか?](https://aws.amazon.com/jp/premiumsupport/knowledge-center/troubleshoot-unhealthy-checks-ecs/)

* タスク内のアプリケーションが正しいレスポンスコードを返していること
* セキュリティグループで遮断されていないこと


[Amazon ECS コンテナインスタンスから自動的にログを収集するにはどうすればよいですか。](https://aws.amazon.com/jp/premiumsupport/knowledge-center/debug-mode-ecs-agent-docker/)

サポート送付用のログコレクターの使用方法について。


[Amazon ECS に最適化された AMI を起動するにはどうすればよいですか?](https://aws.amazon.com/jp/premiumsupport/knowledge-center/launch-ecs-optimized-ami/)

* EC2 インスタンスの起動時に Marketplace から「ecs-optimized」で検索して AMI を探す。
* SSM のパラメータを使用する。

```
aws ssm get-parameters ¥
    --names /aws/service/ecs/optimized-ami/amazon-linux/recommended/image_id ¥
    --region リージョン ¥
    --query "Parameters[0].Value"
```


[Amazon ECS のタスクでコンテナが終了する問題をトラブルシューティングする方法を教えてください。](https://aws.amazon.com/jp/premiumsupport/knowledge-center/ecs-tasks-container-exit-issues/)

* サービスのイベントログを確認する。
* 停止したタスクで停止理由を確認する。
* ログを確認する。


[Amazon ECS タスクで秘密情報や機密情報をコンテナに安全に渡す方法を教えてください。](https://aws.amazon.com/jp/premiumsupport/knowledge-center/ecs-data-security-container-task/)

* Secrets Manager に保存する。
* ECS タスク実行ロールに Systems Manager の読み取り権限を付与する。
* タスク定義で secrets により設定する。


[Amazon EC2 の起動タイプの Amazon ECS タスクと Amazon RDS データベースとの間の接続に関する問題のトラブルシューティング方法を教えてください。](https://aws.amazon.com/jp/premiumsupport/knowledge-center/ecs-task-connect-rds-database/)

* ネットワークモードが host, bridge の場合は、ホストから RDS に接続できるかで切り分け可能。コンテナインスタンスのセキュリティグループが使用されるため。


