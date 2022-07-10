
[Horizontal Pod Autoscaling](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/)

負荷に応じて Pod 数をオートスケーリングできる。

使用方法は以下のドキュメントにて案内されている。

[Horizontal Pod Autoscalerウォークスルー](https://kubernetes.io/ja/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/)


#### EKS ドキュメント

[Horizontal Pod Autoscaler](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/horizontal-pod-autoscaler.html)


```
// Apache のコンテナをデプロイ。Service「php-apache」もあわせて作成。
kubectl apply -f https://k8s.io/examples/application/php-apache.yaml

// HPA リソースを作成。CPU 使用率 50 % を追跡
kubectl autoscale deployment php-apache --cpu-percent=50 --min=1 --max=10

// 負荷をかけるテスト
kubectl run -i \
    --tty load-generator \
    --rm --image=busybox \
    --restart=Never \
    -- /bin/sh -c "while sleep 0.01; do wget -q -O- http://php-apache; done"

// HPA リソースの確認
kubectl get hpa php-apache
```
