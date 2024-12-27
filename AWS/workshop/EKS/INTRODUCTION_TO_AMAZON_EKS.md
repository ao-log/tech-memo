
[INTRODUCTION TO AMAZON EKS](https://eks-for-aws-summit-online.workshop.aws/)

## 2. クラスターの作成

[クラスターの作成](https://catalog.us-east-1.prod.workshops.aws/workshops/f5abb693-2d87-43b5-a439-77454f28e2e7/ja-JP/020-create-cluster/20-create-cluster)

* EKS クラスターの作成
* マネージドノードグループで作成している
```shell
AWS_REGION=$(aws configure get region)
eksctl create cluster \
  --name=ekshandson \
  --version 1.30 \
  --nodes=3 --managed \
  --region ${AWS_REGION} --zones ${AWS_REGION}a,${AWS_REGION}c
```


[便利なツールと設定](https://catalog.us-east-1.prod.workshops.aws/workshops/f5abb693-2d87-43b5-a439-77454f28e2e7/ja-JP/020-create-cluster/30-configure-useful-tools) 

コマンドの補完、現在の Namespace 表示など便利ツールがいくつかあるので、導入しておくと良い。



## 3. クラスターの確認

[クラスターの確認](https://catalog.us-east-1.prod.workshops.aws/workshops/f5abb693-2d87-43b5-a439-77454f28e2e7/ja-JP/030-explore-cluster)

```shell
// クラスターの基本情報を確認
kubectl cluster-info

// クラスターに所属するノードを確認
kubectl get node

// Namespace を確認
kubectl get namespace
// デフォルトの Namespace を変更
kubens kube-system

// Pod を確認
kubectl get pod -n default

// より多くの情報を表示
kubectl get pod -n kube-system -o wide
// Pod の詳細を確認
kubectl describe pod -n kube-system ${POD_NAME}
// YAML 形式で出力
kubectl get pod -n kube-system ${POD_NAME} -o yaml
// jq コマンドと組み合わせて欲しい情報を出力
kubectl get pod -n kube-system ${POD_NAME} -o json | jq -r '.metadata.name'

// Service, Deployment, DaemonSet を確認
kubectl get service -A
kubectl get deployment -A
kubectl get daemonset -A
```



## 4. サンプルアプリのデプロイ

[サンプルアプリのデプロイ](https://catalog.us-east-1.prod.workshops.aws/workshops/f5abb693-2d87-43b5-a439-77454f28e2e7/ja-JP/040-deploy-sample-app)

* frontend, backend の構成。backend では DynamoDB でデータを永続化

#### frontend

・Deployment
```yaml
frontend_repo=$(aws ecr describe-repositories --repository-names frontend --query 'repositories[0].repositoryUri' --output text)
cat <<EOF > frontend-deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: frontend
  namespace: frontend
spec:
  selector:
    matchLabels:
      app: frontend
  replicas: 2
  template:
    metadata:
      labels:
        app: frontend
    spec:
      containers:
      - name: frontend
        image: ${frontend_repo}:latest
        imagePullPolicy: Always
        ports:
        - containerPort: 5000
        env:
        - name: BACKEND_URL
          value: http://backend.backend:5000/messages
EOF
```
* `spec.template` の部分が Pod 定義のテンプレート
* `spec.template.metadata.labels` の部分で Pod に app=frontend というラベルを付与している

・Service
```yaml
cat <<EOF > frontend-service-lb.yaml
apiVersion: v1
kind: Service
metadata:
  name: frontend
  namespace: frontend
spec:
  type: LoadBalancer
  selector:
    app: frontend
  ports:
  - protocol: TCP
    port: 80
    targetPort: 5000
EOF
```
* `spec.type` で LoadBalancer を指定
* `spec.selector` の部分で指定されたラベル (app=frontend) を持つ Pod がこの Service からの割り振り対象となる

・ログの確認
```
kubectl logs -n frontend <Pod名>
```


#### backend

・Deployment
```yaml
AWS_REGION=$(aws configure get region)
backend_repo=$(aws ecr describe-repositories --repository-names backend --query 'repositories[0].repositoryUri' --output text)
cat <<EOF > backend-deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: backend
  namespace: backend
spec:
  selector:
    matchLabels:
      app: backend
  replicas: 2
  template:
    metadata:
      labels:
        app: backend
    spec:
      containers:
      - name: backend
        image: ${backend_repo}:latest
        imagePullPolicy: Always
        ports:
        - containerPort: 5000
        env:
        - name: AWS_DEFAULT_REGION
          value: ${AWS_REGION}
        - name: DYNAMODB_TABLE_NAME
          value: messages
EOF
```

・Service
```yaml
cat <<EOF > backend-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: backend
  namespace: backend
spec:
  type: ClusterIP
  selector:
    app: backend
  ports:
  - protocol: TCP
    port: 5000
    targetPort: 5000
EOF
```
* 同一 Namespace の Pod からは、`metadata.name` で指定されている Service の名前でこの Service にアクセスできる
* 別の Namespace の Pod からは、<Service 名>.<Namespace 名> でこの Service にアクセスできる


#### IRSA

* IRSA 用のサービスアカウント、IAM ロールを作成
```shell
AWS_ACCOUNT_ID=$(aws sts get-caller-identity --output text --query Account)
eksctl create iamserviceaccount \
  --name backend-dynamodb-access \
  --namespace backend \
  --cluster ekshandson \
  --attach-policy-arn arn:aws:iam::${AWS_ACCOUNT_ID}:policy/backend-dynamodb-access \
  --approve
```
* backend の Deployment に以下追加。
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: backend
  namespace: backend
spec:
  selector:
    matchLabels:
      app: backend
  replicas: 2
  template:
    metadata:
      labels:
        app: backend
    spec:
+     serviceAccountName: backend-dynamodb-access
      containers:
      ...
```


