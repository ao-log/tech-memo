
[AWS Control Tower でコントロールを適用する際のベストプラクティス](https://aws.amazon.com/jp/blogs/news/best-practices-for-applying-controls-with-aws-control-tower/)

* 適切なコントロールを適用するために、ワークロードと OU を把握して評価する
  * ビジネス要件などと照らし合わせた上で、OU 構造を整理すること
* コンプライアンスフレームワークとの連携を検討する
  * NIST 800-53、CIS AWS Benchmark、PCI-DSS などのフレームワークがある
* AWS Control Tower のコントロールを有効にする前に、その動作とメカニズムを理解する
  * 予防コントロール
  * 検出コントロール
  * プロアクティブコントロール: プロビジョニングする前に準拠状況を確認できる
* 予防コントロールを検討する前に、検出コントロールを適用する
  * あるべき姿とのギャップを確認してから予防コントロールを適用する
* non-production OU でコントロールをテストする
* 有効化されているコントロールを継続的に監視・テストする積極的なアプローチを採用する
  * AWS Audit Manager を使用して自動評価を実施
  * AWS Identity and Access Management (IAM) Access Analyzer を使用してアクセスパターンを確認
* Policy as Code 戦略を採用し、組織全体でピアレビューを実施する
  * AWS CloudFormation Hook と AWS CloudFormation Guard ルールを組み合わせた AWS Control Tower のプロアクティブコントロールを使用
* 3 種類すべてのコントロールを組み合わせて有効化する多層防御アプローチをとる
* コンプライアンス違反リソースの検出と修復を自動化する
  * ルールへの非準拠時の自動修復アクションを設定する。AWS Control Tower の検出コントロールと AWS Systems Manager Automation を組み合わせて実現。場合によっては SSM ドキュメントの開発が必要
* 独自のコントロールを作成して機能を拡張する


