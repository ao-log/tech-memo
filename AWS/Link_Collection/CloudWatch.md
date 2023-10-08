
[[アップデート] Amazon CloudWatch エージェントが X-Ray と OpenTelemetry のトレース収集機能をサポートしました](https://dev.classmethod.jp/articles/cloudwatch-agent-opentelemetry-traces-x-ray/)

* [CloudWatch エージェント設定ファイルを手動で作成または編集する](https://docs.aws.amazon.com/ja_jp/AmazonCloudWatch/latest/monitoring/CloudWatch-Agent-Configuration-File-Details.html) のリファレンスの通り、traces にて指定する
* トレースの指定があると localhost 上で UDP の 2000 番ポートで LISTEN する
* `cat hoge2.txt > /dev/udp/127.0.0.1/2000` のようなコマンドで localhost の 2000/udp に対して送信できる


[【CloudFormation】SSM ステートマネージャーを使ってCloudWatch Agentのインストールとログ出力設定を自動化してみる](https://dev.classmethod.jp/articles/cloudformation-ssm-manager-automate-cloudwatch-agent-installation/)

* 自作の SSM ドキュメントで各種設定を実施
  * CloudWatch エージェントをインストール。SSM ドキュメント `AWS-ConfigureAWSPackage` にて action: install, name: AmazonCloudWatchAgent を設定
  * SSM ドキュメント `AmazonCloudWatch-ManageAgent` にて CloudWatch エージェントの設定を適用
    * SSM パラメータで CloudWatch エージェント用の設定を格納
* ステートマネージャで上記自作 SSM ドキュメントを指定。ターゲットは各 EC2 インスタンスとする

