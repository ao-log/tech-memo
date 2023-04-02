
[Transit Gateway の動作](https://docs.aws.amazon.com/ja_jp/vpc/latest/tgw/how-transit-gateways-work.html)

* VPC 間を接続するアーキテクチャとすることができる
* 各 VPC のルートテーブルでは、異なる VPC を宛先とするルートが必要。宛先は tgw とする
* Trangit Gateway もルートテーブルを持つ。各 VPC の CIDR に対し、VPC のアタッチメントをターゲットとする
* VPC サブネット内に ENI をデプロイする仕組み。VPN, Direct Connect Gateway などをアタッチメントすることも可能


[例](https://docs.aws.amazon.com/ja_jp/vpc/latest/tgw/TGW_Scenarios.html)

各ユースケースごとのルートテーブルの例などが載っている。



# 参考

* Document
  * [Transit Gateway とは](https://docs.aws.amazon.com/ja_jp/vpc/latest/tgw/what-is-transit-gateway.html)
* Black Belt
  * [AWS Trangit Gateway](https://pages.awscloud.com/rs/112-TZM-766/images/20191113_AWS-BlackBelt_Transit_Gateway.pdf)


