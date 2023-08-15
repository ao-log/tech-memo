# BlackBelt

[20200422 AWS Black Belt Online Seminar Amazon Elastic Container Service (Amazon ECS)](https://pages.awscloud.com/rs/112-TZM-766/images/20200422_BlackBelt_Amazon_ECS_Share.pdf)

* P11, 12: コンテナオーケストレータに対して API を実行することで各種操作を行う。
* P20-: EC2 起動タイプについて
* P25-: Fargate について
* P30-: タスク定義について
  * P43-: ネットワークモード
  * P48-: データボリューム
* P53-: コンテナの実行方法　タスク or サービス
  * P55: タスクの配置
  * P56: サービススケジューラ戦略
  * P57: キャパシティプロバイダー
  * P58: サービスの Auto Scaling


[20190731 Black Belt Online Seminar Amazon ECS Deep Dive](https://pages.awscloud.com/rs/112-TZM-766/images/20190731_AWS-BlackBelt_AmazonECS_DeepDive_Rev.pdf)

* P14: シークレットをコンテナ内のアプリに渡す場合の推奨の方法
  * Secrets Manager を利用し、タスク定義では environmentの valueFrom を使用して Secrets の ARN を記載
* P21: サービスディスカバリの方法
  * 要件次第である。
  * ELB を使用するのが一つの手。
  * 一方で、ECS Service Discovery は ECS が Route 53 に対して自動登録、削除をする仕組み。
* P32: サイドカーのような依存関係のあるコンテナの制御方法
  * タスク定義の dependsOn で依存関係を指定する
  * startTimeout: 依存関係の解決の再試行を止めるまでの時間 
  * stopTimeout: コンテナが SIGTERM で終了しなかった場合に SIGKILL されるまでの時間
* P41: スケジュールされたタスクのエラーハンドリング方法
  * 要件次第だが、StepFunctions により実行しエラーハンドリングする
  * 単に検知だけでよければ EventBridge を使用。
* P46: 自分たちでカスタマイズしたデプロイを行う方法
  * External Deployent Contoller を使用する。
* P64: EC2 起動タイプで awsvpc を使用している場合に起動できるタスク数が少ない
  * ENI Trunking 機能を有効化する
* P70: コンテナの起動まで時間がかかるためヘルスチェックが失敗する
  * ヘルスチェックの猶予時間でアプリケーションにあった時間を設定する
* P75: Fargate で起動するタスクのサイズを選ぶにあたってのリソース使用状況の把握方法
  * Container Insight を使用することでタスク、コンテナ単位のリソース使用状況を確認可能。


[20191127 AWS Black Belt Online Seminar Amazon CloudWatch Container Insights で始めるコンテナモニタリング入門](https://pages.awscloud.com/rs/112-TZM-766/images/20191127_AWS-BlackBelt_Container_Insights.pdf)

* P19: Container Insight はタスク、コンテナレベルでのモニタリングが可能
* P46: パフォーマンスログは CloudWatch Logs へ送られる。
* P51: ユースケース 1. ECS タスクに配置するコンテナリソースのサイジング
  * コンテナごとのリソース使用状況を確認し、適切なサイズに設定する。
* P57: 特定のタスクだけで発生している問題の調査
  * アプリケーションログの表示からログを確認する。


