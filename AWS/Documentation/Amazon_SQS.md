# BlackBelt

[Amazon Simple Queue Service](https://pages.awscloud.com/rs/112-TZM-766/images/20190717_AWS-BlackBelt_AmazonSQS.pdf)

* 利用ケース
  * 一時的なリクエスト増への対応
  * プロデューサ、コンシューマ間の依存関係の低減
  * 重い処理を切り出す
  * SNS → SQS で複数キューにファンアウトすることで、並列化が可能
* キュー
  * FIFO キューだと 1 回のみの配信。スタンダードキューだと少なくとも 1 回の配信
  * 2 回以上の配信に備えて冪等性の確保が重要
  * メッセージ取得の際はロングポーリング or ショートポーリング。通常はロングポーリング
* 機能
  * 可視性タイムアウト: コンシューマが取得したメッセージに対して、一定期間(デフォルト 30 秒) 他のコンシューマからアクセスをブロックする機能
  * 遅延キュー: キューに送信してから一定時間経過後に利用可能になる機能。キューに対して適用
  * メッセージタイマ: 特定のメッセージに対して、一定時間経過後に利用可能になる機能
  * デッドレターキュー: 正常処理できないメッセージの分離先となるキュー
  * 暗号化: サーバサイド暗号化でメッセージの暗号化が可能
  * アクセス制御: IAM ポリシー、SQS ポリシーによる制御が可能
  * メタデータ: メッセージにメタデータを付与可能
  * メトリクス: メッセージ数などのメトリクスがある。コンシューマのスケールなどの用途で利用できる
* SQS, Kinesis の使い分け
  * ストリーミングデータであれば Kinesis を用いる



# 参考

* Document
  * [Amazon Simple Queue Serviceとは?](https://docs.aws.amazon.com/ja_jp/AWSSimpleQueueService/latest/SQSDeveloperGuide/welcome.html)
* Black Belt
  * [Amazon Simple Queue Service](https://pages.awscloud.com/rs/112-TZM-766/images/20190717_AWS-BlackBelt_AmazonSQS.pdf)


