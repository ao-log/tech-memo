
https://www.envoyproxy.io/

* 特徴
  * L7 で稼働するプロキシ。モダンなサービス志向アーキテクチャ向けの通信バス。ネットワークはアプリケーションに対して透過的であるべきという考え方のもと作られている
  * メモリ使用量が少ないコンテナ。あらゆるアプリケーション言語、フレームワークと共に稼働させることができる
  * HTTP/2, GRPC のサポート
  * 高度な機能を持ったロードバランシング。リトライ、サーキットブレーカー、レート制限、リクエストのコピーなど
  * 構成を動的に管理することが可能。構成管理用の API が提供されている
  * L7 の詳細な可観測性、分散トレーシングのサポート、DB に対する可観測性

ドキュメントにも特徴が列挙されている。

https://www.envoyproxy.io/docs/envoy/latest/intro/what_is_envoy


[用語](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/intro/terminology)

* Downstream: Envoy に接続し、リクエストを送信し、レスポンスを受信する
* Upstream: Envoy からリクエストを受信し、レスポンスを送信する


