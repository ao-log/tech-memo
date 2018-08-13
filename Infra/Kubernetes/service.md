公式ページを読みつつ、気になる箇所をつまみ食いしたメモです。
https://kubernetes.io/docs/concepts/services-networking/service/

### サービスの定義

```yaml
kind: Service
apiVersion: v1
metadata:
  name: my-service
spec:
  selector:
    app: MyApp
  ports:
  - protocol: TCP
    port: 80
    targetPort: 9376
```

| 項目 | 説明 |
| --- | --- |
| サービス名 | my-services |
| selector | MyApp ラベルのついた Pod にトラフィックを転送する |
| TargetPort | Pod に転送する際の宛先ポート |
