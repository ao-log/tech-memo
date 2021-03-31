
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



