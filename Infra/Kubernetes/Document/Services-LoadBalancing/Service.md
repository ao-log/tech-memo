
[Service](https://kubernetes.io/ja/docs/concepts/services-networking/service/)

* ```.spec.selector``` で指定したラベルを持つ Pod へトラフィックを転送。
* ClusterIP が割り当てられる。クラスター内からのみ名前解決可能。
* デフォルトは TCP。
* 名前解決
  * 同一ネームスペース内はサービス名で名前解決できる。
  * <サービス名>.<ネームスペース名> で名前解決できる。


#### サンプル

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  selector:
    app: MyApp
  ports:
  - protocol: TCP
    port: 80
    targetPort: 9376
```

セレクターなしで作成することも可能。その場合は、Endpoints オブジェクトによってターゲットの設定を行う。
```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  ports:
  - protocol: TCP
    port: 80
    targetPort: 9376
```

```yaml
apiVersion: v1
kind: Endpoints
metadata:
  name: my-service
subsets:
  - addresses:
      - ip: 192.0.2.42
    ports:
      - port: 9376
```


#### A レコードを返す対応になっていない経緯

* DNS の実装がレコードの TTL をうまく扱わず、期限が切れた後も名前解決の結果をキャッシュするという長い歴史がある。
* いくつかのアプリケーションでは DNS ルックアップを一度だけ行い、その結果を無期限にキャッシュする。
* アプリケーションとライブラリーが適切なDNS名の再解決を行ったとしても、DNS レコード上の0もしくは低い値の TTL が DNS に負荷をかけることがあり、管理が難しい。


#### Headless Service

* .spec.clusterIP を None にすることで Headless Service となる。
* ターゲットの IP アドレスが直接返却される。そのため DNS キャッシュに注意。


#### サービスの種類

* ClusterIP: クラスター内部からのみ疎通性がある。
* NodePort: 各ノードの IP アドレスにて静的なポートを公開。ClusterIP も作られる。
  * 30000-32767 の範囲。
  * ```.spec.externalTrafficPolicy``` が Local の場合はノード到達後に別ノードに転送しない。デフォルトの Cluster の場合は他ノードにある Pod もターゲットとなる。
* LoadBalancer: クラウドプロバイダーのロードバランサを作成。NodePort, ClusterIP も作られる。
* ExternalName: 外部のドメインに対する CNAME を作成。


#### NodePort

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  type: NodePort
  selector:
    app: MyApp
  ports:
      # デフォルトでは利便性のため、 `targetPort` は `port` と同じ値にセットされます。
    - port: 80
      targetPort: 80
      # 省略可能なフィールド
      # デフォルトでは利便性のため、Kubernetesコントロールプレーンはある範囲から1つポートを割り当てます(デフォルト値の範囲:30000-32767)
      nodePort: 30007
```


#### LoadBalancer

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  selector:
    app: MyApp
  ports:
  - protocol: TCP
    port: 80
    targetPort: 9376
  clusterIP: 10.0.171.239
  type: LoadBalancer
status:
  loadBalancer:
    ingress:
    - ip: 192.0.2.127
```

annotations でクラウドプロバイダー固有の属性を設定する。


#### External Name

CNAME 的な役割。

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-service
  namespace: prod
spec:
  type: ExternalName
  externalName: my.database.example.com
```


