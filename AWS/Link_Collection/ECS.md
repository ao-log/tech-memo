
### AWS Blog

#### Categories

[Containers](https://aws.amazon.com/jp/blogs/containers/)

[Category: Amazon Elastic Container Service](https://aws.amazon.com/jp/blogs/news/category/compute/amazon-elastic-container-service/)



#### Blogs

[詳解: Amazon ECSのタスクネットワーク](https://aws.amazon.com/jp/blogs/news/under-the-hood-task-networking-for-amazon-ecs/)

* ネットワークスタックはネットワークの namespace を通じて設定される。
* ECS Agent から CNI プラグインを呼び出して操作。複数のプラグインを順に呼び出す。
* 最初に呼び出されるものは ecs-eni プラグイン。ENI を namespace にアタッチ。
* ecs-bridge と ecs-ipamプラグインにより認証情報エンドポイントに HTTP リクエストできるように設定。
* 処理の流れ
  * eth1 を作成。
  * eth1 をタスク用の namespace に移動。
* pause コンテナを作成。ネットワークスタックの設定とコンテナ内で実行されるコマンドのレースコンディションを防ぐため。
* pause コンテナの network namespace をセットアップ。
* タスク定義の各コンテナを開始。


[詳細: Fargate データプレーン](https://aws.amazon.com/jp/blogs/news/under-the-hood-fargate-data-plane/)

* Docker Engine を Containerd に置き換えた

他にも Fargate の実装にあたっては以下のカスタマイズをおこなっている。

* awsvpc のサポート。Container Networking Interface (CNI) プラグインを利用しネットワークをセットアップ。
* FireLens のサポート。
* ECS タスクのメタデータエンドポイント。統計情報などを採取できるようにローカルの HTTP エンドポイントを提供。

以下の特徴がある。

* pause コンテナは不要となった。Containerd は、コンテナの起動時に外部で作成されたネットワーク名前空間の使用をサポートしているため。
* さまざまな宛先へのログ送信のサポート。Containerd の Shim ログ記録プラグインを拡張している。amazon-ecs-shim-loggers-for-containerd リポジトリ。
* ランタイムは runC のほか FireCracker に切り替えられる構造。
* Fargate エージェント。

コンテナ起動時の流れがシーケンス図で示されている。

* network namespace の作成
* Customer ENI のプロビジョニング
* AWS Secrets Manager からシークレットの取得
* コンテナイメージのプル
* containerd へコンテナ開始の指示
* コンテナランタイムへコンテナ開始の指示。logging shim plugin の開始。

Fargate データプレーンは Fargate Agent, Containerd。こちらは Fargate VPC に繋がっている。
カスタマータスクゾーンにて container runtime shim を介して runc によりコンテナが起動。Customer ENI により Customer の VPC に繋がっている。


[ECS のアプリケーションを正常にシャットダウンする方法](https://aws.amazon.com/jp/blogs/news/graceful-shutdowns-with-ecs/)

* エントリプロセス
  * コンテナでは、Dockerfile の ENTRYPOINT および CMD ディレクティブで指定されたプロセス (エントリプロセスと呼ぶ) が、コンテナ内の他のすべてのプロセスの親になる。
  * ENTRYPOINT が bin/sh -c my-app の場合は、sh はコンテナが停止したときに SIGTERM シグナルを受信するものの my-app に渡されない。
  * shell を安全に使用するには２つの方法がある。
    * 1) シェルスクリプトを介して実行される実際のアプリケーションに exec というプレフィックスをつける。
    * 2) tini (Amazon ECS-Optimized AMI の Docker ランタイムに同梱されている) や dumb-init などの専用のプロセスマネージャーを使用する。
  * あるコマンドを exec をつけて実行すると、子プロセスを生成するのではなく現在実行中のプロセス (この場合は shell ) の内容を新しい実行可能プロセスに置き換える。
  * tini や dumb-init などのプロセスマネージャーを使用すると、これらのプログラムは SIGTERM を受け取ると、アプリケーションを含むすべての子プロセスグループに SIGTERM を送信する。
* SIGTERM シグナルの処理
  * デフォルトの停止シグナルは SIGTERM だが、Dockerfile に STOPSIGNAL ディレクティブを追加することで上書きできる。
  * ALB から draining され unused となったあとに SIGTERM が実行される（単一の ALB に登録されている場合）。
  * EC2 インスタンスが drain 状態になっても RunTask によって実行されたタスクは draining されない。
* スポットインスタンスの中断
  * ECS_ENABLE_SPOT_INSTANCE_DRAINING を true に設定することでスポットインスタンスの中断通知を受信すると、ECS エージェントはインスタンスを DRAINING 状態にする
  * スポットインスタンスの中断通知を受信すると ALB から登録解除される。
  * 登録解除後に SIGTERM なので、登録解除の遅延は 120 秒未満にする必要がある。


[詳解: Amazon Elastic Container Service と AWS Fargate のタスク起動レートの向上](https://aws.amazon.com/jp/blogs/news/under-the-hood-amazon-elastic-container-service-and-aws-fargate-increase-task-launch-rates/)


[詳解 FireLens – Amazon ECS タスクで高度なログルーティングを実現する機能を深く知る](https://aws.amazon.com/jp/blogs/news/under-the-hood-firelens-for-amazon-ecs-tasks/)

* FireLens を使えば、ファイルを編集して S3 に再アップロードするだけでよい。イメージをビルドし直す必要はない。
* コンテナの標準出力ログは、Fluentd Docker ログドライバーを介して Unix ソケット経由で FireLens コンテナに送信される。
* FireLens コンテナは、Fluentd Forward Protocol メッセージを TCP ソケットで LISTEN している
* タスク起動時に設定ファイルが自動設定される。
  * ログソース。ログソースは Unix および TCP ソケット
  * ECS メタデータを追加するトランスフォーマー
  * カスタムログを include
  * タスク定義で設定した内容に応じて OUTPUT プラグインの設定


[Fluent Bit による集中コンテナロギング](https://aws.amazon.com/jp/blogs/news/centralized-container-logging-fluent-bit/)


[Amazon ECS クラスターの Auto Scaling を深く探る](https://aws.amazon.com/jp/blogs/news/deep-dive-on-amazon-ecs-cluster-auto-scaling/)


[新機能 – AWS ECS Cluster Auto ScalingによるECSクラスターの自動スケーリング](https://aws.amazon.com/jp/blogs/news/aws-ecs-cluster-auto-scaling-is-now-generally-available/)


[カスタムメトリクスを用いた Amazon Elastic Container Service (ECS) のオートスケーリング](https://aws.amazon.com/jp/blogs/news/amazon-elastic-container-service-ecs-auto-scaling-using-custom-metrics/)

* リクエストをポーリング、処理するようなワークロードではカスタムメトリクスによる対応が必要。
* スケーリングメトリクスを計算する AWS Lambda 関数をトリガーする。EventBridge によって定期的に実行し、SQS メトリクスをポーチングし、ECS タスクで実行中のキャパシティを考慮して計算する。カスタムメトリクスとして CloudWatch Metrics に送信。
* ターゲット追跡スケーリングポリシーでカスタムメトリクスを使用。


[CloudWatch と Prometheus のカスタムメトリクスに基づく Amazon ECS サービスのオートスケーリング](https://aws.amazon.com/jp/blogs/news/autoscaling-amazon-ecs-services-based-on-custom-cloudwatch-and-prometheus-metrics/)カスタム


[Amazon ECS向けAmazon CloudWatch Container Insightsについて](https://aws.amazon.com/jp/blogs/news/introducing-container-insights-for-amazon-ecs/)


[AWS Distro for OpenTelemetry コレクターを使用したクロスアカウントの Amazon ECS メトリクス収集](https://aws.amazon.com/jp/blogs/news/using-aws-distro-for-opentelemetry-collector-for-cross-account-metrics-collection-on-amazon-ecs/)


[Bottlerocket のセキュリティ機能 〜オープンソースの Linux ベースオペレーティングシステム〜](https://aws.amazon.com/jp/blogs/news/security-features-of-bottlerocket-an-open-source-linux-based-operating-system/)


[Amazon ECS でのデーモンサービスの改善](https://aws.amazon.com/jp/blogs/news/improving-daemon-services-in-amazon-ecs/)


[Amazon ECS と AWS Fargate を利用した Twelve-Factor Apps の開発](https://aws.amazon.com/jp/blogs/news/developing-twelve-factor-apps-using-amazon-ecs-and-aws-fargate/)


[New – Amazon ECS Exec による AWS Fargate, Amazon EC2 上のコンテナへのアクセス](https://aws.amazon.com/jp/blogs/news/new-using-amazon-ecs-exec-access-your-containers-fargate-ec2/)


[Amazon ECS deployment circuit breaker のご紹介](https://aws.amazon.com/jp/blogs/news/announcing-amazon-ecs-deployment-circuit-breaker-jp/)


[AWS Cloud Map:アプリケーションのカスタムマップの簡単な作成と維持](https://aws.amazon.com/jp/blogs/news/aws-cloud-map-easily-create-and-maintain-custom-maps-of-your-applications/)


[AWS App Mesh を使用した Amazon ECS でのカナリアデプロイパイプラインの作成](https://aws.amazon.com/jp/blogs/news/create-a-pipeline-with-canary-deployments-for-amazon-ecs-using-aws-app-mesh/)


[AWS CodeDeploy による AWS Fargate と Amazon ECS でのBlue/Greenデプロイメントの実装](https://aws.amazon.com/jp/blogs/news/use-aws-codedeploy-to-implement-blue-green-deployments-for-aws-fargate-and-amazon-ecs/)


[Amazon ECR をソースとしてコンテナイメージの継続的デリバリパイプラインを構築する](https://aws.amazon.com/jp/blogs/news/build-a-continuous-delivery-pipeline-for-your-container-images-with-amazon-ecr-as-source/)


[Amazon ECS on AWS Fargate を利用したコンテナイメージのビルド](https://aws.amazon.com/jp/blogs/news/building-container-images-on-amazon-ecs-on-aws-fargate/)


[Amazon ECS on AWS Fargate のコスト最適化チェックリスト](https://aws.amazon.com/jp/blogs/news/cost-optimization-checklist-for-ecs-fargate/)

* タグを使用するにはアカウント設定で「新しい Amazon リソースネーム (ARN) とリソース識別子 (ID) 形式をオプトイン」しておく必要がある。
* コストエクスプローラーによって可視化。
* Savings Plans によるコスト削減。
* Fargate Spot の使用。
* タスクの適切なサイジング。
* Auto Scaling によって適切なタスク数を稼働。
* 営業時間外にタスクを停止するようにスケジューリング。


[Amazon ECS Fargate/EC2 起動タイプでの理論的なコスト最適化手法](https://aws.amazon.com/jp/blogs/news/theoretical-cost-optimization-by-amazon-ecs-launch-type-fargate-vs-ec2/)



## Black Belt

[[AWS Black Belt Online Seminar] AWS コンテナサービス開始のおしらせ](https://aws.amazon.com/jp/blogs/news/aws-bb-containers-start/)


[[AWS Black Belt Online Seminar] CON246 ログ入門 資料公開](https://aws.amazon.com/jp/blogs/news/aws-black-belt-online-seminar-con246-log/)


[[AWS Black Belt Online Seminar] CON245 Configuration & Secret Management 入門 資料公開](https://aws.amazon.com/jp/blogs/news/aws-black-belt-online-seminar-con245-config/)


[202109 AWS Black Belt Online Seminar Auto Scaling in ECS](https://www.slideshare.net/AmazonWebServicesJapan/202109-aws-black-belt-online-seminar-auto-scaling-in-ecs-250178830)


[202109 AWS Black Belt Online Seminar Amazon ECS Capacity Providers](https://www.slideshare.net/AmazonWebServicesJapan/202109-aws-black-belt-online-seminar-amazon-ecs-capacity-providers)


[202109 AWS Black Belt Online Seminar Amazon Elastic Container Service − EC2 スポットインスタンス / Fargate Spot ことはじめ](https://www.slideshare.net/AmazonWebServicesJapan/202109-aws-black-belt-online-seminar-amazon-elastic-container-service-ec-fargate-spot)



### tori さん

[Amazon ECS でのコンテナデプロイの高速化](https://toris.io/2021/04/speeding-up-amazon-ecs-container-deployments/)

コンテナのデプロイ速度を高速化するテクニックが書かれている。

* ヘルスチェックの間隔、しきい値を少ない値に設定する。
* 登録解除の遅延を少ない値に設定する。
* ECS_CONTAINER_STOP_TIMEOUT を少ない値に設定する。（SIGTERM 後に SIGKILL を送るまでの時間）
* ECS_IMAGE_PULL_BEHAVIOR を prefer-cached あるいは once に設定する。イメージタグをコミットごとに変えている場合に有効な戦略(latest はアンチパターン)
* minimumHealthyPercent を小さめに maximumPercent を多めに設定することで、より少ないステップ数でデプロイが完了するようにする。


[アプリケーション開発者は Amazon ECS あるいは Kubernetes をどこまで知るべきか](https://speakerdeck.com/toricls/you-build-it-you-run-it)




## ECS Agent

[aws/amazon-ecs-agent](https://github.com/aws/amazon-ecs-agent)

[amazon-ecs-agent/proposals/eni.md](https://github.com/aws/amazon-ecs-agent/blob/master/proposals/eni.md)



## 記事

[[アップデート] 実行中のコンテナに乗り込んでコマンドを実行できる「ECS Exec」が公開されました](https://dev.classmethod.jp/articles/ecs-exec/)

使用するには以下の設定が必要。

* タスクロールに SSM 関連の権限を追加
* ECSサービスで「enableExecuteCommand」の設定を有効にする

以下のコマンドで接続。

```
$ aws ecs execute-command \
    --cluster クラスター名 \
    --task XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX \
    --container nginx \
    --interactive \
    --command "コマンド"
```


[【レポート】Amazon ECS deployment circuit breaker を使った自動ロールバック #AWSSummit](https://dev.classmethod.jp/articles/awssummit2021-ecs-deployment-circuit-breaker/)

Amazon ECS deployment circuit breaker

* デプロイタイプ、ローリングアップデートでサポートされる。
* ステージ 1, 2 の 2 段階で評価。
  * ステージ1: タスクが RUNNING 状態に遷移することなく停止した場合。
  * ステージ2: RUNNING 状態に遷移したタスクがヘルスチェックに合格した場合。
* しきい値は 10 <= Desired Count * 0.5 <= 200。


[詳細解説「AWS Cloud Map」とは #reinvent](https://dev.classmethod.jp/articles/cloud-map-perfect/)

* AWS Cloud Mapへの問い合わせ方式
  * Cloud Mapへの問い合わせ方式は、2種類。
    * DNSクエリ
    * APIコール
      * AWSのARN（Amazon Resource Name）などでアクセスするリソースに対して付与
* 設定が必要なリソース
  * 名前空間
  * サービス
  * サービスインスタンス
* 名前空間
  * 次のいずれかを選択。
    * API呼び出し
    * API呼び出しとVPCのDNSクエリ: VPC の指定も必要
    * API呼び出しと公開DNSクエリ：インターネットからの名前解決が可能。Route 53 のヘルスチェックが利用可能。
* サービス
  * サブドメインのようなイメージ
  * ルーティングポリシーを設定
  * ヘルスチェック方法を設定
* サービスインスタンス
  * 3 つ指定方法がある
    * IP アドレス
    * CNAME
    * リソースを特定するための情報(ARN など)


[CloudWatch Logs Insights でコンテナ単位のCPU・メモリ使用量などを確認する](https://dev.classmethod.jp/articles/ways-to-check-fargate-cpu-usage/)

* Container Insight の [View performance logs] から Logs Insights に遷移する。既にクエリが入力済みの状態になっている。


[Container Insights でコンテナ単位のCPU・メモリ使用率を表示させる](https://dev.classmethod.jp/articles/how-to-check-container-cpu-usage-by-container-insights/)

* コンテナ単位で表示させるにはタスク定義でコンテナの CPU、メモリを設定する必要がある。




