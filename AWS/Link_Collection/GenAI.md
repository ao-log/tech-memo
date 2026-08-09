

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


[開発以外にも使える !? Bedrock Engineer の AI エージェントをカスタマイズしてみよう !](https://aws.amazon.com/jp/builders-flash/202602/customize-bedrock-engineer-ai-agents/)

* Bedrock Engineer でカスタムの AI Agent 作成ができる


[エンタープライズにおける AI エージェント: Amazon Bedrock AgentCore を活用したベストプラクティス](https://aws.amazon.com/jp/blogs/news/ai-agents-in-enterprises-best-practices-with-amazon-bedrock-agentcore/)

* この機能で何ができるかではなく、課題を出発点とする
* 初期の計画段階で明確にしておくべきこと
  * 「やるべきこと」「やるべきでないこと」の明確な定義と文書化
  * エージェントの回答のトーン
  * ツール、パラメータ、ナレッジベースの明確な定義
  * Ground truth データセット
* まず PoC で実際に運用した時の問題を見つけ出す
* オブザーバビリティは初日から組み込んでおく
* AgentCore Gateway は、ツールの所在に関係なく統一されたエントリポイントを提供
* 評価の自動化。Ground truth データセットに対する正答率やレイテンシー
* 一つのエージェントに詰め込まず、マルチエージェント。エージェント間のコンテキスト共有は AgentCore Memory で
* AgentCore Runtime により各セッションを専用のコンピューティングとメモリを備えた独立したマイクロ仮想マシン (microVM) 上で実行できる
* AgentCore Evaluations による評価の実行メカニズムの簡素化



# Kiro

[Kiro 導入ガイド：始める前に知っておくべきすべてのこと](https://aws.amazon.com/jp/blogs/news/kiroweeeeeeek-in-japan-day-1-implementation-guide/)

* Amazon Q Developer と Kiro は独立したサブスクリプション体系を持つ別製品
* 既に Amazon Q Developer Pro サブスクリプションを利用している場合は Kiro CLI と Kiro IDE が Kiro Pro プラン相当の範囲内で有効
* Q Developer CLI は Kiro CLI にアップデートされる
* AWS 請求に一本化したい場合は、AWS IAM Identity Center での認証が必須。GitHub などでも認証できるが、その場合の支払い方法はクレジットカード


[Amazon Q Developer の IDE プラグインから Kiro に乗り換える準備](https://aws.amazon.com/jp/blogs/news/kiroweeeeeeek-in-japan-day-2-q-dev-ide-to-kiro/)

* Rules 機能が進化した Steering 機能。開発標準を定義できるので、チーム開発時に有用
* コードレビュー機能はない。Amazon Inspector コードセキュリティによるコードレビュー機能、Amazon Q Developer の IDE プラグインなどの方法を検討する必要がある
* Spec 機能では、要件（Requirements）、設計（Design）、タスク（Tasks）の 3 フェーズを経て、構造化された開発プロセスを実現


[Kiro を組織で利用するためのセキュリティとガバナンス](https://aws.amazon.com/jp/blogs/news/kiroweeeeeeek-in-japan-day-3-security-governance/)

* IAM Identity Center を用いた認証が可能
* セキュリティ、プライバシー
  * 使用 Tier によりプロンプト、返答が保存されるリージョンが異なる
  * データ転送時、保管時は暗号化される
  * 使用 Tier により、利用情報が製品の改善に使用される場合がある。オプトアウト可能。Kiro for enterprise の場合は自動的にオプトアウト
* PrivateLink を介した接続が可能


[Kiroを使ったペアプログラミングのすすめ](https://aws.amazon.com/jp/blogs/news/kiroweeeeeeek-in-japan-day-4-pair-programming/)

* 議論した内容のホワイトボードを撮影し、画像を Kiro に読み込ませる。画像認識機能により文字起こし
* 生成された仕様書を二人で同じ画面を見ながらレビュー。一人は設計の妥当性、もう一人は仕様との適合性確認と役割分担。より深いレビューが可能になり、承認も適当にならなかった。疑問点、改善点も議論し、すぐに Kiro に修正依頼することで、仕様の精度が飛躍的に向上
* 英語で発表する必要があったので、Hook で英語ドキュメント作成を自動化
* Git 統合されているので、前回のレビューからの変更点なども、即座に確認できる


[インフラエンジニアのあなたも！Shell スクリプト開発で Kiro を使ってみよう](https://aws.amazon.com/jp/blogs/news/kiroweeeeeeek-in-japan-day-5-kiro-for-shell-scripting/)

* Spec mode
  * 仕様書作成
    * 自然言語で要求を伝え仕様書を生成してもらう
    * 仕様書は EARS 記法 (Easy Approach to Requirements Syntax) 
    * 仕様書の内容を変更したい場合は Kiro に依頼して直してもらう
  * 設計フェーズ
  * タスク実行時に生成されたコードの修正を行いたい場合は、まず設計を直す。その後タスクを再実行する
* 既存コードの内容を読み込ませて、説明させることもできる
* Kiro がインフラエンジニアに向いている理由
  * ドキュメント化の自動化ができる
  * ベストプラクティスを適用できる
  * 学習コストが低い


[Amazon Q Developer CLI から Kiro CLI へ : 知っておくべき変更点](https://aws.amazon.com/jp/blogs/news/kiroweeeeeeek-in-japan-day-6-amazon-q-developer-cli-to-kiro-cli/)

* 変更履歴は `kiro-cli version --changelog=all` で確認できる
* おすすめ機能
  * カスタムエージェントにより用途ごとにツールの許可などを設定できる


[イベントストーミングから要件・設計・タスクへ。Kiro を活用した仕様駆動開発](https://aws.amazon.com/jp/blogs/news/eventstorming-with-kiro/)

* イベントストーミングでは、色分けされた付箋と付箋間の遷移のルールに従って、業務フローを可視化
* イベントストーミングの画像を Kiro に添付し、要件作成を指示することで requirements.md を作成してくれる
  * 「不明な点があれば質問してください。」と質問することで、曖昧さを解消するための質問を返してくれる
  * requerements.md は用語集、8 つの要件(EARS 記法)、受け入れ基準が記述されている
* 設計フェーズでは design.md を作成。
  * アーキテクチャ、データモデルなどが設計書として記述される 
  * 正確性プロパティが自動生成される。それぞれのプロパティには、以下の情報が含まレル
    * プロパティの説明（任意の〜に対して、〜すべきである）
    * 検証対象となる要件番号
* 実装タスクでは task.md を作成。プロパティテストはオプション扱いのためグレーになっている


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
* MVP (Minimum Viable Product) 思考で Spec 自体を小さく保つことが重要。アプリケーション層とインフラ層の Spec を分離した


[スピードと品質の両立 – Kiro が加速する開発、GitLab AI が支えるレビュー。新時代の開発パートナーシップ設計](https://aws.amazon.com/jp/blogs/news/kiroweeeeeeek-in-japan-day-10-gitlab-partnerships/)

* Kiro 側で Spec の仕様書やコード実装を管理
* GitLab + Amazon Bedrock でコードレビュー


[進化し続ける Kiro の仕様駆動開発を一度立ち止まって整理する](https://aws.amazon.com/jp/builders-flash/202606/organize-kiro-sdd-evolution/)

* 仕様駆動開発についての、Spec の分割の仕方や、Feature Spec, Bugfix Spec、などの項目ごとの選択肢やどれを選択するかの考え方などが整理されている記事


[バグ修正と既存アプリの上に構築するための新しい Spec タイプ](https://aws.amazon.com/jp/blogs/news/specs-bugfix-and-design-first/)

* デザインファーストワークフロー、バグ修正 の 2 つの新しい Spec がリリースされた


[Kiro Powers であなたにパワーを授けよう](https://aws.amazon.com/jp/builders-flash/202602/enpower-you-with-kiro-powers/)

* MCP サーバの課題としてコンテキスト消費がある
* Powers はツールをオンデマンドで読み込む
* 必須要素は `POWER.md`。keywords に記述したキーワードを含むプロンプトで該当する POWER をアクティブ化する
* Kiro IDE で使用可能


[Kiro の ACP 対応により特化型 IDE にも AI を](https://aws.amazon.com/jp/blogs/news/kiro-adopts-acp/)

* Agent Client Protocol（ACP）は、コードエディタと AI コーディングエージェント間の通信を標準化するプロトコル
* コードエディタから Kiro CLI を呼び出せるようになる


[Kiro CLI 2.0: デザイン刷新、ヘッドレス CI/CD パイプライン、Windows サポート](https://aws.amazon.com/jp/blogs/news/cli-2-0/)

* Kiro CLI 2.0 がリリースされた
* ヘッドレスモードに対応した
* Windows サポート
* UX の改善


[Kiro にハーネスを付ける ~ 制御された全力疾走のすすめ](https://aws.amazon.com/jp/builders-flash/202605/add-harness-kiro/)

* コンテキストエンジニアリング: モデルが推論時に何を見ているか (RAG、ツール定義、Few-shot 例、会話履歴、状態管理) を体系的に設計する技術
* ハーネスエンジニアリング: 複数セッションにまたがる環境全体を最適化
  * 実装と評価を別々のエージェントが行う。GAN (敵対的生成ネットワーク) に着想を得ている


["伝わらない"を解消 ! 生成 AI で磨くAWS サポートケース起票スキル ~ 生成 AI 活用による問題解決の効率化 ~](https://aws.amazon.com/jp/builders-flash/202602/genai-support-case-issue-skills/)

* 「技術的なお問い合わせに関するガイドライン」に記載された内容に沿った形でサポートケースを起票するために、生成 AI を活用してサポートの起票文をセルフレビューする方法を紹介している記事
* 「技術的なお問い合わせに関するガイドライン」をベースにしつつ、効果的にセルフレビューできるように Amazon Quick Suite の Amazon Quick Research でプロンプトを改善
* 方法
  * 4 通りある
    * チャット内で都度プロンプトを記述する
    * カスタムエージェントの prompt にて、プロンプト内容を記載したファイルを指定する
    * MCP で support-case-reviewer を設定する
    * Amazon Quick Suite のカスタムチャットエージェントを利用


[エスツーアイ株式会社様の AWS 生成 AI 事例「Kiro を活用した経費精算システムの迅速な開発」のご紹介](https://aws.amazon.com/jp/blogs/news/genai-case-study-s2i/)


[aws-observability](https://github.com/kirodotdev/powers/tree/main/aws-observability)


[ログもアラームもトレースもセキュリティ監査も、「とりあえず Kiro に聞いてみ」でよかった話](https://zenn.dev/aws_japan/articles/kiro-aws-observability-power)



# DevOps Agent

[AWS DevOps Agent はインシデント対応の迅速化とシステム信頼性の向上に役立ちます (プレビュー)](https://aws.amazon.com/jp/blogs/news/aws-devops-agent-helps-you-accelerate-incident-response-and-improve-system-reliability-preview/)

* 問題発生時に、メトリクスやログから GitHub や GitLab での最近のコードデプロイまで、運用ツールチェーン全体のデータを自動的に関連付け、考えられる根本原因を特定し、的を絞った緩和策を推奨
* Slack チャンネルを使ってステークホルダーに最新情報を伝えたり、詳細な調査スケジュールを管理できる
* Amazon CloudWatch、Datadog、Dynatrace、New Relic、Splunk などの一般的なサービスと連携してオブザーバビリティデータを取得し、GitHub Actions や GitLab CI/CD と統合してデプロイとそのクラウドリソースへの影響を追跡できる


[AWS DevOps Agent を本番環境にデプロイするためのベストプラクティス](https://aws.amazon.com/jp/blogs/news/best-practices-for-deploying-aws-devops-agent-in-production/)


[AWS DevOps Agent によるエージェンティック AI を活用した自律的インシデント対応](https://aws.amazon.com/jp/blogs/news/leverage-agentic-ai-for-autonomous-incident-response-with-aws-devops-agent/)

* CloudWatch Alarm が発報すると、インシデントを自動診断し、最近のコードデプロイに起因する DynamoDB の書き込みスロットリングを特定。オンデマンドキャパシティへの切り替えまたはロールバックの提案を自律的に Slack に投稿
* コンテキスト
  * DevOps Agent はリソーストポロジーだけでなく、テレメトリ、デプロイタイムライン、インフラストラクチャおよびアプリケーションコードも把握
* コントロール
  * アクセス可能な範囲は IAM 権限で制御
  * すべての推論ステップとアクションは、エージェントが記録後に変更できない不変の監査ジャーナルに記録
* 利便性
  * Agent Space に設定が入っているので、使用するエンジニアの PC 内の構成依存がない
* コラボレーション
  * 自律的なチームメンバー的に動作する。インシデントがトリガーされると自動的に調査が開始する
* 継続的学習
  * クラウドインフラストラクチャ、テレメトリデータ、コードリポジトリをスキャンして、アプリケーショントポロジーを継続的に学習・更新
  * 過去の調査を分析してパターンを特定し、将来のトラブルシューティングワークフローを最適化
* コスト効率
  * エージェントがタスクに積極的に取り組んでいる時間に対してのみ課金


[AI エージェントをプロトタイプから製品へ: AWS DevOps Agent 開発で得た教訓](https://aws.amazon.com/jp/blogs/news/from-ai-agent-prototype-to-product-lessons-from-building-aws-devops-agent/)


[AWS DevOps Agent はどこまで障害の原因を特定できるのか？re:Invent 2025 の新機能を検証](https://zenn.dev/10q89s/articles/2bae74f7a6cbdc)


# Q Developer

[Amazon Q Developer CLIにAWSアカウントを調査・操作させてみた](https://dev.classmethod.jp/articles/exploring-aws-with-q-cli/)

* 新規スクリプトの作成
  * CloudFormation スタックから drawio 形式で AWS 構成図を書いてくれる
  * アクセスログの分析ができる
  * プロンプトにより CloudFormation テンプレートの生成
* 既存スクリプトの解析
  * 仕様を解析してテキストに起こしてくれる
  * スクリプトの問題点も洗い出してくれる


# AWS Security Agent

[AWS Security Agent のオンデマンドペネトレーションテストの一般提供を開始](https://aws.amazon.com/jp/blogs/news/aws-security-agent-on-demand-penetration-testing-now-generally-available/)


# Others

[生成系 AI アプリケーションでベクトルデータストアが果たす役割とは](https://aws.amazon.com/jp/blogs/news/the-role-of-vector-datastores-in-generative-ai-applications/)

* ベクトルデータストアは、ベクトルを大規模に保存し、問い合わせを行うためのシステムであり、効率的な最近傍クエリアルゴリズムと適切なインデックスにより、データ検索を改善


[Spec-Driven Presentation Maker — 伝えたいことを先に設計し、スライド構築は AI に任せる](https://aws.amazon.com/jp/blogs/news/spec-driven-presentation-maker-ja/)

* 何を伝えたいか、伝わるか、が本質的に重要
* 設計から開始する
  * 1. ブリーフィング — 聞き手は誰か、何を伝えたいか、聞き手にどうなってほしいかを定義します
  * 2. アウトライン — 1 スライド 1 メッセージの原則で、各スライドが答えるべき疑問と、その答えを定義します
  * 3. アートディレクション — 色使い、フォント、ビジュアルの方向性を定義します
* 既存 Power Point テーマをそのまま使える


[エンタープライズにおける Amazon Bedrock による生成 AI のオペレーティングモデル](https://aws.amazon.com/jp/blogs/news/generative-ai-operating-models-in-enterprise-organizations-with-amazon-bedrock/)


[アート引越センター株式会社様の AI 活用事例 「 AI 見積りアプリによる引越し見積もりの自動化の実現」のご紹介](https://aws.amazon.com/jp/blogs/news/i3design-art-ai-auto-estimation/)

* 室内を撮影すると、デバイス上で 3D モデルが作成される
* AI が 3D モデル内の家具、家電を検出し、自動的に見積もりが行われる
* 3D モデルは S3 に格納
* EventBridge にて推論処理実行。PointNeXt をベースとした学習済みモデル。Amazon SageMaker Serverless Inference によるサーバーレス環境で非同期推論を実行
* SNS → SQS → Step Functions にて DB に結果格納
* Web 画面は ELB + ECS Fargate


[[プレビュー] AWS Knowledge MCP Serverを使ってみた](https://dev.classmethod.jp/articles/aws-knowledge-mcp-server-available-preview/)


[【実機検証】AWS Observability Kiro Power を ECS Fargate 環境で試したら、AIが脆弱性スキャンと設定ミスを自動で見つけてくれた話](https://qiita.com/YShiba92/items/30dee66127dd5f4677bd)

* 障害対応系の MCP のほか、インシデント対応などのステアリングが同梱されている
* Kiro IDE 上からプロンプトで障害調査できる

