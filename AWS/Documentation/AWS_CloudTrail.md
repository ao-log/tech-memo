# Document

[CloudTrail Lake](https://docs.aws.amazon.com/ja_jp/awscloudtrail/latest/userguide/cloudtrail-lake.html)

従来 Athena で分析していたが、こちらを使用することで SQL を使用したクエリが可能となる



# BlackBelt

[【AWS Black Belt Online Seminar】AWS CloudTrail](https://pages.awscloud.com/rs/112-TZM-766/images/20210119_AWSBlackbelt_CloudTrail.pdf)

* API リクエスト以外にマネジメントコンソールへのログインなども記録される
* データイベントも記録可能。S3 への GET, PUT など
* 証跡ログは S3 バケットや CloudWatch Logs に保存可能
* Organizations との連携で、組織の証跡を作成することが可能
* Config Rules でルールに準拠しているかを確認するとよい
* 過去 90 日間分は管理イベントの記録を無料で参照可能
* 基本的には Athena でクエリして分析。マネジメントコンソールから Athena テーブルを作成できる



# 参考

* Document
  * [AWS CloudTrail とは?](https://docs.aws.amazon.com/ja_jp/awscloudtrail/latest/userguide/cloudtrail-user-guide.html)
* Black Belt
  * [【AWS Black Belt Online Seminar】AWS CloudTrail](https://pages.awscloud.com/rs/112-TZM-766/images/20210119_AWSBlackbelt_CloudTrail.pdf)


