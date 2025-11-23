
# AWS Blog

## Categories

[Containers](https://aws.amazon.com/jp/blogs/containers/)

[Category: Amazon Elastic Container Service](https://aws.amazon.com/jp/blogs/news/category/compute/amazon-elastic-container-service/)



## Blogs

#### 内部的な仕組みの詳細

[詳解: Amazon ECSのタスクネットワーク](https://aws.amazon.com/jp/blogs/news/under-the-hood-task-networking-for-amazon-ecs/)

* ネットワークスタックはネットワークの namespace を通じて設定される。
* ECS Agent から CNI プラグインを呼び出して操作。複数のプラグインを順に呼び出す。
* 最初に呼び出されるものは ecs-eni プラグイン。ENI を namespace にアタッチ。
* ecs-bridge と ecs-ipamプラグインにより認証情報エンドポイントに HTTP リクエストできるように設定。
* 処理の流れ
  * eth1 を作成。
  * eth1 をタスク用の namespace に移動。
* pause コンテナを作成。ネットワークスタックの設定とコンテナ内で実行されるコマンドのレースコンディションを防ぐため。
* pause コンテナの network namespace をセットアップ。
* タスク定義の各コンテナを開始。


[詳細: Fargate データプレーン](https://aws.amazon.com/jp/blogs/news/under-the-hood-fargate-data-plane/)

* Docker Engine を Containerd に置き換えた

他にも Fargate の実装にあたっては以下のカスタマイズをおこなっている。

* awsvpc のサポート。Container Networking Interface (CNI) プラグインを利用しネットワークをセットアップ。
* FireLens のサポート。
* ECS タスクのメタデータエンドポイント。統計情報などを採取できるようにローカルの HTTP エンドポイントを提供。

以下の特徴がある。

* pause コンテナは不要となった。Containerd は、コンテナの起動時に外部で作成されたネットワーク名前空間の使用をサポートしているため。
* さまざまな宛先へのログ送信のサポート。Containerd の Shim ログ記録プラグインを拡張している。amazon-ecs-shim-loggers-for-containerd リポジトリ。
* ランタイムは runC のほか FireCracker に切り替えられる構造。
* Fargate エージェント。

コンテナ起動時の流れがシーケンス図で示されている。

* network namespace の作成
* Customer ENI のプロビジョニング
* AWS Secrets Manager からシークレットの取得
* コンテナイメージのプル
* containerd へコンテナ開始の指示
* コンテナランタイムへコンテナ開始の指示。logging shim plugin の開始。

Fargate データプレーンは Fargate Agent, Containerd。こちらは Fargate VPC に繋がっている。
カスタマータスクゾーンにて container runtime shim を介して runc によりコンテナが起動。Customer ENI により Customer の VPC に繋がっている。


[ECS のアプリケーションを正常にシャットダウンする方法](https://aws.amazon.com/jp/blogs/news/graceful-shutdowns-with-ecs/)

* エントリプロセス
  * コンテナでは、Dockerfile の ENTRYPOINT および CMD ディレクティブで指定されたプロセス (エントリプロセスと呼ぶ) が、コンテナ内の他のすべてのプロセスの親になる。
  * ENTRYPOINT が bin/sh -c my-app の場合は、sh はコンテナが停止したときに SIGTERM シグナルを受信するものの my-app に渡されない。
  * shell を安全に使用するには２つの方法がある。
    * 1) シェルスクリプトを介して実行される実際のアプリケーションに exec というプレフィックスをつける。
    * 2) tini (Amazon ECS-Optimized AMI の Docker ランタイムに同梱されている) や dumb-init などの専用のプロセスマネージャーを使用する。
  * あるコマンドを exec をつけて実行すると、子プロセスを生成するのではなく現在実行中のプロセス (この場合は shell ) の内容を新しい実行可能プロセスに置き換える。
  * tini や dumb-init などのプロセスマネージャーを使用すると、これらのプログラムは SIGTERM を受け取ると、アプリケーションを含むすべての子プロセスグループに SIGTERM を送信する。以下のように指定すると良い
  ```
  ENTRYPOINT ["tini", "--", "/path/to/application"]
  ```
  * `InitProcessEnabled` を有効にした ECS タスクを実行すると、ECS はコンテナの init プロセスとして tini を自動的に実行する
* SIGTERM シグナルの処理
  * デフォルトの停止シグナルは SIGTERM だが、Dockerfile に STOPSIGNAL ディレクティブを追加することで上書きできる。
  * ALB から draining され unused となったあとに SIGTERM が実行される（単一の ALB に登録されている場合）。
  * EC2 インスタンスが drain 状態になっても RunTask によって実行されたタスクは draining されない。
* スポットインスタンスの中断
  * `ECS_ENABLE_SPOT_INSTANCE_DRAINING` を true に設定することでスポットインスタンスの中断通知を受信すると、ECS エージェントはインスタンスを DRAINING 状態にする
  * スポットインスタンスの中断通知を受信すると ALB から登録解除される。
  * 登録解除後に SIGTERM なので、登録解除の遅延は 120 秒未満にする必要がある。


#### リソース量の管理

[詳解: Amazon Elastic Container Service と AWS Fargate のタスク起動レートの向上](https://aws.amazon.com/jp/blogs/news/under-the-hood-amazon-elastic-container-service-and-aws-fargate-increase-task-launch-rates/)


[詳解: Amazon ECS による CPU とメモリのリソース管理](https://aws.amazon.com/jp/blogs/news/how-amazon-ecs-manages-cpu-and-memory-resources/)

二つのアプローチについて書かれている。
* 柔軟性を最大化する: CPU を指定しない。この場合は cpu.shares に 2 が設定される。メモリはソフトリミット以上使用できるため競合し OOM Killer が発生する可能性がある。スワップを設定することにより、OOM Killer の確率を下げられる。
* 制御を最大化する: 


[Fargate のサービスクォータが vCPU ベースに変更になります](https://aws.amazon.com/jp/blogs/news/migrating-fargate-service-quotas-to-vcpu-based-quotas/)


#### Security

[Building secure guardrails for Amazon ECS with AWS IAM and AWS CloudFormation Guard](https://aws.amazon.com/jp/blogs/containers/building-secure-guardrails-for-amazon-ecs-with-aws-iam-and-aws-cloudformation-guard/)


#### Networking

[AWS Cloud Map:アプリケーションのカスタムマップの簡単な作成と維持](https://aws.amazon.com/jp/blogs/news/aws-cloud-map-easily-create-and-maintain-custom-maps-of-your-applications/)


[ECS Anywhere のプライベート接続を確立する方法](https://aws.amazon.com/jp/blogs/news/how-to-establish-private-connectivity-with-ecs-anywhere/)

* ECS Anywhere で稼働させているホストとコントロールプレーンとの通信をプライベートにする方法についての記事
* SSM Agent, ECS Agent 用に各エンドポイントとの疎通性が必要
* AWS Direct Connect とサイト間 VPN の 2 つのオプションがある
* アーキテクチャ
  * オンプレミス
    * カスタマーゲートウェイが必要
    * 各 VPC エンドポイント(ecs-a, ecs-t, ecs, ssm, smmessages, ec2messages)への DNS クエリを Route 53 インバウンドリゾルバーエンドポイントに転送
  * AWS
    * vgw 経由で通信
    * ECS, SSM 用の VPC エンドポイント
      * com.amazonaws.<region>.ecs-agent
      * com.amazonaws.<region>.ecs-telemetry
      * com.amazonaws.<region>.ecs
      * com.amazonaws.<region>.ssm
      * com.amazonaws.<region>.ssmmessages
      * com.amazonaws.<region>.ec2messages
    * ECR を使用する場合は更に以下の VPC エンドポイント
      * com.amazonaws.<region>.ecr.dkr
      * com.amazonaws.<region>.ecr.api
      * com.amazonaws.<region>.s3
    * awslogs を使用する場合は更に以下の VPC エンドポイント
      * com.amazonaws.<region>.logs
    * VPC 内に Route 53 リゾルバーのインバウンドエンドポイント


[AWS App Mesh から Amazon ECS Service Connect への移行](https://aws.amazon.com/jp/blogs/news/migrating-from-aws-app-mesh-to-amazon-ecs-service-connect/)

* App Mesh は 2026 年 9 月 30 日にサポート終了
* ECS Service Connect では外れ値検出、リトライの機能がある。また、メトリクスも取得される
* App Mesh は次の仕組み
  * Envoy を使用
  * Virtual Service → Virtual Router → Virtual Node の経路
  * サービスディスカバリは Cloud Map
* 移行
  * ECS サービスの再作成が必要
  * ダウンタイムを回避するために Blue/Green 戦略を使用するのがよい
  * トラフィックの移行は Route 53 の加重ルーティング、CloudFront の継続的デプロイ、ALB の加重ターゲットグループなどがある


[Amazon ECS announces IPv6-only support](https://aws.amazon.com/jp/blogs/containers/amazon-ecs-announces-ipv6-only-support/)

* IPv6 only サブネットでは IPv6 アドレス、dualstack サブネットでは IPv6, IPv4 アドレスが割り当てられる
* IPv6 only サブネットの場合は ECR からのイメージプルは dualstack VPC Endpoint である必要がある
* Cloud Map は IPv6 をサポートしている
* RDS, EFS なども IPv6 をサポートしている
* Internet 向には [DNS64/NAT64](https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateway-nat64-dns64.html) が必要
* IPv4 で構成されたサービスから移行するには新しい IPv6 サービスを作成し、トラフィックを移行するのがおすすめ。


#### Build

[Building better container images](https://aws.amazon.com/jp/blogs/containers/building-better-container-images/)

* コンテナイメージは最適化され、セキュアで信頼できることが肝要
* プラクティス
  * 信頼できるソースからベースイメージを取得する
  * イメージは随時最新にする
  * latest ではなく、個別のタグをつける
  * ECR では暗号化のほか、ライフサイクル管理、レプリケーション、キャッシング、脆弱性スキャンなどの機能がある
  * Bottlerocket は最小限のソフトウェアしか含まれないため、よりセキュア
  * コンテナイメージの署名
  * レイヤー数を少なくすることで複雑性を少なくし、イメージサイズも肥大化しなくなる
  * マルチステージビルド
  * シークレット管理。Secrets Manager などを活用
  * scratch イメージの使用
  * 不要パッケージ、ファイルの削除
  * ディストロレスイメージの使用
  * コンテナ内プロセスを非 root ユーザで動作


[BuildKit クライアント用の Amazon ECR でのリモートキャッシュサポートの発表](https://aws.amazon.com/jp/blogs/news/announcing-remote-cache-support-in-amazon-ecr-for-buildkit-clients/)

* キャッシュは一時的なビルド環境では使用できないのが欠点
* ECR は OCI 準拠のレジストリだが、BuildKit のリモートキャッシュエクスポートは OCI 形式ではないためプッシュできなかった
* BuildKit 0.12 にてリモートビルドキャッシュを OCI 互換の方法で生成および保存できるソリューションを提供
* `docker build` 時に `--cache-to`, `--cache-from` オプションにより指定


#### Deploy

[Migrating from AWS CodeDeploy to Amazon ECS for blue/green deployments](https://aws.amazon.com/jp/blogs/containers/migrating-from-aws-codedeploy-to-amazon-ecs-for-blue-green-deployments/)

* CodeDeploy では未対応だった機能
  * Service Connect のサポート
  * ヘッドレスサービスのサポート。例えば、キュープロセッシングサービス
  * EBS の構成のサポート
  * 複数の ELB をアタッチした構成のサポート
  * 同一のリスナーポートで本番、テストトラフィックに対応できる
* 一方で ECS の blue/green で未対応のもの
  * all-at-once でのデプロイのみ
  * CodeDeploy ではトラフィックの Green 側への移行前の wait time を設定できたが、ECS blue/green ではライフサイクルフックでの対応が必要
* 移行方法
  * `UpdateService` にて、`deploymentStrategy` を `BLUE_GREEN` にする
  * 既存 ALB に新しいリスナーを作成し `BLUE_GREEN` の ECS サービスに転送する。既存リスナーのポート番号を変更した後、`BLUE_GREEN` のリスナーポート番号を本番用のものに変更する
  * ALB と ECS サービスを新規作成し、DNS の向き先を新規作成した ALB に変更する


[Extending deployment pipelines with Amazon ECS blue green deployments and lifecycle hooks](https://aws.amazon.com/jp/blogs/containers/extending-deployment-pipelines-with-amazon-ecs-blue-green-deployments-and-lifecycle-hooks/)


[AWS CodeDeploy による AWS Fargate と Amazon ECS でのBlue/Greenデプロイメントの実装](https://aws.amazon.com/jp/blogs/news/use-aws-codedeploy-to-implement-blue-green-deployments-for-aws-fargate-and-amazon-ecs/)


[Amazon ECR をソースとしてコンテナイメージの継続的デリバリパイプラインを構築する](https://aws.amazon.com/jp/blogs/news/build-a-continuous-delivery-pipeline-for-your-container-images-with-amazon-ecr-as-source/)


[Amazon ECS on AWS Fargate を利用したコンテナイメージのビルド](https://aws.amazon.com/jp/blogs/news/building-container-images-on-amazon-ecs-on-aws-fargate/)


[AWS App Mesh を使用した Amazon ECS でのカナリアデプロイパイプラインの作成](https://aws.amazon.com/jp/blogs/news/create-a-pipeline-with-canary-deployments-for-amazon-ecs-using-aws-app-mesh/)


[ECS Blueprints で Amazon ECS ベースのワークロードを加速しよう](https://aws.amazon.com/jp/blogs/news/accelerate-amazon-ecs-based-workloads-with-ecs-blueprints/)

* [ECS Blueprints](https://github.com/aws-ia/ecs-blueprints) for AWS Cloud Development Kit (AWS CDK)
* CDK で backend を構築する際の[コード](https://github.com/aws-ia/ecs-blueprints/tree/main/cdk/examples/backend_service)
  * container_image="public.ecr.aws/aws-containers/ecsdemo-nodejs" になっている。よって、イメージは事前に用意しておく必要がある。ECS クラスター、サービス、タスク定義などを作成してくれる


[Amazon ECS increases applications resiliency to unpredictable load spikes](https://aws.amazon.com/jp/about-aws/whats-new/2023/10/amazon-ecs-applications-resiliency-unpredictable-load-spikes/)

* ヘルスチェックに失敗した ECS タスクを停止する前にタスクを起動する。その際、新規起動したタスクが healthy となるのを待ってから古い ECS タスクを停止する


[A deep dive into Amazon ECS task health and task replacement](https://aws.amazon.com/jp/blogs/containers/a-deep-dive-into-amazon-ecs-task-health-and-task-replacement/)

* unhealthy タスクの置き換え
  * 2023年10月20日より `maximumPercent` を unhealthy なタスクの置き換えに可能な限り使用する
  * `maximumPercent` が 200 で、8 個中 4 個のタスクがクラッシュした場合。可能な限り早く 4 個のタスクを起動する
  * `maximumPercent` が 200 で、8 個中 4 個のタスクが ELB ヘルスチェックに失敗した場合。4 個のタスクを起動する。healthy になった後に unhealthy なタスクを停止する
  * `maximumPercent` が 150 で、8 個中全てのタスクが ELB ヘルスチェックに失敗した場合。4 個のタスクを起動する。その後、もし 12 個のタスクが healthy になり Auto Scaling により Desired が 10 になった場合は、タスクを 2 個停止する
  * `maximumPercent` が 100 で、タスクがフリーズした場合。タスクが停止する場合は Stopped になった後に新規タスクを起動する
  * `maximumPercent` が 150 で、ローリングアップデート実行中の場合。4 タスクを起動する。旧タスクが unhealthy になった場合は新タスクにて置き換える動作となる
  * `maximumPercent` が 150 で、8 個中全てのタスクが unhealthy になった場合。4 個のタスクを起動するが、これらも unhealthy になる場合。unhealthy なタスクをランダムに 4 個停止する。再び 4 個のタスクを起動する
* 従来は unhealthy なタスクをまず停止していた。背景として EC2 インスタンス上にタスクが詰め込まれており余裕がなかったことがあるが、最近は Fargate であったりキャパシティープロバイダーが使用されている。そのため、`maximumPercent` を可能な限り使用し、unhealthy なタスクも置き換えタスクが healthy になるまでキープする
* unhealthy なタスクを停止すると、残っている healthy な ECS タスクに負荷が集中し、unhelathy になることにつながる


#### Managed Instance

[コンテナ化されたアプリケーション用の Amazon ECS マネージドインスタンスの発表](https://aws.amazon.com/jp/blogs/news/announcing-amazon-ecs-managed-instances-for-containerized-applications/)

* 14 日ごとに開始される定期的なセキュリティパッチの実装
* Bottlerocket 上で動作
* インスタンスタイプを非常に柔軟に選択できる


[Deep Dive: Amazon ECS マネージドインスタンスのプロビジョニングと最適化](https://aws.amazon.com/jp/blogs/news/deep-dive-amazon-ecs-managed-instances-provisioning-and-optimization/)

* 全ての AZ で分散させた後に binpack で配置するので、最小限のインスタンス台数で済む
* 利用率の低いインスタンスは自動的にドレインされ、タスクは既存インスタンスもしくは新規の適切なサイズのインスタンスに配置される



#### Auto Scaling

[Amazon ECS クラスターの Auto Scaling を深く探る](https://aws.amazon.com/jp/blogs/news/deep-dive-on-amazon-ecs-cluster-auto-scaling/)


[新機能 – AWS ECS Cluster Auto ScalingによるECSクラスターの自動スケーリング](https://aws.amazon.com/jp/blogs/news/aws-ecs-cluster-auto-scaling-is-now-generally-available/)


[カスタムメトリクスを用いた Amazon Elastic Container Service (ECS) のオートスケーリング](https://aws.amazon.com/jp/blogs/news/amazon-elastic-container-service-ecs-auto-scaling-using-custom-metrics/)

* リクエストをポーリング、処理するようなワークロードではカスタムメトリクスによる対応が必要。
* スケーリングメトリクスを計算する AWS Lambda 関数をトリガーする。EventBridge によって定期的に実行し、SQS メトリクスをポーチングし、ECS タスクで実行中のキャパシティを考慮して計算する。カスタムメトリクスとして CloudWatch Metrics に送信。
* ターゲット追跡スケーリングポリシーでカスタムメトリクスを使用。


[CloudWatch と Prometheus のカスタムメトリクスに基づく Amazon ECS サービスのオートスケーリング](https://aws.amazon.com/jp/blogs/news/autoscaling-amazon-ecs-services-based-on-custom-cloudwatch-and-prometheus-metrics/)カスタム


[Scale your Amazon ECS using different AWS native services!](https://aws.amazon.com/jp/blogs/containers/scale-your-amazon-ecs-using-different-aws-native-services/)


#### Faster Container Startup

[Improving Amazon ECS deployment consistency with SOCI Index Manifest v2](https://aws.amazon.com/jp/blogs/containers/improving-amazon-ecs-deployment-consistency-with-soci-index-manifest-v2/)
[SOCI Index Manifest v2 を用いた一貫性のある Amazon ECS デプロイメントの実現](https://aws.amazon.com/jp/blogs/news/improving-amazon-ecs-deployment-consistency-with-soci-index-manifest-v2/)

* SOCI v2 ではコンテナイメージと SOCI インデックスの双方向の関連付けにより、一貫性のあるコンテナの実行が可能となる


[AWS Fargate Enables Faster Container Startup using Seekable OCI](https://aws.amazon.com/jp/blogs/aws/aws-fargate-enables-faster-container-startup-using-seekable-oci/)

* 概要
  * スケールアウト時の所要時間の長さは課題の一つ
  * [Research Paper](https://www.usenix.org/conference/fast16/technical-sessions/presentation/harter) によるとコンテナイメージのダウンロードはスタートアップ処理の 76 % の時間を占める。しかし、コンテナが稼働するのに有益なファイルの平均は 6.4 % ほど
  * 従来はコンテナレジストリからイメージをダウンロードし展開していた
  * この方法への対応方法としては Lazy loadinig。アプリケーションの起動と並行してダウンロードを行う
  * 去年、[コンテナイメージを遅延読み込みする Seekable OCI の紹介](https://aws.amazon.com/jp/about-aws/whats-new/2022/09/introducing-seekable-oci-lazy-loading-container-images/) のアナウンスを行った
  * [soci-snapshotter](https://github.com/awslabs/soci-snapshotter/tree/main) を OSS として公開。これは containerd で Lazy loading を有効にするプラグインである
  * Fargate は Seekable OCI (SOCI) をサポートした
  * SOCI は既存コンテナにファイルのインデックスを作成する
  * コンテナイメージを全てダウンロードすることなく、個別のファイルを展開することを可能にする
  * SOCI インデックスはコンテナイメージとは別に生成、保存されるので、コンテナイメージのコンバートのような作業は不要で、イメージの署名内容も影響を受けない
  * Fargate は自動的に SOCI インデックスを検出し、コンテナイメージのプルの完了を待つことなくコンテナプロセスを稼働する
* Let's started
  * Use AWS SOCI Index Builder は CloudFormation で構築できる。Lambda 関数により SOCI インデックスをプッシュする
  * soci-snapshotter が提供する `soci` コマンドにより SOCI インデックスを作成することも可能
  * `nerdctl` によりイメージをプルする。理由としては Docker Engine はデフォルトでは Docker Engine のイメージストアに格納し、containerd のイメージストアには格納しないため
  * SOCI インデックスは SOCI インデックスマニフェスト、zTOCs のセットから構成される。SOCI インデックスマニフェストでは、コンテナイメージのマニフェストの一つのレイヤーごとに ztoc のレイヤーが対応している
  * SOCI インデックスを作成するために `soci` コマンドを使用する。`soci create <イメージ>` で作成可能
  * ztoc skipped となったレイヤについてはプルが完了するのを待ち、その後コンテナプロセスが起動する。スキップされなかったレイヤは lazy loading の対象
  * `sudo soci push --user AWS:$PASSWORD $ECRSOCIURI` により SOCI 関連アーティファクトをプッシュする。ECR のコンソール画面では Artifact Type が SOCI Index, Image Index となっているオブジェクトがプッシュされている
  * SOCI インデックスがあるイメージの場合は RunTask で createdAt、startedAt の間に要する時間が短くなる


#### Fluent Bit

[詳解 FireLens – Amazon ECS タスクで高度なログルーティングを実現する機能を深く知る](https://aws.amazon.com/jp/blogs/news/under-the-hood-firelens-for-amazon-ecs-tasks/)

* FireLens を使えば、ファイルを編集して S3 に再アップロードするだけでよい。イメージをビルドし直す必要はない。
* コンテナの標準出力ログは、Fluentd Docker ログドライバーを介して Unix ソケット経由で FireLens コンテナに送信される。
* FireLens コンテナは、Fluentd Forward Protocol メッセージを TCP ソケットで LISTEN している
* タスク起動時に設定ファイルが自動設定される。
  * ログソース。ログソースは Unix および TCP ソケット
  * ECS メタデータを追加するトランスフォーマー
  * カスタムログを include
  * タスク定義で設定した内容に応じて OUTPUT プラグインの設定

[Fluent Bit による集中コンテナロギング](https://aws.amazon.com/jp/blogs/news/centralized-container-logging-fluent-bit/)


#### Observability

[AWSLogs コンテナログドライバーのノンブロッキングモードによるログ損失の防止](https://aws.amazon.com/jp/blogs/news/preventing-log-loss-with-non-blocking-mode-in-the-awslogs-container-log-driver/)

* デフォルトはブロッキングモード
* ノンブロッキングモードにした場合のログ損失を試験した結果を載せている


[Amazon ECS向けAmazon CloudWatch Container Insightsについて](https://aws.amazon.com/jp/blogs/news/introducing-container-insights-for-amazon-ecs/)


[オブザーバビリティが強化された Container Insights が Amazon ECS で利用可能に](https://aws.amazon.com/jp/blogs/news/container-insights-with-enhanced-observability-now-available-in-amazon-ecs/)

* アカウント設定の箇所で有効化できる。ECS クラスター単位で設定することも可能
* Container Insigths の画面にて、クラスター、インスタンス、サービス、タスクファミリー、タスク、コンテナごとに画面を確認可能
* フィルターオプションで対象を絞ることができる


[AWS Distro for OpenTelemetry コレクターを使用したクロスアカウントの Amazon ECS メトリクス収集](https://aws.amazon.com/jp/blogs/news/using-aws-distro-for-opentelemetry-collector-for-cross-account-metrics-collection-on-amazon-ecs/)


[Centralized Amazon ECS task logging with Amazon OpenSearch](https://aws.amazon.com/jp/blogs/containers/centralized-amazon-ecs-task-logging-with-amazon-opensearch/)

* Fluent Bit により OpenSearch にログ転送を行う


#### FIS

[AWS Fault Injection Simulator の Amazon ECS に関する新機能のお知らせ](https://aws.amazon.com/jp/blogs/news/announcing-aws-fault-injection-simulator-new-features-for-amazon-ecs-workloads/)

* 「CPU 負荷をかける」「ストレージ I/O 負荷をかける」「プロセスの停止」「ネットワークトラフィックの停止」「ネットワークレイテンシーの増加」「パケットロス」などのフォールとインジェクションアクションが追加された
* SSM Agent がサイドカーとして動作し、Run Command により障害試験を行う。そのため、SSM Agent のサイドカーが必要
* CloudWatch アラームと連携し、アラームが発火した時に試験を止めることができる


#### EFS

[Amazon EFS を Amazon ECS と AWS Fargate で使用するための開発者ガイド – パート 1](https://aws.amazon.com/jp/blogs/news/developers-guide-to-using-amazon-efs-with-amazon-ecs-and-aws-fargate-part-1/)

* EFS のマウントは PV 1.4.0 から
* エンドソリューションのパフォーマンス、冗長性、可用性、柔軟性のレベルに応じてボリュームストレージを選定する
* EC2 の場合はこれまでも使えていたが、Fargate でもできるようになった。また EC2 の場合はセットアップが必要だったが、それも省力化されることになる
* S3 にデータを置いている場合は ECS タスクのタスクストレージにダウンロードする必要があるが、このようなユースケースにおいて EFS は便利


[Amazon EFS を Amazon ECS と AWS Fargate で使用するための開発者ガイド – パート 2](https://aws.amazon.com/jp/blogs/news/developers-guide-to-using-amazon-efs-with-amazon-ecs-and-aws-fargate-part-2/)

* EFS のセキュリティの観点は 2 つ
  * ネットワークセキュリティ: EFS マウントターゲットへの疎通性
  * クライアントセキュリティ: 読み書きの検眼があるかどうか
* クライアントセキュリティ
  * EFS 側のリソースベースとクライアント側のアイデンティティベースの 2 つがある
  * 以下のアクションを設定可能
    * elasticfilesystem:ClientMount (読み取り専用アクセス)
    * elasticfilesystem:ClientWrite (読み取り/書き込みアクセス)
    * elasticfilesystem:ClientRootAccess (root アクセス)
  * ClientWrite が許可されていない場合は書き込みできない
  * リソースレベルのポリシーがない場合は、全て許可
  * クライアント側の権限では ECS タスクロールを使用。これがない場合は匿名として識別される
* アクセスポイント
  * アクセスポイントを介した場合、コンテナ内の UID, GID にかかわらず、アクセスポイントに設定された UID, GID でアクセスが行われる
* 対応可能なユースケース
  * ECS サービスごとにアクセスポイントを分けることで、互いに独立したディレクトリに対して読み書き可能。ただし、EFS スループットは共有される
  * 同一のファイルシステムを複数 ECS サービスで利用。ECS サービスによって Read のみ, Read/Write 可能を設定可能
  * 特定のクライアントのみアクセス可能な制限付きのディレクトリを設定できる


[Amazon EFS を Amazon ECS と AWS Fargate で使用するための開発者ガイド – パート 3](https://aws.amazon.com/jp/blogs/news/developers-guide-to-using-amazon-efs-with-amazon-ecs-and-aws-fargate-part-3/)


#### Gen AI

[Automating AI-assisted container deployments with the Amazon ECS MCP Server](https://aws.amazon.com/jp/blogs/containers/automating-ai-assisted-container-deployments-with-amazon-ecs-mcp-server/)

* Amazon ECS MCP Server は development, deployment, operations, troubleshooting, and decommissioning まで対応できる
* 以下のアクションがある
  * containerize_app: Development	Provides guidance to create docker file
  * create_ecs_infrastructure	Deployment: Provisions Amazon ECS infrastructure using CloudFormation
  * get_deployment_status	Deployment: Monitors deployment status
  * ecs_resource_management	Operations: Manages resource inventory
  * ecs_troubleshooting_tool: Troubleshooting	Diagnoses Amazon ECS issues across services and tasks
  * delete_ecs_infrastructure: Decommissioning	Helps with resource cleanup


[Accelerate container troubleshooting with the fully managed Amazon ECS MCP server (preview)](https://aws.amazon.com/blogs/containers/accelerate-container-troubleshooting-with-the-fully-managed-amazon-ecs-mcp-server-preview/)

* [ECS のドキュメント](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-mcp-introduction.html) にもまとめられている
* Remote MCP Server が使用できる。MCP Client → aws-mcp-proxy → MCP Service の構成で使用
  * `~/.kiro/settings/mcp.json` に設定を行う
```json
{
  "ecs-mcp-pdx": {
    "disabled": false,
    "timeout": 60000,
    "command": "uvx",
    "args": [
      "mcp-proxy-for-aws@latest",
      "https://ecs-mcp.us-west-2.api.aws/mcp",
      "--service",
      "ecs-mcp",
      "--profile",
      "default",
      "--region",
      "us-west-2"
    ],
    "type": "stdio"
  }
}
```
* コンソールの Amazon Q からも使用できる。タスク停止時などのトラブルシューティングができる


#### その他の機能

[New – Amazon ECS Exec による AWS Fargate, Amazon EC2 上のコンテナへのアクセス](https://aws.amazon.com/jp/blogs/news/new-using-amazon-ecs-exec-access-your-containers-fargate-ec2/)


[Amazon ECS deployment circuit breaker のご紹介](https://aws.amazon.com/jp/blogs/news/announcing-amazon-ecs-deployment-circuit-breaker-jp/)


[Amazon ECS on AWS Fargate で設定可能な Linux パラメータの追加](https://aws.amazon.com/jp/blogs/news/announcing-additional-linux-controls-for-amazon-ecs-tasks-on-aws-fargate/)

* Fargate タスクにおいても Linux カーネルパラメータを調整できるようになった
  * 以下のカーネルパラメータは要望が多かった
    * `net.core.somaxconn`
    * `net.ipv4.ip_local_port_range`
  * `containerDefinitions.systemControls` にて namespace, value を設定可能
* コンテナ間で PID namespace を共有できるようになった
  * `pidMode` を `task` に設定することで PID namespace を設定可能
  * サイドカーコンテナにアクセスした場合、アプリケーションコンテナのプロセスも確認できる。nginx プロセスが PID 7 になっていたりする
  * `pause` プロセスが PID 1 になる


[AWS Fargate タスクのリタイア通知による運用の可視性の向上](https://aws.amazon.com/jp/blogs/news/improving-operational-visibility-with-aws-fargate-task-retirement-notifications/)

* Fargate タスクがリタイアするまでに待機する日数を指定できる
* アカウント設定の `fargateTaskRetirementWaitPeriod` で設定できる
* Slack にリタイア通知を送る [サンプル](https://github.com/aws-samples/capturing-aws-fargate-task-retirement-notifications/tree/main) が提供されている。イベントのテストは `aws events put-events` で行うことができる
  * Lambda 関数内の以下コードで Slack チャンネルにポストしている
```python
    # Post the message to the Slack channel.
    req = Request(
        slack_uri,
        data=json.dumps(slack_message).encode('utf-8'),
        headers={'content-type': 'application/json'}
    )
```


[Effective use: Amazon ECS lifecycle events with Amazon CloudWatch logs insights](https://aws.amazon.com/jp/blogs/containers/effective-use-amazon-ecs-lifecycle-events-with-amazon-cloudwatch-logs-insights/)



#### Others

[お誕生日おめでとう！AWS Fargate 5 周年](https://aws.amazon.com/jp/blogs/news/happy-5th-birthday-aws-fargate/)


[Amazon ECS の 10 周年を祝う: 10 年間にわたるコンテナ化イノベーションの推進](https://aws.amazon.com/jp/blogs/news/celebrating-10-years-of-amazon-ecs-powering-a-decade-of-containerized-innovation/)


[Bottlerocket のセキュリティ機能 〜オープンソースの Linux ベースオペレーティングシステム〜](https://aws.amazon.com/jp/blogs/news/security-features-of-bottlerocket-an-open-source-linux-based-operating-system/)


[Amazon ECS でのデーモンサービスの改善](https://aws.amazon.com/jp/blogs/news/improving-daemon-services-in-amazon-ecs/)


[Amazon ECS と AWS Fargate を利用した Twelve-Factor Apps の開発](https://aws.amazon.com/jp/blogs/news/developing-twelve-factor-apps-using-amazon-ecs-and-aws-fargate/)


[Amazon ECS on AWS Fargate のコスト最適化チェックリスト](https://aws.amazon.com/jp/blogs/news/cost-optimization-checklist-for-ecs-fargate/)

* タグを使用するにはアカウント設定で「新しい Amazon リソースネーム (ARN) とリソース識別子 (ID) 形式をオプトイン」しておく必要がある。
* コストエクスプローラーによって可視化。
* Savings Plans によるコスト削減。
* Fargate Spot の使用。
* タスクの適切なサイジング。
* Auto Scaling によって適切なタスク数を稼働。
* 営業時間外にタスクを停止するようにスケジューリング。


[Amazon ECS Fargate/EC2 起動タイプでの理論的なコスト最適化手法](https://aws.amazon.com/jp/blogs/news/theoretical-cost-optimization-by-amazon-ecs-launch-type-fargate-vs-ec2/)


[Serverless containers at AWS re:Invent 2024](https://aws.amazon.com/jp/blogs/containers/serverless-containers-at-aws-reinvent-2024/)



## Black Belt

[[AWS Black Belt Online Seminar] AWS コンテナサービス開始のおしらせ](https://aws.amazon.com/jp/blogs/news/aws-bb-containers-start/)


[[AWS Black Belt Online Seminar] CON246 ログ入門 資料公開](https://aws.amazon.com/jp/blogs/news/aws-black-belt-online-seminar-con246-log/)


[[AWS Black Belt Online Seminar] CON245 Configuration & Secret Management 入門 資料公開](https://aws.amazon.com/jp/blogs/news/aws-black-belt-online-seminar-con245-config/)


[202109 AWS Black Belt Online Seminar Auto Scaling in ECS](https://www.slideshare.net/AmazonWebServicesJapan/202109-aws-black-belt-online-seminar-auto-scaling-in-ecs-250178830)


[202109 AWS Black Belt Online Seminar Amazon ECS Capacity Providers](https://www.slideshare.net/AmazonWebServicesJapan/202109-aws-black-belt-online-seminar-amazon-ecs-capacity-providers)


[202109 AWS Black Belt Online Seminar Amazon Elastic Container Service − EC2 スポットインスタンス / Fargate Spot ことはじめ](https://www.slideshare.net/AmazonWebServicesJapan/202109-aws-black-belt-online-seminar-amazon-elastic-container-service-ec-fargate-spot)


## builders.flash

[Web アプリケーションにおける Amazon ECS / AWS Fargate アーキテクチャデザインパターン](https://aws.amazon.com/jp/builders-flash/202409/web-app-architecture-design-pattern/)

* 1. パブリック API サービスパターン
  * ALB - ECS
  * API Gateway - ALB - ECS
* 2. SPA パターン
  * CloudFront - S3
  * CloudFront - API Gateway - NLB - ECS
* 3. SSR パターン
  * ALB - ECS(Rendering) - ECS(Backend API)
* 4. ソーシャルログインパターン
  * ALB と Cognito を連携。Cognito では外部 IdP と連携
* 5. 内部サービス連携パターン
  * 内部 ALB 接続パターン
  * ECS サービス検出パターン
  * App Mesh 接続パターン
  * ECS Service Connect パターン
* 6. ジョブ構成パターン
  * ジョブ定期実行パターン
  * ジョブワークフローパターン (Step Functions によるエラーハンドリングも可能)
  * ジョブ API 実行パターン
* 7. ログ運用パターン
  * CloudWatch Logs 連携パターン
  * ログ長期保管向け S3 連携パターン
* 8. アラート通知パターン
  * アプリケーションエラー通知パターン (CloudWatch Logs の文字列をもとに通知)
    * CloudWatch Logs - SNS - Chatbot - Slack
  * ECS アラート通知パターン
    * タスク状態変更イベントなどを通知
* 9. ECS タスク CI/CD パターン
  * GitHub Actions による CI/CD パターン
  * Code シリーズによる CI/CD パターン
* 10. 開発デバッグパターン
  * ECS exec 利用パターン
  * CloudShell 利用パターン



## tori さん

[Amazon ECS でのコンテナデプロイの高速化](https://toris.io/2021/04/speeding-up-amazon-ecs-container-deployments/)

コンテナのデプロイ速度を高速化するテクニックが書かれている。

* ヘルスチェックの間隔、しきい値を少ない値に設定する。
* 登録解除の遅延を少ない値に設定する。
* ECS_CONTAINER_STOP_TIMEOUT を少ない値に設定する。（SIGTERM 後に SIGKILL を送るまでの時間）
* ECS_IMAGE_PULL_BEHAVIOR を prefer-cached あるいは once に設定する。イメージタグをコミットごとに変えている場合に有効な戦略(latest はアンチパターン)
* minimumHealthyPercent を小さめに maximumPercent を多めに設定することで、より少ないステップ数でデプロイが完了するようにする。


[アプリケーション開発者は Amazon ECS あるいは Kubernetes をどこまで知るべきか](https://speakerdeck.com/toricls/you-build-it-you-run-it)



## ECS Agent

[aws/amazon-ecs-agent](https://github.com/aws/amazon-ecs-agent)

[amazon-ecs-agent/proposals/eni.md](https://github.com/aws/amazon-ecs-agent/blob/master/proposals/eni.md)



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


[【レポート】Amazon ECS deployment circuit breaker を使った自動ロールバック #AWSSummit](https://dev.classmethod.jp/articles/awssummit2021-ecs-deployment-circuit-breaker/)

Amazon ECS deployment circuit breaker

* デプロイタイプ、ローリングアップデートでサポートされる。
* ステージ 1, 2 の 2 段階で評価。
  * ステージ1: タスクが RUNNING 状態に遷移することなく停止した場合。
  * ステージ2: RUNNING 状態に遷移したタスクがヘルスチェックに合格した場合。
* しきい値は 10 <= Desired Count * 0.5 <= 200。


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


[AWS CDKでECS on FargateのCI/CDを実現する際の理想と現実](https://speakerdeck.com/tomoki10/ideal-and-reality-when-implementing-cicd-for-ecs-on-fargate-with-aws-cdk)


[Compute EngineのNested Virtualizationを使ってFirecrackerの開発環境を構築してみた](https://dev.classmethod.jp/articles/setup-firecracker-devenv-on-compute-engine/)


[ECS Service Connectによるサービスの新しいつなぎ方](https://speakerdeck.com/iselegant/a-new-way-to-connect-services-with-ecs-service-connect?slide=29)


[ECS Service Connect はヘルスチェックの結果を元にトラフィックをルーティングしないので、起動に時間がかかるアプリケーションでは注意が必要です](https://dev.classmethod.jp/articles/ecs-service-connect-application-health-check-and-traffic-routing/)

* アプリケーションの起動が安定するまでの間にトラフィックを流す [Issue](https://github.com/aws/containers-roadmap/issues/2334) がある。対策としては 1 個コンテナを増やし、対象コンテナが HEALTHY となるように依存関係を設定するとよい

