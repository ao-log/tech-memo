
# Kiro CLI

[Kiro CLI](https://kiro.dev/docs/cli/)

## Chat

[Chat](https://kiro.dev/docs/cli/chat/)

* `/model` でモデルを選択可能

[Subagents](https://kiro.dev/docs/cli/chat/subagents/)

* サブエージェントは独立して並列に動作させることができ、独自のコンテキストを持つ
* リアルタイムに進捗をトラッキングできる
* サブエージェントの呼び出し時にカスタムエージェントを指定できる。カスタムエージェントを指定しない場合、サブエージェントとしてデフォルトサブエージェントが使用される
* サブエージェントでは read, write, shell ツールおよび MCP サーバを使用できるが、その他のツールは使用できない

[Plan agent](https://kiro.dev/docs/cli/chat/planning-agent/)

* `/plan` で開始する。作成したいアプリ内容をプロンプトで依頼すると、対話型で色々と尋ねてくれて、実装計画を立ててくれる。plan は read-only で動作する

[Manage prompts](https://kiro.dev/docs/cli/chat/manage-prompts/)

* `/prompt` によって、プロンプトの管理ができる
  * `/prompts create --name name [--content content]` でプロンプトの作成
* `@code-review` のように作成したプロンプトを呼び出す
* Local prompts(`project/.kiro/prompts/`), Global prompts(`~/.kiro/prompts/`), MCP prompts の 3 種がある。 

[Context management](https://kiro.dev/docs/cli/chat/context/)

* コンテキスト管理方針
  * コンテキストの消費が大きい内容は Knowledge base の使用を検討する
  * 各会話ごとにコンテキストとして覚えさせる必要がない場合は、セッションコンテキストの使用を検討する
* コンテキストウィンドウの動作
  * コンテキストファイルはモデルのコンテキストウィンドウの 75 % まで。それを超えると自動的に drop する
  * Knowledge base はコンテキストウィンドウを消費しない
* 常に読み込ませたい内容はエージェントの設定ファイルの resources で指定する
```json
{
  "name": "my-agent",
  "description": "My development agent",
  "resources": [
    "file://README.md",
    "file://docs/**/*.md",
    "file://src/config.py"
  ]
}
```
* Knowledge base
  * `kiro-cli settings chat.enableKnowledge true` で有効化
  * `/knowledge add /path/to/large-codebase --include "/*.py" --exclude "node_modules/"` で追加
* `/context`
  * `/context add README.md` で指定ファイルをコンテキストに追加できる。`docs/*.md` のような指定も可
  * `/context show` でコンテキストの使用状況を表示
  * `/context remove src/temp-file.py` で指定ファイルをコンテキストから削除
  * `/context clear` でコンテキストを完全にクリアする
* `/compact` により古い会話を要約し、コンテキストを圧縮する
  * `/chat resume` でオリジナル内容から再開できる
* ベストプラクティス
  * 大きなファイルは token 消費が激しい
  * コンテキストファイルが大きい場合は、用途ごとに分割したり、Knowledge base にすることを検討する

[Responding to messages](https://kiro.dev/docs/cli/chat/responding/)

* `/reply` により、前の会話の特定部分を引用しつつ、当該引用箇所に対する返信や追加の指示ができる

[Managing tool permissions](https://kiro.dev/docs/cli/chat/permissions/)

* `/tools trust read`, `/tools untrust shell` のようにツールごとに実行時の trust 確認を設定できる

[Working with Git](https://kiro.dev/docs/cli/chat/git-aware-selection/)

* `git add` されたファイルの変更などをプロンプトで伝えなくても、自動的に認識してくれるようになっている。また `.gitignore` の内容を尊重する

[Working with images](https://kiro.dev/docs/cli/chat/images/)

* 画像ファイル内容の説明を行わせることが可能

[Security considerations](https://kiro.dev/docs/cli/chat/security/)

* 次のようなリスクがある
  * 破壊的な操作が実行される
  * セキュリティ的に問題のあるコードが生成される
* 推奨される方法
  * trust にしない
  * 開発環境で使用する
  * 機密ファイルはプロジェクト外で管理
  * ステアリングなどで制限する

[Configuration](https://kiro.dev/docs/cli/chat/configuration/)

* Global Scope の設定ファイルの置き場所は `<user-home>/.kiro/`
* Project Scope の設定ファイルの置き場所は `<project-root>/.kiro`
* 各設定ファイル (Global Scope の場合)
  * MCP servers: ~/.kiro/settings/mcp.json
  * Prompts: ~/.kiro/prompts
  * Custom agents: ~/.kiro/agents
  * Steering: ~/.kiro/steering
  * Settings: ~/.kiro/settings/cli.json
* 同じファイルがある場合 Agent > Project > Global の優先度順

[Custom Diff Tools](https://kiro.dev/docs/cli/chat/diff-tools/)


## Custom agents

[Custom agents](https://kiro.dev/docs/cli/custom-agents/)

[Creating custom agents](https://kiro.dev/docs/cli/custom-agents/creating/)

[Agent configuration reference](https://kiro.dev/docs/cli/custom-agents/configuration-reference/)

[Agent Examples](https://kiro.dev/docs/cli/custom-agents/examples/)

[Troubleshooting custom agents](https://kiro.dev/docs/cli/custom-agents/troubleshooting/)


## MCP

[Model Context Protocol (MCP)](https://kiro.dev/docs/cli/mcp/)

[Configuration](https://kiro.dev/docs/cli/mcp/configuration/)

[Examples](https://kiro.dev/docs/cli/mcp/examples/)

[Security](https://kiro.dev/docs/cli/mcp/security/)

[Governance](https://kiro.dev/docs/cli/mcp/governance/)


## Steering

[Steering](https://kiro.dev/docs/cli/steering/)


## Experimental

[Experimental features](https://kiro.dev/docs/cli/experimental/)


## Hooks

[Hooks](https://kiro.dev/docs/cli/hooks/)


## Autocomplete

[Completions & autocomplete](https://kiro.dev/docs/cli/autocomplete/)


## Code Intelligence

[Code Intelligence](https://kiro.dev/docs/cli/code-intelligence/)


## Billing

* [Billing for individuals](https://kiro.dev/docs/cli/billing/)
* [Enterprise billing](https://kiro.dev/docs/cli/enterprise/getting-started/)


## Privacy and security

[Privacy and security](https://kiro.dev/docs/cli/privacy-and-security/)


## Reference

[CLI commands](https://kiro.dev/docs/cli/reference/cli-commands/)

[Slash commands](https://kiro.dev/docs/cli/reference/slash-commands/)

[Built-in tools](https://kiro.dev/docs/cli/reference/built-in-tools/)

[Settings](https://kiro.dev/docs/cli/reference/settings/)


## Migrating

[Upgrading from Amazon Q Developer CLI](https://kiro.dev/docs/cli/migrating-from-q/)

