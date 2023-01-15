
[AWS Lambda Function URLs の提供開始: 単一機能のマイクロサービス向けの組み込み HTTPS エンドポイント](https://aws.amazon.com/jp/blogs/news/announcing-aws-lambda-function-urls-built-in-https-endpoints-for-single-function-microservices/)

* 新しい関数 URL を作成し、任意の関数にマッピング
* 重み付けされたトラフィックシフトと安全なデプロイを実装することもできます
* [Advanced Settings] (高度な設定) の [Enable function URL] (関数 URL を有効化) にチェックを入れます。
* Auth タイプとして [AWS_IAM] または [NONE] を選択
* AuthType [None] を選択します。これは、Lambda が関数を呼び出す前に AWS IAM Sigv4 署名をチェックしないことを意味します。
* AuthType [None] を使用する場合でも、関数のリソースベースのポリシーではパブリックアクセスを明示的に許可する必要がある
* ワンクリックで CORS を有効にすることもできます。

[コンテナ利用者に捧げる AWS Lambda の新しい開発方式 !](https://aws.amazon.com/jp/builders-flash/202103/new-lambda-container-development/?awsf.filter-name=*all)

[[2022年最新版]Lambdaの裏側教えます！！A closer look at AWS Lambda (SVS404-R) #reinvent](https://dev.classmethod.jp/articles/reinvent2020-session-svs404/)

