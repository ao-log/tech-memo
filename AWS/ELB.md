
# ELB

3 種類のロードバランサーをサポート。

* Application Load Balancer
* Network Load Balancer
* Classic Load Balancer


## 用語

#### リスナー

ユーザーが設定したプロトコルとポートを使用してクライアントからの接続リクエストを確認し、ユーザーが定義したルールに基づいて 1 つ以上のターゲットグループにリクエストを転送。

#### ターゲットグループ

指定されたプロトコルとポート番号を使用して、1 つ以上の登録済みのターゲット (EC2 インスタンスなど) にリクエストをルーティング。

## LB と関連するサービス

* Amazon EC2
* Amazon EC2 Auto Scaling
* AWS Certificate Manager 
* Amazon CloudWatch
* Amazon ECS
* Route 53
* AWS WAF

## 製品比較

詳細は [製品の比較](https://aws.amazon.com/jp/elasticloadbalancing/features/#compare) を参照すること。

## ALB

* プロトコルは HTTP、HTTPS、ポートは 1 ～ 65535 をサポート。詳細は、[リスナー設定](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/application/load-balancer-listeners.html) を参照。
* URL パスベースのルーティングをサポート。リスナールールの詳細は、[リスナールール](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/application/load-balancer-listeners.html#listener-rules) を参照。
* WebSocket、HTTP/2 のサポート
* [ターゲットグループへルーティング](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/application/load-balancer-target-groups.html)。Auto Scaling Group を設定可能。
* マルチ AZ を有効にする必要あり
* クロスゾーン負荷分散はデフォルトで有効
* [SSL証明書](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/application/create-https-listener.html#https-listener-certificates)
* [ターゲットグループのヘルスチェック](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/application/target-group-health-checks.html)
* [LB 自身のモニタリング](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/application/load-balancer-monitoring.html)
* [アクセスログ](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/application/load-balancer-access-logs.html)
* [トラブルシューティング](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/application/load-balancer-troubleshooting.html)

#### NLB

* プロトコルは TCP、TLS、UDP、TCP_UDP、ポートは 1 ～ 65535 をサポート。詳細は、[リスナー設定](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/network/load-balancer-listeners.html) を参照。
* [ターゲットグループへルーティング](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/network/load-balancer-target-groups.html)
* クロスゾーン負荷分散はデフォルトで無効
* [L4 の情報に基づいてルーティング](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/network/load-balancer-target-groups.html#target-group-routing-configuration)
* IP アドレスによるターゲット登録をサポート
* [SSL証明書](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/network/create-tls-listener.html#tls-listener-certificates)
* [ターゲットグループのヘルスチェック](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/network/target-group-health-checks.html)
* [LB 自身のモニタリング](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/network/load-balancer-monitoring.html)
* [アクセスログ](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/network/load-balancer-access-logs.html)
* [VPC フローログによるパケットキャプチャ](https://docs.aws.amazon.com/ja_jp/vpc/latest/userguide/flow-logs.html)
* [トラブルシューティング](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/network/load-balancer-troubleshooting.html)

#### CLB

* プロトコルは HTTP、HTTPS、TCP、SSL (セキュア TCP) をサポート。詳細は、[リスナーの設定](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/classic/elb-listener-config.html) を参照。
* ターゲットグループではなく、インスタンスへルーティング
* [クロスゾーン負荷分散可能](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/classic/enable-disable-crosszone-lb.html)
* [SSL証明書](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/classic/elb-listener-config.html#https-ssl-listeners)
* [インスタンスのヘルスチェック](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/classic/elb-healthchecks.html)
* [LB 自身のモニタリング](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/classic/elb-monitor-logs.html)
* [アクセスログ](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/classic/access-log-collection.html)
* [トラブルシューティング](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/classic/elb-troubleshooting.html)


## その他、LB の特徴、機能

* DNS で各 AZ へのルーティング
* ELB にセキュリティグループを設定可能

#### ELB のスケーリング

スケーリングが追いつかない場合、503 ERROR を返す。事前にわかっている場合は、暖機申請、もしくは徐々に負荷を増やしてスケールさせる。

#### インターネット向け or 内部向け

LB はインターネット向け、内部向け、両方を作成できる。

#### ドメイン名の設定

ALB、CLB は IP アドレスが変わる可能性あり。DNS を使用してアクセスすることを推奨。

* Route 53 の場合はエイリアスレコードを設定
* 通常の DNS サーバの場合、CNAME レコードを設定

設定方法

* [Classic Load Balancer のカスタムドメイン名を設定する](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/classic/using-domain-names-with-elb.html)

#### Route 53 DNS フェイルオーバ対応

ユースケースとしては、ELB 配下に正常な EC2 インスタンスがない場合に、S3 上の Sorry ページを参照させる。

#### クロスゾーン負荷分散

LB のあるゾーンだけでなく、有効な全ての AZ の登録済みターゲットに負荷分散する。

#### スティッキーセッション

同じクライアントのリクエストを同じインスタンスに割り振りする機能。ALB、CLB がサポート。

* [スティッキーセッション](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/classic/elb-sticky-sessions.html)

#### 接続元 IP アドレス

バックエンドサーバから見ると、接続元は ELB となる。HTTP ヘッダの「X-Forwarded-For」で参照できる。

#### アイドル接続のタイムアウト

データの送受信がなかった場合のタイムアウト秒数を 1 〜 3600 秒の間で設定可能。デフォルト 60 秒。

* [Classic Load Balancer のアイドル接続のタイムアウトを設定する](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/classic/config-idle-timeout.html)

#### Connection Draining

登録解除中のインスタンスまたは異常の発生したインスタンスにリクエストを送信しないようにする機能。既存の接続を開いたままなので、Graceful に停止できる。

* [Classic Load Balancer の Connection Draining を設定する](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/classic/config-conn-drain.html)


## 価格体系

* ALB の起動時間
* LCU(Load Balancer Capacity Units) の使用量

# 参考

* [AWS Black Belt Online Seminar 2016 Elastic Load Balancing](https://www.slideshare.net/AmazonWebServicesJapan/aws-black-belt-online-seminar-2016-elastic-load-balancing)
* [Elastic Load Balancing ドキュメント](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/?id=docs_gateway)


