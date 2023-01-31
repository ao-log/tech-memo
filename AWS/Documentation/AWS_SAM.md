
# AWS SAM

[チュートリアル: Hello World アプリケーションのデプロイ](https://docs.aws.amazon.com/ja_jp/serverless-application-model/latest/developerguide/serverless-getting-started-hello-world.html)

```
// 対話式で入力していくと、入力内容に従って各ファイルが配備される。
$ sam init

// 依存関係を解消し、.aws-sam/build 下にファイルが配置される。
$ sam build

// デプロイを行う
$ sam deploy --guided
```


# パイプライン

[第一回 コンテナ Lambda の”いろは”、AWS CLI でのデプロイに挑戦 !](https://aws.amazon.com/jp/builders-flash/202103/new-lambda-container-development/)

* Lambda 関数のデプロイは ZIP 関数に固めて行う。```aws lambda invoke``` によって実行できる。
```shell
$ aws lambda create-function \
    --function-name func1 \
    --runtime python3.8 \
    --handler app.handler \
    --zip-file fileb://./package.zip \
    --role ${ROLE_ARN} 
```
* Lambda には Docker イメージをデプロイすることもできる。
```Dockerfile
FROM public.ecr.aws/lambda/python:3.8
COPY app.py   ./
CMD ["app.handler"]  
```

```shell
aws lambda create-function \
     --function-name func1-container  \
     --package-type Image \
     --code ImageUri=${ACCOUNTID}.dkr.ecr.${REGION}.amazonaws.com/func1@${DIGEST} \
     --role ${ROLE_ARN}
```


[コンテナ Lambda を開発、まずは RIC と RIE を使ってみよう !](https://aws.amazon.com/jp/builders-flash/202104/new-lambda-container-development-2/)

* コンテナイメージには Runtime Interface Client (RIC) と Lambda Runtime Interface Emulator (RIE) が含まれている必要がある。[ランタイム API](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/runtimes-api.html) で応答できるようにする必要があるため。AWS  が提供する AWS Lambda のベースイメージには事前に含まれている。
* ローカル環境でテストすることができる。
```
$ curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{}'

{"statusCode": 200, "body": "{\"message\": \"hello world\"}"}
```


[コンテナ Lambda をカスタマイズして、自分好みの PHP イメージを作ろう !](https://aws.amazon.com/jp/builders-flash/202106/new-lambda-container-development-3/)

* AWS 提供の Lambda ベースイメージは AWS Lambda サービス側でチャンクごとにキャッシュされている。よって起動が早くなる効果が期待できる。
* Lambda でサポートされていない言語などを使用したい場合はカスタムイメージで対応できる。


[コンテナ Lambda を AWS SAM でデプロイしよう !](https://aws.amazon.com/jp/builders-flash/202107/new-lambda-container-development-4/?awsf.filter-name=*all)

* ```sam init``` により対話型で各設問に答えていきファイルを自動配備する。コンテナイメージを使用するので Image を選択する。
* SAM テンプレートのプロパティの PackageType を Image に設定している。
* ```sam build``` によりビルドを行う。
* ```sam deploy --guided``` によりデプロイを行う。


[コンテナ Lambda の CI/CD パイプラインの考え方](https://aws.amazon.com/jp/builders-flash/202109/new-lambda-container-development-5/)

* ```sam pipeline bootstrap``` によってパイプラインに必要なアーティファクトリソース (ECR リポジトリや S3 Bucket) と CloudFormation の実行ロールをプロビジョニングできる。
* ``` sam pipeline init``` によりパイプライン用のテンプレートを作成できる。


[コンテナ Lambda の CI/CD パイプラインを SAM Pipeline で作ろう !](https://aws.amazon.com/jp/builders-flash/202110/new-lambda-container-development-6/)

作成する構成
* feature ブランチの更新をトリガーとして Lambda 関数のデプロイを行う。

手順
* CodeCommit のリポジトリは別途作成しておく必要がある。
* ```sam init``` でテンプレート類を生成。以降は sam init で作成したディレクトリ下で作業する。
* ```sam pipeline bootstrap``` で各 AWS リソースを作成。以下のリソースが生成される。
  * S3 バケット
  * ECR リポジトリ
  * IAM リソース
  * ```.aws-sam``` ディレクトリ。```pipelineconfig.toml``` が ```sam pipeline init``` 時に参照される。
* ```sam pipeline init``` にてパイプラインの設定を行う。pipeline.yaml などのファイルが生成される。
* パイプラインを作成する。```codepipeline.yaml``` を指定している。
```shell
$ sam deploy --guided \
 --template codepipeline.yaml \
 --stack-name php-lambda-app-pipeline-stack \
 --capabilities=CAPABILITY_IAM \
 --parameter-overrides="FeatureGitBranch=feature"
```
* feature ブランチに git push するとパイプラインが実行される。パイプラインは以下のステージ構成で作成されている。
  * Source: CodeCommit
  * UpdatePipeline: CloudFormation
  * BuildAndDeployFeatureStack: CodeBuild
* パイプラインが完了すると Lambda 関数がデプロイされている。

仕組みの補足
* 最後のステージにて ```sam build```、```sam deploy``` が実行されている。```buildspec_feature.yml``` は以下のような内容。
```yaml
(snip)
  build:
    commands:
      - sam build --use-container --template ${SAM_TEMPLATE}
      - . ./assume-role.sh ${TESTING_PIPELINE_EXECUTION_ROLE} feature-deploy
      - sam deploy --stack-name $(echo ${FEATURE_BRANCH_NAME} | tr -cd '[a-zA-Z0-9-]')
                    --capabilities CAPABILITY_IAM
                    --region ${TESTING_REGION}
                    --s3-bucket ${TESTING_ARTIFACT_BUCKET}
                    --image-repository ${TESTING_IMAGE_REPOSITORY}
                    --no-fail-on-empty-changeset
                    --role-arn ${TESTING_CLOUDFORMATION_EXECUTION_ROLE}
```


## Black Belt

[20190814 AWS Black Belt Online Seminar AWS Serverless Application Model](https://pages.awscloud.com/rs/112-TZM-766/images/20190814_AWS-Blackbelt_SAM_rev.pdf)

* 簡潔な記述でサーバレスの構成をデプロイできる
* AWS::Serverless::Function の Events プロパティで API Gateway のパス、メソッドを指定可能。直感的に書ける。Events は他には S3、EventBridge など
* CodeUri で指定したディレクトリの内容を ZIP にし、S3 バケットにアップロードされる
* AutoPublishAlias を指定するとデプロイ時に新しい Lambda 関数バージョンを作成し、エイリアスを最新版に変更する
* より詳細に API Gateway の設定を行いたい場合は AWS::Serverless::Api を使用する
* SAM CLI
  * sam init: 雛形作成
  * sam build: .aws-sam/build 内にビルド



# 参考

* Document
  * [AWS Serverless Application Model (AWS SAM) とは](https://docs.aws.amazon.com/ja_jp/serverless-application-model/latest/developerguide/what-is-sam.htmll)
* Black Belt
  * [20190814 AWS Black Belt Online Seminar AWS Serverless Application Model](https://pages.awscloud.com/rs/112-TZM-766/images/20190814_AWS-Blackbelt_SAM_rev.pdf)
    * CloudFormation を拡張している。
    * AWS::Serverless::Function などのリソースタイプによって簡潔にテンプレートを書ける。
    * Event プロパティで Lambda 関数のトリガーを設定。
    * AutoPublishAlias によって新バージョン作成時にエイリアスも新バージョンに向ける。


