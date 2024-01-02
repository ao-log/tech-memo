# ECS



## 用語

* クラスター
* サービス
* タスク
* タスク定義
* ECS エージェント
* 起動タイプ



## Amazon ECS とは

[Amazon Elastic Container Service とは](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/Welcome.html)

管理方法

* マネジメントコンソール
* AWS CLI
* AWS SDK
* AWS Copilot
* ECS CLI
* AWS CDK


### ECS コンポーネント

* クラスター
* サービス
* タスク
* タスク定義
* コンテナエージェント



## はじめに

[Amazon ECS の開始方法](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/getting-started.html)

以下のドキュメントが用意されている。

* Amazon ECS を使用するようにセットアップする
* Amazon ECS で使用するコンテナイメージの作成
* AWS Fargate の Linux コンテナによるコンソールの使用開始
* AWS Fargate の Windows コンテナによるコンソールの使用開始
* Amazon EC2 で Windows を使用してコンソールを開始する


[AWS Fargate の Linux コンテナによるコンソールの使用開始](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/getting-started-fargate.html)

次の流れ。

* クラスターを作成
* タスク定義を作成
* サービスを作成



## ディベロッパーツール

[AWSコパイロットを使用して、Amazon ECS の開始方法](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/getting-started-aws-copilot-cli.html)

サンプルアプリケーションのデプロイ。

```
git clone https://github.com/aws-samples/amazon-ecs-cli-sample-app.git demo-app && \ 
cd demo-app &&                               \
copilot init --app demo                      \
  --name api                                 \
  --type 'Load Balanced Web Service'         \
  --dockerfile './Dockerfile'                \
  --port 80                                  \
  --deploy
```


[AWS CDKを使用してAmazon ECS の開始方法](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/tutorial-ecs-web-server-cdk.html)

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

[Amazon ECS on AWS Fargate](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/AWS_Fargate.html)

* 各 Fargate タスクは、独自の分離境界を持ち、基本となるカーネル、CPU リソース、メモリリソース、または Elastic Network Interface を別のタスクと共有しない。
* `requiresCompatibilities` を `FARGATE` にする必要がある
* **ネットワークモードは awsvpc にする必要あり。タスクごとに ENI が割り当てられる。**
  * パブリックサブネットの場合は `assignPublicIp` が `ENABLED` になっている必要がある
* Windows コンテナは Fargate Spot 未対応
* 負荷分散
  * ALB, NLB が対応。NLB UDP はプラットフォームバージョン 1.4 以上で対応
  * target type を `ip` に設定する必要がある
* Seekable OCI による遅延読み込み
  * [awslabs/soci-snapshotter](https://github.com/awslabs/soci-snapshotter) によりインデックスを作成できる。プッシュする必要がある
  * タスクメタデータエンドポイントにて SOCI による遅延読み込みがされたかを確認できる。`Snapshotter` フィールドがある。また、各コンテナのパスにも `Snapshotter` があり、デフォルト値は `overlayfs` だが SOCI 使用時は `soci` になる


[Task definitions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/fargate-tasks-services.html)

* 一部のタスク定義パラメータは無効、もしくは制限されている。
  * disableNetworking
  * dnsSearchDomains
  * dnsServers
  * dockerSecurityOptions
  * extraHosts
  * gpu
  * ipcMode
  * links
  * placementConstraints
  * privileged
  * maxSwap
  * swappiness
* サポートされているが制限のあるパラメータもある
  * `linuxParameters`
    * `capabilities` は `CAP_SYS_PTRACE` のみ追加可能
    * `devices`, `sharedMemorySize`, `tmpfs` はサポートされない
  * `volumes`
    * `dockerVolumeConfiguration` はサポートされない。ホストボリュームのみサポートされる。
  * `cpu`
    * Windows コンテナでは 1 vCPU 以下にできない
* Network mode
  * `awsvpc` のみ
* Task Operating Systems
  * Amazon Linux 2
  * Amazon Linux 2023
  * Windows Server 2019 Full
  * Windows Server 2019 Core
  * Windows Server 2022 Full
  * Windows Server 2022 Core
* Task CPU architecture
  * ARM or X86_64
  * Windows は X86_64 のみ
* Task CPU and memory
  * タスクレベルでの CPU, Memory 指定が必要
  * CPU, Memory は表に記載の組み合わせのみ可能
* Task resource limits
  * Windows では未サポート
  * Fargate ではデフォルト値が設定される。nofile は上書き可能
* Logging
  * Event logging
    * タスク状態変更イベントなどを EventBridge にてトリガーできる
  * Application logging
    * `awslogs` ログドライバーなどによりログ送信できる
* Amazon ECS task execution IAM role
  * ECR からのイメージプル、CloudWatch Logs のロググループ作成などの使用される
* Task storage
  * EFS がサポートされている
  * エフェメラルストレージをバインドマウント可能


[AWS Fargate platform versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html)

* セキュリティパッチ適用時は既存タスクをリタイアし、新規タスクはパッチ適用済みのプラットフォームで稼働する


[Linux platform versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform-linux-fargate.html)

* LATEST だと 1.4.0 になる
* 新しいタスクは最新のプラットフォームのリビジョン上で動作する
* Linux と Windows はプラットフォームバージョンは別管理
* 1.4.0
  * 2020年11月5日にローンチされた
  * IPv6 を有効化したサブネットの場合、IPv4, IPv6 の両方のアドレスがアサインされる
  * メタデータエンドポイント v4 で追加のメタデータが提供されている
  * NLB の UDP トラフィックを受信可能になった
  * エフェメラルストレージは AWS 所有の AES-256 暗号化アルゴリズムで暗号化されるようになった
  * EFS サポート
  * エフェメラルストレージの最小サイズは 20 GB になった
  * タスク ENI が使用されるようになった。[Task networking for tasks hosted on Fargate](https://docs.aws.amazon.com/AmazonECS/latest/userguide/fargate-task-networking.html)。一部エンドポイントへの疎通はサービス側の ENI が使用されていたのが歴史的な経緯
  * ジャンボフレームのサポート
  * CloudWatch, Container Insights がネットワークのメトリクスを含むようになった
  * Linux Parameter `SYS_PTRACE` のサポート
  * Fargate Container Agent が使用されるようになった
  * コンテナランタイムは Docker から Containerd に変更された
* プラットフォームバージョン移行時の考慮事項
  * ネットワークトラフィックがタスク ENI を使用するように変更されている
  * VPC エンドポイント
    * ECR 使用時は S3, ecr.dkr, ecr.api のエンドポイントが必要
    * Secrets Manager 使用時は Secrets Manager のエンドポイントが必要
    * Systems Manager Parameter 使用時は Sysems Manager のエンドポイントが必要
    * セキュリティグループでエンドポイントと疎通がとれるようになっている必要がある


[Windows platform versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform-windows-fargate.html)

* `platformFamily`, `operatingSystemFamily` にイメージ内容と合致したものを指定する必要がある。[Matching container host version with container image versions](https://learn.microsoft.com/en-us/virtualization/windowscontainers/deploy-containers/version-compatibility?tabs=windows-server-2022%2Cwindows-11#matching-container-host-version-with-container-image-versions) を参照のこと


[AWS Fargate task maintenance](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-maintenance.html)

* サービスを使用している場合
  * `minimum healthy percent` を下回らないようにタスクが置き換えられる。例えば `minimum healthy percent` が 100 の場合は新規タスク起動後に既存タスクを停止する
  * Host issue の場合は通知されない。[Task maintenance](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/AWS_Fargate.html#fargate-task-retirement) に表でまとめられている
* 事前に通知が届くようになっている。AWS Health Dashboard もしくは AWS アカウントに設定された E メールアドレスに送信されるメールから確認可能
* アカウント設定により、リタイアメントまでの待機期間を設定可能。ただし、クリティカルなセキュリティアップデート時は即座にタスクをリタイアする場合がある



## Task definitions

[Task definition states](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-definition-state.html)

* 以下のステータスがある
  * ACTIVE
  * INACTIVE: 新規タスク、サービスの作成ができない。既存タスクへの影響はない
  * DELETE_IN_PROGRESS: 稼働中の ECS リソースがあると削除が完了しない。タスク停止後に、タスク定義削除まで 1 時間まで要する場合がある。また、デプロイ、タスクセットが削除されてからタスク定義削除まで 24 時間まで要する場合がある


### Launch Types

[Amazon ECS 起動タイプ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/launch_types.html)

Fargate, EC2, External の 3 つの起動タイプがある。

Fargate 起動タイプに向いているワークロード

* 低いオーバーヘッドのために最適化する必要がある大規模なワークロード
* 時折バーストが発生する小さなワークロード
* 小さなワークロード
* バッチワークロード

EC2 起動タイプに向いているワークロード

* 一貫して高 CPU コアとメモリ使用量を必要とするワークロード
* 料金のために最適化する必要がある大規模なワークロード

外部起動タイプ

* オンプレミスサーバ、仮想サーバを使用可能


### Container Image

[コンテナイメージ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/container-considerations.html)

* コンテナ内で単一のアプリケーションプロセスを実行
* コンテナの有効期間は、アプリケーションプロセスが実行されている期間
* SIGTERM シグナルを処理できるように実装すること
* ログを stdout, stderr に書き込むように実装すること
* ベストプラクティスとしてイメージごとに一意のタグ付与を推奨。git commit に対応するタグの付与を推奨。ECR ではイミュータブルなイメージタグを有効化することを推奨
* コンテナごとにリソースの共有が必要な場合は、一つのタスク定義内に複数のコンテナを定義する。そうでない場合はタスク定義ごとにコンテナを分けることで、スケーリングなどを独立して行えるメリットを得られる


### Task definition

[タスク定義](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/architect-task-def.html)

* タスク定義ファミリーごとに異なる IAM ロールの使用を推奨


### Networking Models

[Amazon EC2 インスタンスでホストされているタスクのタスクネットワーキング](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-networking.html)

* awsvpc を推奨
* EC2 Windows は default, awsvpc のみサポート
* default モードは Windows のみ
  * Windows 上の Docker の組み込み仮想ネットワークを使用
  * Windows の組み込み仮想ネットワークは nat Docker ネットワークドライバーを使用


『ベストプラクティスガイド』の [ネットワークモードの選択](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/bestpracticesguide/networking-networkmode.html) も各ネットワークモードの参考になる。

[awsvpc](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking-awsvpc.html)

* タスクごとに 1 つの ENI が割り当てられる。VPC フローログに記録される
* 同一タスク内のコンテナは localhost 経由で通信できる
* VPC がデュアルスタックモードでサブネットが IPv6 CIDR ブロックの場合、タスク ENI は IPv6 アドレスも受け取る
* リクエスタマネージド型の ENI になっている。ENI は削除、修正が不可
* アカウント設定で `awsvpcTrunking` を有効にしているとトランク ENI が作成される
* サービスにリンクされたロールが必要
* インスタンスタイプごとに ENI 数のクォータがある。プライマリネットワークインタフェースでも 1 個消費される点にも注意。`awsvpcTrunking` を enabled に設定している場合はより大きな数の ENI をアタッチできる
* タスク ENI にパブリック IP アドレスが付与されない。よって、NAT Gateway もしくは VPC エンドポイントを使用する必要がある
* タスク定義内のコンテナが開始される前に、各タスクに Amazon ECS コンテナエージェントによって追加の pause コンテナが作成される。次に、amazon-ecs-cni-plugins CNI プラグインを実行して pause コンテナのネットワーク名前空間が設定される。その後、エージェントによってタスク内の残りのコンテナが開始されるため、pause コンテナのネットワークスタックが共有される。つまり、タスク内のすべてのコンテナは ENI の IP アドレスによってアドレス可能であり、localhost インターフェイス経由で相互に通信できるようになる。
* ELB サポートは ALB, NLB のみ。CLB はサポートされない。ターゲットのタイプは `ip` にする必要がある
* デュアルスタックモード
 * アカウント設定の `dualStackIPv6` を enabled にしておく必要がある
 * VPC, サブネットが IPv6 で構成されている必要がある


[host](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/networking-networkmode-host.html)

* 推奨されないネットワークモード
* ポートが重複しないようにする必要がある
* コンテナがホストになりすませるようになり、localhost 経由でホスト側にアクセスできるため、セキュリティ面でも問題がある


[bridge](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/networking-networkmode-bridge.html)

* ホスト側の異なるポート番号に関連づけることができる
* 動的ポートマッピングを使用可能。ホストポート未指定の場合は動的ポートマッピングになる。ELB もしくは Cloud Map によりクライアント側からのアクセス用のマッピングを行う


[Task networking for tasks hosted on Fargate](https://docs.aws.amazon.com/AmazonECS/latest/userguide/fargate-task-networking.html)

* プライマリプライベート IP アドレスを備えた ENI が提供される
* オプションでパブリック IP アドレスを付与できる
* VPC がデュアルスタックモードに対応していて IPv6 CIDR ブロックを備えたサブネットを使用する場合、タスクの ENI にも IPv6 アドレスが割り当てられる
* ECR からイメージをプルする場合は、エンドポイントとの疎通性(インターネット or VPC エンドポイント)が必要
* PV 1.4.0 以降では単一のタスク ENI が設定される。PV 1.3.0 では更に Fargate ENI が設定され、こちらを通る通信は VPC Flow Logs で捕捉されない
* Fargate によってアタッチされた ENI は削除不可
* IPv6 を備えたサブネットの場合は、IPv6 アドレスのみが割り当てられる
* タスク ENI はジャンボフレームをサポート


### Data volumes

[タスクでのデータボリュームの使用](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/using_data_volumes.html)

利用可能なボリューム

* EFS
* FSx for Windows File Server
* Docker ボリューム(/var/lib/docker/volumes に作成される Docker マネージドボリューム)
* バインドマウント


[Fargate タスクストレージ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/fargate-task-storage.html)

* Windows PV 1.0.0
  * 最低 20 GB のエフェメラルストレージを受け取る。`ephemeralStorage` により最大 200 GB まで拡張可能
  * コンテナイメージもここに格納されるので、利用できるのはイメージの利用量を差し引いた分
* Linux PV 1.4.0
  * 最低 20 GB のエフェメラルストレージを受け取る。`ephemeralStorage` により最大 200 GB まで拡張可能
  * コンテナイメージもここに格納されるので、利用できるのはイメージの利用量を差し引いた分
  * エフェメラルストレージの使用量はタスクメタデータエンドポイント v4 にて取得できる
  * Container Insights を有効化するとエフェメラルストレージの予約量、使用量が取得できる
* Linux PV 1.3.0
  * 10 GB の Docker Layer ストレージ、ボリュームマウント用の追加の 4 GB のストレージを受け取る


[Amazon EFS ボリューム](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/efs-volumes.html)

* 独自の AMI を使用する場合 `amazon-efs-utils` パッケージをインストールし、`amazon-ecs-volume-plugin` サービスを起動する必要がある
* Fargate
  * PV 1.4.0 から対応
  * スーパーバイザーコンテナにより EFS ボリュームが管理される。このコンテナによりタスクメモリが少しだけ使用される。Container Insights では aws-fargate-supervisor コンテナとして表示される
* 外部インスタンスではサポートされない
* コンテナエージェント設定の `ECS_ENGINE_TASK_CLEANUP_WAIT_DURATION` はデフォルト値の 1 時間よりも短くすることを推奨。FS マウント認証情報の有効期限切れを防ぎ、使用されていないマウントをクリーンアップするのに役立つ
* アクセスポイントを使用することでアクセスポイントを介したすべてのファイルシステム要求に対してユーザーアイデンティティ (ユーザーの POSIX グループなど) を適用できる。ファイルシステムに対して別のルートディレクトリを適用することも可能。


[FSx for Windows File Server ボリューム](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/wfsx-volumes.html)

* 有効なドメインに参加している ECS Windows EC2 インスタンスである必要がある。Fargate は未対応
* Active Directory へのドメイン参加と FSx for Windows File Server ファイルシステムをアタッチするために使用される、認証情報を含む AWS Secrets Manager シークレットまたは SystemsManager パラメータが必要
* `authorizationConfig` にてドメイン、認証情報の ARN を指定する


[Docker ボリューム](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/docker-volumes.html)

* EC2 起動タイプ、または外部インスタンスを使用する場合にのみサポート
* `/var/lib/docker/volumes` にデータボリュームが作成される
* サードパーティー製ドライバーを使用する場合は、コンテナエージェントを起動する前に必ずドライバーをコンテナインスタンスにインストールしアクティブ化しておく必要がある
* driver により使用する Docker ボリュームドライバーを指定。`docker volume ls` により取得できる名前を指定する
```json
    "volumes": [
        {
            "name": "string",
            "dockerVolumeConfiguration": {
                "scope": "string",
                "autoprovision": boolean,
                "driver": "string",
                "driverOpts": {
                    "key": "value"
                },
                "labels": {
                    "key": "value"
                }
            }
        }
    ]
```
* NFS の場合は `local` ドライバーによってマウントできる
```json
"volumes": [
       {
           "name": "NFS",
           "dockerVolumeConfiguration" : {
               "scope": "task",
               "driver": "local",
               "driverOpts": {
                   "type": "nfs",
                   "device": "$NFS_SERVER:/mnt/nfs",
                   "o": "addr=$NFS_SERVER"
               }
           }
       }
   ]
```


[バインドマウント](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/bind-mounts.html)

* `Dockerfile` の `VOLUME` ディレクトリの絶対パスが `containerPath` と同じ場合は、VOLUME 内のデータがデータボリュームにコピーされる
* ボリュームはデフォルトで、タスク停止後の 3 時間後に削除される。この期間は `ECS_ENGINE_TASK_CLEANUP_WAIT_DURATION` により設定可能
* 以下のようなユースケースにも対応可能
  * 空のデータボリュームをマウント。`volumes` にて `name` だけを指定してボリュームを作成するとよい
  * EC2 インスタンス側のライフサイクルのボリュームを作成するには `volumes` にて `host.sourcePath` を指定
  * EC2 の場合は `volumesFrom.sourceContainer` を使用して別コンテナからマウントすることも可能。元コンテナが `mountPoints` でマウントしているディレクトリの他 `Dockerfile` の `VOLUME` の内容も含まれる


### swap

[コンテナスワップ領域の管理](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/container-swap.html)

* Linux EC2 のみ対応
* コンテナ単位でスワップを設定できる
* `maxSwap` にてコンテナが使用できるコンテナのスワップ最大量を MiB 単位で指定できる。0 の場合はスワップを使用しない
* `swappiness` が 100 の場合は積極的にスワップが使用される。0 〜 100 の間で指定でき、デフォルト値は 60。Amazon Linux 2023 ではサポートされない
* スワップの合計使用量はコンテナが予約したメモリの 2 倍まで
* インスタンス側で swap を有効化しておく必要がある。ECS Optimized AMI ではデフォルトでは有効化されていない


### Windows

[Amazon EC2 Windows タスク定義の考慮事項](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/windows_task_definitions.html)

* 以下のパラメータはサポートされない
  * containerDefinitions
    * disableNetworking
    * dnsServers
    * dnsSearchDomains
    * extraHosts
    * links
    * linuxParameters
    * privileged
    * readonlyRootFilesystem
    * user
    * ulimits
  * volumes
    * dockerVolumeConfiguration
  * cpu (Windows ではコンテナレベルでの指定を推奨)
  * memory (Windows ではコンテナレベルでの指定を推奨)
  * proxyConfiguration
  * ipcMode
  * pidMode


### Use cases

[タスク定義のユースケース](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/use-cases.html)

[Amazon ECS での GPU の使用](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-gpu.html)

* 外部インスタンスの登録時には `--enable-gpu` の指定が必要
* エージェント設定ファイルで `ECS_ENABLE_GPU_SUPPORT` の設定が必要
* ECS ではコンテナに対して環境変数 `NVIDIA_VISIBLE_DEVICES` を割り当てる
* タスク定義内のコンテナ定義にて GPU 使用数を指定する
```json
{
  "containerDefinitions": [
     {
        ...
        "resourceRequirements" : [
            {
               "type" : "GPU", 
               "value" : "2"
            }
        ],
     },
...
}
```


[Amazon ECS での動画トランスコーディングの使用](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-vt1.html)

* `placementConstraints` にて特定のインスタンスタイプのみで起動するよう制約をかけることが可能
* `linuxParameters.devices` にて特定の U30 カードを使用するように `/dev` 下のパスを指定


[Amazon ECS で Amazon Linux 2 の AWS ニューロンを使用する](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-inference.html)

* 機械学習のワークロード用に、Amazon EC2 Trn1、Amazon EC2 Inf1、Amazon EC2 Inf2 インスタンスをクラスターに登録できる
  * Amazon EC2 Trn1 インスタンスは、AWS Trainium チップを搭載
  * Amazon EC2 Inf1 インスタンスと Inf2 インスタンスは、AWS Inferentia チップを搭載
  * 機械学習モデルは、専用の Software Developer Kit (SDK) である AWS Neuron を使用してコンテナにデプロイ
* デバイスのパスがインスタンスタイプごとに異なる点に注意


[Amazon ECS での深層学習用 DL1 インスタンスの使用](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-dl1.html)

* Amazon EC2 DL1 インスタンスでは Habana Labs (インテル子会社) の Gaudi アクセラレータを搭載


[Amazon ECSの 64-bit ARM ワークロードの操作](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-arm64.html)

* `runtimePlatform.cpuArchitecture` にて `ARM64` を指定


### Logging

[awslogs ログドライバーを使用する](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/using_awslogs.html)

* awslogs ログドライバーは STDOUT, STDERR の IO ストリームを CloudWatch Logs に送る
* ECS Optimized AMI でない場合は、コンテナエージェントにて `ECS_AVAILABLE_LOGGING_DRIVERS` の設定が必要
* タスク実行ロールに logs:CreateLogStream および logs:PutLogEvents の許可が必要
* タスク定義のオプション
  * awslogs-create-group
  * awslogs-region
  * awslogs-group
  * awslogs-stream-prefix
  * awslogs-datetime-format
  * awslogs-multiline-pattern
  * mode
    * デフォルトは `blocking` モード。CloudWatch Logs へのログのフローが中断された場合、stdout, stderr へのコンテナコードからの呼び出しはブロックされる。アプリケーションが応答しなくなる可能性がある
    * `non-blocking` に設定すると `max-buffer-size` で指定されたメモリ内の中間バッファに保管される。サービスの可用性を確保したいが多少のログ欠損があっても良い場合はこちらのモードを選択する
  * max-buffer-size
    * バッファがいっぱいになるとそれ以上ログは保存できず、保存できなくなったログは失われる


[カスタムログルーティング](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/using_firelens.html)

* Windows は未サポート
* 24224 番ポートで LISTEN するため、当該ポートへのインバウンドトラフィックを許可してはならない。また、ポートマッピングで当該ポートは除外する必要がある
* bridge の場合は FireLens コンテナの方がアプリケーションコンテナよりも先に起動するように依存関係の制御が必要
* デフォルトでは stdout, stderr のコンテナログに以下のメタデータが追加される。`enable-ecs-log-metadata` を `false` に設定することで追加されないようにできる
```
"ecs_cluster": "cluster-name",
"ecs_task_arn": "arn:aws:ecs:region:111122223333:task/cluster-name/f2ad7dba413f45ddb4EXAMPLE",
"ecs_task_definition": "task-def-name:revision",
```
* IAM
  * ログの PUT 先に応じた IAM 設定がタスクロール側に必要
  * 以下の場合はタスク実行ロールの設定が必要
    * ECR からのイメージプル。Secrets Manager からのデータ参照
    * S3 上のカスタム設定ファイルを指定する場合
* Fluentd では `log-driver-buffer-limit` によりメモリにバッファリングされるサイズを指定可能
* 環境変数 `FLUENT_HOST`, `FLUENT_PORT` がコンテナに設定される。これらの環境変数を使用すると stdout を介することなくコードからログルーターに直接ログを送信できる


[FireLens 設定を使用するタスク定義の作成](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/firelens-taskdef.html)

* ログルーターコンテナは `essential` を `true` に設定することを推奨
* ログファイルが生成される。`/fluent-bit/etc/fluent-bit.conf`(Fluent Bit), `/fluentd/etc/fluent.conf`(Fluentd)
* `config-file-type` は `s3` or `file`。ただし、Fargate は `file` のみサポート


### Private registory

[タスクのプライベートレジストリの認証](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/private-auth.html)

* `containerDefinitions.repositoryCredentials.credentialsParameter` にて Secrets Manager の ARN を指定
* タスク実行ロールに `secretsmanager:GetSecretValue` の許可が必要。カスタムの KMS キーを使用する場合は更に `kms:Decrypt` の許可が必要


### Environment Variable

[コンテナへの環境変数の受け渡し](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/taskdef-envfiles.html)

* `containerDefinitions.environmentFiles` にて S3 上のファイルを指定可能。ファイルの拡張子は .env である必要がある


[Secrets Manager の使用](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/secrets-app-secrets-manager.html)


[AWS Systems Manager パラメータストアの使用](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/secrets-app-ssm-paramstore.html)


### Example

[タスク定義の例](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/example_task_definitions.html)

* 以下の例が載っている
  * Web サーバー
  * 各種ログドライバー
  * entryPoint, command の指定内容
  * dependsOn
  * Windows コンテナ



## ECS Cluster

[Amazon ECS clusters](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/clusters.html)

* サービス、タスクを管理する単位
* EC2, Fargate の両起動タイプが共存可能
* サブネットは AZ, Local Zone, Wavelength Zone, AWS Outposts を含められる
* マネジメントコンソールからクラスターを作成すると、CloudFormation のスタックが作られる
* クラスターのデフォルトの Service Connect 名前空間を設定可能
* モニタリングオプション(Container Insights)
* デフォルトのキャパシティープロバイダー戦略を設定可能


[コンソールを使用した Fargate 起動タイプ用のクラスター作成](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/create-cluster-console-v2.html)

デフォルトでは以下の内容でクラスターが作成される。一部は変更可能

* Fargate および Fargate Spot キャパシティープロバイダーを使用
+ 選択したリージョンのデフォルト VPC 内のすべてのデフォルトサブネットでタスクとサービスを起動
* Container Insights は使用しない
* AWS CloudFormation に 3 つのタグが構成
* AWS Cloud Map に、クラスターと同じ名前のデフォルトの名前空間を作成


### Capacity Provider

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
* 6 個までのキャパシティープロバイダーを設定可能
* **キャパシティプロバイダー戦略**により、タスクをどのキャパシティプロバイダーに配置するかを設定できる。base, weight の２つの設定値がある
  * base: 最低何個のタスクを起動するか。一つのキャパシティープロバイダーにのみ設定可能
  * weight: どのキャパシティプロバイダーにタスクを割り当てるかの比率を設定する。0 の場合は、当該キャパシティープロバイダーは使用されない


[AWS Fargate capacity providers](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/fargate-capacity-providers.html)

* キャパシティープロバイダーを作成する必要はない。ただし、クラスターに関連づける必要はある
* マネジメントコンソールより ECS クラスターを Networking only で作成した場合に FARGATE, FARGATE_SPOT が自動的に設定された状態になっている。
* FARGATE_SPOT のキャパシティプロバイダーのタスクがスポットの中断により停止する際は、タスク停止の 2 分前に EventBridge よりワーニングが送られる。また、SIGTERM シグナルがタスクに送られる。キャパシティに空きがある場合は新規タスクの起動を試みる。需要が多い場合は起動が遅くなる場合がある
* FARGATE_SPOT は Windows では未サポート。Linux でも ARM64 の場合は未サポート


[Auto Scaling group capacity providers](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/asg-capacity-providers.html)

* 空の Auto Scaling グループの作成を推奨。既存の Auto Scaling グループの場合、起動済みのインスタンスが正常にキャパシティープロバイダーに登録されないことがある
* インスタンスの重みづけは未サポート
* managed termination protection
  * 使用する際 managed scaling が有効になっている必要がある。そうしないと managed termination protection は動作しない
  * 有効にすることでタスクが存在する EC2 インスタンスについてスケールインから保護することができる。Auto Scaling Group 側でもスケールインからのインスタンス保護の設定が有効になっている必要がある
* **managed scaling を有効にすることで、タスク数に応じて ECS インスタンスがスケールする。**Auto Scaling Group に自動作成されるスケーリングポリシーを変更、削除してはならない
* ウォームプールを使用可能
  * ユーザーデータで `ECS_WARM_POOLS_CHECK` を設定する


### Cluster Auto Scaling

[Amazon ECS クラスターの Auto Scaling](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/cluster-auto-scaling.html)

* スケーリングの制御に `CapacityProviderReservation` のメトリクスを使用する。
  * 計算式: CapacityProviderReservation = (number of instances needed) / (number of running instances) x 100
  * `CapacityProviderReservation` > `targetCapacity` の場合にスケールアウトする
  * `CapacityProviderReservation` < `targetCapacity` の場合にスケールインする。終了するインスタンスは Auto Scaling Group の終了ポリシーにより決定される
* 考慮事項
  * スケーリングポリシーを変更、追加してはならない
  * スケーリングにはサービスにリンクされたロール `AWSServiceRoleForECS` を使用する
  * キャパシティプロバイダーを作成、更新する IAM エンティティには `autoscaling:CreateOrUpdateTags` の許可が必要
  * Auto Scaling グループがら AmazonECSManaged タグを削除してはならない
  * マネージドスケーリングが有効の場合はキャパシティープロバイダーは単一のクラスターにしか関連づけることができない。無効の場合は複数クラスターに関連づけることができる
  * 既存のキャパシティーから配置戦略、配置制約を使用
* Managed termination protection
  * タスクが稼働しているインスタンスはスケールインから保護される。しかし DAEMON タイプは例外
  * Auto Scaling Group にてインスタンスのスケールイン保護が有効になっている必要がある
  * Auto Scaling Group からデタッチする場合は、ECS クラスターからも登録解除が必要
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


### Cluster Concepts

[クラスターの管理](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/clusters-concepts.html)

* クラスターの状態には以下のものがある
  * ACTIVE
  * PROVISIONING
  * DEPROVISIONING
  * FAILED
  * INACTIVE
* インスタンスは一つのクラスターにしか登録できない
* IAM ポリシーで、クラスターごとに権限制御できる
* デフォルトの Service Connect 名前空間を設定可能


### Create Capacity

[キャパシティーの作成](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/create-capacity.html)

* コンテナインスタンスではコンテナエージェントを稼働させる必要がある。ECS Optimized AMI にはインストール済み
* コンテナインスタンスには IAM ロールの設定が必要
* ECS Optimized AMI の 20200430 以降は IMDSv2 がサポートされている
* 各エンドポイントとの疎通性が必要
* コンテナインスタンスを登録解除したあとに別のクラスターに登録してはならない
* 同じコンテナインスタンスを停止して、インスタンスタイプを変更することはできない
* スポットインスタンスを使用する場合は、複数の AZ、混合インスタンスポリシーを使用する


#### Linux Instance

[Linux インスタンス](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-linux.html)

Bottlerocket も選択肢の一つ

* 必須要件
  * Linux Kernel (バージョン 3.10 以上)
  * Amazon ECS コンテナエージェント
  * Docker Daemon (バージョン 1.9.0 以上)
* 推奨
  * ecs-init


[Amazon ECS に最適化された AMI](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-optimized_AMI.html)

* [AMI の Change Log](https://github.com/aws/amazon-ecs-ami/blob/main/CHANGELOG.md)
* [AMI のバージョン履歴](https://github.com/aws/amazon-ecs-ami/releases)
* [Docker Release Note](https://docs.docker.com/engine/release-notes/)
* [Nvidia Driver Document](https://docs.nvidia.com/datacenter/tesla/index.html)
* [ECS Agent の Change Log](https://github.com/aws/amazon-ecs-agent/blob/master/CHANGELOG.md)


[Amazon ECS に最適化された AMI メタデータを取得する](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/retrieve-ecs-optimized_AMI.html)

* AMI ID は SSM パラメータにて取得可能


[AMI ストレージ設定](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-ami-storage-config.html)

* Amazon Linux 2023 AMI 
  * 単一の 30 GiB のボリューム
  * デフォルトファイルシステムは xfs
  * Docker は overlay2 ストレージドライバー
* Amazon Linux 2 AMI 
  * 単一の 30 GiB のボリューム
  * デフォルトファイルシステムは xfs
  * Docker は overlay2 ストレージドライバー
* Amazon Linux AMI
  * オペレーティングシステム用に 8 GiB ボリュームが /dev/xvda にアタッチ。ルートとしてマウント。
  * Docker によるイメージとメタデータの保存用に 22 GiB のボリュームが /dev/xvdcz に追加でアタッチ。
  * devicemapper を使用

補足: [Use the OverlayFS storage driver](https://docs.docker.com/storage/storagedriver/overlayfs-driver/)

Docker ボリュームの拡張方法

* サイズを増加したデータストレージボリュームで新しいインスタンスを起動するのが簡単
* Amazon Linux の場合はボリュームが分かれているため、別途対応が必要

チューニング箇所

* `ECS_ENGINE_TASK_CLEANUP_WAIT_DURATION` を短くする


[Amazon ECS 最適化 Linux AMI のビルドスクリプト](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-ami-build-scripts.html)

* ビルドスクリプトレポジトリには以下が含まれている
  * HashiCorp Packer テンプレート
  * Amazon ECS に最適化された AMI の各 Linux バリアントを生成するためのビルドスクリプト
  * AMI を構築するための Makefile


[Bottlerocket](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-bottlerocket.html)

* 最小限のパッケージのみが含まれている
* パッケージマネージャは含まれていない。イミュータブルな設計思想
* 管理ツールを使用することで、SSH アクセスを獲得できる
* コントロールコンテナが稼働している


[Linux コンテナ向け gMSAs を使用する](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/linux-gmsa.html)

* Linux コンテナの Active Directory 認証をサポート。グループ管理サービスアカウント(gMSA)を使用
* コンテナインスタンス側で credentials-fetcher デーモンが必要。Active Directory ドメインコントローラーから認証情報を取得し、gMSA 認証情報をコンテナインスタンスに転送
* Fargate ではサポートされていない


[Amazon ECS コンテナエージェントをインストールする](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-agent-install.html)

* Amazon Linux 2
  * `amazon-linux-extras` コマンドでインストール可能
  * systemctl で ecs サービスを起動できる
  * `curl -s http://localhost:51678/v1/metadata | python -mjson.tool` コマンドにて稼働状況を確認可能
* 非 Amazon Linux EC2 インスタンス
  * パッケージのファイルをダウンロードし、パッケージマネージャによりインストールする
* コンテナエージェント
  * host ネットワークモードで稼働している
  * コンテナエージェントから稼働したコンテナは http://169.254.169.254 へのアクセスがブロックされる


#### Windows Instance

[Windows インスタンス](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-windows.html)


[Amazon ECS に最適化された AMI](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-optimized_windows_AMI.html)

* 考慮事項
  * 配置制約 `memberOf(ecs.os-type=='windows')` を使用するのが手
  * タスクロールを使用するには、認証情報プロキシを使用する。このプロキシがホストの 80 番ポートを占有するため、ホストポート 80 はタスク側では使用できない
  * ホストとコンテナで Windows バージョンが一致している必要がある


[Windows コンテナでの gMSAs の使用](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/windows-gmsa.html)

* グループマネージドサービスアカウント (gMSA) と呼ばれる特殊なサービスアカウントを使用して、Windows コンテナの Active Directory 認証をサポート


[Amazon ECS に最適化された独自の Windows AMI の作成](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/windows-custom-ami.html)

* EC2 Image Builder を使用して、Amazon ECS 最適化された独自のカスタム Windows AMI を構築できる


#### 外部インスタンス

[外部インスタンス（Amazon ECS Anywhere）](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-anywhere.html)

* EXTERNAL 起動タイプ
* Fedora 32、Fedora 33 は cgroup v2 を使用するため grub 設定変更が必要
* IAM ロールが必要
* `ecs.capability.external` 属性が設定される
* ECS Exec はサポートされている
* 以下は未サポート
  * ELB
  * サービスディスカバリ
  * awsvpc
  * `UpdateContainerAgent` API
  * キャパシティープロバイダー
  * SELinux
  * EFS
  * App Mesh
* エンドポイントとの疎通性が必要
  * ecs-a-*.region.amazonaws.com: タスクを管理
  * ecs-t-*.region.amazonaws.com: タスクとコンテナのメトリクスを管理
  * ecs.region.amazonaws.com: ECS のサービスエンドポイント
  * ssm.region.amazonaws.com: Systems Manager のエンドポイント
  * ec2messages.region.amazonaws.com: Systems Manager エージェントと Systems Manager サービスの間の通信に使用するエンドポイント
  * ssmmessages.region.amazonaws.com: ession Manager サービスでセッションチャネルを作成および削除するために必要なサービスエンドポイント

タスクロール用の iptables 設定
```
$ sysctl -w net.ipv4.conf.all.route_localnet=1
$ iptables -t nat -A PREROUTING -p tcp -d 169.254.170.2 --dport 80 -j DNAT --to-destination 127.0.0.1:51679
$ iptables -t nat -A OUTPUT -d 169.254.170.2 -p tcp -m tcp --dport 80 -j REDIRECT --to-ports 51679
```


### Manage Capacity

[キャパシティ管理](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/manage-capacity.html)

Fargate は AWS 側の管理。EC2 or External は以下項目の管理が必要

* コンテナインスタンスの起動
* コンテナインスタンスのブートストラップ
* 起動時のタスク開始
* ENI トランキングの使用
* メモリの管理
* コンテナインスタンスのリモート管理
* コンテナエージェントと Docker デーモンの両方での HTTP プロキシの使用
* コンテナエージェントの更新
* コンテナインスタンスの登録解除


#### Container Agent

[Amazon ECS Linux コンテナエージェント](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-agent-versions.html)

* 最新バージョンの使用を常に推奨
* ライフサイクル
  * コンテナインスタンスを停止した場合 ACTIVE のままだが、エージェント接続ステータスは FALSE となる
  * 以下のステータスがある
    * ACTIVE 状態の場合 RunTask API によるリクエストを受け入れることができる。
    * DRAINING: 新規タスクが配置されなくなる。サービスのタスクは可能であれば削除される。
    * INACTIVE: コンテナインスタンスを登録解除もしくは終了した場合。1 時間以内はコンテナインスタンスの情報を取得可能


[Amazon ECS コンテナエージェントの設定](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-agent-config.html)

* 環境変数を `/etc/ecs/ecs.config` に設定する
* パラメータの一覧は https://github.com/aws/amazon-ecs-agent/blob/master/README.md を参照
* 機密性の高い情報はインスタンス起動時に S3 から PULL するなどの方法で取得する

旧ドキュメントの内容

* パラメータ例
  * `ECS_DATADIR`: コンテナの状態を保存するディレクトリパス
  * `ECS_ENABLE_TASK_IAM_ROLE`: タスクの IAM ロールを有効化するかどうか
  * `ECS_DISABLE_IMAGE_CLEANUP`: 自動イメージクリーンアップを行うかどうか
  * `ECS_AWSVPC_BLOCK_IMDS`: awsvpc ネットワークモードを使用して起動されるタスクのインスタンスメタデータへのアクセスをブロックするかどうか
  * `ECS_AWSVPC_ADDITIONAL_LOCAL_ROUTES`: awsvpc ネットワークモードでは、これらのプレフィックスへのトラフィックは、タスク Elastic Network Interface ではなく、ホストブリッジ経由でルーティングされる
  * `ECS_TASK_METADATA_RPS_LIMIT`: タスクメタデータエンドポイントのスロットリングに使用する値
  * `ECS_ENABLE_SPOT_INSTANCE_DRAINING`: スポットインスタンスのドレイニングを有効化するかどうか


[コンテナインスタンスのプライベートレジストリの認証](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/private-auth-container-instances.html)

* EC2 起動タイプではプライベートレジストリの認証を行うことができる
* `ECS_ENGINE_AUTH_TYPE` に認証タイプ、`ECS_ENGINE_AUTH_DATA` に認証情報を設定


[コンテナインスタンスのドレイン](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/container-instance-draining.html)

* ドレイニング状態での動作
  * サービスの一部のタスク
    * PENDING 状態のサービスタスクは直ちに停止
    * RUNNING 状態のサービスタスクは `minimumHealthyPercent`、`maximumPercent` に従って置き換え
    * `minimumHealthyPercent` が 100 の場合は、新規起動したタスクが正常と見なされるまで既存タスクを停止しない。ELB を使用している場合は ELB からも healthy と判定される必要がある
  * スタンドアローンのタスク
    * PENDING, RUNNING 状態のタスクは影響を受けない
    * 全てのタスクが STOPPED になるとドレインは完了する
  * コンテナインスタンスのステータスは DRAINING になる。再度使用する場合は ACTIVE に設定する必要がある


[自動化タスクとイメージのクリーンアップ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/automated_image_cleanup.html)

* `ECS_ENGINE_TASK_CLEANUP_WAIT_DURATION`: 停止されたコンテナを削除するまでの時間。デフォルトは 3 時間
* `ECS_DISABLE_IMAGE_CLEANUP`: `true` に設定した場合はイメージの削除は実施されない
* `ECS_IMAGE_CLEANUP_INTERVAL`: イメージをチェックする間隔。デフォルトは 30 分
* `ECS_IMAGE_MINIMUM_CLEANUP_AGE`: イメージが取得されてから削除対象になるまでの最短の時間。デフォルトは 1 時間
* `ECS_NUM_IMAGES_DELETE_PER_CYCLE`: 1 回のクリーンアップサイクルで削除するイメージ数。デフォルトは 5 個


#### Manage Linux

[Linux コンテナインスタンス管理](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/manage-linux.html)


[Amazon ECS Linux コンテナインスタンスの起動](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/launch_container_instance.html)

* `ECS_CONTAINER_INSTANCE_TAGS`: タグは ECS 用。EC2 API では抽出されないタグ
* `ECS_CONTAINER_INSTANCE_PROPAGATE_TAGS_FROM`: EC2 インスタンスからタグを伝搬
* `ECS_ENABLE_SPOT_INSTANCE_DRAINING` を `true` に設定することでスポット中断時に DRAINING へと遷移できる
* 既存のインスタンス上でパラメータを更新する場合は `ecs.config` を修正した後に ecs サービスを再起動する
* コンテナエージェントのメタデータの取得は `curl http://localhost:51678/v1/metadata` から可能


[Amazon EC2 ユーザーデータを使用してコンテナインスタンスをブートストラップする](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/bootstrap_container_instance.html)

* コンテナエージェントの設定は `/etc/ecs/ecs.config` により行う
* Docker デーモンの設定は `/etc/docker/daemon.json` により行う


[コンテナインスタンス起動時のタスクの開始](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/start_task_at_launch.html)

* `StartTask` API により特定のコンテナインスタンス上でタスクを起動できる
* 起動したインスタンス上でタスクを起動する方法としては、コンテナエージェントから提供されるコンテナインスタンスの ARN を取得して、`StartTask` API を実行する


[Elastic Network Interface のトランキング](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/container-instance-eni.html)

* `awsvpc` ではタスクごとに ENI がアタッチされる
* インスタンスタイプごとに ENI 数の上限がある。プライマリネットワークインタフェースは `awsvpc` 用としては使用できない
* アカウント設定で `awsvpcTrunking` を有効化すると追加の ENI を使用できる
* 考慮事項
  * コンテナエージェント 1.28 以降
  * Windows 未サポート
  * インスタンスの設定において、リソースベースの IPv4 DNS リクエストがオフになっている必要がある。修正は `aws ec2 modify-private-dns-name-options --instance-id i-xxxxxxx --no-enable-resource-name-dns-a-record --no-dry-run`
  * 共有サブネットは未サポート
  * コンテナインスタンスとタスクが同一 VPC である必要がある
  * コンテナインスタンスの起動時に失敗すると REGISTRATION_FAILED になる。コンテナインスタンスを describe し statusReason の確認が必要
  * コンテナインスタンスは追加の ENI を受け取り、デフォルトのセキュリティグループが使用される。ECS によって ENI が管理されている
  * サービスにリンクされたロールが必要
  * `PutAccountSettingDefault` は AWS アカウント全体に適用される
  * `PutAccountSetting` は特定の IAM エンティティのみに適用できる。ルートユーザーでないと実行できない
  * 一部のインスタンスファミリーでは未サポート


[コンテナインスタンスメモリ管理](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/memory-management.html)

* Docker の `ReadMemInfo()` 関数を使用してオペレーティングシステムで使用可能な合計メモリのクエリを実行
* `ECS_RESERVED_MEMORY` によりシステム用に予約するメモリ量を設定可能


[AWS Systems Manager を使用してコンテナインスタンスをリモート管理する](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ec2-run-command.html)

* SSM の RunCommand を使用する場合、ecsInstanceRole にて許可が必要


[Linux コンテナインスタンス用の HTTP プロキシ設定](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/http_proxy_config.html)

* `/etc/ecs/ecs.config`
  * `HTTP_PROXY=10.0.0.131:3128`
  * `NO_PROXY=169.254.169.254,169.254.170.2,/var/run/docker.sock`
* `/etc/systemd/system/ecs.service.d/http-proxy.conf`
  * `Environment="HTTP_PROXY=10.0.0.131:3128/"`
  * `Environment="NO_PROXY=169.254.169.254,169.254.170.2,/var/run/docker.sock"`
* `/etc/systemd/system/docker.service.d/http-proxy.conf`
  * `Environment="HTTP_PROXY=http://10.0.0.131:3128"`
  * `Environment="NO_PROXY=169.254.169.254"`
* `/etc/sysconfig/docker`
  * `export HTTP_PROXY=http://10.0.0.131:3128`
  * `export NO_PROXY=169.254.169.254,169.254.170.2`


[Amazon ECS コンテナエージェントをアップデートする](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-agent-update.html)


[Amazon ECS 対応 AMI での Amazon ECS コンテナエージェントのアップデート](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/agent-update-ecs-ami.html)

* 推奨される順
  * 最新の AMI から起動する
  * `ecs-init` を最新にする。yum update にて更新できる。更新後は `docker` サービスの再起動が必要
  * `UpdateContainerAgent` API を実行する。PENDING → STAGING → STAGED → UPDATING → UPDATED へと遷移。失敗時は FAILED となる


[Amazon ECS コンテナエージェントの手動更新（Amazon ECS 最適化以外の AMI の場合）](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/manually_update_agent.html)

* ecs-agent サービスを停止してからコンテナイメージをプルする。`docker pull public.ecr.aws/ecs/amazon-ecs-agent:latest`。`docker run` コマンドにより ecs-agent を起動する


#### Manage Windows

[Windows コンテナインスタンス管理](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/manage-windows.html)


[Amazon ECS Windows コンテナインスタンスの起動](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/launch_window-container_instance.html)

ユーザーデータの内容
```powershell
<powershell>
Import-Module ECSTools
Initialize-ECSAgent -Cluster your_cluster_name -EnableTaskIAMRole -EnableTaskENI -AwsvpcBlockIMDS -AwsvpcAdditionalLocalRoutes '["ip-address"]'
</powershell>
```

* `EnableTaskENI`: トランク ENI の使用
* `AwsvpcBlockIMDS`: タスクコンテナの IMDS へのアクセスをブロック
* `AwsvpcAdditionalLocalRoutes`: タスクの名前空間に追加のルートを設定
* `[Environment]::SetEnvironmentVariable("ECS_ENABLE_SPOT_INSTANCE_DRAINING", "true", "Machine")` によりスポット中断時に DRAINING へと遷移


[Amazon EC2 ユーザーデータを使用して Windows コンテナインスタンスをブートストラップする](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/bootstrap_windows_container_instance.html)

* `[Environment]::SetEnvironmentVariable("ECS_ENABLE_AWSLOGS_EXECUTIONROLE_OVERRIDE", $TRUE, "Machine")` により awslogs が使用可能になる


[Windows インスタンスに接続する](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/instance-windows-connect.html)

* RDP により接続


[Windows コンテナインスタンス用の HTTP プロキシ設定](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/http_proxy_config-windows.html)

ユーザーデータにて以下のように設定
```
$proxy = "http://proxy.mydomain:port"
[Environment]::SetEnvironmentVariable("HTTP_PROXY", $proxy, "Machine")
[Environment]::SetEnvironmentVariable("NO_PROXY", "169.254.169.254,169.254.170.2,\\.\pipe\docker_engine", "Machine")
```


[Amazon EC2 でバックアップされたコンテナインスタンスの登録解除](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/deregister_container_instance.html)

* タスク実行中でも登録解除可能。当該タスクは ECS による情報収集の対象外になる
* サービスタスクの場合は、代わりの新規タスクを起動しようとする。ALB から登録解除され awsvpc の場合は ENI が解除される


#### Manage External

[外部コンテナインスタンス管理](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/manage-external.html)


[クラスターへの外部インスタンスの登録](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-anywhere-registration.html)

* AWS Systems Manager マネージドインスタンスとしての登録が必要
* CLI で対応する場合は以下の流れ
  * IAM ロールに対応する Systems Manager のアクティベーションペアを作成
  * ECS Anywhere のインストールスクリプトをダウンロードし実行
* 別の ECS クラスターに登録する際の流れ
  * ecs サービスを停止
  * `/var/lib/ecs/data/agent.db` を削除
  * ecs サービスを起動


[外部インスタンスの登録を解除する](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-anywhere-deregistration.html)

* 次の流れで対応する
  * コンテナインスタンスをドレイン
  * コンテナインスタンスを登録解除
  * SSM のマネージドインスタンスを登録解除
  * ecs サービスを停止
  * 各ディレクトリ、ファイルを削除


[外部インスタンス上のAWS Systems Managerエージェントと Amazon ECS コンテナエージェントを更新しています](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-anywhere-updates.html)

* 次の流れで対応する
  * `ecs-init` パッケージのダウンロード
  * パッケージインストール
  * ecs サービスの再起動



## Account Settings

[アカウント設定](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-account-settings.html)

* アカウントレベルもしくは IAM ユーザー、ロールレベルで特定の機能をオプトイン、オプトアウトできる
* 設定できる項目
  * Amazon リソースネーム (ARN) と ID: 新しい ARN フォーマットになる
  * AWS VPC トランキング
  * CloudWatch Container Insights: クラスターのデフォルト設定で Container Insights が有効になる。クラスターごとに無効に設定することも可能
  * デュアルスタック VPC IPv6: awsvpc のタスクでプライマリプライベート IPv4 アドレスに加えて IPv6 アドレスを備えたタスクの提供をサポート
  * Fargate FIPS-140 コンプライアンス: [AWS Fargate 連邦情報処理標準 (FIPS-140)](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-fips-compliance.html)
  * タグリソース認可: タグ付け時に `ecs:TagResource` の許可が必要
  * AWS Fargate タスク廃止の待機時間: リタイアするまでの期間(0, 7, 14 日)を設定できる。重要なセキュリティパッチの場合は直ちにリタイアされる場合がある


[アカウント設定の変更](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-modifying-longer-id-settings.html)

コンソール上の操作で変更可能


[デフォルトの Amazon ECS アカウント設定の復元](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-reverting-account.html)

デフォルトのアカウント設定を復元可能


[AWS CLI を使用したアカウント設定管理](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/account-setting-management-cli.html)

* put-account-setting-default
  * アカウント全体に適用される
* put-account-setting
  * 自分の IAM エンティティの設定を変更
  * ルートユーザーの場合は対象を指定しなければアカウント全体が対象になる。特定の IAM エンティティのみに設定することも可能


## Scheduling Tasks

[Amazon ECS タスクのスケジューリング](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/scheduling_tasks.html)

* サービススケジューラ
  * タスクが異常となると置き換えタスクが起動する。置き換えタスクが HEALTHY となった後に異常タスクを停止する。ただし、`maximumPercent` が 100 以下の場合は異常タスクをランダムに停止してから置き換えタスクを起動する
  * サービススケジューラ戦略
    * レプリカ: デフォルトでは AZ 間で分散される
    * デーモン: タスク配置制約を満たすインスタンス上に配置される。Fargate は未サポート
* 手動でタスク実行: RunTask の API によりタスクを起動する
* cron ライクなスケジューラ: EventBridge Scheduler を使用してスケジュールを作成できる
* カスタムスケジューラ: Blox のような OSS がある。StartTask ではコンテナインスタンスを指定してタスクを起動する


[Amazon ECS コンソールを使用したスタンドアロンタスクの実行](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs_run_task-v2.html)

* キャパシティープロバイダー戦略もしくは起動タイプを設定
* パブリック IP アドレスの有効化/無効化を設定
* タスク配置戦略は以下から選ぶことができる。
  * AZ Balanced Spread
  * AZ Balanced BinPack
  * BinPack
  * One Task Per Instance
  * Custom
    * タスク配置戦略の設定
    * タスク配置制約は 10 個まで指定可能
* タスクロール、タスク実行ロールを上書き可能
* コンテナコマンド、環境変数を上書き可能
* タグ付けの指定。クラスター、タスク定義のタグ名で自動タグ付けすることも可能


[コンソールを使用してタスクを停止](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/stop-task-console-v2.html)

コンソールの操作からタスクを停止可能


### タスクの配置

[Amazon ECS タスクの配置](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-placement.html)

* サービスから起動されている場合はデフォルトのタスク配置戦略は `attribute:ecs.availability-zone` を使用した `spread`
* Fargate はタスク配置戦略、タスク配置制約をサポートしない。AZ にまたがって分散される。Fargate, Fargate Spot とでスプレッドの動作は別々の管理
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
* 次の情報をタスクグループ名として使用
  * スタンドアロンタスク: タスク定義ファミリ名。もしくは、カスタムタスクグループ名を指定可能
  * サービスの一部として起動されたタスク: サービス名


[タスク配置戦略](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-placement-strategies.html)

* 以下のタスク配置戦略がある。
  * binpack: 一台に詰め込むような動作。field には memory などを指定可能。スケールイン時は利用可能リソースが一番多いインスタンスのタスクを停止する
  * random: ランダムに配置
  * spread: InstanceId もしくは attribute:ecs.availability-zone を設定する。スケールイン時は AZ 間のバランスを保つようにタスクを停止する
* `CreateService`, `UpdateService`, `RunTask` で指定可能


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

以下の例の場合、タスクグループ「service:daemon-service」が存在するコンテナインスタンス上が配置対象になる
```json
"placementConstraints": [
    {
        "expression": "task:group == service:daemon-service",
        "type": "memberOf"
    }
]
```


[クラスタークエリ言語](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/cluster-query-language.html)

タスク配置制約の構文

* 式の構文 `subject operator [argument]`
* subject
  * agentConnected
  * agentVersion
  * attribute:attribute-name
  * ec2InstanceId
  * registeredAt: コンテナインスタンスの登録日
  * runningTasksCount
  * task:group
* 複合した制約も設定可能 `(expression1 or expression2) and expression3`
* 例
  * 引数リスト: `attribute:ecs.availability-zone in [us-east-1a, us-east-1b]`
  * 複合式: `attribute:ecs.instance-type =~ g2.* and attribute:ecs.availability-zone != us-east-1d`
  * アフィニティ: `task:group == service:production`
  * アンチアフィニティ: `not(task:group == database)`


### スケジューリングされたタスク

[スケジュールされたタスク](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/scheduled_tasks.html)

EventBridge スケジューラのスケジュールによってタスクを起動することが可能。


[EventBridge スケジューラのスケジュールされたタスク](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/scheduled_tasks-eventbridge-scheduler.html)

以下のパラメータを設定可能。
* スケジュール
* ターゲット(RunTask)
  * ECS クラスター
  * タスク定義
  * キャパシティープロバイダー戦略もしくは起動タイプ
  * サブネットなどのネットワーク設定
  * タスク配置戦略、タスク配置制約
  * タグ
  * リトライポリシーとデッドレターキュー
    * イベントの最大保持時間
    * 最大再試行回数
  * 暗号化
  * スケジューラ用の IAM ロール


[EventBridge ルールを使用してスケジュールされたタスク](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/scheduled_tasks-event-bridge.html)

* こちらよりも EventBridge スケジューラがおすすめ
* タスクロール、コンテナコマンド、環境変数の上書きが可能


### タスクのライフサイクル

[タスクのライフサイクル](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-lifecycle.html)

* `lastStatus`, `desiredStatus` がある
* 以下のライスサイクルの状態がある
  * PROVISIONING: awsvpc の ENI プロビジョニング
  * PENDING: リソースに空きがない場合はこの状態のままとなる
  * ACTIVATING: サービスディスカバリの設定、複数ターゲットグループへの登録設定
  * RUNNING: 実行中
  * DEACTIVATING: 複数ターゲットグループからの登録解除
  * STOPPING: SIGTERM を送信。StopTimeout 経過後に SIGKILL を送信
  * DEPROVISIONING: ENI のデタッチ、削除
  * STOPPED: タスク停止済み
  * DELETED: `describe-tasks` で表示される
* 複数のターゲットグループを使用している場合は Activationg, Deactivating を経由する



## Services

[Amazon ECS サービス](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs_services.html)

次の機能がある。
* タスク数の維持
* ELB の背後に配置
* タスク起動に連続して失敗した場合はサービス起動の調整ロジックが働く
* サービスディスカバリ
* サービスのステータス
  * ACTIVE
  * DRAINING: サービスを削除したが、まだ実行中のタスクがある場合
  * INACTIVE: サービスを削除し、全てのタスクが STOPPING, STOPPED となっている場合

サービススケジューラ戦略は次の２つ。
* レプリカ: タスク数を維持
* デーモン: コンテナインスタンスごとに一つのタスク

**デーモン**
* タスク配置制約を満たすコンテナインスタンス上にタスクを配置する
* 複数のタスクが同一コンテナインスタンス上で稼働する場合、まず DAEMON のリソースから確保される
* DRAINING ステータスのインスタンス上には配置されない。また、DRANING 状態に遷移した場合は DAEMON タスクも停止する
* DRAINING 時は最後に停止するように動作する


[コンソールを使用したサービスの作成](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/create-service-console-v2.html)

* 設定項目
  * キャパシティープロバイダー戦略もしくは起動タイプ
  * awsvpc の場合はサブネット、セキュリティグループ設定
  * デプロイ失敗時の動作（サーキットブレイカー、ロールバック）
  * Blue/Green デプロイの場合は、デプロイ設定（トラフィックの移行方法の指定）
  * タスク定義
  * タスク数
  * タグ（マネージドタグの有効化、タグの追加・削除）
  * ... など

タスク配置戦略

* [AZ Balanced Spread (AZ バランススプレッド)] - アベイラビリティーゾーン間およびアベイラビリティーゾーン内のコンテナインスタンス間でタスクを分散します
* [AZ Balanced BinPack (AZ バランスビンパック)] - 利用可能な最小メモリでアベイラビリティーゾーン間およびコンテナインスタンス間でタスクを分散します
* [BinPack (ビンパック)] - CPU またはメモリの最小利用可能量に基づいてタスクを配置します
* [One Task Per Host (ホストごとに 1 つのタスク)] - 各コンテナインスタンスのサービスから最大 1 タスクを配置します
* [カスタム] - 独自のタスク配置戦略を定義します。設定ドキュメントの例については、「Amazon ECS タスクの配置」を参照してください

ネットワークモード
* **EC2 起動タイプの場合、awsvpc ネットワークモードはパブリック IP アドレスを使用する ENI を提供しない**。よって、NAT ゲートウェイなどを用意する必要あり。

ELB
* 動的なポートマッピングにより、単一のコンテナインスタンス上で複数のポートを使用可能

サービスの auto scaling
* ターゲット追跡ポリシーでは以下を設定可能。
  * ECSServiceAverageCPUUtilization—サービスの CPU 平均使用率
  * ECSServiceAverageMemoryUtilization—サービスのメモリ平均使用率
  * ALBRequestCountPerTarget—Application Load Balancer ターゲットグループ内のターゲットごとに完了したリクエストの数


[コンソールを使用したサービスの更新](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/update-service-console-v2.html)

* [新しいデプロイの強制] により、サービス設定を維持したままタスクを置き換えられる
* maximum percent パラメータは  PENDING, RUNNING, STOPPING ステータスのタスク数をカウント
* 置き換えるときは、ELB から登録解除し drain 完了を待つ。その後、 SIGTERM が送信され SIGKILL が送信される
* 置き換えの動作
  * 置き換えタスクが HEALTHY となった後に異常のあるタスクを停止する
  * 置き換えタスクが UNHEALTHY の場合は、置き換えタスクもしくは既存の異常タスクのいずれかを停止する
  * `maximumPercent` が 100 以下の場合は、異常タスクをランダムに 1 個停止してから置き換えタスクを開始
  * ローリングアップデート中はまず異常なタスクを停止する動作


[コンソールを使用してブルー/グリーンデプロイ設定を更新します。](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/update-blue-green-deployment-v2.html)

* CodeDeploy のアプリケーション名、デプロイグループ名、デプロイ設定を変更可能
* CodeDeploy のライフサイクルフックを設定可能


[コンソールでサービスの削除](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/delete-service-v2.html)

* サービスを削除すると、タスク数も自動的に 0 までスケールダウンする


### Deployment Types

[デプロイタイプ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/deployment-types.html)

次の 3 つがある。

* ローリング更新
* CodeDeploy を使用した Blue/Green デプロイ
* 外部デプロイ


[ローリング更新](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/deployment-type-ecs.html)


[デプロイサーキットブレーカー](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/deployment-circuit-breaker.html)

* ロールバックの動作
  * デプロイ失敗時は COMPLETED 状態のデプロイを探す。ロールバック開始すると当該デプロイのステータスが COMPLETED から IN_PROGRESS へと遷移する。つまり、このデプロイは COMPLETED ではない状態になったので、以降のロールバック対象にはならない
  * DescribeServices で状態を確認できる。rolloutState、rolloutStateReason が該当
  * 定常状態にならない場合は FAILED 状態に移行。FAILED 状態のデプロイでは、新しいタスクは起動されない
* `SERVICE_DEPLOYMENT_FAILED` のイベントによりデプロイ失敗を通知できる
* ローリングアップデートのみサポート
* CLB を使用している場合はサポートされない
* 失敗の閾値
  * 閾値は (タスク数 / 2) で計算される。ただし、最小の閾値は 10, 最大は 200
  * 2 段階ある
    * RUNNING 状態の場合は次の段階に遷移。RUNNING にならなかった場合は障害カウントを +1
    * ヘルスチェックを実行。失敗した場合は障害カウントを +1


[CloudWatch アラーム](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/deployment-alarm-failure.html)

* 指定した CloudWatch アラームが ALARM 状態になった場合にロールバックする。以下のようにサービスを作成する
```shell
aws ecs create-service \
     ...
     --deployment-configuration "alarms={alarmNames=[alarm1Name,alarm2Name],enable=true,rollback=true}" \
```
* ALARM 状態になった場合は、デプロイは FAILED になり、新規タスクは起動しなくなる
* アラームを使用するデプロイが失敗した場合にもイベントが送信される。reason フィールドにはロールバックのためのデプロイが開始されたことが示される
* サーキットブレーカー、CloudWatch アラームのどちらかで基準が満たされるとデプロイの失敗となる。その場合は基準を満たした側のロールバック設定を参照し、ロールバックするかどうかを決定する
* ローリングアップデートのみサポート
* ECS 側から `DescribeAlarms` を実行しアラームのステータスを取得する動作となっている。よってスロットリング時に見逃す可能性がある点に注意
* `HTTPCode_ELB_5XX_Count` などのメトリクスを使用するのが手


[CodeDeploy を使用した Blue/Green デプロイ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/deployment-type-bluegreen.html)

* CodeDeply の Blue/Green Deployment の考慮事項
  * デプロイ時に Green のタスクセットを作成する。テストトラフィックを Green のタスクセットに ModifyListener したあと、本番用トラフィックを Blue のタスクセットから Green のタスクセットに ModifyListener する
  * トラフィックの移行は一括、線形、Canaly から選択可能。ただし、NLB では `CodeDeployDefault.ECSAllAtOnce` のみ設定可能
  * CLB はサポートされていない
  * サービスの Auto Scaling と併用でき、デプロイ中もスケーリングできる。しかし、デプロイが失敗する場合がある
  * ECS サービスのサービスロールには CodeDeploy 関連のアクション許可が必要


[外部デプロイ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/deployment-type-external.html)

* 以下の API によって管理される
  * サービス管理（CreateService, UpdateService, DeleteService）
  * タスク設定管理（CreateTaskSet, UpdateTaskSet, UpdateServicePrimaryTaskSet, DeleteTaskSet）
* サポート
  * ALB, NLB をサポート
  * Fargate, DAEMON タイプは未サポート
* サービス定義パラメータ
  * おおむね通常のサービスと同様。scale というパラメータがあり、タスクセットに配置して実行し続けるために必要なタスク数の浮動小数点パーセンテージを指定する。これにより、タスクセットごとに希望タスク数に対して何 % 分を起動するかを指定できる。例えばデプロイ完了時に旧タスクセットに対し 0 を設定するような使い方になる


### Load Balancing

[サービスの負荷分散](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/service-load-balancing.html)

* 特徴
  * 一つのサービスを複数のロードバランサのターゲットグループに登録可能
  * 動的ポートマッピングが可能
  * ALB ではパスベースのルーティングと優先ルールをサポート


[ロードバランサーの種類](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/load-balancer-types.html)

考慮事項
* ターゲットグループは 5 個まで
* Fargate の場合は、ターゲットタイプとして ip を指定する必要がある。
* ELB のサブネットはコンテナインスタンスが存在する全ての AZ を含む必要がある
* サービスの ELB 設定はローリング更新の場合のみ変更可能
* **タスクがヘルスチェックの条件を満たさない場合は、タスクは停止され、再度起動される。**
* NLB と awsvpc の組み合わせの場合、送信元 IP アドレスは NLB のプライベートアドレスとなる。よって、タスク側で NLB のプライベートアドレスを許可するしかないが、その場合は世界中からのアクセス可能な状態になる（NLB 側でセキュリティグループを設定できずフィルタリングできないため）
* NLB の UDP は Linux プラットフォーム 1.4.0、もしくは Windows プラットフォーム 1.0.0 が対応
* 登録解除の遅延よりもタスク定義の stopTimeout を長くすると良い
* ホストポートは一時ポート範囲から動的に選択される


[ロードバランサーの作成](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/create-load-balancer.html)


[サービスに複数のターゲットグループを登録する](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/register-multiple-targetgroups.html)

* 考慮事項
  * ターゲットグループは 5 個まで
  * ALB, NLB をサポート
  * ローリング更新のみサポート


### Service Auto Scaling

[サービスの Auto Scaling](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/service-auto-scaling.html)

* 次の 3 つがサポートされている。
  * ターゲット追跡スケーリングポリシー
  * ステップスケーリングポリシー
  * スケジュールに基づくスケーリング
* サービスにリンクされたロール `AWSServiceRoleForApplicationAutoScaling_ECSService` が Application Auto Scaling 側で作成され使用される
* 考慮事項
  * デプロイ中はスケールインしないようになっている。スケールアウトはされる
  * スケールアウトを中断する場合は `register-scalable-target` を使用する。デプロイ後に `register-scalable-target` を実行し再開されるのを忘れないように
  * クールダウン期間をサポート。この期間が過ぎるまでは必要な容量を再度増加しない
  * 容量を超えた状態でスケールインすると Max の値まで一気にスケールインされる。容量の最小値よりも少ない状態ですケースアウトすると Mix の値まで一気にスケールアウトされる


[ターゲット追跡スケーリングポリシー](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/service-autoscaling-targettracking.html)

* 考慮事項
  * メトリクスデータが不十分な場合はスケールインしない
  * 複数のスケーリングポリシーがある場合、スケールアウトはどれか 1 つでも条件を満たすと行われるが、スケールインは全てが条件を満たす必要がある
  * 自動的に作成された CloudWatch アラームは編集、削除をしてはならない
  * Blue/Green デプロイでは `ALBRequestCountPerTarget` はサポートされない


[ステップスケーリングポリシー](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/service-autoscaling-stepscaling.html)



### Interconnecting

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
* `HealthCheckCustomConfig` を設定することで、コンテナレベルのヘルスチェック結果が CloudMap のカスタムヘルスチェック API に送信される
* ECS サービスを新規作成する場合にのみ設定可能


### Task Scale-In Protection

[タスクスケールイン保護](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-scale-in-protection.html)

* タスクをスケールインから保護するよう設定できる
* `$ECS_AGENT_URI/task-protection/v1/state` にて `ProtectionEnabled` 属性を設定できる。保護期間は `expiresInMinutes` で設定できデフォルトは 2 時間。最短 1 分、最長で 48 時間
* 考慮事項
  * ローリングアップデート時に `protectionEnabled` がクリアされるか有効期限が失効するまで当該タスクは停止しない
  * Blue/Green デプロイでは `protectionEnabled` が設定されたタスクがある場合はクリアされるか有効期限が失効するまで Blue 側のタスクが残る
* IAM
  * タスクロールにて設定が必要
    * ecs:GetTaskProtection: 設定内容の取得
    * ecs:UpdateTaskProtection: 設定内容の更新

[タスクスケールインプロテクションのエンドポイント](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-scale-in-protection-endpoint.html)

* `${ECS_AGENT_URI}/task-protection/v1/state` に対する PUT リクエストによりスケールインからの保護を設定できる
* 60 分間保護する例
```
curl --request PUT --header 'Content-Type: application/json' ${ECS_AGENT_URI}/task-protection/v1/state --data '{"ProtectionEnabled":true}'      

curl --request PUT --header 'Content-Type: application/json' ${ECS_AGENT_URI}/task-protection/v1/state --data '{"ProtectionEnabled":true,"ExpiresInMinutes":60}'
```


### Service Throttle Logic

[サービスの調整ロジック](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/service-throttle-logic.html)

サービスタスクが繰り返し起動に失敗した場合にタスクを起動する頻度を調整するロジックがある。繰り返しタスクの起動が失敗する場合、その後の再起動の試行間隔は最大 15 分まで段階的に増加する。


## Resource tagging

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


[Amazon ECS のサービスクォータ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/service-quotas.html)

* リソース作成数の上限のほか、1 分あたりに起動できるタスク数も定められている
* リファレンス
  * [AWS Fargate スロットリングのクォータ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/userguide/throttling.html)
  * [Request throttling for the Amazon ECS API](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/request-throttling.html)


[AWS Fargate で使用する Amazon ECS でサポートされているリージョン](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/AWS_Fargate-Regions.html)


[Amazon ECS 使用状況レポート](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/usage-reports.html)

* Cost Explorer によりコスト、使用状況のグラフを表示できる



## Monitoring

[Amazon ECS のモニタリング](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs_monitoring.html)

* モニタリング計画を立てる必要がある


[モニタリングツール](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/monitoring-automated-manual.html)

* 自動モニタリングツール
  * CloudWatch Alarm
  * CloudWatch Logs
  * EventBridge
  * CloudTrail
* 手動モニタリングツール
  * Trusted Advisor
  * Compute Optimizer


### CloudWatch Metrics

[Amazon ECS CloudWatch メトリクス](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/cloudwatch-metrics.html)

* EC2 の場合、インスタンスロールに `ecs:StartTelemetrySession` が必要`
* コンテナエージェント設定で `ECS_DISABLE_METRICS` を `true` にしている場合はメトリクスは収集されない
* 名前空間: AWS/ECS
  * CPUReservation:
    * ディメンション: `ClusterName` のみ。(CPU 予約量 ÷ コンテナインスタンスのリソース量)。EC2 起動タイプのみ
  * CPUUtilization:
    * ディメンション: `ClusterName`。(CPU 使用量 ÷ コンテナインスタンスのリソース量)。ACTIVE, DRAINIG のコンテナインスタンス分が集計される。EC2 起動タイプのみ
    * ディメンション: `ClusterName`、`ServiceName`。(サービスに属している CPU ユニット数の使用量 ÷ サービスの CPU 予約量)。ACTIVE, DRAINIG のコンテナインスタンス分が集計される。EC2、Fargate 両方に対応
  * MemoryReservation: CPUReservation の Memory 版
  * MemoryUtilization: CPUUtilization の Memory 版
* クラスター予約
  * コンテナ定義の CPU, Memory, GPU の各コンテナの合計量が予約され、クラスターに登録されているリソース量の割合として計算される
  * タスク定義で `memoryReservation` が指定されている場合はそちらを計算に使用する。ない場合は `memory` を使用
  * メモリの合計サイズには tmpfs, sharedMemorySize のボリュームサイズも含まれている
* クラスター使用率
  * コンテナインスタンスのリソース量に対する、リソース使用量の割合として計算される
* サービス使用率
  * タスク定義で `memoryReservation` が指定されている場合はそちらを計算に使用する。ない場合は `memory` を使用
  * 例えば 512 CPU ユニットを指定している場合、インスタンス上で 1 タスクのみ稼働している場合は 512 CPU ユニットを超えて実行できる。そのため CPUUtilization のメトリクスは 100 % を超えることがある
  * メモリは `memory` を超えて使用することはできない。しかし、`memoryReservation` を超えた使用は可能


### Events

[Amazon ECS イベントおよびEventBridge](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/cloudwatch_event_stream.html)


[ECS のイベント](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs_cwe_events.html)

* コンテナインスタンスの状態変更イベント
  * イベントの契機
    * RunTask などの API によりタスクを起動した時
    * ECS サービススケジューラがタスクを起動、停止した時
    * コンテナエージェントにより `SubmitTaskStateChange` が実行された時
    * `DeregisterContainerInstance` によるコンテナインスタンスの登録解除
    * EC2 インスタンス停止によるタスクの停止時
    * コンテナインスタンスとして登録される時
    * コンテナエージェントが接続、切断するとき
    * コンテナエージェントのバージョンを更新するとき
* タスク状態変更イベント
  * イベントの契機
    * RunTask などの API によりタスクを起動した時
    * ECS サービススケジューラがタスクを起動、停止した時
    * コンテナエージェントにより `SubmitTaskStateChange` が実行された時
    * `DeregisterContainerInstance` によるコンテナインスタンスの登録解除
    * 基盤となるコンテナインスタンスの停止時
    * タスク内のコンテナのステータスが変わった時
    * Fargate Spot の中断通知を受け取った時
* サービスアクションイベント
  * INFO
    * SERVICE_STEADY_STATE
    * TASKSET_STEADY_STATE
    * CAPACITY_PROVIDER_STEADY_STATE
    * SERVICE_DESIRED_COUNT_UPDATED
  * WARN
    * SERVICE_TASK_START_IMPAIRED
    * SERVICE_DISCOVERY_INSTANCE_UNHEALTHY
  * ERROR
    * SERVICE_DAEMON_PLACEMENT_CONSTRAINT_VIOLATED
    * ECS_OPERATION_THROTTLED
    * SERVICE_DISCOVERY_OPERATION_THROTTLED
    * SERVICE_TASK_PLACEMENT_FAILURE
    * SERVICE_TASK_CONFIGURATION_FAILURE
* サービス展開状態変更イベント
  * INFO
    * SERVICE_DEPLOYMENT_IN_PROGRESS
    * SERVICE_DEPLOYMENT_COMPLETED
  * ERROR
    * SERVICE_DEPLOYMENT_FAILED
* コンテナインスタンス状態イベント、タスク状態変更イベントでは detail オブジェクトの version フィールドにてイベントを識別可能。リソースの状態が変わるたびにインクリメントされる


[イベントの処理](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs_cwet_handling.html)

* イベントは少なくとも 1 回は送信される。つまり複数回送信される可能性がある
* Lambda をトリガーして DynamoDB に保存するサンプルスクリプトがある


### Container Insights

[Container Insights](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/cloudwatch-container-insights.html)

* containerInsights アカウント設定をオプトインして作成されたすべての新しいクラスターに対してデフォルトで有効化されている
* 運用データは、パフォーマンスログイベントとして収集される。JSON スキーマのエントリとなっている。CloudWatch はこのデータから CloudWatch メトリクスを作成する
* ネットワークメトリクスは EC2 の場合はネットワークモード none, host では採取されない

* リファレンス
  * [Container Insights のメトリクス](https://docs.aws.amazon.com/ja_jp/AmazonCloudWatch/latest/monitoring/Container-Insights-metrics-ECS.html)
  * [パフォーマンスログイベントの例](https://docs.aws.amazon.com/ja_jp/AmazonCloudWatch/latest/monitoring/Container-Insights-reference-performance-logs-ECS.html)


[CloudWatch Container Insights を使用して、Amazon ECS ライフサイクルイベントを表示するには](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/lifecycle-metrics.html)


### Container Instance Health

[コンテナインスタンスのヘルス](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/container-instance-health.html)

* コンテナインスタンスのヘルスステータスはコンテナエージェントによって実施される
* 1 分間に 2 回実施
* チェックが失敗すると IMPAIRED になる

コンテナインスタンスのヘルスステータスは `DescribeContainerInstances` によって取得可能。


### Trace Data

[アプリケーショントレースデータの収集](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/trace-data.html)

* OpenTelemetry サイドカーコンテナ用の AWS Distroを使用してトレースデータを収集し、AWS X-Ray にルーティングできる

* リファレンス
  * [Setting up AWS Distro for OpenTelemetry Collector in Amazon Elastic Container Service](https://aws-otel.github.io/docs/setup/ecs)


### Application Metrics

[アプリケーションメトリクスを収集する](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/metrics-data.html)

アプリケーションのメトリクスを OpenTelemetry サイドカーコンテナ用 AWS Distro を使用して CloudWatch または Amazon Managed Service for Prometheus へルーティングすることが可能。


[アプリケーションメトリクスを Amazon CloudWatch にエクスポートする](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/application-metrics-cloudwatch.html)

* カスタムメトリクスをエクスポートできる
* OpenTelemetry のコンテナをサイドカーで動作させることで実現できる
* ロググループ `/aws/ecs/application/metrics` にエクスポートされ、メトリクスは名前空間 `ECS/AWSOTel/Application` で表示できる
* OpenTelemetry SDK を使用して計測する必要がある


[アプリケーションメトリクスを Amazon Managed Service for Prometheus にエクスポートする](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/application-metrics-prometheus.html)

* カスタムアプリケーションメトリクスを Amazon Managed Service for Prometheus へエクスポートすることをサポート
* Amazon Managed Grafana ダッシュボードを使用して閲覧可
* OpenTelemetry のコンテナをサイドカーで動作させることで実現できる


### CloudTrail

[AWS CloudTrail を使用した Amazon ECS API コールのログ記録](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/logging-using-cloudtrail.html)

公開 API は CloufTrail に記録される


### Container Agent

[Amazon ECS コンテナエージェントの詳細分析](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-agent-introspection.html)

* メタデータを表示
```
curl -s http://localhost:51678/v1/metadata | python -mjson.tool
```


### Compute Optimizer

[AWS Compute Optimizer 推奨事項](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-recommendations.html)

* タスクの CPU とメモリサイズおよびコンテナの CPU、コンテナメモリ、およびその予約サイズに関する推奨事項が得られる
* リファレンス
  * [Viewing recommendations for Amazon ECS services on Fargate](https://docs.aws.amazon.com/ja_jp/compute-optimizer/latest/ug/view-ecs-recommendations.html)



## Security

[Amazon Elastic Container Service のセキュリティ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/security.html)


### IAM

[Amazon Elastic コンテナサービスのアイデンティティ とアクセス管理](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/security-iam.html)


[IAM を使用するAmazon Elastic Container Service](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/security_iam_service-with-iam.html)

* Condition 要素で使用できる条件キー
  * aws:RequestTag/${TagKey} - リクエストで渡すことができるタグのキーバリューのペアを制御
  * aws:ResourceTag/${TagKey} - リソースにアタッチされたタグに基づいてアクセス制御
  * aws:TagKeys
  * ecs:ResourceTag/${TagKey}
  * ecs:cluster
  * ecs:container-instances
  * ecs:container-name
  * ecs:enable-execute-command
  * ecs:enable-service-connect
  * ecs:namespace
  * ecs:service
  * ecs:task-definition
  * ecs:account-setting


[Amazon Elastic Container Service のアイデンティティベースのポリシーの例](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/security_iam_id-based-policy-examples.html)

* ベストプラクティス
  * 最小権限
  * ポリシーに条件を追加し、更に制限する
  * IAM Access Analyzer を使用して IAM ポリシーを検証
  * MFA を使用


[Amazon Elastic Container Service に関する AWS 管理ポリシー](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/security-iam-awsmanpol.html)

* AmazonECS_FullAccess
* AmazonEC2ContainerServiceforEC2Role: コンテナインスタンス用。bridge の場合はコンテナからアクセスできないようにするには iptables の設定が必要。ただし、iptables では host ネットワークモードだと遮断できない
* AmazonEC2ContainerServiceEventsRole: EventBridge 用
* AmazonECSTaskExecutionRolePolicy: タスク実行ロール用
* AmazonECSServiceRolePolicy: サービスロール用
* AWSApplicationAutoscalingECSServicePolicy: Application Auto Scaling のサービスにリンクされたロールにアタッチされる
* AWSCodeDeployRoleForECS: CodeDeploy のサービスにリンクされたロールにアタッチされる
* AWSCodeDeployRoleForECSLimited: CodeDeploy のサービスにリンクされたロールにアタッチされる


[サービスにリンクされたロール](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/using-service-linked-roles.html)

* `AWSServiceRoleForECS` という名前
* Amazon ECS がユーザーに代わって AWS API を呼び出すために標準的に使用するロール。
* 以下のような用途
  * ENI のライフサイクル管理
  * ELB への登録、登録解除
  * Service Discovery への登録、登録解除
  * Auto Scaling
* 手動で作成する必要はなく、ECS クラスターを作成したタイミングなどで自動的に作成される
* 当該ロールを使用するリソースがある場合は、当該ロールを削除できない


[Amazon ECS コンソールに必要なアクセス許可](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/console-permissions.html)


[タスク実行ロール](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task_execution_IAM_role.html)

* コンテナエージェント用のロール
* ECR からのイメージのプル、CloudWatch Logs へのログ送信などで使用
* シークレットを使用する場合は、SSM, Secrets Manager などへの許可が必要。カスタマーマネージドキーで暗号化されている場合は KMS の権限も必要
* インターフェイスエンドポイントへのアクセスを制限する場合は `aws:SourceVpc` を使用


[タスクロール](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-iam-roles.html)

* EC２ インスタンスのインスタンスプロファイルのようなもの。タスク内で使用できる権限を設定したロール。
* タスク定義で設定するか、RunTask API で上書き可能
* コンテナエージェントにより `AWS_CONTAINER_CREDENTIALS_RELATIVE_URI` 環境変数が各コンテナに設定される
* 認証情報プロバイダーを使用するたびに `/var/log/ecs/audit.log` に記録される
* インスタンスプロファイルへのアクセスを禁止するには以下の対応が必要
  * awsvpc: `ECS_AWSVPC_BLOCK_IMDS` を `true` に設定
  * bridge: iptables により drop。```sudo yum install -y iptables-services; sudo iptables --insert DOCKER-USER 1 --in-interface docker+ --destination 169.254.169.254/32 --jump DROP```
* EC2, 外部インスタンスで必要な対応
  * コンテナエージェントは host ネットワークモードで稼働
  * bridge, default の場合は `ECS_ENABLE_TASK_IAM_ROLE=true`
  * host の場合は `ECS_ENABLE_TASK_IAM_ROLE_NETWORK_HOST=true`
  * コンテナエージェントに転送する設定
   ```
   sudo sysctl -w net.ipv4.conf.all.route_localnet=1
   sudo iptables -t nat -A PREROUTING -p tcp -d 169.254.170.2 --dport 80 -j DNAT --to-destination 127.0.0.1:51679
   sudo iptables -t nat -A OUTPUT -d 169.254.170.2 -p tcp -m tcp --dport 80 -j REDIRECT --to-ports 51679
   ```


[タスク用の Windows IAM ロールの追加設定](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/windows_task_IAM_roles.html)

* コンテナエージェント起動時に `-EnableTaskIAMRole` オプションが必要
* タスク認証情報プロバイダーの IAM ロールは、コンテナインスタンスのポート 80 を使用。よって、コンテナはポート 80 を使用できなくなる
* Powershell の場合は以下の例のようなスクリプトでセットアップが必要
```
$gateway = (Get-NetRoute | Where { $_.DestinationPrefix -eq '0.0.0.0/0' } | Sort-Object RouteMetric | Select NextHop).NextHop
$ifIndex = (Get-NetAdapter -InterfaceDescription "Hyper-V Virtual Ethernet*" | Sort-Object | Select ifIndex).ifIndex
New-NetRoute -DestinationPrefix 169.254.170.2/32 -InterfaceIndex $ifIndex -NextHop $gateway -PolicyStore ActiveStore # credentials API
New-NetRoute -DestinationPrefix 169.254.169.254/32 -InterfaceIndex $ifIndex -NextHop $gateway -PolicyStore ActiveStore # metadata API
```


[コンテナインスタンスの IAM ロール](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/instance_IAM_role.html)

* コンテナインスタンスの登録などの用途
* 管理ポリシーとして mazonEC2ContainerServiceforEC2Role が提供されている


[ECS Anywhere IAM ロール](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/iam-role-ecsanywhere.html)

* 以下のポリシーがあるとよい
  * AmazonEC2ContainerServiceforEC2Role
  * AmazonSSMManagedInstanceCore


[Amazon ECS CodeDeploy IAM ロール](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/codedeploy_IAM_role.html)

* 以下の管理ポリシーが用意されている
  * AWSCodeDeployRoleForECS
  * AWSCodeDeployRoleForECSLimited: より制限されたポリシー
* タスクロール、タスク実行ロールの上書きを行う場合は `iam:PassRole` の許可が必要


[Amazon ECS CloudWatch Events IAM ロール](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/CWE_IAM_role.html)

* 以下の管理ポリシーが用意されている
  * AmazonEC2ContainerServiceEventsRole
* タスクロール、タスク実行ロールの上書きを行う場合は `iam:PassRole` の許可が必要


[リソース作成時にタグ付けするための許可を付与する](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/supported-iam-actions-tagging.html)

* タグづけには `ecs:TagResource` の許可が必要


[リソースタグを使用して Amazon ECS リソースへのアクセスを制御する](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/control-access-with-tags.html)

* 以下の例のようなポリシーにより、当該タグが付与されたリソースに対してのみアクションが許可もしくは拒否される
```
"StringEquals": { "aws:ResourceTag/environment": "production" }
```


[ポリシーの例](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/iam-policies-ecs-console.html)


[Amazon Elastic Container Service のアイデンティティとアクセスのトラブルシューティング](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/security_iam_troubleshoot.html)

* 既存のロールをサービスに渡す場合は `iam:PassRole` が必要


### Logging, Monitoring

[Amazon Elastic Container Service でのログとモニタリング](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-logging-monitoring.html)

* 以下のサービスが提供されている
  * CloudWatch Alarm
  * CloudWatch Logs
  * EventBridge
  * CloudTrail
  * Trusted Advisor
  * Compute Optimizer


### Compliance

[Amazon Elastic Container Service でのコンプライアンス検証](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-compliance.html)

* AWS Artifact を使用して、サードパーティーの監査レポートをダウンロードできる
* コンプライアンスに役立つドキュメント、AWS サービスが提供されている
  * AWS Config
  * AWS Security Hub
  * AWS Audit Manager


[AWS Fargate 連邦情報処理標準 (FIPS-140)](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-fips-compliance.html)

* FIPS-140 は、機密情報を保護する暗号モジュールのセキュリティ要件を規定する米国およびカナダ政府の標準
* 転送中のデータと保管中のデータの暗号化に使用できる、一連の検証済みの暗号化関数を定めている
* FIPS-140 コンプライアンスをオンにすると、FIPS-140 に準拠した態様で Fargate でワークロードを実行できる
* AWS GovCloud (US) リージョンでのみ利用可能
* 制約
  * operatingSystemFamily は LINUX のみ
  * cpuArchitecture は X86_64 のみ
  * プラットフォームバージョンは 1.4.0 以降
* コンテナ内で `cat /proc/sys/crypto/fips_enabled` で「1」になっている場合は FIPS が使用されている


### Infrastructure Security

[Amazon Elastic Container Service におけるインフラストラクチャセキュリティ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/infrastructure-security.html)


[Amazon ECS インターフェイス VPC エンドポイント (AWS PrivateLink)](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/vpc-endpoints.html)

* 必要なエンドポイント
  * Fargate
    * com.amazonaws.region.ecr.dkr
    * com.amazonaws.region.ecr.api (プラットフォームバージョン 1.3.0 では不要)
    * S3 のエンドポイント
    * 場合によって必要なエンドポイント
      * Secrets Manager
      * logs
      * ecs-agent (Service Connect 使用時。VPC エンドポイント未使用時は ecs-sc エンドポイントが使用される)
  * コンテナインスタンス用
    * com.amazonaws.region.ecs-agent
    * com.amazonaws.region.ecs-telemetry
    * com.amazonaws.region.ecs


[Amazon ECR エンドポイントとクォータ](https://docs.aws.amazon.com/ja_jp/general/latest/gr/ecr.html)

* ecr, api.ecr
  * `CreateRepository`, `DescribeImages` などに使用される
* Docker および OCI クライアントエンドポイント
  * Docker の pull, push などに使用されるエンドポイント


[Amazon ECR Public エンドポイントとクォータ](https://docs.aws.amazon.com/ja_jp/general/latest/gr/ecr-public.html)

* ecr-public, api.ecr-public
  * `CreateRepository`, `DescribeImages` などに使用される


[Amazon ECS エンドポイントとクォータ](https://docs.aws.amazon.com/ja_jp/general/latest/gr/ecs-service.html)

* ecs
* ecs-sc
  * Service Connect 用のエンドポイント



### Best Practices

[セキュリティのベストプラクティス](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/security-bestpractices.html)


[AWS Identity and Access Management](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/security-iam-bestpractices.html)

* 最小権限アクセスのポリシーに従う
* クラスターリソースを管理上の境界として使用
* 自動パイプラインを作成し、手動での変更、アクセスを制限する
* ポリシー条件を使用
* Amazon ECS API へのアクセスを定期的に監査


[Amazon ECS タスクで IAM ロールを使用する](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/security-iam-roles.html)

* タスクロールの使用を推奨
  * コンテナエージェントはコンテナ内の環境変数 `AWS_CONTAINER_CREDENTIALS_RELATIVE_URI` に認証情報 ID の URI を設定する仕組み
  * 有効期限は 6 時間で、コンテナエージェントによって自動的にローテーションされる
* タスク実行ロール
  * イメージの取得先は ECR にしたほうがレート制限にかかりにくい
  * Fargate の場合はプライベートリポジトリはタスク定義のコンテナ定義内で `repositoryCredentials` により認証が必要
* 推奨事項
  * Amazon EC2 メタデータへのアクセスをブロックする
    * awsvpc: `ECS_AWSVPC_BLOCK_IMDS` を `true` に設定
    * bridge: iptables
    * host: `ECS_ENABLE_TASK_IAM_ROLE_NETWORK_HOST` を `true` に設定
  * awsvpc ネットワークモードを使用: タスクレベルでネットワークを分離し、セキュリティグループも設定できるため
  * IAM アクセスアドバイザーを使用してロールを絞り込む
  * AWS CloudTrail に不審なアクティビティがないか監視


[ネットワークセキュリティ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/security-network.html)

* 転送中の暗号化
  * App Mesh だと Envoy プロキシ間は TLS で接続される
  * Nitro インスタンス間の通信は暗号化される
  * ALB で SNI を使用する
  * TLS で End-to-end を暗号化
* タスクネットワーク
  * awsvpc を推奨。セキュリティグループをアタッチできるため
* AWS PrivateLink
  * エンドポイントポリシーを設定できる
* Amazon ECS コンテナエージェントの設定
  * ECS_AWSVPC_BLOCK_IMDS
  * ECS_ENABLE_TASK_IAM_ROLE_NETWORK_HOST
* 推奨事項
  * ネットワークの伝送路を暗号化
  * awsvpc でセキュリティグループの使用
  * ネットワークの分離が必要な場合は、クラスターごとに VPC を分ける
  * AWS PrivateLink エンドポイントを使用
  * VPC フローログの取得と分析


[シークレットの管理](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/security-secrets-management.html)

* 以下の箇所でシークレットを使用できる
  * コンテナの環境変数
  * ログ設定
  * コンテナレジストリの認証
* 推奨事項
  * シークレット情報は SSM Parameter Store, Secrets Manager の使用を検討する。タスク実行ロールに権限が必要権限が必要
  * 環境変数は docker inspect で読み取られるリスクがある。S3 からシークレット情報の記載されたオブジェクトをダウンロードするのも手
  * サイドカーコンテナにてボリュームにシークレット情報を書き出し、アプリケーションコンテナにてボリュームをマウントする


[API オペレーションで一時的なセキュリティ認証情報を使用する](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/temp-credientials.html)

* sts から取得した一時的な認証情報で API 実行できる


[コンプライアンスとセキュリティについて](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/security-compliance.html)

* 決済カード業界のデータセキュリティ基準 (PCI DSS)
  * [Architecting on Amazon ECS for PCI DSS Compliance](https://d1.awsstatic.com/whitepapers/compliance/architecting-on-amazon-ecs-for-pci-dss-compliance.pdf)
* HIPAA (米国の医療保険の相互運用性と説明責任に関する法令)
* 推奨事項
  * 早めに社内のコンプライアンスプログラムの所有者に働きかけ、AWS 責任共有モデルを使用してコンプライアンス統制の所有権を特定する必要がある


[ログ記録とモニタリング](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/security-logging-and-monitoring.html)

* bridge ネットワークモードの場合、アプリケーションコンテナの前に FireLens 設定のコンテナが起動する必要がある


[AWS Fargate のセキュリティ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/security-fargate.html)

* 2020年5月28日以降に起動されるタスクでは、エフェメラルストレージは AWS Fargate 管理の暗号化キーを使用して AES-256 暗号化アルゴリズムで暗号化されている
* Fargate では SYS_PTRACE の capability のみ追加可能 


[AWS Fargate のセキュリティに関する考慮事項](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/fargate-security-considerations.html)

* 個々の Fargate タスクの環境は独立している
* 特権コンテナは使用不可
* capability の多くが制限されている
* 基盤となるホストにはアクセスできない


[EC2 コンテナインスタンスのセキュリティに関する考慮事項](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ec2-security-considerations.html)

* タスクロールは最小権限にする
* awsvpc の場合は `ECS_AWSVPC_BLOCK_IMDS` を有効化


[タスクとコンテナのセキュリティ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/security-tasks-containers.html)

* 推奨事項
  * イメージの最小化、distroless イメージの使用
  * イメージのスキャン
  * 特別な権限の削除。setuid, setgid など
  * キュレーションされたイメージのセットを用意しておく
  * アプリケーションとライブラリをスキャンする
  * 静的コード分析
  * 非 root ユーザーで実行
  * 読み取り専用のルートファイルシステムを使用
  * CPU とメモリの使用量を制限
  * ECR でイミュータブルタグの使用
  * 特権コンテナを使用しない。`ECS_DISABLE_PRIVILEGED` を `true` にすることで対応可能
  * capability を drop する
  * ECR に格納されるイメージの暗号化にカスタマーマネージドキーを使用


[ランタイムセキュリティ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/security-runtime.html)

* Windows コンテナの場合
  * [Secure Windows containers](https://learn.microsoft.com/en-us/virtualization/windowscontainers/manage-containers/container-security)
* Linux コンテナの場合
  * `linuxParameters` により不要な capability を drop する
  * SELinux labels もしくは AppArmor profile を `dockerSecurityOptions` により適用
    * AppArmor は enforcement または complain モードで実行
* 推奨事項
  * ランタイムの防御にサードパーティーソリューションを使用する


[AWS パートナー](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/security-partners.html)



## ECS メタデータ

[Amazon ECS メタデータ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-metadata-endpoint.html)

* 以下のメタデータエンドポイントがある
  * コンテナメタデータファイル: コンテナに関する情報
  * タスクメタデータエンドポイント: タスクのメタデータ、Docker の統計情報


[Amazon ECS コンテナメタデータファイル](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/container-metadata.html)

* メタデータファイルはホストインスタンスで作成され、Docker ボリュームとしてコンテナにマウントされる
* コンテナメタデータファイルは、コンテナがクリーンアップされるときにホストインスタンスでクリーンアップされる
* `ECS_ENABLE_CONTAINER_METADATA` を `true` にすることで有効化できる
* 所定のファイルパスから参照できるようになる
  * Linux
    * ホスト: /var/lib/ecs/data/metadata/cluster_name/task_id/container_name/ecs-container-metadata.json
    * コンテナ: /opt/ecs/metadata/random_ID/ecs-container-metadata.json
  * Windows
    * ホスト: C:\ProgramData\Amazon\ECS\data\metadata\task_id\container_name\ecs-container-metadata.json
    * コンテナ: C:\ProgramData\Amazon\ECS\metadata\random_ID\ecs-container-metadata.json
  * コンテナでは環境変数 `ECS_CONTAINER_METADATA_FILE` にパスが設定される
* タスク ARN、ImageID、IPv4Addresses などの情報が JSON 形式で保存されている
* コンテナ開始後 1 秒くらいで作成される。それまでの間は READY ステータスではなくファイルの内容は不完全な状態となっている


[EC2 のタスクで使用できるタスクメタデータ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ec2-metadata.html)

* メタデータエンドポイントのバージョンは 2 〜 4 まで。新しいバージョンは ECS Agent のバージョンが新しくないと使用できない
* バージョン 3 から Docker 統計も提供される


[タスクメタデータエンドポイントバージョン 4](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-metadata-endpoint-v4.html)

* タスクメタデータとネットワークレートの統計は CloudWatch に送信され、 AWS Management Console で表示できる
* コンテナには環境変数 `ECS_CONTAINER_METADATA_URI_V4` がセットされる
* 以下のようなエンドポイントパスが用意されている
  * `${ECS_CONTAINER_METADATA_URI_V4}`
  * `${ECS_CONTAINER_METADATA_URI_V4}/task`
  * `${ECS_CONTAINER_METADATA_URI_V4}/taskWithTags`
  * `${ECS_CONTAINER_METADATA_URI_V4}/stats`: awsvpc, bridge の場合は追加のネットワーク統計が含まれる
  * `${ECS_CONTAINER_METADATA_URI_V4}/task/stats`
* Docker 統計は ```${ECS_CONTAINER_METADATA_URI_V4}/stats```。[Docker Stats](https://docs.docker.com/engine/api/v1.30/#operation/ContainerStats) 参照のこと


[タスクメタデータエンドポイントバージョン 3](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-metadata-endpoint-v3.html)

* 現在アクティブにメンテナンスされていない


[タスクメタデータエンドポイントバージョン 2](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-metadata-endpoint-v2.html)

* 現在アクティブにメンテナンスされていない


[Fargate のタスクで使用できるタスクメタデータ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/fargate-metadata.html)

* タスクメタデータエンドポイントバージョン 4: プラットフォームバージョン 1.4.0 以降を使用するタスクで利用可能
* タスクメタデータエンドポイントバージョン 3: プラットフォームバージョン 1.1.0 以降を使用するタスクで利用可能


[Fargate のタスク用のタスクメタデータエンドポイントバージョン 4](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-metadata-endpoint-v4-fargate.html)

* コンテナには環境変数 `ECS_CONTAINER_METADATA_URI_V4` がセットされる
* 以下のようなエンドポイントパスが用意されている
  * `${ECS_CONTAINER_METADATA_URI_V4}`
  * `${ECS_CONTAINER_METADATA_URI_V4}/task`
  * `${ECS_CONTAINER_METADATA_URI_V4}/stats`: awsvpc, bridge の場合は追加のネットワーク統計が含まれる
  * `${ECS_CONTAINER_METADATA_URI_V4}/task/stats`


[Fargate のタスク用のタスクメタデータエンドポイントバージョン 3](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-metadata-endpoint-v3-fargate.html)



## ECS Integrations

[Amazon ECS と統合された AWS のサービス](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-integrations.html)


[Amazon ECR を Amazon ECS で使用](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecr-repositories.html)


[Amazon ECS クラスターのLocal Zones、Wavelength Zone、およびAWS Outposts](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/cluster-regions-zones.html)


[AWS OutpostsのAmazon Elastic Container Service](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-on-outposts.html)


[AWSAmazon ECS のDeep Learning Containers](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/deep-learning-containers.html)

* Deep Learning Containersは、TensorFlow、NVIDIA CUDA(GPUインスタンス用)、Intel MKL(CPUインスタンス用)のライブラリを使用し、最適化された環境を実現
* Amazon Elastic Inference を併用することで、推論コストを削減できる


[Amazon ECS での AWS ユーザー通知の使用](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/using-user-notifications.html)

* EventBridge にて各イベントのトリガーをもとにターゲットの処理を実行できる



## チュートリアル

[チュートリアル: AWS CLI を使用して Fargate Linux タスクでクラスターの作成](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ECS_AWSCLI_Fargate.html)


[チュートリアル: AWS CLI を使用して Fargate Windows タスクでクラスターを作成](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ECS_AWSCLI_Fargate_windows.html)


[チュートリアル: AWS CLI を使用して EC2 タスクのクラスターを作成する](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ECS_AWSCLI_EC2.html)


[チュートリアル: AWS Management Console と Amazon ECS コンソールによるクラスターの自動スケーリングの使用](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/tutorial-cluster-auto-scaling-console.html)


[チュートリアル: Secrets Manager のシークレットを使用して機密データを指定する](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/specifying-sensitive-data-tutorial.html)

* タスク定義では containerDefinitions.secrets にて valueFrom により指定
* タスク実行ロールに Secrets Manager の権限が必要


[チュートリアル:サービスディスカバリを使用して、サービスの作成](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/create-service-discovery.html)

* サービスディスカバリリソースの作成
  * [名前空間] - [サービス] の構造。ECS 側でサービスディスカバリを有効化することで、タスク起動時にサービス内にサービスディスカバリインスタンスとして登録される
  * Route 53 にホストゾーンが作成され、サービスディスカバリインスタンスに対応した A レコードが登録される仕組み


[チュートリアル: Blue/Green デプロイを使用するサービスの作成](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/create-blue-green.html)

* ALB もしくは NLB に対応
* サービス作成時に deploymentController を CODE_DEPLOY にする
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


[チュートリアル:タスク停止イベントを Amazon SNS に送信](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs_cwet2.html)

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


[チュートリアル: コンソールを使用した Amazon ECS での Amazon EFS ファイルシステムの使用](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/tutorial-efs-volumes.html)

タスク定義にて volumes により EFS を指定し、コンテナ定義の mountPoints によりマウントポイントを指定
```json
            "mountPoints": [
                {
                    "containerPath": "/usr/share/nginx/html",
                    "sourceVolume": "efs-html"
                }
            ],
            "name": "nginx",
            "image": "nginx"
        }
    ],
    "volumes": [
        {
            "name": "efs-html",
            "efsVolumeConfiguration": {
                "fileSystemId": "fs-1324abcd",
                "transitEncryption": "ENABLED"
            }
        }
    ],
```



[チュートリアル: Amazon ECS で FSx for Windows File Server ファイルシステムを使用](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/tutorial-wfsx-volumes.html)

* Windows Active Directory (AD) を作成する手順となっている
* インスタンスロールに `AmazonSSMDirectoryServiceAccess` が必要
* FSx for Windows File Server を作成する際に AD を指定している
* タスク定義にてマウントの設定が可能
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


[チュートリアル: Windows コンテナ用に Amazon ECS に Fluent Bit をデプロイする](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/tutorial-deploy-fluentbit-on-windows.html)

* ユーザーデータにて以下のような指定を行う
```
<powershell>
Import-Module ECSTools
Initialize-ECSAgent -Cluster cluster-name -EnableTaskENI -EnableTaskIAMRole -LoggingDrivers '["awslogs","fluentd"]'
</powershell>
```
* DAEMON スケジューリング戦略で起動する


[Fargate AWS CLI キャパシティープロバイダーの例](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/fargate-cli-examples.html)

* クラスター作成時にキャパシティープロバイダーを関連づけることができる
```
aws ecs create-cluster \
     --cluster-name FargateCluster \
     --capacity-providers FARGATE FARGATE_SPOT \
     --region us-west-2
```
* `PutClusterCapacityProviders` API により既存のクラスターのキャパシティープロバイダーの関連づけを更新できる
  * 追加する新しいキャパシティープロバイダーに加えて、既存のキャパシティープロバイダーをすべて指定する必要がある
  * 関連づけを解除する場合、既存のタスクが使用していない状況である必要がある
```
aws ecs put-cluster-capacity-providers \
     --cluster FargateCluster \
     --capacity-providers FARGATE FARGATE_SPOT existing_capacity_provider1 existing_capacity_provider2 \
     --default-capacity-provider-strategy existing_default_capacity_provider_strategy \
     --region us-west-2
```


[チュートリアル: AWS CLI を使用するドメインレス gMSA で Windows コンテナを使用する](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/tutorial-gmsa-windows.html)

* タスク定義にて credentialSpecs にて S3 オブジェクトのパスを指定する
* タスク実行ロールにて S3 への読み取り権限が必要
* コンテナエージェント起動前に環境変数 `ECS_GMSA_SUPPORTED` の設定が必要


## トラブルシューティング

[デバッグ用にAmazon ECS Exec を使用](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-exec.html)

* readonlyRootFilesystem が有効化されていると使用できない
* ゾンビプロセスをクリーンアップするにはタスク定義に `initProcessEnabled` フラグ設定を推奨
* タスクロールに SSM の権限が必要（タスク実行ロールではない）
* ログ記録を有効化できる。ECS クラスターで設定する
* アイドルタイムアウトは 20 分
* ECS Exec は一部の CPU, メモリを使用する
* サービス作成時に `--enable-execute-command` フラグが必要
* トラブルシューティングには Amazon ECS Execチェッカーが便利


[ECS Anywhere の問題のトラブルシューティング](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-anywhere-troubleshooting.html)

* うまくいかなかった場合は登録解除してから再度インストールスクリプトを実行すること
* 最もよくある原因はネットワークもしくは IAM


[停止されたタスクでのエラーの確認](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/stopped-task-errors.html)

* 停止理由の確認には以下の情報を使用するとよい
  * タスク停止理由
    * マネジメントコンソールからタスクを表示することでエラーを確認できる
    * 1 時間以内に停止したタスクしか確認できない
    * それ以前を確認するにはタスク状態変更イベントを CloudWatch Logs に送信するような設定を事前にしておく必要がある
  * ECS サービスのイベント
  * コンテナログ
  * (EC2 の場合) コンテナインスタンスのログ
* [停止理由] を確認することでタスクの停止理由が分かる。例えば以下のような停止理由がある。
  * Task failed ELB health checks in (elb elb-name)
  * Scaling activity initiated by (deployment deployment-id)
  * Host EC2 (instance id) stopped/terminated
  * Container instance deregistration forced by user
  * Essential container in task exited
* コンテナの欄を確認することで停止理由が分かる。


[CannotPullContainer task errors](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task_cannot_pull_image.html)

コンテナイメージのプルに失敗したエラー。以下のようなエラー原因が考えられる
* コンテナレジストリとの疎通性がない
  * パブリックサブネットでは自動割り当てパブリック IP を有効にする必要がある。プライベートサブネットの場合は無効にする
  * インターネットへのアウトバウンド疎通性がない場合、S3 や各エンドポイントとの疎通性が必要
* エンドポイントポリシーにて許可されていない
* イメージが格納されていない
* ディスク容量不足
  * `json-file` ログドライバーを使用している場合は容量を圧迫しているかもしれない。当該ログドライバーのオプションとして `max-size` を設定できる
* コンテナレジストリ側のレートリミット


[サービスイベントメッセージ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/service-event-messages.html)

* サービスイベントログ
  * 最新 100 件のイベントが表示される
  * 重複したイベントメッセージは原因が解決されるか 6 時間が経過すると省略される。6 時間以内に解消されなかった場合は、別のサービスイベントメッセージが表示される
  * Auto Scaling イベントには Message のプレフィックスが付与される
* イベントメッセージ例:
  * 定常状態になった場合: `service service-name) has reached a steady state.`
  * リソース不足によりタスクを配置できない
    * CPU, Memory, GPU, ENI
    * ポート番号が埋まっている
    * コンテナインスタンスの属性不足
      * awslogs 使用時にコンテナエージェント設定で有効化していない場合など
  * ECS Agent が接続されていない: 対策は ecs サービスの再起動
  * ELB ヘルスチェックの失敗
  * 一貫してタスクを正常に起動できない: サービスの調整ロジックにより起動間隔が長くなっていく
  * API レート制限
  * `minimumHealthyPercent`, `maximumPercent` の設定不備によりタスクの起動、停止ができない
  * Quota に抵触
    * タスク数、合計 vCPU 数、Memory
  * サポートされていない構成
    * サポートされていないリージョン、AZ
    * Fargate Spot で ARM アーキテクチャを選択
  * Desired Count よりも多くのタスクがスケールインから保護されており、タスクを停止できない
  * 容量不足


[指定された CPU またはメモリの値が無効](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task-cpu-memory-error.html)

* `requiresCompatibilities` が EC2 の場合、CPU ユニットは 128 〜 10240 の間で設定可能（0.125 vCPU ～ 192 vCPU で設定できる）
* `requiresCompatibilities` が Fargate の場合、CPU, Memory に指定できる値のペアが固定されている


[CannotCreateContainerError: API error (500): devmapper](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/CannotCreateContainerError.html)

* Amazon Linux の場合は `/dev/xvdcz` の 22 GiB のボリュームが Docker 用となる
  * サイズを増加したデータボリュームでインスタンスを起動するのが楽
  * Docker が使用するボリュームグループにストレージを追加し、その論理ボリュームを拡張することも可能
* `ECS_ENGINE_TASK_CLEANUP_WAIT_DURATION` でタスク停止後に、コンテナがクリーンアップされるまでの時間を設定可能。デフォルトは 3 時間
* コンテナ内で使用されていないデータブロックを削除するには `fstrim` コマンドを実行


[サービスロードバランサーのトラブルシューティング](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/troubleshoot-service-load-balancers.html)

* タスクが停止する場合の原因の候補
  * サービスにリンクされたロールが作成されていない
  * セキュリティグループの設定不備
  * ELB で設定されていない AZ 上にコンテナインスタンスが作成されている
  * ELB のヘルスチェック設定不備
  * タスク定義のコンテナ名、ポート番号を変更した場合


[サービスの自動スケーリングのトラブルシューティング](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/troubleshoot-service-auto-scaling.html)

* デプロイの進行中
  * スケールインプロセスは自動的にオフにされる
  * スケールアウトプロセスはオフにならないので、必要に応じて中断しておく必要がある


[Docker デバッグ出力の有効化](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/docker-debug-mode.html)

* デバッグ出力の有効化方法
  * `/etc/sysconfig/docker` にて `OPTIONS` に `-D` フラグを追加
  * Docker デーモン、ECS Agent をリスタートし反映


[Amazon ECS ログファイルの場所](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/logs.html)

* /var/log/ecs/ecs-agent.log.timestamp
  * 1 時間ごとにローテーションされ、24 世代分(1 日分)保持される
  * 環境変数
    * `ECS_LOGFILE`: パス変更可能
    * `ECS_LOGLEVEL`: ログレベル。デフォルトは info
    * `ECS_LOGLEVEL_ON_INSTANCE`: インスタンスのログファイルに記録されるログの詳細レベル
    * `ECS_LOG_DRIVER`: ログドライバー。デフォルトは `json-file`
    * `ECS_LOG_ROLLOVER_TYPE`: ローテーション方法の指定。`size` or `hourly`
    * `ECS_LOG_OUTPUT_FORMAT`: ログ出力形式。デフォルトは `logfmt`
    * `ECS_LOG_MAX_FILE_SIZE_MB`: ログファイルの最大サイズ。`hourly` の場合は無視される
    * `ECS_LOG_MAX_ROLL_COUNT`: ログファイルを保存する世代数。デフォルトは 24
* /var/log/ecs/ecs-init.log
  * `ecs-init` プロセスのログ
* /var/log/ecs/audit.log
  * `GetCredentials` API によりクレデンシャルの要求を行ったログ


[Amazon ECS ログコレクター](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/ecs-logs-collector.html)

インスタンス上のログを収集できるツール。サポートへ送付する際などに使用。


[エージェントのイントロスペクション診断](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/introspection-diag.html)

* 以下コマンドによりタスクの状態を取得できる
```
curl http://localhost:51678/v1/tasks | python -mjson.tool
```


[Docker 診断](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/docker-diags.html)

診断用の Docker コマンド
```
$ docker ps -a
$ docker logs コンテナID
$ docker inspect コンテナID
```


[AWS Fargate スロットリングのクォータ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/throttling.html)

* トークンバケットアルゴリズムが使用される
  * バケットサイズは 100。よって、100 個までのタスクを同時起動できる
  * リフィルレートは毎秒 20。よって持続的な起動レードは毎秒 20 個まで
* ECS Fargate, EKS Fargate 分が合算される


[API の失敗の理由](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/api_failures_messages.html)

* 失敗理由
  * Describe*
    * MISSING: 当該リソースが存在していない
  * DescribeTasks
    * TaskFailedToStart: AGENT: エージェントが接続されていない
    * TaskFailedToStart: ATTRIBUTE: 例えば awsvpc ネットワークモードを使用するタスクであるものの、指定したサブネット内に `ecs.capability.task-eni` 属性を持つインスタンスがない
    * TaskFailedToStart: EMPTY CAPACITY PROVIDER: キャパシティープロバイダーが空もしくはキャパシティープロバイダーのインスタンスがクラスターに登録されていない
  * RunTask, StartTask
    * LOCATION: 指定したサブネットがコンテナインスタンスのサブネットと異なる AZ にある
  * UpdateTaskProtection
    * DEPLOYMENT_BLOCKED: 1 つ以上の保護されたタスクが原因でサービスのデプロイが安定した状態にならないため、タスク保護を設定できない
    * TASK_NOT_VALID: 指定されたタスクは ECS サービスの一部ではない


[タスク用の IAM ロールのトラブルシューティング](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/troubleshoot-task-iam-roles.html)

* `Unable to locate credentials` エラーが表示された場合の理由
  * コンテナインスタンスでタスクロールが有効になっていない
  * `ECS_TASK_METADATA_RPS_LIMIT` による流量制限にかかっている


## References

[Task definition parameters](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/task_definition_parameters.html)

* ファミリー
  * family: 複数バージョンのタスク定義の名前のようなもの
* 起動タイプ
  * requiresCompatibilities: EC2 | FARGATE | EXTERNAL
* タスクロール
  * taskRoleArn: Windows ではコンテナインスタンス起動時に `-EnableTaskIAMRole` が必要
* タスク実行ロール
  * executionRoleArn: タスクの実行に使用するロール（ECR, CloudWatch Logs など）
* ネットワークモード
  * networkMode:
    * Linux のデフォルトは bridge。none, bridge, awsvpc, host から選択可能
    * Windows のデフォルトは default。Windows では default, awsvpc のみから選択可能
    * none: 外部接続性を持たない。コンテナ定義でポートマッピングを設定できない
    * bridge: Docker ビルトインの仮想ネットワークを使用
    * host: Docker ビルトインの仮想ネットワークをバイパスし、コンテナポートとホストポートを直接マッピングする
      * ダイナミックポートマッピングは使用不可
      * hostPort の指定が必要
      * ルートユーザーを使用してコンテナを実行してはならない
    * awsvpc: ENI がタスクに割り当てられる
      * Fargate の場合はこのモードになる
      * NetworkConfiguration の指定が必要
    * default: Docker ビルトインの仮想ネットワークを使用。nat Docker ネットワークドライバーが使用される
* ランタイムプラットフォーム
  * operatingSystemFamily
    * Fargate では必須。LINUX, WINDOWS_SERVER_2019_FULL, WINDOWS_SERVER_2019_CORE, WINDOWS_SERVER_2022_FULL, WINDOWS_SERVER_2022_CORE
    * EC2 の場合は LINUX, WINDOWS_SERVER_2022_CORE, WINDOWS_SERVER_2022_FULL, WINDOWS_SERVER_2019_FULL, WINDOWS_SERVER_2019_CORE, WINDOWS_SERVER_2016_FULL, WINDOWS_SERVER_2004_CORE, WINDOWS_SERVER_20H2_CORE
    * サービスの platformFamily 値と一致する必要がある
  * cpuArchitecture: X86_64 or ARM64
    * Fargate Spot, Windows は対応。[考慮事項](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/userguide/ecs-arm64.html#ecs-arm64-considerations)
* タスクサイズ
  * EC2 では省略可能。Fargate では必須
  * Windows の場合は無視されるため、コンテナレベルでの指定を推奨
  * cpu
    * EC2, 外部インスタンスで許可されている値は 0.25 〜 10 vCPU の間
    * Fargate の場合は、表に指定されている CPU, Memory の組のみ指定可能
  * memory
    * ハードリミット
* コンテナ定義
  * name: コンテナ名
  * image
    * repository-url/image:tag または repository-url/image@digest
  * memory
    * ハードリミット
    * Fargate の場合はオプション
    * EC2 の場合:
      * タスクレベルもしくはコンテナレベルのどちらかの指定が必要
      * `memoryReservation` 未指定時は `memory` での指定量をスケジューリング時に使用する
    * Docker デーモン 20.10.0 以降によって、コンテナ用として 6 MiB 以上のメモリが予約されるため、6 MiB 以下は指定しないようにする
  * memoryReservation
    * Windows ではサポートされない
  * portMappings
    * awsvpc では `hostPort` は空もしくは `containerPort` と同じ値にする
    * appProtocol: Service Connect で使用。"HTTP" | "HTTP2" | "GRPC"
    * containerPort
      * Windows では 3150 は指定不可
      * EC2 の場合 `hostPort` 未指定時は動的ポートマッピングになる
    * containerPortRange
      * 動的ポートマッピングのホストポート範囲にバインドされるコンテナのポート番号の範囲
      * bridge, awsvpc のみ。EC2, Fargate の両方に対応
      * 最大 100 個のポートレンジまで指定可能
      * ポートの数が多い場合は、Docker デーモン設定ファイルの docker-proxy をオフにすることを推奨
    * hostPortRange
      * Docker によって割り当てられる
    * hostPort
      * Fargate　の場合は空もしくは `containerPort` と同じ値にする
      * EC2 の場合、未指定時は動的ポートマッピングになる
      * 一時ポート範囲は `/proc/sys/net/ipv4/ip_local_port_range` にリストされている
      * エフェメラルポート範囲では、ホストポートを指定しないこと。自動割り当て用に予約済みのため
      * デフォルトの予約済みポートは、SSH 用の 22、Docker ポートの 2375 および 2376、Amazon ECS コンテナエージェントポートの 51678-51680
      * コンテナインスタンスには、デフォルトの予約済みポートを含めて、一度に最大 100 個の予約済みポートを割り当てられる
      * name: Service Connect で使用
      * protocol: tcp or udp。デフォルトは tcp
  * repositoryCredentials
    * credentialsParameter: プライベートリポジトリの認証情報が含まれているシークレットの Amazon リソースネーム (ARN)
  * healthCheck
    * コンテナイメージに組み込まれた HEALTHCHECK については ECS コンテナエージェントはモニタリングしない
    * コンテナの `healthStatus` は `HEALTHY`, `UNHEALTHY`, `UNKNOWN` のいずれかとなる
    * タスクの `healthStatus`
      * UNHEALTHY — 1 つ以上の必須コンテナのヘルスチェックが失敗
      * UNKNOWN — タスク内で実行されている必須コンテナはすべて UNKNOWN 状態。他の必須コンテナは UNHEALTHY 状態ではない
      * HEALTHY — タスク内のすべての必須コンテナがヘルスチェックを正常に完了
    * スタンドアローンのタスクは `healthStatus` にかかわらず稼働を続行する
    * コンテナエージェントが ECS サービスに接続できない場合は、サービスはコンテナを `UNHEALTHY` として報告する
    * command: stderr が出力されない終了コード 0 は成功。0 以外の終了コードは失敗
    * interval: 間隔
    * timeout: タイムアウト
    * retries: デフォルトは 3 回。1 〜 10 回を指定できる
    * startPeriod: コンテナ起動時にブートストラップのための時間を提供。デフォルトは無効。0 〜 300 秒を指定可能
  * cpu
    * `CpuShares` にマップされる。Agent バージョンが 1.2.0 以上の場合は 0, 1, null は 2 CPU share として渡される
    * Fargate では省略可能
    * Windows では絶対クォータとして強制される。0, null を指定した場合は Docker に 0 として渡されるが 1 CPU の 1 % と解釈される
  * gpu
    * GPU 数を指定
    * Fargate, Windows では未サポート
  * Elastic Inference accelerator
    * value は deviceName と一致
    * Fargate, Windows では未サポート
  * essential
    * true になっているコンテナが停止するとタスクも停止する
  * entryPoint
    * `Entrypoint` にマップされる
  * command
    * `Cmd` にマップされる
  * workingDirectory
    * `WorkingDir` にマップされる
  * environmentFiles
    * 最大 10 ファイルを指定可能
    * ファイルの拡張子は .env である必要がある
    * Windows では未サポート
    * コンテナ定義にて環境変数が設定されている場合はそちらが優先される
    * value
      * S3 オブジェクトの ARN
    * type
      * s3 のみ
  * environment
    * name
    * value
  * secrets
    * name
    * valueFrom
      * Secrets Manager の ARN もしくは SSM Parameter Store の ARN
  * disableNetworking
    * `NetworkDisabled` にマップされる
  * links
    * bridge の場合のみサポート
    * awsvpc, Windows は未サポート
    * `"links": ["name:internalName", ...]` のように指定
  * hostname
    * `Hostname` にマップされる
    * awsvpc は未サポート
  * dnsServers
    * `Dns` にマップされる
    * awsvpc, Windows は未サポート
  * dnsSearchDomains
    * `DnsSearch` にマップされる
    * DNS 検索ドメインのリスト
    * awsvpc, Windows は未サポート
  * extraHosts
    * `ExtraHosts` にマップされる
    * `/etc/hosts` にホスト、IP アドレスの組みが追加される
    * awsvpc, Windows は未サポート
    * hostname
      * ホスト名
    * ipAddress
      * IP アドレス
  * readonlyRootFilesystem
    * `ReadonlyRootfs` にマップされる
    * ルートファイルシステムが読み取り専用でマウントされる
    * Windows は未サポート
  * mountPoints
    * `Volumes` にマップされる
    * Windows コンテナは $env:ProgramData と同じドライブに全部のディレクトリがマウントされる
    * sourceVolume
      * ボリューム名
    * containerPath
      * マウントするパス
    * readOnly
      * デフォルトは `false`
  * volumesFrom
    * 別のコンテナからマウントするデータボリューム
    * sourceContainer
      * ソースコンテナ名
    * readOnly
      * デフォルトは `false`
  * logConfiguration
    * `LogConfig` にマップされる
    * EC2 起動タイプの場合 `ECS_AVAILABLE_LOGGING_DRIVERS` にて使用可能なログドライバーを登録する必要がある
    * logDriver
      * awslogs など
      * Fargate の場合は awslogs、splunk、awsfirelens のみ
    * options
      * FireLens の場合は log-driver-buffer-limit を指定可能
    * secretOptions
      * ログ設定に渡すシークレット
      * name
        * 環境変数として渡す値
      * valueFrom
  * firelensConfiguration
    * options
      * `"enable-ecs-log-metadata": "true"` のような設定
    * type
      * fluentd | fluentbit
  * credentialSpecs
    * CredSpec ファイルの参照先
    * MyARN の箇所を ARN に置き換えて指定
    * S3 または SSM の ARN を指定
    * credentialspecdomainless:MyARN
      * 各タスクは異なるドメインに参加可能
    * credentialspec:MyARN
      * コンテナインスタンスをドメインに参加させる必要がある
  * privileged
    * ホストコンテナインスタンスに対する昇格されたアクセス権限が付与される
    * セキュリティ上、使用は推奨されない
    * `Privileged` にマップされる
  * user
    * `User` にマップされる
    * `host` ネットワークモードの場合は root (UID 0)での実行は推奨されない
    * Windows では未サポート
  * dockerSecurityOptions
    * "no-new-privileges" | "apparmor:PROFILE" | "label:value" | "credentialspec:CredentialSpecFilePath"
    * Fargate では未サポート
    * Linux では SELinux, AppArmor のカスタムラベルを参照できる
    * Active Directory 認証用のコンテナを設定する認証情報仕様ファイルを参照できる
    * コンテナエージェント側で `ECS_SELINUX_CAPABLE=true` もしくは `ECS_APPARMOR_CAPABLE=true` の設定が必要
  * ulimitx
    * `Ulimits` にマップされる
    * Fargate では nofile のみ指定可能
    * name
      * "core" | "cpu" | "data" | "fsize" | "locks" | "memlock" | "msgqueue" | "nice" | "nofile" | "nproc" | "rss" | "rtprio" | "rttime" | "sigpending" | "stack"
    * hardLimit
    * softLimit
  * dockerLabels
    * `Labels` にマップされる
    * コンテナに追加するラベルのキー/値のマップ
  * linuxParameters
    * Windows では未サポート
    * Fargate では `SYS_PTRACE` の追加のみ
    * capabilities
      * add
      * drop
  * devices
    * コンテナに公開するホストデバイス
    * Fargate, Windows は未サポート
    * hostPath
    * containerPath
    * permissions
      * read | write | mknod
  * initProcessEnabled
    * PID 1 として使われるべき init プロセスを指定できる
    * [Docker Reference](https://docs.docker.jp/engine/reference/run.html#init)
  * maxSwap
    * 0 を指定した場合はスワップは使用されない
    * 省略時はコンテナインスタンスのスワップ設定を使用
    * Fargate では未サポート
  * sharedMemorySize
    * /dev/shm ボリュームのサイズ値 (MiB) 
    * Fargate では未サポート
  * swappiness
    * スワップの動作を調整するパラメータ
    * Amazon Linux 2023 ではサポートされない
    * 0 の場合は必要な場合を除きスワップされない
    * 100 の場合は頻繁にスワップが行われる
    * 未指定時のデフォルトは 60
  * tmpfs
    * Fargate では未サポート
    * containerPath
    * mountOptions
    * size
  * dependsOn
    * containerName
    * condition
      * START
        * 依存元コンテナが開始していること
      * COMPLETE
        * 依存元コンテナが終了していること。初期化スクリプトコンテナなどのユースケース
      * SUCCESS
        * COMPLETE に加えて exit code 0 で終了していること
      * HEALTHY
        * 依存元コンテナがヘルスチェックに成功していること
  * startTimeout
    * 依存元のコンテナが startTimeout の間に目標のステータスにならない場合、コンテナを開始しない
  * stopTimeout
    * コンテナが終了しなかった場合に SIGKILL 発行までの時間
    * 未指定時、コンテナエージェント側で `ECS_CONTAINER_STOP_TIMEOUT` が設定されている場合はそちらを使用。コンテナエージェント側でも設定されていない場合は 30 秒
  * systemControls
    * Windows では未サポート
    * awsvpc, host の場合、かつ複数コンテナを含むタスクの場合は非推奨
      * awsvpc: 最後に起動したコンテナの設定が全コンテナで使用される
      * host: ネットワーク名前空間の systemControls はサポートされない
    * IPC リソース名前空間を設定している場合
      * host IPC モード: IPC 名前空間の systemControls はサポートされない
      * task IPC モード: IPC 名前空間の systemControls 値がタスク内のすべてのコンテナに適用される
    * namespace
      * "kernel.msgmax" | "kernel.msgmnb" | "kernel.msgmni" | "kernel.sem" | "kernel.shmall" | "kernel.shmmax" | "kernel.shmmni" | "kernel.shm_rmid_forced", and Sysctls that start with "fs.mqueue.*"
    * value
  * interactive
    * `OpenStdin` にマップされる
    * `true` の場合、stdin または tty を割り当てる必要があるコンテナ化されたアプリケーションをデプロイ可能
  * pseudoTerminal
    * `true` の場合、TTY が割り当てられる
* Elastic Inference accelerator name
  * deviceName
  * deviceType
* タスク配置制約
  * Fargate では未サポート
  * expression
    * クラスタークエリ言語
  * type
    * memberOf を使用
* proxyConfiguration
  * App Mesh 用設定
  * Windows では未サポート
  * type
    * APPMESH のみ
  * containerName
    * App Mesh プロキシとして使用するコンテナ名
  * properties
    * IgnoredUID
    * IgnoredGID
    * AppPorts
    * ProxyIngressPort
    * ProxyEgressPort
    * EgressIgnoredPorts
    * EgressIgnoredIPs
  * name
  * value
* volumes
  * 使用できるボリュームの種類はバインドマウント、Docker ボリューム(`/var/lib/docker/volumes` に作成される)
  * name
  * host
    * Windows コンテナは $env:ProgramData と同じドライブに全部のディレクトリをマウントできる
    * sourcePath
  * dockerVolumeConfiguration
    * scope
      * task | shared
    * autoprovision
    * driver
    * driverOpts
    * labels
  * efsVolumeConfiguration
    * fileSystemId
    * rootDirectory
    * transitEncryption
      * ENABLED | DISABLED
    * transitEncryptionPort
    * authorizationConfig
      * accessPointId
      * iam
        * タスクロールの使用可否を設定
        * 使用する場合は transitEncryption の有効化が必要
        * ENABLED | DISABLED
  * FSxWindowsFileServerVolumeConfiguration
    * fileSystemId
    * rootDirectory
    * authorizationConfig
      * credentialsParameter
      * domain
* Tags
  * key
  * value
* ephemeralStorage
  * Fargate におけるエフェメラルストレージサイズの指定
* ipcMode
  * Fargate, Windows は未サポート
  * host | task | none
  * host: ホストと同じ IPC リソースを共有
  * task: タスク内のコンテナ間で IPC リソースを共有
  * none: 各コンテナの IPC リソースはプライベートになる
* pidMode
  * Windows では未サポート
  * host | task
  * Fargate では task のみ
  * サイドカーのモニタリングのようなユースケース



[サービス定義パラメータ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/service_definition_parameters.html)

設定可能項目。

* launchType
* capacityProviderStrategy
  * capacityProvider
  * weight
  * base: base を設定できるのは 1 つのキャパシティプロバイダー
* taskDefinition: ローリングアップデート時は指定が必要
* platformFamily
* platformVersion: 未指定時は LATEST
* cluster: 未指定時は default
* serviceName
* schedulingStrategy: REPLICA or DAEMON
* desiredCount
* deploymentConfiguration
  * maximumPercent
    * DAEMON の場合: 100 にする必要がある
    * CODE_DEPLOY, EXTERNAL かつ EC2 の場合: DRAINING 状態の場合は RUNNING のままとできるタスク数の上限
    * Fargate の場合: 使用されない
  * minimumHealthyPercent
    * ELB 無しの構成の場合
      * RUNNING 状態に達した後 40 秒間待ってから、正常性の最小割合の合計にカウント
      * タスク内のすべての必須コンテナがヘルスチェックに合格すると、タスクは正常と見なされる
    * ELB ありの場合
      * ヘルスチェックが定義されている必須コンテナがない場合: ELB ヘルスチェックが正常ステータスを返すのを待ってからカウント
      * ヘルスチェックが定義されている必須コンテナがある場合: タスクが正常な状態となりELB ヘルスチェックが正常ステータスを返すのを待ってからカウント
    * CODE_DEPLOY, EXTERNAL かつ EC2 の場合: DRAINING 状態の場合は RUNNING のままとできるタスク数の下限
    * Fargate の場合: 使用されない
* deploymentController: ECS | CODE_DEPLOY | EXTERNAL
* placementConstraints
  * type
  * expression
* placementStrategy
  * type: random | spread | binpack
  * field
    * random の場合: このフィールドは使用されない
    * spread の場合: instanceId または attribute:ecs.availability-zone などのコンテナインスタンスに適用される任意のプラットフォームまたはカスタム属性
    * binpack: cpu | memory
* tags
  * key
  * value
* enableECSManagedTags: デフォルトは `false`
* propagateTags: TASK_DEFINITION | SERVICE
* networkConfiguration: awsvpc の場合に必要
  * awsvpcConfiguration
    * subnets
    * securityGroups
    * assignPublicIP
* healthCheckGracePeriodSeconds
* loadBalancers: `deploymentController` が `ECS` の場合のみ変更可能
  * targetGroupArn
  * loadBalancerName
  * containerName
  * containerPort
* role: サービスロール
  * このパラメーターは、サービスの単一のターゲットグループで ELB を使用していて、タスク定義が awsvpc を使用していない場合にのみ許可される
  * 未指定時はサービスにリンクされたロールが使用される。awsvpc ではサービスにリンクされたロールが必須
* serviceConnectConfiguration
  * enabled
  * namespace
  * services
    * portName
    * discoveryName
    * clientAliases
      * port
      * dnsName
    * ingressPortOverride
    * logConfiguration
* serviceRegistries
  * registryArn
  * port
  * containerName
  * containerPort
* clientToken: 冪等性確保のために使用される



#### 追加リソースのサポート

* [Working with GPUs on Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-gpu.html)
  * エージェント設定ファイルで ECS_ENABLED_GPU_SUPPORT を true に設定する必要がある。
* [Working with inference workloads on Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-inference.html)



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




