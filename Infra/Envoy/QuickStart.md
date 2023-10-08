
[Quick start](https://www.envoyproxy.io/docs/envoy/latest/start/quick-start/)


## Run Envoy

[Run Envoy](https://www.envoyproxy.io/docs/envoy/latest/start/quick-start/run-envoy)

* `envoy` コマンドを通して操作できる。コンテナイメージを使用する場合は、おそらく Entrypoint に設定されている
* `-c` オプションで設定ファイルを指定。デモ用の設定ファイルは以下の内容。10000 ポートで LISTEN する
```yaml
static_resources:

  listeners:
  - name: listener_0
    address:
      socket_address:
        address: 0.0.0.0
        port_value: 10000
    filter_chains:
    - filters:
      - name: envoy.filters.network.http_connection_manager
        typed_config:
          "@type": type.googleapis.com/envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager
          stat_prefix: ingress_http
          access_log:
          - name: envoy.access_loggers.stdout
            typed_config:
              "@type": type.googleapis.com/envoy.extensions.access_loggers.stream.v3.StdoutAccessLog
          http_filters:
          - name: envoy.filters.http.router
            typed_config:
              "@type": type.googleapis.com/envoy.extensions.filters.http.router.v3.Router
          route_config:
            name: local_route
            virtual_hosts:
            - name: local_service
              domains: ["*"]
              routes:
              - match:
                  prefix: "/"
                route:
                  host_rewrite_literal: www.envoyproxy.io
                  cluster: service_envoyproxy_io

  clusters:
  - name: service_envoyproxy_io
    type: LOGICAL_DNS
    # Comment out the following line to test on v6 networks
    dns_lookup_family: V4_ONLY
    load_assignment:
      cluster_name: service_envoyproxy_io
      endpoints:
      - lb_endpoints:
        - endpoint:
            address:
              socket_address:
                address: www.envoyproxy.io
                port_value: 443
    transport_socket:
      name: envoy.transport_sockets.tls
      typed_config:
        "@type": type.googleapis.com/envoy.extensions.transport_sockets.tls.v3.UpstreamTlsContext
        sni: www.envoyproxy.io
```
* `--config-yaml` オプションにより設定をオーバーライドできる。メイン設定とマージされる動作
```yaml
admin:
  address:
    socket_address:
      address: 127.0.0.1
      port_value: 9902
```
* `--mode validate` オプションにより設定ファイルのバリデーションができる
* `--component-log-level upstream:debug,connection:trace` のようなオプションによりログ出力内容のレベル設定ができる


## Configuration: Static

[Configuration: Static](https://www.envoyproxy.io/docs/envoy/latest/start/quick-start/configuration-static)

* 静的な設定で開始するには static_resources として listeners, clusters の設定が必要
* listeners
  * ここに記載したポート番号で LISTEN する
  * route_config の内容に従って宛先の cluster に送信する
* cluster
  * endpoints により宛先のエンドポイントを指定する


## Envoy admin interface

[Envoy admin interface](https://www.envoyproxy.io/docs/envoy/latest/start/quick-start/admin)

* admin
  * admin サーバを有効化する設定
  * `http://localhost:9901/stats` により統計情報を確認できる
  * `http://localhost:9901/config_dump` により設定ファイルの内容を出力できる


