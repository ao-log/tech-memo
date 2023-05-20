# Document

[Amazon S3 Glacier とは](https://docs.aws.amazon.com/ja_jp/amazonglacier/latest/dev/introduction.html)

* ストレージ階層は三つ
  * S3 Glacier Instant Retrieval
  * S3 Glacier Flexible Retrieval
  * S3 Glacier Deep Archive


[データモデル](https://docs.aws.amazon.com/ja_jp/amazonglacier/latest/dev/amazon-glacier-data-model.html)

* ボールト: バケットのような概念
* アーカイブ: オブジェクトのような概念
* ジョブ: ボールトのインベントリ獲得やアーカイブの取得を実施


[S3 Glacier ボールトロック](https://docs.aws.amazon.com/ja_jp/amazonglacier/latest/dev/vault-lock.html)

* ボールトロックポリシー: 1 年経過するまで削除を拒否するようなポリシーを作成できる
* ボールトアクセスポリシー: クロスアカウント権限の付与などの権限管理ができるリソースベースのポリシー




# 参考

* Document
  * [Amazon S3 Glacier とは](https://docs.aws.amazon.com/ja_jp/amazonglacier/latest/dev/amazon-glacier-data-model.html)
* Black Belt
  * [20190220 AWS Black Belt Online Seminar Amazon S3 / Glacier](https://www.slideshare.net/AmazonWebServicesJapan/20190220-aws-black-belt-online-seminar-amazon-s3-glacier)

