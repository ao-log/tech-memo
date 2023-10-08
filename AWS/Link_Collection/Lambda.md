
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


[コンテナ開発者向けの AWS Lambda](https://aws.amazon.com/jp/blogs/news/aws-lambda-for-the-containers-developer/)

* コンテナイメージを使用して Lambda を設定している
* Lambda なのでイベント駆動
* コンテナの初期化がされておりアイドル状態の場合は当該コンテナにイベントを割り当てる
* コンテナの初期化がされておらずアイドル状態のコンテナが存在しない場合は、新規にコンテナを起動する


[AWS Lambda 関数の実行の仕組みを知ろう !](https://aws.amazon.com/jp/builders-flash/202308/learn-lambda-function-execution/?awsf.filter-name=*all)

* イベントソースによって呼び出され方が異なる。同期呼び出し、非同期呼び出し、ポーリング
* `lambda_handler` 関数が最初に実行される。イベントオブジェクト、コンテキストオブジェクトを受け取る。イベントオブジェクトはイベントソースからのデータ、コンテキストオブジェクトは Lambda ランタイムからの情報
* ハンドラーの実行時間は最大 15 分まで
* 初回起動時はコールドスタート。実行環境の作成、初期化なども行われる
+ しばらく実行環境は残っている。この場合はウォームスタート
* 実行環境の同時実行数は Service Quotas の Concurrent executions の設定値まで
* Lambda 関数ごとに「予約済同時実行数」を設定できる。設定値までの実行環境が作成可能
* バージョンを発行することができる。発行されたバージョンのコードは変更不可。LATEST, DEFAULT のエイリアスを設定可能。トラフィックの重みづけを設定可能
* プロビジョニングされた同時実行数により、ウォームスタートされた状態で指定値分の実行環境を使用できる。コストには注意
* Lambda SnapStart によりコールドスタートの時間短縮が期待できる
