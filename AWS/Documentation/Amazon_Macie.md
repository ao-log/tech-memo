# BlackBelt

[【AWS Black Belt Online Seminar】Amazon Macie](https://pages.awscloud.com/rs/112-TZM-766/images/20200812_AWS-BlackBelt-Macie.pdf)

* S3 バケットをスキャンして機密情報を含むオブジェクトをレポート
* 機密情報は姓名、誕生日、クレジットカード番号などの個人情報などや鍵情報など
* 暗号化されていないオブジェクト、バージョニングの設定状況などをチェックできる
* スキャンはジョブを作成して定義。ジョブでスケジュール、サンプリング深度(全体の何 % をスキャンするか)、フィルタ条件などを設定
* 暗号化されている場合もスキャン可能だが、SSE-KMS の場合は Macie のサービスリンクロール AWSServiceRoleForAmazonMacie に KMS へのアクセス権が必要。クライアントサイド暗号化、カスタマー提供型のサーバーサイド暗号化（SSE-C）の場合には復号不可


# 参考

* Document
  * [Amazon Macie とは](https://docs.aws.amazon.com/ja_jp/macie/latest/user/what-is-macie.html)
* Black Belt
  * [【AWS Black Belt Online Seminar】Amazon Macie](https://pages.awscloud.com/rs/112-TZM-766/images/20200812_AWS-BlackBelt-Macie.pdf)


