
こちらの手順に習ってGKE の動作を確認しました。
https://github.com/GoogleCloudPlatform/gke-gobang-app-example

以下の手順は、上記 URL のものを写経したものになります。

# 環境変数の設定

```
export PROJECT_ID=$(gcloud config list project --format "value(core.project)")
```

# コンテナクラスタの作成

### クラスタ作成

[コンピューティング] → [Kubernetes Engine]
「クラスタを作成」をクリック
今回は各パラメータをデフォルトのままとし、「作成」をクリック。

**Memo**  
クラスタ作成はコマンドラインからもできるようですが、ここでは管理コンソール上から実施しています。

### 設定ファイルの取得

kubectl を利用できるようにするために必要。

```shell
$ gcloud container clusters get-credentials gobang-cluster --zone=us-central1-a
Fetching cluster endpoint and auth data.
kubeconfig entry generated for gobang-cluster.
```

### ノードの確認

3 台のノードが稼働していることを確認します。

```
$ kubectl get nodes
NAME                                            STATUS    ROLES     AGE       VERSION
gke-gobang-cluster-default-pool-d255e238-gsmr   Ready     <none>    5m        v1.9.7-gke.3
gke-gobang-cluster-default-pool-d255e238-qnh9   Ready     <none>    5m        v1.9.7-gke.3
gke-gobang-cluster-default-pool-d255e238-x5x7   Ready     <none>    5m        v1.9.7-gke.3
```

# Docker イメージ

### ビルド

```
$ docker build -t frontend:v1.0 frontend/
$ docker build -t backend:v1.0 backend-dummy/
$ docker build -t backend:v1.1 backend-smart/
```

### Container Registory 用の別名をつける

```
$ docker tag frontend:v1.0 gcr.io/$PROJECT_ID/frontend:v1.0
$ docker tag backend:v1.0 gcr.io/$PROJECT_ID/backend:v1.0
$ docker tag backend:v1.1 gcr.io/$PROJECT_ID/backend:v1.1
```

### Container Registory push

```
$ gcloud docker -- push gcr.io/$PROJECT_ID/frontend:v1.0
$ gcloud docker -- push gcr.io/$PROJECT_ID/backend:v1.0
$ gcloud docker -- push gcr.io/$PROJECT_ID/backend:v1.1
```

v1.1 については、v1.0 と差分のあるレイヤ分だけ push する動きになります。

# Deployment

デプロイします。

```
$ kubectl create -f config/frontend-deployment.yaml
deployment "frontend-node" created
$ kubectl create -f config/backend-deployment.yaml
deployment "backend-node" created
```

backend-deployment.yaml は以下の内容になっています。
メタデータのほか、コンテナをどう稼働させるかの情報（イメージ、稼働ポート、レプリカ数）を定義しています。

```yaml
kind: Deployment
metadata:
  labels:
    name: backend-node
  name: backend-node
spec:
  replicas: 3
  template:
    metadata:
      labels:
        name: backend-node
    spec:
      containers:
      - image: gcr.io/<PROJECT ID>/backend:v1.0
        name: backend-node
        ports:
        - containerPort: 8081
```

pod の稼働状況を確認します。

```
$ kubectl get pods
NAME                             READY     STATUS    RESTARTS   AGE
backend-node-7469fb6f64-7mwt8    1/1       Running   0          1m
backend-node-7469fb6f64-bwssd    1/1       Running   0          1m
backend-node-7469fb6f64-ng8ms    1/1       Running   0          1m
frontend-node-58d8f858d9-2xwd7   1/1       Running   0          1m
frontend-node-58d8f858d9-dvhqx   1/1       Running   0          1m
frontend-node-58d8f858d9-jt5f5   1/1       Running   0          1m
```

# Service

```
$ kubectl create -f config/frontend-service.yaml
service "frontend-service" created
$ kubectl create -f config/backend-service.yaml
service "backend-service" created
```

frontend-service.yaml は以下の内容になっています。type に LoadBalancer を指定することで、Cloud Load Balancing 経由でアクセスできるようになります。

```yaml
$ cat config/frontend-service.yaml
apiVersion: v1
kind: Service
metadata:
  labels:
    name: frontend-service
  name: frontend-service
spec:
  ports:
  - protocol: TCP
    port: 80
    targetPort: 8080
  selector:
    name: frontend-node
  type: LoadBalancer
```

```
$ kubectl get services
NAME               TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
backend-service    ClusterIP      10.43.248.250   <none>        8081/TCP       32s
frontend-service   LoadBalancer   10.43.255.75    <pending>     80:31753/TCP   43s
kubernetes         ClusterIP      10.43.240.1     <none>        443/TCP        1h
```

# ローリングアップデート

backend を v1.0 から v1.1 にアップデートします。yaml のイメージ指定を
 v1.1 にしたものを apply します。

```
$ kubectl apply -f config/backend-deployment-v1_1.yaml
```

次のコマンドで履歴を確認できます。

```
$ kubectl describe deployment/backend-node
```


# 削除

```
$ kubectl delete service frontend-service
$ kubectl delete service backend-service

$ kubectl delete deployment frontend-node
$ kubectl delete deployment backend-node
```

# 参考

https://cloud.google.com/kubernetes-engine/docs/?hl=JA
https://github.com/GoogleCloudPlatform/gke-gobang-app-example
