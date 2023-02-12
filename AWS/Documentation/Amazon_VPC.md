
# VPC

[VPC とサブネット](https://docs.aws.amazon.com/ja_jp/vpc/latest/userguide/VPC_Subnets.html)

* パブリックサブネット、プライベートサブネット
  * インターネットゲートウェイへのルートがあるサブネットは、パブリックサブネット。
  * インターネットゲートウェイへのルートがないサブネットは、プライベートサブネット。
* VPC に CIDR ブロックを設定する。
* VPC 内に複数のサブネットを作成可能。
* サブネット
  * 複数の AZ にまたがることはできない。
  * ネットワーク ACL への関連付けが必要


[IP アドレス指定](https://docs.aws.amazon.com/ja_jp/vpc/latest/userguide/vpc-ip-addressing.html)

* プライベート IPv4 アドレス
  * 各インスタンスには、プライベートの DNS ホスト名が割り当てられる。
  * プライマリプライベート IP アドレスを指定しない場合、サブネットの範囲内で使用可能な IP アドレスが選択される。
  * セカンダリ IP アドレスの割り当てが可能
* パブリック IPv4 アドレス
  * サブネット設定でパブリック IP アドレスを受信する設定になっている場合、インスタンス起動時にパブリック IP アドレスがプライマリネットワークインタフェースに割り当てられる。
  * 固定アドレスを割り当てたい場合は Elastic IP アドレスを割り当てる。



## セキュリティ

[セキュリティグループ](https://docs.aws.amazon.com/ja_jp/vpc/latest/userguide/VPC_SecurityGroups.html)

* 許可のみ設定可能。拒否は設定できない。
* インバウンド、アウトバウンド、それぞれについて設定可能。
  * デフォルトではアウトバウンドはすべて許可するルールになっている。
* ステートフル。往路で許可された場合、復路は自動的に許可される。
* セキュリティグループはネットワークインタフェースを対象に関連付ける。


[ネットワーク ACL](https://docs.aws.amazon.com/ja_jp/vpc/latest/userguide/vpc-network-acls.html)

* ネットワーク ACL はサブネットを対象に関連付ける。
* ルールの番号の低い順に評価する。
* ステートレス。インバウンドトラフィックの応答のパケットが許可されるかはアウトバウンドのルールに従う



## VPC のコンポーネント

[VPC のネットワーキングコンポーネント](https://docs.aws.amazon.com/ja_jp/vpc/latest/userguide/VPC_Networking.html)

[Elastic Network Interface](https://docs.aws.amazon.com/ja_jp/vpc/latest/userguide/VPC_ElasticNetworkInterfaces.html)

* プライマリネットワークインタフェースはデタッチ不可。

[ルートテーブル](https://docs.aws.amazon.com/ja_jp/vpc/latest/userguide/VPC_Route_Tables.html)

* サブネットとの関連付け
  * 明示的に関連付けていない場合、暗黙的にメインルートテーブルに関連付けられる。
  * 1 つのサブネットに 1 つのルートテーブルの関係。

[プレフィックスリスト](https://docs.aws.amazon.com/ja_jp/vpc/latest/userguide/managed-prefix-lists.html)

* プレフィックスリストは、1 つ以上の CIDR ブロックのセット

[インターネットゲートウェイ](https://docs.aws.amazon.com/ja_jp/vpc/latest/userguide/VPC_Internet_Gateway.html)

* 冗長性と高い可用性を備えており、水平スケーリングが可能。
* インターネットへルーティング可能。
* パブリック IPv4 アドレスが割り当てられているインスタンスに対する NAT を行う役割もある。
* 使用するには VPC にアタッチし、ルートテーブルに igw へのルートを設定する必要あり。

[NAT Gateway](https://docs.aws.amazon.com/ja_jp/vpc/latest/userguide/vpc-nat-gateway.html)

* NAT Gateway はパブリックサブネットに配置する必要がある。
* セキュリティグループを関連付けることはできない。

[DHCP オプションセット](https://docs.aws.amazon.com/ja_jp/vpc/latest/userguide/VPC_DHCP_Options.html)

以下の設定項目がある。

* domain-name-servers: 最大 4 つまでのドメインネームサーバーまたは AmazonProvidedDNS の IP アドレス。デフォルトは AmazonProvidedDNS。
* domain-name
* ntp-servers: 最大 4 つまでの Network Time Protocol (NTP) サーバーの IP アドレス。Amazon Time Sync Service は、169.254.169.123 で使用可能。
* netbios-name-servers
* netbios-node-type

[DNS](https://docs.aws.amazon.com/ja_jp/vpc/latest/userguide/vpc-dns.html)

* enableDnsHostnames: パブリック IP アドレスを持つインスタンスが、対応するパブリック DNS ホスト名を取得するかどうかの設定。
* enableDnsSupport: DNS 解決がサポートされているかどうかの設定。
  * true の場合、Amazon が提供する DNS サーバー (IP アドレス 169.254.169.253) へのクエリ、またはリザーブド IP アドレス (VPC IPv4 ネットワークの範囲に 2 をプラスしたアドレス) へのクエリが成功。
* 所定の書式でホストの FQDN を割り当てる。
  * パブリックホスト名
    * us-east-1 リージョン: ec2-public-ipv4-address.compute-1.amazonaws.com
    * その他のリージョン: ec2-public-ipv4-address.region.compute.amazonaws.com
  * プライベートホスト名
    * us-east-1 リージョン: ip-private-ipv4-address.ec2.internal
    * その他のリージョン: ip-private-ipv4-address.region.compute.internal

[VPC ピアリング接続](https://docs.aws.amazon.com/ja_jp/vpc/latest/userguide/vpc-peering.html)

2 つの VPC 間でプライベートなトラフィックのルーティングを可能にするネットワーキング接続。

* リージョンをまたいで構築可能
* 異なる AWS アカウント間でも可能
* 直接 Peering している VPC とのみ通信可能。2 HOP は不可。

[Elastic IP アドレス](https://docs.aws.amazon.com/ja_jp/vpc/latest/userguide/vpc-eips.html)



## VPC エンドポイント

[VPC エンドポイント](https://docs.aws.amazon.com/ja_jp/vpc/latest/userguide/vpc-endpoints.html)

VPC 内のリソースから各サービス、エンドポイントに対してプライベート接続を可能とする技術。

ゲートウェイエンドポイントは次のものが対応。エンドポイントポリシーを設定可能。
* Amazon S3
* DynamoDB

その他のサービスはインターフェイスエンドポイント。インターフェイスエンドポイントを作成することで、サブネット内に ENI が作成される。セキュリティグループでフィルタリング可能。



# BlackBelt

[20201021 AWS Black Belt Online Seminar Amazon VPC](https://www.slideshare.net/AmazonWebServicesJapan/20201021-aws-black-belt-online-seminar-amazon-vpc)

* P29: VPC に設定するアドレスレンジは作成後の変更不可。2 個目以降は追加、削除できる。
* P36: サブネットで利用できないアドレス。
  * .0: ネットワークアドレス
  * .1: VPC ルータ
  * .2: Amazon Provided DNS
  * .3: AWS で予約
  * .255: ブロードキャストアドレス(ブロードキャストはサポートされていない)
* P42: VPC 内からはプライベートアドレスが解決される。IGW からアクセスが有る場合はパブリックアドレスが使用される。
* P44: セキュリティグループはステートフル。インスタンス(ENI)単位で設定。すべてのルールを評価。Allow のみ設定可能。
* P45: NACL はステートレス。サブネット単位で設定。ルールの番号順に評価。Allow/Deny を設定可能。
* P49: カスタマーマネージドプレフィックスリスト: プレフィックスリストとはアドレスブロックをまとめたもの。
* P51: Ingress Routeing: IGW/VGW に対するトラフィックを特定の EC2 インスタンスに向けることが可能。
* P55: Amazon Provided DNS
  * CIDR 範囲のアドレスに + した値、もしくは 169.254.169.253 を使用可能。
* P56: Route 53 Resolver for Hybrid Clouds
* P57: 
  * Enable DNS resolution: 基本は enable とする。enable だと VPC の DNS 機能を使用可能。
  * Enable DNS hostname: true にしないと有効にならない。true だと DNS 名が割り当てられるようになる。
* P58: Amazon Time Sync Service: 169.254.169.123 で使用可能。



# 参考

* Document
  * [Amazon VPC とは?](https://docs.aws.amazon.com/ja_jp/vpc/latest/userguide/what-is-amazon-vpc.html)
  * [Amazon VPC actions](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/OperationList-query-vpc.html)
* サービス紹介ページ
  * [Amazon VPC](https://aws.amazon.com/jp/vpc/)
  * [よくある質問](https://aws.amazon.com/jp/vpc/faqs/)
* [ナレッジセンター](https://aws.amazon.com/jp/premiumsupport/knowledge-center/#Amazon_Virtual_Private_Cloud_.28Amazon_VPC.29)
* Black Belt
  * [20201021 AWS Black Belt Online Seminar Amazon VPC](https://www.slideshare.net/AmazonWebServicesJapan/20201021-aws-black-belt-online-seminar-amazon-vpc)
  * [20190417 AWS Black Belt Online Seminar Amazon VPC Advanced](https://www.slideshare.net/AmazonWebServicesJapan/20190417-aws-black-belt-online-seminar-amazon-vpc-advanced?ref=https://aws.amazon.com/)

