
[ResourceQuota](https://kubernetes.io/ja/docs/concepts/policy/resource-quotas/)

ネームスペース内で使用可能なリソース数、リソース量の合計を制限できる。


#### リファレンス

[ResourceQuota v1 core](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#resourcequota-v1-core)


#### サンプル

リソース数の制限。

```yaml
apiVersion: v1
kind: ResourceQuota
metadata:
  name: object-counts
spec:
  hard:
    configmaps: "10"
    persistentvolumeclaims: "4"
    pods: "4"
    replicationcontrollers: "20"
    secrets: "10"
    services: "10"
    services.loadbalancers: "2"
```

リソース量の制限。

```yaml
apiVersion: v1
kind: ResourceQuota
metadata:
  name: compute-resources
spec:
  hard:
    requests.cpu: "1"
    requests.memory: 1Gi
    limits.cpu: "2"
    limits.memory: 2Gi
    requests.nvidia.com/gpu: 4
```    


#### クォータのスコープ

スコープを設定している場合は、以下に一致するリソースを対象に使用量の合計値を計測する。

* Terminating: .spec.activeDeadlineSeconds >= 0 である Pod に一致。
* NotTerminating: .spec.activeDeadlineSeconds が nil である Pod に一致。
* BestEffort: BestEffort(requests/limits 未指定) の QoS の Pod に一致。
* NotBestEffort: BestEffor ではない QoS の Pod に一致。
* PriorityClass: 指定された優先度クラスと関連付いている Pod に一致。

[ScopeSelector v1 core](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#scopeselector-v1-core)

サンプル。

```yaml
  scopeSelector:
    matchExpressions:
      - scopeName: PriorityClass
        operator: In
        values:
          - middle
```


#### PriorityClass

scopeSelector に一致する Pod が集計対象となる。

サンプル。

```yaml
apiVersion: v1
kind: List
items:
- apiVersion: v1
  kind: ResourceQuota
  metadata:
    name: pods-high
  spec:
    hard:
      cpu: "1000"
      memory: 200Gi
      pods: "10"
    scopeSelector:
      matchExpressions:
      - operator : In
        scopeName: PriorityClass
        values: ["high"]
```

Pod での指定。

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: high-priority
spec:
  containers:
  - name: high-priority
    image: ubuntu
    command: ["/bin/sh"]
    args: ["-c", "while true; do echo hello; sleep 10;done"]
    resources:
      requests:
        memory: "10Gi"
        cpu: "500m"
      limits:
        memory: "10Gi"
        cpu: "500m"
  priorityClassName: high
```


