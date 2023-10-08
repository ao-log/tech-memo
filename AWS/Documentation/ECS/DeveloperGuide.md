# ECS


## 用語

* クラスター
* サービス
* タスク
* タスク定義
* ECS エージェント
* 起動タイプ


## ECS とは

[Amazon Elastic Container Service とは](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/Welcome.html)

EC2, Fargate 二つの起動タイプがある。

Fargate 起動タイプに向いているワークロード

* 低いオーバーヘッドのために最適化する必要がある大規模なワークロード
* 時折バーストが発生する小さなワークロード
* 小さなワークロード
* バッチワークロード

EC2 起動タイプに向いているワークロード

* 一貫して高 CPU コアとメモリ使用量を必要とするワークロード
* 料金のために最適化する必要がある大規模なワークロード
* アプリケーションは永続的ストレージにアクセスする必要があります
* インフラストラクチャを直接管理する必要があります

管理方法

* マネジメントコンソール
* AWS CLI
* AWS SDK
* AWS Copilot
* ECS CLI
* AWS CDK


## ECS コンポーネント

* クラスター
* サービス
* タスク
* タスク定義
* コンテナエージェント


## チュートリアル

[Getting started with Amazon ECS using Amazon EC2](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/getting-started-ecs-ec2.html)

次の流れについて書かれている。

* タスク定義の作成
* ECS クラスターの作成（ここでは、EC2 Linux + Networking で作成している）
* サービスの作成


[Getting started with Amazon ECS using Fargate](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/getting-started-fargate.html)


[AWS CDK の使用](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/tutorial-ecs-web-server-cdk.html)

次のようなコードで ALB に紐づいた ECS サービスを作成可能。

```ts
import * as cdk from '@aws-cdk-lib';
import { Construct } from 'constructs';

import * as ecs from '@aws-cdk-lib/aws-ecs';
import * as ecsp from '@aws-cdk-lib/aws-ecs-patterns';

export class HelloEcsStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    new ecsp.ApplicationLoadBalancedFargateService(this, 'MyWebServer', {
      taskImageOptions: {
        image: ecs.ContainerImage.fromRegistry('amazon/amazon-ecs-sample'),
      },
      publicLoadBalancer: true
    });
  }
}
```


## Fargate

[https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/AWS_Fargate.html](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/AWS_Fargate.html)

* 各 Fargate タスクは、独自の分離境界を持ち、基本となるカーネル、CPU リソース、メモリリソース、または Elastic Network Interface を別のタスクと共有しない。
* 一部のタスク定義パラメータは無効、もしくは制限されている。
  * disableNetworking
  * dnsSearchDomains
  * dnsServers
  * dockerSecurityOptions
  * extraHosts
  * gpu
  * ipcMode
  * links
  * pidMode
  * placementConstraints
  * privileged
  * systemControls
  * maxSwap
  * swappiness
* dockerVolumeConfiguration はサポートされない。ホストボリュームのみサポートされる。
* **ネットワークモードは awsvpc にする必要あり。タスクごとに ENI が割り当てられる。**
* タスクごとに CPU, メモリの指定が必要。タスクレベル CPU とメモリの有効な組み合わせの表を参照のこと。
* ulimit は nofile のみ上書き可能
* タスクストレージ
  * PV 1.4.0 以降は 20 GB のエフェメラルディスク。
  * PV 1.3.0 以前は 10 GB の Docker レイヤーストレージ。4 GB のリュームをマウント可能（タスク定義の volumes, mountPoints, volumesFrom で指定）
* **ALB, NLB を使用する際はターゲットグループのターゲットタイプを ip にする必要がある。** NLB の UDP 転送は PV 1.4.0 以上かつ一部リージョンのみ対応。
  * プライベートレジストリはシークレットマネージャに認証情報を設定し、タスク定義で設定することで利用可能。 


## Platform Version

プラットフォームのバージョンに関する仕様は次のドキュメントをご参照のこと。

[AWS Fargate platform versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html)

PV 1.2.0 とそれ以前のバージョンは 2020/12/14 に Deprecated となる。


[AWS Fargate platform versions scheduled for deprecation](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform-versions-retired.html)

PV 1.4.0 については BlackBelt にもまとめられている。


[AWS Black Belt Online Seminar Container Services Update P10](https://www.slideshare.net/AmazonWebServicesJapan/20200624-aws-black-belt-online-seminar-container-services-update/10)

* コンテナランタイムが Docker から containerd に変更。
* ECS Agent ではなく Fargate Agent。
* ECR からのログイン情報、SSM, Secrets Manager からの情報取得は、サービス側の ENI ではなく Task ENI を通るようになった。
* タスクメタデータエンドポイントの /stats からモニタリング情報を取得できる。
* UDP をルーティング可能
* タスクストレージサイズの変更。デフォルト 20 GB。



## ECS Cluster

[Amazon ECS clusters](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/clusters.html)

* サービス、タスクを管理する単位
* EC2, Fargate の両起動タイプが共存可能
* マネジメントコンソールからクラスターを作成すると、CloudFormation のスタックが作られる
* クラスターのデフォルトの Service Connect 名前空間を設定可能


[コンソールを使用した Fargate 起動タイプ用のクラスター作成](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/create-cluster-console-v2.html)

デフォルトでは以下の内容でクラスターが作成される。一部は変更可能

* Fargate および Fargate Spot キャパシティープロバイダーを使用
+ 選択したリージョンのデフォルト VPC 内のすべてのデフォルトサブネットでタスクとサービスを起動
* Container Insights は使用しない
* AWS CloudFormation に 3 つのタグが構成
* AWS Cloud Map に、クラスターと同じ名前のデフォルトの名前空間を作成


#### Capacity Provider

[Amazon ECS capacity providers](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-capacity-providers.html)

**Capacity Provider**

* クラスターは複数のキャパシティプロバイダーを持つことができる。キャパシティープロバイダーを作成した後、`PutClusterCapacityProviders` によりクラスターに関連づける対応が必要
* デフォルトのキャパシティプロバイダー戦略がクラスターに設定されており、サービスもしくはスタンドアローンのタスクにおいて、カスタムキャパシティープロバイダー戦略もしくは起動タイプが設定されていない場合に使用される
* Fargate の場合は、FARGATE、FARGATE_SPOT を使用できる。
* EC2 起動タイプの場合は、次の 3 つの設定項目がある
  * Auto Scaling group
  * managed scaling
  * managed termination protection

**Capacity Provider Strategy**

* サービス、またはタスクの設定時に、どのキャパシティプロバイダー戦略を使用するかを設定可能
* **キャパシティプロバイダー戦略**により、タスクをどのキャパシティプロバイダーに配置するかを設定できる。base, weight の２つの設定値がある
  * base: 最低何個のタスクを起動するか。一つのキャパシティープロバイダーにのみ設定可能
  * weight: どのキャパシティプロバイダーにタスクを割り当てるかの比率を設定する


[AWS Fargate capacity providers](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/fargate-capacity-providers.html)

* マネジメントコンソールより ECS クラスターを Networking only で作成した場合に FARGATE, FARGATE_SPOT が自動的に設定された状態になっている。
* FARGATE_SPOT のキャパシティプロバイダーのタスクがスポットの中断により停止する際は、タスク停止の 2 分前に EventBridge よりワーニングが送られる。また、SIGTERM シグナルがタスクに送られる。キャパシティに空きがある場合は新規タスクの起動を試みる
* FARGATE_SPOT は Windows では未サポート。Linux でも ARM64 の場合は未サポート


[Auto Scaling group capacity providers](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/asg-capacity-providers.html)

* 空の Auto Scaling グループの作成を推奨。既存の Auto Scaling グループの場合、起動済みのインスタンスが正常にキャパシティープロバイダーに登録されないことがある
* managed termination protection
  * 使用する際 managed scaling が有効になっている必要がある。そうしないと managed termination protection は動作しない
  * 有効にすることでタスクが存在する EC2 インスタンスについてスケールインから保護することができる
* **managed scaling を有効にすることで、タスク数に応じて ECS インスタンスがスケールする。** スケーリング用のリソースは変更してはならない
* ウォームプールを使用可能
  * ユーザーデータで `ECS_WARM_POOLS_CHECK` を設定する


[Amazon ECS クラスターの Auto Scaling](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/cluster-auto-scaling.html)

* スケーリングの制御に `CapacityProviderReservation` のメトリクスを使用する。
  * 計算式: CapacityProviderReservation = (number of instances needed) / (number of running instances) x 100
  * `CapacityProviderReservation` > `targetCapacity` の場合にスケールアウトする
  * `CapacityProviderReservation` < `targetCapacity` の場合にスケールインする。
* 考慮事項
  * スケーリングポリシーを変更、追加してはならない
  * スケーリングにはサービスにリンクされたロール `AWSServiceRoleForECS` を使用する
  * キャパシティプロバイダーを作成、更新する IAM エンティティには `autoscaling:CreateOrUpdateTags` の許可が必要
  * Auto Scaling グループがら AmazonECSManaged タグを削除してはならない
* Managed termination protection
  * タスクが稼働しているインスタンスはスケールインから保護される。しかし DAEMON タイプは例外
* マネージドスケールアウトの動作
  * どのキャパシティープロバイダーを使用するかを決定
  * Auto Scaling グループが複数のインスタンスタイプから構成される場合は vCPU, メモリ, ENI, GPU 数からインスタンスタイプを決定
  * インスタンス数の算出。binpack で計算
  * `CapacityProviderReservation` メトリクスを更新
  * `CapacityProviderReservation` と `targetCapacity` を比較
  * スケールアウトがオーバープロビジョニングすることを防ぐ動作が組み込まれている。Auto Scaling は全てのインスタンスについて `instanceWarmupPeriod` が経過したかを確認する。スケールアウトは `instanceWarmupPeriod` が経過するまでの間はブロックされる。
  * 考慮事項
    * インスタンス数が 0 の場合は 2 台にスケールアウトする
    * `instanceWarmupPeriod` は EC2 インスタンスが起動し ECS Agent が開始するのに十分な時間とする。さもなければオーバープロビジョニングが発生しうる


## Amazon ECS Task definitions

**TODO**

タスクの起動方法を定義。以下のような項目を設定する。

* Docker イメージ
* CPU、メモリ量（タスク、もしくはタスク内のコンテナ）
* 起動タイプ
* Docker ネットワーキングモード
* ロギング設定
* コンテナ終了時の動作
* コンテナ開始時に実行するコマンド
* データボリューム
* IAM ロール

一つのタスク内に複数コンテナを設定可能。

タスク定義は更新することはできない。代わりに新しいリビジョンを作成して対応する。

タスク定義のリビジョンは INACTIVE にすることもできる。ただし、そのリビジョンを使用しているサービスからはそのリビジョンを引き続き使用できる。できなくなるのは、そのリビジョンからの新たなタスク生成のみ。


[Task definition parameters](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html)

パラメータのリファレンス。

**Family**

* family: タスク定義名。

**Task Role**

* taskRoleArn: タスク内のコンテナが利用可能な権限を設定するロール。(**EC2 のインスタンスプロファイルのようなもの**)

**Task Execution Role**

* executionRoleArn: タスクの実行に使用するロール（**ECR や CloudWatch Logs**）。

**Network Mode**

* networkMode: none, bridge, awsvpc, host がある。
  * none: 外部接続性を持たない。ポートマッピングも設定できない。
  * bridge: Docker ビルトインの仮想ネットワークを使用。
  * host: Docker ビルトインの仮想ネットワークをバイパスし、コンテナポートとホストポートを直接マッピングする。
  * awsvpc: ENI にタスクが割り当てられる。Fargate の場合はこのモードになる。

**Standard Container Definition Parameters**

* name: コンテナに付ける名前
* image: イメージの指定
* memory: メモリ量(MiB)。指定量を超えて使用しようとした場合、コンテナは kill される。
* memoryReservation: メモリ量のソフトリミット。
* portMappings: コンテナとホストのポートのマッピング。
  * containerPort: コンテナのポート。
  * hostPort: ホストのポート。**無指定でかつコンテナポートが設定されている場合は、エフェメラルポートから割り当てる。**
  * protocol: tcp or udp。

**Advanced Container Definition Parameters**

* healthCheck: タスクが手動起動された場合は、ヘルスチェック結果に関わらず継続稼働する。サービスから起動された場合は、unhealthy のタスクを終了し、新たなタスクを起動する。
  * command: ヘルスチェックの判定に使用するコマンド。
  * interval: 間隔。デフォルト 30 秒。5 〜 300 秒で設定可能。
  * timeout: 失敗と判定するまで待つ時間。デフォルトは 5 秒。2 〜 60 秒で設定可能。
  * retries: unhealthy と判定するまでにリトライする回数。デフォルトは 3 秒。1 〜 10 回で設定可能。
  * startPeriod  コンテナ起動時にヘルスチェックを猶予する時間。
* cpu: タスクが使用できる CPU。インスタンス上に複数タスクがある場合はこの設定値に基づいた比率となる。
* gpu: GPU の個数。
* essential: true に設定されている場合、このコンテナがダウンすると他のコンテナも停止させる。
* entryPoint: Docker コンテナ起動時の Entrypoint にマップされる。
* command: Docker コンテナ起動時の Cmd にマップされる。
* workingDirectory: Docker コンテナ起動時の WorkingDir にマップされる。
* environmentFiles: docker run オプションの --env-file にマップされる。S3 上のオブジェクトを指定する。Fargate 起動タイプでは使用不可。
* environment: Docker コンテナ起動時の Env にマップされる。
* secrets: シークレットマネージャ or SSM パラメータストアの ARN を指定。
* disableNetworking: Docker コンテナ起動時の NetworkDisabled にマップされる。
* links: network mode が bridge の場合のみ使用可。ポートマッピング無しでコンテナが相互に通信できるようになる。
* hostname: hostname を設定できる。
* dnsServers: DNS サーバを設定できる。
* dnsSearchDomains: DNS の検索対象ドメインを設定できる。
* extraHosts: /etc/hosts に追記される。
* readonlyRootFilesystem: root ファイルシステムを read only に設定できる。
* mountPoints: データボリュームをマウントできる。
  * sourceVolume: マウントするボリューム名。
  * containerPath: コンテナのマウントパス。
  * readOnly: readonly でマウントするかどうか。
* volumesFrom: 異なるコンテナからデータボリュームをマウントできる。
* logConfiguration: ログに関する設定
  * logDriver: "awslogs","fluentd","gelf","json-file","journald","logentries","splunk","syslog","awsfirelens" から設定可能。
* privileged: コンテナに特権を与える。
* user: コンテナ内で使用するユーザ名。
* dockerSecurityOptions: SELinux, Apparmor 用のカスタムラベルを設定。
* ulimits: ulimit でを設定。
* dockerLabels: Docker コンテナ起動時の Labels にマップされる。

**Other Container Definition Parameters**

* linuxParameters: Linux パラメータの設定。
* dependsOn: コンテナの依存関係を設定。
* startTimeout: 依存関係の解決を行うまでに待機する時間。依存先のコンテナがタイムアウト秒数内に期待する状態にならない場合、ギブアップしコンテナを起動しない。
* stopTimeout: コンテナ終了までに待機する時間。
* systemControls: カーネルパラメータを設定。
* interactive: true の場合、stdin or tty が割り当てられる。
* pseudoTerminal: true の場合、TTY が割り当てられる。

**Volumes**

Docker volumes — Docker-managed volume でコンテナインスタンスの /var/lib/docker/volumes に作られるもの。EC2 起動タイプでのみサポートされる。

**Bind mounts** — ホストマシンのファイル、ディレクトリをコンテナにマウントするもの。EC2 or Fargate 起動タイプでサポートされる。

* name: ボリューム名。
* host: ボリュームマウントを使用する際に設定。
* dockerVolumeConfiguration: Docker Volumes を使用する際に設定。
* efsVolumeConfiguration: EFS を使用する際に設定。

**Task placement constraints**

Fargate の場合はサポートされない。**Fargate の場合はデフォルトで spread across Availability Zones となる。**

* expression: 制約に関する記述。Cluster query language で記述。
* type: 制約のタイプ。

**Launch types**

* requiresCompatibilities: EC2 or FARGATE

**Task size**

* cpu: 1024 = 1 vcpu
* memory: 1024 = 1 GB

**Proxy configuration**

* proxyConfiguration: App Mesh proxy の設定。

**Other task definition parameters**

* ipcMode: IPC resource namespace を使用。
* pidMode: process namespace を使用。


#### 追加リソースのサポート

* [Working with GPUs on Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-gpu.html)
  * エージェント設定ファイルで ECS_ENABLED_GPU_SUPPORT を true に設定する必要がある。
* [Working with inference workloads on Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-inference.html)


#### データボリュームの使用

[Using data volumes in tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_data_volumes.html)


[Fargate Task Storage](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/fargate-task-storage.html)

* PV 1.4.0 以降は 20 GiB のエフェメラルストレージ。200 GiB に拡張可能。
* PV 1.3.0 以前は 10 GiB の Docker レイヤストレージ。ボリュームマウント用の 4 GiB の領域。


[Amazon EFS ボリューム](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/efs-volumes.html)

* Amazon ECS-optimized AMI version 20200319 with container agent version 1.38.0 からサポート。
* Fargate の場合は PV 1.4.0 からサポート。


[Docker volumes](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/docker-volumes.html)

* コンテナインスタンスの /var/lib/docker/volumes が使用される。
* EC2 起動タイプのみで対応。


[Bind mounts](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/bind-mounts.html)

* ホストマシンの特定のファイル、ディレクトリをコンテナにマウント可能。
* EC2、Fargete の両起動タイプで使用可能。
* コンテナイメージ内の所定ディレクトリの内容を他のコンテナと共有することも可能。Dockerfile では当該ディレクトリを VOLUME で記載しておく必要がある。
* volumesFrom を使用して他のコンテナと VOLUME で記載したディレクトリを共有することも可能。


[コンテナスワップ領域の管理](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/container-swap.html)

* コンテナのメモリ不足を避ける用途で有用。
* EC2 起動タイプのみで使用可能。



#### Networking

[Task Networking with the awsvpc Network Mode](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html)

* awsvpc ネットワークモードの場合、タスクに一つ ENI が割り当てられる。そのため、VPC フローログにも記録される。この ENI はリクエスタマネージド型。
* awsvpcTrunking をオプトインしている場合は、trunk の ENI をアタッチする。
* タスク内のコンテナ間通信は localhost で行うことができる。
* Linux インスタンスの場合はタスク ENI にパブリック IP アドレスが付与されない。よって NAT Gateway もしくは VPC エンドポイントを使用する構成にする必要がある。
* タスク定義内のコンテナが開始される前に、各タスクに Amazon ECS コンテナエージェントによって追加の pause コンテナが作成される。次に、amazon-ecs-cni-plugins CNI プラグインを実行して pause コンテナのネットワーク名前空間が設定され流。その後、エージェントによってタスク内の残りのコンテナが開始されま流。この手順により、pause コンテナのネットワークスタックが共有される。つまり、タスク内のすべてのコンテナは ENI の IP アドレスによってアドレス可能であり、localhost インターフェイス経由で相互に通信できるようになる。
* ELB サポートは ALB, NLB のみ。CLB はサポートされない。

『ベストプラクティスガイド』の [ネットワークモードの選択](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/bestpracticesguide/networking-networkmode.html) も各ネットワークモードの参考になる。

ブリッジの場合は、同一インスタンス内で複数タスクのコンテナが同一ポートを使用したい場合、動的ポートマッピングを使用するとよい。



#### Logging

[Using the awslogs log driver](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_awslogs.html)

* awslogs ログドライバーは STDOUT, STDERR の IO ストリームを CloudWatch Logs に送る。
* タスク実行ロールに logs:CreateLogStream および logs:PutLogEvents の許可が必要。
* タスク定義のオプション
  * awslogs-create-group
  * awslogs-region
  * awslogs-group
  * awslogs-stream-prefix
  * awslogs-datetime-format
  * awslogs-multiline-pattern
  * mode
  * max-buffer-size


[Custom log routing](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html)

* Firelens をログルーターとして使用することも可能。サイドカーとして稼働させる。
* OUTPUT が様々なものに対応している。firehose など。


[FireLens 設定を使用するタスク定義の作成](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/firelens-taskdef.html)

* config-file-type でカスタム設定ファイルのソースの場所を指定。s3 or file。EC2 の場合は s3 も指定できる。Fargate は file のみ。


#### Authentication

[Private registry authentication for tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/private-auth.html)

* プライベートレジストリの認証を行うことが可能。認証情報は Secrets Manager に格納しておく。


#### secret

[Specifying sensitive data using Secrets Manager](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/specifying-sensitive-data-secrets.html)


[Specifying sensitive data using Systems Manager Parameter Store](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/specifying-sensitive-data-parameters.html)

認証情報は Secrets Manager、Systems Manager Parameter Store に格納して、取り出すことが可能。環境変数もしくはログ設定情報として使用可能。


#### Environment Variables

[Specifying environment variables](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/taskdef-envfiles.html)

次の設定方法が可能。

* environment
* environmentFiles(S3 上のオブジェクトを指定)


## アカウント設定

[アカウント設定](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-account-settings.html)

* アカウントレベルもしくは IAM ユーザー、ロールレベルで特定の機能をオプトイン、オプトアウトできる
* 設定できる項目
  * Amazon リソースネーム (ARN) と ID
  * AWS VPC トランキング
  * CloudWatch Container Insights
  * デュアルスタック VPC IPv6
  * Fargate FIPS-140 コンプライアンス
  * タグリソース認可


## ECS container instances

[Amazon ECS container instances](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ECS_instances.html)

* ECS エージェントの導入が必須。Amazon ECS-optimized AMI の場合は既に導入済み。
* 適切な IAM ロールの設定が必要。
* Linux Amazon ECS-optimized AMI version 20200430 and later の場合、Amazon EC2 Instance Metadata Service Version 2 (IMDSv2) がサポートされている。それ以前の AMI の場合は IMDSv1。
* コンテナ側で expose するポートに対して、インバウンド通信の許可が必要。
* Amazon ECS service endpoint との接続性が必要。パブリック IP アドレスもしくはインタフェース VPC エンドポイントを用意する必要あり。
* コンテナインスタンス内で固有の情報を持っているので、登録解除してから別クラスターに登録することはしないように。
* インスタンスタイプの変更は不可。

**ライフサイクル**

* ACTIVE 状態の場合 RunTask API によるリクエストを受け入れることができる。
* FALSE: コンテナインスタンスを停止ししばらく経つと FALSE に遷移する。
* DRAINING: 新規タスクが配置されなくなる。サービスのタスクは可能であれば削除される。
* INACTIVE: コンテナインスタンスを登録解除もしくは終了した場合。1 時間以内はコンテナインスタンスの情報を取得可能

コンテナインスタンスを停止した場合 ACTIVE のままだが、エージェント接続ステータスは FALSE となる


[AMI ストレージ設定](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-ami-storage-config.html)

* Amazon Linux 2023
  * 1 つの 30 GiB のルートボリュームが付属
  * Amazon ECS に最適化された Amazon Linux 2 AMI のデフォルトファイルシステムは xfs を使用しており、Docker は overlay2 ストレージドライバーを使用
* Amazon Linux 2
  * 1 つの 30 GiB のルートボリュームが付属
  * Amazon ECS に最適化された Amazon Linux 2 AMI のデフォルトファイルシステムは xfs を使用しており、Docker は overlay2 ストレージドライバーを使用
* Amazon Linux AMI
  * オペレーティングシステム用に 8 GiB ボリュームが /dev/xvda にアタッチ。ルートとしてマウント。
  * Docker によるイメージとメタデータの保存用に 22 GiB のボリュームが /dev/xvdcz に追加でアタッチ。
  * devicemapper を使用

補足: [Use the OverlayFS storage driver](https://docs.docker.com/storage/storagedriver/overlayfs-driver/)

Docker 側では `/etc/docker/daemon.json` を以下のように設定。
```json
{
  "storage-driver": "overlay2"
}
```


[Amazon ECS 最適化 Linux AMI のビルドスクリプト](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-ami-build-scripts.html)

* ビルド方法は OSS 化されている。https://github.com/aws/amazon-ecs-ami


[Amazon ECS 最適化 Bottlerocket AMI](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-bottlerocket.html)

* コンテナの実行に必要な最小数のパッケージのみが含まれている
* パッケージマネージャが含まれていない


#### コンテナインスタンスの設定

[Launching an Amazon ECS Container Instance](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_container_instance.html)

コンテナエージェントの設定ファイルにクラスターの設定が必要。
```bash
#!/bin/bash
echo ECS_CLUSTER=your_cluster_name >> /etc/ecs/ecs.config
```


[Using Spot Instances](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/container-instance-spot.html)

次の設定を行うことで、スポットの中断の通知を受けた場合に(2 分前に発報される) ECS インスタンスを DRAINING 状態にすることができる。
```bash
ECS_ENABLE_SPOT_INSTANCE_DRAINING=true
```


[Elastic network interface trunking](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/container-instance-eni.html)

* c5.large は通常プライマリネットワークインタフェースを含めて 3 つまでの ENI をアタッチ可能。
* アカウント設定で **awsvpcTrunking** にオプトインすることで、ENI のリミットを増やすことができる。c5.large の場合 12 個だが、プライマリネットワークインタフェースとトランクネットワークインタフェースで 1 個ずつ使うので、タスクで使用可能となるのは 10 個となる。
* Windows コンテナは未サポート。
* リソースベースの IPv4 DNS リクエストがオフになっている必要がある。
* 共有サブネットは未サポート。
* 一部インスタンスタイプは実サポート。


[Container Instance Memory Management](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/memory-management.html)

コンテナに割り当て可能なメモリは Docker の ReadMemInfo() により取得する。また、コンテナエージェント側で ECS_RESERVED_MEMORY に MiB を設定することで、指定量分をタスクの割当対象から除外できる。減じた量が、そのインスタンスに配置できるメモリ量となる。


[Windows インスタンス](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-windows.html)


[Amazon ECS Windows コンテナインスタンスの起動](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/launch_window-container_instance.html)

ユーザーデータの例
```powershell
<powershell>
Import-Module ECSTools
Initialize-ECSAgent -Cluster your_cluster_name -EnableTaskIAMRole -EnableTaskENI -AwsvpcBlockIMDS -AwsvpcAdditionalLocalRoutes '["ip-address"]'
</powershell>
```


[外部インスタンス(Amazon ECS Anywhere)](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-anywhere.html)

* EXTERNAL 起動タイプ
* ELB、サービスディスカバリ未サポート
* awsvpc 未サポート
* インスタンスには `ecs.capability.external` 属性が設定される
* `ecs-a-*`、`ecs-t-*` などのエンドポイントとの疎通性が必要


[モニタリング](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/using_cloudwatch_logs.html)

* コンテナインスタンスのログを CloudWatch Logs に送信可能。
* CLoudWatch Agent を導入し設定する必要がある。


[Container instance draining](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/container-instance-draining.html)

* **DRAINING 状態になったインスタンスには新規タスクは割り当てられない。**
* サービスから起動したタスクの場合
  * Pending のタスクは即座に停止させられる。
  * 利用可能なコンテナインスタンス容量がある場合は置き換えタスクが実行される。
  * 最小ヘルス率が 100 % 未満の場合は希望数を無視してタスクを最小ヘルス率の割合まで停止する。100 % の場合はタスクの停止は発生しない。
  * 最大率が 100 % よりも大きい場合は draining する前にタスクを起動する。100 % の場合は、draining タスクの停止までは新規タスクを起動できない。
* スタンドアローンのタスクの場合
  * タスクは停止しないので手動停止が必要



## Amazon ECS Container Agent

* ECS Agent。ソースコードは GitHub 上にある。[aws/amazon-ecs-agent](https://github.com/aws/amazon-ecs-agent)
* Fargate PV 1.4.0 は Fargate Agent が使われる。


[Installing the Amazon ECS Container Agent](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-install.html)

Amazon Linux 2、Amazon Linux の場合はパッケージとして提供されている。
```
sudo amazon-linux-extras install -y ecs; sudo systemctl enable --now ecs
```

インストール後はメタデータにアクセスできるか確認する。
```
curl -s http://localhost:51678/v1/metadata | python -mjson.tool
```

* ecs-init はホストネットワークモードでコンテナエージェントのコンテナを起動する


[Updating the Amazon ECS Container Agent on an Amazon ECS-optimized AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/agent-update-ecs-ami.html)

アップデート方法はいくつかある。

* **インスタンスを Terminate し、最新の  Amazon ECS-optimized Amazon Linux 2 AMI を使用。**
* ecs-init パッケージを最新にする。
* UpdateContainerAgent API を使用する。


[Amazon ECS Container Agent Configuration](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-config.html)

* パラメータは `/etc/ecs/ecs.config` ファイルで設定可能。
* パラメータのリファレンスは [GitHub リポジトリ](https://github.com/aws/amazon-ecs-agent/blob/master/README.md) を参照のこと
* パラメータ例
  * ECS_DATADIR: コンテナの状態を保存するディレクトリパス
  * ECS_ENABLE_TASK_IAM_ROLE: タスクの IAM ロールを有効化するかどうか
  * ECS_DISABLE_IMAGE_CLEANUP: 自動イメージクリーンアップを行うかどうか
  * ECS_AWSVPC_BLOCK_IMDS: awsvpc ネットワークモードを使用して起動されるタスクのインスタンスメタデータへのアクセスをブロックするかどうか
  * ECS_AWSVPC_ADDITIONAL_LOCAL_ROUTES: awsvpc ネットワークモードでは、これらのプレフィックスへのトラフィックは、タスク Elastic Network Interface ではなく、ホストブリッジ経由でルーティングされる
  * ECS_TASK_METADATA_RPS_LIMIT: タスクメタデータエンドポイントのスロットリングに使用する値
  * ECS_ENABLE_SPOT_INSTANCE_DRAINING: スポットインスタンスのドレイニングを有効化するかどうか


[Private Registry Authentication for Container Instances](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/private-auth-container-instances.html)

* プライベートレジストリの認証を設定可能。
* コンテナエージェント側での設定が必要のため、Fargate では対応不可


[Automated Task and Image Cleanup](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/automated_image_cleanup.html)

* 停止したタスクのイメージ、ログやデータボリュームなどは所定の期間が経過したあとに削除される
* 動作はパラメータで調整可能。


[Amazon ECS Container Metadata File](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/container-metadata.html)

* コンテナメタデータを有効化。(ECS_ENABLE_CONTAINER_METADATA=true)することで、所定のファイルパスから参照できるようになる。
* Linux の場合は `/var/lib/ecs/data/metadata/cluster_name/task_id/container_name/ecs-container-metadata.json`


[Amazon ECS タスク メタデータ エンドポイント](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-metadata-endpoint.html)

* 例えばタスクメタデータバージョン 4 の場合は以下のようなエンドポイントパスが用意されている
  * `${ECS_CONTAINER_METADATA_URI_V4}`
  * `${ECS_CONTAINER_METADATA_URI_V4}/task`
  * `${ECS_CONTAINER_METADATA_URI_V4}/taskWithTags`
  * `${ECS_CONTAINER_METADATA_URI_V4}/stats`
  * `${ECS_CONTAINER_METADATA_URI_V4}/task/stats`


[タスクスケールインプロテクションのエンドポイント](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-scale-in-protection-endpoint.html)

* URI `$ECS_AGENT_URI/task-protection/v1/state` への PUT によりスケールインプロテクションを設定可能
* 有効化する際に保護する期間も設定できる


[Amazon ECS Container Agent Introspection](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-introspection.html)

* コンテナインスタンスのメタデータは port 51678 に対して HTTP リクエストを行うことにより取得可能。様々なクエリをかけられる
```
curl -s http://localhost:51678/v1/metadata | python -mjson.tool
```
* Docker 統計は ```${ECS_CONTAINER_METADATA_URI_V4}/stats```。[Docker Stats](https://docs.docker.com/engine/api/v1.30/#operation/ContainerStats) 参照のこと


[HTTP Proxy Configuration](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/http_proxy_config.html)

* HTTP プロキシを設定することが可能
* 以下ファイルにプロキシの指定を行う
  * `/etc/ecs/ecs.config`
  * `/etc/systemd/system/ecs.service.d/http-proxy.conf`
  * `/etc/systemd/system/docker.service.d/http-proxy.conf`



## Scheduling tasks

[Amazon ECS タスクのスケジューリング](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/scheduling_tasks.html)

* サービススケジューラ
  * サービススケジューラ戦略
    * レプリカ: デフォルトでは AZ 間で分散される
    * デーモン: タスク配置制約を満たすインスタンス上に配置される
* 手動でタスク実行: RunTask の API によりタスクを起動する
* cron ライクなスケジューラ: EventBridge Scheduler を使用してスケジュールを作成できる
* カスタムスケジューラ: Blox のような OSS がある。StartTask ではコンテナインスタンスを指定してタスクを起動する


#### タスクの配置

[Amazon ECS コンソールを使用したスタンドアロンタスクの実行](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs_run_task-v2.html)

* タスク配置戦略は以下から選ぶことができる。
  * AZ Balanced Spread
  * AZ Balanced BinPack
  * BinPack
  * One Task Per Instance
  * Custom


[Amazon ECS タスクの配置](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-placement.html)

* サービスから起動されている場合はデフォルトのタスク配置戦略は `attribute:ecs.availability-zone` を使用した `spread`
* Fargate はタスク配置戦略、タスク配置制約をサポートしない。AZ にまたがって分散される
* タスク配置戦略
  * タスク開始、終了時にインスタンスを選択するためのアルゴリズム
  * ベストエフォートなので、最適なオプションが利用できない場合でもタスク配置を試みる
* タスク配置制約
  * タスク配置中に考慮されるルール

以下順序でインスタンスを選択

1. タスク定義の CPU、GPU、メモリ、ポートの要件を満たすインスタンスを識別
2. タスク配置の制約事項を満たすインスタンスを識別
3. タスク配置戦略を満たすインスタンスを識別
4. タスクを配置するインスタンスを選択


[タスクグループ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-groups.html)

* タスク配置戦略、タスク配置制約でタスクグループを元にした配置ができる
* デフォルトでは、スタンドアロンタスクはタスク定義ファミリ名、サービスの一部として起動されたタスクはサービス名をタスクグループ名として使用


[タスク配置戦略](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-placement-strategies.html)

以下のタスク配置戦略がある。
* binpack: 一台に詰め込むような動作。field には memory などを指定可能。スケールイン時は利用可能リソースが一番多いインスタンスのタスクを停止する
* random: ランダムに配置
* spread: InstanceId もしくは attribute:ecs.availability-zone を設定する。スケールイン時は AZ 間のバランスを保つようにタスクを停止する


[Amazon ECS タスク配置の制約事項](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-placement-constraints.html)

* タスク配置時に考慮されるルール。条件を満たさない場合はタスクは PENDING 状態となる
* 制約タイプ
  * distinctInstance
  * memberOf: 式の構文はクラスタークエリ言語で行う
* 属性
  * コンテナインスタンスに設定。名前と値のペアになっている
  * 組み込み属性。自動的に設定されるもの
    * ecs.ami-id
    * ecs.availability-zone
    * ecs.instance-type
    * ecs.os-type: Linux or Windows
    * ecs.os-family: LINUX or WINDOWS_SERVER_<OS_Release>_<FULL or CORE>
    * ecs.cpu-architecture: x86_64 or arm64
    * ecs.vpc-id
    * ecs.subnet-id
  * オプションの属性
    * ecs.awsvpc-trunk-id
    * ecs.outpost-arn
    * ecs.capability.external
  * カスタム属性
    * AWS CLI だと `aws ecs put-attributes` によって設定できる
    * ECS Agent だと `ECS_INSTANCE_ATTRIBUTES` で設定


#### タスクのスケジューリング

[タスクのスケジューリング (cron)](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/scheduled_tasks.html)

EventBridge スケジューラのスケジュールによってタスクを起動することが可能。
以下のパラメータを設定可能。
* タスク定義
* タスク数
* キャパシティープロバイダー戦略もしくは起動タイプ
* サブネットなどのネットワーク設定
* タスク配置戦略、タスク配置制約
* タグ
* リトライポリシーとデッドレターキュー


#### タスクのライフサイクル

[タスクのライフサイクル](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-lifecycle.html)

* 複数のターゲットグループを使用している場合は Activationg, Deactivating を経由する


#### リタイア、リサイクル

[タスクのリタイア](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-retirement.html)

[Fargate タスクリサイクル](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-recycle.html)



## サービス

[Amazon ECS サービス](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs_services.html)

次の機能がある。
* タスク数の維持
* ELB の背後に配置
* タスク起動に連続して失敗した場合はサービス起動の調整ロジックが働く

サービススケジューラ戦略は次の２つ。
* レプリカ: タスク数を維持
* デーモン: コンテナインスタンスごとに一つのタスク

**デーモン**
* タスク配置制約を満たすコンテナインスタンス上にタスクを配置する
* 複数のタスクが同一コンテナインスタンス上で稼働する場合、まずデーモンのリソースから確保される


[サービス定義パラメータ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/service_definition_parameters.html)

**TODO**

設定可能項目。

* 起動タイプ
* キャパシティープロバイダー戦略
* タスク定義
* プラットフォームのバージョン
* クラスター
* サービス名
* スケジュール戦略（レプリカ or デーモン）
* 必要数
* デプロイ設定（最大率、最小ヘルス率）
* デプロイメントコントローラー（ECS(ローリングアップデート), CODE_DEPLOY, EXTERNAL）
* タスクの配置（配置成約、配置戦略）
* タグ
* ネットワーク構成（サブネット、セキュリティグループ、パブリックIP付与）
* ヘルスチェック猶予時間
* ロードバランサ（ターゲットグループ ARN、ロードバランサ名、コンテナ名、コンテナポート）
* IAM ロール(ELB ありの構成で awsvpc を使用していない場合に指定。サービスにリンクされたロールがある場合は、そちらが使われる)
* サービス検出


[サービスの作成](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/create-service.html)

作成方法のチュートリアルのページ。
いくつかの項目の仕様についても書かれている。

タスク配置戦略。

* [AZ Balanced Spread (AZ バランススプレッド)] - アベイラビリティーゾーン間およびアベイラビリティーゾーン内のコンテナインスタンス間でタスクを分散します
* [AZ Balanced BinPack (AZ バランスビンパック)] - 利用可能な最小メモリでアベイラビリティーゾーン間およびコンテナインスタンス間でタスクを分散します
* [BinPack (ビンパック)] - CPU またはメモリの最小利用可能量に基づいてタスクを配置します
* [One Task Per Host (ホストごとに 1 つのタスク)] - 各コンテナインスタンスのサービスから最大 1 タスクを配置します
* [カスタム] - 独自のタスク配置戦略を定義します。設定ドキュメントの例については、「Amazon ECS タスクの配置」を参照してください

ネットワークモード。
* **EC2 起動タイプの場合、awsvpc ネットワークモードはパブリック IP アドレスを使用する ENI を提供しない**。よって、NAT ゲートウェイなどを用意する必要あり。

ELB
* 動的なポートマッピングにより、単一のコンテナインスタンス上で複数のポートを使用可能

サービスの auto scaling
* ターゲット追跡ポリシーでは以下を設定可能。
  * ECSServiceAverageCPUUtilization—サービスの CPU 平均使用率
  * ECSServiceAverageMemoryUtilization—サービスのメモリ平均使用率
  * ALBRequestCountPerTarget—Application Load Balancer ターゲットグループ内のターゲットごとに完了したリクエストの数


[サービスの更新](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/update-service.html)

タスク定義の変更やプラットフォームバージョンなどを変更できる。サービスの更新時の旧タスクの停止、新タスクの起動の動作は、デプロイタイプと最大率、最小ヘルス率によって変わる。

タスクの停止時は**最小ヘルス率**の割合を下回らないように制御される。
また、タスクの起動時は**最大率**で設定した割合まで起動可能。最大率が 100 の場合はもうタスクを起動できないので、まず旧タスクを停止してから新タスクを起動する動作となる。

**タスクの置き換えの際は、ELB からタスクを登録解除し (使用されている場合は) draining が完了するのを待つ。その後、タスクが実行されているコンテナに docker stop と同等のコマンドを発行する。つまり SIGTERM を送る。30 秒経過しても停止していなかった場合は、SIGKILL を送る。**

**新しいデプロイの強制**を行うことでタスク定義等の変更を行うことなく、タスクの入れ替えを発生させることができる。


[デプロイタイプ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/deployment-types.html)

次の 3 つがある。

* ローリング更新
* CodeDeploy を使用した Blue/Green デプロイ
* 外部デプロイ


[ローリング更新](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/deployment-type-ecs.html)

サーキットブレーカーを設定できる。デプロイが失敗した場合にロールバックできる。
* CLB を使用している場合はサポートされない
* 少なくとも 1 つのタスクが起動成功すると発動しない
* DescribeServices で状態を確認できる。rolloutState、rolloutStateReason が該当。ロールアウトの状態は IN_PROGRESS 状態から始まり、成功すると COMPLETED に状態移行、定常状態にならない場合は FAILED 状態に移行。FAILED 状態のデプロイでは、新しいタスクは起動されない
* RUNNING に到達しなかった場合に故障数のカウントを 1 増やす。閾値に達した場合に FAILED に移行
* RUNNING に達した場合はヘルスチェックを行い、失敗した場合は故障数のカウントを 1 増やす
* 閾値は (タスク数 / 2) で計算されるが 10 〜 200 の間に収まらない場合は 10, 200 にセットされる

CloudWatch アラーム
* 指定した CloudWatch アラームが ALARM 状態になった場合にロールバックする。以下のようにサービスを作成する
```shell
aws ecs create-service \
     ...
     --deployment-configuration "alarms={alarmNames=[alarm1Name,alarm2Name],enable=true,rollback=true}" \
```


[CodeDeploy を使用した Blue/Green デプロイ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/deployment-type-bluegreen.html)

CodeDeply の Blue/Green Deployment の考慮事項
* デプロイ時に Green のタスクセットを作成する。テストトラフィックを Green のタスクセットに ModifyLister したあと、本番用トラフィックを Blue のタスクセットから Green のタスクセットに ModifyListener する
* トラフィックの移行は一括、線形、Canaly から選択可能。ただし、NLB では `CodeDeployDefault.ECSAllAtOnce` のみ設定可能
* CLB はサポートされていない
* サービスの Auto Scaling と併用でき、デプロイ中もスケーリングできる。しかし、デプロイが失敗する場合がある


[外部デプロイ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/deployment-type-external.html)


[サービスの負荷分散](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/service-load-balancing.html)

付加機能
* 一つのサービスを複数のロードバランサのターゲットグループに登録可能
* 動的ポートマッピングが可能

考慮事項
* Fargate の場合は、ターゲットタイプとして ip を指定する必要がある。
* **タスクがヘルスチェックの条件を満たさない場合は、タスクは停止され、再度起動される。**
* NLB と Fargate の組み合わせの場合、送信元 IP アドレスは NLB のプライベートアドレスとなる。よって、タスク側で NLB のプライベートアドレスを許可するしかないが、その場合は世界中からのアクセス可能な状態になる（NLB 側でセキュリティグループを設定できずフィルタリングできないため）。
* 登録解除の遅延よりもタスク定義の stopTimeout を長くすると良い


[サービスの Auto Scaling](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/service-auto-scaling.html)

* 考慮事項
  * デプロイ中はスケールインしないようになっている。スケールアウトはされる
  * スケールアウトを中断する場合は ```register-scalable-target``` を使用する。デプロイ後に ```register-scalable-target``` を実行し再開されるのを忘れないように
  * クールダウン期間をサポート
  * 

次の 3 つがサポートされている。
* ターゲット追跡スケーリングポリシー
* ステップスケーリングポリシー
* スケジュールに基づくスケーリング


[Service Connect](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/service-connect.html)

* 特徴
  * サービスディスカバリ、サービスメッシュの両方が構築される
  * VPC DNS に依存しない名前空間内のサービスを参照する
  * Cloud Map に Service Connect エンドポイントが作成される
  * envoy コンテナがインジェクトされる
  * envoy コンテナにより CloudWatch メトリクスが生成される
  * アプリケーションは Service Connect のエンドポイントへの接続のみプロキシを使用
  * プロキシはラウンドロビン負荷分散、外れ値検出、再試行を実行する
* 設定手順
  * タスク定義では portMappings の設定が必要
  * ECS クラスターに名前空間を設定しておくのが簡便な方法
  * サービス定義にて Service Connect の設定を行う
* 考慮事項
  * 以下は未サポート
    * Windows
    * ECS Anywhere
    * Blue/Green デプロイ
    * スタンドアローンのタスク
  * [Service Connect Agent](https://github.com/aws/amazon-ecs-service-connect-agent) が必要
  * デプロイ後に追加された新しいエンドポイントは解決できない。再デプロイが必要


[概念](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/service-connect-concepts.html)

* 接続に関して
  * 同じ名前空間内の接続に適している
  * 以下の場合は接続できるものの Service Connect エンドポイント名を解決できない
    * 他の名前空間の ECS タスク
    * Service Connect が使用されていない ECS タスク
    * ECS 外部のアプリケーション
* 用語
  * ポート名: ポートに名前をつけて識別する。タスク定義側では `containerDefinitions.portMappings.name`、サービス定義側では `serviceConnectConfiguration.services.portName`
  * クライアントエイリアス: サービス定義側。Service Connect エンドポイントのポート番号。更に エンドポイントの DNS 名を割り当て検出名を `serviceConnectConfiguration.services.dnsName` で上書きすることも可能。複数のクライアントエイリアスを設定可能
  * 検出名: サービスディスカバリでどの名前で検出されるかを設定。未設定時はポート名が使用される
  * エンドポイント: 対象サービスに接続するための URL。例としては `http://blog:80`
  * クライアントサービス: 名前空間が設定されている場合は、同名前空間内のエンドポイントを名前解決可能
  * クライアント/サーバーサービス: 名前空間と少なくとも 1 つのエンドポイントが設定されている必要がある。同一名前空間内の他のサービスから名前解決できるようになる
* 設定
  * クラスターにはデフォルトの名前空間を設定可能
  * クライアントサービス作成時は名前空間の選択が必要
  * クライアント/サーバーサービス作成時は名前空間に加えて Service Connect サービス設定が必要
* その他考慮点
  * デプロイ後に追加されたエンドポイントは名前解決できない。名前解決したい場合はデプロイし直す必要がある
  * Service Connect プロキシコンテナのタスクに対する CPU とメモリとして、256 CPU ユニットと少なくとも 64 MiB のメモリを追加することをお勧め
* プロキシ
  * 負荷分散戦略はラウンドロビン
  * 外れ値検知(outlier detection)。パッシブなヘルスチェック。直近 30 秒以内に 5 つ以上の接続が失敗した場合、当該 ECS タスクへは 30 〜 300 秒ルーティングしない
  * 再試行回数は 2 回。2 回目の試行では前の接続先ホストを使用しない
  * デフォルトのタイムアウト値は 15 秒


[チュートリアル: AWS CLI を使用した Fargate での Service Connect の使用](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/create-service-connect.html)

* ECS クラスター作成。下記コマンドにより Cloud Map の名前空間も作成される
  * `aws ecs create-cluster --cluster-name tutorial --service-connect-defaults namespace=service-connect`
* タスク定義作成。`portMappings` の `name`、`appProtocol` が Service Connect の固有設定
```json
{
    "family": "service-connect-nginx",
    "executionRoleArn": "arn:aws:iam::123456789012:role/ecsTaskExecutionRole",
    "networkMode": "awsvpc",
    "containerDefinitions": [
        {
        ...
        "portMappings": [
            {
                "containerPort": 80,
                "protocol": "tcp",
                "name": "nginx",
                "appProtocol": "http"
            }
        ],
        ...
}
```
* サービス作成
```json
{
    "cluster": "tutorial",
    "serviceName": "service-connect-nginx-service",
    "taskDefinition": "service-connect-nginx",
    ...
    "serviceConnectConfiguration": {
        "enabled": true,
        "services": [
            {
                "portName": "nginx",
                "clientAliases": [
                    {
                        "port": 80
                    }
                ]
            }
        ],
        "logConfiguration": {
            "logDriver": "awslogs",
            "options": {
                "awslogs-group": "/ecs/service-connect-proxy",
                "awslogs-region": "us-west-2",
                "awslogs-stream-prefix": "service-connect-proxy"
            }
        }
    }
}
```
* 名前空間内の ECS クライアントアプリケーションは、`portName` および `clientAliases` のポートを使用してこのサービスに接続。例えば `http://nginx:80/`
* 上記例にはないが、`serviceConnectConfiguration.services.clientAliases.dnsName` で指定した FQDN に対して接続することも可能。namespace とは全く関係ない内容でもよい
* 外部アプリケーションからは IP アドレス、ポート番号で接続
* logConfiguration により CloudWatch Logs にプロキシログを送信
* 当該タスクの IP アドレス、ポート番号宛に `curl` すると `server: envoy` ヘッダーがついている


[サービスディスカバリ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/service-discovery.html)

* FQDN で 名前解決できるようになる。CloudMap と連携し、サービスの検出名前空間を設定することでタスク起動時に ConfigMap にインスタンスとして追加され、Route 53 のプライベートホストゾーンに A (or SRV) レコードが設定される仕組み
* 以下のコンポーネントから構成される
  * 名前空間: FQDN のようなもの
  * サービス名: サービス名
  * インスタンス: 実リソース。ECS タスクなど
* awsvpc の場合は A レコードまたは SRV レコード。bridge, host の場合は SRV レコードのみに対応
* 全てのレコードに問題がある場合は、異常なレコードを最大 8 個返却する
* コンテナレベルのヘルスチェック結果が CloudMap のカスタムヘルスチェック API に送信される
* ECS サービスを新規作成する場合にのみ設定可能


[サービスの調整ロジック](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/service-throttle-logic.html)

サービスタスクが繰り返し起動に失敗した場合にタスクを起動する頻度を調整するロジックがある。繰り返しタスクの起動が失敗する場合、その後の再起動の試行間隔は最大 15 分まで段階的に増加する。



## タグ

[Amazon ECS リソースのタグ付け](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-using-tags.html)

* タグ付け方法
  * マネージドタグ
    * スタンドアローンのタスクでは `aws:ecs:clusterName`
    * サービスから起動したタスクでは `aws:ecs:clusterName`、`aws:ecs:serviceName`
    * 新しい ARN 形式のオプトインが必要
    * `RunTask`、`CreateService` にて `enableECSManagedTags` を `true` に設定する必要がある
  * コンソールからリソース作成した際に自動設定される
  * タグの伝搬
    * ECS タスク: タスク定義からの伝搬をサポート
    * ECS サービス: タスク定義からの伝搬、タスクへの伝搬をサポート
  * コンテナインスタンス
    * ユーザーデータにて `ECS_CONTAINER_INSTANCE_TAGS={"tag_key": "tag_value"}` を設定する
    * `ECS_CONTAINER_INSTANCE_PROPAGATE_TAGS_FROM=ec2_instance` を設定しインスタンスから伝播する方法もある
* アカウント設定でオプトインが必要 - [タグ付け認可](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-account-settings.html#tag-resources-setting)
* `ecs:TagResource` の許可が必要



## クォータ

[Amazon ECS のサービスクォータ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/service-quotas.html)



## モニタリング

[CloudWatch のメトリクス](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/cloudwatch-metrics.html)

* 名前空間: AWS/ECS
* ディメンション: ClusterName, ServiceName


[ECS のイベント](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs_cwe_events.html)

EventBridge に送信されるイベントは次の 3 種類。

* コンテナインスタンスの状態変更イベント
* タスク状態変更イベント
* サービスアクションイベント


[Container Insights](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/cloudwatch-container-insights.html)

containerInsights アカウント設定をオプトインして作成されたすべての新しいクラスターに対して有効化されている。
運用データは、パフォーマンスログイベントとして収集される。JSON スキーマのエントリとなっている。CloudWatch はこのデータから CloudWatch メトリクスを作成する。
ネットワークメトリクスは EC2 の場合はネットワークモード none, host では採取されない。

* [Container Insights のメトリクス](https://docs.aws.amazon.com/ja_jp/AmazonCloudWatch/latest/monitoring/Container-Insights-metrics-ECS.html)
* [パフォーマンスログイベントの例](https://docs.aws.amazon.com/ja_jp/AmazonCloudWatch/latest/monitoring/Container-Insights-reference-performance-logs-ECS.html)


[コンテナインスタンスのヘルス](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/container-instance-health.html)

コンテナインスタンスのヘルスステータスは ```DescribeContainerInstances``` によって取得可能。


[アプリケーショントレースデータの収集](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/trace-data.html)

OpenTelemetry 用 AWS Distro と統合して、アプリケーションからトレースデータを収集可能。


[アプリケーションメトリクスを収集する](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/metrics-data.html)

アプリケーションのメトリクスを OpenTelemetry サイドカーコンテナ用 AWS Distro を使用して CloudWatch または Amazon Managed Service for Prometheus へルーティングすることが可能。


[AWS CloudTrail を使用した Amazon ECS API コールのログ記録](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/logging-using-cloudtrail.html)

公開 API は CloufTrail に記録される。



## セキュリティ

[サービスにリンクされたロール](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/using-service-linked-roles.html)

AWSServiceRoleForECS という名前。Amazon ECS がユーザーに代わって AWS API を呼び出すために標準的に使用するロール。


[タスク実行ロール](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task_execution_IAM_role.html)

コンテナエージェント用のロール。
ECR からのイメージのプル、CloudWatch Logs へのログ送信などで使用。


[コンテナインスタンスの IAM ロール](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/instance_IAM_role.html)

ecsInstanceRole。[AmazonEC2ContainerServiceforEC2Role](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/security-iam-awsmanpol.html#security-iam-awsmanpol-AmazonEC2ContainerServiceforEC2Role) の管理ポリシーをアタッチして使用。


[タスクロール](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-iam-roles.html)

EC２ インスタンスのインスタンスプロファイルのようなもの。タスク内で使用できる権限を設定したロール。

タスク内のコンテナからはインスタンスメタデータを通じてコンテナインスタンスの IAM ロールにアクセス可能。
bridge の場合は、以下のような iptables 設定により遮断することが可能。

```
sudo yum install -y iptables-services; sudo iptables --insert FORWARD 1 --in-interface docker+ --destination 169.254.169.254/32 --jump DROP

// 再起動後も有効化。
sudo iptables-save | sudo tee /etc/sysconfig/iptables && sudo systemctl enable --now iptables
```

awsvpc の場合は、コンテナエージェントの設定ファイルで ECS_AWSVPC_BLOCK_IMDS を true に設定する。


[Amazon ECS インターフェイス VPC エンドポイント (AWS PrivateLink)](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/vpc-endpoints.html)

必要となるエンドポイントについて書かれている。

ECS については以下が必要だが、既にコンテナインスタンスが存在する場合は上から順に作成する必要がある。

* com.amazonaws.region.ecs-agent
* com.amazonaws.region.ecs-telemetry
* com.amazonaws.region.ecs



## チュートリアル

[チュートリアル: Secrets Manager のシークレットを使用して機密データを指定する](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/specifying-sensitive-data-tutorial.html)

* タスク定義では containerDefinitions.secrets にて valueFrom により指定
* タスク実行ロールに Secrets Manager の権限が必要


[チュートリアル:サービスディスカバリを使用して、サービスの作成](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/create-service-discovery.html)

* サービスディスカバリリソースの作成
  * [名前空間] - [サービス] の構造。ECS 側でサービスディスカバリを有効化することで、タスク起動時にサービス内にサービスディスカバリインスタンスとして登録される。
  * Route 53 にホストゾーンが作成され、サービスディスカバリインスタンスに対応した A レコードが登録される仕組み。 


[チュートリアル: Blue/Green デプロイを使用するサービスの作成](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/create-blue-green.html)

* ALB もしくは NLB に対応。
* サービス作成時に deploymentController を CODE_DEPLOY にする。
* appspec.yaml は以下のような内容。
```yaml
version: 0.0
Resources:
  - TargetService:
      Type: AWS::ECS::Service
      Properties:
        TaskDefinition: "arn:aws:ecs:region:aws_account_id:task-definition/first-run-task-definition:7"
        LoadBalancerInfo:
          ContainerName: "sample-app"
          ContainerPort: 80
        PlatformVersion: "LATEST"
```


[チュートリアル:Amazon ECS CloudWatch Events イベントについて](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs_cwet.html)

EventBridge のルールでターゲットを Lambda 関数にするチュートリアル。
Lambda 関数のサンプルは以下の内容。トリガーされたイベントの JSON を出力する内容となっている。
```python
import json

def lambda_handler(event, context):
    if event["source"] != "aws.ecs":
       raise ValueError("Function only supports input from events with a source type of: aws.ecs")
       
    print('Here is the event:')
    print(json.dumps(event))
```


[チュートリアル:タスク停止イベントに Amazon シンプル 通知サービス アラートを送信](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs_cwet2.html)

タスク状態変更イベントのうち、"Essential container in task exited" で停止したイベントを扱うルールを作成する。
以下のようなカスタムイベントパターンとなる。

```json
{
   "source":[
      "aws.ecs"
   ],
   "detail-type":[
      "ECS Task State Change"
   ],
   "detail":{
      "lastStatus":[
         "STOPPED"
      ],
      "stoppedReason":[
         "Essential container in task exited"
      ]
   }
}
```


[チュートリアル: Amazon ECS で FSx for Windows File Server ファイルシステムを使用](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/tutorial-wfsx-volumes.html)

* Windows Active Directory (AD) を作成する手順となっている。
* FSx for Windows File Server を作成する際に AD を指定している。
* タスク定義にてマウントの設定が可能。
```json
{
  "containerDefinitions": [
      {
          ...
          "mountPoints": [
              {
                  "sourceVolume": "fsx-windows-dir",
                  "containerPath": "C:\\fsx-windows-dir",
                  "readOnly": false
              }
          ]

  "volumes": [
    {
        "name": "fsx-windows-vol",
        "fsxWindowsFileServerVolumeConfiguration": {
            "fileSystemId": "fs-0eeb5730b2EXAMPLE",
            "authorizationConfig": {
                "domain": "example.com",
                "credentialsParameter": "arn:arn-1234"
            },
            "rootDirectory": "share"
        }
    }
```



## トラブルシューティング

[デバッグ用にAmazon ECS Exec を使用](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-exec.html)

* readonlyRootFilesystem が有効化されていると使用できない。
* ゾンビプロセスをクリーンアップするにはタスク定義に ```initProcessEnabled``` フラグ設定を推奨。
* タスクロールに SSM の権限が必要（タスク実行ロールではない）。
* ログ記録を有効化できる。ECS クラスターで設定する。


[停止されたタスクでのエラーの確認](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/stopped-task-errors.html)

マネジメントコンソールからタスクを表示することでエラーを確認できる。ただし 1 時間以内に停止したタスクしか分からないので、それ以前を確認するにはタスク状態変更イベントを CloudWatch Logs に送信するような設定を事前にしておく必要がある。

* [停止理由] を確認することでタスクの停止理由が分かる。例えば以下のような停止理由がある。
  * Task failed ELB health checks in (elb elb-name)
  * Scaling activity initiated by (deployment deployment-id)
  * Host EC2 (instance id) stopped/terminated
  * Container instance deregistration forced by user
  * Essential container in task exited
* コンテナの欄を確認することで停止理由が分かる。


[CannotPullContainer task errors](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task_cannot_pull_image.html)

コンテナイメージのプルに失敗したエラー。以下のようなエラー原因が考えられる。
* コンテナレジストリとの疎通性がない。
* イメージが格納されていない。
* ディスク容量不足。
* コンテナレジストリ側のレートリミット。


[サービスイベントメッセージ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/service-event-messages.html)

サービスイベントログには最新 100 件のイベントが表示される。

定常状態になった場合は、次のイベントが記録される。
```
service service-name) has reached a steady state. 
```
エラーになった場合もエラーに対応したメッセージが記録されるため、診断の有用な情報となる。


[指定された CPU またはメモリの値が無効](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-cpu-memory-error.html)

以下のような値が表示される場合は、CPU、メモリの設定ミス。
```
An error occurred (ClientException) when calling the RegisterTaskDefinition operation: Invalid 'cpu' setting for task. For more information, see the Troubleshooting section of the Amazon ECS Developer Guide.
```

EC2, Fargate それぞれの場合について設定可能な値がまとめられている。


[Docker デバッグ出力の有効化](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/docker-debug-mode.html)

```/etc/sysconfig/docker``` にて ```OPTIONS``` に ```-D``` フラグを追加する。Docker デーモン、ECS Agent をリスタートし反映する。


[Amazon ECS ログファイルの場所](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/logs.html)

ログファイルの場所がまとめられている。


[ログコレクター](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-logs-collector.html)

インスタンス上のログを収集できるツール。サポートへ送付する際などに使用。


[Docker 診断](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/docker-diags.html)

診断用の Docker コマンド。
```
$ docker ps -a
$ docker logs コンテナID
$ docker inspect コンテナID
```


[AWS Fargate スロットリングのクォータ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/throttling.html)



# 参考

* Document
  * [Amazon Elastic Container Service とは?](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/Welcome.html)
* サービス紹介ページ
  * [Amazon Elastic Container Service](https://aws.amazon.com/jp/ecs/)
  * [よくある質問](https://aws.amazon.com/jp/ecs/faqs/)
* [ナレッジセンター](https://aws.amazon.com/jp/premiumsupport/knowledge-center/#Amazon_Elastic_Container_Service_.28Amazon_ECS.29)
* Black Belt
  * [20200422 AWS Black Belt Online Seminar Amazon Elastic Container Service (Amazon ECS)](https://pages.awscloud.com/rs/112-TZM-766/images/20200422_BlackBelt_Amazon_ECS_Share.pdf)
  * [20190731 Black Belt Online Seminar Amazon ECS Deep Dive](https://pages.awscloud.com/rs/112-TZM-766/images/20190731_AWS-BlackBelt_AmazonECS_DeepDive_Rev.pdf)
  * [20190925 AWS Black Belt Online Seminar AWS Fargate](https://pages.awscloud.com/rs/112-TZM-766/images/20190925_AWS-BlackBelt_AWSFargate.pdf)
  * [20191127 AWS Black Belt Online Seminar Amazon CloudWatch Container Insights で始めるコンテナモニタリング入門](https://pages.awscloud.com/rs/112-TZM-766/images/20191127_AWS-BlackBelt_Container_Insights.pdf)
* [AWS コンテナサービス入門](https://pages.awscloud.com/rs/112-TZM-766/images/C3-01.pdf)




