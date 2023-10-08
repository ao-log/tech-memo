
## Listeners

[Listeners](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/listeners/listeners)

* TCP, UDP リスナーをサポート
* 各リスナーは [filter_chains](https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/listener/v3/listener.proto#envoy-v3-api-field-config-listener-v3-listener-filter-chains) により構成される


## HTTP

[HTTP connection management](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/http/http_connection_management)

* [HTTP connection manager](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_conn_man/http_conn_man#config-http-conn-man) が組み込まれている
* リトライ
  * 以下設定ファイルの例の場合、ホストの再選択を 3 回試行する
```yaml
              routes:
              - match:
                  prefix: "/"
                route:
                  cluster: cluster_0
                  retry_policy:
                    retry_host_predicate:
                    - name: envoy.retry_host_predicates.previous_hosts
                      typed_config:
                        "@type": type.googleapis.com/envoy.extensions.retry.host.previous_hosts.v3.PreviousHostsPredicate
                    host_selection_retry_max_attempts: 3
```


## upstream clusters

[Cluster manager](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/cluster_manager)


[Service discovery](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/service_discovery)


[DNS Resolution](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/dns_resolution)


[Health checking](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/health_checking)


[Connection pooling](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/connection_pooling)


[Load Balancing](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/load_balancing/load_balancing)


[Outlier detection](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/outlier)

* Passive なヘルスチェック。エンドポイントからのレスポンスを確認し正常性を確認する「外れ値検出」に分類されるもの。例えば 5xx が N 回連続した場合に対象を所定の時間分だけ除去する


[Circuit breaking](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/circuit_breaking)

* 背景としてはカスケード障害のような事象から防御したい。マイクロサービスのあるノードがダウンした時に連鎖的に他のコンポーネントにも影響が波及していくこと
* サービス間におけるコネクション数、リトライ数を設定しておき、それ以上はリクエストを送信せず即座に 503 を返すようにする


## Observability

[Statistics](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/observability/statistics)


## Life of a Request

[High level architecture](https://www.envoyproxy.io/docs/envoy/latest/intro/life_of_a_request#high-level-architecture)







