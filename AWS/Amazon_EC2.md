# EC2

## 特徴

* インスタンス内で管理者権限を使える
* 課金は秒単位
* AWS CLI で制御可能


[リージョンとアベイラビリティーゾーン](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/using-regions-availability-zones.html)

AZ は us-east-1a のような識別子で表される。どの AZ に紐づくかは AWS アカウントによって異なる。AZ を特定するには use1-az1 のように表される AZ-ID を使用する。


[Amazon EC2 のベストプラクティス](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/ec2-best-practices.html)

* セキュリティ
  * IAM ユーザー、ID フェデレーションを使用して API へのアクセスを管理。
  * セキュリティグループのルールの最小化
  * 定期的に OS、アプリケーションのパッチ適用
* ストレージ
  * インスタンス削除時に残す設定
  * 暗号化
* バックアップと復旧
  * 定期的にバックアップ
  * フェイルオーバーの試験



## AMI

AMI はリージョン単位に作成される。

AMI には以下の情報が含まれている。

* EBS スナップショット
* 特定の AWS アカウントへの起動許可
* インスタンスにアタッチするボリュームを指定するブロックデバイスマッピング


[仮想化タイプ](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/virtualization_types.html)

HVM, PV の 2 種類がある。


[Amazon Linux](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/amazon-linux-ami-basics.html)

Amazon Linux は、起動時にクリティカルまたは重要なセキュリティ更新をダウンロードおよびインストールするよう設定されている。
cloud-init でどの更新を適用するかを設定できる。

```
#cloud-config
repo_upgrade: security
```

Amazon Linux 2 では、Extras Library を使用してアプリケーションおよびソフトウェア更新をインスタンスにインストールできる。

cloud-init パッケージは、起動時に以下のタスクを実行。

* デフォルトのロケールを設定
* ホスト名を設定
* ユーザーデータの解析と処理
* ホスト プライベート SSH キーの生成
* パブリック SSH キーを .ssh/authorized_keys に追加
* パッケージ管理のためにリポジトリを準備
* ユーザーデータで定義されたパッケージアクションの処理
* ユーザーデータにあるユーザースクリプトの実行
* インスタンスストアボリュームをマウント

cloud-init には複数の形式をサポートしている。

単にスクリプトを実行したい場合: 「#!」または「Content-Type: text/x-shellscript」


[カーネルライブパッチ](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/al2-live-patching.html)

実行中のアプリケーションを再起動や中断せずに、実行中の Linux カーネルにセキュリティの脆弱性や重大なバグのパッチを適用することができるもの。



## インスタンス

#### インスタンスタイプ

[インスタンスタイプ](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/instance-types.html)

* リージョンによって選択できるインスタンスタイプは異なる
* 現行世代は C4, C5 など。

[Nitro Hypervisor](https://aws.amazon.com/jp/ec2/faqs/#Nitro_Hypervisor)

C5, M5, R5 などの最新インスタンスは EC2 ソフトウェアスタック全体を専用ハードウェアへオフロード。
C4, M4, R4 より前は Xen ベースのハイパーバイザー。

**ネットワーキング**

* IPv6 はすべての現行世代、一部旧世代のインスタンスタイプでサポート。
* MTU はすべての現行世代で 9001 をサポート。

**命名規則**

c5d.xlarge

* c: インスタンスファミリー
* 5: インスタンス世代
* d: 追加機能
* xlarge: インスタンスタイプ

**ファミリー**

* c: コンピューティング最適化
* i, d: ストレージ最適化(SSD, HDD)
* x, m: メモリ最適化
* f, p, g: 高速コンピューティング(FPGA, Tesla V100, Tesla M60)

**追加機能**

* d: インスタンスストアを付加(NVMe SSD)
* n: ネットワーク強化
* a: AMD の CPU を搭載

**ベアメタル**

命名規則は、「.metal」


[バーストパフォーマンスインスタンス](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/burstable-performance-instances.html)

[バーストパフォーマンスインスタンスの CPU クレジットとベースライン使用率](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/burstable-credits-baseline-concepts.html)

* ベースラインレベルの CPU 使用率を超える場合はバーストする。CPU クレジットが枯渇している場合はバーストできなくなる。
* ベースラインレベルに満たない場合は CPU クレジットが蓄積していく。
* T2 タイプの場合は起動時に起動クレジットを獲得する。


[Linux 高速コンピューティングインスタンス](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/accelerated-computing-instances.html)

* GPU インスタンス
* AWS Inference
* FPGA インスタンス


[インスタンスタイプの変更](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/ec2-instance-resize.html)

* インスタンスの設定と互換性のあるインスタンスタイプを選ぶ必要がある。
* インスタンスタイプを変更しても、インスタンス ID は変更されない。
* Nitro Hypervisor では、ENA, NVMe ドライバがないと OS 起動に失敗する。


#### 専有オプション

* [ハードウェア専有インスタンス](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/dedicated-instance.html): インスタンスあたりの課金。
* [Dedicated Hosts](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/dedicated-hosts-overview.html): 物理ホストへのインスタンス配置が可能。ライセンス持ち込み可能。ホストあたりの課金。
* [キャパシティ予約](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/ec2-capacity-reservations.html): 予約した分のリソースがほぼ確実に使える。


#### インスタンスのライフサイクル

[Linux インスタンスの休止](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/Hibernate.html)

休止すると、インスタンスメモリの内容が EBS ルートボリュームに保存される。


#### インスタンスの設定

[Amazon Linux インスタンスでのソフトウェアの管理](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/managing-software.html)

パッケージの更新、リポジトリの追加方法など。

[Amazon Linux インスタンスでのユーザーアカウントの管理](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/managing-users.html)

OS ごとのユーザ名の対応表などが載っている。


[EC2 インスタンスのプロセッサのステート制御](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/processor_state_control.html)


[Linux インスタンスの時刻の設定](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/set-time.html)

Amazon Time Sync Service は、VPC で実行されているすべてのインスタンスの 169.254.169.123 IPアドレスで NTP を介して利用可能。


[Amazon Linux インスタンスのホスト名の変更](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/set-hostname.html)

Amazon Linux 2 の場合は hostnamectl コマンドを使用して変更可能。


[起動時に Linux インスタンスでコマンドを実行する](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/user-data.html)

* シェルスクリプトの形式で実行したい場合は、ユーザーデータに #! ではじまる記述をするだけでよい。
* ログは /var/log/cloud-init-output.log にコンソール出力がキャプチャされる。
* ユーザーデータスクリプトを処理すると、/var/lib/cloud/instances/instance-id/ にコピーされ実行される。AMI 作成時は、こちらを削除すること。


[インスタンスメタデータとユーザーデータ](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/ec2-instance-metadata.html)

次の２つの方法がある。IMDSv2 のみを有効にするように構成することもできる。

* インスタンスメタデータサービスバージョン 1 (IMDSv1) – リクエスト/レスポンスメソッド
* インスタンスメタデータサービスバージョン 2 (IMDSv2) – セッション志向メソッド



## モニタリング

[Amazon EC2 のモニタリング](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/monitoring_ec2.html)

モニタリングの設計指針について書かれている。
モニタリングの目的、計画作成など。


[インスタンスのステータスチェック](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/monitoring-system-instance-status-check.html)

[System Status Checks] の CloudWatch メトリクスでは EC2 インスタンスの正常性を監視できる。
システムステータスチェック(基板側のチェック)とインスタンスステータスチェック(インスタンス側のチェック)の 2 種類がある。


[インスタンスの詳細モニタリングを有効または無効にする](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/using-cloudwatch-new.html)

基本モニタリングは 5 分間隔。詳細モニタリングは 1 分間隔。ただし、システムステータスチェックは常に 1 分間隔。


[インスタンスを停止、終了、再起動、または復旧するアラームを作成する](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/UsingAlarmActions.html)

CloudWatch アラームのアクションにより、アラーム状態に遷移したときにインスタンスを停止、終了、再起動、復旧のアクションを行うことができる。



## ネットワーク

[Amazon EC2 インスタンスの IP アドレス指定](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/using-instance-addressing.html)

* プライベート IPv4 アドレスは ENI に関連付けられている。よって、インスタンスを停止しても関連付けられたままとなる。
* サブネットにはパブリック IP アドレスの設定がある(デフォルト以外のサブネットではこの属性のデフォルト値は false)。起動時にも有効、無効を設定できる。
* パブリック IP アドレスは、インスタンス停止時にパブリック IP アドレスプールに戻され、再利用することはできない。


[Elastic IP アドレス](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html)


[Elastic Network Interface](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/using-eni.html)

* サービスによって作られる ENI もある。リクエスタマネージド型のネットワークインターフェイス。
* プライマリネットワークインターフェイスはデタッチできない。
* 送信元/送信先チェックを無効にすると、自身にアドレス指定されていないネットワークトラフィックを処理することが可能になる。


[Linux の拡張ネットワーキング](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/enhanced-networking.html)

* SR-IOV を使用することで、高性能なネットワーキング機能を提供している。


[Elastic Fabric Adapter](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/efa.html)

* TCP トランスポートよりも低く、一貫性の高いレイテンシーを提供し、高いスループットが得られる。
* EFA は、Libfabric 1.11.1 と統合されており、HPC アプリケーション向けに Open MPI 4.0.5 および Intel MPI 2019 Update 7、機械学習アプリケーション向けに Nvidia Collective Communications Library (NCCL) をサポート。


[プレイスメントグループ](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/placement-groups.html)

* Cluster: 密な場所に配置。広帯域、低レイテンシが求められるワークロードに。
* Spread: EC2 インスタンスを別々の物理ホストに分散して配置。
* パーティションプレイスメントグループ: 複数のインスタンスを一つのパーティションにグループ化し、パーティションごとに分散してインスタンスを配置。


[ネットワーク MTU](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/network_mtu.html)



## セキュリティ

[プライベートキーを紛失した場合の Linux インスタンスへの接続](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/replacing-lost-key-pair.html)



## タグ

[リソースとタグ](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/EC2_Resources.html)



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





## トラブルシューティング

[インスタンスのトラブルシューティング](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/ec2-instance-troubleshoot.html)

エラー原因は色々とある。

* リージョン内のキャパシティ不足
* 接続タイムアウト(セキュリティグループ、ルートテーブル、ネットワーク ACL、OS 内、)
など。

なお、設定ミスでログインできなくなった際は、他のインスタンスで EBS をアタッチして設定を修正することで復旧を試みることができる。



## BlackBelt

[20190305 AWS Black Belt Online Seminar Amazon EC2](https://www.slideshare.net/AmazonWebServicesJapan/20190305-aws-black-belt-online-seminar-amazon-ec2)

* 概要
  * P10: Nitro System。ソフトウェアスタックは専用のハードウェアにオフロードされている。KVM ベース。
  * P13: プロセッサは Intel, AMD, ARM がある。
* インスタンスの種類
  * P19: インスタンスタイプの命名規則。
    * c5d.xlarge:
      * c: インスタンスファミリー
      * 5: インスタンス世代
      * d: 追加機能
      * xlarge: インスタンスサイズ
  * P20: インスタンスファミリー
    * 汎用、コンピューティング最適化、ストレージ最適化、メモリ最適化、高速コンピューティング(GPU, FPGA)
  * P27: ベアメタル
  * P28: バースト可能パフォーマンスインスタンス
    * ベースライン CPU 性能を上回るときにクレジットを消費
    * unlimited にすると最大 24 時間分のクレジットを前借りする。前借りクレジットも消費し尽くした場合にバーストする場合は加算請求される。
* 機能とオプション
  * P37:
    * Private IP は停止しても不変(ENI に保存されている)。
    * Public IP は停止すると変わる。
  * P39:
    * 拡張ネットワーキング
* ストレージ
  * P44:
    * インスタンスストア: 停止するとクリアされる。
    * EBS
  * P45: EBS 最適化オプション。EBS 用に帯域が確保される。
* AMI
  * P48: 仮想化方式は HVM(完全仮想化)、PV(準仮想化)がある。PV は古い形式のため推奨されない。
  * P50: Nitro Hypervisor の注意事項
    * ENA, NVMe ドライバが有効である必要がある。acpid が有効でないと停止が正しく行えない。
* その他のオプション
  * P52: 専有オプションは２つ(ハードウェア専有インスタンス、Dedicated Host)
  * P54: プレイスメントグループ
  * P56: ハイパースレッディングの有効化/無効化
  * P57: アクセラレータオプション(Eastic Graphics, Elastic Inference)
* 運用管理
  * P60: ハイバネーション機能。メモリ状態をディスクに書き出しインスタンスを停止できる。
  * P61: CloudWatch, CloudWatch Logs
  * P62: スケジュールイベント(リタイア)
  * P63: Auto Recovery
  * P65: User Data
  * P66: Launch Template
  * P67: インスタンスメタデータ
  * P68: EC2 Fleet(1 回のリクエストで大量のオンデマンド、スポットインスタンスを起動可能)
  * P69: クォータと制限緩和
* 費用
  * P72: EC2 の購入オプション
    * オンデマンドインスタンス
    * リザーブドインスタンス
    * スポットインスタンス
    * 専用ホスト
    * ハードウェア専有インスタンス
  * P74: AMI
    * 無料 OS AMI
    * 商用 OS AMI(ソフトウェア費用もかかる)
    * AWS Marketplace(ソフトウェア費用もかかる)
  * P76: 課金管理の方法
    * 請求書、Detailed Billing Report、Cost Explorer、Trusted Advisor、概算見積もりツール



# 参考

* Document
  * [EC2 ドキュメント - Linux インスタンス用ユーザーガイド](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/concepts.html)
* サービス紹介ページ
  * [Amazon Elastic Block Store](https://aws.amazon.com/jp/ec2/)
  * [よくある質問](https://aws.amazon.com/jp/ec2/faqs/)
* Black Belt
  * [20190305 AWS Black Belt Online Seminar Amazon EC2](https://www.slideshare.net/AmazonWebServicesJapan/20190305-aws-black-belt-online-seminar-amazon-ec2)
  * [20200707 AWS Black Belt Online Seminar Amazon EC2 Deep Dive: AWS Graviton2 Arm CPU搭載インスタンス](https://www.slideshare.net/AmazonWebServicesJapan/20200707-aws-black-belt-online-seminar-amazon-ec2-deep-dive-aws-graviton2-arm-cpu)

