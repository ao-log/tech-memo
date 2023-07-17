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

**TODO**


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

[Fargate で Amazon ECS タスクを実行しているときの Network Load Balancer ヘルスチェックの失敗をトラブルシューティングするにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/fargate-nlb-health-checks)


[Amazon ECS タスクのコンテナヘルスチェックの失敗をトラブルシューティングするにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-task-container-health-check-failures)

* ローカル環境でヘルスチェックに成功することを確認する
  * Dockerfile の HEALTHCHECK 設定でテスト可能
  * ※ ECS は Dockerfile 内の HEALTHCHECK 設定はモニタリングしない
* タスク定義中のコンテナヘルスチェックのコマンド指定が正しいことを確認する
* コンテナ起動に時間がかかる場合の対応
  * タスク定義の `startPeriod` にて十分な時間を指定
* コンテナログを確認


[Fargate での Amazon ECS タスクのヘルスチェックの失敗をトラブルシューティングするにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-fargate-health-check-failures)

[ELB に登録されていて、正常に機能している Amazon ECS タスクが異常とマークされて置き換えられるのはなぜですか?](https://repost.aws/ja/knowledge-center/elb-ecs-tasks-improperly-replaced)

[Fargate での Amazon ECS タスクのロードバランサーのエラーのトラブルシューティング方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-fargate-load-balancer-errors)


## タスク起動失敗のトラブルシューティング

[Amazon ECS タスクが保留状態のままになっているのはなぜですか?](https://repost.aws/ja/knowledge-center/ecs-tasks-stuck-pending-state)

[Fargate の Amazon ECS タスクが [Pending] (保留中) 状態のまま停止している場合のトラブルシューティング方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-fargate-tasks-pending-state)

[Amazon ECS クラスターのタスクが開始されないのはなぜですか?](https://repost.aws/ja/knowledge-center/ecs-run-task-cluster)

[Amazon ECS のサービスで「the closest matching container-instance container-instance-id encountered error 'AGENT'」というエラーを解決する方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-container-instance-agent-error)

[Amazon ECS で「[AWS service] was unable to place a task because no container instance met all of its requirements」(要件をすべて満たすコンテナインスタンスがないため、[AWS のサービス] はタスクを配置できませんでした) というエラーを解決するにはどうすればよいですか。](https://repost.aws/ja/knowledge-center/ecs-container-instance-requirement-error)

[Amazon ECS で「the closest matching container-instance container-instance-id has insufficient CPU units available」というエラーを解決する方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-container-instance-cpu-error)

[Fargate の Amazon ECS のネットワークインターフェイスプロビジョニングエラーを解決する方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-fargate-network-interface-errors)

[Amazon ECS でスケジュールされたタスクに関連する問題をトラブルシューティングするにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-scheduled-task-issues)

[Amazon ECS リソースを起動したときに表示される「ロールを適用し、ロードバランサーに設定されたリスナーを検証できません」という AWS CloudFormation のエラーを解決する方法を教えてください。](https://repost.aws/ja/knowledge-center/assume-role-validate-listeners)

[Amazon ECS クラスターでタスクの起動に失敗する場合の「Image is exist」エラーを解決する方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-cluster-image-not-exist-error)

[Amazon ECS for Fargate で「dockertimeouterror unable transition start timeout after wait 3m0s」というエラーを解決するにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-fargate-docker-timeout-error)

[Amazon ECS の「ResourceInitializationError: failed to validate logger args (ResourceInitializationError: ロガー引数の検証に失敗しました)」エラーを解決するにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-resource-initialization-error)


## タスク停止

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

[Amazon ECS Exec がアクティブ化されている Fargate タスクの SSM エージェントログを取得する方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-exec-ssm-logs-fargate-tasks)

* EFS をマウントすることで実現する
* タスク定義のコンテナ定義におけるコンテナパスは `/var/log/amazon` とする。また起動する ECS タスクは 1 個にする
* EC2 に EFS をマウントしてログを確認する


## その他

[ECS タスクのタグ付けに関連する問題をトラブルシューティングするにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-troubleshoot-tagging-tasks)

[Amazon ECS の API コールに関する一般的なエラーをトラブルシューティングするにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/ecs-api-common-errors)

[コンテナの終了中に Amazon ECS タスクが停止したり、開始できなかったりする場合のトラブルシューティング方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-tasks-container-exit-issues)

[Amazon ECS のブルー/グリーンデプロイに関連する問題のトラブルシューティング方法を教えてください。](https://repost.aws/ja/knowledge-center/ecs-blue-green-deployment)



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



# BlackBelt

[20200422 AWS Black Belt Online Seminar Amazon Elastic Container Service (Amazon ECS)](https://pages.awscloud.com/rs/112-TZM-766/images/20200422_BlackBelt_Amazon_ECS_Share.pdf)

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


[20190731 Black Belt Online Seminar Amazon ECS Deep Dive](https://pages.awscloud.com/rs/112-TZM-766/images/20190731_AWS-BlackBelt_AmazonECS_DeepDive_Rev.pdf)

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


[20191127 AWS Black Belt Online Seminar Amazon CloudWatch Container Insights で始めるコンテナモニタリング入門](https://pages.awscloud.com/rs/112-TZM-766/images/20191127_AWS-BlackBelt_Container_Insights.pdf)

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
  * [20200422 AWS Black Belt Online Seminar Amazon Elastic Container Service (Amazon ECS)](https://pages.awscloud.com/rs/112-TZM-766/images/20200422_BlackBelt_Amazon_ECS_Share.pdf)
  * [20190731 Black Belt Online Seminar Amazon ECS Deep Dive](https://pages.awscloud.com/rs/112-TZM-766/images/20190731_AWS-BlackBelt_AmazonECS_DeepDive_Rev.pdf)
  * [20190925 AWS Black Belt Online Seminar AWS Fargate](https://pages.awscloud.com/rs/112-TZM-766/images/20190925_AWS-BlackBelt_AWSFargate.pdf)
  * [20191127 AWS Black Belt Online Seminar Amazon CloudWatch Container Insights で始めるコンテナモニタリング入門](https://pages.awscloud.com/rs/112-TZM-766/images/20191127_AWS-BlackBelt_Container_Insights.pdf)
* [AWS コンテナサービス入門](https://pages.awscloud.com/rs/112-TZM-766/images/C3-01.pdf)




