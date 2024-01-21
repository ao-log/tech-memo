
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


[プライベートレジストリの認証](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/registry_auth.html)

* `aws ecr get-login-password` コマンドにより認証する。認証トークンの有効時間は 12 時間。
* `GetAuthorizationToken` API が実行されエンコードされたパスワードを含む base64 エンコード認可トークンを取得
* レジストリに対して認証する。AWS アカウント、リージョンごとに異なる

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
* リポジトリ作成テンプレート
* Scanning configuration: ベーシックスキャン or 拡張スキャンを設定


[プライベートレジストリの許可](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/registry-permissions.html)

* デフォルトでは自アカウントのクロスリージョンレプリケーションの権限はある。クロスアカウント時には設定が必要
* ecr:ReplicateImage: クロスアカウントレプリケーションの許可設定
* ecr:BatchImportUpstreamImage: Pull through cache のアクセス許可を付与
* ecr:CreateRepository: リポジトリ作成を許可


[プライベートレジストリのポリシーの例](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/registry-permissions-examples.html)

クロスアカウントレプリケーションの許可。かつ、prod-* で始まる名前のリポジトリのみ許可
```json
{
    "Version":"2012-10-17",
    "Statement":[
        {
            "Sid":"ReplicationAccessCrossAccount",
            "Effect":"Allow",
            "Principal":{
                "AWS":"arn:aws:iam::source_account_id:root"
            },
            "Action":[
                "ecr:CreateRepository",
                "ecr:ReplicateImage"
            ],
            "Resource": [
                "arn:aws:ecr:us-west-2:your_account_id:repository/prod-*"
            ]
        }
    ]
}
```


## リポジトリ

[リポジトリの作成](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/repository-create.html)

設定項目

* タグのイミュータビリティ: 有効にすると、同じタグを使用した後続イメージのプッシュによるイメージタグ上書きを防ぐ
* プッシュ時にスキャン: プッシュ時にスキャンするかどうか
* KMS 暗号化: 有効化した場合、AWS 管理キーもしくはカスタマーマスターキーを選ぶことができる


[リポジトリポリシー](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/repository-policies.html)

* IAM ポリシーもしくはリポジトリポリシーのどちらかで許可が設定されている場合は、許可される。どちらか片方で拒否が設定されている場合は、拒否される
* プル、プッシュを行う前にレジストリに対する `ecr:GetAuthorizationToken` が必要
* `ecr:GetAuthorizationToken` は IAM ポリシー側での許可が必要


[リポジトリ ポリシーの例](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/repository-policy-examples.html)

クロスアカウントのプッシュを許可
```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "AllowCrossAccountPush",
            "Effect": "Allow",
            "Principal": {
                "AWS": "arn:aws:iam::account-id:root"
            },
            "Action": [
                "ecr:BatchCheckLayerAvailability",
                "ecr:CompleteLayerUpload",
                "ecr:InitiateLayerUpload",
                "ecr:PutImage",
                "ecr:UploadLayerPart"
            ]
        }
    ]
}
```

特定のユーザーにはプルのみ許可し、admin-user には全許可
```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "AllowPull",
            "Effect": "Allow",
            "Principal": {
                "AWS": [
                    "arn:aws:iam::account-id:user/pull-user-1",
                    "arn:aws:iam::account-id:user/pull-user-2"
                ]
            },
            "Action": [
                "ecr:BatchGetImage",
                "ecr:GetDownloadUrlForLayer"
            ]
        },
        {
            "Sid": "AllowAll",
            "Effect": "Allow",
            "Principal": {
                "AWS": "arn:aws:iam::account-id:user/admin-user"
            },
            "Action": [
                "ecr:*"
            ]
        }
    ]
}
```

サービスに対する許可の例
```json
{
   "Version":"2012-10-17",
   "Statement":[
      {
         "Sid":"CodeBuildAccess",
         "Effect":"Allow",
         "Principal":{
            "Service":"codebuild.amazonaws.com"
         },
         "Action":[
            "ecr:BatchGetImage",
            "ecr:GetDownloadUrlForLayer"
         ],
         "Condition":{
            "ArnLike":{
               "aws:SourceArn":"arn:aws:codebuild:region:123456789012:project/project-name"
            },
            "StringEquals":{
               "aws:SourceAccount":"123456789012"
            }
         }
      }
   ]
}
```


## イメージ

#### Push image

[イメージのプッシュ](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/docker-push-ecr-image.html)

```shell
docker tag e9ae3c220b23 aws_account_id.dkr.ecr.region.amazonaws.com/my-web-app
```
このようにイメージタグを省略した場合は latest とみなされる。

* レジストリへの認証を行うため `ecr:GetAuthorizationToken` の許可が必要


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

* Lambda ではサポートされない
* リモートパブリックレジストリ内のリポジトリのキャッシュをサポート
* キャッシュイメージの URI からプルすると、リモートリポジトリを 24 時間に 1 回までチェックしキャッシュされたイメージが最新バージョンか確認
* アップストリームからイメージをプルできない状況では、キャッシュが利用される
* マルチアーキテクチャイメージをプルすると、マニフェストリストとマニフェストリストで参照されている各イメージがプルされる。特定のアーキテクチャのみをプルする場合、アーキテクチャに関連付けられたイメージダイジェストまたはタグを使用してイメージをプルすること
* サービスにリンクされた IAM ロールを使用。キャッシュされたイメージのリポジトリを作成し、キャッシュされたイメージをプッシュするために必要なアクセス許可を提供
* レジストリポリシーにより権限を細かく制御可能
* タグのイミュータビリティを手動でオンにした場合、キャッシュされたイメージを更新できない場合がある
* 初回のプルではインターネットへのアウトバウンドの疎通性が必要
* 必要な IAM 許可
  * プライベートレジストリへの認証
  * イメージのプッシュとプルに必要となる Amazon ECR API のアクセス許可
  * 更に以下の許可が必要
    * `ecr:CreatePullThroughCacheRule`: プルスルーキャッシュルールを作成するアクセス許可
      * アイデンティティベースの IAM エンティティで許可
    * `ecr:BatchImportUpstreamImage`: 外部イメージを取得し、プライベートレジストリにインポートするアクセス許可
      * 以下のいずれかで許可
        * レジストリポリシー
        * アイデンティティベース
        * リポジトリポリシー
    * `ecr:CreateRepository`: リポジトリを作成するアクセス許可
      * 以下のいずれかで許可
        * レジストリポリシー
        * アイデンティティベース


[プルスルーキャッシュルールの作成](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/pull-through-cache-creating-rule.html)

* 認証情報が必要な場合は、Secrets Manager にシークレットを格納し、ARN を指定する


[アップストリームリポジトリの認証情報を AWS Secrets Manager シークレットに保存する](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/pull-through-cache-creating-secret.html)

* シークレット名は `ecr-pullthroughcache/` をプレフィックスとする必要がある
* コンテナレジストリごとに手順が用意されている


#### Delete image

[イメージの削除](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/delete_image.html)

* `batch-delete-image` により削除する
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



[プライベートイメージレプリケーションの設定](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/registry-settings-configure.html)

* リージョンごとに個別に設定が必要

プレフィックス prod のイメージのみ、クロスリージョンレプリケーションする場合のレプリケーションルールの例
```json
{
	"rules": [{
		"destinations": [{
			"region": "us-west-1",
			"registryId": "111122223333"
		}],
		"repositoryFilters": [{
			"filter": "prod",
			"filterType": "PREFIX_MATCH"
		}]
	}]
}
```

クロスアカウントに対するレプリケーションルールの例。更に対象アカウント側でレジストリポリシーによる許可が必要
```json
{
    "rules": [
        {
            "destinations": [
                {
                    "region": "us-west-2",
                    "registryId": "444455556666"
                }
            ]
        }
    ]
}
```


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
  * 最初に有効化した際は過去 30 日分のイメージがスキャン対象になる。より古いイメージは `SCAN_ELIGIBILITY_EXPIRED` のステータスになる
  * フィルター
    * 継続スキャン、プッシュ時にスキャン、それぞれについて別のフィルターを設定できる
    * 一致しないリポジトリは、スキャンが無効に設定される
  * 継続スキャン、プッシュ時にスキャンの設定が可能
    * 継続スキャンではスキャン期間を設定可能。デフォルトは Lifetime だが、180 day, 30 day に設定することも可能
  * 手動スキャンはできない
  * スキャン結果はイベントとして発行される
  * Amazon Inspector サービスにリンクされた IAM ロールが使用される。拡張スキャン有効時に Inspector によって自動的に作成される
  + イベント
    * ECR Scan Resource Change: 新しいリポジトリの作成、リポジトリのスキャン頻度の変更、または拡張スキャンがオンになっているリポジトリでのイメージの作成または削除
    * Inspector2 Scan: 初期スキャン完了時
    * Inspector2 Finding: スキャン結果の作成、更新、クローズ時
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

* コンテナエージェントにて許可が必要なポリシー
  * `ecr:BatchGetImage`
  * `ecr:GetDownloadUrlForLayer`
  * `ecr:GetAuthorizationToken`
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


[レプリケーション用の Amazon ECR サービスにリンクされたロール](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/slr-replication.html)

信頼ポリシー
* replication.ecr.amazonaws.com

ポリシー
```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "ecr:CreateRepository",
                "ecr:ReplicateImage"
            ],
            "Resource": "*"
        }
    ]
}
```


[プルスルーキャッシュの Amazon ECR サービスにリンクされたロール](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/slr-pullthroughcache.html)

信頼ポリシー
* pullthroughcache.ecr.amazonaws.com

ポリシー
```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "ECR",
            "Effect": "Allow",
            "Action": [
                "ecr:GetAuthorizationToken",
                "ecr:BatchCheckLayerAvailability",
                "ecr:InitiateLayerUpload",
                "ecr:UploadLayerPart",
                "ecr:CompleteLayerUpload",
                "ecr:PutImage"
            ],
            "Resource": "*"
        },
        {
            "Sid": "SecretsManager",
            "Effect": "Allow",
            "Action": [
                "secretsmanager:GetSecretValue"
            ],
            "Resource": "arn:aws:secretsmanager:*:*:secret:ecr-pullthroughcache/*",
            "Condition": {
                "StringEquals": {
                    "aws:ResourceAccount": "${aws:PrincipalAccount}"
                }
            }
        }
    ]
}
```


[IAM ポリシーの例](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/security_iam_id-based-policy-examples.html)

* 特定のリポジトリのみアクションを許可するなど、いくつかのポリシー例が載っている

us-east-1 の my-repo へのアクセスを許可する例
```json
{
   "Version":"2012-10-17",
   "Statement":[
      {
         "Sid":"ListImagesInRepository",
         "Effect":"Allow",
         "Action":[
            "ecr:ListImages"
         ],
         "Resource":"arn:aws:ecr:us-east-1:123456789012:repository/my-repo"
      },
      {
         "Sid":"GetAuthorizationToken",
         "Effect":"Allow",
         "Action":[
            "ecr:GetAuthorizationToken"
         ],
         "Resource":"*"
      },
      {
         "Sid":"ManageRepositoryContents",
         "Effect":"Allow",
         "Action":[
                "ecr:BatchCheckLayerAvailability",
                "ecr:GetDownloadUrlForLayer",
                "ecr:GetRepositoryPolicy",
                "ecr:DescribeRepositories",
                "ecr:ListImages",
                "ecr:DescribeImages",
                "ecr:BatchGetImage",
                "ecr:InitiateLayerUpload",
                "ecr:UploadLayerPart",
                "ecr:CompleteLayerUpload",
                "ecr:PutImage"
         ],
         "Resource":"arn:aws:ecr:us-east-1:123456789012:repository/my-repo"
      }
   ]
}
```


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
* クロスリージョンは未サポート
* ECR Public は未サポート
* プルスルーキャッシュを使用する場合は、初回のプル時のみインターネットへのアウトバウンドの疎通性が必要
* Windows イメージの場合、ライセンスによって配布が制限されているアーティファクトが含まれており、通常は当該箇所はプッシュされない。Docker デーモンの `--allow-nondistributable-artifacts` フラグを使用することでプッシュされるようになるが、ライセンスの条項に則る必要がある
  * イメージサイズが大きいため、プッシュ時間および料金がかかる点に注意が必要。例えば `mcr.microsoft.com/windows/servercore` イメージは圧縮された状態でも約 1.7 GB となる

各エンドポイントについて

* com.amazonaws.region.ecr.dkr: Docker Registry API 用
* com.amazonaws.region.ecr.api: ECR API 用
* S3
  * リソース `arn:aws:s3:::prod-region-starport-layer-bucket/*` の許可が必要

ECR API エンドポイントポリシーの例。特定のロールに対するレジストリへの認証、イメージのプルのみ許可
```json
{
	"Statement": [{
		"Sid": "AllowPull",
		"Principal": {
			"AWS": "arn:aws:iam::1234567890:role/role_name"
		},
		"Action": [
			"ecr:BatchGetImage",
			"ecr:GetDownloadUrlForLayer",
      "ecr:GetAuthorizationToken"
		],
		"Effect": "Allow",
		"Resource": "*"
	}]
}
```



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

* パフォーマンスの最適化
  * Docker 1.10 以降を使用
  * 小さなイメージ
  * 変更の少ない依存関係を前に配置
  * Dockerfile ではコマンドチェーンを使用する。不要ファイルが残ることを避けやすくなる
  * 最も近いリージョンのエンドポイントを使用する


[Amazon ECR 使用時の Docker コマンドのエラーのトラブルシューティング](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/common-errors-docker.html)

#### イメージプル時の "Filesystem Verification Failed" or "404: Image Not Found"

* エラーメッセージ
  * `Filesystem Verification Failed`: Docker 1.9 以上使用時
  * `404: Image Not Found`: Docler 1.9 未満使用時
* 原理
  * SHA-1 ハッシュが ECR 側で計算された内容と異なるために発生
* 発生理由
  * ローカルディスクがいっぱいである
  * ネットワークエラー。ECR の場合は更に S3 へのネットワークアクセスも必要

#### Amazon ECR からイメージをプルする際のエラー: "ファイルシステムレイヤーの検証に失敗しました"

* エラーメッセージ
  * `Filesystem Layer Verification Failed`
* 原理
  * イメージの 1 つ以上のレイヤーがダウンロードに失敗している
* 発生理由
  * 古いバージョンの Docker 使用時に稀に発生
  * クライアントでネットワークエラーまたはディスクエラーが発生

#### イメージプッシュ時の HTTP 403 Errors or "no basic auth credentials" Error

* エラーメッセージ
  * `HTTP 403 Errors`
  * `no basic auth credentials`
* 原理
  * アクセス権限がない
* 発生理由
  * 別のリージョンに対して認証している
  * リポジトリへのアクセス権がない
  * Token の有効期限切れ
  * wincred 認証情報マネージャーのバグ


[Amazon ECR エラーメッセージのトラブルシューティング](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/common-errors.html)

#### HTTP 429: Too Many Requests or ThrottleException

* エラーメッセージ
  * `HTTP 429: Too Many Requests`
  * `ThrottleException`
* 原理
  * スロットリング
  * 例えば `GetAuthorizationToken` アクションのスロットリングは 20 TPS (トランザクション/秒) 。最大 200 TPS のバースト
* 対策
  * エクスポネンシャルバックオフで再試行

#### HTTP 403: "User [arn] is not authorized to perform [operation]"

* エラーメッセージ
  * `HTTP 403: "User [arn] is not authorized to perform [operation]"`
* 原理
  * 権限不足
* 対策
  * IAM ユーザのポリシー設定を見直すこと

#### HTTP 404: "Repository Does Not Exist" Error

* エラーメッセージ
  * `HTTP 404: "Repository Does Not Exist" error`
* 原理
  * イメージが存在しない
* 対策
  * リポジトリを作成しておく必要がある

#### エラー: Cannot perform an interactive login from a non TTY device (TTY 以外のデバイスから対話型ログインを実行できません)

* エラーメッセージ
  * `Error: Cannot perform an interactive login from a non TTY device`
* 対策
  * AWS CLI バージョン 2 を使用していることを確認


[プルスルーキャッシュに関する問題のトラブルシューティング](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/error-pullthroughcache.html)

リポジトリが存在しない

* エラーメッセージ
  * `Error response from daemon: repository 111122223333.dkr.ecr.us-east-1.amazonaws.com/ecr-public/amazonlinux/amazonlinux not found: name unknown: The repository with name 'ecr-public/amazonlinux/amazonlinux' does not exist in the registry with id '111122223333'`
* 原理
  * リポジトリが存在しない
  * `ecr:CreateRepository` が許可されていない
* 対策
  * URI が正しいか確認する
  * 必要となる IAM アクセス許可がアップストリームイメージをプルする IAM プリンシパルに付与されていること
  * リポジトリが存在することを確認

リクエストされたイメージが見つからない

* エラーメッセージ
  * `Error response from daemon: manifest for 111122223333.dkr.ecr.us-east-1.amazonaws.com/ecr-public/amazonlinux/amazonlinux:latest not found: manifest unknown: Requested image not found`
* 原理
  * アップストリームレジストリにイメージが存在しない
  * アップストリームイメージをプルする IAM プリンシパルに `ecr:BatchImportUpstreamImage`` アクセス許可が付与されていない
  * リポジトリが存在している（存在していない場合はリポジトリがない、というエラーになるはずのため）
* 対策
  * アップストリームのイメージ、タグ指定が正しいか確認する
  * 必要となる IAM アクセス許可がアップストリームイメージをプルする IAM プリンシパルに付与されていること

Docker Hub リポジトリからプルするときに 403 Forbidden

* エラーメッセージ
  * `Error response from daemon: failed to resolve reference "111122223333.dkr.ecr.us-west-2.amazonaws.com/docker-hub/amazonlinux:2023": pulling from host 111122223333.dkr.ecr.us-west-2.amazonaws.com failed with status code [manifests 2023]: 403 Forbidden`
* 原理
  * 使用する URI に /library/ が含まれていない
* 対策
  * 使用する URI に /library/ を含める


[イメージスキャニングの問題のトラブルシューティング](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/userguide/image-scanning-troubleshooting.html)

UnsupportedImageError

* エラーメッセージ
  * `UnsupportedImageError`
* 原理
  * ベーシックスキャンのサポート対象は以下の通り
    * Amazon Linux、Amazon Linux 2、Debian、Ubuntu、CentOS、Oracle Linux、Alpine、RHEL Linux

UNDEFINED 深刻度が返される

* 原理
  * この脆弱性に、CVE ソースによって優先度が割り当てられていなかった
  * この脆弱性に、Amazon ECR が認識しない優先度が割り当てられていた
* 対策
  * ソースから直接 CVE を表示

SCAN_ELIGIBILITY_EXPIRED

* 原理
  * 初回に拡張スキャンを有効化した際のスキャン対象は過去 30 日間にプッシュされたイメージのみ。より古いイメージはスキャン対象外
  * 継続スキャンで設定された日数よりも古いイメージ



# API

[Amazon Elastic Container Registry のアクション、リソース、および条件キー](https://docs.aws.amazon.com/ja_jp/service-authorization/latest/reference/list_amazonelasticcontainerregistry.html)

* BatchCheckLayerAvailability:	指定されたレジストリとリポジトリの複数のイメージレイヤーの可用性を確認する許可を付与
* BatchDeleteImage: 指定したリポジトリ内の指定したイメージのリストを削除する許可を付与
* BatchGetImage: 指定したリポジトリ内の指定したイメージの詳細情報を取得する許可を付与
* BatchGetRepositoryScanningConfiguration: リポジトリのリストのリポジトリスキャン設定を取得するアクセス許可を付与
* BatchImportUpstreamImage [アクセス許可のみ]: アップストリームレジストリからイメージを取得し、プライベートレジストリにインポートするアクセス許可を付与
* CompleteLayerUpload: 指定されたレジストリ、リポジトリ名、およびアップロード ID のイメージレイヤーアップロードが完了したことを Amazon ECR に通知する許可を付与
* CreatePullThroughCacheRule: 新しいプルスルーキャッシュルールを作成するアクセス許可を付与
* CreateRepository: イメージレポジトリを作成する許可を付与
* CreateRepositoryCreationTemplate: リポジトリ作成テンプレートを作成するためのアクセス許可を付与
* DeleteLifecyclePolicy:	指定したライフサイクルポリシーを削除する許可を付与
* DeletePullThroughCacheRule	プルスルーキャッシュルールを削除するアクセス許可を付与	書き込み			
* DeleteRegistryPolicy: レジストリポリシーを削除するアクセス許可を付与
* DeleteRepository: 既存のイメージリポジトリを削除する許可を付与
* DeleteRepositoryCreationTemplate: リポジトリ作成テンプレートを削除するためのアクセス許可を付与
* DeleteRepositoryPolicy* 指定したリポジトリからリポジトリポリシーを削除する許可を付与
* DescribeImageReplicationStatus: レジストリ内のイメージに関するレプリケーションステータス (レプリケーションが失敗した場合の失敗の理由を含む) を取得する許可を付与
* DescribeImageScanFindings:	指定したイメージのイメージスキャンの結果を記述する許可を付与
* DescribeImages: リポジトリ内のイメージに関するメタデータ (例: イメージサイズ、イメージタグ、作成日) を取得する許可を付与
* DescribePullThroughCacheRules: プルスルーキャッシュルールを記述するアクセス許可を付与
* DescribeRegistry: レジストリ設定を記述するアクセス許可を付与
* DescribeRepositories: レジストリ内のイメージリポジトリを記述する許可を付与
* DescribeRepositoryCreationTemplate: リポジトリ作成テンプレートを記述するためのアクセス許可を付与
* GetAuthorizationToken: 指定したレジストリに対して有効なトークンを 12 時間取得する許可を付与
* GetDownloadUrlForLayer: イメージレイヤーに対応するダウンロード URL を取得する許可を付与
* GetLifecyclePolicy: 指定されたライフサイクルポリシーを取得する許可を付与
* GetLifecyclePolicyPreview: 指定されたライフサイクルポリシーのプレビューリクエストの結果を取得する許可を付与
* GetRegistryPolicy: レジストリポリシーを取得するアクセス許可を付与
* GetRegistryScanningConfiguration: レジストリスキャン設定を取得するアクセス許可を付与
* GetRepositoryPolicy: 指定したリポジトリのリポジトリポリシーを取得する許可を付与
* InitiateLayerUpload: イメージレイヤーのアップロードを予定していることを Amazon ECR に通知する許可を付与
* ListImages: 特定のリポジトリのすべてのイメージ ID を一覧表示する許可を付与
* ListTagsForResource: Amazon ECR リソースのタグを一覧表示する許可を付与
* PutImage: イメージに関連付けられたイメージマニフェストを作成または更新する許可を付与
* PutImageScanningConfiguration: リポジトリのイメージスキャン設定を更新する許可を付与
* PutImageTagMutability: リポジトリのイメージタグの変更可能性を更新する許可を付与
* PutLifecyclePolicy: ライフサイクルポリシーを作成または更新する許可を付与
* PutRegistryPolicy: レジストリポリシーを更新するアクセス許可を付与
* PutRegistryScanningConfiguration: レジストリスキャン設定を更新するアクセス許可を付与
* PutReplicationConfiguration: レジストリのレプリケーション設定を更新する許可を付与
* ReplicateImage [アクセス許可のみ]: イメージをレプリケート先レジストリにレプリケートするアクセス許可を付与
* SetRepositoryPolicy: 指定したリポジトリにリポジトリポリシーを適用してアクセス権限を制御する許可を付与
* StartImageScan: イメージスキャンを開始する許可を付与
* StartLifecyclePolicyPreview: 指定したライフサイクルポリシーのプレビューを開始する許可を付与
* TagResource: Amazon ECR リソースにタグを付けるアクセス許可を付与
* UntagResource: Amazon ECR リソースのタグを解除する許可を付与
* UpdatePullThroughCacheRule: プルスルーキャッシュルールを更新するためのアクセス許可を付与
* UploadLayerPart: イメージレイヤー部分を Amazon ECR にアップロードする許可を付与
* ValidatePullThroughCacheRule: プルスルーキャッシュルールを検証するためのアクセス許可を付与



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

