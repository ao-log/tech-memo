
[Ingress](https://kubernetes.io/ja/docs/concepts/services-networking/ingress/)

* Service は L4 で動作。Ingress は L7。
* Nginx Ingress Controller や各クラウドプロバイダーのロードバランサに対応したコントローラーがある。
  * クラウドプロバイダーのロードバランサはクラスタ外からのリクエストも受け付けることができる。ロードバランサからノードポートに転送する仕組み。
  * Nginx Ingress Controller はクラスタ内のトラフィック転送用途。Ingress 用の Pod を起動する。クラスタ外部から通信したい場合は更に Service → Ingress 用の Pod といった経路が必要。
* バックエンドとして Service が必要。Service は NodePort にする必要がある。


## EKS の場合

[Amazon EKS でのアプリケーション負荷分散](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/alb-ingress.html)

* AWS Load Balancer Controller が必要
* ワーカーノードのセキュリティグループの一つにタグ付けが必要
  * キー: kubernetes.io/cluster/cluster-name
  * 値: shared または owned
* サブネットにタグ付けが必要
  * プライベートサブネットの場合
    * キー: kubernetes.io/role/internal-elb
    * 値: 1
  * パブリックサブネットの場合
    * キー: kubernetes.io/role/elb
    * 値: 1

[AWS Load Balancer Controller アドオンのインストール](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/aws-load-balancer-controller.html)

* IAM ポリシーには AWSLoadBalancerControllerIAMPolicy のアタッチが必要
* IAM ロールを作成。信頼ポリシーは以下内容。EKS クラスターの OIDC プロバイダーからのみ許可する。

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "Federated": "arn:aws:iam::111122223333:oidc-provider/oidc.eks.region-code.amazonaws.com/id/EXAMPLED539D4633E53DE1B716D3041E"
            },
            "Action": "sts:AssumeRoleWithWebIdentity",
            "Condition": {
                "StringEquals": {
                    "oidc.eks.region-code.amazonaws.com/id/EXAMPLED539D4633E53DE1B716D3041E:sub": "system:serviceaccount:kube-system:aws-load-balancer-controller"
                }
            }
        }
    ]
}
```

サービスアカウントは以下の内奥。アノテーションで IAM ロールを指定する。
```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    app.kubernetes.io/component: controller
    app.kubernetes.io/name: aws-load-balancer-controller
  name: aws-load-balancer-controller
  namespace: kube-system
  annotations:
    eks.amazonaws.com/role-arn: arn:aws:iam::111122223333:role/AmazonEKSLoadBalancerControllerRole
```

* helm もしくはマニフェストファイルから AWS Load Balancer Controller をインストール。


#### サンプル

[テンプレート例](https://raw.githubusercontent.com/kubernetes-sigs/aws-load-balancer-controller/v2.4.0/docs/examples/2048/2048_full.yaml)

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
