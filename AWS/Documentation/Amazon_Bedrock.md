
[Amazon Bedrock](https://aws.amazon.com/jp/bedrock/)

* 様々な基盤モデル(FM) を使用できる
  * [サポートされている FM の一覧](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/models-supported.html)
  * FM をベースにカスタマイズする。トレーニング用のデータは S3 上に配置する
  * モデルへのアクセスをリクエストする必要がある
* ユースケース
  * テキスト生成
  * バーチャルアシスタント
  * テキスト、画像検索
  * テキスト要約
  * 画像生成
* 微調整や検索拡張生成(RAG) による個人向けカスタマイズが可能
* Chat
  * FM もしくはカスタマイズされたモデルから選ぶことができる
* 用語
  * ファインチューニング: 学習済みの基盤モデルに対し、追加で学習用のデータを渡しより望ましい結果を得られるようにモデルを改良するアプローチ
  * プロンプトエンジニアリング: 出力をコントロールするアプローチ


[Amazon Bedrock とは](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/what-is-bedrock.html)


## プロンプトエンジニアリング

[プロンプトエンジニアリングの概念](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/prompt-engineering-guidelines.html)

* モデルごとにプロンプトガイドが用意されている場合が多い

[Amazon Bedrock のインテリジェントなプロンプトルーティングを理解する](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/prompt-routing.html)

* リクエストごとに各モデルのレスポンス品質を動的に予測し、最適なレスポンス品質のモデルにリクエストをルーティングできる

[プロンプトを設計する](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/design-a-prompt.html)

* シンプルで明確な指示を出すと良い
* 質問または指示をプロンプトの最後に記載する
* どのような出力を希望するかを指示する
* ステップバイステップで考えるように指示すると、段階的に思考し、期待する結果を得られる場合がある



# Black Belt

[Amazon Bedrock Overview](https://pages.awscloud.com/rs/112-TZM-766/images/AWS-Black-Belt_2024_Amazon-Bedrock-Overview_v1.pdf)

* 幅広い基盤モデルを選択可能
* 独自データへの対応方法
  * RAG
    * Knowledge Bases for Amazon Bedrock
      * S3 上のドキュメントをベクトル化しベクトル DB に格納。Bedrock ではベクトル DB から情報取得し、クエリ + ベクトル DB から取得した内容を基盤モデルに問い合わせる
  * 基盤モデルのカスタマイズ
    * ファインチューニング: ラベル付きデータを学習し、特定タスクの精度を高める。出力スタイルの変更などモデルの振る舞いのカスタマイズもできる
    * Continiued Pre-training: 大量のラベルなしデータを学習し、新たなドメイン知識を習得させる
  * Custom Model Import for Amazon Bedrock
    * 別環境で作成したモデルをインポートし Bedrock API で使用できる
* Agent
  * ユーザーの⼊⼒を複数の⼩さなタスクに分割し、タスクごとに適切な API を呼び出すことで回答を⽣成させるアプローチ
  * Agents for Amazon Bedrock
* モデルの評価
  * Model Evaluation on Amazon Bedrock
  * バッチ推論により、非同期に実行し結果を S3 に格納できる
* 適切な利⽤状況の維持
  * Guardrails for Amazon Bedrock
    * 拒否トピック、コンテンツフィルターなど
  * Model Invocation Logging
    * 入力プロンプト、回答を S3 に記録できる
* 安定稼働
  * Provisioned Throughput
* 迅速な実験と評価
  * Amazon Bedrock Studio


[Amazon Bedrock モデル推論 準備編](https://pages.awscloud.com/rs/112-TZM-766/images/AWS-Black-Belt_2024_Amazon-Bedrock-Model-Inference-a_0909_v1.pdf)

[Amazon Bedrock モデル推論 実践編](https://pages.awscloud.com/rs/112-TZM-766/images/AWS-Black-Belt_2024_Amazon-Bedrock-Model-Inference-b_0909_v1.pdf)

* 基盤モデル
  * 膨大なテキストデータで学習されたものが LLM
  * ライフサイクル
    * アクティブ、レガシー、EOL のいずれか
    * EOL になると、そのモデルは実行できなくなる
  * AWS アカウント内で使用できるように、モデルへのアクセスをリクエストする必要がある
* token
  * テキストを LLM に入力するために変換した後の基本的な単位
* 推論パラメータ
  * Temperature: 値が低いほど出現確率が高い token を採用しやすくなる
  * Top P: 出現確率が⾼い順に並べ、合計 P % になるまでの候補から次のトークンを採⽤
  * Top K: 現確率が⾼い上位 K 個から次のトークンを採⽤
  * 堅実な回答をさせるか、想像的な回答をさせるかで調整する
* モデルの実行
  * マネジメントコンソール
    * playground: チャットやテキストの生成
    * Image playground: イメージの生成や編集
  * InvokeModel
    * request body の JSON を組み立てて、InvokeModel API を実行
  * InvokeModelWithReponseStream
    * レスポンスの body が逐一追加されるのでループで取得するなどの処理を組める
    * 体感待ち時間の短縮などのメリットがある
  * Converse API
    * モデルごとの推論パラメータやプロンプト形式の違いを吸収してくれる。モデル側で対応していれば、同じようなリクエストパラメータで対応できる
  * バッチ推論
    * オンデマンドの 50% の料金でスロットリングを回避できる
* モデルを安定的に使用するには？
  * リージョン、アカウント単位のクォータを確認
  * リトライ処理を実装
  * クロスリージョン推論の使用を検討
  * 各種検討の結果不足する場合は Provisioned Throughput を検討
  * 時間をかけて良い、まとまったデータ処理に対してはバッチ推論を検討

[Amazon Bedrock Knowledge Bases](https://pages.awscloud.com/rs/112-TZM-766/images/AWS-Black-Belt_2024_Amazon-Bedrock-Knowledge-Bases_0920_v1.pdf)

* ベクトル検索
  * ベクトル間の距離や⾓度に基づいて関連度を判断する検索⼿法
* Amazon Bedrock Knowledge Bases
  * マネジメントコンソールから簡単に設定可能
  * データソースとして S3 を指定。ベクトルデータベースとして、OpenSearch Serveless などを選択できる
  * 埋め込みへの変換に、複数の埋め込みモデルから選択できる
* Chat with your document 機能
  * ベクトルデータベースの作成なしで RAG を利用可能
* RetrieveAndGenerate API
  * RAG を利用しつつ LLM が結果を応答する
* Retrieve API
  * RAG 部分のみ実行。そのため、LLM への問い合わせは別途実装が必要s



# 参考

* Black Belt
  * [Amazon Bedrock Overview](https://pages.awscloud.com/rs/112-TZM-766/images/AWS-Black-Belt_2024_Amazon-Bedrock-Overview_v1.pdf)
  * [Amazon Bedrock モデル推論 準備編](https://pages.awscloud.com/rs/112-TZM-766/images/AWS-Black-Belt_2024_Amazon-Bedrock-Model-Inference-a_0909_v1.pdf)
  * [Amazon Bedrock モデル推論 実践編](https://pages.awscloud.com/rs/112-TZM-766/images/AWS-Black-Belt_2024_Amazon-Bedrock-Model-Inference-b_0909_v1.pdf)
  * [Amazon Bedrock Knowledge Bases](https://pages.awscloud.com/rs/112-TZM-766/images/AWS-Black-Belt_2024_Amazon-Bedrock-Knowledge-Bases_0920_v1.pdf)
  * [Amazon Bedrock Agents ⾃律型 AI の実現に向けて: 検討編](https://pages.awscloud.com/rs/112-TZM-766/images/AWS-Black-Belt_2024_Amazon-Bedrock-Agents_0930_v1.pdf)
  * [Amazon Bedrock Agents ⾃律型 AI の実現に向けて: 動作理解編](https://pages.awscloud.com/rs/112-TZM-766/images/AWS-Black-Belt_2024_Amazon-Bedrock-Agents-b_1029_v1.pdf)
  * [Amazon Bedrock Agents ⾃律型 AI の実現に向けて: 開発・運⽤編](https://pages.awscloud.com/rs/112-TZM-766/images/AWS-Black-Belt_2024_Amazon-Bedrock-Agents-c_1118_v1.pdf)



#### AWS 外部記事

[生成AIサービス「Amazon Bedrock」とは？できること・使い方](https://business.ntt-east.co.jp/content/cloudsolution/column-487.html)


