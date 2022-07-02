
[ReplicaSet](https://kubernetes.io/ja/docs/concepts/workloads/controllers/replicaset/)

* Deployment を使用することを推奨(Deployment の内部で ReplicaSet が使用されている)
* 指定した Pod 数を維持する。
* ```.spec.selector.matchLabels``` が Pod 数を維持する対象。


#### サンプル

```yaml
apiVersion: apps/v1
kind: ReplicaSet
metadata:
  name: frontend
  labels:
    app: guestbook
    tier: frontend
spec:
  replicas: 3
  selector:
    matchLabels:  # Pod 数を維持する対象
      tier: frontend
    matchExpressions:
      - {key: tier, operator: In, values: [frontend]}
  template:
    metadata:
      labels:
        app: guestbook
        tier: frontend
    spec:
      containers:
      - name: php-redis
        image: gcr.io/google_samples/gb-frontend:v3
        resources:
          requests:
            cpu: 100m
            memory: 100Mi
        env:
        - name: GET_HOSTS_FROM
          value: dns
        ports:
        - containerPort: 80
```

