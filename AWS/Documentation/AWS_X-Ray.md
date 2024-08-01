# Document

[AWS X-Ray](https://docs.aws.amazon.com/ja_jp/xray/latest/devguide/aws-xray.html)

* 用語
  * セグメント: 開始・終了時刻、HTTP リクエスト、レスポンスなどの情報
  * サブセグメント: 外部 HTTP API, SQL 呼び出しなどの追加情報。セグメントを送信しないサービスについても、ダウンストリームノードとして描画できる
  * サービスグラフ: 同じトレース ID を結合した JSON ドキュメント。サービスグラフをもとにサービスマップを生成する
  * トレース: トレース ID を使用して追跡する
  * サンプリング: 全てを収集せず、定められたルールに従って一部のみを収集する
  * トレースヘッダー: X-Amzn-Trace-Id にトレース ID が格納される
  * フィルタ式: 特定のトレースを検索できる
  * グループ: フィルタ式をもとに CloudWatch メトリクスなどを生成する
  * アノテーション: フィルタ式用
  * メタデータ: 情報追加用
  * エラー: Error(4xx), Fault(5xx), Throttle


[インターフェイスを選択する](https://docs.aws.amazon.com/ja_jp/xray/latest/devguide/aws-xray-interface.html)

[AWS Management Console を使用する](https://docs.aws.amazon.com/ja_jp/xray/latest/devguide/aws-xray-interface-console.html)

* マネジメントコンソールからの一通りの操作についての説明がある


[SDK を使用する](https://docs.aws.amazon.com/ja_jp/xray/latest/devguide/aws-xray-interface-sdk.html)

* 2 種類ある
  * ADOT SDK
  * X-Ray SDK



## X-Ray デーモン

[AWS X-Ray デーモン](https://docs.aws.amazon.com/ja_jp/xray/latest/devguide/xray-daemon.html)

* クライアントからのセグメントデータを受信し、X-Ray の API エンドポイントに送信する
* 2000/udp で LISTEN
* クライアントからは `AWS_XRAY_DAEMON_ADDRESS` で宛先を指定可能



## アプリケーションを計測する

[アプリケーションを計測する](https://docs.aws.amazon.com/ja_jp/xray/latest/devguide/xray-instrumenting-your-app.html)


### Python

[Python を使用してアプリケーションを計測する](https://docs.aws.amazon.com/ja_jp/xray/latest/devguide/xray-python.html)


[AWS X-Ray 用の SDK Python](https://docs.aws.amazon.com/ja_jp/xray/latest/devguide/xray-sdk-python.html)


[X-Ray SDK for Python ミドルウェアを使用して受信リクエストをトレースします](https://docs.aws.amazon.com/ja_jp/xray/latest/devguide/xray-sdk-python-middleware.html)

* 次のフレームワークがサポートされている
  * Django
  * Flask
  * Bottle
* 他のフレームワークの場合は手動でセグメントを作成できる
```python
from aws_xray_sdk.core import xray_recorder

# Start a segment
segment = xray_recorder.begin_segment('segment_name')
# Start a subsegment
subsegment = xray_recorder.begin_subsegment('subsegment_name')

# Add metadata and annotations
segment.put_metadata('key', dict, 'namespace')
subsegment.put_annotation('key', 'value')

# Close the subsegment and segment
xray_recorder.end_subsegment()
xray_recorder.end_segment()
```


[ダウンストリームコールを実装するためのライブラリへのパッチ適用](https://docs.aws.amazon.com/ja_jp/xray/latest/devguide/xray-sdk-python-patching.html)

* ダウンストリーム呼び出しを測定するには、パッチを適用する。`requests` や各 DB ライブラリが対応している
* パッチを適用するとサブセグメントが作成され、リクエスト、レスポンスの内容が記録される
```python
import boto3
import botocore
import requests
import sqlite3

from aws_xray_sdk.core import xray_recorder
from aws_xray_sdk.core import patch_all

patch_all()
```


[X-Ray AWS SDK for Python を使用した SDK 呼び出しのトレース](https://docs.aws.amazon.com/ja_jp/xray/latest/devguide/xray-sdk-python-awssdkclients.html)

* botocore をパッチすることで、DynamoDB, SQS, S3 への呼び出しを追跡できる


[X-Ray SDK for Python を使用してダウンストリーム HTTP ウェブサービスの呼び出しをトレースする](https://docs.aws.amazon.com/ja_jp/xray/latest/devguide/xray-sdk-python-httpclients.html)

* 外部への HTTP リクエストもダウンストリーム HTTP 呼び出しとして追跡できる。`requests` などをパッチするだけでよい


[X-Ray SDK for Python を使用したカスタムサブセグメントの生成](https://docs.aws.amazon.com/ja_jp/xray/latest/devguide/xray-sdk-python-subsegments.html)

* 以下の例のようにサブセグメントを追加できる
```python
from aws_xray_sdk.core import xray_recorder

subsegment = xray_recorder.begin_subsegment('annotations')
subsegment.put_annotation('id', 12345)
xray_recorder.end_subsegment()
```
* サブセグメントにも自動的に開始、終了時刻が記録される


[X-Ray SDK for Python を使用してセグメントに注釈とメタデータを追加する](https://docs.aws.amazon.com/ja_jp/xray/latest/devguide/xray-sdk-python-segment.html)
* 現在のセグメントは以下のように参照できる
```python
from aws_xray_sdk.core import xray_recorder
...
document = xray_recorder.current_segment()
```
* 以下のようにアノテーションを追加できる
```python
document.put_annotation("mykey", "my value");
```
* 以下のようにメタデータを追加できる
```python
document.put_metadata("my key", "my value", "my namespace");
```


[サーバーレス環境にデプロイされたウェブフレームワークの計測](https://docs.aws.amazon.com/ja_jp/xray/latest/devguide/xray-sdk-python-serverless.html)

* API Gateway, Lambda の構成におけるサンプルのドキュメント



## 他の AWS サービスとの統合

[他の AWS のサービス と AWS X-Ray との統合](https://docs.aws.amazon.com/ja_jp/xray/latest/devguide/xray-services.html)

* X-Ray との統合は以下のパターンがある
  * アクティブ計測 – 受信リクエストをサンプリングして計測
  * パッシブ計測 – 別のサービスで既にサンプリングされているリクエストを計測
  * リクエストのトレース – すべての受信リクエストにトレースヘッダーを追加してダウンストリームに伝達
  * ツール – X-Ray デーモンを実行して X-Ray SDK からセグメントを受信します。


[AWS Lambda および AWS X-Ray](https://docs.aws.amazon.com/ja_jp/xray/latest/devguide/xray-services-lambda.html)

* 追加の設定なしでトレースされる
* X-Ray SDK を使用する場合はバンドルする必要がある


[Amazon SNS と AWS X-Ray](https://docs.aws.amazon.com/ja_jp/xray/latest/devguide/xray-services-sns.html)

* アクティブトレースを有効にすることで SNS クライアントから発行されたメッセージを X-Ray に送信できる
  * リソースベースのポリシーにて X-Ray のアクションを許可する必要がある


[Amazon SQS と AWS X-Ray](https://docs.aws.amazon.com/ja_jp/xray/latest/devguide/xray-services-sqs.html)

* トレースヘッダーを送信し、コンシューマーに伝達する
* memo
  * Lambda 関数から SQS キューに送信する場合、アクティブトレースを有効にしているとトレース ID が SQS でも伝達されるため、SQS キューのコンシューマについてもトレース可能。ただし、SQS に対するトレースは X-Ray SDK を使用しないとされないため、Lambda Layer に X-Ray SDK を入れておき、Lambda 関数内でインポート、パッチして使用すること



# BlackBelt

[AWS X-Ray](https://pages.awscloud.com/rs/112-TZM-766/images/AWS-Black-Belt_2023_AWS-X-Ray_0228_v1.pdf)

* Observability
  * 三本の柱: Metrics, Log, Trace
* 用語
  * トレース: 一つのリクエストの全体のセグメント
  * セグメント: 各コンポーネントで採取したデータ
  * サブセグメント: 追加の詳細情報
* トレース
  * トレース ID: リクエストごとに一意。最初にセグメントを作成するタイミングで生成
  * セグメント ID: 親セグメント ID　を確認することでセグメントのつながり、順序関係がわかる
  * トレースヘッダーにより伝搬
* 機能
  * サンプリング
  * 注釈: 検索しやすくする目的
  * メタデータ: デバッグ、分析用途
* 分析
  * サービスマップ: 各ノード間の関係、レイテンシ、トレース数などを描画。ノードの色によりエラー状況を確認可能(4xx, 5xx, スロットリング)
    * 各ノードの詳細情報表示: レイテンシ、レスポンスステータス(4xx, 5xx, スロットリング)
  * トレースリスト
    * 上部に URL ごとのトレース分布、レスポンス平均時間が表示される
    * 下部にトレースリストが表示される
      * 各トレースの詳細情報表示: セグメントごとの所要時間
        * 各セグメントの詳細情報表示: セグメントの時間情報、エラー情報、注釈、メタデータなど
  * アナリティクス
    * 応答時間の分布
    * 時系列のトレース数の分布
    * メトリクステーブル
  * インサイト
    * 有効化することで異常を自動検出
    * 概要タブ: 問題の概要、根本原因、影響の確認が可能
    * 調査タブ: より詳細な影響タイムライン、影響分析を確認可能
  * フィルタ式
    * トレースリスト、アナリティクスで使用可能
    * 総所要時間、ステータスコードなどでフィルタ可能
  * グループ
    * サービスマップ、トレースリスト、アナリティクスで使用可能
    * フィルタ式で定義されるトレースのコレクション
    * CloudWatch にトレース数のメトリクスを作成
* データ収集の仕組み
  * アプリケーションから X-Ray デーモンを介して X-Ray API エンドポイントに送信
* X-Ray SDK
* Auto-Instrumentation Agent
  * Java Spring Boot に対応
* 連携している AWS サービス
  * Lambda: HTTP 呼び出しの追加などには X-Ray SDK が必要
  * API Gateway: HTTP リクエストにトレースヘッダを追加
  * ECS: サイドカーで X-Ray デーモンを稼働させる必要がある
  * SQS: SQS、Lambda 構成のようなイベント駆動型の場合も分析が可能
* Amazon CloudWatch ServiceLens
* Amazon CloudWatch RUM
* Amazon CloudWatch Synthetics
* AWS Distro for OpenTelemetry (ADOT)
* Amazon Managed Grafana (AMG): データソースとして X-Ray をサポート


[OpenTelemetryとは](https://pages.awscloud.com/rs/112-TZM-766/images/AWS-Black-Belt_2024_AWS-Distro-for-OpenTelemetry-Part-1_0131_v1.pdf)

* OpenTelemetry Protocol (OTLP)
  * Receier, Procesor, Exporter の流れ



# 参考

* Document
  * [AWS X-Ray](https://docs.aws.amazon.com/ja_jp/xray/latest/devguide/aws-xray.html)
* Black Belt
  * [AWS X-Ray](https://pages.awscloud.com/rs/112-TZM-766/images/AWS-Black-Belt_2023_AWS-X-Ray_0228_v1.pdf)
  * [OpenTelemetryとは](https://pages.awscloud.com/rs/112-TZM-766/images/AWS-Black-Belt_2024_AWS-Distro-for-OpenTelemetry-Part-1_0131_v1.pdf)

