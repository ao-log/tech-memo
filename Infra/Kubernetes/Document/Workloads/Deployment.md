
[Deployment](https://kubernetes.io/ja/docs/concepts/workloads/controllers/deployment/)

* ReplicaSet を使用してローリングアップデートが可能。
* ```spec.selector.matchLabels``` は Deployment が管理する対象のラベル。
* Pod 更新時は、デフォルトで 75 % 〜 125 % の稼働状態を維持する。


#### サンプル

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.14.2
        ports:
        - containerPort: 80
```


#### ローリングアップデート

流れ
* Deployment が新しい ReplicaSet を作成
* Deployment が新しい ReplicaSet をスケールアップ
* Deployment が古い ReplicaSet をスケールダウン

ローリングアップデートのステータス確認。
```
kubectl rollout status deployment <Deployment 名>
```

ロールバック。
```
// 過去のローリングアップデートの履歴を確認。
kubectl rollout history deployment <Deployment 名>

// 過去のバージョンにロールバック。
kubectl rollout undo deployment <Deployment 名> --to-revision <リビジョン番号>
```


#### Deployment 更新の失敗要因

* 不十分なリソースの割り当て
* ReadinessProbe の失敗
* コンテナイメージの取得ができない
* 不十分なパーミッション
* リソースリミットのレンジ
* アプリケーションランタイムの設定の不備


#### 更新戦略

* ```.spec.strategy.type```: アップデート方法
* ```.spec.strategy.rollingUpdate.maxUnavailable```: 利用不可を許容する Pod 数
* ```.spec.strategy.rollingUpdate.maxSurge```: Desired 数よりも多く起動する際に許容する Pod 数
