# Document

## 使用方法

[Amazon DynamoDB とは](https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/Introduction.html)

* フルマネージドの NoSQL データベースサービス
* AWS リージョン内の複数のアベイラビリティーゾーン間で自動的にレプリケート


[Amazon DynamoDB のコアコンポーネント](https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/HowItWorks.CoreComponents.html)

* プライマリキーはテーブルの各項目を一意に識別。プライマリキーはパーティションキー、パーティションとソートキーの 2 種類
* セカンダリインデックスはグローバルセカンダリインデックス、ローカルセカンダリインデックスの 2 種類


[DynamoDB API](https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/HowItWorks.API.html)


[読み込み整合性](https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/HowItWorks.ReadConsistency.html)

* 結果整合性のある読み込み、強力な整合性のある読み込みをサポート


[読み取り/書き込みキャパシティモード](https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html)

* オンデマンド、プロビジョニング済みの 2 種類のモードがある
* オンデマンドモード: リクエストごとの従量課金
* プロビジョニングモード: 1 秒あたりの読み込み、書き込みの回数を指定


[パーティションとデータ分散](https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/HowItWorks.Partitions.html)

* データはパーティションに保存される
* パーティションは、ソリッドステートドライブ (SSD) によってバックアップされ、AWS リージョン内の複数のアベイラビリティーゾーン間で自動的にレプリケートされる、テーブル用のストレージの割り当てのことを指す
* ハッシュ関数からの出力値によって項目が保存されるパーティションが決まる


## DynamoDB の操作

[DynamoDB Auto Scaling によるスループットキャパシティの自動管理](https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/AutoScaling.html)

* Amazon DynamoDB Auto Scaling は AWS Application Auto Scaling サービスを使用し、実際のトラフィックパターンに応じてプロビジョンドスループット性能をユーザーに代わって動的に調節


[Amazon DynamoDB の変更データキャプチャ](https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/streamsmain.html)

[Kinesis Data Streams を使用して DynamoDB への変更をキャプチャする](https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/kds.html)

以下内容を含むデータ項目を送信する。
* 項目が最後に作成、更新、または削除された特定の時刻
* その項目のプライマリキー
* 変更前の項目のイメージ
* 変更後の項目のイメージ


[DynamoDB のオンデマンドバックアップおよび復元の使用](https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/BackupRestore.html)

以下二つの方法がある。
* AWS Backup サービス
* DynamoDB


[DynamoDB での AWS Backup の使用](https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/backuprestore_HowItWorksAWS.html)

機能の一部を抜粋すると以下の通り。
* 定期バックアップ
* クロスアカウント、クロスリージョンへのバックアップ
* ライフサイクルの設定。コールドストレージに移行するなど


[DynamoDB のポイントインタイムリカバリ](https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/PointInTimeRecovery.html)

* 過去 35 日の任意の時点にテーブルを復元できる


## DynamoDB Accelerator (DAX) とインメモリアクセラレーション

[DynamoDB Accelerator (DAX) とインメモリアクセラレーション](https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/DAX.html)

ユースケースはマイクロ秒単位での応答が求められるような場合


## ベストプラクティス

[DynamoDB を使用した設計とアーキテクチャの設計に関するベストプラクティス](https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/best-practices.html)



# BlackBelt

[Amazon DynamoDB](https://pages.awscloud.com/rs/112-TZM-766/images/20170809_AWS-BlackBelt-DynamoDB_rev.pdf)

* RDBMS のスケールの限界を越えるために開発された経緯がある
* SPOF が存在しない。データは 3 AZ に保存される。高信頼性
* ストレージの容量制限がない。また自動的にパーティショニングされる
* 整合性モデル
  * 少なくとも 2 AZ から書き込みが取れた時点で ACK
  * 結果整合性がある読み込みなので、書き込み直後は即座に反映されない可能性がある
  * 整合性がある読み込みオプションを指定できる
* テーブル操作の API。GetItem, PutItem, Query, Scan など
* プライマリキーの持ち方は、パーティションキー、パーティションキー & ソートキーの 2 種類
* テーブル単位でスループットを設定できる
* スループットはパーティションに均等に付与される
* データモデリング
  * 1:1 の場合はパーティションキーのみで
  * 1:N の場合はパーティションキー & ソートキー
* その他機能
  * TTL: TTL を過ぎた項目は自動削除
* DynamoDB Streams
  * 追加、項目、削除の変更履歴を取り出すことができる
  * ユースケースは非同期な集計や、更新時に通知するなど
  * Lambda 関数でストリームイベントに対する処理を行うことも可能 [チュートリアル #1: AWS CLI を使用した Amazon DynamoDB と AWS Lambda での、フィルターを使ったすべてのイベントの処理](https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/Streams.Lambda.Tutorial.html)
  * クロスリージョンレプリケーション可能



# 参考

* Document
  * [Amazon DynamoDB とは](https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/Introduction.html)
* Black Belt
  * [Amazon DynamoDB](https://pages.awscloud.com/rs/112-TZM-766/images/20170809_AWS-BlackBelt-DynamoDB_rev.pdf)
  * [Amazon DynamoDB Advanced Design Pattern](https://pages.awscloud.com/rs/112-TZM-766/images/20181225_AWS-BlackBelt_DynamoDB_rev.pdf)


