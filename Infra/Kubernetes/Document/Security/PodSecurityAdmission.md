
[Pod Security Admission](https://kubernetes.io/docs/concepts/security/pod-security-admission/)

PodSecurityPolicy の後継の位置付け。

参考になった記事
* [PodSecurityPolicyの廃止に備えて、一足先にPodSecurity Admissionを試してみよう!](https://qiita.com/uesyn/items/cf47e12fba5e5c5ea25f)


#### 有効化方法

kube-apiserver で ```--feature-gates="...,PodSecurity=true"``` により有効化する必要がある。
もしくは [Pod Security Admission Webhook](https://git.k8s.io/pod-security-admission/webhook) により対応可能。


#### 使い方

[Pod セキュリティの標準](https://kubernetes.io/ja/docs/concepts/security/pod-security-standards/) で定義された3つのレベル、privileged、baseline、restricted に従って Pod の [Security Context](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/) に制限をかけることができる。

ネームスペースごとに設定することができ、以下のモードを設定できる。

* enforce: ポリシーに違反した場合、Pod の作成が拒否される。
* audit: ポリシー違反時は監査ログに監査アノテーションを追加するトリガーとなるが、許可される。
* warn: ポリシー違反時はユーザーへの警告がトリガーされるが、許可される。

Deployment、Job のようなワークロードオブジェクトに対しては enforce は効力を発揮しない。

[Kubernetes 側での実装箇所](https://github.com/kubernetes/kubernetes/blob/5bd3334ad69a074fafa7f1153e48ae08ca196b07/staging/src/k8s.io/pod-security-admission/admission/admission.go#L338-L379)


#### サンプル

[Enforce Pod Security Standards with Namespace Labels](https://kubernetes.io/docs/tasks/configure-pod-container/enforce-standards-namespace-labels/)
```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: my-baseline-namespace
  labels:
    pod-security.kubernetes.io/enforce: baseline
    pod-security.kubernetes.io/enforce-version: v1.24

    # We are setting these to our _desired_ `enforce` level.
    pod-security.kubernetes.io/audit: restricted
    pod-security.kubernetes.io/audit-version: v1.24
    pod-security.kubernetes.io/warn: restricted
    pod-security.kubernetes.io/warn-version: v1.24
```


## EKS の場合

[Amazon EKS が Kubernetes 1.22 のサポートを開始](https://aws.amazon.com/jp/blogs/news/amazon-eks-now-supports-kubernetes-1-22/)

```
Pod Security Admission
Kubernetes 1.25 では、PodSecurityPolicy が Pod Security Standards (PSS) とPod Security Admission (PSA) に置き換えられます。Kubernetes Pod Security Standards は、Pod の異なる隔離レベルを定義します。これらのスタンダードによって、Pod の動作をどのように制限したいかを明確かつ一貫した方法で定義できます。PSA の取り組みには、PSS で定義された制御を実装するアドミッションコントローラー Webhook プロジェクトが含まれます。AWS は、お客様がこれらの新しいスタンダードのテストを開始できるように、それに応じて Amazon EKS ベストプラクティスガイドを更新しています。
```


[EKS ベストプラクティスガイド - Pod Security Standards (PSS) and Pod Security Admission (PSA)](https://aws.github.io/aws-eks-best-practices/security/docs/pods/#pod-security-standards-pss-and-pod-security-admission-psa)


