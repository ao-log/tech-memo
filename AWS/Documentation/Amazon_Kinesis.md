# BlackBelt

[Amazon Kinesis](https://pages.awscloud.com/rs/112-TZM-766/images/20180110_AWS-BlackBelt-Kinesis.pdf)

* Kinesis Streams
  * アーキテクチャ: 3 AZ の永続ストレージに強い整合性でデータを複製。順序付きイベントストリームとして複数のアプリケーションから同時アクセス可能
  * ストリームを作成し、ストリーム内にシャードがある構成
  * レコードの保持期間はデフォルトで 1 日。最長で 7 日
  * 1 データレコードの最大サイズは 1 MB
  * キャパシティ
    * データ送信側: 1 シャードあたり秒間 1 MB もしくは 1,000 PUT
    * データ受信側: 1 シャードあたり秒間 2 MB もしくは 5 回の読み取りトランザクション
    * シャード数の調整でスループットをコントロール
  * データ入力時に指定するパーティションにより、保存先のシャードが決定される
  * データレコードにはシーケンス番号が付与される
  * プロデューサー
    * Kinesis Agent: モニタリング対象のファイルや送信先のストリームを設定。バッファリングなどの機能がある
    * [Amazon Kinesis Producer Library (KPL)](https://docs.aws.amazon.com/ja_jp/streams/latest/dev/developing-producers-with-kpl.html): Aggregation(複数データを 1 レコードに集約)、Collection(複数レコードをバッファリングして送信) などを実装できる
  * コンシューマー
    * [Amazon Kinesis Client Library (KCL)](https://docs.aws.amazon.com/ja_jp/streams/latest/dev/shared-throughput-kcl-consumers.html)
* Kinesisu Firehose
  * ストリームデータを S3 などに配信
  * 配信ストリームを作成
  * シャードの設定は不要
  * 1 レコードの最大サイズは 1 MB
* Kinesis Analytics
  * アプリケーションを作成。ストリーミングのソースとデスティネーションを指定する
  * SQL クエリを実行できる
  * 前処理用の Lambda を設定できる



# 参考

* Document
  * [Amazon Kinesis のドキュメント](https://docs.aws.amazon.com/ja_jp/kinesis/index.html)
* Black Belt
  * [Amazon Kinesis](https://pages.awscloud.com/rs/112-TZM-766/images/20180110_AWS-BlackBelt-Kinesis.pdf)


