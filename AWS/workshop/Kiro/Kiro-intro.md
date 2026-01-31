
[Kiro ワークショップ: 仕様駆動開発](https://catalog.workshops.aws/kiro-intro/ja-JP)

* ワークショップでは Ubuntu Desktop にリモートデスクトップログインし、Kiro IDE を起動して対応する

[個人セットアップ](https://catalog.workshops.aws/kiro-intro/ja-JP/10-start-workshop/15-on-your-own)

* Kiro IDE をインストール。Windows, macOS, Linux に対応

[MCP: Kiro の機能を拡張](https://catalog.workshops.aws/kiro-intro/ja-JP/10-start-workshop/17-mcp)

* 各 MCP サーバは uvx でインストール
* プロンプトで、どの MCP サーバを使用するか指示するのが効果的


## アプリ設計: 仕様と計画

[アプリ設計: 仕様と計画](https://catalog.workshops.aws/kiro-intro/ja-JP/30-app-design)

[タスク 1: Kiro での効果的プロンプトのマスター](https://catalog.workshops.aws/kiro-intro/ja-JP/30-app-design/30-effective-prompting)

* いいプロンプトを書くために Kiro Chat で壁打ちするのも手

[タスク 2: 仕様駆動開発の理解](https://catalog.workshops.aws/kiro-intro/ja-JP/30-app-design/31-start-designing)

* AI アシストツールでよくある課題として、コンテキストが多くなると、以前の情報を忘れることがある

[タスク 3: Kiro の統合開発環境の理解](https://catalog.workshops.aws/kiro-intro/ja-JP/30-app-design/32-change-model)

[タスク 4: ステアリングの理解](https://catalog.workshops.aws/kiro-intro/ja-JP/30-app-design/33-rules-setup)

* `.kiro/steering/` 内のマークダウンファイルを通じて Kiro にプロジェクトに関する永続的な知識を提供
* コード規約などの定義し、Kiro の動作を一貫したものにすることができる
* ステアリングは適用対象のファイルなどを定めることができる

[タスク 5: プロジェクトスペックの作成](https://catalog.workshops.aws/kiro-intro/ja-JP/30-app-design/35-qdev-requirements)

* スペック機能により要件、設計考慮事項、実装タスクを含む包括的なスペックを作成
* 設計フェーズにて技術アーキテクチャと実装アプローチが定義される
* 実装計画フェーズにて実行可能なタスクに開発作業を分解
* 上記流れにより requirements.md、design.md、tasks.md が生成される
* 上記フェーズ完了後に要件、設計、実装フェーズ間の一貫性を検証を行うのも手

ここまでの作業により、コーディングアシスタントとして Kiro を活用し実装に必要なファイルを作成した


## アプリ構築: 実装

[アプリ構築: 実装](https://catalog.workshops.aws/kiro-intro/ja-JP/50-app-build)

[タスク 1: プロジェクトのスキャフォールド](https://catalog.workshops.aws/kiro-intro/ja-JP/50-app-build/51-scaffold-scratch)

[タスク 2: アプリケーションの構築](https://catalog.workshops.aws/kiro-intro/ja-JP/50-app-build/52-create-interface)

* 初期構造を実装してもらった後、改良のプロンプトにより修正していくことが可能
* 何故そのような実装をおこなったかなどを質問可能

[タスク 3: インフラストラクチャの生成](https://catalog.workshops.aws/kiro-intro/ja-JP/50-app-build/53-infra-as-q)


## アプリテスト: MCP とテスト

[アプリテスト: MCP とテスト](https://catalog.workshops.aws/kiro-intro/ja-JP/70-app-test)

[タスク 1: AI アシスタンスによる包括的テスト](https://catalog.workshops.aws/kiro-intro/ja-JP/70-app-test/71-mcp-setup)

[タスク 2: 統合テストの構築](https://catalog.workshops.aws/kiro-intro/ja-JP/70-app-test/73-switch-profiles)

[タスク 3: 自動テストの実行](https://catalog.workshops.aws/kiro-intro/ja-JP/70-app-test/75-build-tests)

[タスク 4: ユニットテストの完了まで反復](https://catalog.workshops.aws/kiro-intro/ja-JP/70-app-test/77-complete-integration)


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


