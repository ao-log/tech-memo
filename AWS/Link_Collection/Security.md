
[【レポート】AWS における安全な Web アプリケーションの作り方 #AWS-55 #AWSSummit](https://dev.classmethod.jp/articles/awssummit-2021-aws-55/)

* 代表的なセキュリティガイドライン
  * 安全なウェブサイトの作り方
  * OWASP Top 10
* Web アプリケーションのセキュリティ対策
  * 認証・認可、アクセス管理
    * 認証情報を埋め込まない
    * S3 バケットのアクセス管理
    * 認証・認可は既存の仕組みを活用(Amazon Cognito, AWS SSO など)
  * アプリケーションフレームワークを正しく利用
    * SQL プレースホルダ、プリペーアドステートメント
    * XSS 対策
    * セッション管理
  * ログの取得とモニタリング
    * CloudTrail
    * アクセスログ(ELB, CloudFront, API Gateway, S3, VPC Flow Logs, ...)
  * コードレビューの実施
  * セキュリティテストの実施
  * WAF による保護
  * 継続的なセキュリティテスト、脆弱性診断
