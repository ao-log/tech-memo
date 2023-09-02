
[AWS へのシステム移行の要点 (前編)](https://aws.amazon.com/jp/blogs/news/key-points-of-migrating-to-aws-part1/)

* [AWS クラウド導入フレームワーク (AWS CAF)](https://aws.amazon.com/jp/cloud-adoption-framework/) がある
* 評価フェーズにて移行にかかる TCO の算出、移行戦略、移行パスの評価などを行う
* 計画フェーズにて推進組織を立ち上げ、PoC(概念検証)や移行計画策定を実施
* 4 つのステップで進める
  * 1. 移行パターンを整理
    * 許容できる停止時間を見極める
    + 移行パターンを選定。一括移行、分割移行のどちらか
  * 2. データ転送経路を確保
    * 専用線接続、インターネット接続、物理デバイス
  * 3. 利用可能なAWS サービスを選択
    * アプリケーションサーバ
      * リホスト(リフト&シフト): 自動での移行には AWS Application Migration Service (AWS MGN) を利用
      * リロケート: vSphere ワークロードの場合 VMwareCloud on AWS を利用
      * リアーキテクチャ: Lambda, Fargate などにリアーキテクチャ
    * DB
      * 許容停止時間に余裕がある場合
        * 同一エンジンの場合はダンプツール
        * 異なるエンジンの場合は CSV アンロードとロード
      * 許容停止時間に余裕がない場合
        * 同一エンジンで純正のツールを使用できる場合はレプリケーション
        * AWS Database Migration Service (AWS DMS) を使用
    * ファイルサーバ移行
      * EC2: 従来の転送プロトコルを使用(多分 rsync など)
      * S3: AWS CLI, Transfer Family, AWS Snowball Edge または AWS Snowcone
      * ファイルストレージサービス: AWS DataSync, AWS Snowball Edge または AWS Snowcone
  * 4. 移行と付随する作業
    * アプリケーションテスト
    * システム移行時のサポート体制
    * 移行後の正常稼働の評価

