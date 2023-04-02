# BlackBelt

[AWS Site-to-Site VPN](https://pages.awscloud.com/rs/112-TZM-766/images/202110_AWS_Black_Belt_Site-to-Site_VPN.pdf)

* 構成
  * AWS Clinet VPN: VPN Client から Client VPN エンドポイントを通して VPC 内のサブネットにアクセスできる。VPN エンドポイントでは Directory Service でのユーザ認証や CloudWatch Logs へのログ送信、CloudTrail による API 記録が可能
  * Site-to-Site VPN を vgw 接続する構成: 特定の VPC に接続する場合に有用
  * Site-to-Site VPN を tgw 接続する構成: 様々な VPC に接続可能
* ユースケース
  * 拠点 - VPC 間を素早く接続したい
  * 価格重視
  * Direct Connect のバックアップ回線
* 設定
  * カスタマーゲートウェイ用の Config 設定。主要ベンダー用の Config はダウンロード可能
  * ルーティングは動的 or 静的で設定可能。動的(BGP) では片方の IPsec トンネルが使用できない場合にスムーズに切り替わる
  * 事前共有キー or プライベート証明書で認証
  * MTU を 1399 以下になるように設定することを推奨



# 参考

* Document
  * [AWS Site-to-Site VPN の概要](https://docs.aws.amazon.com/ja_jp/vpn/latest/s2svpn/VPC_VPN.html)
* Black Belt
  * [AWS Site-to-Site VPN](https://pages.awscloud.com/rs/112-TZM-766/images/202110_AWS_Black_Belt_Site-to-Site_VPN.pdf)


