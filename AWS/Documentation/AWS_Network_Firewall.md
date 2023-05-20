# BlackBelt

[AWS Network Firewall 入門](https://pages.awscloud.com/rs/112-TZM-766/images/BlackBelt202106_AWS_Network_Firewall_Basic.pdf)

* サブネットに配置するマネージドファイアウォールサービス
* 利用までの流れ
  * Firewall 用のサブネットを作成
  * Firewall の作成
  * Firewall のルール作成
  * ルートテーブルで Firewall を通るように設定
* 機能
  * Firewall
    * ステートフルパケットフィルタ
    * ステートレスパケットフィルタ
    * ドメインリストフィルタ
    * Suricata 互換 IPS
  * 管理機能
    * CloudWatch Metrics
    * flow log, イベント、ルール別のログ
    * S3, CloudWatch Logs などへのログ格納
    * AWS Firewall Manager による一元管理
* 経路
  * NAT Gateway → Firewall endpoint → Internet Gateway



# 参考

* Document
  * [What is AWS Network Firewall?](https://docs.aws.amazon.com/ja_jp/network-firewall/latest/developerguide/what-is-aws-network-firewall.html)
* Black Belt
  * [AWS Network Firewall 入門](https://pages.awscloud.com/rs/112-TZM-766/images/BlackBelt202106_AWS_Network_Firewall_Basic.pdf)
  * [AWS Network Firewall 応⽤編1 新機能︓MSRの活⽤/トラフィックを集約して監査する](https://pages.awscloud.com/rs/112-TZM-766/images/202110_AWS_Black_Belt_Network_Firewall_advanced01.pdf)

