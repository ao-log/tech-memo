# Document

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


[DynamoDB のオンデマンドバックアップおよび復元の使用](https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/BackupRestore.html)


[DynamoDB のポイントインタイムリカバリ](https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/PointInTimeRecovery.html)


[DynamoDB Accelerator (DAX) とインメモリアクセラレーション](https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/DAX.html)

ユースケースはマイクロ秒単位での応答が求められるような場合


[DynamoDB を使用した設計とアーキテクチャの設計に関するベストプラクティス](https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/best-practices.html)



# BlackBelt

[Amazon DynamoDB](https://pages.awscloud.com/rs/112-TZM-766/images/20170809_AWS-BlackBelt-DynamoDB_rev.pdf)

* SPOF が存在しない。データは 3 AZ に保存される
* ストレージの容量制限がない
* 整合性モデル
  * 少なくとも 2 AZ から書き込みが取れた時点で ACK
  * 結果整合性がある読み込みなので、書き込み直後は即座に反映されない可能性がある
  * 整合性がある読み込みオプションを指定できる
* テーブル操作の API。GetItem, PutItem, Query, Scan など
* プライマリキーの持ち方は、パーティションキー、パーティションキー & ソートキーの 2 種類


# 参考

* Document
  * [Amazon DynamoDB とは](https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/Introduction.html)
* Black Belt
  * [Amazon DynamoDB](https://pages.awscloud.com/rs/112-TZM-766/images/20170809_AWS-BlackBelt-DynamoDB_rev.pdf)
  * [Amazon DynamoDB Advanced Design Pattern](https://pages.awscloud.com/rs/112-TZM-766/images/20181225_AWS-BlackBelt_DynamoDB_rev.pdf)


