
#### 関連情報

[Configure Service Accounts for Pods](https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/)


#### リファレンス

[ServiceAccount v1 core](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#serviceaccount-v1-core)


#### ServiceAccount

* Pod 起動時にサービスアカウントを指定しなかった場合は同ネームスペース内の deafult サービスアカウントが使用される。
* Pod が使用しているサービスアカウントは、```kubectl get pods/<podname> -o yaml``` で確認した際、```spec.serviceAccountName``` にて確認できる。
* サービスアカウントの作成時にサービスアカウントと関連する Secret も自動生成される。Secret は証明書と署名された JSON Web Token(JWT) を保持している。
```yaml
apiVersion: v1
data:
  ca.crt: (base64でエンコードされた証明書)
  namespace: ZGVmYXVsdA==
  token: (base64でエンコードされたBearerトークン)
kind: Secret
metadata:
  # ...
type: kubernetes.io/service-account-token
```
* サービスアカウントトークンは Pod の```/var/run/secrets/kubernetes.io/serviceaccount/token``` にマウントされる。```automountServiceAccountToken: false``` の設定によりオプトアウト(つまりマウントしない)ように設定できる。
```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: build-robot
automountServiceAccountToken: false
...
```
* 署名された JWT は、与えられたサービスアカウントとして認証するための Bearer トークンとして使用できる。
* Kubernetes API に問い合わせる際は HTTP リクエストに HTTP ヘッダ ```--header "Authorization: Bearer $TOKEN"``` を含める。
* サービスアカウントはユーザー名 ```system:serviceaccount:(NAMESPACE):(SERVICEACCOUNT)``` で認証される。グループ ```system:serviceaccounts``` と ```system:serviceaccounts:(NAMESPACE)``` に割り当てられる。


## EKS の場合

[サービスアカウントの IAM ロール](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/iam-roles-for-service-accounts.html)

サービスアカウントを IAM ロールと関連づけることができる。
そのため、ワーカーノード単位での IAM ロールの割り当てが不要になる。
Pod 内から AWS の API を実行できるようになる。

* IAM ロールの信頼ポリシーでは以下のように OIDC プロバイダーの Principal を許可する。AssumeRoleWithWebIdentity は一時的な認証情報を取得するための API。
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Federated": "arn:aws:iam::111122223333:oidc-provider/OIDC_PROVIDER"
      },
      "Action": "sts:AssumeRoleWithWebIdentity",
      "Condition": {
        "StringEquals": {
          "OIDC_PROVIDER:sub": "system:serviceaccount:SERVICE_ACCOUNT_NAMESPACE:SERVICE_ACCOUNT_NAME"
        }
      }
    }
  ]
}
```
* サービスアカウント側ではアノーテーションを記載する。
```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  annotations:
    eks.amazonaws.com/role-arn: arn:aws:iam::111122223333:role/iam-role-name
```
* kubelet が token を格納し、Pod にマウントする。token が合計 TTL の 80 % を超えている場合、24 時間を超えている場合に更新する。
* コンテナ内のプロセスからは root アカウントか securityContext を設定することで Web ID トークンファイルを読み取ることができる。
```yaml
kind: Deployment
metadata:
  name: my-app
spec:
  template:
    metadata:
      labels:
        app: my-app
    spec:
      serviceAccountName: my-app
      containers:
      - name: my-app
        image: my-app:latest
      securityContext:
        fsGroup: 1337
```


[詳解: IAM Roles for Service Accounts](https://aws.amazon.com/jp/blogs/news/diving-into-iam-roles-for-service-accounts/)

#### 前提
* IAM OIDC プロバイダーを関連付けておく必要がある。
```
$ eksctl utils associate-iam-oidc-provider --region=us-east-2 --cluster=eks-oidc-demo --approve
```

#### 仕組み
* トークンのマウントに関して
  * Kubernetes 1.12 では、Service Account Token Volume Projection という機能が導入された。
  * EKS の Kubernetes 1.21 以降のクラスターでは BoundServiceAccountTokenVolume フィーチャーフラグがデフォルトで有効化されているため、OIDC 完全準拠 の JWT トークンを Projected Volume として Pod にマウントする。(自動生成された Secret のトークンとは別物)
  * ```/var/run/secrets/kubernetes.io/serviceaccount/token``` は Kubernetes 用のトークン
  * ```/var/run/secrets/eks.amazonaws.com/serviceaccount/token``` は IAM 用のトークン
  * [Pod Identity Webhook](https://github.com/aws/amazon-eks-pod-identity-webhook) により Pod 作成時に Pod に追加のトークンを注入できる。また、以下のような環境変数もセットされる。
```json
  {
    "name": "AWS_DEFAULT_REGION",
    "value": "us-east-2"
  },
  {
    "name": "AWS_REGION",
    "value": "us-east-2"
  },
  {
    "name": "AWS_ROLE_ARN",
    "value": "arn:aws:iam::111122223333:role/eksctl-eks-oidc-demo-addon-iamserviceaccount-Role1-1SJZ3F7H39X72"
  },
  {
    "name": "AWS_WEB_IDENTITY_TOKEN_FILE",
    "value": "/var/run/secrets/eks.amazonaws.com/serviceaccount/token"
  }
```
* 処理のフロー
  * OIDC プロバイダー: JWT を発行
  * Pod 内の AWS SDK: JWT, IAM ロール ARN 情報をもとに AWS STS AssumeRoleWithWebIdentity API オペレーションを実行。
  * IAM: OIDC プロバイダーから公開鍵を取得し検証する。トークン署名を検証するための公開鍵は ```https://OIDC_PROVIDER_URL/.well-known/openid-configuration``` でホスティングされている。OK であれば、一時認証情報を返す。


