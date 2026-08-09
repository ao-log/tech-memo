
# Document

[推論のためのモデルをデプロイする](https://docs.aws.amazon.com/ja_jp/sagemaker/latest/dg/deploy-model.html)

* リアルタイムエンドポイント: インタラクティブで低レイテンシーの要件がある推論ワークロードに最適
* サーバーレスエンドポイント: サーバレス。コールドスタートを許容できるワークロードに最適
* バッチ変換: 大量のデータが事前に利用可能で、永続的なエンドポイントが必要ない場合のオフライン処理に最適
* 非同期エンドポイント: ペイロードサイズが大きく (最大 1 GB)、処理時間が長い (最大 1 時間)、ほぼリアルタイムのレイテンシー要件があるリクエストに最適





# Black Belt

[Black Belt - Amazon SageMaker Basic](https://pages.awscloud.com/rs/112-TZM-766/images/20190206_AWS_BlackBelt_SageMaker_part1.pdf)

* 学習用コード、トレーニング、推論のワークフロー全体をカバー
* ノードブックインスタンスにて、Jupyter Notebook にアクセス
  * サンプルノートブックも多数ある
  * SageMaker SDK により各種操作が可能
    * fit メソッドによりジョブ実行
    * 学習後に deploy メソッドにより推論エンドポイントを作成
    * predict メソットにより推論実行
* 各コンポーネント
  * 開発: Jupyter Notebook や機械学習ライブラリ群がインストール済みのインスタンス
  * 学習: API 実行により学習用インスタンスが起動。複数ジョブの同時実行やハイパーパラメータチューニングに対応
    * コンテナ上で実行される。学習データ、学習済みモデルは S3 上に配置
  * 推論: Auto Scaling や AB テストにも対応
    * コンテナ上で実行される


[Black Belt - Amazon SageMaker Advance](https://pages.awscloud.com/rs/112-TZM-766/images/20190213_AWS_BlackBelt_SageMaker_part2.pdf)

* Amazon SageMaker Ground Truth
  * データにラベルを付与するアノテーション作業を支援するサービス
* 推論用に GPU インスタンスを用意するのは、リソース効率面で高コスト。Elastic Inference のようなアクセラレータを検討する

