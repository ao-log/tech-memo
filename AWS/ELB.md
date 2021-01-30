
# ELB

3 種類のロードバランサーをサポート。

* Application Load Balancer
* Network Load Balancer
* Classic Load Balancer



## ELB 全体のユーザーガイド

[Elastic Load Balancing とは](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/userguide/what-is-load-balancing.html)


[Elastic Load Balancing の仕組み](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/userguide/how-elastic-load-balancing-works.html)

* ALB は複数の AZ を有効にする必要がある。
* **クロスゾーン負荷分散**は ALB では有効になっている。NLB ではデフォルト無効。
* ELB は 1 つ以上の IP アドレスを返す。ALB は IP アドレスが変わる場合がある。
* ALB はラウンドロビン。NLB はフローハッシュアルゴリズム(プロトコル、送信元 IP アドレス・ポート、送信先 IP アドレス・ポート、TCP シーケンス番号に基づく)
* HTTP 接続
  * ALB はクライアントからは HTTP/0.9, HTTP/1.0, HTTP/1.1, HTTP/2.0 をサポート。バックエンド接続では HTTP/1.1 を使用。Keep-Alive はバックエンド接続でサポートされている。
  * ALB は X-Forwarded-For、X-Forwarded-Proto、および X-Forwarded-Port ヘッダをリクエストに追加。
* 作成時に、インターネット向け、内部向けのどちらかを選択。



## 製品比較

詳細は [製品の比較](https://aws.amazon.com/jp/elasticloadbalancing/features/#compare) を参照すること。



## ALB

[Application Load Balancer とは](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/application/introduction.html)


[チュートリアル: AWS CLI を使用して Application Load Balancer を作成する](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/application/tutorial-application-load-balancer-cli.html)

```shell
# ロードバランサの作成
aws elbv2 create-load-balancer --name my-load-balancer \
    --subnets subnet-0e3f5cac72EXAMPLE subnet-081ec835f3EXAMPLE  \
    --security-groups sg-07e8ffd50fEXAMPLE

# ターゲットグループの作成
aws elbv2 create-target-group --name my-targets \
    --protocol HTTP \
    --port 80 \
    --vpc-id vpc-0598c7d356EXAMPLE

# ターゲットグループへ登録
aws elbv2 register-targets --target-group-arn targetgroup-arn  \
    --targets Id=i-0abcdef1234567890 Id=i-1234567890abcdef0

# リスナー作成
aws elbv2 create-listener --load-balancer-arn loadbalancer-arn \
    --protocol HTTP --port 80  \
    --default-actions Type=forward,TargetGroupArn=targetgroup-arn

# リスナールール追加
aws elbv2 create-rule --listener-arn listener-arn --priority 10 \
    --conditions Field=path-pattern,Values='/img/*' \
    --actions Type=forward,TargetGroupArn=targetgroup-arn
```


#### ロードバランサー

[ロードバランサー](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/application/application-load-balancers.html)

* 少なくとも 2 つのサブネットの作成が必要

以下の属性がある。

* access_logs.s3.enabled
* access_logs.s3.bucket
* access_logs.s3.prefix
* deletion_protection.enabled: 削除保護の有効化
* idle_timeout.timeout_seconds: アイドルタイムアウト
* routing.http.desync_mitigation_mode
* routing.http.drop_invalid_header_fields.enabled
* routing.http2.enabled

**アイドルタイムアウト**

接続はクライアント-ELB 間、ELB-バックエンド間の２つがある。
アイドルタイムアウトで指定した秒数以内にデータが送受信されなかった場合は、ロードバランサは接続を閉じる。

バックエンドでは EC2 インスタンス側で Keep-Alive を有効にすることを推奨。また、ロードバランサの設定値よりも長くすることを推奨。
Keep-Alive が有効の場合、タイムアウト期間が終了するまでバックエンド接続を再利用できる。


#### リスナー

[リスナー](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/application/load-balancer-listeners.html)

リスナールールは以下のものを設定可能。

* authenticate-cognito
* authenticate-oidc
* fixed-response
* forward
* redirect

ルールの条件は以下のものを設定可能。

* host-header
* http-header
* http-request-method
* path-pattern
* query-string
* source-ip


[Application Load Balancer の HTTPS リスナーを作成する](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/application/create-https-listener.html)

* 用意されたセキュリティポリシーを設定できる。セキュリティポリシーはプロトコルと暗号の組み合わせ。


#### ターゲットグループ

[ターゲットグループ](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/application/load-balancer-target-groups.html)

ターゲットの種類は次の３つ。

* instance
* ip
* lambda

次の属性がある(ターゲットが instance, ip の場合)

* deregistration_delay.timeout_seconds: **登録解除の遅延**の時間
* load_balancing.algorithm.type: ロードバランサがターゲットを選択する方法。デフォルトは round_robin。least_outstanding_requests も設定可能。
* slow_start.duration_seconds: スロースタートの期間
* stickiness.enabled
* stickiness.lb_cookie.duration_seconds
* stickiness.type


[ヘルスチェック](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/application/target-group-health-checks.html)

以下の設定がある。

* HealthCheckProtocol: HTTP or HTTPS
* HealthCheckPort
* HealthCheckPath: デフォルトは /
* HealthCheckTimeoutSeconds: 失敗とみなすターゲットからの応答がない時間。instance, ip の場合はデフォルト 5 秒。
* HealthCheckIntervalSeconds: 間隔
* HealthyThresholdCount: healthy とみなすまでの連続成功回数
* UnhealthyThresholdCount: unhealthy とみなすまでの連続失敗回数
* Matcher: HTTP ステータスコード。デフォルトは 200

以下のステータスがある。

* initial
* healthy
* unhealthy
* unused
* draining
* unavailable

理由コード。Elb で始まるものはロードバランサ側、Target で始まるものはターゲット側で発生したもの。

* Elb.InitialHealthChecking
* Elb.InternalError
* Elb.RegistrationInProgress
* Target.DeregistrationInProgress
+ Target.FailedHealthChecks
* Target.HealthCheckDisabled
* Target.InvalidState


[ターゲットの登録](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/application/target-group-register-targets.html)

IP アドレスの場合、登録する IP アドレスはターゲットは次の CIDR ブロック内のアドレスである必要がある。

* 10.0.0.0/8 (RFC 1918)
* 100.64.0.0/10 (RFC 6598)
* 172.16.0.0/12 (RFC 1918)
* 192.168.0.0/16 (RFC 1918)


#### モニタリング

[監視](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/application/load-balancer-monitoring.html)


[CloudWatch メトリクス](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/application/load-balancer-cloudwatch-metrics.html)

名前空間は AWS/ApplicationELB。
ディメンションは AvailabilityZone、LoadBalancer、TargetGroup。

メトリクスの統計

* Minimum 統計と Maximum 統計には、個々のロードバランサーノードによって報告される最小値と最大値が反映される。
* Sum 統計は、すべてのロードバランサーノードにおける集計値。

ロードバランサのメトリクス。

* HTTPCode_ELB_4XX_Count: ロードバランサーから送信される HTTP 4XX クライアントエラーコードの数
* HTTPCode_ELB_5XX_Count: ロードバランサーから送信される HTTP 5XX サーバーエラーコードの数。
* HTTPCode_ELB_500_Count、HTTPCode_ELB_502_Count、HTTPCode_ELB_503_Count、HTTPCode_ELB_504_Count
など

ターゲットのメトリクス。

* HealthyHostCount
* UnHealthyHostCount
* HTTPCode_Target_2XX_Count、HTTPCode_Target_3XX_Count、HTTPCode_Target_4XX_Count、HTTPCode_Target_5XX_Count
* RequestCountPerTarget
* TargetConnectionErrorCount: ロードバランサとターゲット間で正常に確立されなかった接続数。
* TargetResponseTime: リクエストがロードバランサーから送信され、ターゲットからの応答を受信するまでの経過時間 (秒)。


[アクセスログ](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/application/load-balancer-access-logs.html)

アクセスログの構文などが書かれている。


#### トラブルシューティング

[トラブルシューティング](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/application/load-balancer-troubleshooting.html)

HTTP ステータスコードごとのトラブルシューティング方法がある。




## BlackBelt

[20191029 AWS Black Belt Online Seminar Elastic Load Balancing (ELB)](https://www.slideshare.net/AmazonWebServicesJapan/20191029-aws-black-belt-online-seminar-elastic-load-balancing-elb)

* 概要
  * P16:
    * リスナーのプロトコル(ALB: HTTP/HTTPS、NLB: TCP, UDP, ILS, TCP_UDP)
    * ターゲットの種類(インスタンス ID、IP アドレス、Lambad 関数)
  * P17: Zone Apex では CNAME 設定不可。ALIAS は可能。
* 高可用性と負荷分散
  * P21: バックエンドへのルーティングアルゴリズム(ALB: ラウンドロビン, NLB: フローハッシュアルゴリズム)
  * P22: クロスゾーン負荷分散(ALB: デフォルトで有効, NLB: デフォルトで無効)
  * P25: ALB のスケールが間に合わない場合は 503 を返す。
* ELB のモニタリング、ログ
  * P27: ヘルスチェックを設定可能(プロトコル、ポート番号、パス、タイムアウト、正常のしきい値、非正常のしきい値、間隔など)
  * P28: CloudWatch メトリクス(HealthyHostCount、UnHealthyHostCount、RequestCount、Latency など)
  * P29: アクセスログ
* コネクション
  * P31: コネクションタイムアウト(ALB はデフォルト 60 秒。NLB は 350 秒固定)
  * P32: 登録解除の遅延。デフォルトは 300 秒。
  * P33: ステッキーセッション
* セキュリティ
  * P35: SSL/TLS Termination
  * P36: 事前定義されたキュリティポリシー
  * P37: TLS サーバ証明書
  * P38: SNI での複数 TLS 証明書のスマートセレクション
* ALB
  * P40: 特徴(HTTP/HTTPS のみ対応、コンテンツベースのルーティング)
  * P41: コンポーネント(リスナー、ルール、ターゲットグループ、ターゲット)
  * P42: コンテンツベースのルーティング(パス、ホストベース、HTTP ヘッダ、HTTP メソッド)
  * P48: クライアントの IP アドレスは X-Forward-For を使用して確認が必要
* NLB
  * P51: 特徴(固定 IP アドレス、送信元 IP アドレスの保持、暖気なしにスパイクに対応可能、セキュリティグループは設定不能)
  * P55: 注意点
    * 古いインスタンスタイプをターゲットにできない
    * アイドルコネクションタイムアウトは 350 秒固定
    * ヘルスチェックはタイムアウトは 10 秒固定。間隔は 10 秒 or 30 秒で作成後の変更不可
    * セキュリティグループなし
* 他サービスとの連携
  * P60: Auto Scaling との連携(インスタンスの登録・登録解除、ELB のヘルスチェック結果の使用)
  * P61: ECS タスクをターゲットに設定可能(EC2 起動タイプの場合は動的ポートマッピングを使用可能)
  * P64: WAF との連携
  * P65: Global Acceleator との連携
  * P66: Lambda 関数をターゲットに設定



# 参考

* AWS Document
  * [Elastic Load Balancing とは](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/userguide/what-is-load-balancing.html)
  * [Application Load Balancer とは](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/application/introduction.html)
  * [Network Load Balancer とは](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/network/introduction.html)
  * [API Reference](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/APIReference/API_Operations.html)
* サービス紹介ページ
  * [Elastic Load Balancing](https://aws.amazon.com/jp/elasticloadbalancing/)
  * [よくある質問](https://aws.amazon.com/jp/elasticloadbalancing/faqs/)
* [ナレッジセンター](https://aws.amazon.com/jp/premiumsupport/knowledge-center/#Elastic_Load_Balancing)
* Black Belt
  * [20191029 AWS Black Belt Online Seminar Elastic Load Balancing (ELB)](https://www.slideshare.net/AmazonWebServicesJapan/20191029-aws-black-belt-online-seminar-elastic-load-balancing-elb)

