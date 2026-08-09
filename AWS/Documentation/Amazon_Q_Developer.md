
# Document

[Amazon Q Developer とは?](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/what-is.html)

* Bedrock が基盤となっている
* 使用方法
  * マネジメントコンソール上の Amazon Q
  * IDE の拡張機能


## 使用開始

[Amazon Q Developer の使用開始](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/getting-started-q-dev.html)


[Q Developer のサービスティア – Free および Pro](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/q-tiers.html)

* Free ティア
  * サインイン方法
    * 個人アカウント (Builder ID) を持つユーザー: IDE, コマンドラインを使用可能
    * IAM Identity Center ID を持つユーザー: マネジメントコンソールから使用可能
    * IAM 認証情報を持つユーザー: マネジメントコンソールから使用可能
* Pro ティア
  * Free ティアよりも制限が緩めになっている。有料バージョン
  * サインイン方法
    * 個人アカウント (Builder ID) を持つユーザー: IDE, コマンドラインを使用可能
    * IAM Identity Center ID を持つユーザー: マネジメントコンソールから使用可能


[個人アカウント (Builder ID) を使用した開始方法](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/getting-started-builderid.html)

* Builder ID でサインインできる。この ID では AWS マネジメントコンソールの使用や IAM 権限の割り当てはできない
* Pro ティアでは制限が緩くなるが、全ての機能を使用できるわけではない。該当機能を使用したい場合は IAM アイデンティティセンター ID でサインインして使用する必要がある


[IAM アイデンティティセンターを使用した開始方法](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/getting-started-idc.html)

[ステップ 1: デプロイオプションを選択する](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/deployment-options.html)

* どこで IAM Identity Center を有効化するか
  * スタンドアロンアカウントにデプロイ
  * Organizations の管理アカウント
  * Organizations のメンバーアカウント

[ステップ 2: Amazon Q Developer Pro にワークフォースをサブスクライブする](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/subscribe-users.html)

* Amazon Q Developer プロファイルを作成
  * サブスクリプションが作成されるが、保留中の状態となる
  * サブスクリプションの画面で [ユーザーとグループの割り当て] を行うと当該ユーザーにメールが送信される。


[Amazon Q Developer Pro サブスクリプション](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/q-admin-setup-subscribe-general.html)

[Amazon Q Developer Pro のリージョンサポート](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/q-admin-setup-subscribe-regions.html)

* Builder ID は us-east-1 でサポート

[Amazon Q Developer Pro サブスクリプションの請求](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/tracking-across-org-cost-usage.html)

* IAM Identity Center かつ Organizations 使用時は組織に対して請求される

[Amazon Q Developer サブスクリプションのステータス](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/q-admin-setup-subscribe-status.html)

* ユーザーのステータスはアクティブ、もしくはキャンセル済み。IAM Identity Center の場合は更に保留中がある

[Amazon Q Developer で使用する開始 URL の表示](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/manage-account-details.html)

* サインインに使用する URL を表示できる

[Amazon Q Developer での暗号化方法の管理](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/manage-encryption.html)

* デフォルトではマネージドキーによる暗号だが、CMK による暗号化も設定可能

[Amazon Q Developer でのプロファイル共有の有効化](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/q-admin-profile-sharing.html)

* プロファイル共有をしておくと、メンバーアカウント使用時も Amazon Q Developer Pro サブスクリプションを使用できるようになる。Amazon Q 使用時に影響し、Free ティアの制限がかかっていない状態となる

[Amazon Q Developer Pro サブスクリプションのトラブルシューティング](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/q-admin-setup-subscribe-troubleshooting.html)

[Amazon Q Developer Pro のサブスクリプション解除](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/q-admin-setup-unsubscribe.html)

[Amazon Q Developer Pro へのアップグレード](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/upgrade-to-pro.html)

[Kiro へのアップグレード](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/upgrade-to-kiro.html)


## On AWS

[AWS アプリやウェブサイトでの Amazon Q Developer の使用](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/q-on-aws.html)

* マネジメントコンソール上で Amazon Q に対し、チャット形式で質問できる
* Pro ティアで使用するには IAM Identity Center でログインする必要がある

[リソースに関する Amazon Q Developer とのチャット](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/chat-actions.html)

* アカウント内の AWS リソースについて一覧を出力するなどの質問をできる

[Amazon Q にリソースのトラブルシューティングを依頼する](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/chat-actions-troubleshooting.html)

* AWS リソースに関するトラブルシューティングが可能。対応可能な AWS サービス及び解決できる問題の種類は当ドキュメントの表にまとめられている

[コストに関するチャット](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/chat-costs.html)

* コストに関する質問が可能。プロンプト例も当ドキュメントにまとめられている

[ネットワークセキュリティに関するチャット](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/chat-network-security.html)

* セキュリティリスクを分析する質問が可能。プロンプト例も当ドキュメントにまとめられている

[E メール送信に関するチャット](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/chat-email.html)

[テレメトリと運用に関するチャット](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/chat-ops.html)

* 指定された AWS サービスのリソースの状態を評価し、これらのリソースで発生した問題やエラーのトラブルシューティングと解決を支援
* 「アラーム」状態のアラームと、アラームをトリガーした基になるテレメトリを特定し、アラーム/アラート/ページの背後にある原因を診断


[Amazon Q Developer Console-to-Code による AWS サービスの自動化](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/console-to-code.html)

* コンソールで操作した内容と同等の CloudFormation テンプレートや CDK コードを生成できる
* Console-to-Code サイドパネルで、[記録を開始] してから[停止] をクリックまでの操作が対象となる


[Amazon Q Developer でコンソールでの一般的なエラーを診断する](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/diagnose-console-errors.html)

* サービスによってはエラー内容の診断ができる。[Amazon Q で診断] をクリックすることで診断できる


[Amazon Q Developer を使用して サポートとチャットする](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/support-chat.html)

* Amazon Q のチャットパネルからサポートケースを起票可能


## IDE

[IDE での Amazon Q Developer の使用](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/q-in-IDE.html)

[IDE に Amazon Q Developer 拡張機能またはプラグインをインストールする](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/q-in-IDE-setup.html)

* IDE の拡張機能、プラグインにより使用可能となる
* Builder ID もしくは IAM Identity Center で認証する


[コードに関する Amazon Q Developer とのチャット](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/q-in-IDE-chat.html)

* IDE 上でチャットを開ける
* コーディングエージェントはデフォルトで有効化されている
* 実行できるタスク例
  * コード生成、修正。コード内容の説明
  * ユニットテスト生成
  * コードレビュー、脆弱性確認
  * コード変換(アップグレードなど)

[Amazon Q Developer でのコードレビュー](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/code-reviews.html)

* ファイルまたはプロジェクト全体をレビュー可能
* デフォルトでは `git diff` のコード差分が対象。差分がない場合はコード全体が対象となる
* 次の種類の問題をレビューする
  * 脆弱性があるかどうか
  * シークレットが含まれていないかどうか
  * IaC コードに設定ミスやコンプライアンス、セキュリティの問題がないかどうか
  * 品質に関する問題が含まれていないかどうか
  * パフォーマンスなどのリスクがあるかどうか
  * ライブラリなどが安全かつ最新であるかどうか

[Amazon Q Developer でコードレビューを開始する](http://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/start-review.html)

* チャットパネルから自然言語でレビューを指示できる
* Few-shot prompting
  * プロンプト内で例を示すことで、例と同じように回答するように誘導できる

[Amazon Q Developer を使用した IDE でのコード変換](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/transform-in-IDE.html)

* JAVA, .NET のアップグレードなどの変換が可能

[Amazon Q Developer でのコードの説明と更新](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/explain-update-code.html)

* 特定のコード行について、説明、リファクタリング、デバッグ、テストの生成、パフォーマンス最適化、プロンプト送信などの操作が可能


[IDE の Amazon Q Developer チャットへのコンテキストの追加](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/ide-chat-context.html)

* チャットウィンドウに @ と入力すると、@workspace などのコンテキストタイプを選択できる
* コンテキストタイプは workspace, フォルダ、ファイル、コード、画像、プロンプトがある
* プロジェクトルールはコンテキストとして自動的に使用される

[IDE の Amazon Q Developer チャットへのワークスペースコンテキストの追加](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/workspace-context.html)

* ワークスペースをコンテキストとして使用するためには、インデックスが作成されている必要がある
* ワークスペースインデックス作成の有効化を行なっておく必要がある
* `@workspace 認可を処理するコードはどこにありますか?` のようなプロンプトで質問する

[Amazon Q Developer チャットで使用するプロンプトをライブラリに保存する](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/context-prompt-library.html)

* `~/.aws/amazonq/prompts` フォルダにプロンプトの md ファイルを配置しておく
* `@Create_sequence_diagram using the files in the @lib folder` のようなプロンプトで質問する

[Amazon Q Developer チャットで使用するプロジェクトルールの作成](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/context-project-rules.html)

* `project-root/.amazonq/rules` に配置する
* 自動的にコンテキストとして使用される

[Amazon Q チャット用のメモリバンクの生成](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/context-memory-bank.html)

* メモリバンクの生成を行うことで、プロジェクトを分析した内容を `.amazonq/rules` に保存する。以下ファイルが生成される
  * product.md – プロジェクトとその機能の概要
  * structure.md – プロジェクトのアーキテクチャ、フォルダの編成、主要コンポーネント
  * tech.md – テクノロジースタック、フレームワーク、依存関係、コーディング標準
  * guidelines.md – プロジェクトの開発標準とパターン

[Amazon Q Developer でのチャット履歴の圧縮](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/ide-chat-history-compaction.html)

* `/compact` により手動圧縮
* `/clear` によりチャット履歴を完全に消去
* コンテキストウィンドウが容量の約 80% に達すると、圧縮を推奨する通知が表示される。圧縮推奨理由の説明と圧縮をすぐ実施するためのボタンが表示される


[Amazon Q Developer によるインライン提案の生成](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/inline-suggestions.html)

* 以下のアシストを受けることができる
  * コード補完
  * コメント内容に応じたコード生成
  * 関数名の入力するとコード生成
  * ユニットテストコード生成


## CLI

Kiro CLI になった。


## MCP

[Amazon Q Developer での MCP の使用](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/qdev-mcp.html)



## Old

[Amazon Q Developer 機能](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/features.html)

* 移行、transform
  * コードの移行、VMware 環境の transform
* 分析
  * Quick Sights でデータの要約
* マネジメント、ガバナンス
  * AWS リソース情報の質問
  * Systems Manager と連携したインスタンス情報の質問
  * 運用調査。AWS 環境全体のリソース、イベント、アクティビティを調査および分析
  * コンソールエラーの診断
* コンピューティング
  * EC2 のインスタンスセレクターでのレコメンド
* データベース
  * 自然言語からのクエリの作成
* ネットワーク、CDN
  * Reachability Analyzer と連携したネットワーク到達性の診断
* Developer Tool
  * コード生成
  * インラインコード提案
  * コードに関するチャット
  * 脆弱性、品質の診断
  * ユニットテスト生成
  * Amazon SageMaker AI Studio のコードに関するチャット
* アプリケーション統合
  * Console-to-Code 機能により、操作内容を自動化
  * ETL スクリプトの記述とデータの統合
* クラウド財務管理
  * コストの分析
* Support
  * チャットで質問対応。サポートケース作成


[IAM Identity Center の開始方法](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/getting-started-idc.html)

* プロファイルにて、どの IAM Identity Center リソースを使用するかを設定
* サブスクリプションにて、個別グループ、ユーザーをアクティブにしていく


[Chatting with Amazon Q Developer about AWS](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/chat-with-q.html)

* マネジメントコンソール上でチャットにより以下内容の質問が可能
  * リソースの状態
  * トラブルシューティング
  * コスト
  * ネットワークセキュリティ(セキュリティ上問題のある設定がないかどうかの診断ができる)
  * メトリクス、アラームの分析
  * サポートケースの作成
* Console-to-Code アイコン
  * 一部のサービスではコンソール上の操作内容と同等の内容の AWS CLI コマンドを生成可能
* エラーメッセージに Amazon Q で診断する旨のボタンが表示される。一部サービスにて対応


[IDE での Amazon Q Developer の使用](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/q-in-IDE.html)

* VS Code の場合は以下設定を行う
  * Amazon Q 拡張機能のインストール
  * IAM Identity Center の場合は、開始 URL を設定し、認証を行う

* IDE 上でできる操作
  * チャット
    * 例
      * AWS のサービスの選択 Limits、ベストプラクティス
      * プログラミング言語の構文やアプリケーション開発を含む一般的なソフトウェア開発の概念
      * コードの説明、コードのデバッグ、ユニットテストを含むコードの記述
  * コード
    * 該当コードを強調表示
      * 説明
      * リファクタリング
      * 修正
      * テストの生成
      * 最適化
      * プロンプトへの送信
    * インライン提案
      * コメントを記述すると、コードを提案
      * 関数名を入力すると、try/except 句を含むコードを生成
      * ユニットテストクラス名から、コードを生成
    * `/dev`
      * 自然言語にて作成したいコードを命令すると、コードが自動生成される
    * `/test`
      * ユニットテストコードの生成
    * `/review`
      * コードベースでセキュリティの脆弱性やコード品質の問題を確認
    * `/doc`
      * ドキュメントの生成


[コマンドラインで Amazon Q Developer を使用する](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/command-line.html)

* q chat
  * アクセス許可設定
    * 次の例のように設定する
      * `/tools trust fs_read`
      * `/tools untrust execute_bash`
      * `/tools trustall`
  * モデルの指定
    * `q chat --model <model name>`


[セキュリティに関する考慮事項とベストプラクティス](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/command-line-chat-security.html)

* `/acceptall`、`/tools trustall` を設定していると、各操作の際に確認を求められないので、意図しない動作を実行してしまうリスクがある
* `/tools untrust fs_read` のように明示的なアクセス許可を求めるようにする
* 本番環境では trustall は使用しない


[Amazon Q Developer での MCP の使用](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/qdev-mcp.html)

* `/tools` により使用可能ツールを確認
* グローバルスコープの場合は `~/.aws/amazonq/mcp.json`、ローカルスコープの場合は `amazonq/mcp.json` に保存される
* MCP 設定ファイルの構造
```json
{
  "mcpServers": {
    "server-name": {
      "command": "command-to-run",
      "args": ["arg1", "arg2"],
      "env": {
        "ENV_VAR1": "value1",
        "ENV_VAR2": "value2"
      },
      "timeout": 60000
    }
  }
}
```

