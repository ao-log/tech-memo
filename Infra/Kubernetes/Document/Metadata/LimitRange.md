
[LimitRange](https://kubernetes.io/ja/docs/concepts/policy/limit-range/)

LimitRange を使用することで、以下のような制限を課すことができる。
* ネームスペース内の Pod、コンテナごとの最小、最大のコンピュートリソースの使用量
* ネームスペース内の PersistentVolumeClaim ごとのストレージの最小、最大の使用量
* ネームスペース内のリソースについての request と limit の比率
* ネームスペース内のデフォルトの request/limit とコンテナ実行時の強制的な適用

既に作成済みの Pod に対しては効力を発揮しない。


#### リファレンス

[LimitRange v1 core](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#limitrange-v1-core)


#### サンプル

```yaml
apiVersion: v1
kind: LimitRange
metadata:
  name: cpu-min-max-demo-lr
spec:
  limits:
  - type: Container
    default:
      memory: 1Gi
    defaultRequest:
      memory: 1Gi
    max:
      memory: 1Gi
    min:
      memory: 500Mi
```

LimitRange 外の指定を行った場合は作成に失敗する。
```
Error from server (Forbidden): error when creating "examples/admin/resource/memory-constraints-pod-2.yaml":
pods "constraints-mem-demo-2" is forbidden: maximum memory usage per Container is 1Gi, but limit is 1536Mi.
```


