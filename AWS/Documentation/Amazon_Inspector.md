# Document

[Amazon Inspector とは](https://docs.aws.amazon.com/ja_jp/inspector/latest/user/what-is-inspector.html)

EC2 インスタンスだけでなく、コンテナイメージもスキャン可能


## リソースのスキャン

[Amazon Inspector を使用した Amazon EC2 インスタンスのスキャン](https://docs.aws.amazon.com/ja_jp/inspector/latest/user/enable-disable-scanning-ec2.html)


[Amazon インスペクターで Amazon ECR コンテナイメージをスキャンする](https://docs.aws.amazon.com/ja_jp/inspector/latest/user/enable-disable-scanning-ecr.html)

* OS パッケージ、プログラミンング言語のパッケージをスキャン可能
* スキャンのタイミングは継続(プッシュ時 + 自動)、プッシュ時から選択可能。新しい CVE アイテムが追加されたタイミングで自動スキャンが実行される



# Blog

[Amazon Inspector と AWS Systems Manager を用いた脆弱性管理と修復の自動化 – Part1](https://aws.amazon.com/jp/blogs/news/automate-vulnerability-management-and-remediation-in-aws-using-amazon-inspector-and-aws-systems-manager-part-1/)

* Amazon Inspector は、Systems Manager (SSM) エージェント を使用して Amazon EC2 インスタンスのソフトウェアアプリケーションインベントリを収集
* AWS Systems Manager Patch Manager を使用すると、SSM エージェントを使用して Systems Manager が管理するノードにパッチを適用するプロセスを自動化できる
* ゼロデイ攻撃に対しては定期実行ではなくオンデマンド実行による対応が必要
* AWS Security Hub カスタムアクションを使用して、選択した EC2 インスタンスに対するオンデマンドの脆弱性パッチ適用のための Systems Manager Automation runbook をトリガー


[Amazon Inspector と AWS Systems Manager を用いて脆弱性の修復パイプラインを構築しよう](https://aws.amazon.com/jp/builders-flash/202210/create-vulnerability-pipeline/?awsf.filter-name=*all)



# BlackBelt

[Amazon Inspector](https://pages.awscloud.com/rs/112-TZM-766/images/AWS-Black-Belt_2023_AmazonInspector_0228_v1.pdf)

* パッケージの脆弱性やネットワーク到達性を継続的なスキャンで検出
* EC2, ECR, Lambda に対応
* パッケージ脆弱性
  * EC2
    * Systems Manager で管理されているマネージドインスタンスをスキャン
    * 新規インスタンスの起動時、ソフトウェアインストール時、Inspector で新たな脆弱性をデータベースに追加時などにスキャン
  * ECR
    * プッシュ時もしくは連続スキャン(CVE 情報がデータベースに追加されるたびにスキャン)
    * 連続スキャンでは Lifetime, 30 日, 180 日から選択可能
    * OS, プログラミング言語のパッケージの両方に対応
  * Lambda
    * Lambda 関数内、Lambda Layer で使用されるアプリケーションパッケージをスキャン
    * 新しい Lambda 関数のデプロイ時、既存 Lambda 関数の更新時、Inspector で新たな脆弱性をデータベースに追加時などにスキャン
* ネットワーク到達性
  * EC2 に対応
  * SSH が公開されていないかなどチェック
* 検出結果
  * active, suppressed, closed の 3 種類のステータス
* ワークフロー
  * Security Hub と統合されている。Security Hub → EventBridge の連携も可能
  * Inspector から EventBridge の通知も可能。スキャン完了時、脆弱性検知時などにイベントが発行される
  * S3 にレポートを保管可能
* 運用
  * Systems Manager による継続的なパッチ適用も重要


[Amazon Inspector (2016)](https://pages.awscloud.com/rs/112-TZM-766/images/20160622_AWS_BlackBelt-Inspector-public.pdf)

* EC2 にエージェントを導入し脆弱性を診断する



# 参考

* Document
  * [Amazon Inspector とは](https://docs.aws.amazon.com/ja_jp/inspector/latest/user/what-is-inspector.html)
* Black Belt
  * [Amazon Inspector](https://pages.awscloud.com/rs/112-TZM-766/images/AWS-Black-Belt_2023_AmazonInspector_0228_v1.pdf)
  * [Amazon Inspector (2016)](https://pages.awscloud.com/rs/112-TZM-766/images/20160622_AWS_BlackBelt-Inspector-public.pdf)


