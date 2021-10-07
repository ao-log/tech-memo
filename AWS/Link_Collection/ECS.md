
### AWS Blog

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


[詳解: Amazon ECSのタスクネットワーク](https://aws.amazon.com/jp/blogs/news/under-the-hood-task-networking-for-amazon-ecs/)


[詳解 FireLens – Amazon ECS タスクで高度なログルーティングを実現する機能を深く知る](https://aws.amazon.com/jp/blogs/news/under-the-hood-firelens-for-amazon-ecs-tasks/)

* FireLens を使えば、ファイルを編集して S3 に再アップロードするだけでよい。イメージをビルドし直す必要はない。
* コンテナの標準出力ログは、Fluentd Docker ログドライバーを介して Unix ソケット経由で FireLens コンテナに送信される。
* FireLens コンテナは、Fluentd Forward Protocol メッセージを TCP ソケットで LISTEN している
* タスク起動時に設定ファイルが自動設定される。
  * ログソース。ログソースは Unix および TCP ソケット
  * ECS メタデータを追加するトランスフォーマー
  * カスタムログを include
  * タスク定義で設定した内容に応じて OUTPUT プラグインの設定


[CloudWatch と Prometheus のカスタムメトリクスに基づく Amazon ECS サービスのオートスケーリング](https://aws.amazon.com/jp/blogs/news/autoscaling-amazon-ecs-services-based-on-custom-cloudwatch-and-prometheus-metrics/)


[AWS App Mesh を使用した Amazon ECS でのカナリアデプロイパイプラインの作成](https://aws.amazon.com/jp/blogs/news/create-a-pipeline-with-canary-deployments-for-amazon-ecs-using-aws-app-mesh/)


[Amazon ECS on AWS Fargate を利用したコンテナイメージのビルド](https://aws.amazon.com/jp/blogs/news/building-container-images-on-amazon-ecs-on-aws-fargate/)


[[AWS Black Belt Online Seminar] AWS コンテナサービス開始のおしらせ](https://aws.amazon.com/jp/blogs/news/aws-bb-containers-start/)



### tori さん

[Amazon ECS でのコンテナデプロイの高速化](https://toris.io/2021/04/speeding-up-amazon-ecs-container-deployments/)

コンテナのデプロイ速度を高速化するテクニックが書かれている。

* ヘルスチェックの間隔、しきい値を少ない値に設定する。
* 登録解除の遅延を少ない値に設定する。
* ECS_CONTAINER_STOP_TIMEOUT を少ない値に設定する。（SIGTERM 後に SIGKILL を送るまでの時間）
* ECS_IMAGE_PULL_BEHAVIOR を prefer-cached あるいは once に設定する。イメージタグをコミットごとに変えている場合に有効な戦略(latest はアンチパターン)
* minimumHealthyPercent を小さめに maximumPercent を多めに設定することで、より少ないステップ数でデプロイが完了するようにする。




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




