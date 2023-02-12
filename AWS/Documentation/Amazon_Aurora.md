## 特徴

[Amazon Aurora とは](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)

* MySQL および PostgreSQL と互換性がある
* MySQL のスループットの 5 倍、PostgreSQL のスループットの 3 倍を実現
* ストレージは、必要に応じて最大サイズの 128 tebibytes (TiB) まで自動拡張。


[Amazon Aurora DB クラスター](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/AuroraUserGuide/Aurora.Overview.html)

* プライマリ DB インスタンス: 読み書き両方をサポート
* Aurora レプリカ: 読み取りのみサポート。プライマリが使用できない場合、Aurora レプリカにフェイルオーバーされる。


[Amazon Aurora 接続管理](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/AuroraUserGuide/Aurora.Overview.Endpoints.html)

エンドポイントのタイプ

* クラスターエンドポイント: プライマリ DB インスタンスに接続
* リーダーエンドポイント: Aurora レプリカにロードバランシング
* カスタムエンドポイント: 選択した DB インスタンスのセット
* インスタンスエンドポイント: 特定の DB インスタンスに接続


[Amazon Aurora ストレージと信頼性](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/AuroraUserGuide/Aurora.Overview.StorageReliability.html)

* データはクラスターボリュームに保存される。


[Amazon Aurora でのレプリケーション](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/AuroraUserGuide/Aurora.Replication.html)

* Aurora レプリカは最大 15 個まで作成可能。



## Aurora DB クラスターの設定

[Amazon Aurora Serverless を使用する](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.html)

* アプリケーションのニーズに応じてコンピューティング性能が自動でスケールする。


[Amazon Aurora グローバルデータベースの使用](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html)

* 1 つのプライマリ AWS リージョン（データを管理）と最大 5 つのセカンダリ AWS リージョン（読み取り専用）で構成



## Aurora DB クラスターの管理

[Aurora DB クラスターボリュームのクローン作成](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/AuroraUserGuide/Aurora.Managing.Clone.html)

* クローンではコピーオンライトプロトコルが使用される。



## Aurora MySQL の操作

[Aurora DB クラスターのバックトラック](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Managing.Backtrack.html)

* バックアップからデータを復元しないで、DB クラスターを特定の時刻までバックトラックできる。



# BlackBelt

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
