
[Pod Security Policies](https://kubernetes.io/docs/concepts/security/pod-security-policy/)

Kubernetes v1.21 から deprecated となり、v1.25 から削除される。
Pod Security Admission への移行が推奨される。

[PodSecurityPolicy の Admission Controller](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/#podsecuritypolicy) を有効化しておく必要がある。


#### 使い方

PodSecurityPolicy を作成し、サービスアカウントと関連づけて使用するなどの方法がある。
ネームスペース単位で設定することも可能。
Deployment によって作成された Pod はコントローラ側に権限が必要になる点に注意。

ClusterRole のサンプルは以下の通り。PodSecurityPolicy を使用する権限を設定する。ClusterRoleBinding もしくは RoleBinding によりサービスアカウントとバインディングして使用。

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

以下の例のようにネームスペース単位で設定することもできる。```system:serviceaccounts:<namespace>``` の指定。
```yaml
apiVersion: rbac.authorization.k8s.io/v1
# This cluster role binding allows all pods in the "development" namespace to use the baseline PSP.
kind: ClusterRoleBinding
metadata:
  name: psp-baseline-namespaces
roleRef:
  kind: ClusterRole
  name: psp-baseline
  apiGroup: rbac.authorization.k8s.io
subjects:
- kind: Group
  name: system:serviceaccounts:development
  apiGroup: rbac.authorization.k8s.io
- kind: Group
  name: system:serviceaccounts:canary
  apiGroup: rbac.authorization.k8s.io
```


#### サンプル

[Pod Security Standards](https://kubernetes.io/docs/concepts/security/pod-security-standards/) に従ったサンプルは以下の通り。

[privileged](https://raw.githubusercontent.com/kubernetes/website/main/content/en/examples/policy/privileged-psp.yaml)
```yaml
apiVersion: policy/v1beta1
kind: PodSecurityPolicy
metadata:
  name: privileged
  annotations:
    seccomp.security.alpha.kubernetes.io/allowedProfileNames: '*'
spec:
  privileged: true
  allowPrivilegeEscalation: true
  allowedCapabilities:
  - '*'
  volumes:
  - '*'
  hostNetwork: true
  hostPorts:
  - min: 0
    max: 65535
  hostIPC: true
  hostPID: true
  runAsUser:
    rule: 'RunAsAny'
  seLinux:
    rule: 'RunAsAny'
  supplementalGroups:
    rule: 'RunAsAny'
  fsGroup:
    rule: 'RunAsAny'
```

* [baseline](https://raw.githubusercontent.com/kubernetes/website/main/content/en/examples/policy/baseline-psp.yaml)
* [restricted](https://raw.githubusercontent.com/kubernetes/website/main/content/en/examples/policy/restricted-psp.yaml)


## EKS の場合

[ポッドのセキュリティポリシー](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/pod-security-policy.html)

デフォルトのセキュリティポリシー eks.privileged があるが、特に制限は実施されていない。


