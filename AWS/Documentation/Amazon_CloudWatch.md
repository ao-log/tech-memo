
# CloudWatch

## CloudWatch Metrics

[Amazon CloudWatch の概念](https://docs.aws.amazon.com/ja_jp/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html)

**名前空間**

命名規則は AWS/service。

**メトリクスの保持**

* 期間が 60 秒未満のデータポイントは、3 時間使用可能。
* 期間が 60 秒 (1 分) のデータポイントは、15 日間使用可能。
* 期間が 300 秒 (5 分) のデータポイントは、63 日間使用可能。
* 期間が 3600 秒 (1 時間) のデータポイントは、455 日 (15 か月) 間使用可能。

**ディメンション**

メトリクスのアイデンティティの一部である名前と値のペア。「Server=Prod, Domain=Frankfurt」のように複数の名前/値のペアから構成される場合もある。

**統計**

* Minumum
* Maximum
* Sum
* Average
* SampleCount

など。

**単位**

Bytes、Seconds、Count、Percent など。

**期間**

各統計は、指定された期間に収集されたメトリクスデータの集約を表している。



## Dashboard

[Amazon CloudWatch ダッシュボードの使用](https://docs.aws.amazon.com/ja_jp/AmazonCloudWatch/latest/monitoring/CloudWatch_Dashboards.html)



## Metrics


[カスタムメトリクスを発行する](https://docs.aws.amazon.com/ja_jp/AmazonCloudWatch/latest/monitoring/publishingMetrics.html)

独自のメトリクスを送信可能。



## Alarm

[Amazon CloudWatch でのアラームの使用](https://docs.aws.amazon.com/ja_jp/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html)

次の３つの状態がある。

* OK – メトリクスや式は、定義したしきい値を下回っている。
* ALARM – メトリクスや式は、定義したしきい値を超えている。
* INSUFFICIENT_DATA – アラームが開始直後であるか、メトリクスが利用できないか、データが不足していてアラームの状態を判定できない。

**欠落データの処理方法**

接続が失われた場合、サーバーがダウンした場合、メトリクスがデータを断続的にのみ報告する設計になっている場合など、データポイントがレポートされない場合がある。欠落データについては次のいずれかの処理方法を行うように設定できる。デフォルトは missing。

* notBreaching – 良好として扱う。
* breaching – 不良として扱う。
* ignore – 現在のアラーム状態を維持する。
* missing – アラーム評価範囲内のすべてのデータポイントがない場合、アラームは INSUFFICIENT_DATA に移行。

**その他の機能**

* アラームの有効/無効化: AWS CLI の disable-alarm-actions および enable-alarm-actions コマンドを使用
* アラームの動作をテスト: AWS CLI の set-alarm-state コマンド
* アラームの履歴表示: AWS CLI の describe-alarm-history コマンド


[インスタンスを停止、終了、再起動、または復旧するアラームを作成する](https://docs.aws.amazon.com/ja_jp/AmazonCloudWatch/latest/monitoring/UsingAlarmActions.html)

アラーム状態になった場合に、インスタンスを停止、終了、再起動、復旧するように設定することが可能。



## CloudWatch Agent

[CloudWatch エージェントを使用して Amazon EC2 インスタンスとオンプレミスサーバーからメトリクスとログを収集する](https://docs.aws.amazon.com/ja_jp/AmazonCloudWatch/latest/monitoring/Install-CloudWatch-Agent.html)

CloudWatch エージェントをインストールして稼働させることでメモリ使用量などのメトリクスも採取できるようになる。また、CloudWatch Logs へのログ送信も設定可能。


[CloudWatch エージェントのインストール](https://docs.aws.amazon.com/ja_jp/AmazonCloudWatch/latest/monitoring/install-CloudWatch-Agent-on-EC2-Instance.html)

CloudWatch エージェントのインストール方法。


[ウィザードを使用して CloudWatch エージェント設定ファイルを作成する](https://docs.aws.amazon.com/ja_jp/AmazonCloudWatch/latest/monitoring/create-cloudwatch-agent-configuration-file-wizard.html)

設定ファイルはウィザードを使用して対話式で答えていくことで生成することが可能。


[CloudWatch エージェント設定ファイルを手動で作成または編集する](https://docs.aws.amazon.com/ja_jp/AmazonCloudWatch/latest/monitoring/CloudWatch-Agent-Configuration-File-Details.html)

設定ファイルのリファレンス。


[CloudWatch エージェントにより収集されるメトリクス](https://docs.aws.amazon.com/ja_jp/AmazonCloudWatch/latest/monitoring/metrics-collected-by-CloudWatch-agent.html)

収集されるメトリクスの一覧。



## CloudWatch Logs

[ロググループとログストリームを操作する](https://docs.aws.amazon.com/ja_jp/AmazonCloudWatch/latest/logs/Working-with-log-groups-and-streams.html)

ロググループ、ログストリームの階層構造になっている。

* ロググループの設定上、デフォルトでは永遠に保存する。保管期間は変更可能。


[フィルターを使用したログイベントからのメトリクスの作成](https://docs.aws.amazon.com/ja_jp/AmazonCloudWatch/latest/logs/MonitoringLogData.html)

ログをフィルタリングして CloudWatch メトリクスを生成できる。いくつか例がある。

* [例: 語句の出現回数をカウントする](https://docs.aws.amazon.com/ja_jp/AmazonCloudWatch/latest/logs/CountOccurrencesExample.html)


[サブスクリプションを使用したログデータのリアルタイム処理](https://docs.aws.amazon.com/ja_jp/AmazonCloudWatch/latest/logs/Subscriptions.html)

Amazon Kinesis ストリーム、Amazon Kinesis Data Firehose ストリーム、AWS Lambda などの他のサービスに配信することが可能。


[Amazon S3 へのログデータのエクスポート](https://docs.aws.amazon.com/ja_jp/AmazonCloudWatch/latest/logs/S3Export.html)

Amazon S3 へエクスポート可能。



## アプリケーションパフォーマンスモニタリング (APM)

[アプリケーションパフォーマンスモニタリング (APM)](https://docs.aws.amazon.com/ja_jp/AmazonCloudWatch/latest/monitoring/CloudWatch-Application-Monitoring-Intro.html)

* メトリクス、トレースなどの収集ができる


[カスタムセットアップを使用して Amazon ECS で Application Signals を有効にする - サイドカー戦略を使用してデプロイする](https://docs.aws.amazon.com/ja_jp/AmazonCloudWatch/latest/monitoring/CloudWatch-Application-Signals-ECS-Sidecar.html)

* 次の流れ
  * サービスリンクロールを作成
  * ECS タスクロールに `CloudWatchAgentServerPolicy` をアタッチ
  * CloudWatch Agent の設定ファイルを SSM Parameter Store にアップロード
  * タスク定義設定
    * 各コンテナで `opentelemetry-auto-instrumentation-python` をマウント
    * init コンテナでファイルを所定のパスに配置
    * サイドカーで CloudWatch Agent を稼働
    * 環境変数にて各種設定変更可能



# BlackBelt

* [20190326 AWS Black Belt Online Seminar Amazon CloudWatch](https://pages.awscloud.com/rs/112-TZM-766/images/20190326_AWS-BlackBelt_CloudWatch.pdf)

* CloudWatch
  * P13:
    * メトリクス: 時系列のデータポイント
    * 名前空間: 例) AWS/EC2
    * ディメンション: 例) instance-id=i-xxxxxxxx
  * P19: EC2 インスタンスは基本モニタリングで 5 分、詳細モニタリグで 1 分。取得頻度ごとに使用期間が決まっている。1 分未満だと 3 時間。
  * P20: 取得される情報は統計情報。Maximum, Minimum, Sum, Average など。
* CloudWatch Alarms
  * P29: アラームの状態は OK, ALARM, INSUFFICIENT_DATA の 3 種類
* CloudWatch Logs
  * P42: 階層はロググループ、ログストリーム、ログイベント。



# 参考

* Document
  * [Amazon CloudWatch とは](https://docs.aws.amazon.com/ja_jp/AmazonCloudWatch/latest/monitoring/WhatIsCloudWatch.html)
  * [Amazon CloudWatch Logs とは](https://docs.aws.amazon.com/ja_jp/AmazonCloudWatch/latest/logs/WhatIsCloudWatchLogs.html)
* サービス紹介ページ
  * [Amazon CloudWatch](https://aws.amazon.com/jp/cloudwatch/)
  * [よくある質問](https://aws.amazon.com/jp/cloudwatch/faqs/)
* [ナレッジセンター](https://aws.amazon.com/jp/premiumsupport/knowledge-center/#Amazon_CloudWatch)
* Black Belt
  * [20190326 AWS Black Belt Online Seminar Amazon CloudWatch](https://pages.awscloud.com/rs/112-TZM-766/images/20190326_AWS-BlackBelt_CloudWatch.pdf)

