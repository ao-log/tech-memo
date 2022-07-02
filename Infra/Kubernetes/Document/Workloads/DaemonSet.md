
[DaemonSet](https://kubernetes.io/ja/docs/concepts/workloads/controllers/daemonset/)

* 各ノード上に 1 つずつ Pod を起動。
* ```.spec.template.spec.nodeSelector``` を設定した場合は、マッチするノード上に Pod を起動。
* ```.spec.template.spec.affinity``` を設定した場合は、マッチするノード上に Pod を起動。
* ```.spec.updateStrategy``` に従って Pod を更新する。OnDelete の場合は Pod が停止して再起動されない限りは Pod が更新されない。


#### サンプル

```yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: fluentd-elasticsearch
  namespace: kube-system
  labels:
    k8s-app: fluentd-logging
spec:
  selector:
    matchLabels:
      name: fluentd-elasticsearch
  template:
    metadata:
      labels:
        name: fluentd-elasticsearch
    spec:
      tolerations:
      - key: node-role.kubernetes.io/master
        operator: Exists
        effect: NoSchedule
      containers:
      - name: fluentd-elasticsearch
        image: quay.io/fluentd_elasticsearch/fluentd:v2.5.2
        resources:
          limits:
            memory: 200Mi
          requests:
            cpu: 100m
            memory: 200Mi
        volumeMounts:
        - name: varlog
          mountPath: /var/log
        - name: varlibdockercontainers
          mountPath: /var/lib/docker/containers
          readOnly: true
      terminationGracePeriodSeconds: 30
      volumes:
      - name: varlog
        hostPath:
          path: /var/log
      - name: varlibdockercontainers
        hostPath:
          path: /var/lib/docker/containers
```
