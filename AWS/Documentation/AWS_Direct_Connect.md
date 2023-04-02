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
* Direct Connect Gateway では複数 VPC への接続を提供可能
  * Direct Connect Gateway あたりの VGW 数は 10 個まで
* VGW では単一の VPC への接続を提供可能
* パートナーの提供サービス
  * 占有型と共有型がある
* 予算面でバックアップの Direct Connect 回線を用意できない場合はインターネット VPN を利用
* BFD(Bidirectional Forwarding Detection) を利用し障害を検知
* 利用できるメトリクス
  * 接続(Connection のアップ/ダウン、データ転送量など)
  * 仮想インターフェイス(VIF のデータ転送量など)



# 参考

* Document
  * [AWS Direct Connect とは](https://docs.aws.amazon.com/ja_jp/directconnect/latest/UserGuide/Welcome.html)
* Black Belt
  * [AWS Direct Connect](https://pages.awscloud.com/rs/112-TZM-766/images/20210209-AWS-Blackbelt-DirectConnect.pdf)


