
[AWS Copilot コマンドラインインターフェイスの使用](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/AWS_Copilot.html)

* Copilot CLI を通して各種操作を行う


[AWSコパイロットを使用して、Amazon ECS の開始方法](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/getting-started-aws-copilot-cli.html)

サンプルをデプロイする流れ
```
git clone https://github.com/aws-samples/amazon-ecs-cli-sample-app.git demo-app && \ 
cd demo-app &&                               \
copilot init --app demo                      \
  --name api                                 \
  --type 'Load Balanced Web Service'         \
  --dockerfile './Dockerfile'                \
  --port 80                                  \
  --deploy
```

Dockerfile は以下の内容

```Dockerfile
FROM nginx
COPY index.html /usr/share/nginx/html
```

#### コマンド

```sh
// アプリケーションを一覧を表示
copilot app ls

// アプリケーション内の環境およびサービスに関する情報を表示
copilot app show

// 環境に関する情報を表示
copilot env ls

// エンドポイント、キャパシティー、関連リソースなど、サービスに関する情報を表示
copilot svc show

// アプリケーション内のすべてのサービスのリスト
copilot svc ls

// デプロイされたサービスのログを表示
copilot svc logs

// サービスのステータスを表示
copilot svc status
```



[FRONTEND RAILS APP](https://ecsworkshop.com/microservices/frontend/#create-a-ci-cd-pipeline)

CI/CD パイプラインと git ワークフローを完全に自動化する方法のワークショップ

以下のリポジトリを使用。

https://github.com/aws-containers/ecsdemo-frontend

