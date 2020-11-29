
# AWS Batch

## 用語

* ジョブ: ジョブ定義の内容を元にジョブの実行内容を設定する。ジョブは ECS コンテナインスタンス上のコンテナで実行される。
* ジョブ定義: ジョブの実行方法を設定するもの。
* ジョブキュー: ジョブの投入先。ジョブキューには一つもしくは複数のコンピューティング環境を設定することができる。コンピューティング環境ごとに優先度を設定できる。
* コンピューティング環境: コンピューティングリソースのセット。インスタンスタイプやスポットインスタンス、vCPU 数などを設定可能。


## ジョブ

[ジョブ](https://docs.aws.amazon.com/ja_jp/batch/latest/userguide/jobs.html)


[ジョブの送信](https://docs.aws.amazon.com/ja_jp/batch/latest/userguide/submit_job.html)

設定項目

* ジョブ名
* ジョブ定義
* ジョブキュー
* ジョブタイプ(単一、配列)
* ジョブの依存関係
* 再試行回数
* タイムアウト時間
* パラメータ
* vCPU 数
* メモリ
* GPU 数
* コマンド
* 環境変数


[ジョブの状態](https://docs.aws.amazon.com/ja_jp/batch/latest/userguide/job_states.html)

* SUBMITTED: ジョブが送信済みでスケジューラの評価待ち。
* PENDING: 依存関係があり、まだ実行できない状態。
* RUNNABLE: 実行可能状態。コンピューティングリソースが使用可能であれば次の状態に遷移する。
* STAGING: ジョブがホストにスケジュールされており、コンテナ初期化処理中。
* RUNNING: ジョブ実行中。0 以外の終了コードで終了し再試行が設定されている場合は RUNNABLE に遷移。
* SUCCEEDED: ジョブが終了コード 0 で終了。
* FAILED: ジョブが終了コード 0 以外で終了。


[AWS Batch ジョブ環境変数](https://docs.aws.amazon.com/ja_jp/batch/latest/userguide/job_env_vars.html)

いくつかの環境変数が自動設定される。例をいくつか記載する。

+ AWS_BATCH_JOB_ARRAY_INDEX: 配列ジョブのインデックス番号。
* AWS_BATCH_JOB_ID: AWS Batch ジョブ ID。
* AWS_BATCH_JOB_NODE_INDEX: マルチノードジョブのノードインデックス番号。


[ジョブの再試行の自動化](https://docs.aws.amazon.com/ja_jp/batch/latest/userguide/job_retries.html)

ジョブがジョブキューに送信されて RUNNING 状態になると、1 回の試行とみなされる。
ジョブが終了コード 0 以外あるいは要害などの理由で失敗した場合、再試行設定で指定した試行回数分、リトライされる。


[ジョブの依存関係](https://docs.aws.amazon.com/ja_jp/batch/latest/userguide/job_dependencies.html)

ジョブの依存関係を設定できる。例えばジョブ A の終了後にジョブ B が実行するように制御可能。
配列ジョブの場合は、 SEQUENTIAL タイプの依存関係を指定可能。この場合は、子ジョブがインデックス 0 から開始し、インデックス 0 のジョブが完了してからインデックス 1 のジョブを開始する動作となる。


[ジョブのタイムアウト](https://docs.aws.amazon.com/ja_jp/batch/latest/userguide/job_timeouts.html)

タイムアウトの時間を設定し、時間に達したらジョブを自動的に打ち切ることができる。60 秒以上に設定する必要がある。SIGTERM が発行され、30 秒以内に停止していなかった場合は SIGKILL が発行される。タイムアウトはベストエフォートとなっており、数秒遅く発動する場合がある。


[配列ジョブ](https://docs.aws.amazon.com/ja_jp/batch/latest/userguide/array_jobs.html)

モンテカルロシミュレーションジョブ、パラメータスイープジョブ、大規模なレンダリングジョブなどのユースケースに向いている。

ジョブ内でのジョブの識別は AWS_BATCH_JOB_ARRAY_INDEX 環境変数によって行う。


[マルチノードの並列ジョブ](https://docs.aws.amazon.com/ja_jp/batch/latest/userguide/multi-node-parallel-jobs.html)

AWS Batch マルチノードの並列ジョブは、IP ベースのノード間通信をサポートするフレームワークのすべて (Apache MXNet、TensorFlow, Caffe2 や Message Passing Interface (MPI) など) と互換性がある。

**考慮事項**

* 単一の AZ でクラスタープレイスメントグループを作成して、コンピューティング環境と関連付けると良い
* スポットインスタンスを使用するコンピューティング環境はサポートされない


[GPU ジョブ](https://docs.aws.amazon.com/ja_jp/batch/latest/userguide/gpu-jobs.html)

サポートされているインスタンスタイプを使用しなかった場合は、RUNNABLE で固まる可能性がある。


## ジョブ定義

[ジョブ定義の作成](https://docs.aws.amazon.com/ja_jp/batch/latest/userguide/create-job-definition.html)

設定項目

* ジョブ定義名
* ジョブ試行
* タイムアウト
* パラメータ
* ジョブロール
* コンテナイメージ
* コマンド
* vCPU
* メモリ
* GPU
* セキュリティ
  * 特権の使用
  * コンテナ内で使用するユーザー
* マウントポイント
  * コンテナ内のパス
  * 対象のボリューム
  * 読み取り専用の有無
* ボリューム
  * 名前
  * インスタンス側のパス
* 環境変数
* ulimit
* Linux パラメータ（デバイスマッピング）


[複数ノードの並列ジョブ定義を作成する](https://docs.aws.amazon.com/ja_jp/batch/latest/userguide/multi-node-job-def.html)

マルチノードの並列ジョブは、「複数ノード設定が必要なジョブ」を選択して設定する。

ノード数、ノードグループの設定が追加で必要。


[ジョブ定義のパラメータ](https://docs.aws.amazon.com/ja_jp/batch/latest/userguide/job_definition_parameters.html)

ジョブ定義のパラメータのリファレンス。


[awslogs ログドライバーを使用する](https://docs.aws.amazon.com/ja_jp/batch/latest/userguide/using_awslogs.html)

awslogs ログドライバーにより STDOUT, STDERR の内容を CloudWatch Logs に送ることができる。


[機密データの指定](https://docs.aws.amazon.com/ja_jp/batch/latest/userguide/specifying-sensitive-data.html)

機密データは AWS Secrets Manager のシークレットまたは AWS Systems Managerパラメータストアのパラメータに格納する。

ジョブ定義内で Secrets Manager もしくは Systemd Manager の ARN を指定し、環境変数に設定することが可能。ログ設定の箇所でも設定可能。



## ジョブキュー

[ジョブキューの作成](https://docs.aws.amazon.com/ja_jp/batch/latest/userguide/create-job-queue.html)

ジョブキューの作成時に複数のコンピューティング環境を設定可能。また、優先順位を設定可能。

ジョブキューに優先順位を設定することが可能。一つのコンピューティング環境が複数のジョブキューに関連付けられている場合、優先度の高いジョブキューからジョブがスケジューリングされる。

[ジョブのスケジュール設定](https://docs.aws.amazon.com/ja_jp/batch/latest/userguide/job_scheduling.html)

ジョブの実行順は、他のジョブへの依存関係がすべて満たされている限り、送信順とほぼ同じ。



## コンピューティング環境

[https://docs.aws.amazon.com/ja_jp/batch/latest/userguide/compute_environments.html](https://docs.aws.amazon.com/ja_jp/batch/latest/userguide/compute_environments.html)

コンピューティング環境には、コンテナ化されたバッチジョブを実行するための Amazon ECS コンテナインスタンスが含まれている。

ジョブキューに関連付けられたコンピューティングごとに順番がある。最初のコンピューティング環境に使用可能なリソースがある場合、そのコンピューティング環境にスケジュールする。使用可能なリソースがない場合、次のコンピューティング環境にスケジュールする。

**マネージド型のコンピューティング環境**

○インスタンスの種類  
オンデマンドインスタンスかスポットインスタンスかを選択できる。
上限価格を設定し、スポットインスタンスの料金がオンデマンド料金の指定されたパーセンテージを下回る場合にのみスポットインスタンスを起動するように設定可能。

○ECS コンテナインスタンスとして  
環境の作成時に指定した VPC およびサブネット内で Amazon ECS コンテナインスタンスを起動する仕組みとなっている。

**アンマネージド型コンピューティング環境**

コンピューティング環境に紐付いている ECS クラスターに手動でコンテナインスタンスを追加する方式。


[コンピューティングリソース AMI](https://docs.aws.amazon.com/ja_jp/batch/latest/userguide/compute_resource_AMIs.html)

AMI を自前で用意している場合は上記ドキュメントの内容を満たすように作成する必要がある。


[起動テンプレートのサポート](https://docs.aws.amazon.com/ja_jp/batch/latest/userguide/launch-templates.html)

コンピューティング環境には、起動テンプレートを設定可能。
自前 AMI を作成しなくても、起動テンプレートで対応できる場合がある(EBS やユーザーデータなど)。


[コンピューティング環境の作成](https://docs.aws.amazon.com/ja_jp/batch/latest/userguide/create-compute-environment.html)

設定項目(マネージド型の場合)

* マネージド or アンマネージド
* サービスロール
* インスタンスロール
* EC2 キーペア
* プロビジョニングモデル(EC2 or スポット)
* 許可されたインスタンスタイプ
* 配分戦略
* 起動テンプレートとバージョン
* 最小 vCPU
* 希望 vCPU
* 最大 vCPU
* 独自 AMI を使用するかどうか
* ネットワーキング(VPC、サブネット、セキュリティグループ)


[配分戦略](https://docs.aws.amazon.com/ja_jp/batch/latest/userguide/allocation-strategies.html)

マネージド型のコンピューティング環境がどのようにインスタンスを起動するかの戦略。３つから選択できる。

**BEST_FIT**  
最もコストの低いインスタンスタイプを使用。追加インスタンスが使用できない場合、使用可能になるまで待機する。

**BEST_FIT_PROGRESSIVE**  
ジョブの要件を満たすのに十分な大きさの追加インスタンスを選択する。より低コストのインスタンスタイプを優先する。

**SPOT_CAPACITY_OPTIMIZED**  
ジョブの要件を満たすのに十分な大きさの追加インスタンスを選択する。中断される可能性が低いインスタンスタイプを優先。スポットを使用する設定の場合のみ使用可能。


[メモリ管理](https://docs.aws.amazon.com/ja_jp/batch/latest/userguide/memory-management.html)

メモリはインスタンスに搭載されている量全てを使用することはできない。
ECS コンテナエージェントは、Docker ReadMemInfo() 関数を使用し、OS で使用可能な合計量を取得している。


[EFA](https://docs.aws.amazon.com/ja_jp/batch/latest/userguide/efa.html)

EFA を使用するすべてのインスタンスは、同じクラスタープレイスメントグループに属している必要がある。
ジョブ定義においては、EFA デバイスがコンテナにパススルーされるように設定する必要がある



## トラブルシューティング

[トラブルシューティング](https://docs.aws.amazon.com/ja_jp/batch/latest/userguide/troubleshooting.html)

**INVALID コンピューティング環境**

マネージド型のコンピューティング環境の設定に不備があると INVALID 状態となる。サービスロール、スポットフリートロールなどの設定不備が原因の場合がある。

**RUNNABLE 状態でジョブが止まる**

ジョブが RUNNING に遷移しない場合。以下の項目が原因となっている場合がある。

* awslogs ログドライバが、コンピューティングリソースで設定されていない。
* リソースが不十分である。
* コンピューティングリソースのインターネットアクセスがない。
* Amazon EC2 インスタンス制限に到達



## ナレッジセンター

* [ナレッジセンター](https://aws.amazon.com/jp/premiumsupport/knowledge-center/#AWS_Batch)

[AWS Batch ジョブが RUNNABLE ステータスで止まっているのはなぜですか?](https://aws.amazon.com/jp/premiumsupport/knowledge-center/batch-job-stuck-runnable-status/)

[AWS Batch で INVALID コンピューティング環境を修正する方法は?](https://aws.amazon.com/jp/premiumsupport/knowledge-center/batch-invalid-compute-environment/)

[マネージド型のコンピューティング環境で Amazon EFS ボリュームを AWS Batch にマウントする方法を教えてください。](https://aws.amazon.com/jp/premiumsupport/knowledge-center/batch-mount-efs/)

* 起動テンプレートのユーザーデータでマウントするように設定する。
* ジョブ定義にて、volumes、mountPoints の設定を行う



## Black Belt

[20190911 AWS Black Belt Online Seminar AWS Batch](https://www.slideshare.net/AmazonWebServicesJapan/20190911-aws-black-belt-online-seminar-aws-batch)

* P49: スポットインスタンスの活用
  * ジョブの再試行を設定可能
  * 工夫としては、一つのジョブを複数に分割する。もしくは途中経過を EFS に出力するチェックポイント方式とする
* P52: コンテナイメージ作成の自動化
  * CodePipeline を使用。CodeCommit にプッシュし、CodeBuild でビルド & ECR へのプッシュを行うパイプラインを構築する
* P53: StepFunctions を使用。トリガーとなる処理の後に Batch を実行したり、Batch の前段で前処理を実行するなど。
* P55: GPU を用いた GPU のエンコード処理
  * [Deploy an 8K HEVC pipeline using Amazon EC2 P3 instances with AWS Batch](https://aws.amazon.com/jp/blogs/compute/deploy-an-8k-hevc-pipeline-using-amazon-ec2-p3-instances-with-aws-batch/)
* P57: マルチノード MPI
  * [Building a tightly coupled molecular dynamics workflow with multi-node parallel jobs in AWS Batch](https://aws.amazon.com/jp/blogs/compute/building-a-tightly-coupled-molecular-dynamics-workflow-with-multi-node-parallel-jobs-in-aws-batch/)
* P58: 機械学習基盤
  * [Scalable deep learning training using multi-node parallel jobs with AWS Batch and Amazon FSx for Lustre](https://aws.amazon.com/jp/blogs/compute/scalable-deep-learning-training-using-multi-node-parallel-jobs-with-aws-batch-and-amazon-fsx-for-lustre/)



# 参考

* [AWS Batch とは](https://docs.aws.amazon.com/ja_jp/batch/latest/userguide/what-is-batch.html)
* [AWS Batch の特徴](https://aws.amazon.com/jp/batch/features/)
* [よくある質問](https://aws.amazon.com/jp/batch/faqs/)
* [ナレッジセンター](https://aws.amazon.com/jp/premiumsupport/knowledge-center/#AWS_Batch)
* Black Belt
  * [20190911 AWS Black Belt Online Seminar AWS Batch](https://www.slideshare.net/AmazonWebServicesJapan/20190911-aws-black-belt-online-seminar-aws-batch)
* [API Reference](https://docs.aws.amazon.com/ja_jp/batch/latest/APIReference/Welcome.html)

