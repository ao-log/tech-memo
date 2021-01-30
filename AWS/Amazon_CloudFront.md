
# Amazon CloudFront

[Amazon CloudFront とは何ですか?](https://docs.aws.amazon.com/ja_jp/AmazonCloudFront/latest/DeveloperGuide/Introduction.html)

コンテンツの配信を高速化するウェブサービス。
世界中のエッジロケーションを経由してコンテンツを配信できる。
エッジロケーションにキャッシュされるので可用性、信頼性の向上の効果も得られる。


[BlackBelt: 用語集のスライド](https://www.slideshare.net/AmazonWebServicesJapan/20201028-aws-black-belt-online-seminar-amazon-cloudfront-deep-dive/19)

* ビューワー: クライアント/Web ブラウザ
* ディストリビューション（SSL/TLS 証明書、WAF Web ACL、ログ用 S3 バケットなどを設定可能）
  * オリジン
    * カスタムオリジン: Web サーバー、ELB など。
    * S3
  * Behavior: キャッシュ動作設定（URL パスパターンごとに作成。キャッシュポリシー、オリジンリクエストポリシーなどを設定可能）


[ユースケース](https://docs.aws.amazon.com/ja_jp/AmazonCloudFront/latest/DeveloperGuide/IntroductionUseCases.html)

* 静的ウェブサイトの高速なコンテンツ配信
* オンデマンドビデオ、ライブストリーミングビデオの配信
* 特定のフィールドの暗号化
* Lambda@Edge によるカスタマイズ



## 開始方法

[簡単な CloudFront ディストリビューションの開始方法](https://docs.aws.amazon.com/ja_jp/AmazonCloudFront/latest/DeveloperGuide/GettingStarted.SimpleDistribution.html)

オリジンを S3 バケットとし、CloudFront 側で 24 時間キャッシュする構成のチュートリアル。



## ディストリビューションの使用

[ディストリビューションを作成または更新する場合に指定する値](https://docs.aws.amazon.com/ja_jp/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html)

このドキュメントで、ディストリビューション作成、更新時にどのようなパラメータを設定可能化を確認できる。
以下のような設定項目がある。

* オリジンの設定（ドメイン名、など）
* キャッシュ動作（パスパターン、TTL、など）
* ディストリビューション設定（証明書、CNAME、WAF、ログ設定、など）


[さまざまなオリジンの使用](https://docs.aws.amazon.com/ja_jp/AmazonCloudFront/latest/DeveloperGuide/DownloadDistS3AndCustomOrigins.html)

オリジンとしては、S3 バケット、ELB、EC2 インスタンスなどを設定可能。


[代替ドメイン名 (CNAME) を追加してカスタム URL を使用する](https://docs.aws.amazon.com/ja_jp/AmazonCloudFront/latest/DeveloperGuide/CNAMEs.html)

事前にドメイン名をドメインプロバイダーに登録しておく必要あり。
以下の項目を設定する。

* CNAME
* SSL 証明書

上記設定後、DNS 側を設定する。Route 53 の場合はエイリアスレコードで対応可能。



## ポリシーの使用

次の２つがある。

* キャッシュポリシー(キャッシュキーに含まれるもの)
* オリジンリクエストポリシー(キャッシュキーには含まれないが、オリジンへのリクエストに使用したいもの)

[キャッシュキーの管理](https://docs.aws.amazon.com/ja_jp/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html)

以下のような設定項目がある。

* 最小 TTL
* 最大 TTL
* デフォルト TTL
* キャッシュキーに含まれる値 (URL クエリ文字列、HTTP ヘッダー、Cookie) 

キャッシュ効率の低下を避けるため、必要最小限のヘッダーを指定することを推奨。


[オリジンリクエストの制御](https://docs.aws.amazon.com/ja_jp/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html)

キャッシュミスが発生した場合、オリジンリクエストが発生する。
オリジンに送信するリクエストに含まれる値 (URL クエリ文字列、HTTP ヘッダー、Cookie) を制御できる。



## コンテンツの追加、削除、または置き換え

[既存のコンテンツと CloudFront ディストリビューションを更新する](https://docs.aws.amazon.com/ja_jp/AmazonCloudFront/latest/DeveloperGuide/UpdatingExistingObjects.html)

既存のファイルを更新する場合、バージョン識別名をファイル名またはディレクトリ名に含めて、コンテンツを容易に制御できるようにすることを推奨。有効期限切れを待つ必要がなくなり、無効化の料金も不要になるため。


[コンテンツを削除して CloudFront が配信できないようにする](https://docs.aws.amazon.com/ja_jp/AmazonCloudFront/latest/DeveloperGuide/RemovingObjects.html)

通常、キャッシュヒットする間は、エッジにキャッシュとして残り続ける。ファイルを無効化するオペレーションが可能となっている。
対応方法は [ファイルの無効化](https://docs.aws.amazon.com/ja_jp/AmazonCloudFront/latest/DeveloperGuide/Invalidation.html) に記載されている。


[デフォルトのルートオブジェクトの指定](https://docs.aws.amazon.com/ja_jp/AmazonCloudFront/latest/DeveloperGuide/DefaultRootObject.html)

ルート URL にアクセスした場合、特定のオブジェクトを返すように設定可能。


[圧縮ファイルの供給](https://docs.aws.amazon.com/ja_jp/AmazonCloudFront/latest/DeveloperGuide/ServingCompressedFiles.html)

コンテンツが圧縮されるように設定可能。



## コンテンツへのセキュアなアクセスとアクセス制限の設定

### 代替ドメイン名と HTTPS の使用

[CloudFront で HTTPS リクエストを供給する方法の選択](https://docs.aws.amazon.com/ja_jp/AmazonCloudFront/latest/DeveloperGuide/cnames-https-dedicated-ip-or-sni.html)

次の 2 つの方法がある。

* Server Name Indication (SNI) を使用する – 推奨
* 各エッジロケーションの専用 IP アドレスを使用する

SNI はリクエストとドメインを関連付ける技術。ブラウザ側も対応している必要がある。
SNI を使用できない場合は、専用 IP アドレスを使用する設定とする。


### 署名付き URL と署名付き Cookie を使用したコンテンツの制限

[署名付き URL の使用](https://docs.aws.amazon.com/ja_jp/AmazonCloudFront/latest/DeveloperGuide/private-content-signed-urls.html)

署名付き URL には有効期限を設定可能。

CloudFront キーペアが必要。CloudFront 側ではパブリックキーにより署名を検証し、問題がある場合は拒否する。


### Amazon S3 コンテンツへのアクセスの制限

[オリジンアクセスアイデンティティを使用して Amazon S3 コンテンツへのアクセスを制限する](https://docs.aws.amazon.com/ja_jp/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-s3.html)

S3 バケットをオリジンアクセスアイデンティティを使用して設定することで、CloudFront 経由でのみアクセス可能にできる。



## キャッシュの最適化と可用性

[キャッシュヒット率の向上](https://docs.aws.amazon.com/ja_jp/AmazonCloudFront/latest/DeveloperGuide/cache-hit-ratio.html)

リクエストヘッダーや Cookie をキャッシュキーとして使用すると、レスポンスが同じ場合でも、ヘッダーや Cookie ごとに別々にキャッシュされる。そのため、特定のキャッシュキーのみに絞ることでキャッシュヒット率の向上を見込める。



## トラブルシューティング

[トラブルシューティング](https://docs.aws.amazon.com/ja_jp/AmazonCloudFront/latest/DeveloperGuide/Troubleshooting.html)

[オリジンからのエラーレスポンスのトラブルシューティング](https://docs.aws.amazon.com/ja_jp/AmazonCloudFront/latest/DeveloperGuide/troubleshooting-response-errors.html)



## リクエストとレスポンスの動作

[Amazon S3 オリジンに対するリクエストと応答の動作](https://docs.aws.amazon.com/ja_jp/AmazonCloudFront/latest/DeveloperGuide/RequestAndResponseBehaviorS3Origin.html)

* クライアント側が X-Forwarded-For を含めない場合でも、CloudFront 側で追加し、オリジンに転送する。
* 有効期限切れの場合、S3 バケットにオブジェクトの情報を問い合わせる。キャッシュにあるものが最新の場合は、HTTP ステータスコード 304(変更なし) がオリジンから返却される。
* オリジン接続タイムアウト（接続を試行する時間、回数を設定可能）
* オリジン応答タイムアウト（レスポンスを受け取るまで
待機時間）はデフォルトでは 30 秒。GET, HEAD の場合は試行回数の設定値分リトライする。


[CloudFront がオリジンからの HTTP 4xx および 5xx ステータスコードを処理してキャッシュする方法](https://docs.aws.amazon.com/ja_jp/AmazonCloudFront/latest/DeveloperGuide/HTTPStatusCodes.html)

カスタムエラーページの設定有無により動作が異なる。



## カスタムエラー応答の生成

[カスタムエラー応答の生成](https://docs.aws.amazon.com/ja_jp/AmazonCloudFront/latest/DeveloperGuide/GeneratingCustomErrorResponses.html)

HTTP ステータスコードに対応して、カスタムエラーメッセージを表示するように設定可能。



## Blackbelt

[20201028 AWS Black Belt Online Seminar Amazon CloudFront deep dive](https://www.slideshare.net/AmazonWebServicesJapan/20201028-aws-black-belt-online-seminar-amazon-cloudfront-deep-dive)

* P19: 用語集
* P20: CloudFront の設定項目
  * ディストリビューションに関連する設定
  * オリジンに関連する設定
  * Behavior に関連する設定
  * ディストリビューションに関連する機能
* P22: ディストリビューション
  * CNAME, エイリアスレコードを使用して代替ドメイン名の指定が可能。ACM で発行した証明書も利用可能。
  * WAF 連携可能
  * アクセスログ
* P31: オリジン
  * Origin Shild による負荷の軽減
  * オリジン応答タイムアウト、Keep-Alive Timeout
  * カスタムヘッダーの付与、上書きが可能
  * OAI(Origin Access Identity) によるオリジンの保護
  * オリジンフェイルオーバー
* P42: Behavior
  * URL パスごとに設定可能。
  * キャッシュポリシー(TTL、キャッシュキーとして使用する項目)
  * オリジンリクエストポリシー(キャッシュキーには含めないが、オリジンリクエストに含めるものを設定)
  * 圧縮
* P66: ディストリビューションに関連する機能
  * エラーレスポンスのカスタマイズ
  * 地域制限
  * キャッシュファイルの無効化



# 参考

* Document
  * [Amazon CloudFront とは](https://docs.aws.amazon.com/ja_jp/AmazonCloudFront/latest/DeveloperGuide/Introduction.html)
* サービス紹介ページ
  * [Amazon CloudFront](https://aws.amazon.com/jp/cloudfront/)
  * [よくある質問](https://aws.amazon.com/jp/cloudfront/faqs/)
* [ナレッジセンター](https://aws.amazon.com/jp/premiumsupport/knowledge-center/#Amazon_CloudFront)
* Black Belt
  * [20190730 AWS Black Belt Online Seminar Amazon CloudFrontの概要](https://www.slideshare.net/AmazonWebServicesJapan/20190730-aws-black-belt-online-seminar-amazon-cloudfront)
  * [20201028 AWS Black Belt Online Seminar Amazon CloudFront deep dive](https://www.slideshare.net/AmazonWebServicesJapan/20201028-aws-black-belt-online-seminar-amazon-cloudfront-deep-dive)
