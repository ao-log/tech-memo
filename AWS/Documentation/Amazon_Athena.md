# BlackBelt

[Amazon Athena](https://pages.awscloud.com/rs/112-TZM-766/images/20200617_BlackBelt_Amazon_Athena.pdf)

* S3 上のデータに対して標準 SQL を使用した分析ができる
* ユースケース
  * ALB ログ、CloudFront ログ、VPC フローログなどに対するクエリのような活用方法もできる
  * 継続的に S3 にアップロードされるデータに対して ETL
* 使用方法
  * テーブル定義が必要
    * デフォルトでは AWS Glue Data Catalog 上のテーブル定義を使用
    * Athena DDL は HiveQL 形式で記述
  * クエリ
    * 実行したクエリの結果とメタデータ情報は自動的に S3 バケットに保存される
    * テクニック
      * 列指向フォーマットの使用、圧縮などによりスキャン量を減らすことが効果的
      * 大量に小さいファイルがある場合は 128 MB 以上の塊にまとめる
* Federated Query
  * Lambda で動作するコネクタを利用して実行
  * DynamoDB, Redshift, HBase, MySQL など様々なデータソースを使用可能
* UDF: ユーザ独自の関数を定義し SQL クエリから呼び出すことが可能。関数は Lambda で実行
* Machine Learning with Amazon Athena: クエリ内で SageMaker ML モデルを呼び出し推論を実行可能



# 参考

* Document
  * [Amazon Athena とは](https://docs.aws.amazon.com/ja_jp/athena/latest/ug/what-is.html)
* Black Belt
  * [Amazon Athena](https://pages.awscloud.com/rs/112-TZM-766/images/20200617_BlackBelt_Amazon_Athena.pdf)


