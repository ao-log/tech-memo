
#   ECR

[Amazon Elastic Container Registry とは](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/what-is-ecr.html)

ECR はコンテナイメージのレジストリサービス。

## 用語

* レジストリ: AWS アカウントごとにレジストリが用意されている。レジストリ内にリポジトリを作ることができる。
* リポジトリ: Docker イメージの他、OCI イメージ、OCI 互換のイメージも保存可能。
* リポジトリポリシー: リポジトリ、リポジトリ内のイメージに対してアクセス権を設定可能。


## AWS CLI を使用した開始方法

[AWS CLI を使用した Amazon ECR の開始方法](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/getting-started-cli.html)

```shell
# デフォルトレジストリに対して認証
aws ecr get-login-password --region region | docker login --username AWS --password-stdin aws_account_id.dkr.ecr.region.amazonaws.com

# リポジトリ作成
aws ecr create-repository \
    --repository-name hello-world \
    --image-scanning-configuration scanOnPush=true \
    --region ap-northeast-1

# イメージのタグ付け
docker tag hello-world:latest aws_account_id.dkr.ecr.ap-northeast-1.amazonaws.com/hello-world:latest

# イメージのプッシュ
docker push aws_account_id.dkr.ecr.ap-northeast-1.amazonaws.com/hello-world:latest

# イメージのプル
docker pull aws_account_id.dkr.ecr.ap-northeast-1.amazonaws.com/hello-world:latest

# リポジトリからイメージを削除
aws ecr batch-delete-image \
      --repository-name hello-world \
      --image-ids imageTag=latest

# リポジトリを削除
aws ecr delete-repository \
      --repository-name hello-world \
      --force
```


## レジストリ

[レジストリ](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/Registries.html)

レジストリの URL は ```https://aws_account_id.dkr.ecr.region.amazonaws.com```

```aws ecr get-login-password``` コマンドにより認証する。認証トークンの有効時間は 12 時間。


## リポジトリ

[リポジトリの作成](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/repository-create.html)

設定項目

* タグのイミュータビリティ: 有効にすると、同じタグを使用した後続イメージのプッシュによるイメージタグ上書きを防ぐ。
* プッシュ時にスキャン: プッシュ時にスキャンするかどうか
* KMS 暗号化: 有効化した場合、AWS 管理キーもしくはカスタマーマスターキーを選ぶことができる。


[リポジトリポリシー](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/repository-policies.html)

IAM ポリシーもしくはリポジトリポリシーのどちらかで許可が設定されている場合は、許可される。
どちらか片方で拒否が設定されている場合は、拒否される。


[リポジトリ ポリシーの例](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/repository-policy-examples.html)


## イメージ

[イメージのプッシュ](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/docker-push-ecr-image.html)

```shell
docker tag e9ae3c220b23 aws_account_id.dkr.ecr.region.amazonaws.com/my-web-app
```
このようにイメージタグを省略した場合は latest とみなされる。


[helm チャートをプッシュ](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/push-oci-artifact.html)

helm チャートをプッシュすることが可能。


[イメージにもう一度タグを付ける](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/image-retag.html)

プッシュ済みのイメージにタグを付与することが可能。

```shell
# イメージマニフェストを取得
MANIFEST=$(aws ecr batch-get-image --repository-name amazonlinux --image-ids imageTag=latest --query 'images[].imageManifest' --output text)

# イメージマニフェストを指定してタグを付与
aws ecr put-image --repository-name amazonlinux --image-tag 2017.03 --image-manifest "$MANIFEST"
```


[ライフサイクルポリシー](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/LifecyclePolicies.html)

イメージのクリーンアップの自動化方法を定義できる。リアルタイムに削除されるわけではなく、24 時間以内にイメージが expire される。
日数、世代数などの条件で設定可能。


[イメージタグのイミュータビリティ](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/image-tag-mutability.html)

リポジトリにこの設定を行うことで、既に存在しているタグ付きイメージをプッシュした場合に ImageTagAlreadyExistsException エラーが返される。


[イメージスキャン](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/image-scanning.html)

オープンソースの Clair プロジェクトから CVE データベースを使用してスキャンしている。
リポジトリの設定でプッシュ時にスキャンするようにすることができる。また手動でスキャンすることも可能。なお、イメージに対するスキャンは 1 日 1 回しかできない。


[Amazon Linux コンテナイメージ](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/amazon_linux_container_image.html)

Amazon Linux のコンテナイメージは ECR のリポジトリで公開されている。
次のコマンドで pull することができる。

```shell
docker pull 137112412989.dkr.ecr.us-east-1.amazonaws.com/amazonlinux:latest
```


## セキュリティ

[IAM ポリシーの例](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/security_iam_id-based-policy-examples.html)

特定のリポジトリのみアクションを許可するなど、いくつかのポリシー例が載っている。


[タグベースのアクセスコントロールを使用する](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/ecr-supported-iam-actions-tagging.html)

タグを使用したアクセスコントロールが可能。
例としては、特定のタグが付与されたリポジトリには ECR の各アクションを実行不可のように設定可能。


[暗号化](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/encryption-at-rest.html)

KMS による暗号化設定について書かれている。


[インタフェース VPC エンドポイント](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/vpc-endpoints.html)

* EC2 起動タイプを使用する場合: ECS のエンドポイントも必要
+ Fargate 起動タイプ(プラットフォームバージョン 1.3 以下) の場合: com.amazonaws.region.ecr.dkr と S3 のエンドポイント
+ Fargate 起動タイプ(プラットフォームバージョン 1.4 以降) の場合: com.amazonaws.region.ecr.dkr、com.amazonaws.region.ecr.api と S3 のエンドポイント
* awslogs ドライバーを使用して CloudWatch Logs へログを送信する場合: CloudWatch Logs VPC エンドポイントが必要



## モニタリング

[EventBridge](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/ecr-eventbridge.html)

イメージプッシュの完了時、スキャン完了時、イメージ削除時などにイベントが送信される。


[CloudTrail](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/logging-using-cloudtrail.html)

記録される API について、いくつかの例が載っている。

* CreateRepository: リポジトリ作成時
* PutImage: イメージのプッシュ (プッシュ時は InitiateLayerUpload、UploadLayerPart、CompleteLayerUpload も記録される)
* BatchGetImage: イメージのプル (プル時は GetDownloadUrlForLayer も記録される)
* PolicyExecutionEvent: ライフサイクルポリシーの実行内容が記録される



## トラブルシューティング

[トラブルシューティング](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/troubleshooting.html)

[Amazon ECR 使用時の Docker コマンドのエラーのトラブルシューティング](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/common-errors-docker.html)

#### イメージプル時の "Filesystem Verification Failed" or "404: Image Not Found"

* ローカルディスクがいっぱいである
* ネットワークエラー。ECR の場合は更に S3 へのネットワークアクセスも必要

#### イメージプッシュ時の HTTP 403 Errors or "no basic auth credentials" Error

* 別のリージョンに対して認証している
* リポジトリへのアクセス権がない


[Amazon ECR エラーメッセージのトラブルシューティング](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/common-errors.html)

#### HTTP 429: Too Many Requests or ThrottleException

スロットリングされているために発生しているエラー。エクスポネンシャルバックオフで再試行するのが有効。

#### HTTP 403: "User [arn] is not authorized to perform [operation]"

```aws ecr get-login``` 時のエラー。IAM ユーザのポリシー設定を見直すこと。

#### HTTP 404: "Repository Does Not Exist" Error

リポジトリが存在しないため。ECR の仕様として先にリポジトリを作成しておく必要がある。



## ナレッジセンター

[Amazon ECS の Amazon ECR エラー「CannotPullContainerError: API error」を解決する方法を教えてください。](https://aws.amazon.com/jp/premiumsupport/knowledge-center/ecs-pull-container-api-error-ecr/)

* 起動タイプに Amazon ECR エンドポイントへのアクセス権がない
  * プライベートサブネットにある場合、NAT ゲートウェイへのルートがあるかどうか。もしくは PrivateLink が設定されているかどうか
  * パブリックサブネットにある場合、EC2 起動タイプの場合はインスタンスにパブリック IP アドレスが割り当てられているかどうか、Fargate の場合は自動割当パブリック IP が ENABLED になっているかどうか。
  * PrivateLink を使用している場合は、ECR のインターフェイス VPC エンドポイントに関連付けられたセキュリティグループにおいてコンテナインスタンスもしくは Fargate タスクからの HTTPS インバウンド通信が許可されているかどうか
  * コンテナインスタンスもしくは Fargate タスクにアタッチされたセキュリティグループにおいて HTTPS のアウトバウンド通信が許可されているかどうか
* Amazon ECR リポジトリポリシーで、リポジトリイメージへのアクセスが制限されている
* AWS Identity and Access Management (IAM) ロールに、イメージをプルまたはプッシュするための適切なアクセス許可がない
* イメージが見つからない
* Amazon Simple Storage Service (Amazon S3) アクセス許可が、Amazon Virtual Private Cloud (Amazon VPC) ゲートウェイエンドポイントポリシーによって拒否されている


# 参考

* Document
  * [Amazon Elastic Container Registry とは](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/what-is-ecr.html)
  * [API Reference](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/APIReference/Welcome.html)
* サービス紹介ページ
  * [Amazon Elastic Container Registry](https://aws.amazon.com/jp/ecr/)
  * [よくある質問](https://aws.amazon.com/jp/ecr/faqs/)

