
[Pod](https://kubernetes.io/ja/docs/concepts/workloads/pods/)

* 複数のコンテナを含めることができる。コンテナ間は localhost で通信可能。
* init コンテナを使用できる


#### リファレンス

[Reference - Pod v1 core](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#podspec-v1-core)


#### サンプル

```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: hello
spec:
  template:
    # これがPodテンプレートです
    spec:
      containers:
      - name: hello
        image: busybox
        command: ['sh', '-c', 'echo "Hello, Kubernetes!" && sleep 3600']
      restartPolicy: OnFailure
    # Podテンプレートはここまでです
```


[Pod のライフサイクル](https://kubernetes.io/ja/docs/concepts/workloads/pods/pod-lifecycle/)

restartPolicy は Always、OnFailure、または Never。

3 種類の Probe がある。
* livenessProbe
* readinessProbe
* startupProbe

詳細は、[Liveness Probe、Readiness ProbeおよびStartup Probeを使用する](https://kubernetes.io/ja/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/) を参照のこと。


**preStop**:  
Pod の削除時は preStop を実施した後に SIGTERM を送信できる。デフォルトでは preStop + SIGTERM が 30 秒以内に終了しなかった場合に SIGKILL が送信される。
サービスの除外後に SIGTERM という順番が保証されるわけではないため、preStop で数秒 sleep するような対応が必要。

* [Podの終了](https://kubernetes.io/ja/docs/concepts/workloads/pods/pod-lifecycle/#pod-termination)
* [コンテナフック](https://kubernetes.io/ja/docs/concepts/containers/container-lifecycle-hooks/#hook-details)


#### Init コンテナ

[Init コンテナ](https://kubernetes.io/ja/docs/concepts/workloads/pods/init-containers/)

Init コンテナは上に書いたものから順に実行される。

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: myapp-pod
  labels:
    app: myapp
spec:
  containers:
  - name: myapp-container
    image: busybox:1.28
    command: ['sh', '-c', 'echo The app is running! && sleep 3600']
  initContainers:
  - name: init-myservice
    image: busybox:1.28
    command: ['sh', '-c', "until nslookup myservice.$(cat /var/run/secrets/kubernetes.io/serviceaccount/namespace).svc.cluster.local; do echo waiting for myservice; sleep 2; done"]
  - name: init-mydb
    image: busybox:1.28
    command: ['sh', '-c', "until nslookup mydb.$(cat /var/run/secrets/kubernetes.io/serviceaccount/namespace).svc.cluster.local; do echo waiting for mydb; sleep 2; done"]
```


#### 環境変数

[.spec.containers.env](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#envvar-v1-core)

```yaml
    env:
      - name: TESTENV
        value: TESTVALUE
```

[valueFrom](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#envvarsource-v1-core)

シークレットを使用した場合。

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: secret-env-pod
spec:
  containers:
  - name: mycontainer
    image: redis
    env:
      - name: SECRET_USERNAME
        valueFrom:
          secretKeyRef:
            name: mysecret
            key: username
      - name: SECRET_PASSWORD
        valueFrom:
          secretKeyRef:
            name: mysecret
            key: password
  restartPolicy: Never
```
