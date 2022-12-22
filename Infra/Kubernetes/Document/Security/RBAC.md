
[Using RBAC Authorization](https://kubernetes.io/ja/docs/reference/access-authn-authz/rbac/)

Role Based Access Control(RBAC) では、ロールをベースとして、アクセス権限を設定する。


#### RBAC の有効化

RBAC を有効化するには kube-apiserver にて以下のオプションを設定する。
```
kube-apiserver --authorization-mode=Example,RBAC --other-options --more-options
```


#### オブジェクトの種類

* Role: 許可する権限を設定する。スコープはネームスペース内。
* ClusterRole: 許可する権限を設定する。スコープはクラスター全体。そのため、node なども対象となる。```/healthz``` なども設定可能。
* RoleBinding: Role とユーザの紐付け。ClusterRole と紐付けることもできるがスコープはネームスペース内となる。
* ClusterRoleBinding: ClusterRole とユーザの紐付け。

Binding 作成後は roleRef の参照先を変更することはできない。


#### サンプル

Pod の読み取り権限を設定した Role。
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  namespace: default
  name: pod-reader
rules:
- apiGroups: [""] # "" はコアのAPIグループを示します
  resources: ["pods"]
  verbs: ["get", "watch", "list"]
```

Secrets の読み取り権限を設定した ClusterRole。
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  # 「namespace」はClusterRolesがNamespaceに属していないため、省略されています
  name: secret-reader
rules:
- apiGroups: [""]
  resources: ["secrets"]
  verbs: ["get", "watch", "list"]
```

ユーザ ```jane``` に ```pod-reader``` の Role を付与。ネームスペースは ```default```。
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: read-pods
  namespace: default
subjects:
- kind: User
  name: jane
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: Role #Role または ClusterRole
  name: pod-reader
  apiGroup: rbac.authorization.k8s.io
```

グループ ```manager``` の全てのユーザに対し全てのネームスペースの Secrets の読み取り権限を付与。
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: read-secrets-global
subjects:
- kind: Group
  name: manager
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: ClusterRole
  name: secret-reader
  apiGroup: rbac.authorization.k8s.io
```


#### Role の集約

複数の Role を集約することができる。
aggregationRule にて対象となる Role のラベルを指定する。

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: monitoring
aggregationRule:
  clusterRoleSelectors:
  - matchLabels:
      rbac.example.com/aggregate-to-monitoring: "true"
rules: [] # コントロールプレーンは自動的にルールを入力します
```

デフォルト Role の admin, edit, view に権限を追加することも可能。
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: aggregate-cron-tabs-edit
  labels:
    # デフォルトRoleの「admin」と「edit」にこれらの権限を追加する。
    rbac.authorization.k8s.io/aggregate-to-admin: "true"
    rbac.authorization.k8s.io/aggregate-to-edit: "true"
rules:
  ...
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: aggregate-cron-tabs-view
  labels:
    # デフォルトRoleの「view」にこれらの権限を追加します。
    rbac.authorization.k8s.io/aggregate-to-view: "true"
rules:
 ...
```


#### rules[].verbs

[Determine the Request Verb](https://kubernetes.io/docs/reference/access-authn-authz/authorization/#determine-the-request-verb)

* POST: ```create```
* GET, HEAD: ```get``` (for individual resources), ```list``` (for collections, including full object content), ```watch``` (for watching an individual resource or collection of resources)
* PUT: ```update```
* PATCH: ```patch```
* DELETE: ```delete``` (for individual resources), deletecollection (for collections)


#### rules[].apiGroups

一覧の取得方法。
```
kubectl get --raw /apis | jq -r '.groups[].name'
```


#### rules[].resources

一覧の取得方法
```
kubectl get --raw /api/v1 | jq -r '.resources[].name'
```


#### ユーザー向け Role

ユーザー向けの Role が用意されている。

* cluster-admin: スーパーユーザ用途。system:masters に Binding されている。ClusterRoleBinding だとクラスター全体、RoleBinding だとネームスペース内が対象となる。
* admin: ネームスペース内のほとんどの読み取り/書き込みの操作が許可されている。リソースクォータ、ネームスペース自体への書き込みアクセスは許可されていない。
* edit: ネームスペース内のほとんどの読み取り/書き込みの操作が許可されている。Role または RoleBinding の表示、変更は許可されていない。
* view: ネームスペース内のほとんどの読み取りの操作が許可されている。Role または RoleBinding の表示は許可されていない。


#### コアコンポーネントのロール

* system:kube-scheduler
* system:volume-scheduler
* system:kube-controller-manager
* system:node
* system:node-proxier


#### 他のコンポーネントのロール

* system:auth-delegator	
* system:heapster
* system:kube-aggregator
* system:kube-dns
* system:kubelet-api-admin
* system:node-bootstrapper
* system:node-problem-detector
* system:persistent-volume-provisioner


#### 組み込みコントローラーのRole

コントローラごとに対応する Role が存在する。

* system:controller:attachdetach-controller
* system:controller:certificate-controller
* system:controller:clusterrole-aggregation-controller
* system:controller:cronjob-controller
* system:controller:daemon-set-controller
* system:controller:deployment-controller
...




