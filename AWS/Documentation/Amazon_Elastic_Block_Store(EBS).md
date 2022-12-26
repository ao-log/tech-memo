
## 特徴

* AZ ごとに作成する。
* EBS 最適化を有効化すると、他のトラフィックから分離し、EBS 専用の帯域を確保できる(デフォルトで有効化されている場合がほとんど)。
* アタッチ中でも、タイプ、サイズ、IOPS を変更可能。拡張後は、OS からファイルシステム拡張が必要。IOPS 変更は徐々に反映される。変更後は 6 時間は変更不可。
* 暗号化可能。



## EBS ボリュームの概要

[Amazon EBS ボリュームの種類](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/ebs-volume-types.html)

* SSD: 汎用 SSD、プロビジョンド IOPS
* HDD: スループット最適化 HDD(シーケンシャルアクセス、例えば Hadoop に向いている)、コールド HDD(ログやバックアップのアーカイブ先に向いている)

**gp3**

* 3,000 IOPS と 125 MiB/秒 の一貫したベースラインレートを提供
* 追加の IOPS, スループットを追加可能(追加料金が発生)
* プロビジョニングされたボリュームサイズに対するプロビジョンド IOPS の最大比率は、GiB あたり 500 IOPS

**gp2**

* 3,000 IOPS にバースト可能
* 最低 100 IOPS が確保される。
* 1 GiB あたり 3 IOPS ずつ増加。

I/O クレジットおよびバーストパフォーマンス

* 初期 I/O クレジットバランスは 540 万 I/O クレジット(30 分間の 3,000 IOPS)
* ベースライン(1 GiB あたり 3 IOPS) を超える場合にクレジットを消費
* 最大 3,000 IOPS までバースト
* Nitro 世代のインスタンスにアタッチされている場合、バーストバランスは報告されない。

**io1, io2**

* リクエストされたボリュームサイズに対するプロビジョンド IOPS の最大比率 (GiB 単位) は、io1 ボリュームの場合は 50:1、io2 ボリュームの場合は 500:1

**st1**

* スループット最適化 HDD (st1) 

**sc1**

* Cold HDD (sc1) 


[Linux で Amazon EBS ボリュームを使用できるようにする](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/ebs-using-volumes.html)

フォーマット、マウントの手順。


[ボリュームのステータスのモニタリング](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/monitoring-volume-status.html)

以下のステータスがある。

* ok
* warning
* impaired
* insufficient-data



## スナップショット

[Amazon EBS スナップショット](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/EBSSnapshots.html)

* スナップショットは S3 上に保存される
* 増分バックアップ
* 共有可能
* 暗号化可能


[Amazon EBS スナップショットの作成](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/ebs-creating-snapshot.html)

* ルートデバイスとして機能する EBS ボリュームのスナップショットを作成する場合は、スナップショットを取る前にインスタンスを停止(することを推奨)
* 休止が有効にされているインスタンスからスナップショットを作成することはできない


[Amazon EBS スナップショットのコピー](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/ebs-copy-snapshot.html)

* スナップショットはコピー可能。コピー先を別リージョンとすることも可能。


[スナップショットライフサイクルの自動化](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/automating-snapshots.html)

* Amazon Data Lifecycle Manager を使用して、Amazon EBS ボリュームをバックアップするスナップショットの作成、保持、削除を自動化可能



## 機能

[Amazon EBS 高速スナップショット復元](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/ebs-fast-snapshot-restore.html)

* この機能を使うには Amazon EBS 高速スナップショット復元を有効にする必要あり。
* スナップショットからボリュームを作成時に、ボリュームは作成時に完全に初期化された状態となる。ファーストタッチペナルティを受けなくなる。


[Amazon EBS 最適化インスタンス](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/ebs-optimized.html)

* EBS 専用の帯域を確保する。



## パフォーマンス

[I/O の特性とモニタリング](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/ebs-io-characteristics.html)

重要なメトリクスは以下のもの。

* BurstBalance
* VolumeReadBytes
* VolumeWriteBytes
* VolumeReadOps
* VolumeWriteOps
* VolumeQueueLength


[Amazon EBS の Amazon CloudWatch メトリクス](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/using_cloudwatch_ebs.html)



## インスタンスストア

[Amazon EC2 インスタンスストア](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/InstanceStorage.html)



## その他

[ブロックデバイスマッピング](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html)



## BlackBelt

[AWS Black Belt Online Seminar Amazon EBS 資料及び QA 公開](https://aws.amazon.com/jp/blogs/news/webinar-bb-amazon-ebs-2019/)

* 概要
  * P8: 99.999% の可用性。
  * P9: 容量は 1 GiB 単位で設定。AZ ごとに独立。同一 AZ 内のインスタンスからのみ利用可能。
  * P11: S3 にスナップショットを作成し、任意の AZ に復元可能。
  * P12: ボリュームのデータは AZ 内で複数のハードウェアにレプリケートされている。
  * P13: EBS 最適化インスタンスでは EBS 用に独立した帯域を確保する。
  * P14: Nitro 世代のインスタンスは EBS への書き込みや暗号化をハードウェアにオフロードしている。
  * P16: ボリュームタイプ: gp2, io1, st1, sc1
  * P20-26:
    * gp2: 1 GiB あたり 3 IOPS ずつ性能が上がる。ただし、最小 100 IOPS は保証される。3,000 IOPS までバースト可能。
    * BurstBalance のメトリクスで確認可能。0 になると枯渇。
  * P27:
    * io1: IOPS 値を指定可能。
  * P28:
    * st1: スループットを要求するビッグデータ処理に最適。
  * P32: 256 KiB までの連続したアクセスを 1 IOPS としてカウント。例えば 32 KiB のランダムな 8 回のアクセスは 8 IOPS としてカウントされる。
* パフォーマンス
  * P33: 律速する要素は３つ。EC2 インスタンスのスループット。EBS ボリュームの IOPS。EBS ボリュームのスループット。
* 監視
  * P52: OS のリソース確認、CloudWatch メトリクス 
* NVMe SSD
  * P55: Nitro 世代のインスタンスでは NVMe デバイスとして EBS ボリュームが認識される。NVMe ドライバを使用して PCI バスをスキャンしてアタッチされた EBS を検出。
* EBS の機能
  * P59: EC2 にアタッチ中もサイズ、IOPS、ボリュームタイプを変更可能。サイズの縮小はできない。
  * P60: IOPS の変更は徐々に反映される。一度変更すると 6 時間は変更不可。
  * P62: スナップショット。作成時は静止点を設けることを推奨。
  * P63: スナップショットの作成指示をした時点のデータがバックアップされる。
  * P65: フルバックアップ後は増分がバックアップされる。
  * P69: リージョン間でスナップショットをコピーできる。
  * P72: スナップショットの作成方法(AWS CLI, AWS SDK, マネジメントコンソール, Systems Manager, Amazon Data Lifecycle Manager, AWS Backup)
  * P79: 暗号化はハードウェアで行われる。そのため、パフォーマンスへの影響は極めて小さい。



# 参考

* Document
  * [EBS ドキュメント](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/AmazonEBS.html)
* サービス紹介ページ
  * [Amazon Elastic Block Store](https://aws.amazon.com/jp/ebs/)
  * [よくある質問](https://aws.amazon.com/jp/ebs/faqs/)
* Black Belt
  * [AWS Black Belt Online Seminar Amazon EBS 資料及び QA 公開](https://aws.amazon.com/jp/blogs/news/webinar-bb-amazon-ebs-2019/)


