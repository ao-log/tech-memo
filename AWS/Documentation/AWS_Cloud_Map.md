# Document

[AWS Cloud Map とは?](https://docs.aws.amazon.com/ja_jp/cloud-map/latest/dg/what-is-cloud-map.html)

* 名前空間、サービス、インスタンスの 3 段の構造になっている
* ECS と緊密に統合されており、ECS タスクの起動、停止に応じてインスタンスの作成、削除が行われる


[AWS Cloud Map を使用する](https://docs.aws.amazon.com/ja_jp/cloud-map/latest/dg/using-cloud-map.html)

* 名前空間
  * 検出方法: API コールもしくは DNS
  * パブリック、プライベートの DNS 名前空間を作成した場合、Route 53 のホストゾーンが自動的に作成される
* サービス
  * DNS の場合は、サブドメイン名になる。名前空間が `example.com` でサービス名が `backend` の場合は `backend.example.com`
  * ヘルスチェックの設定
    * Route 53 ヘルスチェック
    * カスタムヘルスチェック
    * 設定しない
* インスタンス
  * DNS の場合は DNS クエリで名前解決できる
  * [DiscoverInstances](https://docs.aws.amazon.com/cloud-map/latest/api/API_DiscoverInstances.html) によりインスタンスの情報を取得できる
  * ヘルスチェックを設定している場合は healthy なインスタンスのみが返却される


[名前空間の作成](https://docs.aws.amazon.com/ja_jp/cloud-map/latest/dg/creating-namespaces.html)

* HTTP 名前空間の場合は `DiscoverInstances` リクエストを使用して検出する。DNS を使用して検出することはできない
* DNS 名前空間の場合は Route 53 ホストゾーンに `service-name.namespace-name` のレコードが作成される

```shell
// HTTP 名前空間の場合
aws servicediscovery create-http-namespace --name name-of-namespace

// DNS 名前空間の場合
aws servicediscovery create-private-dns-namespace --name name-of-namespace --vpc vpc-xxxxxxxxx
```


[名前空間の作成時に指定する値](https://docs.aws.amazon.com/ja_jp/cloud-map/latest/dg/namespaces-values.html)

* Namespace name
  * API コール
  * API コールと VPC の DNS クエリ
  * API コールとパブリック DNS クエリ
* Namespace description
* Instance discovery
  * API コール
  * API コールと VPC の DNS クエリ
  * API コールとパブリック DNS クエリ
* Tags
* VPC


[サービスの作成](https://docs.aws.amazon.com/ja_jp/cloud-map/latest/dg/creating-services.html)

```shell
aws servicediscovery create-service \
    --name service-name \
    --namespace-id  ns-xxxxxxxxxxx \
    --dns-config "NamespaceId=ns-xxxxxxxxxxx,RoutingPolicy=MULTIVALUE,DnsRecords=[{Type=A,TTL=60}]"
```


[サービスの作成時に指定する値](https://docs.aws.amazon.com/ja_jp/cloud-map/latest/dg/services-values.html)

* Service name
  * API コール 
  * API コールと VPC の DNS クエリまたは API コールとパブリック DNS クエリ
* Service description
* Service discovery configuration
  * API および DNS: SRV レコードを作成、もしくは `DiscoverInstances` による検出
  * API のみ: `DiscoverInstances` による検出のみ
* Routing policy (DNS の場合のみ)
  * 加重ルーティング、複数値回答ルーティングなどから選択
  * 副数値回答ルーティングの場合: 最大 8 のインスタンスを返却。
* Record type (DNS の場合のみ)
  * A
  * AAAA
  * CNAME
  * SRV
* TTL (DNS の場合のみ)
* Health check options
  * ヘルスチェックなし
  * Route 53 ヘルスチェック: パブリック DNS 名前空間のみ
  * カスタムヘルスチェック
* Failure threshold (DNS の場合のみ)
* Health check protocol (DNS の場合のみ)
  * HTTP
  * HTTPS
  * TCP
* Health check path (DNS の場合のみ)
  * HTTP, HTTPS の場合のみ
* Tags


[インスタンスの登録](https://docs.aws.amazon.com/ja_jp/cloud-map/latest/dg/registering-instances.html)

```shell
aws servicediscovery register-instance \
    --service-id srv-xxxxxxxxx \
    --instance-id myservice-xx \
    --attributes=AWS_INSTANCE_IPV4=172.2.1.3,AWS_INSTANCE_PORT=808
```


[インスタンスの登録時または更新時に指定する値](https://docs.aws.amazon.com/ja_jp/cloud-map/latest/dg/instances-values.html)

* Instance type
  * IP アドレス
  * EC2 インスタンス
  * 別のリソースを特定するための情報
* Service instance ID
  * サービス内で一意の値にする必要がある
* IPv4 address
* IPv6 address
* Port
* EC2 instance ID
* Custom attributes



# 参考

* Document
  * [AWS Cloud Map とは?](https://docs.aws.amazon.com/ja_jp/cloud-map/latest/dg/what-is-cloud-map.html)


