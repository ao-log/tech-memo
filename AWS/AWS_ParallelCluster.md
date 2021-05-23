
# AWS ParallelCluster

## セットアップ

[AWS ParallelCluster のセットアップ](https://docs.aws.amazon.com/ja_jp/parallelcluster/latest/ug/getting_started.html)

[AWS ParallelCluster のインストール](https://docs.aws.amazon.com/ja_jp/parallelcluster/latest/ug/install.html)

pip でインストール可能。

```
python3 -m pip install --upgrade aws-parallelcluster
```


[AWS ParallelCluster の設定](https://docs.aws.amazon.com/ja_jp/parallelcluster/latest/ug/getting-started-configuring-parallelcluster.html)

```
// 対話型で ParallelCluster の設定ファイルを作成
pcluster configure

// クラスターを作成
pcluster create mycluster
```


[ベストプラクティス](https://docs.aws.amazon.com/ja_jp/parallelcluster/latest/ug/best-practices.html)

**マスターノード**

* クラスターサイズが大きい場合は、マスターノードにもコンピュートリソースが多めに必要
* NFS サーバとしても使用するので、ワークフローの処理に十分なネットワーク帯域幅、EBS 帯域幅を持つインスタンスが必要

**ネットワークパフォーマンス**

* プレイスメントグループを使用
* 拡張ネットワーキングをサポートするインスタンスタイプを使用
* 帯域幅のあるインスタンスタイプを使用



## ParallelCluster を使用する

### [pcluster](https://docs.aws.amazon.com/ja_jp/parallelcluster/latest/ug/pcluster.html)

* [pcluster configure](https://docs.aws.amazon.com/ja_jp/parallelcluster/latest/ug/pcluster.configure.html) ParallelCluster の設定
* [pcluster create](https://docs.aws.amazon.com/ja_jp/parallelcluster/latest/ug/pluster.create.html) クラスターを作成
* [pcluster list](https://docs.aws.amazon.com/ja_jp/parallelcluster/latest/ug/pcluster.list.html) スタックの一覧
* [pcluster update](https://docs.aws.amazon.com/ja_jp/parallelcluster/latest/ug/pcluster.update.html) クラスターの更新
* [pcluster delete](https://docs.aws.amazon.com/ja_jp/parallelcluster/latest/ug/pcluster.delete.html) クラスターの削除
* [pcluster instances](https://docs.aws.amazon.com/ja_jp/parallelcluster/latest/ug/pcluster.instances.html) クラスター内のインスタンスを一覧
* [pcluster start](https://docs.aws.amazon.com/ja_jp/parallelcluster/latest/ug/pcluster.start.html) 停止されているコンピュートノードの開始
* [pcluster status](https://docs.aws.amazon.com/ja_jp/parallelcluster/latest/ug/pcluster.status.html) クラスターの現在の状態を表示
* [pcluster stop](https://docs.aws.amazon.com/ja_jp/parallelcluster/latest/ug/pcluster.stop.html) コンピュートノードを停止
* [pcluster ssh](https://docs.aws.amazon.com/ja_jp/parallelcluster/latest/ug/pcluster.ssh.html) クラスターへの SSH 接続
* [pcluster createami](https://docs.aws.amazon.com/ja_jp/parallelcluster/latest/ug/pcluster.createami.html) カスタム AMI を作成
* [pcluster dcv](https://docs.aws.amazon.com/ja_jp/parallelcluster/latest/ug/pcluster.dcv.html) ヘッドノードの NICE DCV サーバとやり取り
* [pcluster version](https://docs.aws.amazon.com/ja_jp/parallelcluster/latest/ug/pcluster.version.html) ParallelCluster のバージョンを表示


[カスタムブートストラップアクション](https://docs.aws.amazon.com/ja_jp/parallelcluster/latest/ug/pre_post_install.html)

メインブートストラップアクションの前後で任意のスクリプトを実行することが可能。
マスターノード、コンピュートノードで処理を分岐できる。


[Amazon S3 の使用](https://docs.aws.amazon.com/ja_jp/parallelcluster/latest/ug/s3_resources.html)


[スポットインスタンスの使用](https://docs.aws.amazon.com/ja_jp/parallelcluster/latest/ug/spot.html)

cluster_type = spot に設定されている場合、スポットインスタンスを使用する。


[AWS ParallelCluster の IAM ロール](https://docs.aws.amazon.com/ja_jp/parallelcluster/latest/ug/iam.html)

クラスターの作成に必要な権限がまとめられている。


[サポートされているスケジューラ](https://docs.aws.amazon.com/ja_jp/parallelcluster/latest/ug/schedulers.html)

ParallelCluster における SGE および Torque のサポートは 2021年12月31日以降は中止される。


[Slurm 複数キューモードの ガイド](https://docs.aws.amazon.com/ja_jp/parallelcluster/latest/ug/multiple-queue-mode-slurm-user-guide.html)

ParallelCluster 2.9.0 では複数のキューモードと新しいスケーリングアーキテクチャが導入されている。

ノード

* POWER_SAVING 状態: idle~ のように表示される。ジョブを投入すると割り当てられたノードは POWERUP 状態へと遷移する。
* POWER_UP 状態: idle# のように表示される。
* 利用可能となったノードは idle のように表示される。
* POWER_DOWN　状態: idle% のように表示される。Dynamic Node は scaledown_idletime で設定した時間の後にこの状態に遷移する。
* 疎通が取れないノードは offline 状態となり down* のように表示される。

ジョブ

* CF, CONFIGURING: POWER_SAVING 状態のときに投入されたジョブの状態
* R, RUNNING: 実行中

キュー

* デフォルトキューには * がついている。

次のようにインスタンスタイプ指定の投入も可能。
```
sbatch --wrap "sleep 300" -p ondemand -N 10 -C "[c5.2xlarge*8&t2.xlarge*2]"
```

パーティションは次の二つの状態がある。

* INACTIVE
* UP


[AWS Batch](https://docs.aws.amazon.com/parallelcluster/latest/ug/awsbatchcli.html)

AWS Batch を使用する場合は各種コマンドがヘッドノードにインストールされる。awsbsub コマンドなど。


[Multiple queue mode](https://docs.aws.amazon.com/parallelcluster/latest/ug/queue-mode.html)

AWS ParallelCluster 2.9.0 からマルチキューモードが導入された。
Slurm を使用し、queue_settings が設定されているときにサポートされる。

[compute_resource] は Static Node と Dynamic Node に分けられる。
Static Node の台数は 1 〜 min_count。Dynamic Node の台数は 1 〜 (max_count - min_count)

ホスト名は以下の規則となる。

$HOSTNAME=$QUEUE-$STATDYN-$INSTANCE_TYPE-$NODENUM

ホスト名、FQDN は Route 53 ホストゾーン $HOSTNAME.$CLUSTERNAME.pcluster で作成される。


[Integration with Amazon CloudWatch Logs](https://docs.aws.amazon.com/parallelcluster/latest/ug/cloudwatch-logs.html)

バージョン 2.6.0 以降はデフォルトで主要ログが CloudWatch Logs に送信されるように設定されている。
バージョン 2.10.0 では CloudWatch Dashboard が作られる。


[Elastic Fabric Adapter](https://docs.aws.amazon.com/parallelcluster/latest/ug/efa.html)

Elastic Fabric Adapter (EFA) は低レイテンシで同一サブネットのインスタンス間のネットワーク疎通を行う能力を持った OS をバイパスするネットワークデバイスである。

[cluster] セクションで enable_efa = compute と記述することで設定可能。
一部のインスタンスタイプ、OS でサポートされている。


[Enable Intel MPI](https://docs.aws.amazon.com/parallelcluster/latest/ug/intelmpi.html)

デフォルトでは Open MPI を使用する設定になっている。
Intel MPI を使用するには ```module load intelmpi``` を実行する。


[Intel HPC Platform Specification](https://docs.aws.amazon.com/parallelcluster/latest/ug/intel-hpc-platform-specification.html)

AWS ParallelCluster は Intel HPC Platform Specification に準拠している。
このページに OS、インスタンスタイプ、ディスクサイズなどの要件がまとめられている。


[Arm Performance Libraries](https://docs.aws.amazon.com/parallelcluster/latest/ug/arm-performance-libraries.html)

AWS ParallelCluster 2.10.1 から Arm Performance Libraries が使用できるようになった。


[Connect to the head node through NICE DCV](https://docs.aws.amazon.com/parallelcluster/latest/ug/dcv.html)

NICE DCV はリモートのサーバに接続して使用する可視化テクノロジーである。
dcv_settings を設定することで使用可能となる。


[Using pcluster update](https://docs.aws.amazon.com/parallelcluster/latest/ug/using-pcluster-update.html)



## Configuration

config ファイルの設定方法。

* ~/.parallelcluster/config
* -c or --config command line option
* AWS_PCLUSTER_CONFIG_FILE 

ドキュメント上にサンプルがある。
[examples](https://docs.aws.amazon.com/parallelcluster/latest/ug/examples.html)

サンプルはこちらにもある。  
https://github.com/aws/aws-parallelcluster/blob/v2.10.3/cli/src/pcluster/examples/config

[grobal], [aws] セクションは必須。
[cluster], [vpc] セクションのうち、どちらか片方は必要。

#### [grobal] セクション

* cluster_template: どの [cluster] セクションを使用するか
* update_check: 
* sanity_check: 設定のバリデーションを行う

#### [aws] セクション

次の優先順位でリージョンを決定。

* pcluster の -r or --region parameter
* 環境変数 AWS_DEFAULT_REGION
* 設定ファイルの [aws] セクションの設定
* ~/.aws/config の [default] セクションの設定

#### [aliases] セクション

ssh コマンドのカスタマイズ。

#### [cluster] セクション

* additional_cfn_template
* additional_iam_policies
* base_os: alinux, alinux2, centos7, centos8, ubuntu1604, ubuntu1804
* cluster_resource_bucket
* cluster_type: ondemand, spot
* compute_instance_type
* compute_root_volume_size
* custom_ami
* cw_log_settings: [cw_log] セクションの指定
* dashboard_settings: [dashboard] セクションの指定
* dcv_settings: [dcv] の指定
* desired_vcpus: スケジューラが awsbatch の場合のみ
* disable_cluster_dns: デフォルトではホストゾーンが作られるが、disable にすることで作られない
* disable_hyperthreading: head, compute ノードでハイパースレッディングを無効化。
* ebs_settings: [ebs] セクションの指定
* ec2_iam_role: additional_iam_policies を推奨
* efs_settings: [efs] セクションの指定
* enable_efa: compute ノードの EFA 有効可。
* enable_efa_gdr: compute ノードの Elastic Fabric Adapter (EFA) support for GPUDirect RDMA 有効可。
* enable_intel_hpc_platform: Intel Parallel Studio がインストールされる。
* encrypted_ephemeral
* ephemeral_dir
* extra_json
* fsx_settings: [fsx] セクションの指定
* iam_lambda_role
* initial_queue_size: compute ノードの初期台数。queue_settings がある場合は [compute_resource] セクションで指定すること。
* key_name: EC2 インスタンスのキーペア名
* maintain_initial_size
* master_instance_type
* master_root_volume_size
* max_queue_size: compute ノードの最大台数。queue_settings がある場合は [compute_resource] セクションで指定すること。
* max_vcpus: スケジューラが awsbatch の場合のみ
* min_vcpus: スケジューラが awsbatch の場合のみ
* placement: cluster or compute。queue_settings がある場合は [queue] セクションで指定すること。
* placement_group: DYNAMIC or An existing Amazon EC2 cluster placement group name
* post_install
* post_install_args
* pre_install
* pre_install_args
* proxy_server
* queue_settings: [queue] セクションの指定
* raid_settings
* s3_read_resource
* s3_read_write_resource
* scaling_settings: [scaling] セクションの指定
* scheduler: awsbatch, sge, slurm, torque
* shared_dir
* spot_bid_percentage
* spot_price
* tags
* template_url
* vpc_settings: [vpc] セクションの指定

#### [compute_resource] セクション

scheduler が slurm に設定されているときのみ有効。

* initial_count
* instance_type
* max_count
* min_count
* spot_price

#### [cw_log] セクション

* enable
* retention_days

#### [dashboard] セクション

* enable

#### [dcv] セクション

* access_from
* enable
* port

#### [ebs] セクション

head ノードのボリューム。これは NFS で compute ノードにもシェアされる。

* shared_dir
* ebs_kms_key_id
* ebs_snapshot_id
* ebs_volume_id
* encrypted
* volume_iops
* volume_size
* volume_throughput
* volume_type

#### [efs] セクション

* efs_fs_id
* efs_kms_key_id
* encrypted
* performance_mode
* provisioned_throughput
* shared_dir
* throughput_mode

#### [fsx] セクション

* auto_import_policy
* automatic_backup_retention_days
* copy_tags_to_backups
* daily_automatic_backup_start_time
* deployment_type
* drive_cache_type
* export_path
* fsx_backup_id
* fsx_fs_id
* fsx_kms_key_id
* import_path
* imported_file_chunk_size
* per_unit_storage_throughput
* shared_dir
* storage_capacity
* storage_type
* weekly_maintenance_start_time

#### [queue] セクション

* compute_resource_settings: [compute_resource] セクションを指定
* compute_type: ondemand or spot
* disable_hyperthreading: default value is false
* enable_efa
* enable_efa_gdr
* placement_group

#### [raid] セクション

* shared_dir
* ebs_kms_key_id
* encrypted
* num_of_raid_volumes
* raid_type
* volume_iops
* volume_size
* volume_throughput
* volume_type

#### [scaling] セクション

* scaledown_idletime

#### [vpc] セクション

* additional_sg
* compute_subnet_cidr
* compute_subnet_id
* master_subnet_id
* ssh_from
* use_public_ips
* vpc_id
* vpc_security_group_id



## How AWS ParallelCluster works

[AWS ParallelCluster processes](https://docs.aws.amazon.com/parallelcluster/latest/ug/processes.html)

コンピュートノードのプロビジョニング、Auto Scaling Group との連携が自動で行われるようになっている。

ライフサイクルとしては、クラスターの作成は CloudFormation によって行われ、クラスターの作成後にジョブを投入できるようになる。sqswatcher, jobwatcher, nodewatcher が動作している。

jobwatcher が毎分スケジューラをモニタリングしスケールアップするかどうかを評価する。

sqswatcher は Auto Scaling によって通知された SQS のメッセージをモニタリングしている。

node_watcher は各コンピュートノードで動作している。scaledown_idletime 経過後にインスタンスを終了する。


[AWS services used by AWS ParallelCluster](https://docs.aws.amazon.com/parallelcluster/latest/ug/aws-services.html)

以下のサービスが使用されている。

* AWS Auto Scaling
* AWS Batch
* AWS CloudFormation
* Amazon CloudWatch
* Amazon CloudWatch Logs
* AWS CodeBuild
* Amazon DynamoDB
* Amazon Elastic Block Store
* Amazon Elastic Compute Cloud
* Amazon Elastic Container Registry
* Amazon EFS
* Amazon FSx for Lustre
* AWS Identity and Access Management
* AWS Lambda
* NICE DCV
* Amazon Route 53
* Amazon Simple Notification Service
* Amazon Simple Queue Service
* Amazon Simple Storage Service
* Amazon VPC


[AWS ParallelCluster Auto Scaling](https://docs.aws.amazon.com/parallelcluster/latest/ug/autoscaling.html)

SGE, Torque の場合の Auto Scaling について書かれているページ。



## Tutorials

[Running your first job on AWS ParallelCluster](https://docs.aws.amazon.com/parallelcluster/latest/ug/tutorials_01_hello_world.html)

SGE の場合のクラスター作成からジョブ実行までの流れ。


[Building a Custom AWS ParallelCluster AMI](https://docs.aws.amazon.com/parallelcluster/latest/ug/tutorials_02_ami_customization.html)

カスタム AMI の作り方の流れ。なお、カスタム AMI は AWS によるアップデートが反映されないため非推奨。


[Running an MPI job with AWS ParallelCluster and awsbatch scheduler](https://docs.aws.amazon.com/parallelcluster/latest/ug/tutorials_03_batch_mpi.html)

AWS Batch でのクラスター作成からジョブ実行までの流れ。
Batch のコマンドは ParallelCluster をインストールすることで利用可能となる。
一方でチュートリアルではヘッドノードにログインして作業する。
Batch ジョブが実行されるリソースと NFS ボリュームが共有されているため。

awsbub でジョブ投入する。
awsbstat でジョブの状態を確認する。
awsbout でジョブの出力を確認する。


[Disk encryption with a custom KMS Key](https://docs.aws.amazon.com/parallelcluster/latest/ug/tutorials_04_encrypted_kms_fs.html)

ebs_kms_key_id, fsx_kms_key_id により EBS, FSx を CMK により暗号化可能。


[Multiple queue mode tutorial](https://docs.aws.amazon.com/parallelcluster/latest/ug/tutorial-mqm.html)

複数のキューを設定可能。また、キュー内で複数の compute resource を設定可能。

```
[cluster multi-queue]
key_name = <Your SSH key name>
base_os = alinux2                   # optional, defaults to alinux2
scheduler = slurm
master_instance_type = c5.xlarge    # optional, defaults to t2.micro
vpc_settings = <Your VPC section>
scaling_settings = demo
queue_settings = spot,ondemand

[queue spot]
compute_resource_settings = spot_i1,spot_i2
compute_type = spot                 # optional, defaults to ondemand

[compute_resource spot_i1]
instance_type = c5.xlarge
min_count = 0                       # optional, defaults to 0
max_count = 10                      # optional, defaults to 10

[compute_resource spot_i2]
instance_type = t2.micro
min_count = 1
initial_count = 2

[queue ondemand]
compute_resource_settings = ondemand_i1
disable_hyperthreading = true       # optional, defaults to false

[compute_resource ondemand_i1]
instance_type = c5.2xlarge
```

sbatch コマンドでジョブを投入する。
squeue コマンドでジョブ実行状況を確認。
sinfo コマンドでノードの状況を確認。

インスタンスタイプを指定した投入も可能。

```
sbatch -N 3 -p spot -C "[c5.xlarge*1&t2.micro*2]" --wrap "srun hellojob.sh"
```



## AWS ParallelCluster troubleshooting

#### ログの保全

クラスター削除時にログを残す場合、次のコマンドオプションで対応可能。

```
pcluster delete —keep-logs <cluster_name> 
```

#### スタック作成失敗の切り分け

クラスター作成に失敗する場合、次のオプションでロールバックを実行しないようにできる。

```
pcluster create mycluster --norollback
```

切り分けに有用なログ。

* /var/log/cfn-init.log 
* /var/log/cloud-init.log
* /var/log/cloud-init-output.log

#### 複数キューモードの切り分け

切り分けに有用なログ。

* /var/log/cfn-init.log
* /var/log/chef-client.log
* /var/log/parallelcluster/slurm_resume.log
* /var/log/parallelcluster/slurm_suspend.log
* /var/log/parallelcluster/clustermgtd
* /var/log/slurmctld.log
* /var/log/cloud-init-output.log
* /var/log/parallelcluster/computemgtd
* /var/log/slurmd.log



## BlackBelt

[20200408 AWS Black Belt Online Seminar AWS ParallelCluster ではじめるクラウドHPC](https://www.slideshare.net/AmazonWebServicesJapan/20200408-aws-black-belt-online-seminar-aws-parallelcluster-hpc)

* HPC on AWS
  * 従量課金で利用可能
  * P19: Elastic Fabric Adapter の登場により MPI の密結合なワークフローも高いパフォーマンスを発揮
* AWS ParallelCluster とは
  * P26: 自前で HPC クラスタを AWS に組むのは大変。それを解決できるのが ParallelCluster。
  * P31-32: 各サービスの使い分け。ParallelCluster は既存の HPC クラスターからの移行が容易。
* AWS ParallelCluster の設定
  * P44: config ファイルの読み方
* パフォーマンス
  * P53: コンピュートの選択
  * P54: ハイパースレッディングの無効化
  * P55: ストレージの選択
  * P56: マスターノードの NFS(EBS, マスターノードのインスタンスタイプのネットワーク帯域)
  * P58: EFS
  * P59: インスタンスストアの活用
  * P60: FSx for Lustre
  * P64: ネットワークパフォーマンス(クラスタープレイスメントグループ、EFA)
* オペレーション
  * P72: NICE-DCV
  * P73: Ganglia
  * P74: リソースのタグ付け
  * P75: カスタムブートストラップ、カスタム AMI
  * P76: CloudWatch Logs
  * P77: IAM
* 他のサービスとの連携
  * P80: AWS Backup による EBS, EFS バックアップ
  * P81: AWS Directory Service によるマルチユーザ対応
  * P82: API Gateway によるジョブ投入口の API 化
  * P83: アカウンティング情報の RDS への保存



# 参考

* Document
  * [AWS ParallelCluster とは?](https://docs.aws.amazon.com/ja_jp/parallelcluster/latest/ug/what-is-aws-parallelcluster.html)
* サービス紹介ページ
  * [AWS ParallelCluster](https://aws.amazon.com/jp/hpc/parallelcluster/)
  * [よくある質問](https://aws.amazon.com/jp/hpc/parallelcluster/faqs/)
* Black Belt
  * [20200408 AWS Black Belt Online Seminar AWS ParallelCluster ではじめるクラウドHPC](https://www.slideshare.net/AmazonWebServicesJapan/20200408-aws-black-belt-online-seminar-aws-parallelcluster-hpc)

