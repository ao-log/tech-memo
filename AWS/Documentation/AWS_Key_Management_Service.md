# Document

[AWS KMS keys ローテーション](https://docs.aws.amazon.com/ja_jp/kms/latest/developerguide/rotate-keys.html)



# BlackBelt

[AWS Key Management Service](https://pages.awscloud.com/rs/112-TZM-766/images/20181121_AWS-BlackBelt-KMS.pdf)

* カスタマーマスターキー(CMK)
  * AES256 ビットの鍵
  * KMS 内部の HSM 上にのみ平文で存在。
  * 最大 4 KB のデータを暗号化、復号
* カスタマーデータキー(CDK)
  * データの暗号化に使用する鍵
  * KMS で生成され、CMK による暗号化をした状態でデータとともに保管
  * データ復号時に CDK も復号
* 鍵のインポートが可能 (BYOK)
* キーポリシー
* Grants
  * CMK の使用を他の AWS Principal に委任できる
* API
  * Encrypt: 4 KB までの平文データを暗号化。生成される暗号文にヘッダが付与される
  * Decrypt: データを復号。ヘッダ内に CMK の情報が含まれているので、CMK の指定は不要
  * GenerateDataKey: 平文の鍵と Encrypt で暗号化された鍵を返す。平文の鍵は暗号化した後、削除すること。暗号化された鍵はデータと共に保管



# 参考

* Document
  * [AWS Key Management Service](https://docs.aws.amazon.com/ja_jp/kms/latest/developerguide/overview.html)
* Black Belt
  * [AWS Key Management Service](https://pages.awscloud.com/rs/112-TZM-766/images/20181121_AWS-BlackBelt-KMS.pdf)
