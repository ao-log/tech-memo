
[AWS Certified Solutions Architect - Professional](https://aws.amazon.com/jp/certification/certified-solutions-architect-professional/)

* 試験時間 3h
* 75 問
* 300 USD の受験料
* 100 〜 1000 点のスケールスコアで 750 点以上で合格

[試験ガイド](https://d1.awsstatic.com/ja_JP/training-and-certification/docs-sa-pro/AWS-Certified-Solutions-Architect-Professional_Exam-Guide.pdf)

[サンプル問題 - 10 問](https://d1.awsstatic.com/ja_JP/training-and-certification/docs-sa-pro/AWS-Certified-Solutions-Architect-Professional_Sample-Questions.pdf)

[AWS Certified Solutions Architect - Professional 公式練習問題集](https://explore.skillbuilder.aws/learn/course/external/view/elearning/13272/aws-certified-solutions-architect-professional-official-practice-question-set-sap-c02-japanese)

[試験準備: AWS 認定 ソリューションアーキテクト - プロフェッショナル](https://explore.skillbuilder.aws/learn/course/external/view/elearning/14951/exam-prep-aws-certified-solutions-architect-professional-sap-c02)


## おさえておくと良いサービス

* AD
* SSO
* VPC, PrivateLink, Egress-Only Internet Gateway
* S3
* KMS
* SES
* Migration Hub
* Organization
* RAM(AWS Resource Access Manager)
* SQS
* Cognito
* VPN
* DX
* Transit Gateway
* Storage Gateway
* Kinesis
* TimeStream


## 公式模擬試験ポイントまとめ

* CORS
* S3 Transfer Acceleration
* PrivateLink エンドポイント
https://docs.aws.amazon.com/ja_jp/whitepapers/latest/building-scalable-secure-multi-vpc-network-infrastructure/aws-privatelink.html
* Aurora Global Database
https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html
* Step Functions
  * Express ワークフローで複数実行できる
* IP アドレス範囲が重複している場合は VPC ピアリングを作成できない。仮に VPC に CIDR を追加しても同様
* 時系列データの分析は Timestream を使用すると高速
* Route 53 は S3 にルーティングできない。静的ホスティングでは可能だが、アップロードは不可。
* CloudWatch Synthetics による外形監視

