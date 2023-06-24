# Document

[AWS Direct Connect とは](https://docs.aws.amazon.com/ja_jp/directconnect/latest/UserGuide/Welcome.html)

[ルーティングポリシーと BGP コミュニティ](https://docs.aws.amazon.com/ja_jp/directconnect/latest/UserGuide/routing-and-bgp.html)

BGP コミュニティ。

パブリック VIF

* インバウンドルーティングポリシーの場合。AWS に広告するプレフィックスの伝搬を制御
  * 7224:9100 – ローカル AWS リージョン
  * 7224:9200 – 大陸内のすべての AWS リージョン
  * 7224:9300 – グローバル (すべてのパブリック AWS リージョン)
* アウトバウンドルーティングポリシーの場合。AWS が広告するプレフィックスに付与される
  * 7224:8100 — AWS のプレゼンスポイントが関連付けられている AWS Direct Connect リージョンと同じリージョンから送信されるルート
  * 7224:8200 — AWS Direct Connect のプレゼンスポイントが関連付けられている大陸と同じ大陸から送信されるルート
  * タグなし - グローバル (すべてのパブリック AWS リージョン)
* NO_EXPORT BGP コミュニティタグをサポート

プライベート VIF。

* AS_PATH プリペンド。Direct Connect 接続が VPC とは異なるリージョン にある場合は機能しない
* ローカル優先設定の BGP コミュニティタグ。AWS に広告するローカルプリファレンスを制御
  * 7224:7100 - 優先設定: 低
  * 7224:7200 - 優先設定: 中
  * 7224:7300 - 優先設定: 高


[プライベート仮想インターフェイスルーティングの例](https://docs.aws.amazon.com/ja_jp/directconnect/latest/UserGuide/private-transit-vif-example.html)

* AWS からの通信は以下の通り決まる
  * CIDR のプレフィックスの一致部分が一番長い経路を選択
  * AS_PATH が短い経路を選択


## 接続

[接続](https://docs.aws.amazon.com/ja_jp/directconnect/latest/UserGuide/WorkingWithConnections.html)

* 専用接続、ホスト接続の 2 種類がある
* 専用接続は 1, 10 , 100 Gbps から指定可能
* ホスト接続は 50 Mbps、100 Mbps、200 Mbps、300 Mbps、400 Mbps、500 Mbps、1 Gbps、2 Gbps、5 Gbps、10 Gbps から指定可能


## 仮想インタフェース

[AWS Direct Connect 仮想インターフェイス](https://docs.aws.amazon.com/ja_jp/directconnect/latest/UserGuide/WorkingWithVirtualInterfaces.html)

* プライベート VIF, パブリック VIF, トランジット VIF の 3 種類


## Direct Connect Gateway

[Direct Connect ゲートウェイの操作](https://docs.aws.amazon.com/ja_jp/directconnect/latest/UserGuide/direct-connect-gateways.html)

* VIF を直接 vgw にアタッチする構成ではなく、Direct Connect Gateway に接続する方式が主流
* ただしパブリック VIF の場合は Direct Connect Gateway は不要
* Direct Connect Gateway は vgw もしくは Transit Gateway に関連づける
* Direct Connect Gateway はグローバルリソース。複数リージョンの VPC を関連づけられる
* 最大 10 個の vgw を収容可能。この場合 Direct Connet Gateway を介した VPC 間の通信はできない

一方で Transig VIF の場合...

* VPC を 5,000 まで接続可能
* VPC 間も通信可能
* 1 Gbps/ホスト接続が必要
* オンプレからの接続は Private VIF と Direct Connect Gateway で行い、VPC 間の通信用途に連携が必要な VPC のみ Transit Gateway を使用するのも Tips の一つ



# BlackBelt

[AWS Direct Connect](https://pages.awscloud.com/rs/112-TZM-766/images/20210209-AWS-Blackbelt-DirectConnect.pdf)

* お客様の専用線の片端と AWS Cloud とを Direct Connect ロケーションで接続するサービス
* Link Aggregation Group を作成可能
* SLA 要件に応じて回復性の設定を行う。2 つのロケーションへの冗長化、ロケーション内での冗長化が可能
* 仮想インターフェイス(VIF)
  * お客様ルータとの間に BGP ピアを確立するためにも使用
  * VIF ごとに VLAN ID を持つ
  * VIF は 3 種類
    * プライベート VIF: VPC 内のプライベート IP アドレスへの接続を提供
    * パブリック VIF: AWS の全リージョンのパブリック IP を介した接続を提供
    * トランジット VIF: Transit Gateway 用の Direct Connect Gateway へ接続を提供
  * 他の AWS アカウントに対して VIF を提供可能
* Transit Gateway
  * Transit VIF - Direct Connect Gateway - Transit Gateway - VPC の経路
  * VPC 間の折り返しは可能
* Direct Connect Gateway では複数 VPC への接続を提供可能
  * Direct Connect Gateway あたりの VGW 数は 10 個まで
  * Transit Gateway をアタッチ可能
  * VPC 間の折り返しは不可
* VGW では単一の VPC への接続を提供可能
* パートナーの提供サービス
  * 占有型と共有型がある
* 高い回復性
  * 異なるロケーションへの分散が基本
* 経路制御
  * LP(Local Preference) で AWS への送信トラフィックの経路を制御。数字が大きい方が優先
  * AS-Path Prepend で AWS からの受信トラフィックの経路を制御
  * Active/Active の場合は LP, MED, AS-Path Prepend が同じになるようにする
  * 予算面でバックアップの Direct Connect 回線を用意できない場合はインターネット VPN を利用。仕様上 Direct Connect の経路が優先される
  * BFD(Bidirectional Forwarding Detection) を利用し障害を検知
* 利用できるメトリクス
  * 接続(Connection のアップ/ダウン、データ転送量など)
  * 仮想インターフェイス(VIF のデータ転送量など)


[オンプレミスとAWS間の冗長化接続](https://pages.awscloud.com/rs/112-TZM-766/images/20200219_BlackBelt_Onpremises_Redundancy.pdf)

* 複数の Direct Connect ロケーションによる冗長化
  * Active-Active, Active-Standby はお客様ルータにて制御
  * Active-Standby の場合、オンプレミス → AWS は Local Preference、AWS → オンプレミスは AS Path Prepend で設定
  * MED による優先制御もできるがサポート対象外
* VGW が持つルートテーブルは参照できないが、サポートに問い合わせると確認可能
* 障害時の経路切り替え時間の短縮には BFD を利用し高速に障害を検知する
* VPN 接続
  * AS パス属性などを使用した場合でも Direct Connect の経路が優先される
  * VPN を優先したい場合は経路を分割して広報する
* Direct Connect Gateway
  * 同時にアタッチ可能な VIF は 30 まで
* マルチリージョン冗長
  * 例えば東京、大阪リージョンの Direct Connect ロケーション
  * Direct Connect Gateway はグローバルリソースなので、二つの VIF を接続可能
* 既存 VPC から Transit Gateway への移行
  * Transit Gateway と VPC を接続しておく
  * Direct Connect Gateway と Transit VIF を接続。Direct Connect Gateway に Transit Gateway をアタッチ
  * Direct Connect ロケーションのルータにて VPC 宛の経路を Private VIF → Transit VIF に変更
  * VPC のルートテーブルにてオンプレ宛の経路を vgw から Transit Gateway に変更



# AWS Summit

[ネットワークデザインパターン Deep Dive](https://pages.awscloud.com/rs/112-TZM-766/images/B1-04.pdf)

* オンプレミスから VPC 内リソースの名前解決
  * Route 53 Inbound Endpoint を作成
  * オンプレ側の DNS リゾルバにて AWS ドメインは Inbound Endpoint に転送するように設定
* VPC 内からオンプレミスへの名前解決
  * Outbound Endpoint からオンプレミスの DNS リゾルバに転送
* オンプレミス、AWS 間の接続を高信頼に(上から順に優先度が高い)
  * デュアル Direct Connect ロケーション
  * 同一通信キャリアで異経路設定
  * 異なる通信キャリア使用
  * 大阪リージョン、マルチリージョン
  * BFD の利用
  * LAG の利用
  * デュアルシステム設計 (同一構成の VPC を 2 面構築)



# 参考

* Document
  * [AWS Direct Connect とは](https://docs.aws.amazon.com/ja_jp/directconnect/latest/UserGuide/Welcome.html)
* Black Belt
  * [AWS Direct Connect](https://pages.awscloud.com/rs/112-TZM-766/images/20210209-AWS-Blackbelt-DirectConnect.pdf)


