
# RDS

#### [マルチ AZ](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/Concepts.MultiAZ.html)

> フェイルオーバーが完了するまでにかかる時間は、 ...(略)... 通常 60 ～ 120 秒です。

> 同期データレプリケーションが発生するため、シングル AZ 配置より書き込みとコミットのレイテンシーが増加する可能性があります


#### [リードレプリカ](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/USER_ReadRepl.html)

> ソース DB インスタンスの DB スナップショットを取得し、レプリケーションを開始する

#### [DB スナップショットの作成](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/USER_CreateSnapshot.html)

> Single-AZ DB インスタンスでこの DB スナップショットを作成すると、I/O が短時間中断します。この時間は、DB インスタンスのサイズやクラスによって異なり、数秒から数分になります。マルチ AZ DB インスタンスは、スタンバイ側でバックアップが作成されるため、この I/O 中断の影響を受けません

#### [ポイントインタイムリカバリ](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/USER_PIT.html)

> RDS は、DB インスタンスのトランザクションログを 5 分ごとに Amazon S3 にアップロード


# 参考

* [[AWS Black Belt Online Seminar] Amazon RDS 資料及び QA 公開](https://aws.amazon.com/jp/blogs/news/webinar-bb-rds-2018/)
* [Amazon Relational Database Service ドキュメント](https://docs.aws.amazon.com/ja_jp/rds/?id=docs_gateway)
