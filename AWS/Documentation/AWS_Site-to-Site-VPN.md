# Document

[AWS Site-to-Site VPN の概要](https://docs.aws.amazon.com/ja_jp/vpn/latest/s2svpn/VPC_VPN.html)



## AWS Site-to-Site VPN の仕組み

[Site-to-Site VPN 接続のトンネルオプション](https://docs.aws.amazon.com/ja_jp/vpn/latest/s2svpn/VPNTunnels.html)

* デッドピア検出 (DPD)
* 事前共有キー (PSK)
など


[Site-to-Site VPN トンネル認証オプション](https://docs.aws.amazon.com/ja_jp/vpn/latest/s2svpn/vpn-tunnel-authentication-options.html)

事前共有キーとプライベート証明書のどちらかを使用して認証可能


[Site-to-Site VPN のルーティングオプション](https://docs.aws.amazon.com/ja_jp/vpn/latest/s2svpn/VPNRoutingTypes.html)

動的もしくは静的ルーティングを選択可能。使用可能な場合は BGP に対応したデバイスの使用をおススメ

プレフィックスが同じ場合は、以下の優先順位となる

* AWS Direct Connect 接続から BGP で伝播されたルート
* Site-to-Site VPN 接続用に手動で追加された静的ルート
* Site-to-Site VPN 接続から BGP で伝播されたルート
* 各 Site-to-Site VPN 接続が BGP を使用しているプレフィックスのマッチングでは、AS PATH が比較され、最短の AS PATH を持っているプレフィックス

その他優先度をコントロールできる BGP の設定項目

* MED: multi-exit discriminator。外部 AS のネイバールータに対して自身の AS 内への優先パスを示す。相手からの通信は MED 値が低いルートが優先される
* Local-Preference: 内部 AS 内のルータに対して、外部 AS への優先パスを示す。高い値を持つルートが優先される
* AS-path-prepending: 自身の AS に対して、外部のルータから受信する際にどの AS 経由とするかを指定



## アーキテクチャ

[アーキテクチャ](https://docs.aws.amazon.com/ja_jp/vpn/latest/s2svpn/site-site-architechtures.html)


[VPN CloudHub](https://docs.aws.amazon.com/ja_jp/vpn/latest/s2svpn/VPN_CloudHub.html)

vgw に複数の VPN を接続。オンプレミス - AWS 間だけでなく、オンプレミスサイト間での通信も可能



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
* CloudHub 構成
  * vgw, tgw に複数の VPN を接続する構成
  * 折り返し通信が必要な場合は ASN を別にする(同一 ASN からの経路情報をルートテーブルに反映しない BGP の特性)



# 参考

* Document
  * [AWS Site-to-Site VPN の概要](https://docs.aws.amazon.com/ja_jp/vpn/latest/s2svpn/VPC_VPN.html)
* Black Belt
  * [AWS Site-to-Site VPN](https://pages.awscloud.com/rs/112-TZM-766/images/202110_AWS_Black_Belt_Site-to-Site_VPN.pdf)


