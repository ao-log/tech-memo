
# Amazon S3

[はじめに](https://docs.aws.amazon.com/ja_jp/AmazonS3/latest/dev/Introduction.html)

* バケット名は全 AWS アカウントを通して一意である必要がある。
* オブジェクトはバケット内にフラットに配置される。「/」でフォルダ構造を表現することが可能。
* バケット作成時にリージョンを選択する。

2020年12月に強力な整合性がサポートされた。

[Amazon S3 アップデート – 強力な書き込み後の読み取り整合性](https://aws.amazon.com/jp/blogs/news/amazon-s3-update-strong-read-after-write-consistency/)



## バケット

[Amazon S3 バケットの使用](https://docs.aws.amazon.com/ja_jp/AmazonS3/latest/dev/UsingBucket.html)

URL は２つの形式がある。パス形式は非推奨になる予定で仮想ホスティング形式が推奨される。
```
// 仮想ホスティング形式
https://bucket-name.s3.Region.amazonaws.com/key_name
// パス形式
https://s3.Region.amazonaws.com/bucket-name/key name
```

* バケットの所有者は AWS アカウント
* オブジェクトのオーナは基本的にはオブジェクトを作成した IAM ユーザが属する AWS アカウント



## アクセスポイント

[アクセスポイント](https://docs.aws.amazon.com/ja_jp/AmazonS3/latest/dev/access-points.html)



## オブジェクト

[ストレージクラス](https://docs.aws.amazon.com/ja_jp/AmazonS3/latest/dev/storage-class-intro.html)

各オブジェクトには、ストレージクラスが関連付けられている。


[バージョニング](https://docs.aws.amazon.com/ja_jp/AmazonS3/latest/dev/ObjectVersioning.html)

一つのオブジェクトについて複数のバージョンを持つことができる。PUT 時は上書きではなく新バージョンの作成となる。DELETE 時は削除せず削除マーカーを挿入する。


[オブジェクトのタグ付け](https://docs.aws.amazon.com/ja_jp/AmazonS3/latest/dev/object-tagging.html)


[ライフサイクル管理](https://docs.aws.amazon.com/ja_jp/AmazonS3/latest/dev/object-lifecycle-mgmt.html)


[Cross-Origin Resource Sharing (CORS)](https://docs.aws.amazon.com/ja_jp/AmazonS3/latest/dev/cors.html)



## セキュリティ

[概要](https://docs.aws.amazon.com/ja_jp/AmazonS3/latest/dev/access-control-overview.html)

* バケット所有者には、他のアカウントが所有するオブジェクトに対するアクセス許可はない。
* 認証されていないリクエストは匿名ユーザーによって行われる
* 匿名ユーザの WRITE を強制的に防ぐにはパブリックアクセスのブロックを使用する


[Amazon S3 がオブジェクトオペレーションのリクエストを許可する方法](https://docs.aws.amazon.com/ja_jp/AmazonS3/latest/dev/access-control-auth-workflow-object-operation.html)

オブジェクトオペレーションに対しては次の順番で評価する。

* ユーザーコンテキスト(IAM ユーザー、ロール)
* バケットコンテキスト(バケットポリシー)
* オブジェクトコンテキスト(ACL)


[アクセスポリシーのオプションを使用するためのガイドライン](https://docs.aws.amazon.com/ja_jp/AmazonS3/latest/dev/access-policy-alternatives-guidelines.html)

* オブジェクト ACL は、バケット所有者が所有していないオブジェクトへのアクセスを管理する唯一の方法


[バケットポリシーの例](https://docs.aws.amazon.com/ja_jp/AmazonS3/latest/dev/example-bucket-policies.html)

バケットポリシー例が載っている。

* クロスアカウントの PUT の許可
* 匿名ユーザーへの GET の許可(ウェブサイトホスティングを行う場合)
* 特定の IP アドレス以外は拒否
* CloudFront からのみ許可
* bucket-owner-full-control が設定されている場合のみ、クロスアカウントの PUT を許可

など。

**注意点**

* バケットポリシーはオブジェクト所有者が所有するオブジェクトにのみ適用される。
* バケット所有者として所有していないオブジェクトが含まれる場合は、ACL による権限設定が必要。


[ユーザーポリシーの例](https://docs.aws.amazon.com/ja_jp/AmazonS3/latest/dev/example-policies-s3.html)

IAM ユーザーのポリシー例が載っている。


[ACL によるアクセス管理](https://docs.aws.amazon.com/ja_jp/AmazonS3/latest/dev/S3_ACLs_UsingACLs.html)

* バケット、オブジェクトそれぞれに ACL が存在。
* バケットまたはオブジェクトの作成時、S3 は所有者に対するフルアクセスコントロールの ACL を設定する動作となる。
* 付与できるアクセス許可は [ACL の概要](https://docs.aws.amazon.com/ja_jp/AmazonS3/latest/dev/acl-overview.html) を参照のこと。READ, WRITE, READ_ACP(ACL の読み込み許可), WRITE_ACP(ACL の書き込み許可), FULL_CONTROL がある。


[S3 のオブジェクトの所有権を使用したアップロードされたオブジェクトの所有権の管理](https://docs.aws.amazon.com/ja_jp/AmazonS3/latest/dev/about-object-ownership.html)

* オブジェクトの所有権は、デフォルトではアップロードアカウントによって所有されたままとなる。
* S3 のオブジェクトの所有権を使用すると、既定アクセスコントロールリスト (ACL) の bucket-owner-full-control を指定して他のアカウントによって書き込まれた新しいオブジェクトは、バケット所有者によって自動的に所有され、フルコントロールが与えられる。


[Amazon S3 のパブリックアクセスブロックの使用](https://docs.aws.amazon.com/ja_jp/AmazonS3/latest/dev/access-control-block-public-access.html)

* パブリックアクセスブロックでは、リソースがどのように作成されたかに関係なく強制的に適用することが可能。

どのような場合にパブリックとみなされるかの例もこのドキュメントに書かれている。



## 静的ウェブサイトのホスティング

[静的ウェブサイトを Amazon S3 上でホスティングする](https://docs.aws.amazon.com/ja_jp/AmazonS3/latest/dev/WebsiteHosting.html)

以下の設定が必要。

* 静的ウェブサイトホスティングの有効化
* アクセス許可の設定
  * パブリックアクセス(アカウントレベルとバケットレベルの両方でオフに設定)
  * バケットポリシーで匿名アクセスを許可
  * ACL(バケット所有者以外が所有するオブジェクトの場合は AllUsers への READ が必要)
* インデックスドキュメントの設定



## 通知

[Amazon S3 イベント通知の設定](https://docs.aws.amazon.com/ja_jp/AmazonS3/latest/dev/NotificationHowTo.html)

イベントのトリガーは以下のものがある。

* 新しいオブジェクトの作成
* オブジェクト削除
* オブジェクト復元(Glacier から)

など。

通知先は以下のものが設定できる。

* SNS トピック
* SQS キュー
* AWS Lambda



## レプリケーション

[レプリケーション](https://docs.aws.amazon.com/ja_jp/AmazonS3/latest/dev/replication.html)

レプリケーションの対象は [Amazon S3 がレプリケートするもの](https://docs.aws.amazon.com/ja_jp/AmazonS3/latest/dev/replication-what-is-isnot-replicated.html) を参照のこと。例えば、レプリケーション設定前に存在していたオブジェクトは対象外。



## リクエストルーティング

[リクエストルーティング](https://docs.aws.amazon.com/ja_jp/AmazonS3/latest/dev/UsingRouting.html)



## Amazon S3 のパフォーマンスの最適化

[設計パターンのベストプラクティス: Amazon S3 のパフォーマンスの最適化](https://docs.aws.amazon.com/ja_jp/AmazonS3/latest/dev/optimizing-performance.html)

* バケット内の**プレフィックスごとに** 1 秒あたり 3,500 回以上の PUT/COPY/POST/DELETE リクエストまたは 5,500 回以上の GET/HEAD リクエストを達成可能


## サーバーアクセスのログ記録

[Amazon S3 サーバーアクセスログ](https://docs.aws.amazon.com/ja_jp/AmazonS3/latest/dev/ServerLogs.html)



## BlackBelt

[20190220 AWS Black Belt Online Seminar Amazon S3 / Glacier](https://www.slideshare.net/AmazonWebServicesJapan/20190220-aws-black-belt-online-seminar-amazon-s3-glacier)

* 概要
  * P20: ストレージクラス
  * P21: 操作(GET/PUT/LIST/COPY/DELETE など)
* Amazon S3 へのアクセス
  * P24: アクセス管理(ユーザーポリシー = IAM、バケットポリシー、ACL)
  * P30: Block Public Access(バケット単位での保護)
  * P32: VPC Endpoint
  * P33: 署名付き URL
  * P35: Web サイトホスティング
* データ保護
  * P38: 暗号化(サーバサイド、クライアントサイド)
  * P39: バージョン管理
  * P40: オブジェクトロック機能
  * P43: クロスリージョンレプリケーション
* データ管理
  * P45: ストレージクラス
  * P46: ライフサイクル管理
  * P51: S3 Analytics(データアクセスパターンの可視化)
  * P54: S3 インベントリ(オブジェクトのリストの取得)
  * P56: S3 バッチオペレーション
  * P57: S3 イベント通知(SNS, SQS, Lambda などに対して通知可能)
  * P59: CLoudWatch Metrics によるモニタリング
  * P61: CloudTrail(データイベント)
  * P62: アクセスログ
* パフォーマンス最適化
  * P64: マルチパートアップロード
  * P66: S3 Transfer Accelration
  * P68: S3 Select
  * P70: リクエストレート
* 料金



# 参考

* [Amazon S3とは](https://docs.aws.amazon.com/ja_jp/AmazonS3/latest/dev/Welcome.html)
* [Amazon S3](https://aws.amazon.com/jp/s3/)
* [よくある質問](https://aws.amazon.com/jp/s3/faqs/)
* [ナレッジセンター](https://aws.amazon.com/jp/premiumsupport/knowledge-center/#Amazon_Simple_Storage_Service_.28Amazon_S3.29)
* Black Belt
  * [20190220 AWS Black Belt Online Seminar Amazon S3 / Glacier](https://www.slideshare.net/AmazonWebServicesJapan/20190220-aws-black-belt-online-seminar-amazon-s3-glacier)

