
[AWS Certified DevOps Engineer - Professional](https://aws.amazon.com/jp/certification/certified-devops-engineer-professional/)

* 試験時間 3h
* 75 問
* 300 USD の受験料
* 100 〜 1000 点のスケールスコアで 750 点以上で合格

[試験ガイド](https://d1.awsstatic.com/ja_JP/training-and-certification/docs-devops-pro/AWS-Certified-DevOps-Engineer-Professional_Exam-Guide.pdf)

[サンプル問題 - 10 問](https://d1.awsstatic.com/ja_JP/training-and-certification/docs-devops-pro/AWS-Certified-DevOps-Engineer-Professional_Sample-Questions.pdf)

[AWS Certified DevOps Engineer - Professional 公式練習問題集](https://explore.skillbuilder.aws/learn/course/external/view/elearning/12514/aws-certified-devops-engineer-professional-practice-question-set-dop-c01-japanese?devops=sec&sec=prep)

[試験準備: AWS Certified DevOps Engineer - Professional](https://explore.skillbuilder.aws/learn/course/external/view/elearning/74/exam-readiness-aws-certified-devops-engineer-professional?devops=sec&sec=prep)


#### 1. SDLC の自動化

* CI/CD
  * 継続的インテグレーション: 継続的にコードのコミットをし、移動的にテスト、ビルドアーティファクトの作成、メインブランチへのマージのサイクルを回す
  * 継続的デリバリー: ビルド、テストが通ったらステージング環境、本番環境にデプロイ。手動での承認プロセスを経る
  * 継続的デプロイ: 明示的な承認無く本番環境にデプロイ
* Code シリーズ
* ビルドプロセスは Jenkins を使用することも可能
* クロスアカウントパイプライン
  * CodeDeploy のクロスアカウント用 IAM ロールなどが必要
  * S3 バケット保存用の KMS キーの設定が必要
* ユニットテストを都度実行する
* 耐障害性のテスト。インスタンスを終了させて回復するか確認するなど
* デプロイ
  * インプレースデプロイは全てのインスタンスを一気に対応する
  * ローリングアップデートは一部ずつデプロイする


#### 2. 設定管理と Infrastructure as Code

* CloudFormation
  * Template の書き方
  * スタック更新時の影響。中断が発生するか、置き換えが発生するか
  * UpdatePolicy: ASG の更新時の動作を設定
  * ネストされたスタック
  * DependsOn
  * CreationPolicy: リソース作成の完了を待機
  * AWS::CloudFormation::WaitConditionHandle
  * スタックポリシー: 特定のリソースの更新を拒否したりできる
  * カスタムリソース: Lambda を使用するなどできる
  * StackSets
  * ヘルパースクリプト:
    * cfn-init: CFn のメタデータを実行する
    * cfn-hup: CFn のメタデータ変更時にインスタンスに反映。デーモンとして稼働させておく必要がある
    * cfn-signal: 処理の終了シグナルを送信できる
* Elastic Beanstalk
  * 構成要素: 環境、アプリケーションバージョン、設定
  * 更新方法: インプレース更新など。CNAME スワップなどの方法もとれる
  * .ebextensions
  * ECS も使用可能
* OpsWorks
  * Chef, Puppet を使用
  * OpsWorks スタック内に複数レイヤーがある構造
* Containers
  * CodePipeline でビルド、プッシュ、ECS デプロイする流れ
* Serverless
  * Lambda
  * API Gateway


#### 3. モニタリングとロギング

* CloudWatch
  * Metrics: 15 か月分保持される。古いデータは解像度が荒い状態に集約される
    * ELB: SurgeQueueLength, SpilloverCount
  * Alarm
  * Logs
    * メトリクスフィルター
  * EventBridge
* ログ
  * VPC フローログ
    * トラフィックの許可、拒否を確認できる
  * ELB アクセスログ
  * Route 53 クエリログ
* EventBridge
  * ルールとターゲット
* CloudTrail
* X-Ray
* Kinesis
  * Data Streams: ストリームデータの処理、分析
  * Data Firehose: ストレージにデータを流し込む
* タグづけ
  * リソースの管理、検索、フィルターに便利


#### 4. ポリシーと標準の自動化

* IAM
  * Condition: IP アドレスや MFA を条件に
* CloudFront
  * SNI
* Networking
  * Security Groups
  * NACL
* データ暗号化
  * S3
  * Glacier: デフォルトで暗号化されている
  * EBS
  * EFS: KMS のみ
* GuardDuty
  * CloudTrail, VPC フローログ, DNS クエリログから異常なパターンを検知
* Inspector
* Systems Manager
  * Run コマンド: インスタンスやタグを対象に事前定義されたコマンドや用意したコマンドを実行できる
  * パッチマネージャ: パッチベースラインを作成し、メンテナンスウィンドウに適用する。適用状況のレポート
  * パラメータ
* Secrets Manager
  * KMS による暗号化
  * ローテーションをサポート
* AWS Config
  * ルールへの準拠状況を継続的に追跡
  * 修復アクションを使用することもできる
* コスト最適化
  * タグ付け
* Trusted Advisor
* Service Catalog


#### 5. インシデントとイベントレスポンス

* ロギング戦略の構築
* 自動復旧
* Auto Scaling
  * デタッチ
  * スタンバイ状態
  * Suspend
  * 終了ポリシー
  * ライフサイクルフック: スケールアウト、スケールイン時などのタイミングでフックの処理を実行できる。発火したイベントをもとに Lambda 関数などで処理する
* ロードバランシング
  * スロースタート
* CodeDeploy


#### 6. 高可用性、耐障害性、災害対策

* 高可用性
  * マルチリージョン、マルチ AZ
  * Auto Scaling
  * DNS フェイルオーバー
  * DynamoDB - グローバルテーブル。キャパシティユニットの自動調整
  * RDS - リードレプリカ。ElastiCache へのキャッシュ。自動レプリケーションは 1 日 1 回
  * クロスリージョンレプリケーション
  * 疎結合化
    * SQS, SNS。SNS から複数の SQS キューにメッセージを送信できる
  * CloudFront
* 災害対策
  * RPO(目標復旧時点)、RTO(目標復旧時間) を設定する
* 障害に備えた設計
  * 単一障害点の排除
  * 高可用性構成
  * モニタリング


## おさえておくと良いサービス

* Code シリーズ
* ELB
* Auto Scaling
* ECS
* CloudFormation
* ServiceCatalog
* Elastic Beanstalk
* OpsWorks
* Lambda
* API Gateway
* CloudWatch
* EventBridge
* Systems Manager
* Organiaztions
* AWS Config
* GuardDuty
* Inspector
* Macie
* CloudTrail
* DynamoDB
* Trusted Advisor


## ポイント

* 用語
  * RPO(目標復旧時点)、RTO(目標復旧時間)
  * 3 層アーキテクチャ(フロント、バックエンド、DB)
* 各サービスごとの連携方法
  * [CodeCommit のイベント](https://docs.aws.amazon.com/ja_jp/codecommit/latest/userguide/monitoring-events.html)
  * [AWS リソースが準拠していない場合に AWS Config を使用して通知を受けるにはどうすればよいですか?](https://aws.amazon.com/jp/premiumsupport/knowledge-center/config-resource-non-compliant/)
  * [CloudWatch Logs サブスクリプションフィルターの使用](https://docs.aws.amazon.com/ja_jp/AmazonCloudWatch/latest/logs/SubscriptionFilters.html)
* 各サービスの制約
  * Lambda 関数のタイムアウト時間は 15 分
  * ポイントインタイムリカバリは 35 日まで
* デプロイ関連の機能
  * API Gateway
    * シンプルディストリビューションはトラフィックの分割はできない。加重ディストリビューションだとステージごとにトラフィックを指定比率で分割できる
    * カナリアデプロイの機能がある。[API Gateway の Canary リリースデプロイの設定](https://docs.aws.amazon.com/ja_jp/apigateway/latest/developerguide/canary-release.html)
  * Elastic Beanstalk
    * .ebextensions はアプリケーションの展開前に実行される。[Elastic Beanstalk Linux プラットフォームの拡張](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/platforms-linux-extend.html)
* 運用
  * RDS の自動スナップショットは 1 日 1 回。バックアップウインドウ中に実施される。[バックアップの使用](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html)
  * DynamoDB はテーブルに TTL を設定できる
* その他ポイント
  * 機能面で問題のない選択肢でも、非機能要件を満たせない場合は不正解の場合がある。例えばリアルタイムではないなど。
  * コスト最適化の観点も重要。機能面で満たせても、高価な場合はより安価な選択肢を優先する


