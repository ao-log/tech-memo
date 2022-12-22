
## [Policies](https://kubernetes.io/docs/concepts/policy/)

#### [Limit Ranges](https://kubernetes.io/docs/concepts/policy/limit-range/)

デフォルトでは、コンテナのリソースは制限されていない。
リソースクォータを使用することによって、クラスター管理者はリソース消費と作成をネームスペースに基づいて制限することができる。
ネームスペース内では、Pod もしくはコンテナはネームスペースのリソースクォータによって定義された CPU、メモリの範囲内でできるだけ多くを消費することができる。
そのため一つの Pod もしくはコンテナが利用可能なリソースを占有できる恐れがある。
LimitRange はネームスペースないのリソースの割り当てを制限するポリシーである。

LimitRange を使用することで、以下のような制限を課すことができる。
* ネームスペース内の Pod、コンテナごとの最小、最大のコンピュートリソースの使用量
* ネームスペース内の PersistentVolumeClaim ごとのストレージの最小、最大の使用量
* ネームスペース内のリソースについての request と limit の比率
* ネームスペース内のデフォルトの request/limit とコンテナ実行時の強制的な適用

**Limit Range の有効化**  
LimitRange サポートは Kubernetes 1.10 以降にデフォルトで有効化された。
LimitRange が強制されるのは LimitRange オブジェクトがネームスペース内に存在する場合である。
 LimitRange オブジェクトの名前は有効な DNS のサブドメインでなければならない。

**Limit Range の概要**
* 管理者はネームスペース内に一つの LimitRange を作成する。
* ユーザーは Pod, コンテナ、PersistentVolumeClaims のようなリソースをネームスペース内に作成する。
* LimitRanger admission コントローラーはデフォルトとリミットを全ての Pod とコンテナに強制し、LimitRange で設定した最小、最大、比率を超えないことを保証する。
* リソースを作成もしくは更新する際、LimitRange の制約に違反すると、API リクエストは HTTP ステータスコード 403 で失敗し、破られた制約について説明するメッセージが表示される。
* もし LimitRange がネームスペース内に適用されると、ユーザーは reqest もしくは limit をそれらの値に指定しなければならない。そうしなければ Pod の作成は拒否され、LimitRange の違反が発生する場合がある。LimitRange の違反は Pod の Admission ステージでのみ発生し、稼働中の Pod では発生しない。

以下は LimitRange を使用することによって作成できるポリシーの例である。

* 8 GiB の RAM、16 core の CPU を持つ 2 ノードのクラスターにおいて、ネームスペース内の Pod について CPU は 100m の request と最大 500m の limit、メモリを 200Mi の request と最大 600Mi に制限する。
* spec 内で cpu と memory の request を定義していないコンテナについて、CPU のデフォルトの limit と request を 150m にし、メモリのデフォルト request を 300 Mi にする。

ネームスペースのトータルの limits が Pod/コンテナの Limit の合計よりも小さい場合は、リソースについて競合が発生する可能性がある。
その場合、コンテナ、Pod は作成されない。

LimitRange に対する競合や変更は、既に作成済みのリソースについては影響しない。



#### [Resource Quotas](https://kubernetes.io/docs/concepts/policy/resource-quotas/)

ResourceQuota オブジェクトによってリソースクォータをネームスペースごとに設定できる。
クォータ制約に違反している場合は 403 FORBIDDEN が返される。
cpu, memory のリソースクォータが設定されている場合は、requests, limits の指定が必要。デフォルト値の設定のためには LimitRanger の Admission コントローラを使用できる。
以下のようなポリシーを設定できる。

* 32 GiB RAM、16 コアのキャパシティーを持つクラスターで、A チームに 20 GiB、10 コアを割り当て、B チームに10 GiB、4 コアを割り当て、将来の割り当てのために 2 GiB、2 コアを予約しておく。
* "testing"という名前空間に対して1コア、1 GiB RAMの使用制限をかけ、"production"という名前空間には制限をかけない。

クラスターの総キャパシティが名前空間のクォータの合計よりも少ない場合は、競合が発生する場合があり、その場合は先着順となる。

**リソースクォータの有効化**  
kube-apiserver にて --enable-admission-plugins= の値に ResourceQuota が含まれている場合に有効になる。

**リソースクォータの計算**  
名前空間におけるコンピュートリソースの合計を設定できる。

* limits.cpu
* limits.memory
* requests.cpu
* requests.memory
* hugepages-<size>
* cpu: requests.cpuと同じ
* memory: requests.memoryと同じ

**拡張リソースのためのリソースクォータ**  
拡張リソースに対しては requests. のみ設定できる。

リソース名が nvidia.com/gpu で GPU の上限を 4 とするとき以下のように設定する。
```
requests.nvidia.com/gpu: 4
```

**ストレージのリソースクォータ**  

* requests.storage: PVC に対するクォータ
* persistentvolumeclaims: PVC 数
* <storage-class-name>.storageclass.storage.k8s.io/requests.storage: <storage-class-name> に関連する PVC に対するクォータ
* <storage-class-name>.storageclass.storage.k8s.io/persistentvolumeclaims: <storage-class-name> に関連する PVC 数

Kubernetes v 1.8 よりローカルのエフェメラルストレージに対するリソースクォータのサポートが α 版の機能として提供された。

* requests.ephemeral-storage:
* limits.ephemeral-storage:
* ephemeral-storage: requests.ephemeral-storage と同じ

**オブジェクト数に対するクォータ**  

設定可能なリソースの例は以下の通り。

* count/persistentvolumeclaims
* count/services
* count/secrets
* count/configmaps
* count/replicationcontrollers
* count/deployments.apps
* count/replicasets.apps
* count/statefulsets.apps
* count/jobs.batch
* count/cronjobs.batch

オブジェクトの作りすぎの防止に有効。
Secret が大量にある場合はサーバ、コントローラの起動を妨げることになる。また、不適切な CronJob からの大量のジョブ作成を防止できる。Pod については IP アドレスの枯渇を防止できる。

**クォータのスコープ**  
クォータは、列挙された scope の共通部分と一致する場合にのみリソースの使用量を計測する。

* Terminating: .spec.activeDeadlineSeconds >= 0 である Pod に一致。
* NotTerminating: .spec.activeDeadlineSeconds が nil である Pod に一致。
* BestEffort: ベストエフォート型のサービス品質の Pod に一致。
* NotBestEffort: ベストエフォート型のサービス品質でない Pod に一致。
* PriorityClass: 指定された優先度クラスと関連付いている Pod に一致。

BestEffort スコープはリソースクォータを次のリソースに対するトラッキングのみに制限する。

* pods

Terminating、NotTerminating、NotBestEffort、PriorityClass スコープは、リソースクォータを次のリソースに対するトラッキングのみに制限する。

* pods
* cpu
* memory
* requests.cpu
* requests.memory
* limits.cpu
* limits.memory

同じクォータで Terminating と NotTerminating の両方のスコープを指定することはできない。
BestEffort と NotBestEffort についても両方のスコープを指定することはできない。

scopeSelectorはoperator フィールドにおいて下記の値をサポート。

* In
* NotIn
* Exists
* DoesNotExist

scopeSelectorの定義において scopeName に下記のいずれかの値を使用する場合、operator に Exists を指定する。

* Terminating
* NotTerminating
* BestEffort
* NotBestEffort

operator が In または NotIn の場合、values フィールドには少なくとも 1 つの値が必要。

```yaml
  scopeSelector:
    matchExpressions:
      - scopeName: PriorityClass
        operator: In
        values:
          - middle
```

operator が Exists または DoesNotExist の場合、values フィールドを指定しないこと。

**PriorityClass毎のリソースクォータ**  
Pod の優先度に基づいて Pod のシステムリソースの消費をコントロールできる。
scopeSelector に一致する Pod が集計対象となる。

Priority class によってクォータがスコープされる場合は、クォータは以下リソースのみ制限できる。

* pods
* cpu
* memory
* ephemeral-storage
* limits.cpu
* limits.memory
* limits.ephemeral-storage
* requests.cpu
* requests.memory
* requests.ephemeral-storage

サンプル。
* リソースクォータが 3 つ。それぞれ、PriorityClass が high, medium, low。
* 優先度が高いほど、ハードリミットが高くなっている。

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
- apiVersion: v1
  kind: ResourceQuota
  metadata:
    name: pods-medium
  spec:
    hard:
      cpu: "10"
      memory: 20Gi
      pods: "10"
    scopeSelector:
      matchExpressions:
      - operator : In
        scopeName: PriorityClass
        values: ["medium"]
- apiVersion: v1
  kind: ResourceQuota
  metadata:
    name: pods-low
  spec:
    hard:
      cpu: "5"
      memory: 10Gi
      pods: "10"
    scopeSelector:
      matchExpressions:
      - operator : In
        scopeName: PriorityClass
        values: ["low"]
```

次のコマンドでクォータの状況を確認できる。
```
kubectl describe quota
```

Pod においては PriorityClassName によって指定する。
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

**requests vs limits**  
requests.cpu、requests.memory については各コンテナはリソースに対する明示的な要求を行う。
limits.cpu、limits.memory については各コンテナはリソースに対する明示的な制限を行う。

**クォータの確認と設定**  

リソース量のリミットのサンプル。
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

オブジェクト数のリミットのサンプル。
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

count/<resource>.<group> というシンタックスにより、名前空間に依存した全ての主要なリソースに対するクォータをサポートしている。
```
$ kubectl create quota test \
    --hard=count/deployments.apps=2,count/replicasets.apps=4,count/pods=3,count/secrets=4 \
    --namespace=myspace

$ kubectl describe quota --namespace=myspace
Name:                         test
Namespace:                    myspace
Resource                      Used  Hard
--------                      ----  ----
count/deployments.apps        1     2
count/pods                    2     3
count/replicasets.apps        1     4
count/secrets                 1     4
```

**クォータとクラスター容量**  
ResourceQuota はクラスター容量に依存せず、絶対値として表される。

以下のようなポリシーが必要となる場合がある。
* 複数チーム間でクラスターリソースの総量を分けあう。
* 各テナントが必要な時にリソース使用量を増やせるようにするが、偶発的なリソースの枯渇を防ぐために上限を設定する。
* 1つの名前空間に対してリソース消費の需要を検出し、ノードを追加し、クォータを増加させる。

上記のようなポリシーは、クォータ使用量の監視、他のシグナルに従ってクォータ値を調整するコントローラを記述することによって、ResourceQuota をビルディングブロックのように使用して実装できる。

ResourceQuota はクラスターリソースを分割できるが、ノードに対しては何の制限も行わない。

**デフォルトで優先度クラスの消費を制限する**  

特定の PriorityClass を特定のネームスペースでのみ使用できるように設定できる。
kube-apiserver の --admission-control-config-file で以下のファイルにパスを通す必要がある。
```yaml
apiVersion: apiserver.config.k8s.io/v1
kind: AdmissionConfiguration
plugins:
- name: "ResourceQuota"
  configuration:
    apiVersion: apiserver.config.k8s.io/v1
    kind: ResourceQuotaConfiguration
    limitedResources:
    - resource: pods
      matchScopes:
      - scopeName: PriorityClass
        operator: In
        values: ["cluster-services"]
```

ResourceQuota は以下のように設定し、kube-system ネームスペースに適用する。
```yaml
apiVersion: v1
kind: ResourceQuota
metadata:
  name: pods-cluster-services
spec:
  scopeSelector:
    matchExpressions:
      - operator : In
        scopeName: PriorityClass
        values: ["cluster-services"]
```

この場合は Pod 作成は以下の場合に許可される。

* priorityClassName が指定されていない場合
* priorityClassName が指定されており、cluster-services 以外が指定されている場合
* priorityClassName が cluster-services に設定されており、kube-system ネームスペースに作成される場合。かつリソースクォータのチェックに通過した場合

priorityClassName を cluster-services にした場合、kube-system 以外の名前空間での Pod 作成は拒否される。



#### [Process ID Limits And Reservations](https://kubernetes.io/docs/concepts/policy/pid-limiting/)

Kubernetes は Pod が使用できる PID 数の制限を書けることができる。
OS とデーモン用に割り当て可能な PID 数を予約することもできる。

PID はノードにおいて基本的なリソースである。
他のリソースの制限なしに、タスク制限にヒットすることは簡単で、このことはホストマシンを不安定にする要因となる。

クラスタ管理者は Pod が PID を浪費しホスト上で稼働するデーモンを妨げないように保証する必要がある。
加えて、ノード上の他のワークロードに影響を与えないように制限をすることも重要である。

ある種の OS は PID 数の制限は 32768 個である。
/proc/sys/kernel/pid_max によって増やすことができる。

PID が使用可能な PID の制限は kubelet によって設定できる。
例えば、ホスト OS が 262144 個の PID を最大値に設定しており、250 個の Pod が稼働すると想定される場合は、Pod の予算を 1000 PID に制限することでノードに渡って使用可能な PID を使い切ることを防ぐことができる。
管理者がオーバーコミットをしたい場合は、追加のリスクを負うことになる。
いずれにせよ、単一の Pod がマシンをダウンさせることはできない。
この種のリミットは単純な fork 爆弾をがクラスター全体に影響を与えることを防ぐことができる。

Pod ごとの PID 制限は Pod を他の Pod から保護することを可能とするが、ノード全体に影響を与えないことを保証するものではない。
Pod ごとの制限は PID の枯渇を防ぐものではない。

ノードのオーバーヘッド用に PID を予約することもできる。
これは CPU、メモリを OS のために予約するのと似ている。

PID 制限はコンピュートリソースの requests, limits の重要な兄弟である。
しかし、指定方法は異なっており、Pod のリソース制限とするのではなく、kubelet の制限として行う。

**Node Pid limits**  
Kubernetes はシステム用に PID 数を予約することができる。
予約を構成するために、kubelet のコマンドラインオプション ```--system-reserved``` と ```--kube-reserved``` にパラメータ pid=<number> を設定する。
指定した値は、システム用と Kubernetes のシステムデーモン用に予約される PID 数を宣言する。

**Pod Pid limits**  
Kubernetes は Pod 用にプロセス数の制限を行うことができる。
この制限はノードレベルで指定する。
制限を構成するため、kubelet の --pod-max-pids オプションに設定するか、PodPidLimit を kubelet の構成ファイルにて設定する。

**PID based eviction**  
Pod が不正な量のリソースを消費した際に終了するように kubelet を構成することができる。
この昨日は eviction と呼ばれる。
リソース使用量が超えた時にさまざまな eviction のシグナルを構成することができる。
pid.available eviction シグナルは PID 数の閾値を構成する。
ソフト、ハードの eviction ポリシーを設定できる。
しかし、ハードの eviction ポリシーを使用した場合においても PID 数が急速に成長した場合は、ノードは PID リミットに到達し不安定な状態になりうる。
eviction シグナルの値は定期的に計算され、制限を強制しない。

PID 制限 - Pod ごと、ノードごとにハードリミットを制限できる。
制限にと王達した場合、ワークロードは新しい PID を取得しゆおとした場合に失敗する。
これは Pod のリスケジュールを導くことになるかもしれない。
これはワークロードに依存し、liveless, rediness probe の設定状況やどのように失敗しるかにもよる。
しかし、リミットが正確に設定されている場合、一つの Pod が暴走しても、他の Pod のワークロードとシステムプロセスが PID を枯渇に陥らないことを保証できる。



#### [Node Resource Managers](https://kubernetes.io/docs/concepts/policy/node-resource-managers/)

レイテンシが重要で高スループットのワークロードをサポートするため、Kubernetes は Resource Managers のスイートを提供している。
このマネージャは特定の要求を行う Pod についてリソースの割り当てを最適化することを狙いとしている。

メインのマネージャであるトポロジーマネージャは Kubelet のコンポーネントとなっており、ポリシーを通してリソースマネジメントを行う。

ここのマネージャの構成は以下のドキュメントに記載されている。

* [CPU Manager Policies](https://kubernetes.io/docs/tasks/administer-cluster/cpu-management-policies/)
* [Device Manager](https://kubernetes.io/docs/concepts/extend-kubernetes/compute-storage-net/device-plugins/#device-plugin-integration-with-the-topology-manager)
* [Memory Manager Policies](https://kubernetes.io/docs/tasks/administer-cluster/memory-manager/)


