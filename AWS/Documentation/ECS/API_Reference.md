
[Amazon Elastic Container Service API Reference](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/Welcome.html)


## Actions

[CreateService](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_CreateService.html)

* サービスの特徴
  * desired count を下回ると、desired count を満たすようにタスクを起動する
  * ELB と連携可能
  * ELB を使用していない場合は RUNNING 状態の場合に healthy だと扱う
  * スケジューラ戦略は `REPLICA`, `DAEMON` の 2 通り
  * `REPLICA` だと `minimumHealthyPercent` のデフォルト値は 100。`DAEMON` の場合は 0。
  * `maximum percent` の値は RUNNING, PENDING 状態で起動できるタスクの上限数(デプロイコントローラが ecs の場合)
  * デプロイコントローラが CODE_DEPLOY or EXTERNAL の場合、`minimum healthy percent`、`maximum percent` は DRAINNNG 状態の際に RUNNING 状態となれるタスク数の下限、上限を示す
  * Fargate 起動タイプでは `minimum healthy percent`、`maximum percent` は使用されない
* Request Parameters
  * capacityProviderStrategy
    * 指定された場合は `launchType` は省略される
    * `capacityProviderStrategy`、`launchType` が無指定の場合はクラスターの `defaultCapacityProviderStrategy` が使用される
  * healthCheckGracePeriodSeconds
    * サービススケジューラが ECS タスク起動後に ELB ヘルスチェックの unhealthy を無視する秒数
    * ELB を使用しない構成では `startPeriod` の使用を推奨
  * loadBalancers
    * デプロイコントローラが ECS の場合
      * ELB ターゲットグループの ARN を指定する。複数指定可能
      * 複数のターゲットグループを指定する場合は、サービスにリンクされたロールが必要
    * デプロイコントローラが CODE_DEPLOY の場合
      * 二つのターゲットグループを指定する
      * デプロイ中タスクセットの状態を PRIMARY のセットしターゲットグループと関連づける。置き換え用タスクセットを別のターゲットグループと関連づける。
    * awsvpc の場合はターゲットタイプは instance にできず ip にする必要がある。awsvpc では EC2 ではなく ENI に関連づけられるため
  * networkConfiguration
    * awsvpc の場合のみ必要
  * role
    * ELB を使用し awsvpc ではない場合のみ指定が許可される
    * サービスにリンクされたロールが存在する場合は、role で指定したロールは使用されずサービスにリンクされたロールが使用される
    * パスを「/」以外にしている場合、ロールの ARN もしくはパス付きでロール名を指定(ロール名が bar でパスが /bar/ の場合は、/foo/bar) が必要


## Data Types

[CapacityProvider](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_CapacityProvider.html)

* autoScalingGroupProvider
* status
  * ACTIVE | INACTIVE
* updateStatus
  * DELETE_IN_PROGRESS | DELETE_COMPLETE | DELETE_FAILED


[Cluster](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_Cluster.html)

* configuration
  * executeCommandConfiguration
* defaultCapacityProviderStrategy
* serviceConnectDefaults
  * Service Connect のデフォルトのネームスペースを設定
  * Service Connect を有効化したサービスかつ ServiceConnectConfiguration を無指定の場合はデフォルトのネームスペースが設定される
  * マネージドプロキシコントローラによりメトリクス、ログが収集される
* settings
  * 現状 Container Insights の有効化/無効化のみ。アカウント設定の内容を上書きできる
* status
  * ACTIVE, PROVISIONING, DEPROVISIONING, FAILED, INACTIVE


[Container](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_Container.html)

**TODO**


[ContainerDefinition](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_ContainerDefinition.html)

**TODO**


[ContainerInstance](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_ContainerInstance.html)

* agentConnected
  * エージェントが unhealthy もしくは stopped の場合 false となる。false となったインスタンスはタスク配置対象外
* agentUpdateStatus
  * PENDING | STAGING | STAGED | UPDATING | UPDATED | FAILED
* attachments
  * コンテナインスタンスにアタッチされたリソース
* attributes
  * 属性。コンテナエージェントによる登録時もしくは `PutAttributes` により設定される
* capacityProviderName
* healthStatus
  * コンテナインスタンスのヘルス状態。OK | IMPAIRED | INSUFFICIENT_DATA | INITIALIZING
* registeredResources
  * コンテナインスタンス上で利用可能なリソース量の合計
* remainingResources
  * 残リソース量
* status
  * REGISTERING, REGISTRATION_FAILED, ACTIVE, INACTIVE, DEREGISTERING, or DRAINING.
  * REGISTERING: awsvpcTrunking にオプトインしている場合、コンテナインスタンスの登録時のステータス
  * REGISTRATION_FAILED: 失敗理由は statusReason で確認可能
  * DEREGISTERING: ENI のデプロビジョニング中
  * INACTIVE: インスタンスの登録解除が完了
  * ACTIVE: タスクを受け入れられる状態
  * DRAINING: スタンドアローンタスクは配置対象外。サービスから起動しているタスクは可能であれば除かれる


[ContainerInstanceHealthStatus](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_ContainerInstanceHealthStatus.html)

* overallStatus
  * OK | IMPAIRED | INSUFFICIENT_DATA | INITIALIZING


[Deployment](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_Deployment.html)

デプロイコントローラが ECS の時だけ使用される

* rolloutState
  * デプロイ中は IN_PROGRESS。steady state になると COMPLETED。サーキットブレイカーが発動すると FAILED
* rolloutStateReason
  * ロールアウト理由
* status
  * PRIMARY: 最新のデプロイ
  * ACTIVE: 新規デプロイ実行中
  * INACTIVE: デプロイ完了し置き換え済み


[HealthCheck](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_HealthCheck.html)

コンテナヘルスチェックに関するオブジェクト。コンテナエージェントはタスク定義に指定されたヘルスチェックのみモニタリングしレポートする。コンテナイメージ内に埋め込まれたヘルスチェックは対象外。

コンテナの `healthStatus` の値は以下の通り。

* HEALTHY: ヘルスチェックにパス
* UNHEALTHY: ヘルスチェックが失敗
* UNKNOWN: ヘルスチェックの評価中もしくはコンテナヘルスチェックが未設定

タスクの `healthStatus` の値は以下の通り。non-essential コンテナのヘルスチェック結果はタスクのヘルスチェック結果に影響を及ぼさない

* HEALTHY: 全ての essential コンテナはヘルスチェクにパス
* UNHEALTHY: 一つもしくは複数の essential コンテナのヘルスチェックが失敗
* UNKNOWN: essentiak コンテナはヘルスチェックの評価中もしくは non-essenntial コンテナのみにヘルスチェックが設定されているもしくはヘルスチェックが設定されているコンテナがない

その他の考慮事項

* スタンドアローンのタスクの場合は、ヘルスチェックの結果に関わらず稼働し続ける
* ECS Agent が ECS サービスと接続できない場合、サービスはコンテナを UNHEALTHY と報告する

コンテンツ

* command
* interval
* retries
* startPeriod
* timeout


[LinuxParameters](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_LinuxParameters.html)

**TODO**


[LogConfiguration](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_LogConfiguration.html)

**TODO**


[ManagedScaling](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_ManagedScaling.html)

* instanceWarmupPeriod: 新規起動した EC2 インスタンスが Auto Scaling グループの CloudWatch メトリクスに寄与できるようになるまでの時間(インスタンス起動後に次のスケールアウトが発生するまでの期間？)
* maximumScalingStepSize: スケールアウト時に同時に起動するインスタンスの最大数。スケールインには寄与しないパラメータ
* minimumScalingStepSize: スケールアウト時に同時に起動するインスタンスの最小数。スケールインには寄与しないパラメータ。Desired より多くなる見込みの場合でもこのパラメータの分だけ起動する
* targetCapacity: 0 〜 100 で指定。例えば 90 で指定した場合は 10 % 分の余剰リソースを起動する


[PlacementConstraint](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_PlacementConstraint.html)

* expression
  * 2000 字まで
* type
  * distinctInstance | memberOf


[PlacementStrategy](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_PlacementStrategy.html)

* field
  * spread の場合は instanceId、プラットフォーム、`attribute:ecs.availability-zone` のようなカスタム属性
  * binpack の場合は cpu or memory
  * random の場合は field は使われない
* type
  * random | spread | binpack


[PortMapping](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_PortMapping.html)

[Create a container](https://docs.docker.com/engine/api/v1.35/#tag/Container/operation/ContainerList) に記載されている `PortBindings` にマッピングされる。

* appProtocol
  * Service Connect 使用時に使用される。Service Connect プロキシにプロトコル固有のコネクションハンドリングを設定する
* containerPort
  * awsvpc の場合は `containerPort` の設定だけでよい
* containerPortRange
  * コンテナ側のポート範囲
* hostPort
  * bridge の場合、無指定の場合は動的にポートが割り当てられる
  * 一時ポート範囲は `/proc/sys/net/ipv4/ip_local_port_range` にリストされている


[RuntimePlatform](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_RuntimePlatform.html)

* cpuArchitecture
  * X86_64 | ARM64
* operatingSystemFamily
  * WINDOWS_SERVER_2019_FULL | WINDOWS_SERVER_2019_CORE | WINDOWS_SERVER_2016_FULL | WINDOWS_SERVER_2004_CORE | WINDOWS_SERVER_2022_CORE | WINDOWS_SERVER_2022_FULL | WINDOWS_SERVER_20H2_CORE | LINUX


[Service](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_Service.html)

* taskSets
  * CODE_DEPLOY, EXTERNAL の場合のみ含まれる情報


[ServiceConnectClientAlias](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_ServiceConnectClientAlias.html)

* port: Service Connect 向けの LISTEN ポート番号。同 namespace 内のタスクからこのポート番号を使用できる
* dnsName:
  * サービスに接続するための DNS 名
  * 未指定時は `discoveryName.namespace` になる
  * `discoveryName` 未指定時は `portName.namespace` になる


[ServiceConnectConfiguration](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_ServiceConnectConfiguration.html)

* enabled
* logConfiguration
* namespace: 名前空間名もしくは Cloud Map の ARN
* services
  * 他の ECS サービスから接続するための名前
  * クライアントとしてのみ使用する場合は、この設定は不要
  * `ServiceConnectService` オブジェクトをリストで指定。つまり複数設定できる


[ServiceConnectService](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_ServiceConnectService.html)

* portName: タスク定義の `portMappings` を指定
* clientAliases: `ServiceConnectClientAlias` を複数指定可能
* discoveryName: Cloud Map に作成されるサービス名。未指定時は `portMappings` が使用される
* ingressPortOverride: 指定したポート番号宛の通信を `portMapping` で指定されたポート番号にバイパスする


[ServiceConnectServiceResource](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_ServiceConnectServiceResource.html)

* discoveryArn: Cloud Map の名前空間の ARN
* discoveryName: Cloud Map に作成されるサービス名。未指定時は `portMappings` が使用される


[ServiceRegistry](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_ServiceRegistry.html)

複数のサービスレジストリには登録できない。

* containerName: bridge, host の場合は containerName, containerPort の組を指定。awsvpc の場合で SRV DNS レコードを使用している場合は containerName, containerPort の組、もしくは port を指定
* containerPort
* port
* registryArn


[SystemControl](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_SystemControl.html)

ネットワークモードが awsvpc もしくは host の場合に、複数コンテナが含まれるタスクでネットワーク関連のパラメータを指定することは推奨しない。awsvpc の場合はあるコンテナに定義したパラメータが全コンテナに適用される。もしくはコンテナごとに別々のパラメータを適用した場合は最後に起動したコンテナのパラメータが適用される。host の場合はコンテナインスタンスのカーネルパラメータに適用されるため、コンテナインスタンス上の全コンテナに影響を及ぼす。

* namespace
* value


[Task](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_Task.html)

**TODO**


[TaskDefinition](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_TaskDefinition.html)

**TODO**


[TaskSet](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_TaskSet.html)

* stabilityStatus
  * 以下を満たす場合は STEADY_STATE になる。そうでない場合は STABILIZING
    * runningCount が computedDesiredCount と同じ場合
    * pendingCount が 0 の場合
    * DRAINING 状態のコンテナインスタンスで稼働中のタスクがない場合
    * 全てのタスクが ELB、サービスディスカバリ、コンテナヘルスチェックに healthy となっている場合
* status
  * PRIMARY: 本番トラフィックを提供しているタスクセット
  * ACTIVE: 本番トラフィックを提供していないタスクセット
  * DRAINING: タスクセット内のタスクを停止中で、ターゲットグループから登録解除中


[Tmpfs](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_Tmpfs.html)

* containerPath
  * tmpfs ボリュームマウント先の絶対パス
* size
* mountOptions


[Ulimit](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_Ulimit.html)

Fargate の場合は上書きできるのは `nofile` のみ。`nofile` のデフォルトのソフトリミットは 1024, ハードリミットは 4096

* hardLimit
* name
  * core | cpu | data | fsize | locks | memlock | msgqueue | nice | nofile | nproc | rss | rtprio | rttime | sigpending | stack
* softLimit


[VersionInfo](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_VersionInfo.html)

コンテナインスタンスの Docker, コンテナエージェントのバージョン情報

* agentHash
  * amazon-ecs-agent GitHub リポジトリの Git コミットハッシュ
* agentVersion
* dockerVersion


[VOLUME](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_Volume.html)

* dockerVolumeConfiguration: Docker ボリュームを使用している場合に使用される。バインドマウント時は host が使用される
  * autoprovision
  * driver: 使用する Docker Volume Driver
  * driverOpts
  * labels
  * scope: task | shared。shared の場合はタスク停止後も残される
* efsVolumeConfiguration
  * fileSystemId
  * authorizationConfig
    * accessPointId
    * iam: ENABLED | DISABLED。EFS マウント時にタスクロールを使用するか否か
  * rootDirectory
  * transitEncryption: ENABLED | DISABLED。EFS、ECS ホスト間で暗号化するか否か
  * transitEncryptionPort
* fsxWindowsFileServerVolumeConfiguration
  * authorizationConfig
    * credentialsParameter: Secrets Manager or SSM Parameter の ARN
    * domain
  * fileSystemId
  * rootDirectory
* host: バインドマウント時に使用される。host パラメータが空の場合は Docker デーモンがホストパスをアサインする。しかし、タスク停止後の永続化は保証されない
* name


