
# EventBridge

[Amazon EventBridge とは](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/what-is-amazon-eventbridge.html)

アプリケーションをさまざまなソースからのデータと接続できるようにするサーバーレスイベントバスサービス。

CloudWatch Events と同じ API になっている。
サードパーティの SaaS アプリケーションからデータを接続できるように拡張し、Event Bridge という新しい名前でリリースされた。



## 用語

* イベント: 状態が変更した際などに発生。定期的にイベントを生成することも可能。
* ルール: 一致したイベントを検出し、ターゲットに振り分けるもの。
* ターゲット: Lambda 関数、SNS トピック、SQS キュー、StepFunctions、ECS タスクなどに振り分け可能。
* イベントバス: イベントを受信するバス。ルールが検出できるのは紐付いたイベントバスのもののみ。AWS アカウントにはデフォルトのイベントバスが一つある。
* パートナーイベントソース: AWS アカウントにイベントを送信するために使用される。

イベントソースからイベントが発生し、イベントバスに送信。ルールに一致したイベントについてターゲットに送信する流れ。

[Black Belt](https://www.slideshare.net/AmazonWebServicesJapan/20200122-aws-black-belt-online-seminar-amazon-eventbridge/21) に上記流れの図がある。


## はじめに

[AWSサービスのルールの作成](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/create-eventbridge-rule.html)

ルール作成時に設定する項目。

* パターンの定義
  * イベントパターン or スケジュール
* イベントバス
* ターゲット
  * イベント発行元のサービス名
  * 入力の設定
    * 一致したイベント、一致したイベントの一部、定数(JSON テキスト) or 入力トランスフォーマー
  * IAM ロールの選択（EventBridge がターゲットにイベントを送信する際に必要となるアクセス許可）
  * Retry ポリシーとデッドレターキュー:
    * Retry ポリシー
      * イベントの最大有効期限: 1 分から 24 時間の間の値
      * 再試行: 0〜185
    * デッドレターキュー
      * SQS を選択。使用しない場合は [None]


[AWS CloudTrail を使用して AWS API コールでトリガーする EventBridge ルールの作成](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/create-eventbridge-cloudtrail-rule.html)

CloudTrail に記録される特定の API コールをルールのトリガー元にすることが可能。


[スケジュールに従ってトリガーする EventBridge ルールの作成](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/create-eventbridge-scheduled-rule.html)

定期的なスケジュールをルールのトリガー元にすることが可能。


[イベントバスの作成](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/create-event-bus.html)

アカウントには一つのデフォルトイベントバスがある。

イベントバスには二種類ある。

* パートナーイベントバス: パートナーによって生成されたイベントを受信。
* カスタムイベントバス: イベントバスごとに 100 個のルール数上限がある。上限を超えて作成したい場合や、イベントバスごとにアクセス許可の設定内容を変えたい場合に使用。


[EventBridge ルールの削除または無効化](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/delete-or-disable-rule.html)

ルールは一時的に無効化することが可能。


[イベント再試行ポリシーとデッドレターキューの使用](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/rule-dlq.html)

イベントはターゲットに正常に配信できないことがある。原因はターゲットリソースが使用できない、権限がない、ネットワークの状態など。

正常に配信されない場合は再試行を行う。再試行の所要時間、回数は再試行ポリシーによって決まる。

配信失敗した場合はデッドレターキューに送信可能。
一部のイベントは再試行せずにデッドレターキューに送信される。例えば、アクセス権限がないような場合。


[EventBridgeアーカイブとイベントの再生](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/eb-archives.html)

アーカイブを作成し、イベントに一致するフィルタパターンを設定しておくことで、フィルタ条件に一致するイベントはアーカイブされる。

アーカイブされたイベントは再生することが可能。再生を作成する場合、イベントの再生元のアーカイブ、イベントの再生の開始時間と終了時間、イベントの再生先のターゲットイベントバスまたは 1 つ以上のルールを指定できる。


## チュートリアル

[チュートリアル: AWS Systems Manager 実行コマンドにイベントを中継](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/ec2-run-command.html)

Auto Scaling で起動されるインスタンスにコマンドを実行する例。

* イベント
  * イベントタイプ: インスタンスの起動と削除
  * EC2 Instance-launch Lifecycle Action
* ターゲット: SSM 実行コマンド
  * ドキュメン: AWS-RunShellScript
  * Target key: tag:environment
  * Target value(s): production


[チュートリアル: Amazon S3 オブジェクトレベル操作のログを記録する](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/log-s3-data-events.html)  

CloudTrail 側でデータイベントを記録するように証跡を作成しておく必要がある。


[チュートリアル: インプットトランスフォーマーを使用してイベントターゲットに渡される内容をカスタマイズする](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/eventbridge-input-transformer-tutorial.html)

インプットトランスフォーマーを使用することで、イベントから取得されるテキストをターゲットに渡す前にカスタマイズ可能。


[自動Amazon EBSスナップショットのスケジュール](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/take-scheduled-snapshot.html)

cron 式でルールをトリガーし、ターゲットを [EC2 CreateSnapshot API 呼び出し] にすることが可能。


[AWS Lambda 関数をスケジュールする](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/run-lambda-schedule.html)

cron 式でルールをトリガーし、ターゲットを Lambda 関数にすることが可能。


[Amazon Kinesisストリームにイベントを中継する](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/relay-events-kinesis-stream.html)

ターゲットを Kinesis ストリームにすることが可能。


[ファイルが Amazon S3 バケットにアップロードされたときに Amazon ECS タスクを実行する](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/eventbridge-tutorial-ecs.html)

* イベント
  * サービス名: S3
  * イベントタイプ: オブジェクトレベルのオペレーション
  * 特定のオペレーション: Put Object
  * 特定のバケット: バケット名
* ターゲット: ECS タスク
  * クラスター、タスク定義などの情報


[自動化されたビルドのスケジュール](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/eventbridge-tutorial-codebuild.html)

ターゲットを CodeBuild にすることが可能。


[Amazon EC2インスタンスの状態変更のログ記録](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/eventbridge-tutorial-cloudwatch-logs.html)

ターゲットを CLoudWatch Logs にすることが可能。


## ルールのスケジュール式

[ルールのスケジュール式](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/scheduled-events.html)

#### cron 式

フィールドは左から順に 分、時間、日、月、曜日、年

例
```shell
# 毎日午後 12:00 UTC
cron(0 12 * * ? *)
```

#### rate 式

指定した時間間隔で実行

例
```shell
# 1 分間隔
rate(1 minute)
# 1 時間間隔
rate(1 hour)
```


## イベントとイベントパターン

[EventBridge でのイベントとイベントパターン](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/eventbridge-and-event-patterns.html)


[AWS イベント](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/aws-events.html)

各 AWS サービスの状態が変わった時にイベントが発生するほか、CloudTrail で特定の API が記録された場合にイベントを発生させることができる。
イベントは JSON 形式になっている。

```json
{
  "version": "0",
  "id": "6a7e8feb-b491-4cf7-a9f1-bf3703467718",
  "detail-type": "EC2 Instance State-change Notification",
  "source": "aws.ec2",
  "account": "111122223333",
  "time": "2017-12-22T18:43:48Z",
  "region": "us-west-1",
  "resources": [
    "arn:aws:ec2:us-west-1:123456789012:instance/i-1234567890abcdef0"
  ],
  "detail": {
    "instance-id": " i-1234567890abcdef0",
    "state": "terminated"
  }
}
```
* source: イベントを発生させたサービスを識別
* detail-type: どの種類のイベント化を識別するための情報
* detail: イベントの内容


[イベントパターン](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/filtering-examples-structure.html)

どのイベントを抽出するかのフィルタリングの仕方を定義したもの。

例としては次のような内容。
```json
{
  "source": [ "aws.ec2" ],
  "detail-type": [ "EC2 Instance State-change Notification" ],
  "detail": {
    "state": [ "running" ]
  }
}
```



## インプットトランスフォーマー

[ターゲット入力の変換](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/transform-input.html)

インプットトランスフォーマーにより、イベントの内容を加工することが可能。



## サポートされている AWS サービスからのイベント

[サポートされている AWS サービスからの EventBridge イベントの例](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/event-types.html#ecs-event-types)

各 AWS サービスの EventBridge のイベント例がまとめられている。



## PutEvents

* [PutEvents](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/add-events-putevents.html)
* [API Reference - PutEvents](https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_PutEvents.html)

PutEvents の API によりイベントを送ることができる。



## インターフェイス VPC エンドポイント

[EventBridge とインターフェイス VPC エンドポイントの使用](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/eventbridge-and-interface-VPC.html)

EventBridge 用のプライベートインタフェース VPC エンドポイントが用意されている。
作成しておくことで、EventBridge が他の AWS サービスの API を実行する際に使用するエンドポイントとしてこちらが使われる。



## CloudWatch メトリクス

[CloudWatch メトリクスの使用状況のモニタリング](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/eventbridge-monitoring-cloudwatch-metrics.html)

CloudWatch メトリクスの一覧。


## ターゲット

[Amazon EventBridgeターゲット](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/eventbridge-targets.html)

ターゲットとして、Lambda 関数、SNS トピック、SQS キュー、StepFunctions、ECS タスクなど、多くのサービスを選択可能。


## アイデンティティとアクセスの管理

[EventBridge でのアイデンティティベースのポリシー (IAM ポリシー) の使用](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/iam-identity-based-access-control-eventbridge.html)

EventBridge がターゲットにアクセスするためには、ターゲットにアクセスするための IAM ロールを指定し必要な検眼を付与する必要がある。(※ Lambda, SNS, SQS, CloudWatch Logs 以外の場合)


[EventBridge のリソースベースのポリシーを使用する](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/resource-based-policies-eventbridge.html)

Lambda, SNS, SQS, CloudWatch Logs はリソースベースのポリシーで許可する必要がある。


[EventBridge の権限リファレンス](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/permissions-reference-eventbridge.html)

API ごとの説明が一覧になっている。


[詳細に設定されたアクセスコントロールのための IAM ポリシー条件の使用](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/policy-keys-eventbridge.html#events-limit-access-control)

「events:PutRule」「events:PutTargets」アクションのポリシー設定例が載っている。



## トラブルシューティング

[Amazon EventBridge のトラブルシューティング](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/eventbridge-troubleshooting.html#delayed-event-delivery)

#### ルールはトリガーされたが、Lambda 関数が呼び出されなかった

Lambda 関数に対する適切なアクセス権限が設定されていることを確認する。

#### 予定した時刻にトリガーされなかった

時刻は UTC による指定となっている。JST ではないので注意。

#### ターゲットへのイベントの配信で遅延が発生

最大 24 時間に渡ってイベントの再配信をしようとする。
24 時間経過するとそれ以上はスケジュールされず、FailedInvocations メトリクスが CloudWatch で発行される。



# 参考

* Document
  * [Amazon EventBridge とは](https://docs.aws.amazon.com/ja_jp/eventbridge/latest/userguide/what-is-amazon-eventbridge.html)
  * [Amazon EventBridge API Reference](https://docs.aws.amazon.com/eventbridge/latest/APIReference/Welcome.html)
* サービス紹介ページ
  * [Amazon EventBridge](https://aws.amazon.com/jp/eventbridge/)
  * [よくある質問](https://aws.amazon.com/jp/eventbridge/faqs/)
* Black Belt
  * [20200122 AWS Black Belt Online Seminar Amazon EventBridge](https://www.slideshare.net/AmazonWebServicesJapan/20200122-aws-black-belt-online-seminar-amazon-eventbridge)

