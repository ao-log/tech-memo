

## 特徴

* インスタンス内で管理者権限を使える
* 課金は秒単位
* AWS CLI で制御可能

## リージョン、AZ

[リージョンとアベイラビリティーゾーン](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/using-regions-availability-zones.html)

AZ は us-east-1a のような識別子で表される。どの AZ に紐づくかは AWS アカウントによって異なる。AZ を特定するには use1-az1 のように表される AZ-ID を使用する。

## システム基盤

[Nitro Hypervisor](https://aws.amazon.com/jp/ec2/faqs/#Nitro_Hypervisor)

C5, M5, R5 などの最新インスタンスは EC2 ソフトウェアスタック全体を専用ハードウェアへオフロード。
C4, M4, R4 より前は Xen ベースのハイパーバイザー。

## インスタンスタイプ

[インスタンスタイプ](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/instance-types.html)

* リージョンによって選択できるインスタンスタイプは異なる

#### 命名規則

c5d.xlarge

* c: インスタンスファミリー
* 5: インスタンス世代
* d: 追加機能
* xlarge: インスタンスタイプ

#### ファミリー

* c: コンピューティング最適化
* i, d: ストレージ最適化(SSD, HDD)
* x, m: メモリ最適化
* f, p, g: 高速コンピューティング(FPGA, Tesla V100, Tesla M60)

#### 追加機能

* d: インスタンスストアを付加(NVMe SSD)
* n: ネットワーク強化
* a: AMD の CPU を搭載

#### ベアメタル

命名規則は、「.metal」

#### バースト可能なインスタンス

T2, T3 が該当。CPU クレジットを消費する方式。

#### A1 インスタンス

ARM プロセッサ。マイクロサービス、Web サーバなど多数の小規模インスタンスを使用する用途に最適。

## CPU

* CPU 最適化オプション: 起動時に CPU コア数、ハイパースレッドをオフに指定できるオプション。

## アクセラレータ

* Elastic Graphics: GPU をアタッチ可能
* Elastic Inference: アタッチすることで推論処理を高速化

## メモリ

* ハイバネーション機能: メモリの状態をディスクに書き出して停止可能な機能。一部インスタンスタイプ、OS で対応。

## ネットワーク、セキュリティ

* インスタンスには公開鍵のみ起動時にコピーして配置。
* セキュリティグループでフィルタリング
* ENI はインスタンスによって割り当て可能な数が異なる。
* 拡張ネットワークにより、帯域、レイテンシを改善可能。Intel 82599VF, Elastic Network Adopter、Elastic Fablic Adopter(HPC 向け) がある。

[Elastic Fabric Adapter](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/efa.html)

## ストレージ

### EBS

* EBS 最適化オプションは EBS への IO は専用の帯域で通信できるオプション。
* EBS-Backed インスタンス、Instance store-backed の二種類がある。EBS-Backed が推奨。

https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/RootDeviceStorage.html

### インスタンスストア

揮発性のストレージ。起動時のみアタッチ可能。起動後はアタッチ不可。無料で使える。


### AMI

[Amazon マシンイメージ (AMI)](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/AMIs.html)

#### 概要

* OS 起動に必要なイメージ。
* リージョン単位で保存される。

#### AMI に含まれている情報

* EBS スナップショット
* 特定の AWS アカウントへの起動許可
* インスタンス起動のデバイスマッピング

#### 備考

* Nitro Hypervisor では、ENA, NVMe ドライバがないと OS 起動に失敗する

## インスタンスのライフサイクル

[インスタンスのライフサイクル](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/ec2-instance-lifecycle.html)

リタイアという状態がある。ハード側で回復不可能な障害が検知された場合、予定されたリタイア日に強制的に停止または削除される。


## インスタンスの配置

#### 専有オプション

|項目|説明|
|---|---|
| [ハードウェア専有インスタンス](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/dedicated-instance.html) | インスタンスあたりの課金。 |
| [Dedicated Hosts](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/dedicated-hosts-overview.html) | 物理ホストへのインスタンス配置が可能。ライセンス持ち込み可能。ホストあたりの課金。 |
| [キャパシティ予約](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/ec2-capacity-reservations.html) | 予約した分のリソースがほぼ確実に使える。 |

#### プレイスメントグループ

[プレイスメントグループ](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/placement-groups.html)

|項目|説明|
|---|---|
| Cluster | 密な場所に配置。広帯域、低レイテンシが求められるワークロードに。 |
| Spread | EC2 インスタンスを別々の物理ホストに分散して配置。 |
| パーティションプレイスメントグループ | 複数のインスタンスを一つのパーティションにグループ化し、パーティションごとに分散してインスタンスを配置。 |

## 運用

#### 障害検知、復旧

* ホスト側で回復不可能な障害が検出された場合、インスタンスリタイアが予定される。
* インスタンスの異常は StatusCheckFailed_System、StatusCheckFailed_System で検知される。CloudWatch アラームで「Recover this Instance」アクションを指定することで自動復旧する。

#### 便利機能

|項目|説明|
|---|---|
| [ユーザーデータ](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/user-data.html) | 起動時にスクリプトを実行する。 |
| 起動テンプレート | 起動時の設定をテンプレ化。Auto Scaling などで使用できる。 |
| インスタンスメタデータ | 自インスタンスの情報を採取。次のアドレス「http://169.254.169.254/latest/meta-data」にアクセスすることで採取できる。 |

## ベストプラクティス

[ベストプラクティス](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/ec2-best-practices.html)

#### セキュリティ

* IAM で EC2 インスタンスへのアクセスを制御
* セキュリティグループで制限
* パスワードログインを禁止

#### バックアップ & リストア

* スナップショット、AMI の作成

## トラブルシューティング

[インスタンスのトラブルシューティング](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/ec2-instance-troubleshoot.html)


## 課金

#### 購入オプション

* オンデマンドインスタンス
* リザーブドインスタンス
* スポットインスタンス
* Dedicated Hosts
* ハードウェア専有インスタンス

参考

* [概算見積もりツール](https://calculator.s3.amazonaws.com/index.html?lng=ja_JP)
* [料金ページ](https://aws.amazon.com/jp/ec2/pricing/)


# 参考

* [[AWS Black Belt Online Seminar] Amazon EC2 資料及び QA 公開](https://aws.amazon.com/jp/blogs/news/webinar-bb-amazon-ec2-2019/)
* [EC2 ドキュメント](https://docs.aws.amazon.com/ec2/index.html)
