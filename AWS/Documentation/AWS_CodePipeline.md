
# CodePipeline

[AWS CodePipeline とは](https://docs.aws.amazon.com/ja_jp/codepipeline/latest/userguide/welcome.html)

[概念](https://docs.aws.amazon.com/ja_jp/codepipeline/latest/userguide/concepts.html)

* パイプライン：　一連のステージから構成されている。
* ステージ：　アクションから構成されている。アクションを並列に実行することも可能。
* アクション：　source、build、test、deploy、approval、および invoke のいずれか。
* トランジション：　ステージ間のポイントのこと。
* アーティファクト：　ソースコード、定義ファイルなど。前のステージの出力アーティファクトを次のステージの入力アーティファクトとして渡すことが可能。


[パイプライン実行の仕組み](https://docs.aws.amazon.com/ja_jp/codepipeline/latest/userguide/concepts-how-it-works.html)

ソースコードを変更したり、パイプラインを手動で起動したりするときに、実行をトリガー可能。
スケジュールした Amazon CloudWatch Events ルールを使用して実行をトリガーすることも可能。

* ルール 1: 実行の処理中はステージがロックされる
* ルール 2: 後続の実行はステージのロックが解除されるまで待機する
* ルール 3: 待機中の実行はより最近の実行によって置き換えられる



## チュートリアル

[チュートリアル: CodePipeline を使用した Amazon ECS 標準デプロイ](https://docs.aws.amazon.com/ja_jp/codepipeline/latest/userguide/ecs-cd-pipeline.html)

CodeCommit → CodeBuild(ECR へ push) → ECS へデプロイの流れ。


[チュートリアル: Amazon ECR ソースと、ECS と CodeDeploy 間のデプロイを含むパイプラインを作成する](https://docs.aws.amazon.com/ja_jp/codepipeline/latest/userguide/tutorials-ecs-ecr-codedeploy.html)

Source(CodeCommit, ECR) → ECS(Blue/Green) のデプロイの流れ。



## パイプラインの使用

[パイプラインの使用](https://docs.aws.amazon.com/ja_jp/codepipeline/latest/userguide/pipelines.html)

パイプラインの実行を開始する方法。

* ソースが変更された場合
* 手動
* スケジュール

ソースアクションとして選べるもの。

* S3
* Bitbucket
* CodeCommit
* ECR
* GitHub, GitHub Enterprise


## アクション

[アクションの使用方法](https://docs.aws.amazon.com/ja_jp/codepipeline/latest/userguide/actions.html)

ステージにはアクションを設定することができる。次のアクションを設定可能。

* Source
* Build
* Test
* Deploy
* Approval
* Invoke



## パイプライン構造リファレンス

[パイプライン構造リファレンス](https://docs.aws.amazon.com/ja_jp/codepipeline/latest/userguide/reference-pipeline-structure.html)

アクションのカテゴリごとにどのアクションプロバイダーを使用できるかの一覧を確認可能。

**パイプラインの構造の要件**

* パイプラインに少なくとも 2 つのステージを含める必要がある
* パイプラインの最初のステージには、少なくとも 1 つのソースアクションが含まれている必要がある
* ソースアクションは、パイプラインの最初のステージにのみ含める
* 各パイプラインのいずれかのステージには、必ずソースアクション以外のアクションを含める

**アクションの構造の要件**

* アクションの入力アーティファクトは、前述のアクションで宣言された出力アーティファクトと完全に一致する必要がある。



## アクション構造リファレンス

[アクション構造リファレンス](https://docs.aws.amazon.com/ja_jp/codepipeline/latest/userguide/action-reference.html)


[CodeCommit](https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodeCommit.html)

**Configuration Parameters**

* RepositoryName
* BranchName

**Output artifacts**

出力アーティファクトは ZIP ファイルとして生成される。


[AWS CodeBuild](https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodeBuild.html)

**Configuration parameters**

* ProjectName

**Input artifacts**

Primary ソースアーティファクト内の buildspec.yml を使用。

**Output artifacts**

buildspec.yml の artifacts セクションで設定したものが出力アーティファクトとなる。


[ECS, CodeDeploy Blue/Green](https://docs.aws.amazon.com/ja_jp/codepipeline/latest/userguide/action-reference-ECSbluegreen.html)

**Configuration Parameters**

* ApplicationName
* DeploymentGroupName
* TaskDefinitionTemplateArtifact
* AppSpecTemplateArtifact

**Input Artifacts**

imageDetail.jsonファイルで、イメージ URI をイメージにマッピング。
ECR ソースアクションを使用している場合は自動作成される。


[Amazon Elastic Container Service](https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-ECS.html)

**Configuration Parameters**

* ClusterName
* ServiceName

**Input Artifacts**

imagedefinitions.json が必要。



## BlackBelt

[20201111 AWS Black Belt Online Seminar AWS CodeStar & AWS CodePipeline](https://www2.slideshare.net/AmazonWebServicesJapan/20201111-aws-black-belt-online-seminar-aws-codestar-aws-codepipeline)

* P30: CodePipeline の各用語の定義
  * ステージ、アクション、入力/出力アーティファクト、トランジション
* P31: サービスロールを設定。アーティファクトストア、暗号化キーはオプション
* P39: イベントに関する通知を Amazon SNS で通知



# 参考

* Document
  * * [CodePipeline とは](https://docs.aws.amazon.com/ja_jp/codepipeline/latest/userguide/welcome.html)
* サービス紹介ページ
  * [AWS CodeDeploy](https://aws.amazon.com/jp/codepipeline/features/)
  * [よくある質問](https://aws.amazon.com/jp/codepipeline/faqs/)
* [ナレッジセンター](https://aws.amazon.com/jp/premiumsupport/knowledge-center/#AWS_CodePipeline)
* Black Belt
  * [20201111 AWS Black Belt Online Seminar AWS CodeStar & AWS CodePipeline](https://www2.slideshare.net/AmazonWebServicesJapan/20201111-aws-black-belt-online-seminar-aws-codestar-aws-codepipeline)

