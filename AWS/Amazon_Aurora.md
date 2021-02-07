## 特徴

[Amazon Aurora とは](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)

* MySQL のスループットの 5 倍、PostgreSQL のスループットの 3 倍を実現
* 基本ストレージは、必要に応じて最大 64 tebibytes (TiB) まで自動的に拡張。

## 構成

* ストレージとコンピューティングを分離したアーキテクチャ。
* プライマリ DB インスタンス(Writer/Reader)
* Aurora レプリカ(Reader)
* データのコピーを 3 つの AZ に保持。

## エンドポイント

* クラスターエンドポイント: プライマリ DB インスタンスに接続。
* 読み取りエンドポイント: Aurora レプリカに接続。
* カスタムエンドポイント: ユーザが作成できるもの。選択したインスタンスのセットからなるエンドポイントを作成できる。
* インスタンスエンドポイント: 特定のインスタンスをエンドポイントとできる。

## Aurora Serverless

* コンピーティングの部分を自動的に伸縮するサービス。可変、予測不能なワークロードに向く。

## バックアップ

* 自動バックアップ可能。
* バックアップ保存期間を超える場合はスナップショットで対応。
* バックアップはパフォーマンスへの影響なく取得可能。

## セキュリティ

* 暗号化の機能。KMS によってキーをローテーション可能。
* SSL/TLS 接続可能。現時点では CA-2019 が最新。接続時に使用する証明書は [SSL/TLS を使用した DB クラスターへの接続の暗号化](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL.html)からダウンロードできる。

## 機能

* バックトラック: 特定時点に巻き戻すことができる
* DB 認証: IAM DB 認証をすることも可能(AWS CLI でトークンを発行し、そのトークンで mysql コマンドなどで接続)
* 拡張モニタリング
* Performance Insight
* 暗号化




Amazon Aurora DB クラスターの管理
から




## BlackBelt

[20190424 AWS Black Belt Online Seminar Amazon Aurora MySQL](https://www.slideshare.net/AmazonWebServicesJapan/20190424-aws-black-belt-online-seminar-amazon-aurora-mysql)

* P11: Aurora のストレージ
  * 自動でスケールアップ
  * 3 つの AZ に 6 つのデータのコピーを作成
* P13: ストレージノードクラスタ
  * 各ログレコードは Log Sequence Number を持っている。不足、重複レコードを判別可能。ストレージノード間でゴシッププロトコルを使用し補完を行う。
* P14: 2 つのコピーに障害が発生しても読み書きに影響なし。3 つのコピーに障害が発生しても読み込み可能。
* P20: 6 つのコピーは全て同一ではない。3 つのフルセグメント。3 つのテールセグメント。



# 参考

* Document
  * [Amazon Aurora とは](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
* サービス紹介ページ
  * [Amazon Aurora](https://aws.amazon.com/jp/rds/aurora/)
  * [よくある質問](https://aws.amazon.com/jp/rds/aurora/faqs/)
* Black Belt
  * [20190424 AWS Black Belt Online Seminar Amazon Aurora MySQL](https://www.slideshare.net/AmazonWebServicesJapan/20190424-aws-black-belt-online-seminar-amazon-aurora-mysql)
  * [20190828 AWS Black Belt Online Seminar Amazon Aurora with PostgreSQL Compatibility](https://www.slideshare.net/AmazonWebServicesJapan/20190828-aws-black-belt-online-seminar-amazon-aurora-with-postgresql-compatibility-168930538)
  * [20200929 AWS Black Belt Online Seminar Amazon Aurora MySQL Compatible Edition ユースケース毎のスケーリング手法](https://www.slideshare.net/AmazonWebServicesJapan/20200929-aws-black-belt-online-seminar-amazon-aurora-mysql-compatible-edition)
