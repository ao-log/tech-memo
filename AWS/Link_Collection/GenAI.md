

# Amazon Q Developer CLI

[Amazon Q Developer CLI での超高速な新しいエージェント型のコーディング体験](https://aws.amazon.com/jp/blogs/news/introducing-the-enhanced-command-line-interface-in-amazon-q-developer/)

* q chat でプロンプトで指示することで、プロンプトで依頼した内容に沿って環境構築やコードを作成してくれる
* AWS 内のリソースの検索のような仕事も、指示できる


[生成 AI で生成 AI アプリケーションを生成しよう！](https://aws.amazon.com/jp/blogs/news/generate-genai-apps-using-genai/)

* Amazon Q Developer CLI を使った コーディング支援で動くものを作ることは可能。ただし、プロダクションレベルの可読性・保守性は別途検討が必要
* Amazon Q Developer CLI が必要
  * IAM Identity Center の ID か AWS Builder ID が必要
* q chat と入力して、Amazon Q Developer CLI を起動
  * プロンプトで改修内容を命令すると、ブランチを作成するなど改修を行なってくれる。変更の都度、対話型で yes/no を設定できる
  * 成果物でエラーが出る場合は、エラー内容をプロンプトで伝える


[[アップデート]Amazon Q Developer CLIでMCPがサポートされました！](https://dev.classmethod.jp/articles/amazon-q-developer-cli-mcp-support-hands-on/)

* 設定ファイル `~/.aws/amazonq/mcp.json`
* [AWS Documentation MCP Server](https://awslabs.github.io/mcp/servers/aws-documentation-mcp-server/)


# Bedrock

[Amazon Bedrock Agents で MCP サーバーを活用する](https://aws.amazon.com/jp/blogs/news/harness-the-power-of-mcp-servers-with-amazon-bedrock-agents/)

* Model Context Protocol (MCP) は、LLM がデータソースやツールに接続するための標準化された方法を提供する
* クライアント/サーバのアーキテクチャ
* MCP サーバは特定の用途に関するプロンプト処理に特化。用途ごとに MCP サーバがある


[Nyantech マルチエージェントでぴったりのオトモを見つけよう !](https://aws.amazon.com/jp/builders-flash/202502/create-nyantech-multi-agent/)

* 司令塔となるエージェントから、各役割に特化したエージェントを必要に応じて呼び出す仕組み
* show trace で司令塔エージェントでの処理内容を確認できる
* Knowledge Bases でナレッジベースを作成。S3 上に PDF を格納しておく。PDF にはお猫様の画像及びデータが含まれている。Amazon OpenSearch Service Serverless のコレクションが作成される
* Agent にはナレッジベースの指定を行い、どのような出力を行うかの指示を行う
* 司令塔用のエージェントに指示を行う。どの場合にどのエージェントを使用するかも指示に含める



[Amazon Bedrock の新機能マルチエージェントで「わが家の AI 技術顧問」を作ろう !](https://aws.amazon.com/jp/builders-flash/202503/create-ai-advisor-with-bedrock/)

* マルチエージェントシステムは、複数のエージェントを駆使して目的を達成するアプローチ
* Bedrock の機能でナレッジベースを作成できる。データソースは S3 上の PDF、Vector Store は Amazon Aurora PostgreSQL Serverless。これで簡単に RAG API を作成できる。RAG 担当の Agent を作成する
* 2 個目の Agent が外部サイト検索用。Agent のアクショングループで作成される Lambda 関数内で tavily_search を用いて外部サイト検索
* 監督用の Agent を作成し、上記 2 個の Agent に対して適切に使い分けるようにする指示を与えておく


# Q Developer

[Amazon Q Developer CLIにAWSアカウントを調査・操作させてみた](https://dev.classmethod.jp/articles/exploring-aws-with-q-cli/)

* CloudFormation スタックから drawio 形式で AWS 構成図を書いてくれる
* アクセスログの分析ができる
* プロンプトにより CloudFormation テンプレートの生成


# Others

[生成系 AI アプリケーションでベクトルデータストアが果たす役割とは](https://aws.amazon.com/jp/blogs/news/the-role-of-vector-datastores-in-generative-ai-applications/)

* ベクトルデータストアは、ベクトルを大規模に保存し、問い合わせを行うためのシステムであり、効率的な最近傍クエリアルゴリズムと適切なインデックスにより、データ検索を改善


[エンタープライズにおける Amazon Bedrock による生成 AI のオペレーティングモデル](https://aws.amazon.com/jp/blogs/news/generative-ai-operating-models-in-enterprise-organizations-with-amazon-bedrock/)


[アート引越センター株式会社様の AI 活用事例 「 AI 見積りアプリによる引越し見積もりの自動化の実現」のご紹介](https://aws.amazon.com/jp/blogs/news/i3design-art-ai-auto-estimation/)

* 室内を撮影すると、デバイス上で 3D モデルが作成される
* AI が 3D モデル内の家具、家電を検出し、自動的に見積もりが行われる
* 3D モデルは S3 に格納
* EventBridge にて推論処理実行。PointNeXt をベースとした学習済みモデル。Amazon SageMaker Serverless Inference によるサーバーレス環境で非同期推論を実行
* SNS → SQS → Step Functions にて DB に結果格納
* Web 画面は ELB + ECS Fargate


[[プレビュー] AWS Knowledge MCP Serverを使ってみた](https://dev.classmethod.jp/articles/aws-knowledge-mcp-server-available-preview/)


