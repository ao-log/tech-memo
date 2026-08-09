
[AWS Certified Solutions Architect - Professional](https://aws.amazon.com/jp/certification/certified-solutions-architect-professional/)

* 試験時間 3h
* 75 問
* 300 USD の受験料
* 100 〜 1000 点のスケールスコアで 750 点以上で合格

[試験ガイド](https://d1.awsstatic.com/ja_JP/training-and-certification/docs-sa-pro/AWS-Certified-Solutions-Architect-Professional_Exam-Guide.pdf)

[サンプル問題 - 10 問](https://d1.awsstatic.com/ja_JP/training-and-certification/docs-sa-pro/AWS-Certified-Solutions-Architect-Professional_Sample-Questions.pdf)

[AWS Certified Solutions Architect - Professional 公式練習問題集](https://explore.skillbuilder.aws/learn/course/external/view/elearning/13272/aws-certified-solutions-architect-professional-official-practice-question-set-sap-c02-japanese)

[試験準備: AWS 認定 ソリューションアーキテクト - プロフェッショナル](https://awscertificationpractice.benchprep.com/app/aws-certified-solutions-architect-professional-official-practice-question-set-sap-c02-v2#exams)


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
* [S3 Transfer Acceleration](https://docs.aws.amazon.com/ja_jp/AmazonS3/latest/userguide/transfer-acceleration.html)
  * Transfer Acceleration は、世界中からの S3 汎用バケットへの転送の速度を最適化するように設計されている 
* PrivateLink エンドポイント
https://docs.aws.amazon.com/ja_jp/whitepapers/latest/building-scalable-secure-multi-vpc-network-infrastructure/aws-privatelink.html
* Aurora Global Database
https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html
* Step Functions
  * Express ワークフローで複数実行できる
* IP アドレス範囲が重複している場合は VPC ピアリングを作成できない。仮に VPC に CIDR を追加しても同様
* 時系列データの分析は Timestream を使用すると高速
* Route 53 は S3 にルーティングできない。静的ホスティングでは可能だが、アップロードは不可。
* アクセス許可の境界は IAM グループには設定できない
* CloudWatch Synthetics による外形監視


## その他ポイント

* RPO(目標復旧時点)、RTO(目標復旧時間)
* Route 53
  * [ルーティングポリシー](https://docs.aws.amazon.com/ja_jp/Route53/latest/DeveloperGuide/routing-policy.html)
  * [ヘルスチェック](https://docs.aws.amazon.com/ja_jp/Route53/latest/DeveloperGuide/health-checks-types.html)
  * [フェイルオーバールーティング](https://docs.aws.amazon.com/ja_jp/Route53/latest/DeveloperGuide/routing-policy-failover.html)
* API Gateway
  * [エンドポイントタイプ](https://docs.aws.amazon.com/ja_jp/apigateway/latest/developerguide/http-api-vs-rest.html#http-api-vs-rest.differences.endpoint-type)。その他機能一覧もおさえておくべき
  * [HTTP API](https://docs.aws.amazon.com/ja_jp/apigateway/latest/developerguide/http-api.html)
  * [REST API](https://docs.aws.amazon.com/ja_jp/apigateway/latest/developerguide/apigateway-rest-api.html)
* CloudFront
  * [CloudFront Functions と Lambda@Edge の違い](https://docs.aws.amazon.com/ja_jp/AmazonCloudFront/latest/DeveloperGuide/edge-functions-choosing.html)
* S3
  * [マルチリージョンアクセスポイント](https://docs.aws.amazon.com/ja_jp/AmazonS3/latest/userguide/MultiRegionAccessPoints.html)
* SQS
  * [機能と特徴](https://docs.aws.amazon.com/ja_jp/AWSSimpleQueueService/latest/SQSDeveloperGuide/features-capabilities.html)。デッドレターキュー、可視性タイムアウト、遅延キューなど
* [AWS Transfer Family](https://docs.aws.amazon.com/ja_jp/transfer/latest/userguide/what-is-aws-transfer-family.html)
* [AWS DataSync](https://docs.aws.amazon.com/ja_jp/datasync/latest/userguide/what-is-datasync.html)
* [AWS Application Discovery Service](https://docs.aws.amazon.com/ja_jp/application-discovery/latest/userguide/what-is-appdiscovery.html)
  * エージェントレス検出、エージェントベースの検出

