# Document

[AWS Service Catalog とは?](https://docs.aws.amazon.com/ja_jp/servicecatalog/latest/adminguide/introduction.html)


[使用開始](https://docs.aws.amazon.com/ja_jp/servicecatalog/latest/adminguide/getstarted.html)

一連の流れを確認できるチュートリアル。



# BlackBelt

[AWS Service Catalog](https://pages.awscloud.com/rs/112-TZM-766/images/20180718_AWS_BlackBelt_AWSServiceCatalog_public.pdf)

* 概念
  * 製品: CloudFormation テンプレートをパッケージ化したもの。バージョン管理可能
  * ポートフォリオ: 製品の集合。ポートフォリオの単位でユーザーへのアクセスを許可
  * 制約: デプロイ方法を制限するもの。パラメータに制約をかけるなど
  * プロビジョニングされた製品: 起動済みの製品のインスタンス
* 管理者用のコンソールとエンドユーザー用のコンソールがある。管理者はポートフォリオや製品を登録。エンドユーザーは製品の検索、起動
* IAM 権限
  * 管理者用
    * AWSServiceCatalogAdminFullAccess: 管理コンソールビューへのフルアクセス権と、製品とポートフォリオの作成および管理
の権限を付与
    * ServiceCatalogAdminReadOnlyAccess: 管理者コンソールビューへのフルアクセス権を付与。製品とポートフォリオを作
成または管理するためのアクセス権はなし。
  * エンドユーザー用
    * AWSServiceCatalogEndUserFullAccess: エンドユーザーコンソールビューへのフルアクセス権を付与。製品を起動し、プ
ロビジョニング済み製品を管理するアクセス権を付与
    * ServiceCatalogEndUserAccess: エンドユーザーコンソールビューへの読み取り専用アクセス権を付与。製品を起
動し、プロビジョニング済み製品を管理するアクセス権はなし



# 参考

* Document
  * [AWS Service Catalog とは?](https://docs.aws.amazon.com/ja_jp/servicecatalog/latest/adminguide/introduction.html)
* Black Belt
  * [AWS Service Catalog](https://pages.awscloud.com/rs/112-TZM-766/images/20180718_AWS_BlackBelt_AWSServiceCatalog_public.pdf)


