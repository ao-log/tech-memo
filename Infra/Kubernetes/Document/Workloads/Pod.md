
[Pod](https://kubernetes.io/ja/docs/concepts/workloads/pods/)

* 複数のコンテナを含めることができる。コンテナ間は localhost で通信可能。


#### リファレンス

[Pod v1 core](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#podspec-v1-core)


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


#### Pod のライフサイクル

・Pod のフェーズ  
[Pod のフェーズ](https://kubernetes.io/ja/docs/concepts/workloads/pods/pod-lifecycle/#pod-phase) に表でまとめられている。

* Pending: Pod の起動を試みている状態
* Running: Pod が稼働している状態
* Suceeded: 正常終了
* Failed: 異常終了
* Unknown: Pod の状態を取得できていない

・コンテナのステータス  
[コンテナのステータス](https://kubernetes.io/ja/docs/concepts/workloads/pods/pod-lifecycle/#container-states) にまとめられている。

* Waiting
* Running
* Terminated

・Pod のステータス  
[Pod の Condition](https://kubernetes.io/ja/docs/concepts/workloads/pods/pod-lifecycle/#pod-conditions) にまとめられている。

[PodCondition v1 core](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#podcondition-v1-core) を参照のこと。

* type: PodScheduled, ContainersReady, Initialized or Ready
* status: True, False or Unknown
* lastProbeTime
* lastTransitionTime
* reason: 機械可読用のメッセージ
* message: human readable なメッセージ

・Pod の終了  
[Pod の終了](https://kubernetes.io/ja/docs/concepts/workloads/pods/pod-lifecycle/#pod-termination) に詳しく書かれている。

* preStop が設定されている場合は実行される。
* kubelet からコンテナランタイムをトリガーし、各コンテナの PID 1 に SIGTERM を送信する。
* コントロールプレーンが Endpoint オブジェクトから終了中の Pod を削除する。
* 猶予期間が過ぎると、kubelet からコンテナランタイムをトリガーし、各コンテナの実行中のプロセスに SIGKILL を送信する。
* API Server から Pod のオブジェクト情報を削除する。

```kubectl delete pod``` を実行する際、```--grace-period``` オプションに 0 を指定すると即座に強制終了する。


#### preStop

Pod の削除時は preStop を実施した後に SIGTERM を送信できる。デフォルトでは preStop + SIGTERM が 30 秒以内に終了しなかった場合に SIGKILL が送信される。
サービスの除外後に SIGTERM という順番が保証されるわけではないため、preStop で数秒 sleep するような対応が必要。

* [Podの終了](https://kubernetes.io/ja/docs/concepts/workloads/pods/pod-lifecycle/#pod-termination)
* [コンテナフック](https://kubernetes.io/ja/docs/concepts/containers/container-lifecycle-hooks/#hook-details)


#### ガベージコレクション

* [ガベージコレクション](https://kubernetes.io/ja/docs/concepts/workloads/controllers/garbage-collection/)
* [失敗したPodのガベージコレクション](https://kubernetes.io/ja/docs/concepts/workloads/pods/pod-lifecycle/#pod-garbage-collection)

従属オブジェクト(ReplicaSet により起動された Pod など) は metadata.ownerReferences フィールドを持つ。

Background, Forground の二つのモードがある。
Background の場合は ガベージコレクタがバックグラウンドで非同期に Pod を削除する。


#### 再起動ポリシー

restartPolicy で設定。

* Always: 常に再起動
* OnFailure: 終了コード 0 以外の場合に再起動
* Never: 再起動しない

Pod 内の全てのコンテナが適用対象となる。
Pod 内のコンテナが終了しああとコンテナの再起動を行う。
連続した場合は 5 分を上限とし指数バックオフ遅延(10秒, 20秒, 40秒, ...)を行う。
ただしコンテナが 10 分間稼働するとバックオフタイマーをリセットする。


#### ヘルスチェック

[Pod のライフサイクル](https://kubernetes.io/ja/docs/concepts/workloads/pods/pod-lifecycle/)

3 種類の Probe がある。
* livenessProbe: コンテナが稼働しているかどうかを判定。失敗時は restartPolicy に従う。
* readinessProbe: コンテナがリクエストに応答できるかどうかを判定。失敗時は当該 Pod を Service から除外。
* startupProbe: コンテナ内のアプリケーションが起動したかどうかを判定。失敗時は restartPolicy に従う。

診断方法は以下の通り。

* ExecAction: コンテナ内で特定のコマンドを実行。ステータス 0 を成功と判定。
* TCPSocketAction: Pod の IP の特定のポートに TCP チェックを行う。TCP コネクションを確立できる場合に成功と判定。
* HTTPGetAction: Pod の IP の特定のポートとパスに対して、HTTP GET のリクエストを送信。 レスポンスのステータスコードが 200 以上400 未満の際に成功と判定。

使用方法は [Liveness Probe、Readiness ProbeおよびStartup Probeを使用する](https://kubernetes.io/ja/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/) を参照するとよい。

HTTP GET による Liveness Probe の例。[ヘルスチェック用のテストイメージ - server.go](https://github.com/kubernetes/kubernetes/blob/master/test/images/agnhost/liveness/server.go) では /healthz に対して初めの 10 秒は 200, 以降は 500 を返すように実装されている。
readinessProbeの場合も livenessProbe と同様のプロパティとなっている。

```yaml
apiVersion: v1
kind: Pod
metadata:
  labels:
    test: liveness
  name: liveness-http
spec:
  containers:
  - name: liveness
    image: k8s.gcr.io/liveness
    args:
    - /server
    livenessProbe:
      httpGet:
        path: /healthz
        port: 8080
        httpHeaders:
        - name: X-Custom-Header
          value: Awesome
      initialDelaySeconds: 3
      periodSeconds: 3
```


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


#### requests, limits

[コンテナのリソース管理](https://kubernetes.io/ja/docs/concepts/configuration/manage-resources-containers/)

以下の指定により CPU, Memory のリソース量を設定できる。requests はノードが確保する量。limits は使用できる上限。

* spec.containers[].resources.limits.cpu
* spec.containers[].resources.limits.memory
* spec.containers[].resources.requests.cpu
* spec.containers[].resources.requests.memory

requests を超えてリソースを使用することもできる。オーバーコミットしているといえる。しかし、特にノードの CPU を使い果たさないように limits は程々の値に抑えるようにしたほうがいい。


