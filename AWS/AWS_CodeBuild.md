
# CodeBuild

## はじめに

[AWS CodeBuild とは何ですか?](https://docs.aws.amazon.com/ja_jp/codebuild/latest/userguide/welcome.html)

CodeBuild はマネージド型のビルドサービス。

実行する方法。

* マネジメントコンソール
* AWS CLI
* AWS SDK
* AWS CodePipeline(ビルドステージまたはテストステージ)

[AWS CodeBuild の概念](https://docs.aws.amazon.com/ja_jp/codebuild/latest/userguide/concepts.html)

ビルドを実行した時の流れ。

1. ビルドプロジェクトを実行。
2. ソースコードをビルド環境にダウンロード。buildspec に基づいてビルド。
3. ビルド出力がある場合、出力アーティファクトを S3 バケットにアップロード。ビルド通知を Amazon SNS トピックに送信するなども可能。また、ビルド実行中に CloudWatch Logs にビルドのログを送信。



## 開始方法

[コンソールを使用した AWS CodeBuild の開始方法](https://docs.aws.amazon.com/ja_jp/codebuild/latest/userguide/getting-started.html)

JAR ファイルをビルドするサンプル。

#### ステップ 1: 2 つの S3 バケットを作成する
入力、出力用にそれぞれ S3 バケットを作成

#### ステップ 2: ソースコードの作成

以下の構成でソースコードを作成。
```
(root directory name)
    |-- pom.xml
    `-- src
         |-- main
         |     `-- java
         |           `-- MessageUtil.java
         `-- test
               `-- java
                     `-- TestMessageUtil.java
```

#### ステップ 3: buildspec ファイルを作成する

buildspec.yml ファイルを作成し、ステップ 2 で作成したディレクトリのトップに配置。
```yaml
version: 0.2

phases:
  install:
    runtime-versions:
      java: corretto11
  pre_build:
    commands:
      - echo Nothing to do in the pre_build phase...
  build:
    commands:
      - echo Build started on `date`
      - mvn install
  post_build:
    commands:
      - echo Build completed on `date`
artifacts:
  files:
    - target/messageUtil-1.0.jar
```

#### ステップ 4: ソースコードと buildspec ファイルをアップロードする
ステップ 2、3 で作成したファイル類から ZIP ファイルを作成し、入力用の S3 バケットにアップロード。

#### ステップ 5: ビルドプロジェクトを作成する
ビルドプロジェクトを作成する。

* ソースプロバイダー: S3
* バケット名: 入力用の S3バケット
* オブジェクト: ステップ ４で作成した ZIP ファイル
* 環境イメージ: Managed image
* オペレーティングシステム: Amazon Linux 2
* ランタイム: Standard (標準)
* イメージ: aws/codebuild/amazonlinux2-x86_64-standard:3.0
* サービスロール: New service role (新しいサービスロール)
* Buildspec: Use a buildspec file
* アーティファクトのタイプ: Amazon S3
* バケット名: 出力用の S3 バケット

#### ステップ 6: ビルドを実行する
ビルドを開始する。

#### ステップ 7: 要約されたビルド情報を表示する
フェーズごとのステータスを確認可能。

#### ステップ 8: 詳細なビルド情報を表示する
CloudWatch Logs の末尾 10,000 行のログを確認可能。

#### ステップ 9: ビルド出力アーティファクトを取得する
出力用の S3 バケット内にビルドの成果物が生成されていることを確認する。


[AWS CLI を使用した AWS CodeBuild の開始方法](https://docs.aws.amazon.com/ja_jp/codebuild/latest/userguide/getting-started-cli.html)

```shell
# パラメータテンプレートを生成 
aws codebuild create-project --generate-cli-skeleton
# ビルドプロジェクトの作成
aws codebuild create-project --cli-input-json file://create-project.json

# ビルドの開始
aws codebuild start-build --project-name project-name

# 要約されたビルド情報の表示
aws codebuild batch-get-builds --ids id
```

## サンプル

[CodeBuild でソースプロバイダにアクセストークンを使用する](https://docs.aws.amazon.com/ja_jp/codebuild/latest/userguide/sample-access-tokens.html)

アクセストークンを使用して GitHubまたは Bitbucket に接続する方法のサンプル。
ソースプロバイダーを選択し、画面の指示に従って認証情報を入力する。

[Amazon ECR サンプル](https://docs.aws.amazon.com/ja_jp/codebuild/latest/userguide/sample-ecr.html)

ビルドを行うコンテナイメージとして自前のものを使用するサンプル。
サービスロールに ECR の権限を付与する必要あり。また、CodeBuild プロジェクトにおいて、次の例のように設定する。

```json
  "environment": {
    "type": "LINUX_CONTAINER",
    "image": "account-ID.dkr.ecr.region-ID.amazonaws.com/your-Amazon-ECR-repo-name:tag",
    "computeType": "BUILD_GENERAL1_SMALL"
  },
```

[AWS CodeDeploy サンプル](https://docs.aws.amazon.com/ja_jp/codebuild/latest/userguide/sample-codedeploy.html)

ソースは次の内容で用意。
```
(root directory name)
     `-- my-app
           |-- buildspec.yml
           |-- appspec.yml
           |-- pom.xml
           `-- src    
                ...
```                                               

buildspec.yaml の出力アーティファクトは次の内容になる。出力にビルドの成果物と appspec.yml を指定する。
```yaml
artifacts:
  files:
    - target/my-app-1.0-SNAPSHOT.jar
    - appspec.yml
  discard-paths: yes
```

[AWS CodePipeline を CodeBuild の複数の入力ソースおよび出力アーティファクトと統合するサンプル](https://docs.aws.amazon.com/ja_jp/codebuild/latest/userguide/sample-pipeline-multi-input-output.html)

CodePipeline でビルドステージに複数の入力、複数の出力を行うサンプル。

* CodePipeline 側では入力ソースの一つを PrimarySource として指定する必要がある。PrimarySource 内の buildspec.yaml を使用する。
* 入力ソースは環境変数にパスが格納されている。プライマリソースは $CODEBUILD_SRC_DIR。その他は $CODEBUILD_SRC_DIR_yourInputArtifactName
* 出力アーティファクトは buildspec.yaml 内で次のように指定する。artifact1, artifact2 の部分は CodePipeline の outputArtifacts の文字列と同一である必要がある。

```yaml
artifacts:
  secondary-artifacts:
    artifact1:
      base-directory: $CODEBUILD_SRC_DIR
      files:
        - source1_file
    artifact2:
      base-directory: $CODEBUILD_SRC_DIR_source2
      files:
        - source2_file
```

[AWS Elastic Beanstalk サンプル](https://docs.aws.amazon.com/ja_jp/codebuild/latest/userguide/sample-elastic-beanstalk.html)

CodePipeline を使用して CodeBuild でビルドし、Elastic Beanstalk の環境にデプロイ可能。

ソースは次の内容で用意。
```
(root directory name)
     `-- my-app
           |-- .ebextensions
           |     `-- fix-path.config
           |-- buildspec.yml
           |-- pom.xml
           `-- src    
                ...
```                                               

buildspec.yaml の出力アーティファクトは次の内容になる。出力にビルドの成果物と .ebextensions を指定する。

```yaml
artifacts:
  files:
    - my-web-app.war
    - .ebextensions/**/*
  base-directory: 'target/my-web-app'
```

[Docker サンプル](https://docs.aws.amazon.com/ja_jp/codebuild/latest/userguide/sample-docker.html)

Docker イメージをビルドし、ECR にプッシュするサンプル。

```Dockerfile
FROM golang:1.12-alpine AS build
#Install git
RUN apk add --no-cache git
#Get the hello world package from a GitHub repository
RUN go get github.com/golang/example/hello
WORKDIR /go/src/github.com/golang/example/hello
# Build the project and send the output to /bin/HelloWorld 
RUN go build -o /bin/HelloWorld

FROM golang:1.12-alpine
#Copy the build's output binary from the previous build container
COPY --from=build /bin/HelloWorld /bin/HelloWorld
ENTRYPOINT ["/bin/HelloWorld"]
```

buildspec.yaml。環境変数はビルドプロジェクトにおいて設定する。
```yaml
version: 0.2

phases:
  pre_build:
    commands:
      - echo Logging in to Amazon ECR...
      - aws ecr get-login-password --region $AWS_DEFAULT_REGION | docker login --username AWS --password-stdin $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com
  build:
    commands:
      - echo Build started on `date`
      - echo Building the Docker image...          
      - docker build -t $IMAGE_REPO_NAME:$IMAGE_TAG .
      - docker tag $IMAGE_REPO_NAME:$IMAGE_TAG $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/$IMAGE_REPO_NAME:$IMAGE_TAG      
  post_build:
    commands:
      - echo Build completed on `date`
      - echo Pushing the Docker image...
      - docker push $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/$IMAGE_REPO_NAME:$IMAGE_TAG
```



## ビルドを計画する

[ビルドを実行する](https://docs.aws.amazon.com/ja_jp/codebuild/latest/userguide/planning.html)

計画段階で考慮が必要なことは次の通り。

* ソースプロバイダ。指定できるものは以下の通り。
  * CodeCommit
  * Amazon S3
  * GitHub
* CodeBuild の入出力（ソースプロバイダ、出力アーティファクト）
* ビルドに必要なランタイムは何か。
* 使用する AWS リソース（サービスロールへの権限の付与）
* VPC と連携させるか。

#### buildspec のリファレンス

[ビルド仕様 (buildspec) に関するリファレンス](https://docs.aws.amazon.com/ja_jp/codebuild/latest/userguide/build-spec-ref.html)

デフォルトの buildspec ファイルの名前は buildspec.yml で、ソースディレクトリのルートに配置する必要がある。
ファイル名と場所は変更可能。

構文は以下の通り。
```yaml
version: 0.2

run-as: Linux-user-name

env:
  shell: shell-tag
  variables:
    key: "value"
    key: "value"
  parameter-store:
    key: "value"
    key: "value"
  exported-variables:
    - variable
    - variable
  secrets-manager:
    key: secret-id:json-key:version-stage:version-id
  git-credential-helper: no | yes

proxy:
  upload-artifacts: no | yes
  logs: no | yes

batch:
  fast-fail: false | true
  # build-list:
  # build-matrix:
  # build-graph:
        
phases:
  install:
    run-as: Linux-user-name
    runtime-versions:
      runtime: version
      runtime: version
    commands:
      - command
      - command
    finally:
      - command
      - command
  pre_build:
    run-as: Linux-user-name
    commands:
      - command
      - command
    finally:
      - command
      - command
  build:
    run-as: Linux-user-name
    commands:
      - command
      - command
    finally:
      - command
      - command
  post_build:
    run-as: Linux-user-name
    commands:
      - command
      - command
    finally:
      - command
      - command
reports:
  report-group-name-or-arn:
    files:
      - location
      - location
    base-directory: location
    discard-paths: no | yes
    file-format: report-format
artifacts:
  files:
    - location
    - location
  name: artifact-name
  discard-paths: no | yes
  base-directory: location
  secondary-artifacts:
    artifactIdentifier:
      files:
        - location
        - location
      name: secondary-artifact-name
      discard-paths: no | yes
      base-directory: location
    artifactIdentifier:
      files:
        - location
        - location
      discard-paths: no | yes
      base-directory: location
cache:
  paths:
    - path
    - path
```

上記構文のうち一部について説明する。

**version**

buildspec のバージョンを表します。0.2 を使用することが推奨される。

**env**

環境変数を指定。

* env/variables: プレーンテキストで良い場合
* env/parameter-store: SSM のパラメータストアに格納している場合
* env/secrets-manager: Secrets Manager に格納している場合
* env/exported-variables: 実際に export する環境変数を列挙

**phases/install**

ビルド環境でのパッケージのインストールなどを行うためのフェーズ。
例えば以下のように runtime-versions の指定が可能。利用可能なランタイムはリファレンスを参照すること。

```yaml
phases:
  install:
    runtime-versions:
      java: corretto8
```

**phases/pre_build**

ビルドの前に作業を行うためのフェーズ。ECR へのサインインなど。

**phases/build**

ビルドを行うためのフェーズ。

**phases/post_build**

ビルドの後に作業を行うためのフェーズ。ECR へのプッシュなど。

**reports**

レポート（テストレポートなど）を行うために使用。

**artifacts/files**

出力アーティファクトの場所。
個別のファイルを列挙することもできるし、全てを対象とする場合は '**/*' のような再帰的に表す形式で指定する。

**artifacts/secondary-artifacts**

２つ以上の出力アーティファクトがある場合に使用。

**cache/paths**

S3 バケットにアップロードしたいキャッシュのパスを指定。


[CodeBuild に用意されている Docker イメージ](https://docs.aws.amazon.com/ja_jp/codebuild/latest/userguide/build-env-ref-available.html)

Docker イメージのリストとランタイムが一覧されている。Docker イメージの内容は GitHub 上で公開されている。

Docker イメージ例：

* aws/codebuild/amazonlinux2-x86_64-standard:3.0

ランタイム例：

* golang: 1.14


[ビルド環境のコンピューティングタイプ](https://docs.aws.amazon.com/ja_jp/codebuild/latest/userguide/build-env-ref-compute-types.html)

コンピューティング環境例

* build.general1.medium


[ビルド環境の環境変数](https://docs.aws.amazon.com/ja_jp/codebuild/latest/userguide/build-env-ref-env-vars.html)

いくつかの環境変数が用意されている。以下は一例。

* CODEBUILD_SRC_DIR: ビルドのディレクトリパス
* CODEBUILD_BUILD_SUCCEEDING: 現在のビルドが成功かどうか。


[AWS CodeBuild エージェントを使用したローカルでのテストおよびデバッグ](https://docs.aws.amazon.com/ja_jp/codebuild/latest/userguide/use-codebuild-agent.html)

ローカルマシンでビルドをテスト、デバッグ可能。



## VPC サポート

[Amazon Virtual Private Cloud による AWS CodeBuild の使用](https://docs.aws.amazon.com/ja_jp/codebuild/latest/userguide/vpc-support.html)

通常のビルドは VPC 外で実行されるが、VPC 内で実行するように設定することもできる。
ユースケースとしては次の通り。

* VPC 内のリソースに対してテストを実行する。
* VPC 内にあるセルフホストしているリポジトリを使用している。
* 固定 IP アドレスからの接続が必要な外部サービスに対して NAT Gateway を介してアクセスする。

プロジェクトの設定において VPC を設定可能。



## AWS CodeBuild でのビルドプロジェクトとビルドの使用

[ビルドプロジェクトの作成 (コンソール)](https://docs.aws.amazon.com/ja_jp/codebuild/latest/userguide/create-project-console.html)

設定項目。

* プロジェクト名
* 説明
* ソースプロバイダー(S3, CodeCommit, Bitbucket, GitHub, GitHub Enterprise Server)
* 環境
  * イメージ(マネージドイメージの場合は OS, ランタイム, イメージ, ランタイムバージョン)
  * Privileged
  * サービスロール
  * 追加設定
    * タイムアウト
    * VPC
    * コンピューティング
    * 環境変数(Systems Manager パラメータストア, AWS Secrets Manager を使用することも可能)
* buildspec
* バッチ設定
* アーティファクト
  * タイプ(S3 バケット名などの情報)
  * 追加設定
    * 暗号化キー
    * キャッシュタイプ
* ログ
  * CloudWatch Logs
  * S3


[AWS CodeBuild でのキャッシュのビルド](https://docs.aws.amazon.com/ja_jp/codebuild/latest/userguide/build-caching.html)

次の２つをキャッシュの置き場所として利用可能。

* Amazon S3 キャッシュ
* ローカルキャッシュ


[ビルドの実行 (コンソール)](https://docs.aws.amazon.com/ja_jp/codebuild/latest/userguide/run-build-console.html)

ソースプロバイダーの中からどのバージョン、コミット ID を使用するかを設定する。
また、ビルドプロジェクトの設定の多くを上書き可能。


[Session Manager で実行中のビルドを表示する](https://docs.aws.amazon.com/ja_jp/codebuild/latest/userguide/session-manager.html)

ビルドを一時停止し、Session Manager でコンテナに接続可能。
次のように buildspec.yml ファイルにて、ブレークポイントを設定する。
ビルド開始時に [セッション接続を有効にする] を設定した場合は、ブレークポイントで処理が中断されるようになる。

```yaml
phases:
  pre_build:
    commands:
    ...
      - codebuild-breakpoint
```

処理を再開する場合はコンテナ内から次のコマンドを実行する。
```
$ codebuild-resume
```


# 参考

* Document
  * [CodeBuild とは](https://docs.aws.amazon.com/ja_jp/codebuild/latest/userguide/welcome.html)
* サービス紹介ページ
  * [AWS CodeBuild](https://aws.amazon.com/jp/codebuild/)
  * [よくある質問](https://aws.amazon.com/jp/codebuild/faqs/)
* [ナレッジセンター](https://aws.amazon.com/jp/premiumsupport/knowledge-center/#AWS_CodeBuild)
* Black Belt
  * [AWS Black Belt Online Seminar 2017 AWS Code Services ( CodeCommit, CodeBuild )](https://www.slideshare.net/AmazonWebServicesJapan/aws-black-belt-online-seminar-2017-aws-code-services-codecommit-codebuild)

