
# StepFunctions

ステートマシンを作り実行できるサービス。



# 導入

[AWS Step Functions の開始方法](https://docs.aws.amazon.com/ja_jp/step-functions/latest/dg/getting-started.html)

ASL(Amazon State Language) で記述する。
```json
{
  "Comment": "A Hello World example of the Amazon States Language using Pass states",
  "StartAt": "Hello",
  "States": {
    "Hello": {
      "Type": "Pass",
      "Result": "Hello",
      "Next": "World"
    },
    "World": {
      "Type": "Pass",
      "Result": "World",
      "End": true
    }
  }
}
```

様々なチュートリアルが用意されている。

[チュートリアル](https://docs.aws.amazon.com/ja_jp/step-functions/latest/dg/tutorials.html)

サンプルプロジェクトも用意されている。

[Sample Projects for Step Functions](https://docs.aws.amazon.com/ja_jp/step-functions/latest/dg/create-sample-projects.html)

#### AWS Batch の例

[バッチジョブの管理 (AWS Batch、Amazon SNS)](https://docs.aws.amazon.com/ja_jp/step-functions/latest/dg/batch-job-notification.html)

上記ドキュメントは、AWS Batch ジョブを送信後、ジョブの結果 (成功または失敗) に基づき Amazon SNS 通知を送信する方法について例。

ドキュメントからの引用となるが ASL は以下の例のとおりとなる。
```json
{
  "Comment": "An example of the Amazon States Language for notification on an AWS Batch job completion",
  "StartAt": "Submit Batch Job",
  "TimeoutSeconds": 3600,
  "States": {
    "Submit Batch Job": {
      "Type": "Task",
      "Resource": "arn:aws:states:::batch:submitJob.sync",
      "Parameters": {
        "JobName": "BatchJobNotification",
        "JobQueue": "arn:aws:batch:us-east-1:123456789012:job-queue/BatchJobQueue-7049d367474b4dd",
        "JobDefinition": "arn:aws:batch:us-east-1:123456789012:job-definition/BatchJobDefinition-74d55ec34c4643c:1"
      },
      "Next": "Notify Success",
      "Catch": [
          {
            "ErrorEquals": [ "States.ALL" ],
            "Next": "Notify Failure"
          }
      ]
    },
    "Notify Success": {
      "Type": "Task",
      "Resource": "arn:aws:states:::sns:publish",
      "Parameters": {
        "Message": "Batch job submitted through Step Functions succeeded",
        "TopicArn": "arn:aws:sns:us-east-1:123456789012:batchjobnotificatiointemplate-SNSTopic-1J757CVBQ2KHM"
      },
      "End": true
    },
    "Notify Failure": {
      "Type": "Task",
      "Resource": "arn:aws:states:::sns:publish",
      "Parameters": {
        "Message": "Batch job submitted through Step Functions failed",
        "TopicArn": "arn:aws:sns:us-east-1:123456789012:batchjobnotificatiointemplate-SNSTopic-1J757CVBQ2KHM"
      },
      "End": true
    }
  }
}
```

* "StartAt" が "Submit Batch Job" なので、"Submit Batch Job" からはじめる。
* "Submit Batch Job" のジョブを実行。成功時は "Notify Success" へ。エラーをキャッチしたときは "Notify Failure" へ。



# 仕組み

[標準ワークフローとExpress ワークフロー](https://docs.aws.amazon.com/ja_jp/step-functions/latest/dg/concepts-standard-vs-express.html)

・標準ワークフロー
最大期間は 1 年。
状態を内部的に持つことができるが、同時に 1 つだけしか実行できない。

・Express ワークフロー
最大期間は 5 分。
状態を持たない。最低 1 回の実行がされる。

[状態](https://docs.aws.amazon.com/ja_jp/step-functions/latest/dg/concepts-states.html)

* 個別の状態では、入力に基づいて決定を行い、アクションを実行して、出力を他の状態に渡すことができる。
* ワークフローは ASL(Amazon State Language) で記述する。
* 各状態には、必ずその状態のタイプを示す Type フィールドがある。
* 各状態 (Succeed または Fail 状態を除く) には Next フィールドが必要。代わりに End フィールドを指定して終了状態にする。

[Amazon ステートメント言語](https://docs.aws.amazon.com/ja_jp/step-functions/latest/dg/concepts-amazon-states-language.html)

Amazon ステートメント言語 仕様の例が載っている。

[Task](https://docs.aws.amazon.com/ja_jp/step-functions/latest/dg/amazon-states-language-task-state.html)

* Resource (必須): URI (特に、実行する特定のタスクを一意に識別する ARN)。
* Parameters (オプション): 接続されたリソースの API アクションに情報を渡すのに使用。
* Retry (オプション): 状態でランタイムエラーが発生した場合の再試行ポリシーを定義。
* Catch (オプション): 状態にランタイムエラーが発生し、その再試行ポリシーが使い果たされたか定義されていない場合に実行される。



[Step Functions の入出力処理](https://docs.aws.amazon.com/ja_jp/step-functions/latest/dg/concepts-input-output-filtering.html)

Step Functions の実行は JSON テキストを入力として受け取り、その入力をワークフローの最初の状態に渡す。
各状態ごとに JSON を入力として受け取る。
そして通常 JSON として次の状態に渡す。

[エラー処理](https://docs.aws.amazon.com/ja_jp/step-functions/latest/dg/concepts-error-handling.html)

デフォルトでは、状態でエラーが報告されると、AWS Step Functions の実行全体が失敗する。
リトライを設定したり、エラーを Catch して処理を継続することも可能。

[サービスの AWS Step Functions との統合](https://docs.aws.amazon.com/ja_jp/step-functions/latest/dg/concepts-service-integrations.html)

Step Functions から呼び出せるサービス。
Lambda, Batch, ECS タスクなど。

[呼び出し AWS Step Functions 他のサービスから](https://docs.aws.amazon.com/ja_jp/step-functions/latest/dg/concepts-invoke-sfn.html)

呼び出し元。
Lambda, API Gateway, EventBridge など。



# ベストプラクティス

[Step Functions のベストプラクティス](https://docs.aws.amazon.com/ja_jp/step-functions/latest/dg/sfn-best-practices.html)

* タイムアウトを設定する。理由としては、正常にレスポンスが帰ってこない場合に処理が止まってしまうのを防ぐため。
* 入力が大きいサイズの場合は、S3 上のオブジェクトを渡すことも可能。



# 参考

* [20190522 AWS Black Belt Online Seminar AWS Step Functions](https://www.slideshare.net/AmazonWebServicesJapan/20190522-aws-black-belt-online-seminar-aws-step-functions)
* [StepFunctions ドキュメント](https://docs.aws.amazon.com/ja_jp/step-functions/latest/dg/welcome.html)

