

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


[より豊かなコンテキストのための Model Context Protocol (MCP) による Amazon Q Developer CLI の拡張](https://aws.amazon.com/jp/blogs/news/extend-the-amazon-q-developer-cli-with-mcp/)


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


[Claude Code on AWS パターン解説 – Amazon Bedrock / AWS Marketplace](https://aws.amazon.com/jp/blogs/news/claude-code-on-aws-patterns/)

* Claude Code on AWS パターン 1: Amazon Bedrock との連携
  * Claude Code 起動時に環境変数を設定しておけば Bedrock が使用される
  * OIDC provider で認証 → ID トークンを Amazon Cognito に渡す → Cognito から一時的な AWS クレデンシャルを生成、の流れで認証可能
  * 利用状況などを確認できる Dashboard がある
  * token 使用量に基づいた従量課金。スモールスタートに適している
* パターン2: AWS Marketplace での購入
  * GUI アプリケーションとしての使用も可能
  * コーディングだけでなく、設計、レビュー、ドキュメント作成なども効率化したい場合、このパターンが適している
  * 予測可能で予算が立てやすいコスト体系


# Kiro

[Kiro 導入ガイド：始める前に知っておくべきすべてのこと](https://aws.amazon.com/jp/blogs/news/kiroweeeeeeek-in-japan-day-1-implementation-guide/)

* Amazon Q Developer と Kiro は独立したサブスクリプション体系を持つ別製品
* 既に Amazon Q Developer Pro サブスクリプションを利用している場合は Kiro CLI と Kiro IDE が Kiro Pro プラン相当の範囲内で有効
* Q Developer CLI は Kiro CLI にアップデートされる
* AWS 請求に一本化したい場合は、AWS IAM Identity Center での認証が必須。GitHub などでも認証できるが、その場合の支払い方法はクレジットカード


[Amazon Q Developer の IDE プラグインから Kiro に乗り換える準備](https://aws.amazon.com/jp/blogs/news/kiroweeeeeeek-in-japan-day-2-q-dev-ide-to-kiro/)

* Rules 機能が進化した Steering 機能。開発標準を定義できるので、チーム開発時に有用
* コードレビュー機能はない。Amazon Inspector コードセキュリティによるコードレビュー機能などの方法を検討する必要がある
* Spec 昨日では、要件（Requirements）、設計（Design）、タスク（Tasks）の 3 フェーズを経て、構造化された開発プロセスを実現


[Kiro を組織で利用するためのセキュリティとガバナンス](https://aws.amazon.com/jp/blogs/news/kiroweeeeeeek-in-japan-day-3-security-governance/)

* IAM Identity Center を用いた認証が可能
* セキュリティ、プライバシー
  * 使用 Tier によりプロンプト、返答が保存されるリージョンが異なる
  * データ転送時、保管時は暗号化される
  * 使用 Tier により、利用情報が製品の改善に使用される場合がある。オプトアウト可能。Kiro for enterprise の場合は自動的にオプトアウト
* PrivateLink を介した接続が可能


[インフラエンジニアのあなたも！Shell スクリプト開発で Kiro を使ってみよう](https://aws.amazon.com/jp/blogs/news/kiroweeeeeeek-in-japan-day-5-kiro-for-shell-scripting/)

* Spec mode で自然言語で要求を伝え仕様書を生成してもらう
* 設計書の作成
* Kiro によるコードの自動生成


[イベントストーミングから要件・設計・タスクへ。Kiro を活用した仕様駆動開発](https://aws.amazon.com/jp/blogs/news/eventstorming-with-kiro/)


[Kiro における負債にならない Spec ファイルの扱い方](https://aws.amazon.com/jp/blogs/news/kiroweeeeeeek-in-japan-day-8-spec-bps/)

* Spec ファイルの構成
  * requirements.md：EARS (Easy Approach to Requirements Syntax) 記法での要件
  * design.md：構造やデータフローなど実装方針
  * tasks.md：要件と設計を基に Kiro が生成する実装タスク群
* Spec ファイルの分け方
  * 巨大な一つの Spec ファイルを作成するのではなく、複数の Spec ファイルを機能ごとに作成することを推奨
  * Kiro と壁打ちしながら Spec ファイルの分け方の案を出してもらいつつ分類していくのも手
* 要件変更時の対応
  * 要件の更新：requirements.md ファイルを直接修正するか、Spec モードを開始して Kiro に新しい要件や設計要素を追加するよう指示
  * 設計の更新: design.md ファイルに移動し「Refine」を選択
  * タスクの更新: tasks.mdファイルに移動し、「Update tasks」を選択
* Vibe モードで構築したコンテキストから Spec ファイルを作成することも可能
* Spec ファイルもバージョン管理推奨


[AWS Japan の新卒が Kiro でマネコン上の操作を支援する Chrome 拡張機能をチーム開発してみた！](https://aws.amazon.com/jp/blogs/news/kiroweeeeeeek-in-japan-day-9-team-development/)

* ソリューションの開発は [PRFAQ](https://aws.amazon.com/jp/blogs/news/propelling-innovation-the-people-culture-and-process-imperatives/) を書くところから
* ステアリング機能によりコーディング規約を定義
  * ステアリングファイルに AI からの質問やヒアリングを促す指示を含めると効果的
  * LLM は古いバージョンのフレームワークやライブラリに基づくコードを生成することがある。そのため、バージョン情報の明示が重要
  * ステアリングを育てていく継続的な改善が不可欠
* MVP (Minimum Viable Product) 思考で Spec 自体を小さく保つことが重要重要


[スピードと品質の両立 – Kiro が加速する開発、GitLab AI が支えるレビュー。新時代の開発パートナーシップ設計](https://aws.amazon.com/jp/blogs/news/kiroweeeeeeek-in-japan-day-10-gitlab-partnerships/)

* Kiro 側で Spec の仕様書やコード実装を管理
* GitLab + Amazon Bedrock でコードレビュー


[エスツーアイ株式会社様の AWS 生成 AI 事例「Kiro を活用した経費精算システムの迅速な開発」のご紹介](https://aws.amazon.com/jp/blogs/news/genai-case-study-s2i/)


# DevOps Agent

[AWS DevOps Agent はインシデント対応の迅速化とシステム信頼性の向上に役立ちます (プレビュー)](https://aws.amazon.com/jp/blogs/news/aws-devops-agent-helps-you-accelerate-incident-response-and-improve-system-reliability-preview/)

* 問題発生時に、メトリクスやログから GitHub や GitLab での最近のコードデプロイまで、運用ツールチェーン全体のデータを自動的に関連付け、考えられる根本原因を特定し、的を絞った緩和策を推奨
* Slack チャンネルを使ってステークホルダーに最新情報を伝えたり、詳細な調査スケジュールを管理できる
* Amazon CloudWatch、Datadog、Dynatrace、New Relic、Splunk などの一般的なサービスと連携してオブザーバビリティデータを取得し、GitHub Actions や GitLab CI/CD と統合してデプロイとそのクラウドリソースへの影響を追跡できる


# Q Developer

[Amazon Q Developer CLIにAWSアカウントを調査・操作させてみた](https://dev.classmethod.jp/articles/exploring-aws-with-q-cli/)

* 新規スクリプトの作成
  * CloudFormation スタックから drawio 形式で AWS 構成図を書いてくれる
  * アクセスログの分析ができる
  * プロンプトにより CloudFormation テンプレートの生成
* 既存スクリプトの解析
  * 仕様を解析してテキストに起こしてくれる
  * スクリプトの問題点も洗い出してくれる


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


