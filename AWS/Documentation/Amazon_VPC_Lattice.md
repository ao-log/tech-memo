
# Document

[Amazon VPC Lattice とは](https://docs.aws.amazon.com/ja_jp/vpc-lattice/latest/ug/what-is-vpc-lattice.html)

* VPC Lattice では、マイクロサービスをサービスと呼んでいる
* コンポーネント
  * サービス: ELB と同じような概念。DNS 名が割り当てられ、リスナー、ルールがある
  * ターゲットグループ: ELB のターゲットグループと同じようなもの
  * Listener: プロトコル、ポート番号で LISTEN
  * ルール: Listener のトラフィック転送ルール
  * サービスネットワーク: サービスの集合の単位
  * サービスディレクトリ: 全 VPC Lattice サービスの中央のレジストリ
  * 認証ポリシー: 


[VPC Lattice の仕組み](https://docs.aws.amazon.com/ja_jp/vpc-lattice/latest/ug/how-it-works.html)

* サービスをサービスネットワークに関連づけると、サービスネットワークに関連づいた VPC やサービスから検出できるようになる。受信用の設定。
* VPC をサービスネットワークに関連づけると、クライアントとして利用できるようになり、関連づけられたサービスにリクエストを送信できるようになる。送信用の設定はじめる


