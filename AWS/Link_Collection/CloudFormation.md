
## AWS Blog

[AWS CloudFormation の新機能 — 障害ポイントからスタックオペレーションをすばやく再試行する](https://aws.amazon.com/jp/blogs/news/new-for-aws-cloudformation-quickly-retry-stack-operations-from-the-point-of-failure-2/)



## スライド

[CI/CDプロセスにCloudFormationを本気導入するために考えるべきこと](https://speakerdeck.com/hamadakoji/cdpurosesunicloudformationwoben-qi-dao-ru-surutamenikao-erubekikoto)

* CloudFormation における CI/CD
  * CI: アーティファクト作成前に異常なテンプレートを排除
    * CloudFormation Linter: プロパティの補完、リソースの依存関係を図示
    * cfn-python-lint: テンプレートの静的解析
    * CloudFormation Guard: セキュリティポリシー評価
  * CD:
    * 実装方法は二つ: CodePipeline or CodeBuild から AWS CLI で。
    * Rain がイケてそう？



## 記事

[CloudFormationのCLI実行ツール Rain がイケてそうなので紹介したい](https://dev.classmethod.jp/articles/aws-cloudformation-rain/)

* rain build AWS::S3::Bucket > S3.yml でサンプルテンプレートを作れる。
* rain deploy S3.yml。進捗も見れる。前回の作成に失敗したスタックがあっても自動的に削除してくれる。更新もこれでできる。
* rain rm S3。スタックの削除。


[【小ネタ】 AWS CloudFormationテンプレートでAWSアカウントごとにリソース作成有無を決定する](https://dev.classmethod.jp/articles/cfn-create-resources-depending-on-accounts/)

* Conditions, Mappings を使って実現


[[アップデート] CloudFormation StackSetsでスタックを複数リージョンに”同時デプロイ”が可能に！【時短テク】](https://dev.classmethod.jp/articles/deploy-concurrently-across-multiple-aws-regions-using-cfn-stacksets/)



[AWS CloudFormation と AWS SAM を使用したサーバーレスアプリケーションの開発とデプロイ](https://aws.amazon.com/jp/blogs/news/build-deploy-serverless-app-using-aws-serverless-application-model-and-aws-cloudformation/)




