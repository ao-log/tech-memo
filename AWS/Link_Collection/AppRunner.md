
[App Runner VPCネットワーク接続時における可観測性](https://aws.amazon.com/jp/blogs/news/observability-for-aws-app-runner-vpc-networking/)

* ADOT を通して X-Ray にトレースデータを送信している。


[AWS App Runner でセマンティックバージョニングに基づいた継続的デプロイメントの実現](https://aws.amazon.com/jp/blogs/news/enable-continuous-deployment-based-on-semantic-versioning-using-aws-app-runner/)

* 継続的なデプロイメントの重要な側面の一つにセマンティックバージョニングがあ流。ソフトウェアのリリースにバージョンナンバーを割り当てる仕組みのこと
* セマンティックバージョニングの一般的なルール
  * 当該リリースが後方互換を含まない場合 (API コンストラクトの破棄など)、MAJOR バージョンをインクリメント
  * 当該リリースが後方互換を含む場合、MINOR バージョンをインクリメント
  * 当該リリースがバグフィックスのみ含まれる場合、PATCH バージョンをインクリメント
* App Runner では特定のイメージタグを追跡するため、そのままだとセマンティックバージョニングに対応できない。以下のアーキテクチャで対応
  * ECR の PUSH イベントをトリガーにした EventBridge ルールを作成。SQS → Lambda

