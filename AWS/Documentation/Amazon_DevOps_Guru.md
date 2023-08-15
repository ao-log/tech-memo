# Document

[ユーザーガイド](https://docs.aws.amazon.com/ja_jp/devops-guru/latest/userguide/welcome.html)

* 機械学習モデルを使用し、運用データ、アプリケーションメトリクス、ログから異常を検知できる


[High Level Workflow](https://docs.aws.amazon.com/ja_jp/devops-guru/latest/userguide/high-level-workflow.html)

概ね以下の流れで利用する

1. 分析対象の決定。アカウント単位、CloudFormation スタック、タグなどで指定
2. DevOps Guru による分析
3. イベント通知


[DevOps Guru の概念](https://docs.aws.amazon.com/ja_jp/devops-guru/latest/userguide/concepts.html)

* アノマリー: 異常な振る舞い、メトリクス
* インサイト: 分析した結果得られる異常に関する情報。事後対応型、事前対応型の 2 タイプ
* レコメンデーション: レコメンデーションも提供される


[コスト見積り](https://docs.aws.amazon.com/ja_jp/devops-guru/latest/userguide/cost-estimate.html)

* DevOps Guru が分析するのに必要な月額コストの見積もりを行うことができる
* 分析対象のリソースが 1 か月稼働し続けていることを前提としている

料金補足
* [Amazon DevOps Guru の料金](https://aws.amazon.com/jp/devops-guru/pricing/)


[タグを使用して DevOps Guru アプリケーション内のリソースを識別する](https://docs.aws.amazon.com/ja_jp/devops-guru/latest/userguide/working-with-resource-tags.html)

* DevOps Guru の分析対象を指定するのにタグを使用することができる。どのリソースがモニタリング用にグループ化されるかをタグキー、タグの値で指定できる
* 「devops-guru-」で始まるタグキー名とする必要がある
* 分析時は一つのキーのみ指定可能
  * All account resources: 全リソースを分析する。選択したタグキーについてはタグの値によってグループ化される
  * Choose specific tag values: 指定したタグのリソースを分析する。
* インサイトにてタグによるフィルターが可能


# 参考

* Document
  * [ユーザーガイド](https://docs.aws.amazon.com/ja_jp/devops-guru/latest/userguide/welcome.html)
  * [Amazon DevOps Guru のよくある質問](https://aws.amazon.com/jp/devops-guru/faqs/)


