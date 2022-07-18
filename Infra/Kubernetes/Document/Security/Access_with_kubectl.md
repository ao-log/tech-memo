
#### kubectl によるアクセス

[Accessing for the first time with kubectl](https://kubernetes.io/docs/tasks/access-application-cluster/access-cluster/#accessing-for-the-first-time-with-kubectl)

kubectl proxy なしの場合。

サービスアカウント用の secret を作成する。
```yaml
Version: v1
kind: Secret
metadata:
  name: default-token
  annotations:
    kubernetes.io/service-account.name: default
type: kubernetes.io/service-account-token
```

token を用いてアクセスする。--insecure フラグになっている点に注意。本来は ~/.kube ディレクトリ下の証明書を使用する。
```shell
APISERVER=$(kubectl config view --minify -o jsonpath='{.clusters[0].cluster.server}')
TOKEN=$(kubectl get secret default-token -o jsonpath='{.data.token}' | base64 --decode)

curl $APISERVER/api --header "Authorization: Bearer $TOKEN" --insecure
```


## EKS の場合

[クラスター認証](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/cluster-auth.html)

IAM アイデンティティと RoleBinding とを紐づけることができる。
ConfigMap の ```aws-auth``` に関係を記載する。

Qiita の記事 [EKSでIAMユーザをアクセスコントロール (RBAC)](https://qiita.com/taishin/items/dfb9a5620f37ffb74fe9) も分かりやすかった。


[クラスターへの IAM ユーザーおよびロールアクセスを有効にする](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/add-user-role.html)

コントロールプレーンでは [Authenticator for Kubernetes](https://github.com/kubernetes-sigs/aws-iam-authenticator#aws-iam-authenticator-for-kubernetes) が実行されている。
ConfigMap の ```aws-auth``` から認証情報を読み取る。

EKS クラスターを作成した IAM アイデンティティは RBAC で system:masters が自動的に付与されている。
この IAM アイデンティティの情報を参照する方法はない。

IAM ユーザー、ロールを追加する場合には ConfigMap の ```aws-auth``` を編集する対応が必要。
```shell
$ kubectl edit -n kube-system configmap/aws-auth
```


[Amazon EKS の kubeconfig を作成する](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/create-kubeconfig.html)

以下のコマンドで kubeconfig を作成、更新できる。
```
$ aws eks update-kubeconfig --region region-code --name cluster-name
```


