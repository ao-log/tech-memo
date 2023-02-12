
# Document

[AWS Lambda とは](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/welcome.html)


[Lambda の機能](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/gettingstarted-features.html)

* 自動でスケーリングする。詳細は[Lambda 関数のスケーリング](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/invocation-scaling.html)を参照のこと
* 同時実行数の設定。事前にプロビジョニングしたい場合は、プロビジョンされた同時実行数を設定すること
* 専用のエンドポイントとなる URL の割り当て
* 非同期呼び出し。詳細は[非同期呼び出し](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/invocation-async.html)を参照のこと


## 許可

[Lambda でのアクセス許可](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/lambda-permissions.html)

* 「実行ロール」で、関数内で使用するアクションに対する許可を行う
* リソースベースのポリシーによりクロスアカウントアクセスなどを設定可能。AWS の他のサービスが呼び出し元の場合も許可設定が必要


## Lambda ランタイム

[Lambda ランタイム](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/lambda-runtimes.html)

* プログラミング言語の古いバージョンなどは廃止されていくので注意


## 関数のデプロイ

[Lambda 関数のデプロイ](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/lambda-deploy-functions.html)

次の方法がある
* ZIP ファイルのアップロード
* コンテナイメージ


## 関数の設定

[Lambda レイヤーの作成と共有](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/configuration-layers.html)


[Lambda 関数オプションの設定](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/configuration-function-common.html)

以下のような設定項目がある
* 関数のバージョン
* アクセス許可
* 環境変数
* VPC
* 同時実行数


[Lambda 関数のバージョン](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/configuration-versions.html)

* バージョンを公開すると、ほとんどの設定項目はロックされ設定変更できなくなる


[Lambda 関数のエイリアス](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/configuration-aliases.html)

* エイリアスではバージョンごとのトラフィック比率などを設定できる。2 つのバージョンまで指すことが可能


## AWS Lambda 関数の管理

[Lambda の予約済み同時実行数の管理](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/configuration-concurrency.html)

同時実行コントロールは以下の 2 種類がある
* 予約同時実行: 最大実行数まで実行できることを保証
* プロビジョニング済み同時実行: 事前にプロビジョニングし、即応答できるようにする


[VPC 内のリソースにアクセスするように Lambda 関数を設定する](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/configuration-vpc.html)

* 指定したサブネットに [Lambda Hyperplane ENI](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/foundation-networking.html#foundation-nw-eni) を作成する



[Lambda 関数のデータベースアクセスの設定](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/configuration-database.html)

* データベースへの接続数を抑制するために有用


## 関数を呼び出す

[Lambda 関数を呼び出す](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/lambda-invocation.html)

* 同期呼び出しでは、関数のレスポンスを返却する。非同期呼び出しではイベントをキューに入れて処理するものの、すぐにレスポンスを返却する


## Go の使用

[Go による Lambda 関数の構築](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/lambda-golang.html)


[Go の AWS Lambda 関数ハンドラー](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/golang-handler.html)

* 関数が呼び出されると、Lambda はハンドラーメソッドを実行する

```golang
package main

import (
        "fmt"
        "context"
        "github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
        Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
        return fmt.Sprintf("Hello %s!", name.Name ), nil
}

func main() {
        lambda.Start(HandleRequest)
}
```


## 他のサービスでの利用

[CloudWatch Logs で Lambda を使用する](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/services-cloudwatchlogs.html)

* CloudWatch Logs 側でサブスクリプションフィルターを作成し Lambda 関数を呼び出すように設定する
* Lambda 関数側ではリソースポリシーで logs を許可しておく
* [CloudWatch Logs サブスクリプションフィルターの使用](https://docs.aws.amazon.com/ja_jp/AmazonCloudWatch/latest/logs/SubscriptionFilters.html#LambdaFunctionExample) にも例がある


[Amazon DynamoDB で AWS Lambda を使用する](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/with-ddb.html)

* DynameDB Streams のレコードを処理できる


[Amazon S3 での AWS Lambda の使用](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/with-s3.html)

* オブジェクトの作成、削除などのイベントをトリガーに Lambda 関数を呼び出す
* 例えばイメージのサムネールの作成などのユースケース。加工後イメージは別バケットに出力すること。でないと、無限に Lambda 関数が実行されることになりうる



# 参考

* Document
  * [AWS Lambda とは](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/welcome.html)
* Black Belt
  * [Let’s Dive Deep into AWS Lambda Part 1, 2](https://pages.awscloud.com/rs/112-TZM-766/images/20190402_AWSBlackbelt_AWSLambda%20Part1&2.pdf)
  * [サーバーレス イベント駆動アーキテクチャ](https://pages.awscloud.com/rs/112-TZM-766/images/20200610_AWS_BlackBelt_Building_Event_driven_Architectures_on_AWS.pdf)
  * [Serverless モニタリング](https://pages.awscloud.com/rs/112-TZM-766/images/20190820_AWS-Blackbelt_Serverless_Monitoring.pdf)
  * [実践的サーバーレスセキュリティプラクティス](https://pages.awscloud.com/rs/112-TZM-766/images/20190813_AWS-BlackBelt_ServerlessSecurityPractice.pdf)
  * [形で考えるサーバーレス設計 サーバーレス ユースケースパターン解説](https://pages.awscloud.com/rs/112-TZM-766/images/20201118_AWS_BlackBet_Serverless_Usecase_Patterns.pdf)
  * [AWS Serverless Application Model (AWS SAM)](https://pages.awscloud.com/rs/112-TZM-766/images/20190814_AWS-Blackbelt_SAM_rev.pdf)

  