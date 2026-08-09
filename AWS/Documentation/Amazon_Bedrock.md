
# Document

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


[クイックスタート](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/getting-started.html)

* Bedrock API を環境変数にセットし、モデルを呼び出す
* Boto3 の場合は converse API



## モデル

[モデルの概要](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/model-cards.html)


[詳細なモデル情報](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/models.html)


[API の互換性](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/models-api-compatibility.html)

* InvokeModel
* InvokeModelWithResponseStream: レスポンスがリアルタイムストリーム
* Converse: モデルに依存しない統一的なインターフェース


[エンドポイントの可用性](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/models-endpoint-availability.html)

* `bedrock-mantle.{region}.api.aws`: OpenAI 互換エンドポイントと Anthropic Messages API
* `bedrock-runtime.{region}.amazonaws.com`: InvokeModel/Converse/Chat Completions/Messages API を使用して Amazon Bedrock でホストされているモデルの推論リクエストを行うためのリージョン固有のエンドポイント


[リージョン別の可用性](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/models-region-compatibility.html)

* リージョン内、地理的 (Geo)、グローバルの 3 つのオプションがある


### モデルごとの使用方法

[モデル推論パラメータとレスポンス](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/model-parameters.html)


[Anthropic Claude Messages API](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/model-parameters-anthropic-claude-messages.html)

* システムプロンプト: 特定の目標やロールを指定するなど、Anthropic Claude に、コンテキストと手順を提供できる


[モデルのライフサイクル](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/model-lifecycle.html)

* アクティブ
* レガシー
* End-of-Life (EOL) 



## 構築

[Invoke API を使用した推論](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/inference-api.html)


[推論パラメータでレスポンスの生成に影響を与える](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/inference-parameters.html)

* 温度: 予測出力の確率分布の形状に影響。モデルがより確率の低い出力を選択する可能性にも影響
* トップ K: モデルが次のトークンについて検討する最も可能性の高い候補の数
* トップ P: モデルが次のトークンについて考慮する最も可能性の高い候補のパーセンテージ
* レスポンスの長さ: 生成されたレスポンスで返されるトークンの最小数または最大数を指定
* ペナルティ: レスポンス内の出力にどの程度ペナルティを課すかを指定
* 停止シーケンス: モデルがそれ以上トークンを生成しないようにする文字シーケンスを指定


[API キー](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/api-keys.html)

* AWS 認証情報の代わりにベアラートークンを使用して API リクエストを認証rできる



## モデルのカスタマイズ

[モデルをカスタマイズしてユースケースのパフォーマンスを向上させる](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/custom-models.html)

* 教師ありファインチューニング: ラベル付きデータを提供してモデルをトレーニング
* 強化ファインチューニング
* 蒸留: より大規模でインテリジェントなモデル (教師と呼ばれる) から、より小規模かつ高速でコスト効率の高いモデル (生徒と呼ばれる) に知識を移行


[事前トレーニング済みのモデルを Amazon Bedrock にインポートする](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/import-pre-trained-model.html)

* Amazon S3 バケットからモデルをインポートする場合は、モデルファイルを Hugging Face 重み形式で準備



## セキュリティ、ガードレール、オブザーバビリティ

### ガードレール

* ガードレールを使用するには `InvokeModel` のパラメータで `guardrailIdentifier`, `guardrailVersion` を指定する

[Amazon Bedrock ガードレールを使用して有害なコンテンツを検出してフィルタリングする](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/guardrails.html)

* コンテンツフィルター
  * 入力、レスポンスの有害な内容を検出してフィルタリングできる
  * ヘイト、侮辱、性的、暴力、不正行為などのカテゴリごとにフィルタ強度を設定出来る
  * プロンプト攻撃に対するフィルタ強度を設定できる
* 拒否されたトピック: 定義したトピックが含まれる場合はブロックできる
* 単語フィルター: 指定したワードに合致した場合にブロックできる
* 機密情報フィルター: 個人情報をブロックまたはマスクできる
* コンテキストグラウンディングチェック: ハルシネーションのように不正確な内容を検出できる
* 自動推論チェック: ハルシネーションを検出したり、修正を提案したりできる


[モデル推論リクエストで特定のガードレールの使用を強制する](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/guardrails-permissions-id.html)

以下の例のようなポリシーでガードレールの使用を強制できる
```json
{
    "Version":"2012-10-17",
    "Statement": [
        {
            "Sid": "InvokeFoundationModelStatement1",
            "Effect": "Allow",
            "Action": [
                "bedrock:InvokeModel",
                "bedrock:InvokeModelWithResponseStream"
            ],
            "Resource": [
                "arn:aws:bedrock:us-east-1::foundation-model/*"
            ],
            "Condition": {
                "StringEquals": {
                    "bedrock:GuardrailIdentifier": "arn:aws:bedrock:us-east-1:123456789012:guardrail/guardrail-id:1"
                }
            }
        },
```


[ガードレールをデプロイする](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/guardrails-deploy.html)

* バージョンを作成する。呼び出し側ではバージョンもあわせて指定


### モニタリング

[CloudWatch Logs と Amazon S3 を使用してモデル呼び出しをモニタリングする](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/model-invocation-logging.html)

* すべての呼び出しの呼び出しログ、モデル入力データ、モデル出力データを収集できる



## 容量とパフォーマンス

[容量とパフォーマンス](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/capacity-limits-cost-optimization.html)


[バッチ推論を使用して複数のプロンプトを処理する](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/batch-inference.html)


[クロスリージョン推論によりスループットを向上させる](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/cross-region-inference.html)


[推論プロファイルを使用してモデル呼び出しリソースを設定する](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/inference-profiles.html)

* 推論時にプロファイルを指定する。そうすることで、ユーザーごとのコスト、モデルの使用量を追跡できる
* 推論プロファイルの種類
  * クロスリージョン（システム定義）推論プロファイル: Bedrockで事前定義され、モデルへのリクエストをルーティングできる複数のリージョンを含む推論プロファイル
  * アプリケーション推論プロファイル: コストとモデルの使用を追跡するためにユーザーが作成する推論プロファイル


[Amazon Bedrock のプロビジョンドスループットでモデル呼び出し容量を増やす](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/prov-throughput.html)

* プロビジョンドスループットを購入して、固定コストでモデルのより高いレベルのスループットをプロビジョニングできる
* 次の例のようにモデル ID にプロビジョンドモデルの ARN を指定する
```
aws bedrock-runtime invoke-model \
    --model-id ${provisioned-model-arn} \
    ...
```


[モデル推論を高速化するためのプロンプトキャッシュ](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/prompt-caching.html)



## その他の機能

### Amazon Bedrock Data Automation

[Amazon Bedrock Data Automation を使用して非構造化データを有意義なインサイトに変換する](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/bda.html)

* ドキュメント、動画などの非構造コンテンツからインサイトを抽出できる


[Bedrock Data Automation の仕組み](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/bda-how-it-works.html)

* 標準出力: デフォルトではドキュメントのサマリ、音声からの文字起こしなどを行う
* カスタム出力
  * ドキュメント、音声、イメージのみ対応
  * ブループリントを使用して抽出する情報を正確に定義
  * ブループリントは、ファイルから取得するフィールドのリストで構成されている
* プロジェクト
  * 出力内容を指定
  * `InvokeDataAutomationAsync` で呼び出される


[プロジェクト使用中にドキュメントを分割する](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/bda-document-splitting.html)

* ドキュメントの分割をサポート
* ブループリントのマッチングのされ方。このようにマッチングされるので、多種のドキュメントを単一のバッチで処理できる
  * 以下情報を元に判断
    * ブループリント名
    * ブループリントの説明
    * ブループリントのフィールド
* ベストプラクティス
  * マッチングに役立つように、ブループリントの名前と説明を明確かつ詳細に記述
  * 関連するブループリントを複数提供すると、BDA は最適なものを選択してくれる。ドキュメントの形式が大幅に異なる場合に有効
  * プロジェクトに同じタイプの 2 つのブループリント (2 つの W2 ブループリントなど) を含めないこと。パフォーマンス低下につながるため

[ブループリントを使用してさまざまな IDP タスクを実現する](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/idp-cases.html)

* カスタム出力用のブループリントを作成できる
* 分類、抽出、正規化、変換、検証用のブループリントを作成できる


[グラウンドトゥルースでブループリントを最適化する](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/bda-optimize-blueprint-info.html)

* 最適化プロセスを開始するときは、サンプルアセットとそれに対応するグラウンドトゥルースデータ、つまり各フィールドに対して抽出する予定の正しい値を指定する


### Amazon Bedrock Knowledge Bases

[Amazon Bedrock ナレッジベースでデータを取得して AI レスポンスを生成する](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/knowledge-base.html)

* RAG を構築できる


[Amazon Bedrock ナレッジベースの仕組み](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/kb-how-it-works.html)

* 元のデータをチャンクごとに区切り、埋め込みモデルにより埋め込みを作成してベクトルデータベースに保存する
* 実行時はクエリ内容を埋め込みに変換する


[Amazon Bedrock ナレッジベースを使用してデータソースから情報を取得する](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/kb-how-retrieval.html)

* Retrieve: レスポンスを配列として返す
* RetrieveAndGenerate: 自然言語のレスポンス、特定のソースチャンクへの引用を返す
* GenerateQuery: 自然言語のユーザークエリを構造化データストアに適した形式のクエリに変換


[ナレッジベースのコンテンツのチャンキングの仕組み](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/kb-chunking.html)

* 標準のチャンキング: デフォルトのチャンキングは約 300 トークンのテキストチャンクに分割
* 階層的チャンキング: 情報を子チャンク、親チャンクのネスト構造で整理。検索中は子チャンクを最初に取得し、包括的なコンテキストをモデルに提供するためにより広範な親チャンクに置き換えられる。包括的なコンテキスト提供が目的
* セマンティックチャンキング: テキストを意味のあるチャンクに分割
* マルチモーダルコンテンツのチャンキング: 


### Amazon Bedrock Model Evaluation

[Amazon Bedrock リソースのパフォーマンスを評価する](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/evaluation.html)

* パフォーマンスと有効性を評価
* ジョブを実行することで評価を行い helpfulness などのスコアをレポートできる
* 評価の種類
  * 自動モデル評価ジョブ
  * 人間ベースのモデル評価
    * S3 出力バケットで CORS 設定を指定する必要がある
  * LLM-as-a-judge


[Amazon Bedrock での自動モデル評価ジョブの作成](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/evaluation-automatic.html)

[Model evaluation task types in Amazon Bedrock](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/model-evaluation-tasks.html)

* モデル評価ジョブごとに 1 つのタスクタイプを選択できる
* 組み込みのデータセットを選ぶこともできる


[LLM-as-a-judge を使用してモデルのパフォーマンスを評価する](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/evaluation-judge.html)

[メトリクスを使用してモデルのパフォーマンスを把握する](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/model-evaluation-metrics.html)

* 正確性 (Builtin.Correctness)、完全性 (Builtin.Completeness) などのメトリクスのスコアで評価する



### Amazon Bedrock Prompt Management

[Amazon Bedrock でプロンプト管理を使用して再利用可能なプロンプトを構築して保存する](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/prompt-management.html)

* `InvokeModel` のパラメータで `modelId` にプロンプトの ARN、`promptVariables` に変数を指定できる


### Amazon Bedrock Agent

* エージェントは、基盤モデル (FM)、データソース、ソフトウェアアプリケーション、ユーザーとの会話の間のインタラクションを調整
* 自律的に何を呼び出して実行するかの計画を立てる


### プロンプトエンジニアリング

[プロンプトエンジニアリングの概念](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/prompt-engineering-guidelines.html)

* モデルごとにプロンプトガイドが用意されている場合が多い

[Amazon Bedrock のインテリジェントなプロンプトルーティングを理解する](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/prompt-routing.html)

* リクエストごとに各モデルのレスポンス品質を動的に予測し、最適なレスポンス品質のモデルにリクエストをルーティングできる

[プロンプトを設計する](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/design-a-prompt.html)

* シンプルで明確な指示を出すと良い
* 質問または指示をプロンプトの最後に記載する
* どのような出力を希望するかを指示する
* ステップバイステップで考えるように指示すると、段階的に思考し、期待する結果を得られる場合がある


### Amazon Bedrock Flows

[Amazon Bedrock フローを使用してエンドツーエンドの生成 AI ワークフローを構築する](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/flows.html)


[フローのノードタイプ](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/flows-nodes.html)

* すべてのフローには Flow input ノードが 1 つだけ含まれており、そこから始める必要がある
* Prompt ノード: フローで使用するプロンプトを定義
* Agent ノード: プロンプトをエージェントに送信できる
* ナレッジベースノード: Amazon Bedrock Knowledge Bases からナレッジベースにクエリを送信できる
* S3 Storage ノード: Amazon S3 バケットへのフローにデータを保存できる
* S3 Retrieval ノード: Amazon S3 ロケーションからデータを取得してフローに導入
* Lambda function ノード: ビジネスロジックを実行するようにコードを定義できる Lambda 関数を呼び出す


## リファレンス/アドバンスト

[Amazon Bedrock のリランカ―モデルを使用してクエリレスポンスの関連性を向上する](https://docs.aws.amazon.com/ja_jp/bedrock/latest/userguide/rerank.html)

* リランカーモデルでは、チャンクのクエリへの関連性を計算し、計算したスコアに基づいて結果を並べ替える



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

* モデル
  * ベースモデル: 
  * カスタムモデル: Bedrock で fine-tuning もしくは continuous-pre-training したもの。プロビジョンドスループットのみ
  * インポートされたモデル: Bedrock 以外の環境でカスタマイズしたもの。オンデマンド実行のみ

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

[Amazon Bedrock Agents ⾃律型 AI の実現に向けて: 検討編](https://pages.awscloud.com/rs/112-TZM-766/images/AWS-Black-Belt_2024_Amazon-Bedrock-Agents_0930_v1.pdf)

* Agent を使用することでタスクの実行計画も Agent 側になる。Agent を使用しない場合は人間が組み立てる必要がある
* ワークフローが自明な場合は AI を使用せず Step Functions などでワークフローを組んだ方がいい
* (自分用メモ) 要は普通の LLM のチャット形式で使えるのだが、独自の業務フローなどを実行したいユースケースに向いている。社内システムと連携する Lambda 関数や独自情報が入った Knowledge Base を構築しておくことで、それらを必要に応じて必要となるアウトプットを出してくれる

[Amazon Bedrock Agents ⾃律型 AI の実現に向けて: 動作理解編](https://pages.awscloud.com/rs/112-TZM-766/images/AWS-Black-Belt_2024_Amazon-Bedrock-Agents-b_1029_v1.pdf)

* 前処理、オーケストレーション、後処理の流れで進んでいく
* 前処理では害を及ぼす内容がないかなどを確認し、後続処理に渡すかどうかを判断
* オーケストレーションにて自律計画を立てて実行していく
  * Action Group により Lambda 関数を呼び出して所定の処理を行なったりできる
  * Bedrock Knowledge Bases もできる
  * コードを生成し、サンドボックス内で実行する
  * 不足している情報があると判断すると、ユーザーに質問で返す
* 後処理で、好ましい回答となるように調整

[Amazon Bedrock Agents ⾃律型 AI の実現に向けて: 開発・運⽤編](https://pages.awscloud.com/rs/112-TZM-766/images/AWS-Black-Belt_2024_Amazon-Bedrock-Agents-c_1118_v1.pdf)

* 開発フローは、初期作成、変更、準備、テスト、デプロイの順
* 準備
  * 必要に応じて以下を用意する。これらにより社内システム連携なども実現できる
    * Lambda Parser 用や Action Group 用の Lambda 関数の実装
    * Knowledge Base の作成    
* テストではチャット形式で試験できる
* 呼び出す時はエイリアスを指定。エイリアスとバージョンを紐付ける


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


