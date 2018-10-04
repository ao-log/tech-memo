
# セクション 1: クラウド ソリューション環境の設定

### 1.1 クラウド プロジェクトとアカウントを設定する。次のような作業があります。

###### プロジェクトの作成

* [プロジェクトの作成と管理](https://cloud.google.com/resource-manager/docs/creating-managing-projects?hl=ja&visit_id=636741699649337684-270988070&rd=1)

###### プロジェクト内に事前定義された IAM 役割へのユーザーの割り当て

* [プロジェクト メンバーに対するアクセス権の付与、変更、取り消し](https://cloud.google.com/iam/docs/granting-changing-revoking-access)
* [事前定義された役割](
https://cloud.google.com/iam/docs/understanding-roles#predefined_roles)

###### G Suite ID へのユーザーのリンク

###### プロジェクト内での API の有効化

* [API の有効化と無効化](https://cloud.google.com/apis/docs/enable-disable-apis)

###### 1 つ以上の Stackdriver アカウントのプロビジョニング

### 1.2 課金設定を管理する。次のような作業があります。

###### 1 つ以上の請求先アカウントの作成

* [
請求先アカウントの作成、変更、閉鎖](https://cloud.google.com/billing/docs/how-to/manage-billing-account?hl=ja&visit_id=636741688280889601-943112649&rd=1)

###### 請求先アカウントへのプロジェクトのリンク

* [リンクされているプロジェクトの確認](https://cloud.google.com/billing/docs/how-to/view-linked?hl=ja)

###### 課金の予算とアラートの設定

* [
予算とアラートの設定](https://cloud.google.com/billing/docs/how-to/budgets?hl=ja)

###### 日 / 月単位の料金見積もりを目的とする請求関連のエクスポートの設定

* [
課金データのファイルへのエクスポート](https://cloud.google.com/billing/docs/how-to/export-data-file?hl=JA)
* [課金データの BigQuery へのエクスポート](https://cloud.google.com/billing/docs/how-to/export-data-bigquery?hl=JA)


### 1.3 コマンドライン インターフェース（CLI）、具体的には Cloud SDK をインストールし構成する（デフォルト プロジェクトの設定など）。

* [Google Cloud SDK ドキュメント](https://cloud.google.com/sdk/docs/?hl=ja#Getting_Started)

* [Cloud SDK の初期化](https://cloud.google.com/sdk/docs/initializing?hl=ja)

* [SDK のプロパティの管理](https://cloud.google.com/sdk/docs/properties?hl=ja)

プロパティ

```
$ gcloud config list
[compute]
region = us-east1
zone = us-east1-d
[core]
account = user@google.com
disable_usage_reporting = False
project = example-project
[metrics]
command_name = gcloud.config.list
```

プロジェクトを設定

```
$ gcloud config set project [PROJECT]
```

ゾーンを設定

```
$ gcloud config set compute/zone us-east1-b
```

# セクション 2: クラウド ソリューションの計画と構成

### 2.1 料金計算ツールを使用して GCP プロダクトの使用量を計画し、見積もる。

* [料金計算ツール](https://cloud.google.com/products/calculator/?authuser=3&hl=ja)
* [料金](https://cloud.google.com/compute/pricing?hl=ja-)

### 2.2 コンピューティング リソースを計画し、構成する。次のような内容を考察します。

###### ワークロードに適したコンピューティング サービスの選択（Compute Engine、Kubernetes Engine、App Engine など）

###### 必要に応じたプリエンプティブ VM とカスタム マシンタイプの使用

* [プリエンプティブ VM インスタンス](https://cloud.google.com/compute/docs/instances/preemptible?hl=ja)
* [カスタムマシンタイプ](https://cloud.google.com/compute/docs/machine-types?hl=ja#custom_machine_types)

### 2.3 データ ストレージ オプションを計画し、構成する。次のような内容を考察します。

###### プロダクトの選択（Cloud SQL、BigQuery、Cloud Spanner、Cloud Bigtable など）

###### ストレージ オプションの選択（Regional、Multi-regional、Nearline、Coldline など）

* [ストレージクラス](https://cloud.google.com/storage/docs/storage-classes?hl=ja)

### 2.4 ネットワーク リソースを計画し、構成する。次のようなタスクがあります。

###### 負荷分散オプションの差別化

* [負荷分散](https://cloud.google.com/compute/docs/load-balancing/)

###### 可用性を考慮したネットワーク内のリソース ロケーションの特定

###### Cloud DNS の構成

* [Cloud DNS](https://cloud.google.com/dns/docs/?hl=ja)

# セクション 3: クラウド ソリューションのデプロイと実装

### 3.1 Compute Engine リソースをデプロイし、実装する。次のようなタスクがあります。

###### Cloud Console と Cloud SDK（gcloud）を使用したコンピューティング インスタンスの起動（ディスクの割り当て、可用性ポリシー、SSH 認証鍵など）

###### インスタンス テンプレートを使用した、自動スケーリングされるマネージド インスタンス グループの作成

###### インスタンス用のカスタム SSH 認証鍵の生成 / アップロード

###### Stackdriver Monitoring と Logging のための VM の構成

###### コンピューティングの割り当ての評価と増加のリクエスト

###### モニタリングとロギング用の Stackdriver Agent のインストール

### 3.2 Kubernetes Engine リソースをデプロイし、実装する。次のようなタスクがあります。

###### Kubernetes Engine クラスタのデプロイ

###### ポッドを使用した Kubernetes Engine へのコンテナ アプリケーションのデプロイ

###### Kubernetes Engine アプリケーションのモニタリングとロギングの構成

### 3.3 App Engine リソースと Cloud Functions リソースをデプロイし、実装する。次のようなタスクがあります。

###### App Engine へのアプリケーションのデプロイ（スケーリング構成、バージョン、トラフィック分割など）

###### Google Cloud イベント（Cloud Pub/Sub イベント、Cloud Storage オブジェクト変更通知イベントなど）を受信する Cloud Function のデプロイ

### 3.4 データ ソリューションをデプロイし、実装する。次のようなタスクがあります。

###### プロダクトによるデータシステムの初期化（Cloud SQL、Cloud Datastore、BigQuery、Cloud Spanner、Cloud Pub/Sub、Cloud Bigtable、Cloud Dataproc、Cloud Storage など）

###### データの読み込み（コマンドラインによるアップロード、API による転送、インポート / エクスポート、Cloud Storage からのデータの読み込み、Cloud Pub/Sub へのデータのストリーミングなど）

### 3.5 ネットワーキング リソースをデプロイし、実装する。次のようなタスクがあります。

###### サブネットを使用した VPC の作成（カスタムモード VPC、共有 VPC など）

###### カスタム ネットワーク構成を使用した Compute Engine インスタンスの起動（内部専用 IP アドレス、限定公開の Google アクセス、静的外部 IP アドレスとプライベート IP アドレス、ネットワーク タグなど）

###### VPC 用の上りおよび下りファイアウォール ルールの作成（IP サブネット、タグ、サービス アカウントなど）

###### Cloud VPN を使用した Google VPC と外部ネットワーク間の VPN の作成

###### アプリケーションへのアプリケーション ネットワーク トラフィックを分散するロードバランサの作成（グローバル HTTP(S) ロードバランサ、グローバル SSL プロキシ ロードバランサ、グローバル TCP プロキシ ロードバランサ、リージョン ネットワーク ロードバランサ、リージョン内部ロードバランサなど）

### 3.6 Cloud Launcher を使用してソリューションをデプロイする。次のようなタスクがあります。

###### Cloud Launcher カタログの閲覧とソリューションの詳細の表示

###### Cloud Launcher Marketplace ソリューションのデプロイ

### 3.7 Deployment Manager を使用してアプリケーションをデプロイする。次のようなタスクがあります。

###### アプリケーションのデプロイを自動化する Deployment Manager テンプレートの開発

###### Deployment Manager テンプレートの起動による自動的な GCP リソースのプロビジョニングとアプリケーションの構成

# セクション 4: クラウド ソリューションの正常なオペレーションの確保

### 4.1 Compute Engine リソースを管理する。次のようなタスクがあります。

###### 単一 VM インスタンスの管理（起動、停止、構成の編集、インスタンスの削除など）

###### インスタンスへの SSH / RDP

###### 新しいインスタンスへの GPU の接続と CUDA ライブラリのインストール

###### 現在実行されている VM のインベントリ（インスタンス ID、詳細）の表示

###### スナップショットの操作（VM からのスナップショットの作成、表示、削除など）

###### イメージの操作（VM またはスナップショットからのイメージの作成、表示、削除など）

###### インスタンス グループの操作（自動スケーリング パラメータの設定、インスタンス テンプレートの割り当てや作成、インスタンス グループの削除など）

###### 管理インターフェースの操作（Cloud Console、Cloud Shell、GCloud SDK など）

### 4.2 Kubernetes Engine リソースを管理する。次のようなタスクがあります。

###### 現在実行されているクラスタのインベントリ（ノード、ポッド、サービス）の表示

###### コンテナ イメージ リポジトリの閲覧とコンテナ イメージの詳細の表示

###### ノードの操作（ノードの追加、編集、削除など）

###### ポッドの操作（ポッドの追加、編集、削除など）

###### サービスの操作（サービスの追加、編集、削除など）

###### 管理インターフェースの操作（Cloud Console、Cloud Shell、Cloud SDK など）

### 4.3 App Engine リソースを管理する。次のようなタスクがあります。

###### アプリケーションのトラフィック分割パラメータの調整

###### 自動スケーリング インスタンスのスケーリング パラメータの設定

###### 管理インターフェースの操作（Cloud Console、Cloud Shell、Cloud SDK など）

### 4.4 データ ソリューションを管理する。次のようなタスクがあります。

###### データ インスタンスからデータを取得するクエリの実行（Cloud SQL、BigQuery、Cloud Spanner、Cloud Datastore、Cloud Bigtable、Cloud Dataproc など）

###### BigQuery クエリのコストの見積もり

###### データ インスタンスのバックアップと復元（Cloud SQL、Cloud Datastore、Cloud Dataproc など）

###### Cloud Dataproc または BigQuery 内のジョブ ステータスの確認

###### Cloud Storage バケット間でのオブジェクトの移動

###### ストレージ クラス間での Cloud Storage バケットの変換

###### Cloud Storage バケットのオブジェクト ライフサイクル管理ポリシーの設定

###### 管理インターフェースの操作（Cloud Console、Cloud Shell、Cloud SDK など）

### 4.5 ネットワーキング リソースを管理する。次のようなタスクがあります。

###### 既存の VPC へのサブネットの追加

###### CIDR ブロック サブネットの拡張による IP アドレスの追加

###### 静的外部または内部 IP アドレスの予約

###### 管理インターフェースの操作（Cloud Console、Cloud Shell、Cloud SDK など）

### 4.6 モニタリングとロギングを行う。次のようなタスクがあります。

###### リソース指標に基づく Stackdriver アラートの作成

###### Stackdriver カスタム指標の作成

###### ログを外部システムにエクスポートするためのログシンクの構成（オンブレミスまたは BigQuery など）

###### Stackdriver のログの表示とフィルタリング

###### Stackdriver の特定のログメッセージの詳細の表示

###### Cloud Diagnostics を使用したアプリケーションの問題の調査（Cloud Trace データの表示、Cloud Debug を使用したアプリケーションのポイントインタイムの表示など）

###### Google Cloud Platform のステータスの表示

###### 管理インターフェースの操作（Cloud Console、Cloud Shell、Cloud SDK など）

# セクション 5: アクセスとセキュリティの構成

### 5.1 Identity and Access Management（IAM）を管理する。次のようなタスクがあります。

###### アカウントの IAM 割り当ての表示

###### アカウントまたは Google グループへの IAM 役割の割り当て

###### カスタム IAM 役割の定義

### 5.2 サービス アカウントを管理する。次のようなタスクがあります。

###### スコープの制限によるサービス アカウントの管理

###### VM インスタンスへのサービス アカウントの割り当て

###### 別のプロジェクトのサービス アカウントへのアクセス権の付与

### 5.3 プロジェクトとマネージド サービスの監査ログを表示する。
