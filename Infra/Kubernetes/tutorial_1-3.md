こちらのチュートリアルを行いました。
https://kubernetes.io/docs/tutorials/kubernetes-basics/

その時のメモです。
このページでは、1 〜 3 章分を記載しています。


# Create a Cluster

### 概要

Kubernetes は高可用性のクラスタをコーディネートします。Kubernetes による抽象化により、コンテナ化されたアプリケーションを個々のマシンを意識せずにデプロイできます。Kubernetes はクラスタ内へのアプリケーションコンテナの配置、スケジューリングをより効果的に行います。

Kubernetes クラスタは以下の二つの要素から構成されます。

* Master がクラスタのコーディネートを行います。
* ノードはアプリケーションを実行するワーカーです。

Master はクラスタの管理の責務を持ちます。スケジューリング、アプリケーションを望む状態にするためのメンテナンス、スケーリング、update するときの退避など様々な活動をコーディネートします。

ノードは物理マシンもしくは VM です。ワーカーマシンとして機能します。各ノードでは Kubelet が動いていて、これはノードを管理し、Kubernetes master と通信するためのエージェントです。ノードにはコンテナ操作をするための Docker, rkt などのツールが必要です。

Kubernetes 常にアプリケーションをデプロイした時、master にアプリケーションコンテナを稼働させることを伝える必要があります。Master はクラスタのノード常でコンテナを実行させるようスケジューリングします。ノードは Master と Kubernetes API を介して通信します。エンドユーザもまた、Kubernetes API を介して直接クラスタ操作ができます。

Kubernetes クラスタは物理、仮想どちらのタイプのマシンにもデプロイできます。Kubernetes のデプロイを試す用途では、Minikube を使うことができます。Minikube は軽量な Kubernetes の実装で、あなたのローカルマシン上に VM を作り、1 ノードのみのシンプルなクラスタを構築します。Linux、macOS、Windows で利用可能です　Minikube の CLI は基本的な起動、停止などのオペレーションができるようになっています。このチュートリアルでは、 Minikube がプレインストールされたオンラインのターミナル上で作業します。

### チュートリアル

minikube を起動します。

```
$ minikube version
$ minikube start
```

kubectl で各情報を参照します。

```
$ kubectl version
$ kubectl cluster-info
```

ノードの一覧を確認します。

```
$ kubectl get nodes
```

# Deploy an app

### 概要

Kubernetes コンテナを実行したら、コンテナ化されたアプリケーションをデプロイすることができます。そのために Kuberbetes の Deployment を作成します。Deployment はどのようにアプリケーションインスタンスを作成し、アップデートするかを指示します。Deployment を定義すると、Kuberbetes master はアプリケーションインスタンスをクラスタ内のノード上で稼働するようにスケジュールします。

アプリケーションインスタンスができたら、Kubernetes Deployment コントローラは各インスタンスを継続的に監視します。もし、ノードが停止したり、削除された場合は、Deployment コントローラが別のノードに交換します。この仕組みがセルフヒーリングメカニズムです。

オーケスレーションが存在する前の世界では、インストールスクリプトがアプリの起動に使われていましたが、それらはマシン障害からは回復するのに役立ちませんでした。アプリケーションインスタンスの起動、インスタンスが稼働することの保証を複数ノードに渡って行うこと。従来のアプリケーション管理とは異なるアプローチだと言えます。


Deployment の作成、管理は Kubernetes のコマンドラインインタフェースを通して行うことができます。Kubectl コマンドです。Kuberctl は Kubernetes API を使って、クラスタ操作ができます。この章では、最も一般的な Kubectl コマンドの使用法である、Deployment の作成を行う事でアプリケーションをクラスタ上で実行させます。

Deployment を作成したら、コンテナイメージの指定と何個のレプリカを起動させるかを指定する必要があります。これらの情報は後で更新することもできます。

ここでは、Node.js アプリケーションがパッケージングされた Docker コンテナを使用します。

### チュートリアル

次のコマンドでコンテナを稼働させます。

```
$ kubectl run kubernetes-bootcamp --image=gcr.io/google-samples/kubernetes-bootcamp:v1 --port=8080
```

このコマンドは以下のことを行っている。

* コンテナを稼働させるのに適したノードを探す。
* そのノード上でアプリケーションを実行するようスケジューリングする。
* もし新たなノードが必要となる場合は、インスタンスをリスケジュールしつつ、クラスタを構成する。

次のコマンドで現在稼働中の deployment 一覧を確認できます。

```
$ kubectl get depliyments
NAME                  DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
kubernetes-bootcamp   1         1         1            1           5m
```

pod 間の通信は同一クラスタ内であれば可能、クラスタ外からは接続不能です。
外から接続可能にするためにプロキシを稼働させます。

```
$ kubectl proxy
```

この状態だと、先ほどデプロイしたコンテナに接続可能です。

```
$ curl http://localhost:8001/version
{
  "major": "",
  "minor": "",
  "gitVersion": "v1.9.0",
  "gitCommit": "925c127ec6b946659ad0fd596fa959be43f0cc05",
  "gitTreeState": "clean",
  "buildDate": "2018-01-26T19:04:38Z",
  "goVersion": "go1.9.1",
  "compiler": "gc",
  "platform": "linux/amd64"
}$
```

pod ごとに作られる endpoint にアクセスすることもできます。

```
$ export POD_NAME=$(kubectl get pods -o go-template --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')

$ echo Name of the Pod: $POD_NAME
Name of the Pod: kubernetes-bootcamp-5dbf48f7d4-6vtck

$ curl http://localhost:8001/api/v1/namespaces/default/pods/$POD_NAME/proxy/
Hello Kubernetes bootcamp! | Running on: kubernetes-bootcamp-5dbf48f7d4-6vtck | v=1
```

# Viewing Pods and Nodes

２章で Deployment を作成した時、Kubernetes は Pod を作成しました。Pod とは複数のコンテナをグループ化したもので、Pod 内のコンテナ間で幾つかのリソースを共有しています。

* 共有ストレージ
* ネットワーキング
* 個々のコンテナの起動方法。イメージのバージョンやポート番号。

Pod には複数のコンテナを含めることができます。それぞれ密結合なコンテナを想定しています。例えば、Node.js アプリと data をフィードするコンテナを含めることができます。Pod 内のコンテナは IP アドレス、ポートの空間を共有します。ノード上へは Pod 単位で配置、スケジュールされます。

Pod は Kubernetes プラットフォームにおいて、原始的な李ニットです。Deployment を作成したら、Deployment は コンテナを中に含む Pod を作成します（コンテナを直接作成しているわけではありません）。各 Pod はノードにひも付き、停止、削除されるまでとどまり続けます。もし、ノード障害が発生した場合は、Pod は他の利用可能なノードにスケジュールされます。

Pod は常に Node 上で稼働します。ノードは Kubernetes のワーカーマシンであり、仮想、物理マシンです。各ノードは Master によって管理され、ノードは複数の Pod を持つ可能性があり、Kubernetes Master が自動的に Pod のスケジューリングをハンドリングします

どの Kubernetes ノードも少なくとも以下の条件を満たします。

* Kubelet は Kubernetes Master とノード間の通信に責務を持ちます。Pod とコンテナの稼働に関して管理します。

* コンテナランタイムはレジストリからコンテナイメージを pull したり、コンテナを展開したり、アプリケーションを実行することに責務を道ます。

2 章では、Kubectl コマンドラインインタフェースを使用してきました。3 章では、デプロイされたアプリケーションや環境の情報を取得します。主要なオペレーションは下記のコマンドで対応します。

* **kubectl get** リソース一覧
* **kubectl describe** リソースの詳細情報
* **kubectl logs** Pod 内のコンテナログを表示
* **kubectl exec** Pod 内のコンテナでコマンドを実行

### チュートリアル

pod の稼働状況を確認するには、次のコマンドを実行します。

```
$ kubectl get pods
NAME                                   READY     STATUS    RESTARTS   AGE
kubernetes-bootcamp-5dbf48f7d4-6vgsx   1/1       Running   0          40s
```

Pod の詳細を確認するには、次のコマンドを実行します。IPアドレス、ポートやイベント出力を確認できます。

```
$ kubectl describe pods
```

Pod のログは次のコマンドで確認します。

```
$ export POD_NAME=$(kubectl get pods -o go-template --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')

$ kubectl logs $POD_NAME
```

Pod 内に 1 コンテナのみの場合は、次のコマンドで、コンテナ上でコマンド実行できます。

```
$ kubectl exec $POD_NAME env
```

コンテナ上で bash セッションを始めることもできます。

```
$ kubectl exec -ti $POD_NAME bash
```
