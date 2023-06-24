
[AWS Certified Advanced Networking - Specialty](https://aws.amazon.com/jp/certification/certified-advanced-networking-specialty/)

* 試験時間 2h 50m
* 65 問
* 300 USD の受験料
* 100 〜 1000 点のスケールスコアで 750 点以上で合格

[試験ガイド](https://d1.awsstatic.com/ja_JP/training-and-certification/docs-advnetworking-spec/AWS-Certified-Advanced-Networking-Specialty_Exam-Guide.pdf)

[サンプル問題 - 10 問](https://d1.awsstatic.com/ja_JP/training-and-certification/docs-advnetworking-spec/AWS-Certified-Advanced-Networking-Specialty_Sample-Questions.pdf)

[AWS Certified Solutions Architect - Professional 公式練習問題集](https://explore.skillbuilder.aws/learn/course/14837/play/69362/aws-certified-advanced-networking-specialty-official-practice-question-set-ans-c01-japanese)


## おさえておくと良いサービス

* Route 53
  * DNSSEC
  * Amazon Route 53 Resolver DNS Firewall
* VPC
  * Interface Endpoint, Gateway Endpoint
  * トラフィックのミラーリング
* Direct Connect
* Transit Gateway
* ELB
* CloudFront
* Global Acceleator
* WAF
* Network Firewall
* AWS Resource Access Manager

一般的な IT 知識

* BGP



## 公式模擬試験ポイントまとめ

* CloudFront は gRPC 未サポート
* Trangit Gateway のゲートウェイアプライアンスモード
* Transit Gateway への VPN の最大速度は 1.25 Gbps
* Route 53 でキー署名キーを作成する場合、カスタマーマネージドキーが必要。カスタマーマネージドキーは us-east-1 に存在する必要がある
* Route 53 の DNSSEC
* ローカルプリファレンスタグ 7224:7300。7224:9300 は顧客がアドバタイズするプレフィックスの伝播距離を制御するために使用
* AS_PATH は単一リージョンの複数の VIF には適しているが、マルチリージョンの複数の VIF には不適
* Amazon Route 53 Resolver DNS Firewall によるドメインのフィルタリング
* Site-to-Site VPN 接続は DNS ログ記録のオプションを提供しない
* Egress-only Gateway

各サービスの FAQ やナレッジセンター記事も読んでおく。
また、コンソール画面の VPC の画面の左メニューに多くのネットワーキング系サービスがあるので、触ることも大事。

