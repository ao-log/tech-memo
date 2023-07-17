
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

