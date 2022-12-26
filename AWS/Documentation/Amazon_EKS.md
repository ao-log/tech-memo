

[Amazon EKS とは何ですか?](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/what-is-eks.html)

コントロールプレーンがマネージドになっている。
2 つ以上の API ノード、3 つの etcd ノードから構成されている。



## 開始方法

[eksctl の開始方法](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/getting-started-eksctl.html)

kubectl, eksctl はバイナリをダウンロードするだけで使用できる。
eksctl はコマンドラインオプションだけでなく、yaml ファイルで設定値を渡すことも可能。


[AWS マネジメントコンソール の開始方法](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/getting-started-console.html)

以下のステップで EKS クラスターを作成している。

* サービス用の IAM ロールの作成
* CloudFormation テンプレートにより VPC とサブネットを作成
* マネジメントコンソール上で EKS クラスターを作成。次の項目を設定。
  * 上記で作成した IAM ロール、サブネット
  * クラスターエンドポイントのアクセス(パブリック、プライベート、パブリックおよびプライベート)
  * ログ記録の構成
* aws eks update-kubeconfig ... による kubeconfig ファイルの作成(.kube/config に設定が追加される)
* ノードグループの作成



## クラスター

[Amazon EKS クラスターの作成](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/create-cluster.html)

３つの作成方法がある。

* eksctl
* マネジメントコンソール
* AWS CLI

シークレットの暗号化を KMS キーで行う場合、鍵を削除するとクラスターを復旧できなくなる。


[Amazon EKS クラスターの Kubernetes バージョンの更新](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/update-cluster.html)

* クラスターの更新時はサービスが短時間中断することがある。
* クラスターの更新時にはアドオンは更新されない。
* コントロールプレーンをアップデートしてから、データプレーン側をアップデートする順番で対応する。


[クラスターの削除](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/delete-cluster.html)

eksctl で対応する場合、サービスの削除 → クラスターの削除の順番。
マネジメントコンソールから対応する場合、クラスターの削除の前に Fargate プロファイル、ノードグループを削除する必要あり。


[Amazon EKS クラスターエンドポイントのアクセスコントロール](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/cluster-endpoint.html)

Kubernetes API サーバーエンドポイントの公開範囲を設定可能。

デフォルトではパブリックになっている。この場合でも IAM, RBAC により保護されている。
また、許可した接続元 IP アドレスのみアクセスできるよう設定することも可能。VPC 内のノードからもアクセスができるようにするには、その IP アドレスも許可対象に加える、もしくはプライベートエンドポイント有効にする必要がある。

プライベートに設定することで、API エンドポイントにアクセス可能な接続元を VPC 内に留めることができる。


[Cluster Autoscaler](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/cluster-autoscaler.html)

クラスター内のノードの数を自動的に調整できる。


[Amazon EKS コントロールプレーンのログ記録](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/control-plane-logs.html)

以下 5 種類のログがある。それぞれ有効化、無効化できる。

* Kubernetes API サーバーコンポーネントログ (api)
* 監査 (audit)
* 認証 (authenticator)
* コントローラーマネージャー (controllerManager)
* スケジューラ (scheduler)


[Amazon EKS Kubernetes バージョン](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/kubernetes-versions.html)

Kubernetes のマイナーバージョンは約 3 ヶ月ごとにリリースされている。
各マイナーバージョンは最初にリリースされてから約 12 ヶ月間サポートされる。
EKS は 4 つの本稼働準備が整ったバージョンをサポートしている。


[プライベートクラスター](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/private-clusters.html)

インターネットへのアウトバウントの疎通性のないプライベートクラスターを構築可能。
以下の要件を満たす必要がある。

* イメージレジストリは ECR か VPC 内のレジストリが必要
* クラスターエンドポイントのプライベート設定
* 必要に応じて VPC エンドポイントを作成
* セルフマネージド型ノードの場合はブートストラップ引数のオプション追加が必要



## ノード

[Amazon EKSコンピューティング](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/eks-compute.html)

ノードグループは次の 3 種類ある。

* マネージド型
* セルフマネージド型
* Fargate


[マネージド型ノードグループ](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/managed-node-groups.html)

* ASG の一部としてプロビジョニングされる。
* Cluster Autoscaler を使用するようにタグ付けされる。
* デフォルトでは EKS 最適化された Amazon Linux 2 AMI を使用。カスタム AMI を使用することも可。
* 1 つのクラスター内に複数のマネージド型ノードグループを作成可能。


[マネージド型ノードグループの作成](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/create-managed-node-group.html)

eksctl もしくはマネジメントコンソールより作成可能。

eksctl の場合のコマンド例。

```
eksctl create nodegroup \
  --cluster <my-cluster> \
  --region <us-west-2> \
  --name <my-mng> \
  --node-type <m5.large> \
  --nodes <3> \
  --nodes-min <2> \
  --nodes-max <4> \
  --ssh-access \
  --ssh-public-key <my-public-key.pub> \
  --managed
```

マネジメントコンソールでは次の設定項目がある。

* ノードグループ名
* インスタンスロール
* 起動テンプレート
* Kubernetes ラベル
* タグ
* AMI のタイプ
* インスタンスタイプ
* Auto Scaling の台数設定(min, max, desired)
* サブネット
* SSH キーペア


[マネージド型ノードグループの更新](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/update-managed-node-group.html)

eksctl もしくはマネジメントコンソール上の操作でノードグループを更新可能。


[マネージド型ノードの更新動作](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/managed-node-update-behavior.html)

マネージド型ノードグループを更新した場合の動作の流れが書かれている。

1. 新しい起動テンプレートを作成
1. 新しい起動テンプレートを使用するように Auto Scaling グループを更新
1. Auto Scaling グループの最大サイズ、必要なサイズを最大 2 倍に増やす
1. 最新の AMI ID でラベル付をされていないノードグループ内のノードに  eks.amazonaws.com/nodegroup=unschedulable:NoSchedule の taint。
1. ノードグループ内のノードをランダムに選択し、Pod を削除。
1. ノードを cordon する。このノードに新しいリクエストを送信しないようにするため。
1. cordon されたノードを terminate。
1. 元のバージョンのノードグループからノードがなくなるまで 5〜7 を繰り返す。
1. Auto Scaling グループのサイズをもとに戻す。


[セルフマネージド型ノード](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/worker.html)


[セルフマネージド型 Amazon Linux ノードの起動](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/launch-workers.html)

eksctl の場合のコマンド例。

```
eksctl create nodegroup \
--cluster <my-cluster> \
--version auto \
--name <al-nodes> \
--node-type <t3.medium> \
--node-ami auto \
--nodes <3> \
--nodes-min <1> \
--nodes-max <4>
```

マネジメントコンソール上から作成する場合。

CloudFormation テンプレートから作成できる。また、ノードがクラスターに参加するには、ConfigMap にインスタンスロールの設定が必要。


[セルフマネージド型ノードの更新](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/update-workers.html)

新しいノードグループを作成したあと元のノードグループを削除する方法がある。
また、既存のノードグループのスタックの AMI を更新することでも対応可能。

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
  1. ConffigMap から元のノードグループのインスタンスロールを削除


[AWS Fargate](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/fargate.html)

いくつか制約がある。

* DaemonSet はサポートされない
* 特権を持つコンテナはサポートされない
* プライベートサブネットでのみサポート


[Amazon EKS を使用した AWS Fargate の開始方法](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/fargate-getting-started.html)

eksctl の場合のコマンド例。

```
eksctl create cluster --name <my-cluster> --version <1.18> --fargate
```

Fargate で実行されている Pod は、クラスターのクラスターセキュリティグループを使用するように設定される。


[AWS Fargate プロファイル](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/fargate-profile.html)

クラスターの Fargate で Pod をスケジュールする前に Fargate Profile の作成が必要。

次のコマンドで Fargate Profile を作成可能。

```
eksctl create fargateprofile \
    --cluster <cluster_name> \
    --name <fargate_profile_name> \
    --namespace <kubernetes_namespace> \
    --labels <key=value>
```

マネジメントコンソールからも作成可能。Pod 実行ロール、サブネット、namespace などを指定する。



## ストレージ

[Storage](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/storage.html)

[ストレージクラス](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/storage-classes.html)

ストレージクラスの作成方法。次のような yaml から作成可能。

```yaml
kind: StorageClass
apiVersion: storage.k8s.io/v1
metadata:
  name: gp2
  annotations:
    storageclass.kubernetes.io/is-default-class: "true"
provisioner: kubernetes.io/aws-ebs
parameters:
  type: gp2
  fsType: ext4 
```


[Amazon EBS CSI ドライバー](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/ebs-csi.html)

次の流れで適用する。

1. IAM ロールの作成(EKS クラスターから EBS 関連の API を呼び出す用途)
1. CSI ドライバーの YAML を適用
1. サービスアカウント ebs-csi-controller-sa に IAM ロールのアノテーションを追記
1. ドライバーの Pod を削除(新しい Pod を適用)



## ネットワーク

[Amazon EKS ネットワーク](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/eks-networking.html)

[クラスター VPC に関する考慮事項](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/network_reqs.html)

サブネットには所定の内容でタグを付与する必要がある。
そうすることで kubernetes から認識できるようになる。

例

* key: kubernetes.io/cluster/<cluster-name>
* value: shared

shared は複数クラスターによるサブネットの仕様を許可。


[Amazon EKS セキュリティグループの考慮事項](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/sec-group-reqs.html)

コントロールプレーン、データプレーンそれぞれにセキュリティグループを設定可能。


[ポッドネットワーキング (CNI)](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/pod-networking.html)

CNI プラグインにより Kubernetes Pod は VPC ネットワーク上と同じ IP アドレスを Pod 内に持つことができるようになる。
各 EC2 インスタンスに aws-node の DaemonSet としてデプロイされる。

プラグインは次の二つの主要なコンポーネントで構成されている。

**L-IPAM デーモン**

ENI のアタッチ、セカンダリ IP アドレスの割当、IP アドレスプールの維持など。

**CNI プラグイン**

[外部ソースネットワークアドレス変換 (SNAT)](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/external-snat.html)

トラフィックが VPC 外のアドレスに向けられる際は、CNI プラグインは、デフォルトで各 Pod のプライベート IP アドレスを Pod が実行されているノードのプライマリ ENI に割り当てられたプライマリプライベートアドレスに SNAT で変換する。

[サービスアカウントの IAM ロールを使用するように VPC CNI プラグインを設定する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/cni-iam-role.html)

サービスアカウント aws-node を作成し、IAM ロール AmazonEKS_CNI_Policy をアタッチする必要がある。

[ポッドのセキュリティグループ](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/security-groups-for-pods.html)

いくつか制約がある。

* Kubernetes バージョン 1.17 以降で対応。
* Fargate では使用不可。
* ソース NAT が無効になる。よって、インターネットアクセスが必要な場合は、プライベートサブネットに配置し NAT Gateway などを使用する必要あり。

使用するには以下の対応が必要

1. CNI プラグイン 1.7.7 以降で対応。
1. EKS クラスターロールに AmazonEKSVPCResourceController のポリシーをアタッチ
1. aws-node の環境変数設定。ENABLE_POD_ENI=true
1. SecurityGroupPolicy の Kind をクラスターにデプロイ
1. Pod をデプロイ

[AWS Load Balancer コントローラー](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/aws-load-balancer-controller.html)

使用するには以下の対応が必要

1. IAM サービスアカウントを作成
1. AWS Load Balancer コントローラーをインストール



## ワークロード

[Horizontal Pod Autoscaler](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/horizontal-pod-autoscaler.html)

Horizontal Pod Autoscaler リソースの作成例。
```
kubectl autoscale deployment php-apache --cpu-percent=50 --min=1 --max=10
```


[ネットワーク負荷分散](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/load-balancing.html)

NLB または CLB が対応している。

サービス設定で明示的にサブネットを指定しない場合は、サブネットにはタグ設定が必要。
プライベートサブネットの場合。

* キー – kubernetes.io/role/internal-elb
* 値 – 1

パブリックサブネットの場合。

* キー – kubernetes.io/role/elb
* 値 – 1

NLB または CLB は Kubernetes インツリー負荷分散コントローラーによって作成される。
次の例のようにアノテーションを記載する。

```yaml
apiVersion: v1
kind: Service
metadata:
  name: sample-service
  annotations:
    service.beta.kubernetes.io/aws-load-balancer-type: nlb-ip
spec:
  ports:
    - port: 80
      targetPort: 80
      protocol: TCP
  type: LoadBalancer
  selector:
    app: nginx
```


[アプリケーションの負荷分散](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/alb-ingress.html)

AWS Load Balancer Controller によって ALB のリソースが作成される。

ALB が対応しているトラフィックモード。

* インスタンス – クラスター内のノードを ALB のターゲットとして登録。ALB に到達するトラフィックは、サービスの NodePort にルーティングされてから、ポッドにプロキシされる。
* IP – ポッドを ALB のターゲットとして登録。ALB に到達するトラフィックは、サービスのポッドに直接ルーティングされる。alb.ingress.kubernetes.io/target-type: ip のアノテーションが必要。

ドキュメントに載っているサンプルは以下のもの。

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
      - image: alexwhen/docker-2048
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
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  namespace: game-2048
  name: ingress-2048
  annotations:
    kubernetes.io/ingress.class: alb
    alb.ingress.kubernetes.io/scheme: internet-facing
    alb.ingress.kubernetes.io/target-type: ip
spec:
  rules:
    - http:
        paths:
          - path: /*
            backend:
              serviceName: service-2048
              servicePort: 80
```



## クラスター認証

[クラスター認証](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/managing-auth.html)

EKS は、 IAM を使用して Kubernetes クラスターに認証を提供する。
実際の認証には、従来の Kubernetes RBAC システムを使用する。

[クラスターのユーザーまたは IAM ロールの管理](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/add-user-role.html)

EKS クラスター作成時にクラスターを作成した IAM ユーザー、ロールは RBAC の設定で system:masters の権限が付与される。

ノードをクラスターに追加させるには aws-auth の ConfigMap に IAM ロールを設定する必要がある。

IAM ユーザー、IAM ロールを追加するには aws-auth の ConfigMap に mapRoles、mapUsers を設定する。

[kubeconfig を作成する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/create-kubeconfig.html)

次のコマンドで kubeconfig を作成、更新可能。
```
$ aws eks --region <region-code> update-kubeconfig --name <cluster_name>
```



## クラスターの管理

[kubectl のインストール](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/install-kubectl.html)

[eksctl のインストール](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/eksctl.html)

[Kubernetes ダッシュボード (ウェブ UI) のデプロイ](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/dashboard-tutorial.html)

[Kubernetes メトリクスサーバーのインストール](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/metrics-server.html)

[Prometheus メトリクス](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/prometheus.html)



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

