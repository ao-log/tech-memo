# Document

[AWS Secrets Manager の概要](https://docs.aws.amazon.com/ja_jp/secretsmanager/latest/userguide/intro.html)

* KMS で暗号化されている


[AWS Secrets Managerシークレットのローテーション](https://docs.aws.amazon.com/ja_jp/secretsmanager/latest/userguide/rotating-secrets.html)

* 指定スケジュールでシークレットをローテーションできる
  * ステージングラベルによってバージョンに紐づけることができる
  * Lambda 関数により処理する
* 処理の流れ
  * シークレットの新しいバージョンを作成。この新バージョンに対しステージングラベル AWSPENDING をセット
  * データベースまたはサービスの認証情報を更新
  * 新バージョンの認証情報で接続できるかテスト
  * 新バージョンのステージングラベルを AWSCURRENT にセット。以前のバージョンには AWSPREVIOUS をセット



# 参考

* Document
  * [AWS Secrets Manager の概要](https://docs.aws.amazon.com/ja_jp/secretsmanager/latest/userguide/intro.html)


