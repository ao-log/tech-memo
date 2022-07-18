
#### Kubernetes API へのアクセスコントロール

[Kubernetes APIへのアクセスコントロール](https://kubernetes.io/ja/docs/concepts/security/controlling-access/)

認証、認可、Admission Control の順に処理される。

* 認証
  * ヘッダとクライアント証明書の片方もしくは両方を調べる。
  * 認証モジュールは クライアント証明書、パスワード、プレーントークン、ブートストラップトークン、JSON Web Tokens など。
  * 認証モジュールのうち、いずれか一つ成功すればよい。
  * 認証失敗時は HTTP ステータスコード 401 で拒否される。
* 認可
  * リクエストにはユーザー名、リアクション、アクションによって影響を受けるオブジェクトの情報を含める必要がある。
  * アクセスが許可されていない場合は、リクエストは HTTP ステータスコード 403 で拒否される。
  * ABACモード、RBACモード、Webhookモードなど、複数の認可モジュールがある。
* Admission Control
  * いずれか一つの Admission Controller が拒否すると、リクエストは失敗する。
  * どのような Admission Controller があるかは [What does each admission controller do?](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/#what-does-each-admission-controller-do) を参照。LimitRanger、ServiceAccount、ResourceQuota、PodSecurityPolicy などがある。


#### 認証

[認証](https://kubernetes.io/ja/docs/reference/access-authn-authz/authentication/)

Kubernetes クラスターには 2 種類のユーザがある。

* 通常のユーザー
* ServiceAccount

Kubernetes は、クライアント証明書、Bearer トークン、認証プロキシー、HTTP Basic 認証を使い、認証プラグインを通して API リクエストを認証する。

**サービスアカウントトークン**  

* サービスアカウントは ServiceAccountAdmission Controller を介してクラスター内の Pod に関連付けられる。
* Bearer トークンは Pod にマウントされる。
* Pod では以下のように ```serviceAccountName``` にて指定する。

```yaml
apiVersion: apps/v1 # このapiVersionは、Kubernetes1.9時点で適切です
kind: Deployment
metadata:
  name: nginx-deployment
  namespace: default
spec:
  replicas: 3
  template:
    metadata:
    # ...
    spec:
      serviceAccountName: bob-the-bot
      containers:
      - name: nginx
        image: nginx:1.14.2
```


#### 認可

[Authorization Overview](https://kubernetes.io/docs/reference/access-authn-authz/authorization/)


#### Admission Control

[Using Admission Controllers](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/)


