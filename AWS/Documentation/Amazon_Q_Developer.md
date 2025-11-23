
# Document

[Amazon Q Developer とは?](https://docs.aws.amazon.com/ja_jp/amazonq/latest/qdeveloper-ug/what-is.html)


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

