
# ストレージクラス

### 概要

https://cloud.google.com/storage/docs/storage-classes?hl=ja

|項目|説明|想定するデータ|
|---|---|---|
| Multi-Regional Storage | マルチリージョンロケーションに配置 | 頻繁にアクセスされるデータ |
| Regional Storage | 単一ロケーションに配置。複数の可用性ゾーンに対し冗長 ||
| Nearline Storage | 取得の課金。早期削除に最小保存期間での課金 | 月1回程度アクセスするデータ置き場を想定 |
| Coldline Storage | 取得の課金。早期削除に最小保存期間での課金 | 年1回程度アクセスするデータ置き場を想定 |

同一リージョン内に他サービスを配置すると、パフォーマンス面で有利で、ネットワークの転送コストも抑えられる。

### オブジェクトごとのストレージクラス

https://cloud.google.com/storage/docs/per-object-storage-class?hl=ja

* バケット単位ではなく、オブジェクト単位でもストレージクラスを設定可能。
* バケットのデフォルトのストレージクラスを変更しても、すでに存在するオブジェクトのストレージクラスは変わらない。
* リージョンのロケーションのオブジェクトを Multi-Reagional Storage クラスに変更することはできない。

# オブジェクトのバージョンニング

https://cloud.google.com/storage/docs/object-versioning?hl=ja

バージョニングを有効にすると、オブジェクトの世代管理ができるようになる。

# オブジェクトのライフサイクル管理

https://cloud.google.com/storage/docs/lifecycle?hl=ja

ある日数経つとストレージクラスを変更する、指定した世代数のみ保存などの処理が可能。

### 操作

* 削除
* ストレージクラスの変更

### 条件

| 条件 | 説明 |
|---|---|
| Age | オブジェクト作成後に経過した日数 |
| CreatedBefore | 指定した日付より前 |
| isLive | true: ライブオブジェクト、false: アーカイブオブジェクト |
| MatchesStoragelass | 指定したストレージクラス |
| NumberOsNewerObjects | 指定数より新しいバージョンの数が多い場合に条件合致。より新しいバージョン数は、ライブオブジェクトだと 0 個、一つ前の世代だと 1 個になる。 |

# Cloud Pub/Sub Notifications for Cloud Storage

https://cloud.google.com/storage/docs/pubsub-notifications?hl=ja

オブジェクトの作成などを検知して、Pub/Sub で通知を送ることができる。

# Cloud Functions との連携

https://cloud.google.com/functions/docs/tutorials/storage?hl=ja

# アクセス制御

### ACL

##### IAM or ACL ?

バケット全てのオブジェクトにアクセス権限を適用する場合。→ IAM が適している。
個々のオブジェクトごとにアクセス権限を設定したい場合。→ ACL が適している。

https://cloud.google.com/storage/docs/access-control/lists?hl=ja

##### 権限

バケット、オブジェクトに適用可能。  
READER、WRITER、OWNER、デフォルトの 4 種。

# gsutil

```shell
# バケットの一覧
$ gsutil ls

# ライフサイクルの確認
$ gsutil lifecycle get gs://BUCKET_NAME/

# バケットの詳細（-L: 詳細、-b: バケットの指定）
$ gsutil list -L -b gs://BUCKET_NAME/
```

# ベストプラクティス

https://cloud.google.com/storage/docs/best-practices?hl=ja

# 料金

https://cloud.google.com/storage/pricing?hl=ja
