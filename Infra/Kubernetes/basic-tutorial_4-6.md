こちらのチュートリアルを行いました。
https://kubernetes.io/docs/tutorials/kubernetes-basics/

その時のメモです。
このページでは、4 〜 6 章分を記載しています。


# Using Service to Expose Your App

Pod はいつかは停止する運命にあります。Pod はライフサイクルを持っています。例えば、ワーカーノードがダウンした時、ノード上で稼働する Pod もまた消失します。レプリケーションコントローラは動的にクラスタをあるべき姿に戻し、新しい Pod を作成することで、アプリケーションの稼働状況を保ちます。別の例を出すと、3 つのレプリカをバックエンドに作成する例を考えてみてください。フロントエンドシステムはバックエンドレプリカの状況を把握する必要はありません。

Service は Pod の論理的な組を定義し、どのようにアクセスするかを定義します。Service は Pod 間のゆるい結合を可能にします。Service は YAML もしくは JSON によって定義されます。Pod の組は LabelSelector によって決定されます。

各 Pod がユニークな IP アドレスを持ちますが、それらの IP アドレスは Service なしではクラスタ外部に公開されません。Service によって、あなたのアプリケーションはトラフィックを受信できるようになります。Service は ServiceSpec のタイプによって公開することもできます。

* ClusterIP (default) - クラスタ内のインターナル IP アドレスでサービスを公開します。このタイプはクラスタ内部のみ通信が届きます。

* NodePort - 選択されたノード上で NAT により同一のポートでサービスを公開します。サービスは <NodeIP>:<NodePort> でアクセス可能になります。

* LoadBalancer - 外部のロードバランサを作成し、Service とグローバルアドレスをひも付けます。

* ExternalName - 任意の名前でサービスを公開します。kube-dns の v1.7 以上が必要になります。

より詳細な情報は Using Source IP tutorial をご参照ください。あるいは、Connecting Applications with Services をご参照ください。

Service がセレクタなしで作られるユースケースが幾つかあります。セレクタなしで作られるということはエンドポイントなしで作られるということです。このことは手動でマッピングできるということです。もう一つのセレクタのない例は、ExternalName を使う場合です。


Service は Pod の組へトラフィックをルーティングします。Service は抽象化されているので、Pod の稼働状況はアプリケーションに影響を与えません。関連する Pod のディスカバリ、ルーティングは Kubernetes Service によってハンドリングされます。

Service は Pod の組とのマッチングをラベルとセレクタによって行います。グルーピングによって、オブジェクトへの論理的な操作が可能になります。ラベルはキーバリューのペアになっていて、オブジェクトに設定され、様々な方法で利用できます。

* オブジェクトの用途指定（development, test, production）
* バージョン、タグの埋め込み
* タグによるオブジェクトの分類

ラベルはオブジェクト生成時に設定されます。いつでも変更可能です。

### チュートリアル

Service の一覧を確認します。これは Minikube 起動時にデフォルトで作られるものです。

```
$ kubectl get services
NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   33s
```

新たな Service を作ります。type は NodePort にしています。（minikube は現時点において LoadBalancer をサポートしていません）

```
$ kubectl expose deployment/kubernetes-bootcamp --type="NodePort" --port 8080
service "kubernetes-bootcamp" exposed
```

先ほど追加した Service が追加されています。

```
$ kubectl get services
NAME                  TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE
kubernetes            ClusterIP   10.96.0.1      <none>        443/TCP          2m
kubernetes-bootcamp   NodePort    10.103.61.38   <none>        8080:30186/TCP   43s
```

外部に公開されているポートを確認するには kubectl describe を実行します。NodePort が外部公開されているポートです。

```
$ kubectl describe services/kubernetes-bootcamp
Name:                     kubernetes-bootcamp
Namespace:                default
Labels:                   run=kubernetes-bootcamp
Annotations:              <none>
Selector:                 run=kubernetes-bootcamp
Type:                     NodePort
IP:                       10.101.199.63
Port:                     <unset>  8080/TCP
TargetPort:               8080/TCP
NodePort:                 <unset>  30486/TCP
Endpoints:                172.18.0.4:8080
Session Affinity:         None
External Traffic Policy:  Cluster
Events:                   <none>
```

環境変数 NODE_PORT にを設定します。

```
$ export NODE_PORT=$(kubectl get services/kubernetes-bootcamp -o go-template='{{(index .spec.ports 0).nodePort}}')

$ echo NODE_PORT=$NODE_PORT
NODE_PORT=30486
```

NodePort にアクセスできるようになっているかどうかテストします。
確かにアクセスできています。

```
$ curl $(minikube ip):$NODE_PORT
Hello Kubernetes bootcamp! | Running on: kubernetes-bootcamp-5dbf48f7d4-4xdhh | v=1
```

Deployment にはラベルがつけられています。

```
$ kubectl describe deployment
Name:                   kubernetes-bootcamp
Namespace:              default
CreationTimestamp:      Sat, 11 Aug 2018 13:10:45 +0000
Labels:                 run=kubernetes-bootcamp
Annotations:            deployment.kubernetes.io/revision=1
...
```

-l で指定したラベルのみに出力を絞ることができます。

```
$ kubectl get pods -l run=kubernetes-bootcamp
NAME                                   READY     STATUS    RESTARTS   AGE
kubernetes-bootcamp-5dbf48f7d4-fflp2   1/1       Running   0          2m

$ kubectl get services -l run=kubernetes-bootcamp
NAME                  TYPE       CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
kubernetes-bootcamp   NodePort   10.105.163.110   <none>        8080:30720/TCP   2m
```

環境変数に pod 名を格納しておきます。

```
$ export POD_NAME=$(kubectl get pods -o go-template --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')
$ echo Name of the Pod: $POD_NAME
Name of the Pod: kubernetes-bootcamp-5dbf48f7d4-4xdhh
```

新しいラベルをつけることができます。

```
$ kubectl label pod $POD_NAME app=v1
pod "kubernetes-bootcamp-5dbf48f7d4-4xdhh" labeled
```

確かに新しいラベルがつけられています。

```
$ kubectl describe pods $POD_NAME
Name:           kubernetes-bootcamp-5dbf48f7d4-4xdhh
Namespace:      default
Node:           host01/172.17.0.11
Start Time:     Sat, 11 Aug 2018 11:58:06 +0000
Labels:         app=v1
                pod-template-hash=1869049380
                run=kubernetes-bootcamp
Annotations:    <none>
Status:         Running
IP:             172.18.0.4
Controlled By:  ReplicaSet/kubernetes-bootcamp-5dbf48f7d4
...
```

Service の削除は、ラベルで指定することができます。

```
$ kubectl delete service -l run=kubernetes-bootcamp
service "kubernetes-bootcamp" deleted
```

サービスの一覧を確認すると、確かに削除されています。

```
$ kubectl get services
NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   8m
```

接続できなくなっています。

```
$ curl $(minikube ip):$NODE_PORT
curl: (7) Failed to connect to 172.17.0.11 port 30486: Connection refused
```

ただし、アプリケーションは稼働したままです。次の通り、コンテナに接続することで確認できます。

```
$ kubectl exec -ti $POD_NAME curl localhost:8080
Hello Kubernetes bootcamp! | Running on: kubernetes-bootcamp-5dbf48f7d4-4xdhh | v=1
```

# Running Multiple Instances of Your App

### 概要

前章では、Deployment を作成しました。1 つの pod のみ稼働させていましたが、トラフィックが増えてきたらアプリケーションをスケールさせる必要があります。スケーリングはレプリカ数を増減させることで実現します。

Deployment の scale out は新しい Pod が作成され、スケジュールされることを保証します。スケーリングは Pod 数を増やすことで、あるべき状態へとします。Kubernetes はオートスケールもサポートしています。しかし、これはこのチュートリアルのスコープ外です。0 にスケールすることも可能です。これは、全ての pod を停止させることを意味します。

複数のインスタンスを起動する場合は、トラフィックの分散方法が必要になります。Service はロードバランサを持っていて、トラフィックを全ての Pod に分散させます。Service は稼働中の Pod をエンドポイントを使うことでモニタリングし、利用可能な Pod にのみトラフィックが送信されていることを保証します。

ローリングアップデートもダウンタイムなしに実現できます。

### チュートリアル

pod を 1 から 4 個に scale out します。

```
$ kubectl get deployments
NAME                  DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
kubernetes-bootcamp   1         1         1            0           7s

$ kubectl scale deployments/kubernetes-bootcamp --replicas=4
deployment "kubernetes-bootcamp" scaled

$ kubectl get deployments
NAME                  DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
kubernetes-bootcamp   4         4         4            4           53s

$ kubectl get pods -o wide
NAME                                   READY     STATUS    RESTARTS   AGE       IP       NODE
kubernetes-bootcamp-5dbf48f7d4-26vrz   1/1       Running   0          1m        172.18.0.7   host01
kubernetes-bootcamp-5dbf48f7d4-bdfgn   1/1       Running   0          1m        172.18.0.5   host01
kubernetes-bootcamp-5dbf48f7d4-qvj25   1/1       Running   0          2m        172.18.0.4   host01
kubernetes-bootcamp-5dbf48f7d4-s542c   1/1       Running   0          1m        172.18.0.6   host01
```

履歴を見ると、スケールしていることが記録されています。

```
$ kubectl describe deployments/kubernetes-bootcamp
Name:                   kubernetes-bootcamp
...
Type    Reason             Age   From                   Message
----    ------             ----  ----                   -------
Normal  ScalingReplicaSet  3m    deployment-controller  Scaled up replica set kubernetes-bootcamp-5dbf48f7d4 to 1
Normal  ScalingReplicaSet  2m    deployment-controller  Scaled up replica set kubernetes-bootcamp-5dbf48f7d4 to 4
```

Service を見ると、エンドポイントに 4 個の Pod の IP アドレスが設定されています。そのため、トラフィックがスケールした各 Pod に届いていると分かります。

```
$ kubectl describe services/kubernetes-bootcamp
Name:                     kubernetes-bootcamp
Namespace:                default
Labels:                   run=kubernetes-bootcamp
Annotations:              <none>
Selector:                 run=kubernetes-bootcamp
Type:                     NodePort
IP:                       10.106.170.139
Port:                     <unset>  8080/TCP
TargetPort:               8080/TCP
NodePort:                 <unset>  32213/TCP
Endpoints:                172.18.0.4:8080,172.18.0.5:8080,172.18.0.6:8080 + 1 more...
Session Affinity:         None
External Traffic Policy:  Cluster
Events:                   <none>
```

アクセスするたびに、異なる Pod に接続していることが分かります。

```
$ export NODE_PORT=$(kubectl get services/kubernetes-bootcamp -o go-template='{{(index .spec.ports 0).nodePort}}')
$ echo NODE_PORT=$NODE_PORT
NODE_PORT=32213

$ curl $(minikube ip):$NODE_PORT
Hello Kubernetes bootcamp! | Running on: kubernetes-bootcamp-5dbf48f7d4-26vrz | v=1
$ curl $(minikube ip):$NODE_PORT
Hello Kubernetes bootcamp! | Running on: kubernetes-bootcamp-5dbf48f7d4-bdfgn | v=1
$ curl $(minikube ip):$NODE_PORT
Hello Kubernetes bootcamp! | Running on: kubernetes-bootcamp-5dbf48f7d4-s542c | v=1
```

スケールダウンすることもできます。

```
$ kubectl scale deployments/kubernetes-bootcamp --replicas=2
deployment "kubernetes-bootcamp" scaled

$ kubectl get deployments
NAME                  DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
kubernetes-bootcamp   2         2         2            2           8m

$ kubectl get pods -o wide
NAME                                   READY     STATUS        RESTARTS   AGE       IP           NODE
kubernetes-bootcamp-5dbf48f7d4-26vrz   1/1       Terminating   0          7m        172.18.0.7   host01
kubernetes-bootcamp-5dbf48f7d4-bdfgn   1/1       Terminating   0          7m        172.18.0.5   host01
kubernetes-bootcamp-5dbf48f7d4-qvj25   1/1       Running       0          8m        172.18.0.4   host01
kubernetes-bootcamp-5dbf48f7d4-s542c   1/1       Running       0          7m        172.18.0.6   host01
```

# Performing a Rolling Update

### 概要

ユーザはアプリケーションをいつでも使えることを想定しており、開発者は1日に何回も新しいバージョンをデプロイできることを想定しています。Kubernetes はローリングアップデートによりこれを実現します。ローリングアップデートにより、ゼロダウンタイムでデプロイすることができます。Pod インスtんすを段階的にアップデートすることで実現しています。新しい Pod は利用可能リソースのあるノード上にスケジュールされます。

前の章では、アプリケーションをスケールさせました。アップデート時にアプリケーションに影響のないようにする必要があります。デフォルトでは、アップデート中は Pod を最大量使えないようになり、新しい Pod が最大値分作られます。ポッド数もしくは割合で設定することができます。Kubernetes 匂いては、アップデートはバージョニングされ、どの Deployment のアップデートも元に戻すことができます。

Deployment がパブリックに公開されている場合、アップデート中はサービスは利用可能ノードに飲みトラフィックをロードバランスします。

ローリングアップデートは以下のアクションを可能にします。

* アプリケーションの別の環境への変化を促進します
* 前のバーションにロールバックします。
* ゼロタイムで継続的インテグレーションとデリバリを行います。

### チュートリアル

Pod 数 4 の状態から始めます。

```
$ kubectl get deployments
NAME                  DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
kubernetes-bootcamp   4         4         4            0           12s
```

```
$ kubectl get pods
NAME                                   READY     STATUS    RESTARTS   AGE
kubernetes-bootcamp-5c69669756-5rppg   1/1       Running   0          14s
kubernetes-bootcamp-5c69669756-gw2vw   1/1       Running   0          14s
kubernetes-bootcamp-5c69669756-lr4wn   1/1       Running   0          14s
kubernetes-bootcamp-5c69669756-m4l29   1/1       Running   0          14s
```

この段階では、v1 のイメージです。

```
$kubectl describe pods
...
Image:          gcr.io/google-samples/kubernetes-bootcamp:v1
...
```

次のコマンドで v2 のコンテナにアップデートします。

```
$ kubectl set image deployments/kubernetes-bootcamp kubernetes-bootcamp=jocatalin/kubernetes-bootcamp:v2
deployment.apps "kubernetes-bootcamp" image updated
```

v2 の新しいコンテナを作成し、v1 のコンテナを落としていく動作となります。
v2 になっているかどうかはコンテナにアクセスすると分かります。

```
$ export NODE_PORT=$(kubectl get services/kubernetes-bootcamp -o go-template='{{(index .spec.ports 0).nodePort}}')
$ echo NODE_PORT=$NODE_PORT
NODE_PORT=32737

$ curl $(minikube ip):$NODE_PORT
Hello Kubernetes bootcamp! | Running on: kubernetes-bootcamp-7799cbcb86-kqvg8 | v=2
```

もしくは、次のコマンドでも確認できます。

```
$ kubectl rollout status deployments/kubernetes-bootcamp
deployment "kubernetes-bootcamp" successfully rolled out
```

試しに、存在しないイメージバージョンへのアップデートを試みます。

```
$ kubectl set image deployments/kubernetes-bootcamp kubernetes-bootcamp=gcr.io/google-samples/kubernetes-bootcamp:v10
deployment.apps "kubernetes-bootcamp" image updated
```

失敗している様子が分かります。

```
$ kubectl get pods
NAME                                   READY     STATUS         RESTARTS   AGE
kubernetes-bootcamp-5f76cd7b94-4ptzf   0/1       ErrImagePull   0          1m
kubernetes-bootcamp-5f76cd7b94-ccg6t   0/1       ErrImagePull   0          1m
kubernetes-bootcamp-7799cbcb86-4xb7p   1/1       Running        0          9m
kubernetes-bootcamp-7799cbcb86-8frrq   1/1       Running        0          9m
kubernetes-bootcamp-7799cbcb86-kqvg8   1/1       Running        0          9m
```

この場合は、ロールアウトして戻します。

```
$ kubectl rollout undo deployments/kubernetes-bootcamp
deployment.apps "kubernetes-bootcamp"

$ kubectl get pods
NAME                                   READY     STATUS    RESTARTS   AGE
kubernetes-bootcamp-7799cbcb86-4xb7p   1/1       Running   0          10m
kubernetes-bootcamp-7799cbcb86-8frrq   1/1       Running   0          10m
kubernetes-bootcamp-7799cbcb86-kqvg8   1/1       Running   0          10m
kubernetes-bootcamp-7799cbcb86-lchcz   1/1       Running   0          5s
```
