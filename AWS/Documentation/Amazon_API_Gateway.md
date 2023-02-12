# Document

[Amazon API Gateway とは何ですか?](https://docs.aws.amazon.com/ja_jp/apigateway/latest/developerguide/welcome.html)


[API Gateway の開始方法](https://docs.aws.amazon.com/ja_jp/apigateway/latest/developerguide/getting-started.html)

1. Lambda 関数の作成
2. HTTP API の作成。統合で Lambda 関数を追加
3. API の呼び出し URL にアクセスする


## チュートリアル、ワークショップ

[Amazon API Gateway のチュートリアルとワークショップ](https://docs.aws.amazon.com/ja_jp/apigateway/latest/developerguide/api-gateway-tutorials.html)


[チュートリアル: Lambda プロキシ統合を使用した Hello World REST API の構築](https://docs.aws.amazon.com/ja_jp/apigateway/latest/developerguide/api-gateway-create-api-as-simple-proxy-for-lambda.html)

* リソースパスが URL パスに対応する
* リソースパスごとに HTTP メソッドを設定する。ANY は各メソッドに対応している
* デプロイする際にステージを指定する
* 次のような URL で HTTP GET でアクセスできる
  * https://r275xc9bmd.execute-api.us-east-1.amazonaws.com/test/helloworld?name=John&city=Seattle


[チュートリアル: Lambda と DynamoDB を使用した CRUD API の構築](https://docs.aws.amazon.com/ja_jp/apigateway/latest/developerguide/http-api-dynamo-db.html)

* 以下のルートを実装している API
  * GET /items/{id}
  * GET /items
  * PUT /items
  * DELETE /items/{id}


## REST API の操作

[Amazon API Gateway での REST API の作成](https://docs.aws.amazon.com/ja_jp/apigateway/latest/developerguide/how-to-create-api.html)

リソース間の関係は CloudFormation テンプレートで書くと分かりやすいかも。
* [AWS::ApiGateway::RestApi](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html)
* [AWS::ApiGateway::Resource](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html)
* [AWS::ApiGateway::Method](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html)
```yaml
ProxyResource:
  Type: 'AWS::ApiGateway::Resource'
  Properties:
    RestApiId: !Ref LambdaSimpleProxy
    ParentId: !GetAtt 
      - LambdaSimpleProxy
      - RootResourceId
    PathPart: '{proxy+}'
ProxyResourceANY:
  Type: 'AWS::ApiGateway::Method'
  Properties:
    RestApiId: !Ref LambdaSimpleProxy
    ResourceId: !Ref ProxyResource
    HttpMethod: ANY
    AuthorizationType: NONE
    Integration:
      Type: AWS_PROXY
      IntegrationHttpMethod: POST
      Uri: !Sub >-
        arn:aws:apigateway:${AWS::Region}:lambda:path/2015-03-31/functions/${LambdaForSimpleProxy.Arn}/invocations
```


[API Gateway API に対してセットアップするエンドポイントタイプを選択する](https://docs.aws.amazon.com/ja_jp/apigateway/latest/developerguide/api-gateway-api-endpoint-types.html)

以下の種類がある
* エッジ最適化 API エンドポイント
* リージョン API エンドポイント
* プライベート API エンドポイント


[API Gateway での REST API へのアクセスの制御と管理](https://docs.aws.amazon.com/ja_jp/apigateway/latest/developerguide/apigateway-control-access-to-api.html)

* リソースポリシーを設定可能
* 認証は Cognito を使用可能


[API Gateway で Lambda 統合を設定する](https://docs.aws.amazon.com/ja_jp/apigateway/latest/developerguide/set-up-lambda-integrations.html)


[Amazon API Gateway での REST API のデプロイ](https://docs.aws.amazon.com/ja_jp/apigateway/latest/developerguide/how-to-deploy-api.html)

* ステージに対してデプロイする。ステージ名は dev 、prod、beta、v2 のように任意のものを指定できる
* デフォルトドメインの場合、発行される URL は https://{restapi-id}.execute-api.{region}.amazonaws.com/{stageName} の形式


[API Gateway の Canary リリースデプロイの設定](https://docs.aws.amazon.com/ja_jp/apigateway/latest/developerguide/canary-release.html)

* ベースバージョンをデプロイしたまま、非本稼働バージョンを Canaly リリースできる
* テスト完了したら非本稼働バージョンを本稼働バージョンに昇格する



# Workshop

[サーバーレスのウェブアプリケーションの構築](https://webapp.serverlessworkshops.io/)



# BlackBelt

[Amazon API Gateway](https://pages.awscloud.com/rs/112-TZM-766/images/20190514_AWS-Blackbelt_APIGateway_rev.pdf)

* API 提供時の課題。インフラの管理、API の管理、認証・認可、流量制限など
* API > リソース > HTTP メソッドの関係。リソースは URL のパスに相当
* ステージにデプロイすることでクライアントからアクセス可能となる。ステージ名は URL の一部になっている
* API は Swagger ファイルのインポートでも作成可能
* メソッドはメソッドリクエスト → 統合リクエスト → 統合バックエンド → 統合レスポンス → メソッドレスポンスのフローとなる
  * メソッドリクエスト: クエリパラメータ、必須 HTTP ヘッダなどを設定
  * 統合リクエスト: ルーティング先バックエンドの設定、リクエストの変換など
  * 統合レスポンス: レスポンス内容の変換など
  * メソッドレスポンス: どの HTTP ステータスコード、HTTP ヘッダを返却するかなど
* 認証・認可: IAM アクセス権限、Lambda オーサライザー、Cognito オーサライザーがある。
  * IAM アクセス権限: Sig v4 の署名送信を要求する方式
  * Lambda オーサライザー: Lambda 認証関数がポリシー & プリンシパル ID を返却する方式
  * Cognito オーサライザー: Cognito ユーザープールで認証を行い、取得したトークンを HTTP ヘッダに指定することを要求する方式
* 統合タイプ
  * Lambda 関数、HTTP、Mock、AWS サービス、VPC リンクがある
  * プロキシ統合: バックエンドから渡された HTTP 応答データをそのまま返却
* その他機能
  * スロットリング
  * カスタムドメイン
  * ステージごとにキャッシュを設定可能
  * リソースポリシー
  * カナリアリリース
  * ステージに WAF の WebACL を設定
  * X-Ray 連携
  * クライアント証明書の作成



# 参考

* Document
  * [Amazon API Gateway とは何ですか?](https://docs.aws.amazon.com/ja_jp/apigateway/latest/developerguide/welcome.html)
* Black Belt
  * [Amazon API Gateway](https://pages.awscloud.com/rs/112-TZM-766/images/20190514_AWS-Blackbelt_APIGateway_rev.pdf)



