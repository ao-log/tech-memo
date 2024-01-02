
#   ECR

[Amazon Elastic Container Registry とは](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/what-is-ecr.html)

ECR はコンテナイメージのレジストリサービス


## 用語

* レジストリ: AWS アカウントごとにレジストリが用意されている。レジストリ内にリポジトリを作ることができる
* リポジトリ: Docker イメージの他、OCI イメージ、OCI 互換のイメージも保存可能
* リポジトリポリシー: リポジトリ、リポジトリ内のイメージに対してアクセス権を設定可能



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

レジストリの URL は `https://aws_account_id.dkr.ecr.region.amazonaws.com`

* `aws ecr get-login-password` コマンドにより認証する。認証トークンの有効時間は 12 時間。
* `GetAuthorizationToken` API が実行されエンコードされたパスワードを含む base64 エンコード認可トークンを取得

Docker Registry HTTP API による操作も可能
```
TOKEN=$(aws ecr get-authorization-token --output text --query 'authorizationData[].authorizationToken')

curl -i -H "Authorization: Basic $TOKEN" https://aws_account_id.dkr.ecr.region.amazonaws.com/v2/amazonlinux/tags/list
```


[プライベートレジストリの設定](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/registry-settings.html)

以下の設定項目がある

* Registry permissions: レプリケーション機能、プルスルーキャッシュに関する許可を設定
* Pull through cache rules: Pull through cache ルール を設定
* Replication: クロスリージョン or クロスアカウントのレプリケーションを設定
* Scanning configuration: ベーシックスキャン or 拡張スキャンを設定


[プライベートレジストリの許可](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/registry-permissions.html)

* デフォルトでは自アカウントのクロスリージョンレプリケーションの権限はある。クロスアカウント時には設定が必要
* ecr:ReplicateImage: クロスアカウントレプリケーションの許可設定
* ecr:BatchImportUpstreamImage: Pull through cache のアクセス許可を付与
* ecr:CreateRepository: リポジトリ作成を許可



## リポジトリ

[リポジトリの作成](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/repository-create.html)

設定項目

* タグのイミュータビリティ: 有効にすると、同じタグを使用した後続イメージのプッシュによるイメージタグ上書きを防ぐ
* プッシュ時にスキャン: プッシュ時にスキャンするかどうか
* KMS 暗号化: 有効化した場合、AWS 管理キーもしくはカスタマーマスターキーを選ぶことができる


[リポジトリポリシー](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/repository-policies.html)

* IAM ポリシーもしくはリポジトリポリシーのどちらかで許可が設定されている場合は、許可される。どちらか片方で拒否が設定されている場合は、拒否される
* プル、プッシュを行う前に `ecr:GetAuthorizationToken` が必要
* `ecr:GetAuthorizationToken` は IAM ポリシー側での許可が必要


[リポジトリ ポリシーの例](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/repository-policy-examples.html)



## イメージ

#### Push image

[イメージのプッシュ](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/docker-push-ecr-image.html)

```shell
docker tag e9ae3c220b23 aws_account_id.dkr.ecr.region.amazonaws.com/my-web-app
```
このようにイメージタグを省略した場合は latest とみなされる。


[マルチアーキテクチャイメージのプッシュ](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/docker-push-multi-architecture-image.html)

以下の例のようにマニフェストを作成し、プッシュする
```
docker manifest create \
  aws_account_id.dkr.ecr.us-west-2.amazonaws.com/my-repository \
  aws_account_id.dkr.ecr.us-west-2.amazonaws.com/my-repository:image_one_tag \
  aws_account_id.dkr.ecr.us-west-2.amazonaws.com/my-repository:image_two

docker manifest push aws_account_id.dkr.ecr.us-west-2.amazonaws.com/my-repository
```


[helm チャートをプッシュ](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/push-oci-artifact.html)

helm チャートをプッシュすることが可能。
以下の例のような流れ。
```
helm create helm-test-chart
rm -rf ./helm-test-chart/templates/*

cd helm-test-chart/templates
cat <<EOF > configmap.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: helm-test-chart-configmap
data:
  myvalue: "Hello World"
EOF

cd ../..
helm package helm-test-chart

aws ecr create-repository \
     --repository-name helm-test-chart \
     --region us-west-2

aws ecr get-login-password \
     --region us-west-2 | helm registry login \
     --username AWS \
     --password-stdin aws_account_id.dkr.ecr.us-west-2.amazonaws.com

helm push helm-test-chart-0.1.0.tgz oci://aws_account_id.dkr.ecr.us-west-2.amazonaws.com/
```


#### Signing

[イメージの署名](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/image-signing.html)

* Notation CLI と Notation 用の AWS Signer プラグインがインストールされていることが前提

以下の流れで署名する
```
// 署名プロファイルの作成
aws signer put-signing-profile --profile-name ecr_signing_profile --platform-id Notation-OCI-SHA384-ECDSA

// Notation CLI を Amazon ECR プライベートレジストリに対して認証
aws ecr get-login-password --region region | notation login --username AWS --password-stdin 111122223333.dkr.ecr.region.amazonaws.com

// 署名
notation sign 111122223333.dkr.ecr.region.amazonaws.com/curl@sha256:ca78e5f730f9a789ef8c63bb55275ac12dfb9e8099e6EXAMPLE  \
    --plugin "com.amazonaws.signer.notation.plugin"  \
    --id "arn:aws:signer:region:111122223333:/signing-profiles/ecrSigningProfileName"
```

署名の検証。検証はトラストストアが必要
```
notation verify 111122223333.dkr.ecr.region.amazonaws.com/curl@SHA256_digest
```

削除時は ORAS CLI を使用してアーティファクトを削除すると、ORAS がイメージインデックスの更新または削除を処理されるので、この方法が推奨される


#### Pull through cache

[プルスルーキャッシュルールの使用](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/pull-through-cache.html)

* リモートパブリックレジストリ内のリポジトリのキャッシュをサポート
* キャッシュイメージの URI からプルすると、リモートリポジトリを 24 時間に 1 回までチェックしキャッシュされたイメージが最新バージョンか確認
* アップストリームからイメージをプルできない状況では、キャッシュが利用される
* マルチアーキテクチャイメージをプルすると、マニフェストリストとマニフェストリストで参照されている各イメージがプルされる。特定のアーキテクチャのみをプルする場合、アーキテクチャに関連付けられたイメージダイジェストまたはタグを使用してイメージをプルすること
* サービスにリンクされた IAM ロールを使用。キャッシュされたイメージのリポジトリを作成し、キャッシュされたイメージをプッシュするために必要なアクセス許可を提供
* レジストリポリシーにより権限を細かく制御可能
* タグのイミュータビリティを手動でオンにした場合、キャッシュされたイメージを更新できない場合がある
* 初回のプルではインターネットへのアウトバウンドの疎通性が必要
* 必要な IAM 許可
  * `ecr:CreatePullThroughCacheRule`: プルスルーキャッシュルールを作成するアクセス許可
  * `ecr:BatchImportUpstreamImage`: 外部イメージを取得し、プライベートレジストリにインポートするアクセス許可
  * `ecr:CreateRepository`: リポジトリを作成するアクセス許可


[プルスルーキャッシュルールの作成](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/pull-through-cache-creating-rule.html)

* 認証情報が必要な場合は、Secrets Manager にシークレットを格納し、ARN を指定する


[アップストリームリポジトリの認証情報を AWS Secrets Manager シークレットに保存する](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/pull-through-cache-creating-secret.html)

* シークレット名は `ecr-pullthroughcache/` をプレフィックスとする必要がある
* コンテナレジストリごとに手順が用意されている


#### Delete image

[イメージの削除](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/delete_image.html)

* batch-delete-image により削除する
* 次のコマンドによりタグを削除する。最後のタグ削除時にイメージが削除される
```
aws ecr batch-delete-image \
     --repository-name my-repo \
     --image-ids imageTag=tag1 imageTag=tag2
```
* イメージダイジェストを指定した場合、イメージとタグは削除される
```
aws ecr batch-delete-image \
     --repository-name my-repo \
     --image-ids imageDigest=sha256:4f70ef7a4d29e8c0c302b13e25962d8f7a0bd304EXAMPLE
```


#### Retag

[イメージにもう一度タグを付ける](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/image-retag.html)

プッシュ済みのイメージにタグを付与することが可能。

```shell
# イメージマニフェストを取得
MANIFEST=$(aws ecr batch-get-image --repository-name amazonlinux --image-ids imageTag=latest --query 'images[].imageManifest' --output text)

# イメージマニフェストを指定してタグを付与
aws ecr put-image --repository-name amazonlinux --image-tag 2017.03 --image-manifest "$MANIFEST"
```


#### Replication

[プライベートイメージのレプリケーション](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/replication.html)

* クロスアカウント、クロスリージョンのレプリケーションが可能
* 複製されるのは設定後にプッシュしたイメージのみ
* 削除は実行されない
* クロスアカウント時にはレプリケーション先にてレジストリポリシーによる許可が必要
* リポジトリ設定は複製されない
* イミュータブル性が有効になっている場合、イメージのタグづけが解除される可能性がある
* フィルター設定により複製対象を絞り込むことが可能
* AWS パーティション間ではレプリケーション不可。us-west-2 のリポジトリは cn-north-1 にレプリケートできない


[レプリケーションステータスの表示](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/replication-status.html)

以下コマンドでレプリケーションステータスを表示できる
```
aws ecr describe-image-replication-status \
     --repository-name repository_name \
     --image-id imageTag=image_tag \
     --region us-west-2
```


#### Lifecycle Policies

[ライフサイクルポリシー](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/LifecyclePolicies.html)

* イメージのクリーンアップの自動化方法を定義できる
* リアルタイムに削除されるわけではなく、24 時間以内にイメージが expire となる
* tagPrefixList で複数タグが指定された場合は、and 条件。or ではない
* imageCountMoreThan では pushed_at_time 順にイメージを並べ、指定したカウントよりも大きいイメージは古い順に expire となる
* より優先度の高いルールにて識別対象となったイメージは削除対象から除外される
* 日数、世代数などの条件で設定可能。パラメータは以下の通り
  * rulePriority: 数値が低いものから評価していく
  * description
  * selection
    * tagStatus: tagged | untagged | any
    * tagPrefixList
    * countType: imageCountMoreThan | sinceImagePushed
    * countUnit: sinceImagePushed の場合のみ。days
    * countNumber: imageCountMoreThan の場合は世代数。sinceImagePushed の場合はイメージの最大期限
  * action
    * type: expire


[ライフサイクルポリシーの例](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/lifecycle_policy_examples.html)

以下設定の場合、タグづけされていないイメージを 1 個だけ保持し 残りは削除する。もっとも新しいもののみが残る
```json
            "selection": {
                "tagStatus": "untagged",
                "countType": "imageCountMoreThan",
                "countNumber": 1
            },
```

以下設定の場合、`tagPrefixList` は and 条件。よって、片方のタグのみ指定されたイメージは削除対象とならない
```json
            "selection": {
                "tagStatus": "tagged",
                "tagPrefixList": ["alpha", "beta"],
                "countType": "sinceImagePushed",
                "countNumber": 5,
                "countUnit": "days"
            },
```

以下設定の場合でも、より高いルールで評価済みのイメージについては削除対象とならない。残りのイメージから削除対象を決める
```json
            "selection": {
                "tagStatus": "any",
                "countType": "imageCountMoreThan",
                "countNumber": 1
            },
```


#### Image tag mutability

[イメージタグのイミュータビリティ](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/image-tag-mutability.html)

リポジトリにこの設定を行うことで、既に存在しているタグ付きイメージをプッシュした場合に `ImageTagAlreadyExistsException` エラーが返される。


#### Image scan

[イメージスキャン](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/image-scanning.html)

* 拡張スキャン
  * 新しい脆弱性が発生するとスキャンされる
  * OS のほかアプリケーションパッケージもスキャン対象
  * 最初に有効化した際は過去 30 日分のイメージがスキャン対象になる
  * フィルター
    * 継続スキャン、プッシュ時にスキャン、それぞれについて別のフィルターを設定できる
    * 一致しないリポジトリは、スキャンが無効に設定される
  * 継続スキャン、プッシュ時にスキャンの設定が可能
    * 継続スキャンではスキャン期間を設定可能。デフォルトは Lifetime だが、180 day, 30 day に設定することも可能
  * 手動スキャンはできない
  * スキャン結果はイベントとして発行される
  * Amazon Inspector サービスにリンクされた IAM ロールが使用される。拡張スキャン有効時に Inspector によって自動的に作成される
* ベーシックスキャン
  * オープンソースの Clair プロジェクトの CVE データベースを使用
  * プッシュ時にスキャンするほか、手動スキャンも設定可能。手動スキャンは 1 日 1 回のみ


#### Image manifest format

[コンテナイメージマニフェストの形式](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/image-manifest-formats.html)

* 次のマニフェスト形式がサポートされる
  * Docker Image Manifest V2 Schema 1 (Docker バージョン 1.9 以前で使用)
  * Docker Image Manifest V2 Schema 2 (Docker バージョン 1.10 以降で使用)
  * Open Container Initiative (OCI) 仕様 (v1.0 以降)
* イメージをプルするときクライアントが理解できない形式の場合は ECR によって理解できる形式に変換される


#### How to use image

[Amazon ECS での Amazon ECR イメージの使用](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/ECR_on_ECS.html)

* 必要なポリシー
  * EC2, 外部インスタンスの場合: インスタンスロールに [AmazonEC2ContainerServiceforEC2Role](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/security-iam-awsmanpol.html#security-iam-awsmanpol-AmazonEC2ContainerServiceforEC2Role)
  * Fargate の場合: タスク実行ロールに [AmazonECSTaskExecutionRolePolicy](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/security-iam-awsmanpol.html#security-iam-awsmanpol-AmazonECSTaskExecutionRolePolicy)


[Amazon EKS での Amazon ECR イメージの使用](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/ECR_on_EKS.html)

* 必要なポリシー
  * EC2 の場合: NodeInstanceRole

Helm チャートにも対応
```
helm install ecr-chart-demo oci://aws_account_id.dkr.ecr.region.amazonaws.com/helm-test-chart --version 0.1.0
```



## Security

[Amazon Elastic Container Regist の AWS 管理ポリシー](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/security-iam-awsmanpol.html)

* AmazonEC2ContainerRegistryFullAccess
* AmazonEC2ContainerRegistryPowerUser
* AmazonEC2ContainerRegistryReadOnly
* AWSECRPullThroughCache_ServiceRolePolicy: サービスにリンクされたロールにアタッチされている
* ECRReplicationServiceRolePolicy: サービスにリンクされたロールにアタッチされている


[Amazon ECR でのサービスにリンクされたロールの使用](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/using-service-linked-roles.html)

* レプリケーションとプルスルーキャッシュの利用に必要なアクセス許可が含まれている


[IAM ポリシーの例](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/security_iam_id-based-policy-examples.html)

* 特定のリポジトリのみアクションを許可するなど、いくつかのポリシー例が載っている


[タグベースのアクセスコントロールを使用する](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/ecr-supported-iam-actions-tagging.html)

* タグを使用したアクセスコントロールが可能
* `aws:RequestTag` によりリソース作成時に指定タグを付与しないと失敗する
* `ecr:ResourceTag` により指定タグのリソースのみにアクセスを制限できる


[暗号化](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/encryption-at-rest.html)

* KMS のカスタマーマネージドキーによる暗号化が可能
  * CreateGrant により ECR がデータキーを使用して暗号化、復号できるようにする
  * プッシュ時にデータキーを生成し、データキーにより暗号化。イメージレイヤーメタデータ、イメージマニフェストと共に保存
  * プル時に暗号化されたデータキーを復号して S3 に送信。イメージレイヤーを復号してプル


[インタフェース VPC エンドポイント](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/vpc-endpoints.html)

* EC2 起動タイプを使用する場合: ECS のエンドポイントも必要
+ Fargate 起動タイプ(プラットフォームバージョン 1.3 以下) の場合: com.amazonaws.region.ecr.dkr と S3 のエンドポイント
+ Fargate 起動タイプ(プラットフォームバージョン 1.4 以降) の場合: com.amazonaws.region.ecr.dkr、com.amazonaws.region.ecr.api と S3 のエンドポイント
* awslogs ドライバーを使用して CloudWatch Logs へログを送信する場合: CloudWatch Logs VPC エンドポイントが必要
* ECR Public は未サポート
* プルスルーキャッシュを使用する場合は、初回のプル時のみインターネットへのアウトバウンドの疎通性が必要
* Windows イメージの場合、ライセンスによって配布が制限されているアーティファクトが含まれており、通常は当該箇所はプッシュされない。Docker デーモンの --allow-nondistributable-artifacts フラグを使用することでプッシュされるようになるが、ライセンスの条項に則る必要がある

各エンドポイントについて

* com.amazonaws.region.ecr.dkr: Docker Registry API 用
* com.amazonaws.region.ecr.api: ECR API 用
* S3
  * リソース `arn:aws:s3:::prod-region-starport-layer-bucket/*` の許可が必要



## Monitoring

[サービスクォータの可視化とアラームの設定](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/monitoring-quotas-alarms.html)

* クォータのメトリクスから CloudWatch アラームを作成し、通知などを行うことが可能


[Amazon ECR 使用状況メトリクス](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/monitoring-usage.html)

* API ごとに CallCount が収集される


[Amazon ECR リポジトリメトリクス](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/ecr-repository-metrics.html)

* リポジトリごとに RepositoryPullCount が収集される


[EventBridge](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/ecr-eventbridge.html)

* 以下の場合にイベントが発行される
  * イメージプッシュの完了時
  * ベーシックスキャン完了時
  * 拡張スキャンのリソース変更時
    * 新しいリポジトリの作成
    * スキャン頻度の変更
    * イメージの作成、削除
  * イメージ削除時


[CloudTrail](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/logging-using-cloudtrail.html)

記録される API について、いくつかの例が載っている。

* CreateRepository: リポジトリ作成時
* PutImage: イメージのプッシュ (プッシュ時は InitiateLayerUpload、UploadLayerPart、CompleteLayerUpload も記録される)
* BatchGetImage: イメージのプル (プル時は GetDownloadUrlForLayer も記録される)
* PolicyExecutionEvent: ライフサイクルポリシーの実行内容が記録される



## Service Quotas

[Amazon ECR のサービスクォータ](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/service-quotas.html)



## Troubleshooting

[トラブルシューティング](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/troubleshooting.html)


[Amazon ECR 使用時の Docker コマンドのエラーのトラブルシューティング](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/common-errors-docker.html)

#### イメージプル時の "Filesystem Verification Failed" or "404: Image Not Found"

* ローカルディスクがいっぱいである
* ネットワークエラー。ECR の場合は更に S3 へのネットワークアクセスも必要

#### イメージプッシュ時の HTTP 403 Errors or "no basic auth credentials" Error

* 別のリージョンに対して認証している
* リポジトリへのアクセス権がない
* Token の有効期限切れ
* wincred 認証情報マネージャーのバグ


[Amazon ECR エラーメッセージのトラブルシューティング](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/common-errors.html)

#### HTTP 429: Too Many Requests or ThrottleException

スロットリングされているために発生しているエラー。エクスポネンシャルバックオフで再試行するのが有効。

#### HTTP 403: "User [arn] is not authorized to perform [operation]"

`aws ecr get-login` 時のエラー。IAM ユーザのポリシー設定を見直すこと。

#### HTTP 404: "Repository Does Not Exist" Error

リポジトリが存在しないため。ECR の仕様として先にリポジトリを作成しておく必要がある。



# ナレッジセンター

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

