# BlackBelt

[AWSアカウントシングルサインオンの設計と運用](https://pages.awscloud.com/rs/112-TZM-766/images/20200722_AWSBlackbelt_aws_sso.pdf)

* ID プロバイダー(IdP) を使用し SSO することができる
* 用語
  * シングルサインオン: 一度のユーザ認証よって、独立した複数のシステム上のリソースが利用可能になる特性
  * ID フェデレーション: 一つの組織を超えて他の管理ドメインのサービスにもログインできるようにする処理のこと。SAML などを用いて実現
* AWS におけるシングルサインオン
  * ユーザーは IdP に対して認証を行う。IdP と AWS のサービスプロバイダー間で事前に信頼関係を構築しておく。ユーザーはロールを引き受ける
  * IdP の認証に通ると SAML アサーションを受信する。SAML アサーションを AWS の SAML 用エンドポイントにポストすることで、クライアントをコンソールにリダイレクトする
  * AWS では ID プロバイダーを設定する。その際にメタデータドキュメントを指定
  * IAM ロールの信頼ポリシーの Principal には ID プロバイダーの ARN を指定
* AWS SSO
  * 管理者がアクセス権限セットを作成し、ユーザーは許可されたアクセス権限セットのロールでコンソールログインができる



# 参考

* Document
  * [IAM Identity Center とは](https://docs.aws.amazon.com/ja_jp/singlesignon/latest/userguide/what-is.html)
* Black Belt
  * [AWSアカウントシングルサインオンの設計と運用](https://pages.awscloud.com/rs/112-TZM-766/images/20200722_AWSBlackbelt_aws_sso.pdf)


