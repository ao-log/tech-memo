
[AWS におけるパブリック IPv4 アドレスの使用状況の特定と最適化](https://aws.amazon.com/jp/blogs/news/identify-and-optimize-public-ipv4-address-usage-on-aws/)

* パブリック IP アドレスの種類
  * Amazon EC2 パブリック IPv4 アドレス
  * Amazon が所有する Elastic IP アドレス
  * サービス管理のパブリック IPv4 アドレス
  * Bring Your Own IP (BYOIP) の 4 種類のパブリック IPv4 アドレス
* 確認方法
  * Cost and Usage Report (CUR)
    * 使用中のパブリック IPv4 アドレス、アイドル状態のパブリック IPv4 アドレスの使用状況データを確認可能
  * Amazon VPC IPAM Public IP Insights
    * 無料の機能。ダッシュボードを通してアカウント内でどのリソース及びサービスがパブリック IPv4 アドレスを使用しているか確認できる
* ベストプラクティス
  * サブネットでのパブリック IPv4 アドレスの自動割り当てを無効にする
  * インスタンスの起動時にパブリック IP アドレスを自動割り当てするよう変更することを検討
  * リモートアクセスでは Amazon EC2 Instance Connect (EIC) Endpoints を利用
  * インバウンドのインターネットトラフィックには、Elastic Load Balancers または AWS Global Accelerator の使用を検討


[『VPC 間通信ができる新サービス VPC Lattice』というタイトルで DevelopersIO 2023 札幌に登壇しました #devio2023](https://dev.classmethod.jp/articles/devio2023-sapporo-vpc-lattice/)

* VPC Lattice
  * 設定
    * ターゲットグループ
    * VPC Lattice Service: ドメイン名が割り当てられる。リスナー設定とターゲットグループの関連づけ
    * VPC Lattice Service Network: VPC との関連付け
* RAM で他アカウントと共有可能


[Amazon VPC のルーティング強化により、VPC 内のサブネット間のトラフィックが調査可能に](https://aws.amazon.com/jp/blogs/news/inspect-subnet-to-subnet-traffic-with-amazon-vpc-more-specific-routing/)

* 特定のアプライアンスを通過したいようなユースケースを実現できるようになった
* 以下のようにルートを追加しアプライアンスの ENI に向くようにする
```
aws ec2 create-route                                  \
     --region $REGION                                 \
     --route-table-id $BASTION_SUBNET_ROUTE_TABLE     \
     --destination-cidr-block 10.0.1.0/24             \
     --network-interface-id $APPLIANCE_ENI_ID
```
* つまり、従来 Local 宛の通信経路を変えることはできなかったが、特定 CIDR 宛の通信を特定 ENI に向けることができるようになった
* アプライアンス用のインスタンスでは `net.ipv4.ip_forward=1` でトラフィック転送できるようにし、Source/dest check を無効化する


[Amazon VPC Block Public Access による VPC セキュリティの強化](https://aws.amazon.com/jp/blogs/news/vpc-block-public-access/)

* VPC からインターネットに対するインバウンド、アウトバウンドトラフィックをブロックできる
* Internet Gateway があったとしても、この設定の方が優先される
* イングレス方向のみをブロックできる
* 除外対象とする VPC, サブネットを設定できる




