
[Secret](https://kubernetes.io/ja/docs/concepts/configuration/secret/)

Secret を定義しているマニフェストは base64 エンコードしているだけなので、Git リポジトリにそのまま上げるような運用はできない。

Pod がシークレットを使用する方法は 3 種類

* ボリューム内のファイルとして、Pod の単一または複数のコンテナにマウントする
* コンテナの環境変数として利用する
* Pod を生成するために kubelet がイメージを pull するときに使用する

Secret には以下の種類がある。

* Opaque: arbitrary user-defined data
* kubernetes.io/service-account-token: service account token
* kubernetes.io/dockercfg: serialized ~/.dockercfg file
* kubernetes.io/dockerconfigjson: serialized ~/.docker/config.json file
* kubernetes.io/basic-auth: credentials for basic authentication
* kubernetes.io/ssh-auth: credentials for SSH authentication
* kubernetes.io/tls: data for a TLS client or server
* bootstrap.kubernetes.io/token: bootstrap token data


#### リファレンス

[Secret v1 core](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#secret-v1-core)


#### サンプル

Opaque タイプの場合。

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: mysecret
  namespace: default
type: Opaque
data:
  username: YWRtaW4=
  password: MWYyZDFlMmU2N2Rm
```


#### Pod からの参照

ファイルとして参照する場合。ボリュームとしてマウントし、デコードされた状態でファイルに格納される。

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: mypod
spec:
  containers:
  - name: mypod
    image: redis
    volumeMounts:
    - name: foo
      mountPath: "/etc/foo"
      readOnly: true
  volumes:
  - name: foo
    secret:
      secretName: mysecret
```

環境変数として使用。

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
