
### VPC

* 全体のネットワーク空間を VPC として定義
* 最初に設定したアドレスレンジは変更不可
* 2 個目以降のアドレスレンジを追加、削除できる

分割例。IT オペレーションモデルに沿うように作成するのがよい。

* アプリケーションごと
* 本番、検証、開発
* 部署ごと、など

### サブネット

* 属する VPC、AZ、アドレス範囲、ネームタグを設定
* 5 つのアドレスは予約されており利用できない(.0: ネットワークアドレス、.1: VPCルータ、.2: DNS、.3: 予約されている、.255: ブロードキャストアドレス)
* パブリックサブネット、プライベートサブネット。パブリック IP アドレスが付与されている EC2 インスタンスでもインスタンス内からはパブリックアドレスは見えない。

### ルートテーブル

* local: 送信先が同一セグメントの場合、同一セグメント内に送る。VPC 作成時にデフォルトで作成される。
* igw: ルーティングテーブルに追加することで、インターネット上に出ていくことができる。インターネットゲートウェイの作成画面で VPC にアタッチしておく必要あり。

### セキュリティグループ

* インスタンス単位で設定
* ステートフル(行きが通れば、帰りも許可)
* Allow のみ設定。ホワイトリスト型
* デフォルトでは同一のセキュリティグループ内の通信のみ許可
* 全てのルールを適用

### Network ACLs

* サブネット単位で設定
* ステートレス(行きと帰り、それぞれ設定が必要)
* Allow, Deny を設定。ブラックリスト型
* デフォルトでは全て許可
* 順番通りにルールを適用

### サブネット内の DHCP

* ENI にアドレスを自動割り当てする

### DNS

次のいずれかのアドレスで DNS を利用可能。

* CIDR のネットワークアドレスに +2 した値
* 169.254.169.253

ホストに DNS 名を割り当てるには、以下の設定が必要。

* Enable DNS resolution
* Enable DNS hostname

### Route 53 Resolver for Hybrid Clouds

オンプレから Direct Connect/VPN 経由で VPC Provided DNS にアクセス可能な DNS エンドポイントを提供。DNS エンドポイントは ENI。
逆方向も可能。逆方向用の ENI を作成する。

### Amazon Time Sync Service

VPC 内の全ての EC2 インスタンスから NTP を利用可能。次のアドレスを設定すると良い。「169.254.169.123 」


### AWS との通信

* IGW 経由で通信する。
* S3, DynamoDB へは VPC エンドポイントを利用可能。

### VPC エンドポイント

**gateway型**

* ルートテーブルの宛先で pl-xxxx(AWSが管理するアドレス範囲)を設定する。
* エンドポイントポリシー(IAM policy と同じ構文)でアクセス制御。

**PrivateLink(interface型)**

* エンドポイント用の IP アドレスが作成される。作成されるアドレスは自動採番。
* DNS がこのアドレスに対して解決することで、ここを通して通信する。セキュリティグループで制御。

### NAT ゲートウェイ

* パブリックサブネットに配置する必要あり
* プライベートサブネット内では ルーティングテーブルに NAT ゲートウェイを追加する必要あり
* EIP を割り当て可能

### VPC Peering

* リージョンをまたいで構築可能
* 異なる AWS アカウント間でも可能
* 直接 Peering している VPC とのみ通信可能。2 HOP は不可。

### VPC Sharing

* 複数アカウントで VPC をシェア

### VPC Flow Logs

ネットワークトラフィックをキャプチャし、CloudWatch Logs に Publish する機能。ElasticSearch、Kibana と連携して可視化することもできる。


## 料金


# 参考

* [[AWS Black Belt Online Seminar] Amazon VPC Basic 資料及び QA 公開](https://aws.amazon.com/jp/blogs/news/webinar-bb-amazon-vpc-basic-2019/)
* [[AWS Black Belt Online Seminar] Amazon VPC Advanced 資料及び QA 公開](https://aws.amazon.com/jp/blogs/news/webinar-bb-amazon-vpc-advanced-2019/)
* [VPC ドキュメント](https://docs.aws.amazon.com/vpc/index.html)
