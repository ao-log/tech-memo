
[Helm Introduction Workshop](https://catalog.us-east-1.prod.workshops.aws/workshops/d75a936d-6a08-46b5-97c1-8666de887924/ja-JP)

* Helm チャートはマニフェストファイルをまとめたアーカイブファイル
* チャートリポジトリはチャートを公開する Web サーバ
* テンプレートエンジンとしての機能。変更可能なパラメータを用意しておき、デプロイ時にパラメータを変更できる
* Helm CLI
  * helm repo add でリポジトリを追加できる
  * helm show values <チャート> でチャートのパラメータのデフォルト値を確認できる
  * helm upgrade --install <チャート> でインストール
  * helm ls -A で全 namespace のチャート一覧を確認できる
* AWS Load Balancer Controller
  * 事前に IAM ポリシー、サービスアカウントの作成が必要
  * サービスアカウントには `metadata.annotations` に IAM ロールの ARN が設定されている。IRSA 用に必要な設定
  * helm から AWS Load Balancer Controller をインストールできる
  * helm history でリビジョンを確認できる
  * helm rollback でリビジョンを指定してロールバックできる
  * helm diff で新しいバージョンにアップグレードした場合の差分を確認できる


