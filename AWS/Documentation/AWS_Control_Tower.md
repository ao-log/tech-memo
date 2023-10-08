# Document

[AWS Control Tower とは](https://docs.aws.amazon.com/ja_jp/controltower/latest/userguide/what-is-control-tower.html)

Control Tower から新規 AWS アカウントを作成することができる。
作成したアカウントには予防および発見的制御 (ガードレール) が適用される。

## 機能

* ランディングゾーン: セキュリティ、コンプライアンスのベストプラクティスに基づくマルチアカウント環境。
* ガードレール: 予防的と発見的の 2 種類がある。必須、強く推奨、選択的の 3 つのガイダンスカテゴリが適用される。
* Account Factory: 新規アカウントをプロビジョニングする。

## 仕組み

[AWS Control Tower の仕組み](https://docs.aws.amazon.com/ja_jp/controltower/latest/userguide/how-control-tower-works.html)

組織の構造。

* ルート
  * Security OU: ログアーカイブアカウント、監査アカウントを含む。共有アカウントとも呼ばれる。
  * Sandbox OU: メンバーアカウントが含まれる。

アカウントごとの役割

* ログアーカイブアカウント: ランディングゾーン内の全てのアカウントから API アクティビティとリソース設定のログが送信されるアカウント。
* 監査アカウント: 他のアカウントをレビューする用途。他のアカウントに直接アクセスできないが、Lambda 関数を通してアクセスすることができる。

ガードレール

* 必須ガードレールは無効化できない。
* 予防ガードレール: アクションの実行を拒否する。SCP により実装されている。
* 検出ガードレール: 定義した状態からずれている場合に検出する。AWS Config によって実装されている。


## AFT

[AWS Control Tower Account Factory for Terraform (AFT) によるアカウントのプロビジョニング](https://docs.aws.amazon.com/ja_jp/controltower/latest/userguide/taf-account-provisioning.html)

**外部資料**

[ついにControl Towerのアカウント発行からカスタマイズまでIaC対応！Account Factory for Terraform (AFT)が新登場 #reinvent](https://dev.classmethod.jp/articles/ct-account-factory-for-terraform/)

* AFT 管理専用の AWS アカウントが必要
* AFT 専用の OU 上に AFT 管理専用のアカウントを配置することを推奨
* 管理アカウント上で AFT モジュールを apply する。[AWS Control Tower Account Factory for Terraform のリソースに関する考慮事項](https://docs.aws.amazon.com/ja_jp/controltower/latest/userguide/aft-resources.html)に記載されているリソースが作成される
* 運用で使用する「新規作成する AWS アカウントの情報」「ベースラインの設定内容」は別リポジトリで管理する
* Service Catalog のアカウントファクトリー用のポートフォリオに AFT の IAM ロールを追加
* AFT リポジトリにコードをプッシュすることで CodePipeline のパイプラインが動作。AWS アカウント発行の処理がされる



# 参考

* Document
  * [AWS Control Tower とは](https://docs.aws.amazon.com/ja_jp/controltower/latest/userguide/what-is-control-tower.html)
* サービス紹介ページ
  * [AWS Control Tower](https://aws.amazon.com/jp/controltower/)
  * [よくある質問](https://aws.amazon.com/jp/controltower/faqs/)
* [ナレッジセンター](https://aws.amazon.com/jp/premiumsupport/knowledge-center/#AWS_Control_Tower)
