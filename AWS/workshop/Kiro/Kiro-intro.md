
[Kiro ワークショップ: 仕様駆動開発](https://catalog.workshops.aws/kiro-intro/ja-JP)

基本的にはタスク管理アプリを開発していくワークショップになっている。


## セットアップと設定

[AWS イベントでのワークショップ実行](https://catalog.workshops.aws/kiro-intro/ja-JP/10-start-workshop/11-aws-event/12-configure-ide)

* AWS イベントからワークショップを実行する場合は、Ubuntu Desktop にリモートデスクトップログインし、Kiro IDE を起動して対応する

[個人セットアップ](https://catalog.workshops.aws/kiro-intro/ja-JP/10-start-workshop/15-on-your-own)
 
* Kiro IDE をインストール。Windows, macOS, Linux それぞれインストーラが用意されている
* チャット入力では `#file`、`#folder`、`#codebase` をコンテキストプロバイダーとして使用できる

[IDE の設定とパスの選択](https://catalog.workshops.aws/kiro-intro/ja-JP/10-start-workshop/15-on-your-own/16-configure-ide)

* コンテキストプロバイダーを試す
  * `#codebase このプロジェクトの全体的な構造はどうなっていますか？`
* ワークショップパス
  * 基本的には「タスク管理システム」を構築する
  * それとは別に既存プロジェクトに適応することも可能

[MCP: Kiro の機能を拡張](https://catalog.workshops.aws/kiro-intro/ja-JP/10-start-workshop/17-mcp)

* uv をインストールすると `uvx` コマンドも使用できるようになる
* チャットで MCP サーバのセットアップするように Kiro に依頼すると良い。`mcp.json` をよしなに作ってくれる
* 各 MCP サーバは uvx でインストールし、ローカル PC 上のプロセスとして起動する
* プロンプトで、どの MCP サーバを使用するか指示するのが効果的


## アプリ設計: 仕様と計画

[アプリ設計: 仕様と計画](https://catalog.workshops.aws/kiro-intro/ja-JP/30-app-design)

[タスク 1: Kiro での効果的プロンプトのマスター](https://catalog.workshops.aws/kiro-intro/ja-JP/30-app-design/30-effective-prompting)

* いいプロンプトとは
  * 特定のファイルを読み込むように指示し、スコープを狭める
  * 何を実装するかを具体的に指示する
  * 明確なスコープと制約を指示する
* プロンプトのパターン
  * 特定機能の実装
  * 既存機能の改善
  * 変更内容の意図や動作内容、トレードオフなどを説明してもらう

* いいプロンプトを書くために Kiro Chat で壁打ちするのも手

[タスク 2: 仕様駆動開発の理解](https://catalog.workshops.aws/kiro-intro/ja-JP/30-app-design/31-start-designing)

* AI アシストツールでよくある課題として、コンテキストが多くなると、以前の情報を忘れることがある
* 仕様駆動開発では `requirements.md`、`design.md`、`tasks.md` を作成する

[タスク 3: Kiro の統合開発環境の理解](https://catalog.workshops.aws/kiro-intro/ja-JP/30-app-design/32-change-model)

* Kiro パネルセクション
  * チャット欄
  * スペック
  * ステアリング
  * フック
* チャットウィンドウ
  * モデル選択
  * オートパイロットモード
* チャットモード
  * スペックモード
  * バイブモード

[タスク 4: ステアリングの理解](https://catalog.workshops.aws/kiro-intro/ja-JP/30-app-design/33-rules-setup)

* `.kiro/steering/` 内のマークダウンファイルを通じて Kiro にプロジェクトに関する永続的な知識を提供
* コード規約などの定義し、Kiro の動作を一貫したものにすることができる
* [ステアリング] セクションの [Generate Steering Docs] をクリックすると、「プロダクト概要 (`product.md`) 」「技術スタック (`tech.md`)」「プロジェクト構造 (`structure.md`) 」の 3 ファイルを生成できる
* `inclusion` によりステアリングは適用対象のファイルなどを定めることができる
  * `manual` の場合、チャットで `#steering-file-name` により呼び出せる

[タスク 5: プロジェクトスペックの作成](https://catalog.workshops.aws/kiro-intro/ja-JP/30-app-design/35-qdev-requirements)

* プロンプトでスペックを作っていく
  * Spec モードの状態で、チャット欄で作成したいアプリ内容をプロンプトで記載する。すると初期スペックが作成され、`requirements.md` が作成される
  * 次に設計フェーズを開発するように Kiro にプロンプトで指示。すると `design.md` が作成される
    * 技術アーキテクチャと実装アプローチが定義される
  * 最後に実装フェーズを作成するように指示。すると `task.md` が作成される。`task.md` を開くと、タスクの一覧ができているので、タスクの上側にある Start をクリックすると AI によりコードが自動生成されたり、ユニットテストが実行されていく。もしくはプロンプトでタスクの何番を実行するように指示することができる
* 上記フェーズ完了後に要件、設計、実装フェーズ間の一貫性を検証を行うのも手

ここまでの作業により、コーディングアシスタントとして Kiro を活用し実装に必要なファイルを作成した


## アプリ構築: 実装

[アプリ構築: 実装](https://catalog.workshops.aws/kiro-intro/ja-JP/50-app-build)

[タスク 1: プロジェクトのスキャフォールド](https://catalog.workshops.aws/kiro-intro/ja-JP/50-app-build/51-scaffold-scratch)

* 実装を指示。`tasks.md の最初のタスクを実装してください。` のようにチャットで指示する

[タスク 2: アプリケーションの構築](https://catalog.workshops.aws/kiro-intro/ja-JP/50-app-build/52-create-interface)

* ここで、フロントエンド、バックエンドを構築していく
* 実装してもらう際は、どのタスクを実行してもらうかや、どの情報を参照するべきかを指示するのが効果的
* 初期構造を実装してもらった後、改良のプロンプトにより修正していくことが可能
* 何故そのような実装をおこなったかなどを質問可能

[タスク 3: インフラストラクチャの生成](https://catalog.workshops.aws/kiro-intro/ja-JP/50-app-build/53-infra-as-q)

* ここで AWS リソースを構築していく
* 詳細な手順はないが CDK 用の Spec を新たに作る流れになるのではないかと思われる


## アプリテスト: MCP とテスト

[アプリテスト: MCP とテスト](https://catalog.workshops.aws/kiro-intro/ja-JP/70-app-test)

[タスク 1: AI アシスタンスによる包括的テスト](https://catalog.workshops.aws/kiro-intro/ja-JP/70-app-test/71-mcp-setup)

* ユニット、統合、E2E などのテストを実装できる

[タスク 2: 統合テストの構築](https://catalog.workshops.aws/kiro-intro/ja-JP/70-app-test/73-switch-profiles)

[タスク 3: 自動テストの実行](https://catalog.workshops.aws/kiro-intro/ja-JP/70-app-test/75-build-tests)

[タスク 4: ユニットテストの完了まで反復](https://catalog.workshops.aws/kiro-intro/ja-JP/70-app-test/77-complete-integration)

* テストと実装修正を繰り返すことで、完全なテストカバレッジに近づいていく


## 上級トピック

[上級トピック](https://catalog.workshops.aws/kiro-intro/ja-JP/80-app-extras)

[タスク 1: フックと自動化の理解](https://catalog.workshops.aws/kiro-intro/ja-JP/80-app-extras/82-simple-hooks)

[タスク 2: 仕様駆動開発の自動化](https://catalog.workshops.aws/kiro-intro/ja-JP/80-app-extras/81-amazonq-context)

[タスク 3: デプロイメントコストの評価](https://catalog.workshops.aws/kiro-intro/ja-JP/80-app-extras/83-evaluating-cost)


## Summary

[ワークショップ結論: 実用的な要点と次のステップ](https://catalog.workshops.aws/kiro-intro/ja-JP/99-summary)

主要なファイル構造
```
.kiro/
├── specs/[feature-name]/
│   ├── requirements.md
│   ├── design.md
│   └── tasks.md
├── steering/
│   └── development-standards.md
└── settings/
    └── mcp.json
```

