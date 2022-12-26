
[ベストプラクティスガイド](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/bestpracticesguide/intro.html)


## ネットワーク

[インターネットへの接続](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/bestpracticesguide/networking-outbound.html)

* EC2 の場合、ネットワークモード awsvpc だとパブリック IP アドレスの付与ができない


[インターネットからのインバウンド接続の受信](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/bestpracticesguide/networking-inbound.html)

ALB の利点

* SSL/TLS 終端ができる。
* 複数の DNS ホスト名を設定できる。
* HTTP パスなどに基づいたルーティング。
* gRPC、web socket もルーティング可能。
* WAF 連携可能。

NLB の利点

* SSL/TLS 終端ができる。
* UDP サポート。

API Gateway の利点

* SSL/TLS 終端ができる。
* リクエストごとに課金。
* HTTP パスに基づいたルーティング。


[ネットワークモードの選択](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/bestpracticesguide/networking-networkmode.html)

#### host

* ホスト側のポートと 1:1 で対応するので、コンテナポート番号の重複ができない。
* ホスト上のプライベートループバックネットワークサービスに接続できる。セキュリティ上のリスクとなる。

#### bridge

* 静的もしくは動的なポートマッピングが可能。動的ポートマッピングの場合、ELB もしくはサービスディスカバリによって動的にアサインされたポート番号と連携可能。
* 動的ポートマッピングの場合、セキュリティグループで許可するポート範囲が広くなることはデメリットの一つ。

#### awsvpc

* タスクごとに ENI を作成。そのため、タスクごとにセキュリティグループを設定できることは利点。
* ENI にはホストと異なる CIDR が設定されていてもよい。
* インスタンスにアタッチ可能な ENI 数が制限されていることがデメリットの一つ。
* IP アドレスの枯渇も考慮すべき点。
* ENI Trunking を使用することで、より多くの ENI をアタッチできるようになる。ただし、ENI Trinking 有効化時にはタスクの起動までの時間が長くなる。


[VPC 内からの接続](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/bestpracticesguide/networking-connecting-vpc.html)

NAT Gateway

* 宛先の制限をできない点はデメリットの一つ。
* NAT Gateway がボトルネックとなる場合は、サブネットのルートごとに別々の NAT Gateway を設定することを検討。


[VPC 内の Amazon ECS サービス間のネットワーキング](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/bestpracticesguide/networking-connecting-services.html)

サービス検出

* bridge の場合は A レコードだけだと不十分。SRV レコードによる解決が必要。
* タスクが停止した場合など、TTL が切れるまではそのタスクの DNS レコードもキャッシュされている。アプリケーション側では、そのことを考慮したリトライロジックが必要。

内部ロードバランサ

* タスクの稼働状況はロードバランサ側で対応されるため、アプリケーション側で考慮しなくてもよい点が利点。
* コストがデメリット。複数のサービスで ELB を共有するのが対応策の一つ。

サービスメッシュ

* Envoy にプロキシする方式。
* Envpy 側で負荷分散、リトライを対応してくれる。
* ロードバランサと比べてホップ数が少ないので、その分のオーバーヘッドを回避できる点も利点。


[アカウント、VPC をまたがったネットワーキング](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/bestpracticesguide/networking-connecting-services-crossaccount.html)

以下のような選択肢がある。

* AWS Transit Gateway
* VPN
* VPC ピアリング
* 共有 VPC


[最適化とトラブルシューティング](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/bestpracticesguide/networking-troubleshooting.html)

トラブルシューティング時に有用なメトリクス、トレース、ログ

* Container Insights によるメトリクスの確認
* X-Ray によるトレーシング
* VPC フローログ

調整可能なパラメータ

* ulimit(ファイルディスクリプタ)
* sysctl net



## Auto Scaling とキャパシティ管理

[タスクサイズの決定](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/bestpracticesguide/capacity-tasksize.html)

* メモリ量は ps, top, CloudWatch メトリクスなどの方法で確認可能。


[サービスの Auto Scaling の設定](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/bestpracticesguide/capacity-autoscaling.html)

* 使用率、飽和率を示すパラメータの特定が必要。需要と相関するメトリクスを選ぶこと。
* 負荷テストを行うことを推奨する。
  * CPU 使用率、メモリ使用率、I/O 操作、I/O キューの深さ、ネットワークスループットなどを監視する。
  * 負荷を段階的に上げていく。
  * 飽和状態に達するリソースの特定。
* ワーカーベースのワークロードのように適切なメトリクスが標準で用意されていない場合は、カスタムメトリクスとしてプッシュする対応をとる。
* Java アプリケーションの場合、できるだけ多くのメモリを JVM ヒープに割り当てることを推奨。最近の JVM バージョンではヒープ内におさまるように自動的に割り当てられるため、メモリベースでのスケーリングは推奨しない。
* .Net や Ruby などもメモリ使用率がスループット、同時実行性と相関しないため、メモリベースでのスケーリングは推奨しない。
* キューから取り出すタイプのアプリケーションの場合は、キューの深さが適切なメトリクスとなる。


[容量と可用性](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/bestpracticesguide/capacity-availability.html)

* スケーリングのターゲットは 60 〜 80 % が目安。
* 複数の AZ を使用することを推奨。

スケーリング速度の最大化のための方策

* イメージサイズの最小化
  * バイナリを配置するだけでよい言語の選定
  * FROM で指定するベースイメージを最小のものにする
  * マルチステージビルドを活用。ビルドの中間生成物は含めないようにする 
  * 「&&」を使用することで 1 レイヤーに複数操作を含める。レイヤー数を少なくできるため
  * データは起動後に S3 からフェッチするなどの方法を検討する
* コンテナレジストリはレイテンシの少なくするように。ECR は同一リージョンを使用する
* ロードバランサのヘルスチェックの間隔、しきい値を小さく設定する
* より大きなインスタンスサイズ、EBS ボリュームの選定
* awsvpc よりも bridge の方が起動が早い

クォータの計画

* ビジネス部門と連携し、ニュース配信などのスケジュールを共有
* 事前に上限緩和をしておく


[Amazon EC2 インスタンスタイプの選択](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/bestpracticesguide/ec2-instance-type.html)

* インスタンスが小さい方が、余剰リソースが少なくなるため、コスト面で有利
* 複数のインスタンスタイプが必要な場合は、ASG を個別に作成するとよい。ECS サービス、タスク側では対応するキャパシティープロバイダーを設定すればよい


[Amazon EC2 スポットと FARGATE_SPOT の使用](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/bestpracticesguide/ec2-and-fargate-spot.html)

* スポットは一時的なダウンタイムを許容するワークロードに向いている。
* 需要が非常に多い場合は、スポットのリソースを起動できない可能性がある点に注意。
* スポットインスタンスの中断へ対応できるとよい。

スポット容量不足を最小限に抑えるための対策

* 複数のリージョン、AZ の使用
* EC2 の場合は複数のインスタンスタイプの使用。容量とコストに最適化された割り当て戦略の使用



## 永続ストレージ

* サブミリ秒のレイテンシが求められ、スケールアウトの考慮が不要な場合は EBS が向いている。

[Amazon EFS](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/bestpracticesguide/storage-efs.html)

* 共有ファイルシステムを必要とするユースケースに向いている。データ分析、メディア処理、コンテンツ管理、ウェブ配信など。
* ミリ秒未満のレイテンシが求められる場合は向かない。
* タスク定義にて EFS ボリュームのマウントを設定する
* ECS タスクが稼働する可能性のある全ての AZ にマウントターゲットを用意するとよい
* アクセス制御の設定方法
  * セキュリティグループ
  * IAM。マウントする権限の設定を行うことができる
  * EFS アクセスポイント
* コスト最適化
  * 複数のアプリケーションで EFS を共用する。アクセス分離はアクセスポイントによって行う。
  * アクセス頻度の低いデータはストレージクラスを移動させる。


[Docker ボリューム](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/bestpracticesguide/storage-dockervolumes.html)

* EBS ボリュームを使用することでデータの永続化をできる。
* EBS ボリュームを使いまわす場合は、同一の AZ にタスクを配置するように配置制約を設定する。
　　attribute:ecs.availability-zone == us-east-1a


[Amazon FSx for Windows ファイルサーバー](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/bestpracticesguide/storage-fsx.html)

* サーバーメッセージブロック（SMB）プロトコルでアクセス可能
* Windows コンテナインスタンスが Active Directory ドメインサービス (AD DS) 上のドメインメンバーとなっている必要がある



## セキュリティ

[AWS Identity and Access Management](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/bestpracticesguide/security-iam.html)

* 最小権限の原則に従う
* Condition 句
  * MFA を要求するように設定
  * 特定のタグが付与されたリソースのみ操作可能とするように設定
* クラスター単位で管理境界にする


[Amazon ECS タスクでの IAM ロールの使用](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/bestpracticesguide/security-iam-roles.html)

#### タスクロール

* ECS Agent がタスクロールをフェッチし、タスク用のクレデンシャルをコンテナ内に格納する仕組み。コンテナ内ではメタデータアクセスによって取得できる。環境変数は AWS_CONTAINER_CREDENTIALS_RELATIVE_URI。

#### タスク実行ロール

* ECS Agent が API を発行する際に使用。イメージのプル、CloudWatch Logs のログストリーム作成など。

#### コンテナインスタンスのロール

* コンテナインスタンスのクラスターへの登録などに必要。init コマンドは OS 上で行われるため。

#### サービスにリンクされたロール

* ENI の作成、ターゲットグループへの登録/登録解除などの用途で必要。

#### 推奨事項

* EC2 メタデータへのアクセスをドロップ
  * EC2 の場合
```
sudo yum install -y iptables-services; sudo iptables --insert FORWARD 1 --in-interface docker+ --destination 192.0.2.0/32 --jump DROP
```
  * awsvpc の場合。ECS_AWSVPC_BLOCK_IMDS を true に設定する。
  * host の場合。ECS_ENABLE_TASK_IAM_ROLE_NETWORK_HOST を false に設定する。
* awsvpc の使用。タスク単位でネットワークを分離できるため。
* IAM アクセスアドバイザーによって、最近使用されていないアクションを確認し、当該アクションを削除。
* CloudTrail のモニタリング。


[ネットワークセキュリティ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/bestpracticesguide/security-network.html)

* 転送時の暗号化。
  * Envoy の場合は Envoy 間で TLS 接続を設定できる。
  * Nitro インスタンス間の通信は自動的に暗号化される。
* awsvpc モードを推奨。タスクごとにセキュリティグループを設定できるため。
* プライベートリンクによって、インターネットを通らない環境での使用。
* ネットワークトラフィックを厳密に分離する必要がある場合は、VPC を分ける。
* VPC フローログの有効化。


[シークレット管理](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/bestpracticesguide/security-secrets-management.html)

* シークレット管理のプラクティス
  * Secrets Manager に格納し、タスク定義で参照設定する。
  * 暗号化されている S3 バケットから取得。
  * シークレットを格納したボリュームをマウント。


[コンプライアンス](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/bestpracticesguide/security-compliance.html)

コンプライアンスに関連する文書へのリンクがまとめられている。


[ログ記録とモニタリング](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/bestpracticesguide/security-logging-and-monitoring.html)

* CloudWatch Alarm, CloudWatch Logs, CloudTrail など各サービスと連携可能
* CloudWatch Logs にログ送信可能。FireLens によって Fluent Bit, Fluentd を設定することも可能。


[AWS Fargate のセキュリティ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/bestpracticesguide/security-fargate.html)

* 2020/5/28 以降に起動されたタスクのエフェメラルストレージでは AEDS-256 暗号化が AWS Fargae によって管理される暗号キーによってなされている。
* capability は SYS_PTRACE を許可できるだけとなっている。


[タスクとコンテナのセキュリティ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/bestpracticesguide/security-tasks-containers.html)

推奨事項

* 最小のイメージの使用。
* イメージのスキャンを行い、脆弱性が含まれていないか検出する。
* setuid, setgid が設定されているファイルを配置しない。
* Docker Hub からイメージを取得しないようにする。管理者によってキュレーションされたイメージのみ使用するようにする。
* アプリケーションパッケージ、ライブラリについてもスキャンを行う。
* 静的コード分析の実行。SQL インジェクションが含まれていないかなどの調査のため。
* 非 root ユーザでコンテナ内プロセスを起動する。
* 読み取り専用のルートファイルシステムの使用。
* CPU とメモリ使用量の Limit を制限。
* イミュータブルなタグの使用。
* privileged を使用しない。
* 不要な Linux capability を削除。
* カスタマー管理キーによる ECR イメージの暗号化。


[ランタイムセキュリティ](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/bestpracticesguide/security-runtime.html)

* seccomp の使用。特定のシステムコール実行を防ぐ効果。


参考 [Docker 向け Seccomp セキュリティプロファイル](https://matsuand.github.io/docs.docker.jp.onthefly/engine/security/seccomp/)






