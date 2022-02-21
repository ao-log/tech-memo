
[The Amazon Builders' Library - Amazon はソフトウェアをどのように構築し、運用するのか](https://aws.amazon.com/jp/builders-library/?cards-body.sort-by=item.additionalFields.sortDate&cards-body.sort-order=desc&awsf.filter-content-category=*all&awsf.filter-content-type=*all&awsf.filter-content-level=*all)

[運用の可視性を高めるために分散システムを装備する](https://aws.amazon.com/jp/builders-library/instrumenting-distributed-systems-for-operational-visibility/?did=ba_card&trk=ba_card)

* 平均レイテンシーの他、99.9 パーセンタイルや 99.99 パーセンタイルなど、 レイテンシーの異常値にも注目。
* チームは関連するすべてのサービスの運用パフォーマンスに目標を設定
* サービス所有者として、システムの動作を測定する必要があります
* 指定されたトレース ID のシステム間でインストルメンテーションを収集するには、必要に応じて事後に、または AWS X-Ray のようなサービスを使用してほぼリアルタイムで行うことができます
* すべての作業のタイマーとカウンターはすべてログファイルに書き込まれます。そこから、ログが処理され、他のシステムによって、後から集約メトリックが計算されます。
* ログ記録の良い習慣
  * 作業単位ごとに 1 つのリクエストのログエントリを作成する。

[負荷制限を使用して過負荷を回避する](https://aws.amazon.com/jp/builders-library/using-load-shedding-to-avoid-overload/?did=ba_card&trk=ba_card)

* サーバーが過負荷になると、受信リクエストをトリアージして、どのリクエストを受け入れ、どのリクエストを拒否するかを決定する機会がある
* 優先順位付けとスロットリングを一緒に使用して、サービスを過負荷から保護


