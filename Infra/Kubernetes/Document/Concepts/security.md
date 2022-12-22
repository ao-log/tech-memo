
## [Security](https://kubernetes.io/docs/concepts/security/)

#### [Overview of Cloud Native Security](https://kubernetes.io/docs/concepts/security/overview/)

4 C と呼ばれる 4 つのレイヤについて考える必要がある。
4 C とは Cloud, Cluster, Containers, Code のこと。
各レイヤのセキュリティモデルは外側のレイヤに基づいている。コードレベルでのみ対処したとしても十分とは言えない。

**Cloud**  
クラウドは Kubernetes クラスターの [trusted computing base](https://en.wikipedia.org/wiki/Trusted_computing_base) である。
もしクラウドレイヤが脆弱であれば、この上に構築されたコンポーネントはセキュアである保証はできない。
各クラウドプロバイダはクラウド内でセキュアにワークロードを実行するためのセキュリティ推奨を作成している。

各クラウドプロバイダはセキュリティの文書を提供している。

**Infrastructure security**  
Kubernetes クラスターのインフラストラクチャをセキュアにするための提案。

* コントロールプレーンへのネットワークアクセス: コントロールプレーンへの全てのアクセスはインターネットに公開せず、クラスター管理に必要な IP アドレスのセットのみ許可するように制限する。
* ノードへのネットワークアクセス: ノードはコントロールの特定ポートもしくは NodePort, LoadBalancer service からのアクセスのみ許可するように制限するべき。可能であれば、パブリックなインターネットには晒さない。
* クラウドプロバイダーの API への Kubernetes アクセス: 各クラウドプロバイダーはコントロールプレーン、ノードに対して異なる権限のセットの付与を必要とする。管理に必要な最小権限を設定するのがベスト。[Kops のドキュメント](https://github.com/kubernetes/kops/blob/master/docs/iam_roles.md#iam-roles) では IAM ポリシー、ロールの情報を提供している。
* etcd へのアクセス: etcd へのアクセスはコントロールプレーンのみに制限されるべき。構成にもよるが、TLS によるアクセスを試みるべき。詳細は [etcd のドキュメント](https://github.com/etcd-io/etcd/tree/main/Documentation) を参照のこと。
* etcd の暗号化: 全てのストレージは暗号化するのが良いプラクティス。etcd はクラスター全体の情報を保持するので、とりわけ暗号化をするべき。

**Cluster**  
Kubernetes のセキュリティはクラスターコンポーネント、アプリケーションの二つの領域がある。

クラスターについては、[Securing a Cluster](https://kubernetes.io/docs/tasks/administer-cluster/securing-a-cluster/) に従うのが良い。

アプリケーションについては、特定のセキュリティの側面にフォーカスする必要がある。
サービス A が重要なサービスとなっており、サービス A がサービス B に依存しているとする。サービス B がリソース枯渇攻撃に対して脆弱だとする。
この場合はサービス B のリソースを制限しないと、サービス A のリスクは高くなる。

以下の表はトピックごとのセキュリティの推奨についてのリンクとなっている。

* [RBAC Authorization (Access to the Kubernetes API)](https://kubernetes.io/docs/reference/access-authn-authz/rbac/)
* [Authentication](https://kubernetes.io/docs/concepts/security/controlling-access/)
* [Application secrets management (and encrypting them in etcd at rest)](https://kubernetes.io/docs/concepts/configuration/secret/https://kubernetes.io/docs/tasks/administer-cluster/encrypt-data/)
* [Ensuring that pods meet defined Pod Security Standards](https://kubernetes.io/docs/concepts/security/pod-security-standards/#policy-instantiation)
* [Quality of Service (and Cluster resource management)](https://kubernetes.io/docs/tasks/configure-pod-container/quality-service-pod/)
* [Network Policies](https://kubernetes.io/docs/concepts/services-networking/network-policies/)
* [TLS for Kubernetes Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/#tls)

**Container**  
コンテナセキュリティはこのガイドのスコープ外である。以下は各トピックの現時点の推奨である。

* Container Vulnerability Scanning and OS Dependency Security: コンテナイメージのビルドのステップで、既知の脆弱性をスキャンすべきである。
* Image Signing and Enforcement: コンテナのコンテンツの信頼性を維持するためにコンテナイメージの署名を行う。
* Disallow privileged users: コンテナを構築するとき、コンテナ実行のための最小権限を持ったユーザーを作成する。
* Use container runtime with stronger isolation: 強固なアイソレーションを提供する [container runtime classes](https://kubernetes.io/docs/concepts/containers/runtime-class/) を選択。

**Code**  
アプリケーションコードは最もコントロールしやすい主要な attack surface の一つ。アプリケーションコードのセキュア化は Kubernetes セキュリティトピックの範囲外。
以下はアプリケーション保護のための推奨事項。

* TLS のアクセスのみに制限: TCP による通信が必要な場合は TLS ハンドシェークを実行すべき。いくつかの例外はあるが、全ての通信を暗号化すべき。さらに一歩進んで、サービス間の通信の暗号化を推奨。mutal TLS authentication もしくは mTLS によって行われ、二つの証明書を保持するサービス間の検証を両サイドにて行う。
* ポート範囲を制限: サービスの提供もしくはメトリクスの収集に必要なポートのみを公開する。
* 3rd パーティの依存コンポーネントのセキュリティ: 3rd パーティのライブラリをスキャンすることは良いプラクティス。各プログラミング言語は自動的にチェックするためのツールを提供している。
* 静的ポート解析: 多くの言語はコードスニペットを解析し潜在的に危険なプラクティスを検出する方法を提供している。可能な限り、自動化されたツールを使用してコードベースをスキャンするべきである。一部のツールは次の URL で提供されている。: https://owasp.org/www-community/Source_Code_Analysis_Tools
* Dynamic probing attacks: いくつかの既知の脆弱性をついて攻撃するツールが提供されている。これには SQL インジェクション、CSRF、XSS が含まれている。有名なツールの一つが OWASP Zed Attack proxy tool である。



#### [Pod Security Standards](https://kubernetes.io/docs/concepts/security/pod-security-standards/)

Pod セキュリティ標準はセキュリティを広くカバーするために 3 つのポリシーを定義している。
これらのポリシーは積み重ねていくものであり許容度の高いものから制限の厳しいものまでさまざまである。
このガイドは各々のポリシーの要求するものを説明する。

* Privileged: 制限されていないポリシーはより高い権限で広範な作業が可能となる。このポリシーは既知の特権の昇格を許可するものである。
* Baseline: 既知の特権昇格を防ぐためのポリシー。デフォルト（最小の指定）の Pod 設定を許容する。
* Restricted: 最小限に制限されたポリシー。Pod を強化するための現在のベストプラクティスに従っている。

**Privileged**  
特権ポリシーは意図的に解放されており、完全に制限されていない。
このポリシーのタイプは、通常特権もしくは信頼されたユーザーによって管理されるシステム、インフラストラクチャレベルのワークロードに適用されることを意図している。

特権ポリシーは制限がないことと定義される。getekeeper のようなデフォルトで許可する仕組みは通常はデフォルトで特権が設定される場合がある。
Pod セキュリティーポリシーのようなデフォルトで拒否するメカニズムの場合は、特権ポリシーは全ての制限を無効化するべきである。

**Baseline**  
ベースラインポリシーは一般的なコンテナワークロードへの適用、既知の特権昇格を防ぐことを意図している。
このポリシーはクリティカルではないアプリケーションの運用者、開発者をターゲットにしている。
以下の項目は無効化するべきである。

Note: In this table, wildcards (*) indicate all elements in a list. For example, spec.containers[*].securityContext refers to the Security Context object for all defined containers. If any of the listed containers fails to meet the requirements, the entire pod will fail validation.

* HostProcess: Windows Pod には Windows ノードへの特権アクセスを可能にする HostPorocess コンテナを実行する機能がある。ベースラインポリシーではホストへの特権アクセスは無効化するべき。HostProcess Pods は Kubernetes v1.22 ではアルファ版の機能である。
* Host Namespaces: ホストのネームスペースの共用は無効化すべき。
  * Restricted Fields
    * spec.hostNetwork
    * spec.hostPID
    * spec.hostIPC
* Privileged Containers: 特権 Pod はほとんどのセキュリティ機構を無にするので、必ず無効化すべき。
  * Restricted Fields
    * spec.containers[*].securityContext.privileged
    * spec.initContainers[*].securityContext.privileged
    * spec.ephemeralContainers[*].securityContext.privileged
* Capabilities: 以下にあげる Capabilities は無効化すべき。
  * Restricted Fields
    * spec.containers[*].securityContext.capabilities.add
    * spec.initContainers[*].securityContext.capabilities.add
    * spec.ephemeralContainers[*].securityContext.capabilities.add
* HostPath Volumes: HostPath ボリュームは禁止すべき。
  * Restricted Fields
    * spec.volumes[*].hostPath
* HostPorts: HostPorts は無効化もしくは最小限に制限すべき。
  * Restricted Fields
    * spec.containers[*].ports[*].hostPort
    * spec.initContainers[*].ports[*].hostPort
    * spec.ephemeralContainers[*].ports[*].hostPort
* AppArmor: サポートされているホストでは、デフォルトの AppArmor プロファイルがデフォルトで適用されている。ベースラインポリシーでは、デフォルトの AppArmor プロファイルの上書き、無効化を避けるべき。もしくは制限を特定の許可されたプロファイルにのみ制限すべき。
  * Restricted Fields
    * metadata.annotations["container.apparmor.security.beta.kubernetes.io/*"]
* SELinux: SELinux の設定、ユーザ、ロールオプションのカスタマイズは禁止すべき。
  * Restricted Fields
    * spec.securityContext.seLinuxOptions.type
    * spec.containers[*].securityContext.seLinuxOptions.type
    * spec.initContainers[*].securityContext.seLinuxOptions.type
    * spec.ephemeralContainers[*].securityContext.seLinuxOptions.type
  * Allowed Values
    * Undefined/""
    * container_t
    * container_init_t
    * container_kvm_t
  * Restricted Fields
    * spec.securityContext.seLinuxOptions.user
    * spec.containers[*].securityContext.seLinuxOptions.user
    * spec.initContainers[*].securityContext.seLinuxOptions.user
    * spec.ephemeralContainers[*].securityContext.seLinuxOptions.user
    * spec.securityContext.seLinuxOptions.role
    * spec.containers[*].securityContext.seLinuxOptions.role
    * spec.initContainers[*].securityContext.seLinuxOptions.role
    * spec.ephemeralContainers[*].securityContext.seLinuxOptions.role
  * Allowed Values
    * Undefined
* /proc マウントタイプ: /proc のマスクは攻撃対象の最小化のため必要。
  * Restricted Fields
    * spec.containers[*].securityContext.procMount
    * spec.initContainers[*].securityContext.procMount
    * spec.ephemeralContainers[*].securityContext.procMount
* Seccomp : seccomp は明示的に Unconfined 設定してはならない。
  * Restricted Fields
    * spec.securityContext.seccompProfile.type
    * spec.containers[*].securityContext.seccompProfile.type
    * spec.initContainers[*].securityContext.seccompProfile.type
    * spec.ephemeralContainers[*].securityContext.seccompProfile.type
* Sysctls: sysctls はセキュリティ機能を無効化したりホスト上の全てのコンテナに影響するため、安全なサブセットとして許可されるものを除いで無効化すべき。コンテナもしくは Pod ないのネームスペースにあり、他の Pod、プロセスから分離されている場合は sysctl は安全だと考えられる。
  * Restricted Fields
    * spec.securityContext.sysctls[*].name
  * Allowed Values
    * Undefined/nil
    * kernel.shm_rmid_forced
    * net.ipv4.ip_local_port_range
    * net.ipv4.ip_unprivileged_port_start
    * net.ipv4.tcp_syncookies
    * net.ipv4.ping_group_range

**restricted**  
制限ポリシーはいくつかの互換性を犠牲にして、Pod を強化するためのベストプラクティスを強制することを意図している。セキュリティがクリティカルとなるアプリケーションの運用者、開発者、信頼性の低いユーザーをターゲットにしている。
以下の項目は強制的に無効化すべきである。

* ベースラインプロファイルの項目全て
* Volume Types: 以下のボリュームタイプのみを許可。
  * Restricted Fields
    * spec.volumes[*]
  * Allowed Values
    * spec.volumes[*].configMap
    * spec.volumes[*].csi
    * spec.volumes[*].downwardAPI
    * spec.volumes[*].emptyDir
    * spec.volumes[*].ephemeral
    * spec.volumes[*].persistentVolumeClaim
    * spec.volumes[*].projected
    * spec.volumes[*].secret
* 特権昇格(v1.8+): set-user-ID、set-group-ID ファイルモードのような特権昇格は許可すべきではない。
  * Restricted Fields
    * spec.containers[*].securityContext.allowPrivilegeEscalation
    * spec.initContainers[*].securityContext.allowPrivilegeEscalation
    * spec.ephemeralContainers[*].securityContext.allowPrivilegeEscalation
  * Allowed Values
    * false
* 非 root での実行: コンテナは非 root ユーザで実行すべきである。
  * Restricted Fields
    * spec.securityContext.runAsNonRoot
    * spec.containers[*].securityContext.runAsNonRoot
    * spec.initContainers[*].securityContext.runAsNonRoot
    * spec.ephemeralContainers[*].securityContext.runAsNonRoot
  * Allowed Values
    * true
* 非 root での実行(v1.23+): runAsUser を 0 に設定するべきではない。
  * Restricted Fields
    * spec.securityContext.runAsUser
    * spec.containers[*].securityContext.runAsUser
    * spec.initContainers[*].securityContext.runAsUser
    * spec.ephemeralContainers[*].securityContext.runAsUser
  * Allowed Values
    * any non-zero value
    * undefined/null
* Seccomp(v 1.19+): seccomp プロファイルは明示的に許可された値の一つに設定すべき。未定義のプロファイル、未設定のプロファイルは禁止される。
  * Restricted Fields
    * spec.securityContext.seccompProfile.type
    * spec.containers[*].securityContext.seccompProfile.type
    * spec.initContainers[*].securityContext.seccompProfile.type
    * spec.ephemeralContainers[*].securityContext.seccompProfile.type
  * Allowed Values
    * RuntimeDefault
    * Localhost
* Capabilities (v1.22+): コンテナは全ての capabilities をドロップするべき。また、NET_BIND_SERVICE のみ許可するべき。
  * Restricted Fields
    * spec.containers[*].securityContext.capabilities.drop
    * spec.initContainers[*].securityContext.capabilities.drop
    * spec.ephemeralContainers[*].securityContext.capabilities.drop
  * Allowed Values
    * Any list of capabilities that includes ALL
  * Restricted Fields
    * spec.containers[*].securityContext.capabilities.add
    * spec.initContainers[*].securityContext.capabilities.add
    * spec.ephemeralContainers[*].securityContext.capabilities.add
  * Allowed Values
    * Undefined/nil
    * NET_BIND_SERVICE


**ポリシーの実装**  
ポリシーの定義とポリシーの実装を切り離すことによって、ポリシーを強制するメカニズムとは独立して、汎用的な理解や複数のクラスターにわたる共通言語とすることができる。
メカニズムが成熟するにつれて、ポリシーごとに以下のように定義される。
個々のポリシーの実施方法はここでは定義されていない。

[Pod Security Admission Controller](https://kubernetes.io/docs/concepts/security/pod-security-admission/)


**FAQ**  

・priviledge と baseline の間にプロファイルがないのは何故？  
三つのプロファイルは安全なものから(restricted)安全でないもの(priviledged) まで直線的に段階が設定されている。
ベースラインを超える特権が必要になる場合はアプリケーションによるため、ニッチなケースに対してプロファイルを提供することはしない。
特権が必要な場合に特権プロファイルを使用するべきという意味ではなく、場合に応じてプロファイルを定義する必要がある。

ただし、将来他のプロファイルが必要になった場合は SIG Auth は再考する。

・セキュリティプロファイルとセキュリティコンテキストの違いは何？  
セキュリティコンテキストは実行時の Pod とコンテナを設定するものである。
セキュリティコンテキストは Pod マニフェスト内の Pod、コンテナ仕様の一部として定義され、コンテナランタイムに渡されるパラメータを示している。

セキュリティプロファイルはコントロールプレーンのメカニズムであり、セキュリティコンテキストやそれ以外の特定の設定を強制するものである。
2021年7月に Pod Security Policies は廃止され、Pod Security Admission Controller が採用されている。

・Windows Pod にはどのプロファイルを適用するべき？  
Kubernetes 内の Winows は Linux ベースのワークロードと比べていくつかの制限や違いがある。
特に Pod の SecurityCOntext フィールドは Windows 環境では効果がない。
よって、現段階においては標準化されたセキュリティポリシーは存在しない。

もし Windows Pod に制限されたプロファイルを適用すると、実行時に Pod に影響が出る場合がある。
制限されたプロファイルは Linux seccomp プロファイルや特権昇格の無効化など Linux 固有の制限を適用する必要がある。
Kubelet やコンテナランタイムはこれらの Linux 固有の値を無視した場合、Windows Pod は制限されたプロファイル内で正常に動作すべきである。
しかし、強制の欠如はベースラインプロファイルと比較して追加の制限がないといえる。

HostProcess Pod を作成するための HostProcess フラグは特権ポリシーに沿ってのみ行われるべきである。
Windows HostProcess Pod の作成はベースライン、制限されたポリシーに置いてはブロックされており、そのためいかなる HostProcess Pod は特権的であると見なされるべきである。

・サンドボックス化された Pod の扱いはどのようにしたらよい？  
現時点においては Pod がサンドボックス化されているかどうかを制御する API 標準はありません。
サンドボックス化された Pod はサンドボックス化されたランタイム(gVisor、Kata Containers)の使用によって特定することは可能だが、サンドボックス化されたランタイムについての標準的な定義は存在しない。

サンドボックス化されたワークロードについての保護ん必要性は、それ以外に対するものとは異なる。
例えば、特権を制限する必要性は小さくなる。ワークロードが元となるカーネルから分離されている場合は。
このことによって強い検眼を必要とするワークロードは依然として隔離された状態となる。

加えて、サンドボックス化されたワークロードの保護はサンドボックス化の方法に強く依存する。
よって、全てのサンドボックス化されたワークロードについては推奨される単一のプロファイルはない。



#### [Pod Security Admission](https://kubernetes.io/docs/concepts/security/pod-security-admission/)

前章の Kubernetes の Pod のセキュリティ標準は Pod に対して異なる分離レベルを定義する。
これらの標準は Pod の動作をどのように制限したいかを明確、かつ一貫した方法で定義できる。

ベータ版の機能として Kubernetes はビルトインの Security admission コントローラーを提供しており、これは  PodSecurityPolicies の後継である。
Pod セキュリティの制限はネームスペースのレベルで適用される。

PodSecurityPolicy の API は deprecated となり Kubernetes 1.25 以降は取り除かれる。

**Built-in Pod Security admission enforcement**  
Kubernetes 1.23 以降、PodSecurity feature gate の機能はベータとなっており、デフォルトで有効化されている。
このページは Kubernetes 1.24 に向けてのドキュメントの一部である。
もし異なるバージョンを使用している場合は、当該リリースのドキュメントを参照のこと。

**Alternative: installing the PodSecurity admission webhook**  
PodSecurity の管理ロジックは admission webhook の有効化によっても利用可能である。
この実装もまたベータとなっている。
ビルトインの PodSecurity admission プラグインを使用できない環境においては、代替としてこの webhook を使用できる。

ビルド前のコンテナイメージ、証明書生成スクリプト、サンプルのマニフェストは次の URL より入手できる。
https://git.k8s.io/pod-security-admission/webhook.

インストール手順。
```
git clone https://github.com/kubernetes/pod-security-admission.git
cd pod-security-admission/webhook
make certs
kubectl apply -k .
```

**Pod Security levels**  
Pod Security admission は Pod セキュリティコンテキストと Pod セキュリティ標準(privileged, baseline, restricted) によって定義された三つのレベルの関連フィールドに要件が設定される。
前章の Pod Security 標準のページを参照すること。

**Pod Security Admission labels for namespaces**  
機能が有効化されるか webhook がインストールされると、ネームスペースに対して admission コントロールモードを定義し、pod のセキュリティ用途で使用することができる。
Kubernetes はラベルのセットを定義している。
ラベルは違反が検出された場合にコントロールプレーンが取るアクションを定義している。

* enforce: ポリシーに違反した場合 pod は拒否される。
* audit: ポリシーに違反した場合、監査アノテーションへの追加がトリガーされる。しかし許可される。
* warn:	ポリシーに違反した場合、ユーザーに警告を発する。しかし許可される。

ネームスペースは任意または全てのモードを設定でき、異なるモードに対して異なるレベルを設定することもできる。
各々のモードでは、ポリシーを決定するための二つのラベルがある。
```yaml
# The per-mode level label indicates which policy level to apply for the mode.
#
# MODE must be one of `enforce`, `audit`, or `warn`.
# LEVEL must be one of `privileged`, `baseline`, or `restricted`.
pod-security.kubernetes.io/<MODE>: <LEVEL>

# Optional: per-mode version label that can be used to pin the policy to the
# version that shipped with a given Kubernetes minor version (for example v1.24).
#
# MODE must be one of `enforce`, `audit`, or `warn`.
# VERSION must be a valid Kubernetes minor version, or `latest`.
pod-security.kubernetes.io/<MODE>-version: <VERSION>
```
mode は enforce, audit, warn のいずれか。LEVEL は privileged, baseline, restricted である。

**Workload resources and Pod templates**  
Pod は Deployment や Job などのオブジェクトによって間接的に作成される。
アークロードオブジェクトは Pod テンプレートを定義し、ワークロードリソースのコントローラーはテンプレートに基づいた Pod を作成する。
違反を早期にキャッチすることを支援するために、audit, warning の両モードはワークロードリソースに適用される。
しかし、enforce モードはワークロードリソースに適用されず、Pod オブジェクトに対してのみ適用される。

**Exemptions**  
Por セキュリティ強制の exemptions を定義することで、特定のネームスペースに関連づけられたポリシーによって禁止されていた Pod の作成を許可することができる。
exemptions は Admission コントローラーによって静的に構成される。

exemptions は明治的に列挙する必要がある。
exemptions を満たした要求は Admission コントローラーによって無視され、enforce, audit, warn の全ての挙動はスキップされる。
exemprions の次元は以下の通りである。

* Usernames: 認証されていないユーザーからのリクエストは無視される。
* RuntimeClassNames: Pod、ワークロードリソースで指定された認証されていないランタイムクラスは無視される。
* Namespaces: 認証されていないネームスペースの Pod、ワークロードリソースは無視される。

ほとんどの Pod はコントローラーによってワークロードリソースへのレスポンスとして作成される。
これが意味することは、exempting は Pod を直接作成する場合にのみ適用される。
コントローラーサービスアカウント(system:serviceaccount:kube-system:replicaset-controller) は除外すべきではない。
そうした場合に、対応するワークロードリソースを区政できる全てのユーザーを暗黙的に除外してしまうため。

以下の Pod のフィールドはポリシーチェックの対象外である。
つまり、Pod の更新要求がこれらのフィールドに対する変更のみであった場合、現在のポリシーレベルに違反している場合においても拒否されることはない。

* seccomp、AppArmor アノテーションを除くメタデータの更新
  * seccomp.security.alpha.kubernetes.io/pod (deprecated)
  * container.seccomp.security.alpha.kubernetes.io/* (deprecated)
  * container.apparmor.security.beta.kubernetes.io/*
* .spec.activeDeadlineSeconds に対する有効な更新
* .spec.tolerations に対する有効な更新



#### [Pod Security Policies](https://kubernetes.io/docs/concepts/security/pod-security-policy/)

PodSecurityPolicy は Kubernetes 1.21 から deprecated になる。1.25 からは取り除かれる。
Pod Security Admission への移行もしくは 3rd パーティの Admission プラグインへの移行を推奨する。

**Pod Security Policy とは**
クラスターレベルのリソース。
各セキュリティ項目について、どのような制限を行うか定義することができる。
Admission コントローラーが必要。

クラスターロールで ```podsecuritypolicies``` リソースに対し ```use``` と設定する。
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: <role name>
rules:
- apiGroups: ['policy']
  resources: ['podsecuritypolicies']
  verbs:     ['use']
  resourceNames:
  - <list of policies to authorize>
```



#### [Security For Windows Nodes](https://kubernetes.io/docs/concepts/security/windows-security/)

Windows OS についてのセキュリティの考慮事項、ベストプラクティスが書かれているページ。

**シークレットの保護**  
Windows では、Secrets からのデータはノードのローカルストレージ上にテキストとして出力される。
以下の点の考慮が必要。
* ACL を用いて Secrets のファイルがある場所をセキュアにする。
* BitLocker によってボリュームレベルの暗号化を行う。

**コンテナユーザ**  
RunAsUsername を使用することによって Windows Pod、コンテナは指定したユーザで実行することができる。
大まかにいうと RunAsUser と同等である。

Windows コンテナは ContainerUser と ContainerAdministrator の二つのデフォルトのユーザアカウントがある。
これらの違いは、[When to use ContainerAdmin and ContainerUser user accounts ](https://docs.microsoft.com/ja-jp/virtualization/windowscontainers/manage-containers/container-security#when-to-use-containeradmin-and-containeruser-user-accounts) のドキュメントに書かれている。

ローカルユーザはコンテナのビルドプロセスの間にコンテナイメージに追加される。
* Nano Server ベースのイメージは ContainerUser がデフォルトとなっている。
* Server Core ベースのイメージは ContainerAdministrator がデフォルトとなっている。
Windows コンテナは Group Managed Service Accounts を使用することによって Active Directory のアイデンティティで実行される。

**Pod レベルのセキュリティ分離**  
Linux に固有なセキュリティコンテキストメカニズム(SELinux, AppArmor, Seccomp, POSIX capabilities) は Windows ノードではサポートされていない。

特権コンテナは Windows ではサポートされていない。
Linux では特権によって実行されるが、HostProcess コンテナが代わりに使用される。



#### [Controlling Access to the Kubernetes API](https://kubernetes.io/docs/concepts/security/controlling-access/)

Kubernetes API には kubectl やクライアントライブラリ、あるいは REST リクエストを用いてアクセスする。
人間のユーザーと Kubernetes サービスアカウントの両方を認証可能。
リクエストが API に到達すると以下のような流れで処理される。

クライアント → Authentition(認証) → Autoriation(認可) → Admission Control

**トランスポート層のセキュリティ**  
API は TLS で保護された 443 番ポートで提供される。

API サーバは証明書を提示する。
この証明書はプライベート CA を使用して署名することも、一般に認知されている CA と連携した公開鍵基盤に基づき署名することも可能。
プライベート CA を用いている場合は、接続を信頼し傍受されていないと確信できるようにクライアントの ~/.kube/config に設定された CA 証明書のコピーが必要。

**認証**  
TLS が確立されると HTTP リクエストは認証のステップに移行する。
クラスター作成スクリプト、クラスター管理者は 1 つまたは複数の Authenticator モジュールを実行できるように API サーバを設定する。
Authenticator については [認証](https://kubernetes.io/ja/docs/reference/access-authn-authz/authentication/) に詳しく記載されている。

認証ステップの入力は HTTP リクエスト全体。しかし、ヘッダとクライアント証明書の両方またはどちらかを検証する。

認証モジュールには、クライアント証明書、パスワード、プレーントークン、ブートストラップトークン、JSON Web Tokens(サービスアカウントで使用)などがある。

複数の認証モジュールを指定することもできる。
その場合は、1 つが成功するまで順番に試行する。

認証できない場合は、HTTP ステータスコード 401 で拒否される。
拒否されなかった場合、ユーザーは特定の username として認証され、そのユーザ名は複数のステップでの判断に使用できる。
いくつかの Authenticator はユーザーのグループメンバシップを提供するが、提供しないものもある。

Kubernetes はアクセスコントロールの決定やリクエストログにユーザー名を使用する。
しかし User オブジェクトを持たず、ユーザー名やその他のユーザーに関する情報を API 内に保存しない。

**認可**  
リクエストが認証された後は、認可される必要がある。
リクエストにはリクエストを行ったユーザー名、アクション、アクションによって影響を受けるオブジェクトを含める必要がある。
既存のポリシーがユーザーが要求したアクションを完了するための権限を持っていると宣言している場合、リクエストは承認される。

例えば、Bob が以下のようなポリシーを持っている場合、彼はネームスペース projectCaribou 内の Pod のみを読み取ることができる。

```json
{
    "apiVersion": "abac.authorization.kubernetes.io/v1beta1",
    "kind": "Policy",
    "spec": {
        "user": "bob",
        "namespace": "projectCaribou",
        "resource": "pods",
        "readonly": true
    }
}
```

Bob が以下のようなリクエストをいた場合、このリクエストは許可される。

```json
{
  "apiVersion": "authorization.k8s.io/v1beta1",
  "kind": "SubjectAccessReview",
  "spec": {
    "resourceAttributes": {
      "namespace": "projectCaribou",
      "verb": "get",
      "group": "unicorn.example.org",
      "resource": "pods"
    }
  }
}
```

Bob がネームスペース projectCaribou のオブジェクトに書き込み(create or update)のリクエストを行った場合は拒否される。
別のネームスペースに対し読み取り(get)を行った場合も拒否される。

Kubernetes の認可では、組織全体またはクラウドプロバイダー全体の既存のアクセスコントロールシステムと対話するために、共通の REST 属性を使用する必要がある。
これらのコントロールシステムは Kubernetes API 以外の API とやり取りする可能性があるため、REST 形式を使用することが重要である。

Kubernetes は ABAC モード、RBAK モード、Webhook モードなど、複数の認可モジュールをサポートしている。
管理者はクラスターを作成する際に API サーバーで使用する認可モジュールを設定する。
複数の認可モジュールが設定されている場合、Kubernetes は各モジュールをチェックし、いずれかのモジュールがリクエストを許可した場合、リクエストを続行することができる。
全てのモジュールがリクエストを拒否した場合はリクエストは拒否される(HTTP ステータスコード 403)。

認可の詳細は [Authorization](https://kubernetes.io/docs/reference/access-authn-authz/authorization/) に詳しく記載されている。

**アドミッションコントロール**  
アドミッションコントロールモジュールは、リクエストを変更したり拒否したりすることができるソフトウェアモジュールである。
認可モジュールが利用できる属性に加えて、アドミッションコントロールモジュールは、作成または修正されるオブジェクトのコンテンツにアクセスすることができる。

アドミッションコントローラーは、オブジェクトの作成、変更、削除、または接続(プロキシ)を行うリクエストに対して動作する。
アドミッションコントローラーは、単にオブジェクトを読み取るだけのリクエストには動作しない。
 複数のアドミッションコントローラーが設定されている場合は、順番に呼び出される。

いずれかのアドミッションコントローラーモジュールが拒否した場合、リクエストは即座に拒否される。
オブジェクトを拒否するだけでなく、アドミッションコントローラーは、フィールドに複雑なデフォルトを設定することもできる。
利用可能なアドミッションコントロールモジュールは、アドミッションコントローラーに記載されている。

リクエストがすべてのアドミッションコントローラーを通過すると、対応する API オブジェクトの検証ルーチンを使って検証され、オブジェクトストアに書き込まれる。

**監査**  
Kubernete sの監査は、クラスター内の一連のアクションを文書化した、セキュリティに関連する時系列の記録を提供する。
クラスターは、ユーザー、Kubernetes API を使用するアプリケーション、およびコントロールプレーン自身によって生成されるアクティビティを監査する。

詳しくは [監査](https://kubernetes.io/ja/docs/tasks/debug-application-cluster/audit/) に記載されている。



#### [Role Based Access Control Good Practices](https://kubernetes.io/docs/concepts/security/rbac-good-practices/)

[Using RBAC Authorization](https://kubernetes.io/docs/reference/access-authn-authz/rbac/#restrictions-on-role-creation-or-update) と併せて読むと良い。

**一般的な Goog Practice**  

* ネームスペース単位で割り当て。ClusterRoleBindings よりは RoleBindings。
* 最小権限。
* cluster-admin は極力使用しない。
* system:mastersグループにユーザーを追加しない。このグループのメンバーであるユーザーは、すべての RBAC 権限チェックをバイパスし、常に無制限のスーパーユーザーアクセス権を持つ。

**特権 token の配布を最小化する**  

* 信頼できない Pod は NodeAffinity、または PodAntiAffinity を使用して、重要な Pod とは同一ノードに乗せないようにする。

**Hardening**  

* system: アカウントには変更を加えるべきではない。

**KubernetesRBAC-特権昇格のリスク**  

以下の点に注意を払うべきである。

* Secrets の一覧表示
* Pod Security Admission によるベースライン、制限付きの適用












