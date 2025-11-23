
[SOCI Index Manifest v2 を用いた一貫性のある Amazon ECS デプロイメントの実現](https://aws.amazon.com/jp/blogs/news/improving-amazon-ecs-deployment-consistency-with-soci-index-manifest-v2/)

* SOCI Index Manifest v2 では、コンテナイメージと SOCI インデックスの間に明示的な関連付けを構築し SOCI を用いた場合でも一貫したデプロイメントを実現できるようになる


[コンテナにおけるデジタル署名](https://aws.amazon.com/jp/blogs/news/cryptographic-signing-for-containers/)

* 典型的なデジタル署名は、署名されるデータの暗号化されたハッシュと署名者のアイデンディティ情報を提供する証明書の、(少なくとも) 2 つの要素で構成される。
* 元のファイルのハッシュは署名者の秘密鍵で暗号化される
* 秘密鍵は署名者を一意に識別し、署名者以外の人が署名を生成できないことを保証する

必要な仕組み

* 承認済みイメージのみがそれぞれの環境に昇格可能であり、最終的に本番環境にデプロイできることを保証する仕組み
* 承認済みイメージが、承認後に改ざんされていないことを確認する仕組み

対応方法

* 有効な署名鍵にアクセスできる開発者やチームだけが、イメージが承認済みであることを示す署名を生成できる


[Streamline container image signatures with Amazon ECR managed signing](https://aws.amazon.com/blogs/containers/streamline-container-image-signatures-with-amazon-ecr-managed-signing/)

* ECR にコンテナイメージを push すると、自動的に AWS Signer により署名されるようにできる
* signing rule を作成し、フィルタに合致したイメージが対象となる
* 署名されるとリポジトリ内に Signature タイプのリソースができる
* 例えば ECS の場合は lifecycleHooks で指定した Lambda 関数で署名を検証する



# builders.flash

[コンテナサイズ最小化のためのベースイメージ再考](https://aws.amazon.com/jp/builders-flash/202502/base-img-for-container-minimization/)

