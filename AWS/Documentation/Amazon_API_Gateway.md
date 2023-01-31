# Document

[Amazon API Gateway とは何ですか?](https://docs.aws.amazon.com/ja_jp/apigateway/latest/developerguide/welcome.html)


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


[サーバーレスのウェブアプリケーションの構築](https://webapp.serverlessworkshops.io/)



# 参考

* Document
  * [Amazon API Gateway とは何ですか?](https://docs.aws.amazon.com/ja_jp/apigateway/latest/developerguide/welcome.html)
* Black Belt
  * [Amazon API Gateway](https://pages.awscloud.com/rs/112-TZM-766/images/20190514_AWS-Blackbelt_APIGateway_rev.pdf)



