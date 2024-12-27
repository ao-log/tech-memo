
[Amazon EKS とは何ですか?](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/what-is-eks.html)


[Amazon EKS アーキテクチャ](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/eks-architecture.html)

* コントロールプレーンがマネージドになっている
* 2 つ以上の API ノード、3 つの etcd ノードから構成されている


[Kubernetes の概念](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/kubernetes-concepts.html)

* コントロールプレーン
  * kube-apiserver
  * etcd キーバリューストア
  * kube-scheduler
  * kube-controller-manager
  * cloud-controller-manager
* ワーカーノード(データプレーン)
  * kubelet
  * コンテナランタイム
  * kube-proxy



## Amazon EKS の使用開始

[Amazon EKS の開始方法 – eksctl](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/getting-started-eksctl.html)

* kubectl, eksctl はバイナリをダウンロードするだけで使用できる
* eksctl はコマンドラインオプションだけでなく、yaml ファイルで設定値を渡すことも可能


[Amazon EKS の開始方法 – AWS Management Console と AWS CLI](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/getting-started-console.html)

コンソールから対応する場合は、以下のステップで EKS クラスターを作成している。

* CloudFormation テンプレートにより VPC とサブネットを作成
* クラスター用の IAM ロールの作成。`AmazonEKSClusterPolicy` などをアタッチ
* マネジメントコンソール上で EKS クラスターを作成。次の項目を設定。
  * 上記で作成した IAM ロール、サブネット
  * クラスターエンドポイントのアクセス(パブリック、プライベート、パブリックおよびプライベート)
  * ログ記録の構成
* `aws eks update-kubeconfig ...` による kubeconfig ファイルの作成(`~/.kube/config` に設定が追加される)
* ノードグループの作成



## クラスター

[Amazon EKS クラスターの作成](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/create-cluster.html)

以下のような設定項目がある。

* Kubernetes バージョン
* クラスターサービスロール
* シークレット暗号化
* ネットワーキング
  * VPC
  * サブネット
  * セキュリティグループ
  * Kubernetes サービスの IP アドレス範囲
  * クラスターエンドポイントアクセス
* オブザーバビリティ
* アドオン


[クラスターのインサイト](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/cluster-insights.html)

* バージョンアップグレードの準備状況に関するインサイトを表示


[Amazon EKS クラスターの Kubernetes バージョンの更新](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/update-cluster.html)

* 事前に更新内容を確認するとともに、テスト用の環境などでテストしておくことを推奨
* 最大で 5 つの使用可能な IP アドレスが必要
* コントロールプレーンをアップデートしてから、データプレーン側をアップデートする順番で対応する
* アドオンについても必要に応じて更新が必要


[クラスターの削除](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/delete-cluster.html)

* 事前に EXTERNAL-IP 値のある Service オブジェクトを削除しておく必要がある
* マネジメントコンソールから対応する場合、クラスターの削除の前に Fargate プロファイル、ノードグループを削除する必要あり


[Amazon EKS クラスターエンドポイントアクセスコントロール](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/cluster-endpoint.html)

* Kubernetes API サーバーエンドポイントの公開範囲を設定可能
* デフォルトではパブリックになっている。この場合でも IAM, RBAC により保護されている。また、許可した接続元 IP アドレスのみアクセスできるよう設定することも可能。VPC 内のノードからもアクセスができるようにするには、その IP アドレスも許可対象に加える、もしくはプライベートエンドポイント有効にする必要がある
* プライベートに設定することで、API エンドポイントにアクセス可能な接続元を VPC 内に留めることができる


[既存のクラスター上でシークレット暗号化を有効にする](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/enable-kms.html)

* Secrets が KMS キーによって暗号化されるようになる
* KMS キーを削除すると、復旧手段がない


[Amazon EKS クラスター の Windows サポートの有効化](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/windows-support.html)


[プライベートクラスターの要件](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/private-clusters.html)

インターネットへのアウトバウントの疎通性のないプライベートクラスターを構築可能。
以下の要件を満たす必要がある。

* クラスターエンドポイントのプライベート設定
* 必要に応じて VPC エンドポイントを作成
* イメージレジストリは ECR か VPC 内のレジストリが必要
* セルフマネージド型ノードの場合はブートストラップ引数のオプション追加が必要


[Amazon EKS Kubernetes のバージョン](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/kubernetes-versions.html)

* Kubernetes のマイナーバージョンは平均して 4 ヶ月ごとにリリースされている
* 各マイナーバージョンは最初にリリースされてから約 14 ヶ月間、標準サポート対象となる
* 標準サポート終了日を過ぎると自動的に延長サポートとなる。期間は約 12 ヶ月。追加料金が発生する。延長サポートが失効する日になると、自動的に現在サポートされている最も古い延長バージョンに自動アップデートされる
* EKS は 4 つ以上のバージョンを標準サポートとするように努めている


[Amazon EKS のプラットフォームバージョン](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/platform-versions.html)

* EKS クラスターバージョンはのコントロールプレーン側のバージョンが該当
* 各 Kubernetes バージョンごとに Amazon EKS プラットフォームバージョンがある。eks.1 から始まり、1 ずつインクリメントされていく
* EKS プラットフォームバージョンの更新は自動的に行われ、サービス中断の影響は発生しない。また、重大な変更は発生しない


[Autoscaling](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/autoscaling.html)

2 つの製品をサポート
* Karpenter
* Cluster Autoscaler



## アクセスを管理する

[アクセスを管理する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/cluster-auth.html)


### Kubernetes API

[IAM ユーザーおよびロールに KubernetesAPIs へのアクセス権を付与する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/grant-k8s-access.html)

* IAM もしくは OIDC プロバイダーにより認証可能
* AWS IAM Authenticator はコントロールプレーンにインストールされている。アクセスエントリもしくは aws-auth ConfigMap により認証できる
* クラスター認証モードは 3 種類
  * アクセスエントリのみ
  * アクセスエントリと ConfigMap の両方
  * ConfigMap のみ


[EKS アクセスエントリを使用して Kubernetes へのアクセスを IAM ユーザーに許可する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/access-entries.html)

* 次のタイプがある
  * STANDARD
  * EC2_LINUX
  * EC2_WINDOWS
  * FARGATE_LINUX
* アクセスポリシーを関連づけることができる。アクセスポリシーにより Kubernetes オブジェクトへの権限を制御


[アクセスポリシーとアクセスエントリとの関連付けおよび関連付け解除](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/access-policies.html)

* 次のアクセスポリシーがある
  * AmazonEKSAdminPolicy
  * AmazonEKSClusterAdminPolicy
  * AmazonEKSAdminViewPolicy
  * AmazonEKSEditPolicy
  * AmazonEKSViewPolicy


[ConfigMap を使用して Kubernetes へのアクセスを IAM ユーザーに許可する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/auth-configmap.html)

aws-auth ConfigMap は deprecated になっている。


[外部 OIDC プロバイダーを使用して Kubernetes へのアクセスをユーザーに許可する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/authenticate-oidc-identity-provider.html)


[kubeconfig ファイルを作成して kubectl を EKS クラスターに接続する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/create-kubeconfig.html)

* クラスター認証では `aws eks get-token` を使用している
* `aws eks update-kubeconfig --region region-code --name my-cluster` により kubeconfig を更新できる


### AWS API

[Kubernetes ワークロードに Kubernetes サービスアカウントを使用して AWS へのアクセスを許可する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/service-accounts.html)

* AWS サービスへのアクセス許可の方法は 2 種類ある
  * サービスアカウントの IAM ロール(IRSA)
  * EKS Pod Identity


#### EKS Pod Identity

[EKS Pod Identity がポッドに AWS サービスへのアクセス権を付与する方法を学ぶ](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/pod-identities.html)

* OIDC プロバイダーを必要としない
* 以下のステップでセットアップ
  * Amazon EKS Pod Identity エージェントのセットアップ
  * IAM ロールを Kubernetes サービスアカウントに割り当てる
  * サービスアカウントを使用して AWS サービスにアクセスするように pods を設定する
  * サポートされる AWS SDK を使用する 
* 考慮事項
  * Fargate, Windows ではサポートされていない
  * アドオンでは IRSA のみ対応


[EKS Pod Identity の詳細を理解する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/pod-id-how-it-works.html)

* EKS Pod Identity の関連付けを持つサービスアカウントを使用する新しい Pod を起動すると、マニフェストに環境変数やボリュームのマウントが追加される


[Amazon EKS Pod Identity エージェントのセットアップ](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/pod-id-agent-setup.html)

* ノードの IAM ロールには `eks-auth:AssumeRoleForPodIdentity` の許可が必要
* EKS Pod Identity のアドオンを追加することで導入できる。DaemonSet オブジェクトがデプロイされる


[IAM ロールを Kubernetes サービスアカウントに割り当てる](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/pod-id-association.html)

* サービスアカウントは以下のような内容で作成する
```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: my-service-account
  namespace: default
```
* IAM ロールの信頼ポリシーは以下のように設定する
```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "AllowEksAuthToAssumeRoleForPodIdentity",
            "Effect": "Allow",
            "Principal": {
                "Service": "pods.eks.amazonaws.com"
            },
            "Action": [
                "sts:AssumeRole",
                "sts:TagSession"
            ]
        }
    ]
}
```

* 関連づけ設定。サービスアカウント、IAM ロールを関連づけ、namespace を指定する
```shell
aws eks create-pod-identity-association ¥
    --cluster-name my-cluster ¥
    --role-arn arn:aws:iam::111122223333:role/my-role ¥
    --namespace default ¥
    --service-account my-service-account
```


[サービスアカウントを使用して AWS サービスにアクセスするように pods を設定する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/pod-id-configure-pods.html)

* Pod の .spec にてサービスアカウントを指定する
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-app
spec:
  selector:
    matchLabels:
      app: my-app
  template:
    metadata:
      labels:
        app: my-app
    spec:
      serviceAccountName: my-service-account
      containers:
      - name: my-app
        image: public.ecr.aws/nginx/nginx:X.XX
```
* 環境変数 `AWS_CONTAINER_AUTHORIZATION_TOKEN_FILE` にサービスアカウントトークンのファイルパスが設定される


[タグに基づいて AWS リソースへの pods アクセス権を付与する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/pod-id-abac.html)


[サポートされる AWS SDK を使用する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/pod-id-minimum-sdk.html)

* EKS Pod Identity を使用するには、SDK のバージョンが所定バージョン以降を使用すること


[EKS Pod Identity が必要とする信頼ポリシーを使用して IAM ロールを作成する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/pod-id-role.html)


#### IRSA

[サービスアカウントの IAM ロール](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/iam-roles-for-service-accounts.html)

* 以下のステップでセットアップ
  * クラスターの IAM OIDC プロバイダーを作成
  * IAM ロールを Kubernetes サービスアカウントに割り当てる
  * サービスアカウントを使用して AWS サービスにアクセスするように pods を設定する
  * サポートされる AWS SDK を使用する 
* 参考情報
  * 動作原理に関する参考記事 [詳解: IAM Roles for Service Accounts](https://aws.amazon.com/jp/blogs/news/diving-into-iam-roles-for-service-accounts/)
  * [上記記事のまとめ](../..//Infra/Kubernetes/Document/Security/ServiceAccount.md)


[クラスターの IAM OIDC プロバイダーを作成する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/enable-iam-roles-for-service-accounts.html)

* eksctl などにより OIDC プロバイダーを作成し、EKS クラスターに関連づけることができる


[Kubernetes サービスアカウントへの IAM ロールの割り当て](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/associate-service-account-role.html)

* サービスアカウントは以下のような内容で作成する
```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: my-service-account
  namespace: default
```
* 更にサービスアカウントにアノテーションの付与が必要
```shell
kubectl annotate serviceaccount -n $namespace $service_account eks.amazonaws.com/role-arn=arn:aws:iam::$account_id:role/my-role
```
* IAM ロールの信頼ポリシーは以下のように設定する
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Federated": "arn:aws:iam::$account_id:oidc-provider/$oidc_provider"
      },
      "Action": "sts:AssumeRoleWithWebIdentity",
      "Condition": {
        "StringEquals": {
          "$oidc_provider:aud": "sts.amazonaws.com",
          "$oidc_provider:sub": "system:serviceaccount:$namespace:$service_account"
        }
      }
    }
  ]
}
```


[Kubernetes サービスアカウントを使用するように Pods を設定するには](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/pod-configuration.html)

* Pod の .spec にてサービスアカウントを指定する
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-app
spec:
  selector:
    matchLabels:
      app: my-app
  template:
    metadata:
      labels:
        app: my-app
    spec:
      serviceAccountName: my-service-account
      containers:
      - name: my-app
        image: public.ecr.aws/nginx/nginx:X.XX
```
* 環境変数 `AWS_WEB_IDENTITY_TOKEN_FILE` にサービスアカウントトークンのファイルパスが設定される
  * kubelet が Pod に代わってトークンをリクエストし格納する。デフォルトでは有効期限の 80 % を超えている場合もしくは 24 時間を超えているとトークンを更新する
  * クラスターの Amazon EKS Pod Identity ウェブフックがアノテーションが付与されたサービスアカウントの Pod をモニタリングしている


[サービスアカウントの AWS Security Token Service エンドポイントを設定する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/configure-sts-endpoint.html)


[クロスアカウントの IAM アクセス許可](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/cross-account-access.html)


[サポートされる AWS SDK の使用](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/iam-roles-for-service-accounts-minimum-sdk.html)

* IRSA を使用するには、SDK のバージョンが所定バージョン以降を使用すること


[OIDC トークンを検証するための署名キーを取得する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/irsa-fetch-keys.html)



## ノード

[Amazon EKS ノード](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/eks-compute.html)

* ノードに含まれているもの
  * コンテナランタイム
  * kubelet
  * kube-proxy
* ノードグループの種類
  * マネージド型
  * セルフマネージド型
  * Fargate


### マネージド型ノードグループ

[マネージド型ノードグループ](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/managed-node-groups.html)

* ASG の一部としてプロビジョニングされる
* Cluster Autoscaler を使用するようにタグ付けされる
* カスタム起動テンプレートを使用可能
* インスタンスにラベルが付与される。プレフィックス `eks.amazonaws.com` が付与されている
* 終了時に Kubernetes API を使用して drain する


[マネージド型ノードグループの作成](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/create-managed-node-group.html)

eksctl もしくはマネジメントコンソールより作成可能。

eksctl の場合のコマンド例。

```shell
eksctl create nodegroup \
  --cluster my-cluster \
  --region region-code \
  --name my-mng \
  --node-ami-family ami-family \
  --node-type m5.large \
  --nodes 3 \
  --nodes-min 2 \
  --nodes-max 4 \
  --ssh-access \
  --ssh-public-key my-key
```

マネジメントコンソールでは次の設定項目がある。

* ノードグループ名
* インスタンスロール
* 起動テンプレート
* Kubernetes labels
* Kubernetes taints
* タグ
* AMI のタイプ
* キャパシティタイプ
* インスタンスタイプ
* ディスクサイズ
* Auto Scaling の台数設定(min, max, desired)
* サブネット
* SSH キーペア


[マネージド型ノードグループの更新](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/update-managed-node-group.html)

* eksctl もしくはマネジメントコンソール上の操作でノードグループを更新可能
* 次のコマンドにより最新の AMI に更新できる
```shell
eksctl upgrade nodegroup \
  --name=node-group-name \
  --cluster=my-cluster \
  --region=region-code
```


[マネージド型ノードの更新動作](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/managed-node-update-behavior.html)

マネージド型ノードグループを更新した場合の動作の流れが書かれている。

1. 新しい起動テンプレートを作成
1. 新しい起動テンプレートを使用するように Auto Scaling グループを更新
1. Auto Scaling グループの最大サイズ、必要なサイズを最大 2 倍に増やす
1. 最新の AMI ID でラベル付をされていないノードグループ内のノードに `eks.amazonaws.com/nodegroup=unschedulable:NoSchedule` の taint
1. ノードグループ内のノードをランダムに選択し、Pod をドレイン
1. Auto Scaling グループの最大サイズ、希望サイズを 1 ずつ減らしていき更新前の台数まで繰り返す


[マネージド型ノードグループでのノードテイント](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/node-taints-managed-node-groups.html)

* ノードグループのプロパティにて taint を設定できる。key, value, effect を設定できる


[起動テンプレートを使用したマネージドノードのカスタマイズ](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/launch-templates.html)

* bootstrap の引数を与えたい場合のユーザーデータ設定例
```shell
MIME-Version: 1.0
Content-Type: multipart/mixed; boundary="==MYBOUNDARY=="

--==MYBOUNDARY==
Content-Type: text/x-shellscript; charset="us-ascii"

#!/bin/bash
set -ex
/etc/eks/bootstrap.sh my-cluster \
  --b64-cluster-ca certificate-authority \
  --apiserver-endpoint api-server-endpoint \
  --dns-cluster-ip service-cidr.10 \
  --kubelet-extra-args '--max-pods=my-max-pods-value' \
  --use-max-pods false

--==MYBOUNDARY==--
```


[マネージド型ノードグループの削除](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/delete-managed-node-group.html)


### セルフマネージド型ノード

[セルフマネージド型ノード](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/worker.html)

* ノードには以下のタグが必要
  * key: `kubernetes.io/cluster/クラスター名`
  * value: `owned`


[セルフマネージド型 Amazon Linux ノードの起動](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/launch-workers.html)

eksctl の場合のコマンド例。

```shell
eksctl create nodegroup \
  --cluster my-cluster \
  --name al-nodes \
  --node-type t3.medium \
  --nodes 3 \
  --nodes-min 1 \
  --nodes-max 4 \
  --ssh-access \
  --managed=false \
  --ssh-public-key my-key
```

* マネジメントコンソール上から作成する場合
  * CloudFormation テンプレートから作成できる。また、ノードがクラスターに参加するには、ConfigMap にインスタンスロールの設定が必要


[セルフマネージド型ノードの更新](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/update-workers.html)

* 2 つの更新方法がある
  * 新しいノードグループを作成したあと元のノードグループを削除
  * 既存のノードグループのスタックの AMI を更新


[新しいノードグループへの移行](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/migrate-stack.html)

**eksctl を使用する場合**

eksctl コマンドで新しいノードグループを削除したあと、古いノードグループを削除すれば良い。

**AWS CLI を使用する場合**

1. 新しいノードグループを作成。
1. 新旧双方のノードグループのセキュリティグループにおいて、互いに通信できるようにインバウンド通信を許可。
1. ConfigMap に新しいノードグループのインスタンスロールを設定。
1. 元のノードグループのノードに taint を付与。
1. クラスターから削除する各ノードを drain する。
1. 元のノードグループを削除
  1. セキュリティグループの新旧双方の許可設定を削除。
  1. 元のノードグループの CloudFormation スタックを削除。
  1. ConfigMap から元のノードグループのインスタンスロールを削除


[既存のセルフマネージド型ノードグループの更新](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/update-stack.html)


### AWS Fargate

[AWS Fargate](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/fargate.html)

いくつか制約がある。

* Fargate Profile に一致するように Pod の設定が必要
* DaemonSet はサポートされない
* 特権を持つコンテナはサポートされない
* プライベートサブネットでのみサポート
* EBS ボリュームはマウントできない
* Job が終了した後も Pod は残り続ける。自動削除するには `.spec.ttlSecondsAfterFinished` の設定が必要


[Amazon EKS を使用した AWS Fargate の使用開始](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/fargate-getting-started.html)

* Fargate で実行されている Pod は、クラスターのクラスターセキュリティグループを使用するように設定される
* Fargate Pod 実行ロール: ECR からのイメージプルや CloudWatch Logs へのログ送信に必要な権限
* Pod をスケジューリングするには Fargate Profile が必要
```shell
eksctl create fargateprofile \
    --cluster my-cluster \
    --name my-fargate-profile \
    --namespace my-kubernetes-namespace \
    --labels key=value
```
* CoreDNS を Fargate で稼働させることが可能
  * Fargate Profile を作成
  ```shell
  aws eks create-fargate-profile \
      --fargate-profile-name coredns \
      --cluster-name my-cluster \
      --pod-execution-role-arn arn:aws:iam::111122223333:role/AmazonEKSFargatePodExecutionRole \
      --selectors namespace=kube-system,labels={k8s-app=kube-dns} \
      --subnets subnet-0000000000000001 subnet-0000000000000002 subnet-0000000000000003
  ```
  * EC2 アノテーションを削除
  ```shell
  kubectl patch deployment coredns \
    -n kube-system \
    --type json \
    -p='[{"op": "remove", "path": "/spec/template/metadata/annotations/eks.amazonaws.com~1compute-type"}]'
  ```


[AWS Fargate プロファイル](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/fargate-profile.html)

* 1 つの Profile に最大 5 つのセレクターを設定可能
* 複数の Profile にマッチする場合は、Pod で明示的に指定することが可能。`eks.amazonaws.com/fargate-profile: my-fargate-profile`
* Fargate Profile のコンポーネント
  * Pod 実行ロール
  * サブネット。プライベートサブネットのみ対応している
  * セレクター。名前空間、ラベルを設定する
* 次のコマンドで Fargate Profile を作成可能。
```shell
eksctl create fargateprofile \
    --cluster my-cluster \
    --name my-fargate-profile \
    --namespace my-kubernetes-namespace \
    --labels key=value
```


[Fargate Pod の設定](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/fargate-pod-configuration.html)

* Kubernetes コンポーネント (kubelet、kube-proxy、および containerd) の各 Pod のメモリ予約に 256 MB を追加
* vCPU, メモリ量の組み合わせが決まっている。無指定時は最小(0.25 vCPU, 0.5 GB Memory)となる
* node を describe した時のサイズは Fargate Pod で使用可能なサイズを反映していない。Pod を describe した時の `.annotations.CapacityProvisioned` を参照すること
* エフェメラルストレージはデフォルトで 20 GiB。175 GiB まで増やすことが可能


[Fargate OS のパッチ適用](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/fargate-pod-patching.html)

* セキュリティパッチが自動適用される
* PDB により同時にダウンする Pod 数を制御可能


[Fargate メトリクス](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/monitoring-fargate-usage.html)

* デフォルトでは `AWS/Usage` 名前空間でアカウント内での Fargate のリソース使用量が取得される
* AWS Distro for OpenTelemetry (ADOT) により各メトリクスを取得可能


[Fargate ログ記録](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/fargate-logging.html)

* Fluent Bit コンテナをサイドカーとして稼働させる必要はない
* 名前空間 `aws-observability` 内の ConfigMap `aws-logging` により Fluent Bit の設定を行う 


### インスタンスタイプ

[Amazon EC2 インスタンスタイプを選択する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/choosing-instance-type.html)

* 一部のインスタンスタイプはサポートされていない
* インスタンスタイプごとに Pod 数の最大数が決まっている。Nitro タイプのインスタンスタイプではこの最大数を増やすことができる。


### EKS 最適化 AMI

[Amazon EKS 最適化 AMI](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/eks-optimized-amis.html)


[Amazon EKS 最適化 Amazon Linux AMI](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/eks-optimized-ami.html)

* 次のコンポーネントが含まれている
  * kubelet
  * AWS IAM Authenticator
  * containerd
* Amazon Linux 2023
  * セルフマネージド型もしくは起動テンプレートを使用している場合、以下の設定が必要
  ```yaml
  ---
  apiVersion: node.eks.aws/v1alpha1
  kind: NodeConfig
  spec:
    cluster:
      name: my-cluster
      apiServerEndpoint: https://example.com
      certificateAuthority: Y2VydGlmaWNhdGVBdXRob3JpdHk=
      cidr: 10.100.0.0/16
  ```
  * IMDSv2 を使用する必要がある。ホップカウントが 1 のため、コンテナからは使用不可。使用するためには 2 以上に設定するか、Amazon EKS Pod Identity による認証を行う


[Amazon EKS 最適化 Amazon Linux AMI ID の取得](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/retrieve-ami-id.html)

* SSM パラメータとして EKS 最適化 AMI の AMI ID を取得可能


[Amazon EKS 最適化 Amazon Linux AMI のビルドスクリプト](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/eks-ami-build-scripts.html)

* AMI のビルドスクリプトは公開されている。[awslabs/amazon-eks-ami](https://github.com/awslabs/amazon-eks-ami)



## ストレージ

[クラスターのためにアプリケーションデータを保存する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/storage.html)


[Amazon EBS を利用して Kubernetes ボリュームを保存する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/ebs-csi.html)

* Fargate Pod にはマウントできない
* OIDC プロバイダーがあることが前提
* 使用手順
  * IAM ロールとサービスアカウントを関連づける
  * アドオンなどを使用して CSI ドライバーをインストール
  * [サンプル](https://github.com/kubernetes-sigs/aws-ebs-csi-driver/tree/master/examples/kubernetes)

dynamic-provisioning のサンプル

```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: ebs-sc
provisioner: ebs.csi.aws.com
volumeBindingMode: WaitForFirstConsumer
```

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: ebs-claim
spec:
  accessModes:
    - ReadWriteOnce
  storageClassName: ebs-sc
  resources:
    requests:
      storage: 4Gi
```

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: app
spec:
  containers:
  - name: app
    image: centos
    command: ["/bin/sh"]
    args: ["-c", "while true; do echo $(date -u) >> /data/out.txt; sleep 5; done"]
    volumeMounts:
    - name: persistent-storage
      mountPath: /data
  volumes:
  - name: persistent-storage
    persistentVolumeClaim:
      claimName: ebs-claim
```


[Amazon EFS で伸縮自在なファイルシステムを保存する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/efs-csi.html)

* OIDC プロバイダーがあることが前提
* 使用手順
  * IAM ロールとサービスアカウントを関連づける
  * アドオンなどを使用して CSI ドライバーをインストール
  * [サンプル](https://github.com/kubernetes-sigs/aws-efs-csi-driver/blob/master/docs/README.md#examples)

dynamic-provisioning のサンプル

```yaml
kind: StorageClass
apiVersion: storage.k8s.io/v1
metadata:
  name: efs-sc
provisioner: efs.csi.aws.com
parameters:
  provisioningMode: efs-ap
  fileSystemId: fs-92107410
  directoryPerms: "700"
  gidRangeStart: "1000" # optional
  gidRangeEnd: "2000" # optional
  basePath: "/dynamic_provisioning" # optional
  subPathPattern: "${.PVC.namespace}/${.PVC.name}" # optional
  ensureUniqueDirectory: "true" # optional
  reuseAccessPoint: "false" # optional
```

```yaml
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: efs-claim
spec:
  accessModes:
    - ReadWriteMany
  storageClassName: efs-sc
  resources:
    requests:
      storage: 5Gi
---
apiVersion: v1
kind: Pod
metadata:
  name: efs-app
spec:
  containers:
    - name: app
      image: centos
      command: ["/bin/sh"]
      args: ["-c", "while true; do echo $(date -u) >> /data/out; sleep 5; done"]
      volumeMounts:
        - name: persistent-storage
          mountPath: /data
  volumes:
    - name: persistent-storage
      persistentVolumeClaim:
        claimName: efs-claim
```


[FSx for Lustre を使用して高性能アプリケーションを保存する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/fsx-csi.html)


[FSx for NetApp ONTAP を使用して高性能アプリケーションを保存する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/fsx-ontap.html)


[Amazon FSx for OpenZFS を使用してデータを保存する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/fsx-openzfs-csi.html)


[Amazon File Cache を使用してレイテンシーを最小化する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/file-cache-csi.html)


[CSI ボリュームのためにスナップショット機能を有効にする](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/csi-snapshot-controller.html)



## ネットワーク

[Amazon EKS ネットワーク](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/eks-networking.html)


[Amazon EKS VPC およびサブネットの要件と考慮事項](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/network_reqs.html)

* クラスター
  * 異なる AZ の二つ以上のサブネットを指定
* VPC の要件、考慮事項
  * 十分な数の IP アドレス
  * VPC の DNS ホスト名、DNS 解決を有効化する必要がある
* サブネットの要件、考慮事項
  * 6 個以上の IP アドレスが必要。16 個以上を推奨
* ノードのサブネットの要件
  * ロードバランサをデプロイする場合は、タグ指定が必要
    * プライベートサブネット - キー: `kubernetes.io/role/internal-elb`, 値: 1
    * パブリックサブネット - キー: `kubernetes.io/role/elb`, 値: 1


[Amazon EKS クラスター VPC の作成](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/creating-a-vpc.html)

* VPC 作成用の CloudFormation サンプルが提供されている


[Amazon EKS セキュリティグループの要件および考慮事項](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/sec-group-reqs.html)

* クラスター作成時にセキュリティグループ `eks-cluster-sg-my-cluster-uniqueID` が作成される
* 次のタグが付与される
  * `kubernetes.io/cluster/my-cluster`	`owned`
  * `aws:eks:cluster-name`	`my-cluster`
  * `Name`	`eks-cluster-sg-my-cluster-uniqueid`
* 次のリソースにセキュリティグループがアタッチされる
  * クラスターの ENI
  * マネージドノードグループのネットワークインターフェイス


### ネットワーキングアドオン

[Amazon EKS ネットワーキングアドオン](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/eks-networking-add-ons.html)

* 次のアドオンはデフォルトでクラスターにインストールされている
  * Amazon VPC CNI plugin for Kubernetes
  * CoreDNS
  * kube-proxy


#### Amazon VPC CNI plugin

[Amazon VPC CNI plugin for Kubernetes Amazon EKS アドオンの使用](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/managing-vpc-cni.html)

* DaemonSet aws-node としてデプロイされる
* 参考記事 [Amazon VPC CNI プラグインでノード 1 台に配置可能な Pod 数を増やすために](https://aws.amazon.com/jp/blogs/news/amazon-vpc-cni-increases-pods-per-node-limits/)


[サービスアカウントの IAM ロールを使用する Amazon VPC CNI plugin for Kubernetes の設定](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/cni-iam-role.html)

* ノードロール、もしくは IRSA で AmazonEKS_CNI_Policy ポリシーを設定する必要がある


[Pod ネットワークのユースケースの選択](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/pod-networking-use-cases.html)


[クラスター、Pods、services 用の IPv6 アドレス](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/cni-ipv6.html)

* IPv6 アドレスを割り当て可能。デュアルスタックには対応していない


[Pods の SNAT](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/external-snat.html)

* VPC 外へ通信する際はノードのプライマリ ENI の IP アドレスに変換する
* プライベートサブネットにて NAT ゲートウェイがある構成では、以下の設定が必要
```
kubectl set env daemonset -n kube-system aws-node AWS_VPC_K8S_CNI_EXTERNALSNAT=true
```


[Kubernetes ネットワークポリシーのクラスターを構成する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/cni-network-policy.html)

* Fargate, Windows は未サポート
* `NETWORK_POLICY_ENFORCING_MODE` を `strict` に設定した場合はデフォルトは拒否ポリシーになる
* クラスターバージョン 1.26 以下の場合は BPF ファイルシステムのマウントが必要
```
sudo mount -t bpf bpffs /sys/fs/bpf
```
* アドオンで導入可能
* NetworkPolicy オブジェクトにより、どこからどこへの通信を許可するかを設定できる
```yaml
kind: NetworkPolicy
apiVersion: networking.k8s.io/v1
metadata:
  namespace: stars
  name: backend-policy
spec:
  podSelector:
    matchLabels:
      role: backend
  ingress:
    - from:
        - podSelector:
            matchLabels:
              role: frontend
      ports:
        - protocol: TCP
          port: 6379
```


[ポッド用のカスタムネットワーク](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/cni-custom-network.html)

* セカンダリネットワークインタフェースを異なるサブネットにしたり、異なるセキュリティグループを設定することが可能
* DaemonSet `aws-node` の環境変数 `AWS_VPC_K8S_CNI_CUSTOM_NETWORK_CFG` を `true` に設定する
* サブネットごとに `ENIConfig` オブジェクトを作成する
```yaml
apiVersion: crd.k8s.amazonaws.com/v1alpha1
kind: ENIConfig
metadata: 
  name: $az_1
spec: 
  securityGroups: 
    - $cluster_security_group_id
  subnet: $new_subnet_id_1
```
* Node オブジェクトにアノテーションの付与が必要
```shell
kubectl annotate node ip-192-168-0-126.us-west-2.compute.internal k8s.amazonaws.com/eniConfig=EniConfigName1
kubectl annotate node ip-192-168-0-92.us-west-2.compute.internal k8s.amazonaws.com/eniConfig=EniConfigName2
```


[Amazon EC2 ノードで使用可能な IP アドレスの量を増やす](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/cni-increase-ip-addresses.html)

* IP プレフィックスをノードに割り当てることで、ノードごとに使用可能な IP アドレスを増やすことができる
* Nitro ベースのインスタンスタイプである必要がある
* DaemonSet `aws-node` に以下の環境変数設定が必要
```
kubectl set env daemonset aws-node -n kube-system ENABLE_PREFIX_DELEGATION=true
```
* 各パラメータの調整。詳細は[こちら](https://github.com/aws/amazon-vpc-cni-k8s/blob/master/docs/prefix-and-ip-target.md)
```
kubectl set env ds aws-node -n kube-system WARM_PREFIX_TARGET=1
kubectl set env ds aws-node -n kube-system WARM_IP_TARGET=5
kubectl set env ds aws-node -n kube-system MINIMUM_IP_TARGET=2
```
* 起動テンプレートを使用しないマネージド型ノードグループの場合は max-pods の値が自動計算される
  * セルフマネージド型の場合は BootstrapArguments で以下の指定を行う
  ```shell
  --use-max-pods false --kubelet-extra-args '--max-pods=110'
  ```
  * マネージド型で AMI ID を指定している場合はユーザーデータで以下の指定を行う
  ```shell
  /etc/eks/bootstrap.sh my-cluster \
    --use-max-pods false \
    --kubelet-extra-args '--max-pods=110'
  ```


[ポッドのセキュリティグループ](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/security-groups-for-pods.html)

* 考慮事項
  * Windows は未サポート
  * ソース NAT が無効になる。よって、インターネットアクセスが必要な場合は、プライベートサブネットに配置し NAT Gateway などを使用する必要あり。
* 使用手順
  1. CNI プラグイン 1.7.7 以降で対応
  1. EKS クラスターロールに `AmazonEKSVPCResourceController` のポリシーをアタッチ
  1. aws-node の環境変数設定。`ENABLE_POD_ENI=true`
  1. SecurityGroupPolicy の Kind をクラスターにデプロイ
  ```yaml
  apiVersion: vpcresources.k8s.aws/v1beta1
  kind: SecurityGroupPolicy
  metadata:
    name: my-security-group-policy
    namespace: my-namespace
  spec:
    podSelector: 
      matchLabels:
        role: my-role
    securityGroups:
      groupIds:
        - my_pod_security_group_id
  ```
  1. Pod をデプロイ
  ```yaml
  apiVersion: apps/v1
  kind: Deployment
  metadata:
    name: my-deployment
    namespace: my-namespace
    labels:
      app: my-app
  spec:
    replicas: 4
    selector:
      matchLabels:
        app: my-app
    template:
      metadata:
        labels:
          app: my-app
          role: my-role
      spec:
        terminationGracePeriodSeconds: 120
        containers:
        - name: nginx
          image: public.ecr.aws/nginx/nginx:1.23
          ports:
          - containerPort: 80
  ---
  apiVersion: v1
  kind: Service
  metadata:
    name: my-app
    namespace: my-namespace
    labels:
      app: my-app
  spec:
    selector:
      app: my-app
    ports:
      - protocol: TCP
        port: 80
        targetPort: 80
  ```


[Pods 用の複数のネットワークインターフェイス](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/pod-multiple-network-interfaces.html)

* Multus CNI により Pod に複数のネットワークインターフェイスを割り当て可能


[互換性のある代替 CNI プラグイン](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/alternate-cni-plugins.html)

* 唯一 EKS でサポートされているのは Amazon VPC CNI plugin for Kubernetes


#### AWS Load Balancer Controller

[AWS Load Balancer Controller とは](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/aws-load-balancer-controller.html)

* 次のリソースのプロビジョニングを実行
  * Ingress により ALB を作成
  * LoadBalancer により NLB を作成 
* コントローラーのインストール方法
  * アドオン
  * helm


[Helm を使用して AWS Load Balancer Controller をインストールする](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/lbc-helm.html)

* IRSA 用の IAM ロールが必要。ロールには `AWSLoadBalancerControllerIAMPolicy` をアタッチ


[Kubernetes マニフェストを使用して AWS Load Balancer Controller アドオンをインストールする](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/lbc-manifest.html)

* cert-manager のインストールが必要


#### CoreDNS

[CoreDNS Amazon EKS アドオンの使用](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/managing-coredns.html)

* デフォルトでは 2 個の Pod がデプロイされる
* レプリカ数は次のように更新できる
```
aws eks update-addon --cluster-name my-cluster --addon-name coredns --addon-version v1.11.1-eksbuild.9 \
    --resolve-conflicts PRESERVE --configuration-values '{"replicaCount":3}'
```


[CoreDNS の自動スケーリング](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/coredns-autoscaling.html)

* アドオンの設定値にて以下の通り設定すると自動スケーリングが有効になる
```json
{
  "autoScaling": {
    "enabled": true
  }
}
```
* クラスター内のノード数、CPU 数に応じてスケールする


[CoreDNS のメトリクス](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/coredns-metrics.html)

* ポート番号 9153 で Prometheus 形式でメトリクスを公開している


#### kube-proxy

[Kubernetes kube-proxy アドオンの使用](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/managing-kube-proxy.html)

* DaemonSet で稼働する


### PrivateLink

[インターフェイスエンドポイント (AWS PrivateLink) を使用して Amazon Elastic Kubernetes Service にアクセス](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/vpc-interface-endpoints.html)

* EKS のインターフェイスエンドポイント
```
com.amazonaws.region-code.eks
```



## ワークロード

[ワークロード](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/eks-workloads.html)


[サンプルアプリケーションをデプロイする](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/sample-deployment.html)

* Sample Deployment
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: eks-sample-linux-deployment
  namespace: eks-sample-app
  labels:
    app: eks-sample-linux-app
spec:
  replicas: 3
  selector:
    matchLabels:
      app: eks-sample-linux-app
  template:
    metadata:
      labels:
        app: eks-sample-linux-app
    spec:
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: kubernetes.io/arch
                operator: In
                values:
                - amd64
                - arm64
      containers:
      - name: nginx
        image: public.ecr.aws/nginx/nginx:1.23
        ports:
        - name: http
          containerPort: 80
        imagePullPolicy: IfNotPresent
      nodeSelector:
        kubernetes.io/os: linux
```

* Sample Service
```yaml
apiVersion: v1
kind: Service
metadata:
  name: eks-sample-linux-service
  namespace: eks-sample-app
  labels:
    app: eks-sample-linux-app
spec:
  selector:
    app: eks-sample-linux-app
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
```


[Vertical Pod Autoscaler でポッドリソースを調整する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/vertical-pod-autoscaler.html)

* 前提として metrics-server が必要
* Vertical Pod Autoscaler をデプロイしておく必要がある
* リソースが不足している場合、Requests が更新され Pod が再起動される


[Horizontal Pod Autoscaler を使用してポッドデプロイをスケールする](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/horizontal-pod-autoscaler.html)

* 前提として metrics-server が必要
* Horizontal Pod Autoscaler リソースの作成例
```shell
kubectl autoscale deployment php-apache --cpu-percent=50 --min=1 --max=10
```


[Network Load Balancers で TCP および UDP トラフィックをルーティングする](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/network-load-balancing.html)

* L4 でのロードバランシング
* デフォルトでは AWS cloud provider load balancer controller により CLB が作成される。NLB の作成も可能
* AWS Load Balancer Controller の使用を推奨
* サービス設定で明示的にサブネットを指定しない場合は、サブネットにはタグ設定が必要
  * プライベートサブネットの場合
    * キー – `kubernetes.io/role/internal-elb`
    * 値 – 1
  * パブリックサブネットの場合。
    * キー – `kubernetes.io/role/elb`
    * 値 – 1
* Service(type: LoadBalancer) のサンプル
```yaml
apiVersion: v1
kind: Service
metadata:
  name: nlb-sample-service
  namespace: nlb-sample-app
  annotations:
    service.beta.kubernetes.io/aws-load-balancer-type: external
    service.beta.kubernetes.io/aws-load-balancer-nlb-target-type: ip
    service.beta.kubernetes.io/aws-load-balancer-scheme: internet-facing
spec:
  ports:
    - port: 80
      targetPort: 80
      protocol: TCP
  type: LoadBalancer
  selector:
    app: nginx
```


[Application Load Balancers でアプリケーションと HTTP トラフィックをルーティングする](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/alb-ingress.html)

* L7 でのロードバランシング
* NLB の場合と同様にサービス設定で明示的にサブネットを指定しない場合は、サブネットにはタグ設定が必要
* トラフィックモード
  * インスタンス: デフォルトで使用される。NodePort にルーティングする
  * IP: Pod に直接ルーティングされる。Fargate の場合に必要。`alb.ingress.kubernetes.io/target-type: ip` のアノテーションが必要
* [サンプル](https://raw.githubusercontent.com/kubernetes-sigs/aws-load-balancer-controller/v2.7.2/docs/examples/2048/2048_full.yaml) は 2048 ゲーム
```yaml
---
apiVersion: v1
kind: Namespace
metadata:
  name: game-2048
---
apiVersion: apps/v1
kind: Deployment
metadata:
  namespace: game-2048
  name: deployment-2048
spec:
  selector:
    matchLabels:
      app.kubernetes.io/name: app-2048
  replicas: 5
  template:
    metadata:
      labels:
        app.kubernetes.io/name: app-2048
    spec:
      containers:
      - image: public.ecr.aws/l6m2t8p7/docker-2048:latest
        imagePullPolicy: Always
        name: app-2048
        ports:
        - containerPort: 80
---
apiVersion: v1
kind: Service
metadata:
  namespace: game-2048
  name: service-2048
spec:
  ports:
    - port: 80
      targetPort: 80
      protocol: TCP
  type: NodePort
  selector:
    app.kubernetes.io/name: app-2048
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  namespace: game-2048
  name: ingress-2048
  annotations:
    alb.ingress.kubernetes.io/scheme: internet-facing
    alb.ingress.kubernetes.io/target-type: ip
spec:
  ingressClassName: alb
  rules:
    - http:
        paths:
        - path: /
          pathType: Prefix
          backend:
            service:
              name: service-2048
              port:
                number: 80
```


[サービスに割り当てることができる外部 IP アドレスを制限する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/restrict-service-external-ip.html)

* externalIPs に設定できる IP アドレスを制限できる。指定した CIDR 範囲外の場合はデプロイが失敗する
* webhook により実現している


[あるリポジトリから別のリポジトリにコンテナイメージをコピーする](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/copy-image-to-repository.html)

* イメージを ECR にプッシュする手順のドキュメント


[Amazon コンテナイメージレジストリ](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/add-ons-images.html)

* Amazon のコンテナイメージレジストリのドメインリスト


### アドオン

[AWSAPIs を使用して EKS アドオンでクラスターコンポーネントをインストール/更新する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/eks-add-ons.html)

* クラスターの作成方法によっては kube-proxy、Amazon VPC CNI plugin for Kubernetes、CoreDNS のセルフマネージド型アドオンがインストールされている
* ClusterRoleBinding `eks:addon-cluster-admin` により ClusterRole `cluster-admin` と Kubernetes Identity `eks:addon-manager` とを関連づけている。eks:addon-manager がアドオンに関する各種操作を行うのに必要


[使用可能な AWS の Amazon EKS アドオン](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/workloads-add-ons-available-eks.html)

* 使用可能なアドオンのリスト
* 各アドオンのページでは必要な IAM 許可やその他前提条件などがまとめられている
* 独立系のソフトウェアベンダのアドオンの情報もある


[Amazon EKS アドオンの作成](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/creating-an-add-on.html)

* 次のコマンド例により設定オプションを確認できる。このスキーマに従って、アドオンの設定内容を上書きできる
```shell
aws eks describe-addon-configuration --addon-name vpc-cni --addon-version v1.12.0-eksbuild.1
```
* アドオンの作成、更新時に `--configuration-values` オプションで設定値を上書きする。`--resolve-conflicts` オプションで競合に対する対応方法を設定できる


[フィールド管理を使用して Amazon EKS アドオン設定をカスタマイズする](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/kubernetes-field-management.html)


[Pod Identity を使用して Amazon EKS アドオンに IAM ロールをアタッチする](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/add-ons-iam.html)

* AWS CLI で `describe-addon-versions` を実行することにより IAM 許可が必要かを確認できる。`requiresIamPermissions` が `true` の場合は必要。`serviceAccount`、`recommendedManagedPolicies` も確認可能


### Others

[デプロイ中にコンテナイメージの署名を検証する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/image-verification.html)

* デプロイ時に署名済みのコンテナイメージを検証したい場合のソリューションへのリンクがあるページ


[Elastic Fabric Adapter を使用して Amazon EKS で機械学習トレーニングを実行する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/node-efa.html)

* EFA を Pod と統合できる


[Amazon EKS で AWSInferentia を使用して ML 推論ワークロードをデプロイする](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/inferentia-support.html)

* EC2 Inf1 インスタンスを使用できる



## クラスターの管理

[クラスターリソースを整理およびモニタリングする](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/eks-managing.html)

* クラスターの運用、管理で有用な情報がまとめられている


[Amazon EKS クラスターのコストをモニタリングして最適化する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/cost-monitoring.html)

* kubecost の使用方法もまとめられている
  * helm もしくはアドオンにより導入可能
  * port-forward を行い、ダッシュボードを表示可能


[KubernetesMetrics Server でリソースの使用状況を表示する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/metrics-server.html)

* HPA などでも使用される
* 監視ソリューション向きではないので、モニタリングには別のモニタリングソフトウェアが必要


[Helm を使用して Amazon EKS にアプリケーションをデプロイする](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/helm.html)


[タグを使用して Amazon EKS リソースを整理する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/eks-using-tags.html)


[Service Quotas](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/service-quotas.html)



## セキュリティ

[Amazon EKS のセキュリティ](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/security.html)


[証明書への署名](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/cert-signing.html)


### IAM

[Amazon EKS の Identity and Access Management](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/security-iam.html)


[Amazon EKS で IAM を使用する方法](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/security_iam_service-with-iam.html)


[Amazon EKS でのアイデンティティベースのポリシーの例](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/security_iam_id-based-policy-examples.html)

* EKS クラスターの作成者にはデフォルトで RBAK で `system:masters` の許可が付与される


[Amazon EKS でのサービスにリンクされたロールの使用](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/using-service-linked-roles.html)


[Amazon EKS クラスターのロールの使用](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/using-service-linked-roles-eks.html)

* `AWSServiceRoleForAmazonEKS`: [AmazonEKSServiceRolePolicy](https://docs.aws.amazon.com/aws-managed-policy/latest/reference/AmazonEKSServiceRolePolicy.html) がアタッチされている


[Amazon EKS ノードグループでのロールの使用](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/using-service-linked-roles-eks-nodegroups.html)

* `AWSServiceRoleForAmazonEKSNodegroup`: [AWSServiceRoleForAmazonEKSNodegroup](https://docs.aws.amazon.com/aws-managed-policy/latest/reference/AWSServiceRoleForAmazonEKSNodegroup.html) がアタッチされている。ノードグループの管理用


[Amazon EKS Fargate プロファイルのロールの使用](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/using-service-linked-roles-eks-fargate.html)

* `AWSServiceRoleForAmazonEKSForFargate`: [AmazonEKSForFargateServiceRolePolicy](https://docs.aws.amazon.com/aws-managed-policy/latest/reference/AmazonEKSForFargateServiceRolePolicy.html) がアタッチされている。Fargate Pod 起動、停止用


[ロールを使用して Kubernetes クラスターを Amazon EKS に接続する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/using-service-linked-roles-eks-connector.html)

* `AWSServiceRoleForAmazonEKSConnector`: [AmazonEKSConnectorServiceRolePolicy](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/using-service-linked-roles-eks-connector.html) がアタッチされている。EKS が EKS クラスターに接続する用途


[Outpost で Amazon EKS ローカルクラスターのロールの使用](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/using-service-linked-roles-eks-outpost.html)

* `AWSServiceRoleForAmazonEKSLocalOutpost`: [AmazonEKSServiceRolePolicy](https://docs.aws.amazon.com/ja_jp/aws-managed-policy/latest/reference/AmazonEKSServiceRolePolicy.html) がアタッチされている


[Amazon EKS クラスター の IAM ロール](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/service_IAM_role.html)

* クラスターの管理用の IAM ロール。例えば [AmazonEKSClusterPolicy](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/service_IAM_role.html) のようなポリシーをアタッチする


[Amazon EKS ノードの IAM ロール](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/create-node-role.html)

* kubelet が使用する IAM ロール
  * 例えば [AmazonEKSWorkerNodePolicy ](https://docs.aws.amazon.com/ja_jp/aws-managed-policy/latest/reference/AmazonEKSWorkerNodePolicy.html)。このポリシーには EKS Pod Identity を使用するための `eks-auth:AssumeRoleForPodIdentity` の許可が含まれている
  * ECR を使用する場合は、[AmazonEC2ContainerRegistryReadOnly](https://docs.aws.amazon.com/ja_jp/aws-managed-policy/latest/reference/AmazonEC2ContainerRegistryReadOnly.html) のようなポリシーをアタッチする


[Amazon EKS Pod 実行 IAM ロール](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/pod-execution-role.html)

* ECS タスク実行ロールのような役割。ECR への許可が含まれている [AmazonEKSFargatePodExecutionRolePolicy](https://docs.aws.amazon.com/ja_jp/aws-managed-policy/latest/reference/AmazonEKSFargatePodExecutionRolePolicy.html) をアタッチする


[Amazon EKS コネクタの IAM ロール](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/connector_IAM_role.html)

* `AmazonEKSConnectorAgentRole` ロールを作成し `AmazonEKSConnectorAgentPolicy` ポリシーをアタッチする


[Amazon Elastic Kubernetes Service に関する AWS 管理ポリシー](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/security-iam-awsmanpol.html)

* マネージドポリシーが多く用意されている


[IAM のトラブルシューティング](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/security_iam_troubleshoot.html)


[デフォルトの Amazon EKS は Kubernetes ロールとユーザーを作成しました。](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/default-roles-users.html)


### Others

[Amazon Elastic Kubernetes Service でのコンプライアンス検証](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/compliance.html)


[Amazon EKS の耐障害性](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/disaster-recovery-resiliency.html)


[Amazon EKS でのインフラストラクチャセキュリティ](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/infrastructure-security.html)


[Amazon EKS での設定と脆弱性の分析](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/configuration-vulnerability-analysis.html)


[Amazon EKS のセキュリティベストプラクティス](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/security-best-practices.html)

* [https://aws.github.io/aws-eks-best-practices/security/docs/](https://aws.github.io/aws-eks-best-practices/security/docs/) に公開されている


[ポッドのセキュリティポリシー](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/pod-security-policy.html)

* PodSecurityPolicy (PSP) は Kubernetes バージョン 1.21 で非推奨となり、Kubernetes 1.25 で削除されている


[ポッドセキュリティポリシー (PSP) の削除に関するよくある質問](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/pod-security-policy-removal-faq.html)

* 組み込みの Kubernetes ポッドセキュリティ標準 (PSS) または Policy as Code ソリューションに移行する必要がある


[Kubernetes で AWS Secrets Manager シークレットを使用する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/manage-secrets.html)

* Kubernetes Secrets Store CSI Driver 向けの AWS Secrets and Configuration Provider (ASCP) を使用することで対応可能
* 詳細は [Amazon Elastic Kubernetes Service で AWS Secrets Manager シークレットを使用する](https://docs.aws.amazon.com/ja_jp/secretsmanager/latest/userguide/integrating_csi_driver.html) を参照すること
* `SecretProviderClass` オブジェクトで SecretsManager のリソース ARN を指定する
* Deployment では次の例のように指定する
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 2
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      serviceAccountName: nginx-deployment-sa
      volumes:
      - name: secrets-store-inline
        csi:
          driver: secrets-store.csi.k8s.io
          readOnly: true
          volumeAttributes:
            secretProviderClass: "nginx-deployment-aws-secrets"
      containers:
      - name: nginx-deployment
        image: nginx
        ports:
        - containerPort: 80
        volumeMounts:
        - name: secrets-store-inline
          mountPath: "/mnt/secrets-store"
          readOnly: true
```


[Amazon EKS Connector の考慮事項](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/security-connector.html)


[Kubernetes リソースを表示する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/view-kubernetes-resources.html)



## Monitoring

[クラスターのパフォーマンスをモニタリングし、ログを表示する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/eks-observe.html)

* モニタリング計画を立てる必要がある
* ログ記録用ツール
  * Amazon CloudWatch Container Insights
  * AWS CloudTrail
  * AWS Fargate ログルーター
* モニタリングツール
  * CloudWatch Container Insights
  * AWS Distro for OpenTelemetry (ADOT)
  * Amazon DevOps Guru
  * AWS X-Ray
  * Amazon CloudWatch オブザーバビリティオペレータ
  * Prometheus


[Prometheus を使用してクラスターのメトリクスをモニタリングする](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/prometheus.html)

* クラスター作成時に Prometheus にメトリクスを送信するオプションを有効化できる
* スクレイパーは、クラスターインフラストラクチャーとコンテナー化されたアプリケーションからデータを収集するように設定されている


[Helm を使用して Prometheus をデプロイする](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/deploy-prometheus.html)


[コントロールプレーンの raw メトリクスを Prometheus 形式で表示する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/view-raw-metrics.html)

* コントロールプレーンのメトリクスは次のコマンドで確認できる
```
kubectl get --raw /metrics
```


[Amazon CloudWatch でクラスターデータをモニタリングする](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/cloudwatch.html)

* EKS アドオンとして Amazon CloudWatch Observability Operator を導入できる


[コントロールプレーンログを CloudWatch Logs に送信する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/control-plane-logs.html)

* 次の 5 つのログについて有効/無効を設定できる
  * API サーバ
  * Audit
  * authenticator
  * controllerManager
  * scheduler


[API コールを AWS CloudTrail イベントとしてログ記録する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/logging-using-cloudtrail.html)


[ADOT Operator を使用してメトリクスとトレースデータを送信する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/opentelemetry.html)

* 詳細は [Getting Started with AWS Distro for OpenTelemetry using EKS Add-Ons](https://aws-otel.github.io/docs/getting-started/adot-eks-add-on) を参照すること
* ブログ記事もある。[AWS Distro for OpenTelemetry の Amazon EKS アドオンを使用したメトリクスとトレースの収集](https://aws.amazon.com/jp/blogs/news/metrics-and-traces-collection-using-amazon-eks-add-ons-for-aws-distro-for-opentelemetry/)
  * EKS の ADOT アドオンとして導入可能
    * カスタムリソース `OpenTelemetryCollector` が作成される
    * X-Ray: Receiver は Service で公開されるので、アプリ側では `AWS_XRAY_DAEMON_ADDRESS` を `observability-collector.aws-otel-eks:2000` にするとよい。
    * Prometheus Receiver: Receiver はサービスディスカバリを行いメトリクスをスクレイピングする。Amazon Managed Service for Prometheus の作成済みワークスペースに送信する
    * CloudWatch にメトリクスを送信するように YAML を記述することも可能
    * Exporter は IRSA により各サービスに送信するためのアクセス許可が必要。サービスアカウントは aws-otel-collector



## EKS Integrations

[AWS CloudFormation を使用して Amazon EKS リソースを作成する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/creating-resources-with-cloudformation.html)


[深層学習コンテナを使用して EKS で TensorFlow モデルをトレーニングして提供する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/deep-learning-containers.html)


[Amazon Detective を使用して EKS のセキュリティイベントを分析する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/integration-detective.html)

* セキュリティに関する検出結果や疑わしいアクティビティの根本原因を分析、調査、および迅速に特定できる


[Amazon GuardDuty で脅威を検出する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/integration-guardduty.html)

* EKS Protection
  * 監査ログをモニタリングしている。脅威検出カバレッジを提供
* Runtime Monitoring
  * GuardDuty エージェントをインストールする必要がある


[AWS Resilience Hub で EKS クラスターの耐障害性を評価する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/integration-resilience-hub.html)

* EKS クラスターの耐障害性を評価できる


[Security Lake を使用して EKS セキュリティデータを一元化して分析する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/integration-securitylake.html)

* セキュリティデータを一元化して分析できるサービス


[Amazon VPC Lattice との安全なクラスター間接続を有効にする](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/integration-vpc-lattice.html)

* Kubernetes の Gateway API を実装したコントローラを使用して VPC Lattice を利用できる


[AWS Local Zones を使用して低レイテンシーの EKS クラスターを起動する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/local-zones.html)



## Troubleshooting

* Insufficient capacity
* Nodes fail to join cluster
  * ノードには次のタグ設定が必要
    * kubernetes.io/cluster/<クラスター名>: owned
  * API エンドポイントへの疎通性の確保が必要
  * VPC の DHCP オプションセット設定の確認
* kubectl の認証エラー
  * kubeconfig ファイルの確認
  * アクセスエントリの設定確認



## Amazon EKS Connector

[Amazon EKS Connector を使用して Kubernetes クラスターを Amazon EKS マネジメントコンソールに接続する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/eks-connector.html)

準拠する Kubernetes クラスターを AWS に登録および接続し、Amazon EKS コンソールで可視化できる。そのため、オンプレミスや EC2 常に構築した Kubernetes クラスターを登録できる



## Outposts

[AWS Outposts を使用して Amazon EKS をオンプレミスにデプロイする](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/eks-outposts.html)



## Related Projects

[オープンソースプロジェクトで Amazon EKS の機能を拡張する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/related-projects.html)



## Roadmap

[Amazon EKS 新機能とロードマップの詳細](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/roadmap.html)



# 参考

* Document
  * [Amazon EKS とは?](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/what-is-eks.html)
* サービス紹介ページ
  * [Amazon EKS の特徴](https://aws.amazon.com/jp/eks/features/)
  * [よくある質問](https://aws.amazon.com/jp/eks/faqs/)
* [ナレッジセンター](https://aws.amazon.com/jp/premiumsupport/knowledge-center/#Amazon_Elastic_Container_Service_for_Kubernetes_.28Amazon_EKS.29)
* Black Belt
  * [20190410 AWS Black Belt Online Seminar Amazon Elastic Container Service for Kubernetes (Amazon EKS)](https://www.slideshare.net/AmazonWebServicesJapan/20190410-aws-black-belt-online-seminar-amazon-elastic-container-service-for-kubernetes-amazon-eks)
* [API Reference](https://docs.aws.amazon.com/ja_jp/eks/latest/APIReference/Welcome.html)
* 外部ドキュメント、ブログ
  * [EKS on FargateでALBからアプリにアクセスする](https://839.hateblo.jp/entry/2019/12/08/172020)

