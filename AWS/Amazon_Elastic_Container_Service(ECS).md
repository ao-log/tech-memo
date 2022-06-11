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

* サービス、タスクを管理する単位。
* EC2, Fargate の両起動タイプが共存可能
* マネジメントコンソールからクラスターを作成すると、CloudFormation のスタックが作られる。


#### Capacity Provider

[Amazon ECS capacity providers](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-capacity-providers.html)

**Capacity Provider**

クラスターは複数のキャパシティプロバイダーを持つことができる。

デフォルトのキャパシティプロバイダー戦略がクラスターに設定されており、サービスもしくはスタンドアローンのタスクにおいて、カスタムキャパシティープロバイダー戦略もしくは起動タイプが設定されていない場合に使用される。

Fargate の場合は、FARGATE、FARGATE_SPOT を使用できる。

EC2 起動タイプの場合は、次の 3 つの設定項目がある。

* Auto Scaling group
* managed scaling
* managed termination protection

**Capacity Provider Strategy**

サービス、またはタスクの設定時に、どのキャパシティプロバイダー戦略を使用するかを設定可能。

**キャパシティプロバイダー戦略**により、タスクをどのキャパシティプロバイダーに配置するかを設定できる。base, weight の２つの設定値がある。

* base: 最低何個のタスクを起動するか。
* weight: どのキャパシティプロバイダーにタスクを割り当てるかの比率を設定する。


[AWS Fargate capacity providers](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/fargate-capacity-providers.html)

* マネジメントコンソールより ECS クラスターを Networking only で作成した場合に FARGATE, FARGATE_SPOT が自動的に設定された状態になっている。
* FARGATE_SPOT のキャパシティプロバイダーのタスクがスポットの中断により停止する際は、タスク停止の 2 分前に EventBridge よりワーニングが送られる。また、SIGTERM シグナルがタスクに送られる。キャパシティに空きがある場合は新規タスクの起動を試みる。


[Auto Scaling group capacity providers](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/asg-capacity-providers.html)

* managed termination protection を使用する際、managed scaling が有効になっている必要がある。そうしないと managed termination protection は動作しない。
* **managed scaling を有効にすることで、タスク数に応じて ECS インスタンスがスケールする。** CloudWatch メトリクスの AWS/ECS/ManagedScaling を使用することで実現している。
* Managed termination protection を有効にすることで、タスクが存在する EC2 インスタンスについてスケールインから保護することができる。
* ECS クラスターを EC2 Linux + Networking で作成すると、ASG も CFn スタックの一部として自動的に作られる。これをキャパシティプロバイダーで設定する際は、スケールインからの保護を設定しておく必要あり。

クラスター作成時にキャパシティプロバイダーも設定する例。
```shell
aws ecs create-cluster \
     --cluster-name ASGCluster \
     --capacity-providers CapacityProviderA CapacityProviderB \
     --default-capacity-provider-strategy capacityProvider=CapacityProviderA,weight=1,base=1 capacityProvider=CapacityProviderB,weight=1 \
     --region us-west-2
```


[Amazon ECS クラスターの Auto Scaling](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/cluster-auto-scaling.html)

* CapacityProviderReservation のメトリクスを使用する。



## Amazon ECS Task definitions

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

以下項目を設定できる。
* Amazon リソースネーム (ARN) と ID
* AWS VPC トランキング
* CloudWatch Container Insights



## ECS container instances

[Amazon ECS container instances](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ECS_instances.html)

* ECS エージェントの導入が必須。Amazon ECS-optimized AMI の場合は既に導入済み。
* 適切な IAM ロールの設定が必要。
* Linux Amazon ECS-optimized AMI version 20200430 and later の場合、Amazon EC2 Instance Metadata Service Version 2 (IMDSv2) がサポートされている。それ以前の AMI の場合は IMDSv1。
* コンテナ側で expose するポートに対して、インバウンド通信の許可が必要。
* Amazon ECS service endpoint との接続性が必要。パブリック IP アドレスもしくはインタフェース VPC エンドポイントを用意する必要あり。

**ライフサイクル**

* ACTIVE 状態の場合 RunTask API によるリクエストを受け入れることができる。
* FALSE: コンテナインスタンスを停止し暫く経つと FALSE に遷移する。
* DRAINING: 新規タスクが配置されなくなる。


[AMI ストレージ設定](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-ami-storage-config.html)

* Amazon Linux 2
  * 1 つの 30 GiB のルートボリュームが付属
  * Amazon ECS に最適化された Amazon Linux 2 AMI のデフォルトファイルシステムは ext4 を使用しており、Docker は overlay2 ストレージドライバーを使用
* Amazon Linux AMI
  * オペレーティングシステム用に 8 GiB ボリュームが /dev/xvda にアタッチ。ルートとしてマウント。
  * Docker によるイメージとメタデータの保存用に 22 GiB のボリュームが /dev/xvdcz に追加でアタッチ。
  * devicemapper を使用。


[Amazon ECS 最適化 Linux AMI のビルドスクリプト](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-ami-build-scripts.html)

* ビルド方法は OSS 化されている。https://github.com/aws/amazon-ecs-ami


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

c5.large は通常プライマリネットワークインタフェースを含めて 3 つまでの ENI をアタッチ可能。
アカウント設定で **awsvpcTrunking** にオプトインすることで、ENI のリミットを増やすことができる。c5.large の場合 12 個だが、プライマリネットワークインタフェースとトランクネットワークインタフェースで 1 個ずつ使うので、タスクで使用可能となるのは 10 個となる。


[Container Instance Memory Management](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/memory-management.html)

コンテナに割り当て可能なメモリは Docker の ReadMemInfo() により取得する。また、コンテナエージェント側で ECS_RESERVED_MEMORY に MiB を設定することで、指定量分をタスクの割当対象から除外できる。減じた量が、そのインスタンスに配置できるメモリ量となる。


[Windows インスタンス](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-windows.html)


[外部インスタンス(Amazon ECS Anywhere)](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-anywhere.html)

* EXTERNAL起動タイプ


[Container instance draining](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/container-instance-draining.html)

* **DRAINING 状態になったインスタンスには新規タスクは割り当てられない。**また、Pending のタスクは即座に停止させられる。
* 最小ヘルス率が 100 % 未満の場合は希望数を無視してタスクを最小ヘルス率の割合まで停止する。100 % の場合はタスクの停止は発生しない。
* 最大率が 100 % よりも大きい場合は draining する前にタスクを起動する。100 % の場合は、draining タスクの停止までは新規タスクを起動できない。


[Deregister a container instance](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deregister_container_instance.html)

コンテナインスタンスは登録解除することも可能。



## Amazon ECS Container Agent

* ECS Agent。ソースコードは GitHub 上にある。[aws/amazon-ecs-agent](https://github.com/aws/amazon-ecs-agent)
* Fargate PV 1.4.0 は Fargate Agent が使われる。

[Installing the Amazon ECS Container Agent](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-install.html)

Amazon Linux 2、Amazon Linux の場合はパッケージとして提供されている。

インストール後はメタデータにアクセスできるか確認する。
```
curl -s http://localhost:51678/v1/metadata | python -mjson.tool
```


[Updating the Amazon ECS Container Agent on an Amazon ECS-optimized AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/agent-update-ecs-ami.html)

アップデート方法はいくつかある。

* **インスタンスを Terminate し、最新の  Amazon ECS-optimized Amazon Linux 2 AMI を使用。**
* ecs-init パッケージを最新にする。
* UpdateContainerAgent API を使用する。

[Amazon ECS Container Agent Configuration](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-config.html)

```/etc/ecs/ecs.config``` ファイルで設定可能。

* ECS_DATADIR: コンテナの状態を保存するディレクトリパス
* ECS_ENABLE_TASK_IAM_ROLE: タスクの IAM ロールを有効化するかどうか
* ECS_DISABLE_IMAGE_CLEANUP: 自動イメージクリーンアップを行うかどうか
* ECS_AWSVPC_BLOCK_IMDS: awsvpc ネットワークモードを使用して起動されるタスクのインスタンスメタデータへのアクセスをブロックするかどうか
* ECS_AWSVPC_ADDITIONAL_LOCAL_ROUTES: awsvpc ネットワークモードでは、これらのプレフィックスへのトラフィックは、タスク Elastic Network Interface ではなく、ホストブリッジ経由でルーティングされる
* ECS_TASK_METADATA_RPS_LIMIT: タスクメタデータエンドポイントのスロットリングに使用する値
* ECS_ENABLE_SPOT_INSTANCE_DRAINING: スポットインスタンスのドレイニングを有効化するかどうか


[Private Registry Authentication for Container Instances](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/private-auth-container-instances.html)

プライベートレジストリの認証を設定可能。


[Automated Task and Image Cleanup](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/automated_image_cleanup.html)

停止したタスクのイメージ、ログやデータボリュームなどは所定の期間が経過したあとに削除される。この動作はパラメータで調整可能。


[Amazon ECS Container Metadata File](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/container-metadata.html)

コンテナメタデータを有効化。(ECS_ENABLE_CONTAINER_METADATA=true)することで、所定のファイルパスから参照できるようになる。


[Amazon ECS Container Agent Introspection](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-introspection.html)

コンテナインスタンスのメタデータは port 51678 に対して HTTP リクエストを行うことにより取得可能。様々なクエリをかけられる。

```
curl -s http://localhost:51678/v1/metadata | python -mjson.tool
```

Docker 統計は ```${ECS_CONTAINER_METADATA_URI_V4}/stats```。[Docker Stats](https://docs.docker.com/engine/api/v1.30/#operation/ContainerStats) 参照のこと。



[HTTP Proxy Configuration](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/http_proxy_config.html)

HTTP プロキシを設定することが可能。



## Scheduling tasks

[Amazon ECS タスクのスケジューリング](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/scheduling_tasks.html)

* タスク配置にはレプリカとデーモンがある。
* RunTask の API によりタスクを起動する。
* StartTask ではコンテナインスタンスを指定してタスクを起動する。


#### タスクの配置

[Amazon ECS タスクの配置](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-placement.html)

・EC2 起動タイプの場合

以下から選ぶことができる。
* AZ Balanced Spread
* AZ Balanced BinPack
* BinPack
* One Task Per Instance
* Custom

以下の順番で稼働インスタンスを選択する。
* タスク定義で要求される CPU、メモリ、ポートの要件を満たすインスタンスを識別。
* タスク配置の制約事項を満たすインスタンスを識別。
* タスク配置戦略を満たすインスタンスを識別。
* タスクを配置するインスタンスを選択。

・Fargate 起動タイプの場合

デフォルトでは、Fargate タスクはアベイラビリティーゾーン間で分散される。


[タスクグループ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-groups.html)

* タスク配置戦略、タスク配置制約でタスクグループを元にした配置ができる。
* デフォルトでは、スタンドアロンタスクはタスク定義ファミリ名、サービスの一部として起動されたタスクはサービス名をタスクグループ名として使用。


[タスク配置戦略](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-placement-strategies.html)

以下のタスク配置戦略がある。
* binpack: 一台に詰め込むような動作。field には memory などを指定可能。スケールイン時はリソースが一番多いインスタンスのタスクを停止する
* random: ランダムに配置
* spread: InstanceId もしくは attribute:ecs.availability-zone を設定する。スケールイン時は AZ 間のバランスを保つようにタスクを停止する


[Amazon ECS タスク配置の制約事項](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-placement-constraints.html)

例えば特定のインスタンスタイプにインスタンスを配置することができる。

制約タイプは以下のものがある。
* distinctInstance
* memberOf
* ecs.os-family

コンテナインスタンスに属性を設定できる。
* ecs.ami-id
* ecs.availability-zone
* ecs.instance-type
* ecs.os-type
* ecs.cpu-architecture
* ecs.vpc-id
* ecs.subnet-id

カスタム属性も設定でき、AWS CLI だと ```aws ecs put-attributes``` によって設定できる。


#### タスクのスケジューリング

[タスクのスケジューリング (cron)](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/scheduled_tasks.html)

CloudWatch Events からタスクを起動することが可能。
以下のパラメータを設定可能。
* タスク定義
* タスク数
* キャパシティープロバイダー戦略もしくは起動タイプ
* サブネットなどのネットワーク設定
* タスク配置戦略、タスク配置制約
* タグ
* ECS exec の有効化
* リトライポリシーとデッドレターキュー


#### タスクのライフサイクル

[タスクのライフサイクル](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-lifecycle.html)

複数のターゲットグループを使用している場合は Activationg, Deactivating を経由する。


#### リタイア、リサイクル

[タスクのリタイア](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-retirement.html)

[Fargate タスクリサイクル](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-recycle.html)



## サービス

[Amazon ECS サービス](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs_services.html)

次の機能がある。

* タスク数の維持
* ELB の背後に配置
* タスク起動に失敗した場合は、起動の試行の増分的な減速を開始

サービススケジューラ戦略は次の２つ。

* レプリカ: タスク数を維持
* デーモン: コンテナインスタンスごとに一つのタスク

**デーモン**
* タスク配置制約を満たすコンテナインスタンス上にタスクを配置する。
* 複数のタスクが同一コンテナインスタンス上で稼働する場合、まずデーモンのリソースから確保される。


[サービス定義パラメータ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/service_definition_parameters.html)

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

* [AZ Balanced Spread (AZ バランススプレッド)] - アベイラビリティーゾーン間およびアベイラビリティーゾーン内のコンテナインスタンス間でタスクを分散します。
* [AZ Balanced BinPack (AZ バランスビンパック)] - 利用可能な最小メモリでアベイラビリティーゾーン間およびコンテナインスタンス間でタスクを分散します。
* [BinPack (ビンパック)] - CPU またはメモリの最小利用可能量に基づいてタスクを配置します。
* [One Task Per Host (ホストごとに 1 つのタスク)] - 各コンテナインスタンスのサービスから最大 1 タスクを配置します。
* [カスタム] - 独自のタスク配置戦略を定義します。設定ドキュメントの例については、「Amazon ECS タスクの配置」を参照してください。

ネットワークモード。

* **EC2 起動タイプの場合、awsvpc ネットワークモードはパブリック IP アドレスを使用する ENI を提供しない**。よって、NAT ゲートウェイなどを用意する必要あり。

ELB

* 動的なポートマッピングにより、単一のコンテナインスタンス上で複数のポートを使用可能。

サービスの auto scaling

ターゲット追跡ポリシーでは以下を設定可能。

* ECSServiceAverageCPUUtilization—サービスの CPU 平均使用率。
* ECSServiceAverageMemoryUtilization—サービスのメモリ平均使用率。
* ALBRequestCountPerTarget—Application Load Balancer ターゲットグループ内のターゲットごとに完了したリクエストの数。


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
* DescribeServices で状態を確認できる。rolloutState、rolloutStateReason が該当。ロールアウトの状態は IN_PROGRESS 状態から始まり、成功すると COMPLETED に状態移行、定常状態にならない場合は FAILED 状態に移行。FAILED 状態のデプロイでは、新しいタスクは起動されない。
* RUNNING に到達しなかった場合に故障数のカウントを 1 増やす。閾値に達した場合に FAILED に移行。
* RUNNING に達した場合はヘルスチェックを行い、失敗した場合は小少数のカウントを 1 増やす。
* 閾値は (タスク数 / 2) で計算されるが 10 〜 200 の間に収まらない場合は 10, 200 にセットされる。


[CodeDeploy を使用した Blue/Green デプロイ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/deployment-type-bluegreen.html)

CodeDeply の Blue/Green Deployment の考慮事項

* デプロイ時に Green のタスクセットを作成する。テストトラフィックを Green のタスクセットに ModifyLister したあと、本番用トラフィックを Blue のタスクセットから Green のタスクセットに ModifyListener する。
* トラフィックの移行は一括、線形、Canaly から選択可能。
* CLB はサポートされていない。
* サービスの Auto Scaling と併用できるが、デプロイが失敗する場合がある。


[サービスの負荷分散](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/service-load-balancing.html)

付加機能

* 一つのサービスを複数のロードバランサのターゲットグループに登録可能。
* 動的ポートマッピングが可能。

考慮事項

* Fargate の場合は、ターゲットタイプとして ip を指定する必要がある。
* **タスクがヘルスチェックの条件を満たさない場合は、タスクは停止され、再度起動される。**
* NLB と Fargate の組み合わせの場合、送信元 IP アドレスは NLB のプライベートアドレスとなる。よって、タスク側で NLB のプライベートアドレスを許可するしかないが、その場合は世界中からのアクセス可能な状態になる（NLB 側でセキュリティグループを設定できずフィルタリングできないため）。
* 登録解除の遅延よりもタスク定義の stopTimeout を長くすると良い


[サービスの Auto Scaling](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/service-auto-scaling.html)

デプロイ中はスケールインしないようになっている。スケールアウトはされる。スケールアウトを中断する場合は ```register-scalable-target``` を使用する。デプロイ後に ```register-scalable-target``` を実行し再開されるのを忘れないように。

次の 3 つがサポートされている。

* ターゲット追跡スケーリングポリシー
* ステップスケーリングポリシー
* スケジュールに基づくスケーリング


[サービスディスカバリ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/service-discovery.html)

FQDN で名前解決できるようになる。CloudMap と連携し、サービスの検出名前空間を設定することで、タスク起動時に ConfigMap にインスタンスとして追加され、Route 53 のプライベートホストゾーンに A レコードが設定される仕組み。


[サービスの調整ロジック](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/service-throttle-logic.html)

サービスタスクが繰り返し起動に失敗した場合にタスクを起動する頻度を調整するロジックがある。繰り返しタスクの起動が失敗する場合、その後の再起動の試行間隔は最大 15 分まで段階的に増加する。



## タグ

[Amazon ECS リソースのタグ付け](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-using-tags.html)

Amazon ECS の新規または既存のタスク、サービス、タスク定義、およびクラスターにタグ付け可能。

[Propagate tags from (タグの伝播元)] オプションを使用して、タスク定義またはサービスからタスクにタグをコピー可能。



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



## ナレッジセンター

[ナレッジセンター](https://aws.amazon.com/jp/premiumsupport/knowledge-center/#Amazon_Elastic_Container_Service_.28Amazon_ECS.29)


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


[Amazon ECS の Amazon ECR エラー「CannotPullContainerError: API error」を解決する方法を教えてください。](https://aws.amazon.com/jp/premiumsupport/knowledge-center/ecs-pull-container-api-error-ecr/)

原因としてありえるのは以下のもの。

* Amazon ECR エンドポイントへのアクセス権がない
* Amazon ECR リポジトリポリシーで制限されている
* IAM ロールに、イメージをプルまたはプッシュするための権限が許可されていない
* イメージが見つからない
* S3 へのアクセスがエンドポイントポリシーによって拒否されている


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



## BlackBelt

[20200422 AWS Black Belt Online Seminar Amazon Elastic Container Service (Amazon ECS)](https://www.slideshare.net/AmazonWebServicesJapan/20200422-aws-black-belt-online-seminar-amazon-elastic-container-service-amazon-ecs)

* P11, 12: コンテナオーケストレータに対して API を実行することで各種操作を行う。
* P20-: EC2 起動タイプについて
* P25-: Fargate について
* P30-: タスク定義について
  * P43-: ネットワークモード
  * P48-: データボリューム
* P53-: コンテナの実行方法　タスク or サービス
  * P55: タスクの配置
  * P56: サービススケジューラ戦略
  * P57: キャパシティプロバイダー
  * P58: サービスの Auto Scaling


[20190731 Black Belt Online Seminar Amazon ECS Deep Dive](https://www.slideshare.net/AmazonWebServicesJapan/20190731-black-belt-online-seminar-amazon-ecs-deep-dive-162160987)

* P14: シークレットをコンテナ内のアプリに渡す場合の推奨の方法
  * Secrets Manager を利用し、タスク定義では environmentの valueFrom を使用して Secrets の ARN を記載
* P21: サービスディスカバリの方法
  * 要件次第である。
  * ELB を使用するのが一つの手。
  * 一方で、ECS Service Discovery は ECS が Route 53 に対して自動登録、削除をする仕組み。
* P32: サイドカーのような依存関係のあるコンテナの制御方法
  * タスク定義の dependsOn で依存関係を指定する
  * startTimeout: 依存関係の解決の再試行を止めるまでの時間 
  * stopTimeout: コンテナが SIGTERM で終了しなかった場合に SIGKILL されるまでの時間
* P41: スケジュールされたタスクのエラーハンドリング方法
  * 要件次第だが、StepFunctions により実行しエラーハンドリングする
  * 単に検知だけでよければ EventBridge を使用。
* P46: 自分たちでカスタマイズしたデプロイを行う方法
  * External Deployent Contoller を使用する。
* P64: EC2 起動タイプで awsvpc を使用している場合に起動できるタスク数が少ない
  * ENI Trunking 機能を有効化する
* P70: コンテナの起動まで時間がかかるためヘルスチェックが失敗する
  * ヘルスチェックの猶予時間でアプリケーションにあった時間を設定する
* P75: Fargate で起動するタスクのサイズを選ぶにあたってのリソース使用状況の把握方法
  * Container Insight を使用することでタスク、コンテナ単位のリソース使用状況を確認可能。


[20191127 AWS Black Belt Online Seminar Amazon CloudWatch Container Insights で始めるコンテナモニタリング入門](https://www.slideshare.net/AmazonWebServicesJapan/20191127-aws-black-belt-online-seminar-amazon-cloudwatch-container-insights)

* P19: Container Insight はタスク、コンテナレベルでのモニタリングが可能
* P46: パフォーマンスログは CloudWatch Logs へ送られる。
* P51: ユースケース 1. ECS タスクに配置するコンテナリソースのサイジング
  * コンテナごとのリソース使用状況を確認し、適切なサイズに設定する。
* P57: 特定のタスクだけで発生している問題の調査
  * アプリケーションログの表示からログを確認する。



# 参考

* Document
  * [Amazon Elastic Container Service とは?](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/Welcome.html)
* サービス紹介ページ
  * [Amazon Elastic Container Service](https://aws.amazon.com/jp/ecs/)
  * [よくある質問](https://aws.amazon.com/jp/ecs/faqs/)
* [ナレッジセンター](https://aws.amazon.com/jp/premiumsupport/knowledge-center/#Amazon_Elastic_Container_Service_.28Amazon_ECS.29)
* Black Belt
  * [20200422 AWS Black Belt Online Seminar Amazon Elastic Container Service (Amazon ECS)](https://www.slideshare.net/AmazonWebServicesJapan/20200422-aws-black-belt-online-seminar-amazon-elastic-container-service-amazon-ecs)
  * [20190731 Black Belt Online Seminar Amazon ECS Deep Dive](https://www.slideshare.net/AmazonWebServicesJapan/20190731-black-belt-online-seminar-amazon-ecs-deep-dive-162160987)
  * [20180220 AWS Black Belt Online Seminar - Amazon Container Services](https://www.slideshare.net/AmazonWebServicesJapan/20180214-aws-black-belt-online-seminar-amazon-container-services)
  * [20190925 AWS Black Belt Online Seminar AWS Fargate](https://www.slideshare.net/AmazonWebServicesJapan/20190925-aws-black-belt-online-seminar-aws-fargate)
  * [20191127 AWS Black Belt Online Seminar Amazon CloudWatch Container Insights で始めるコンテナモニタリング入門](https://www.slideshare.net/AmazonWebServicesJapan/20191127-aws-black-belt-online-seminar-amazon-cloudwatch-container-insights)
* [AWS コンテナサービス入門](https://pages.awscloud.com/rs/112-TZM-766/images/C3-01.pdf)




