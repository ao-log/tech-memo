
[The Amazon Builders' Library - Amazon はソフトウェアをどのように構築し、運用するのか](https://aws.amazon.com/jp/builders-library/?cards-body.sort-by=item.additionalFields.sortDate&cards-body.sort-order=desc&awsf.filter-content-category=*all&awsf.filter-content-type=*all&awsf.filter-content-level=*all)


[運用の可視性を高めるために分散システムを装備する](https://aws.amazon.com/jp/builders-library/instrumenting-distributed-systems-for-operational-visibility/?did=ba_card&trk=ba_card)

* 平均レイテンシーの他、99.9 パーセンタイルや 99.99 パーセンタイルなど、 レイテンシーの異常値にも注目。
* チームは関連するすべてのサービスの運用パフォーマンスに目標を設定
* サービス所有者として、システムの動作を測定する必要があります
* 指定されたトレース ID のシステム間でインストルメンテーションを収集するには、必要に応じて事後に、または AWS X-Ray のようなサービスを使用してほぼリアルタイムで行うことができます
* すべての作業のタイマーとカウンターはすべてログファイルに書き込まれます。そこから、ログが処理され、他のシステムによって、後から集約メトリックが計算されます。
* ログ記録の良い習慣
  * 作業単位ごとに 1 つのリクエストのログエントリを作成する。


[サーバーレスアプリケーション開発におけるエラーハンドリング ~ イベント駆動のデータ加工、連携処理パターン ~](https://aws.amazon.com/jp/builders-flash/202308/serverless-error-handling-3/?awsf.filter-name=*all)

* イベントの生成をトリガーとして、Push 方式で非同期に呼び出すのが特徴
* プロデューサの例は S3、イベントルーターの例は SNS, EventBridge、コンシューマの例は Lambda
* エラーハンドリング
  * イベントルーター
    * データを送信できない場合は AWS 管理のイベントルーターがリトライを試みる
  * Lambda 関数
    * 実行前の呼び出し時のエラーへの対応。Maximum Event Age により設定可能
    * Lambda 関数がエラーを返した場合の対応。Maximum Retry Attempts により設定可能
* データ送信できない場合やリトライポリシーを超えた場合
  * SQS
    * Dead Letter Queue (DLQ)
  * Lambda
    * Lambda Destinations


[負荷制限を使用して過負荷を回避する](https://aws.amazon.com/jp/builders-library/using-load-shedding-to-avoid-overload/?did=ba_card&trk=ba_card)

* サーバーが過負荷になると、受信リクエストをトリアージして、どのリクエストを受け入れ、どのリクエストを拒否するかを決定する機会がある
* 優先順位付けとスロットリングを一緒に使用して、サービスを過負荷から保護


[負荷テスト on AWS のすすめ](https://aws.amazon.com/jp/builders-flash/202309/distributed-test-on-aws-2/?awsf.filter-name=*all)

* 一口に負荷試験と言っても目的に応じたさまざまな種類がある。ピーク負荷試験、限界性能試験、長時間負荷試験など


[マイグレーションの勉強方法を聞いてみた。](https://aws.amazon.com/jp/builders-flash/202309/way-to-learn-migration/?awsf.filter-name=*all)

* プロジェクト管理スキルも重要


#### 外部記事

[DevelopersIOブログの記事配信がCloudFront経由になりました](https://dev.classmethod.jp/articles/developersio-cdn-cloudfront/)




