# Document

[Transit Gateway の動作](https://docs.aws.amazon.com/ja_jp/vpc/latest/tgw/how-transit-gateways-work.html)

* VPC 間を接続するアーキテクチャとすることができる
* 各 VPC のルートテーブルでは、異なる VPC を宛先とするルートが必要。宛先は tgw とする
* Trangit Gateway もルートテーブルを持つ。各 VPC の CIDR に対し、VPC のアタッチメントをターゲットとする
* VPC サブネット内に ENI をデプロイする仕組み。VPN, Direct Connect Gateway などをアタッチメントすることも可能


## 例

[例](https://docs.aws.amazon.com/ja_jp/vpc/latest/tgw/TGW_Scenarios.html)

各ユースケースごとのルートテーブルの例などが載っている。


[アプライアンス VPC](https://docs.aws.amazon.com/ja_jp/vpc/latest/tgw/transit-gateway-appliance-scenario.html)

* 共有サービス VPC でアプライアンスを設定できる
* アプライアンスモードが有効な場合、対称ルーティングとなる


## Transit Gateway の使用

[Direct Connect ゲートウェイへのトランジットゲートウェイアタッチメント](https://docs.aws.amazon.com/ja_jp/vpc/latest/tgw/tgw-dcg-attachments.html)

* Transit VIF を Direct Connect Gateway にアタッチする


[Transit Gateway VPN アタッチメント](https://docs.aws.amazon.com/ja_jp/vpc/latest/tgw/tgw-vpn-attachments.html)

* アタッチメントで VPN を選択し、カスタマーゲートウェイを指定


[Transit Gateway ピアリングアタッチメント](https://docs.aws.amazon.com/ja_jp/vpc/latest/tgw/tgw-peering.html)

* 各リージョンの Transit Gateway を接続する
* ピアリングアタッチメントを作成し、相手側の Transit Gateway で Accept する



# BlackBelt

[AWS Transit Gateway](https://pages.awscloud.com/rs/112-TZM-766/images/20191113_AWS-BlackBelt_Transit_Gateway.pdf)

* 注意点
  * Transit Gateway のアタッチメント ENI に専用のサブネットを用意することをオススメ
  * クロスアカウントや VPN の場合は Route 53 Resolver Endpoint が必要



# AWS Summit

[Transit Gateway Deep Dive アーキテクチャガイド](https://pages.awscloud.com/rs/112-TZM-766/images/B1-05.pdf)

* アタッチメント: VPC や VPN を Transit Gateway にくっつける。これだけだとまだ通信はできない
* ルートテーブル
  * アソシエーション: アタッチした VPC などをルートテーブルに結びつけること
  * プロバゲーション: アタッチした VPC から経路情報をルートテーブルに伝播する
  * プロパゲートすることで通信可能になる
* VPC 側のルートテーブルには経路情報が伝播しないので、別途書く必要がある
* ユースケース
  * インターネット接続用の VPC を設ける構成
    * デフォルトゲートウェイを Transit Gateway としておき、インターネット接続用の VPC 経由で通信。インターネット接続用 VPC のルートテーブルで戻りの経路の宛先を Transit Gateway に設定する必要がある
  * 共有リソースの VPC を用意し、VPC 間の接続を不可とする
    * Transit Gateway で二つのルートテーブルを作成。共有 VPC は各 VPC への経路があるルートテーブル。各 VPC は共有 VPC 宛のみのルートテーブルを使用
  * VPC 間の通信をインライン監査
    * 各 VPC にはミドルボックス向けのルートテーブル。ミドルボックスの VPC では各 VPC への経路があるルートテーブル



# 参考

* Document
  * [Transit Gateway とは](https://docs.aws.amazon.com/ja_jp/vpc/latest/tgw/what-is-transit-gateway.html)
* Black Belt
  * [AWS Trangit Gateway](https://pages.awscloud.com/rs/112-TZM-766/images/20191113_AWS-BlackBelt_Transit_Gateway.pdf)


