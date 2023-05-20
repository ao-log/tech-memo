# BlackBelt

[AWS CloudHSM](https://pages.awscloud.com/rs/112-TZM-766/images/20190723_AWS-Blackbelt-CloudHSM_A.pdf)

* シングルテナント方式のハードウェアセキュリティモジュールを使用した暗号鍵管理サービス
* 暗号化、復号処理のアクセラレーション、オフロードも可能
* リージョン内で稼働
* エンベロープ暗号化
  * マスターキー: データキーの暗号化、復号に使用
  * データキー: データの暗号化、復号に使用
* HSM クラスター内で HSM インスタンスを管理
* loginHSM コマンドでログインした後、各種操作を行う



# 参考

* Document
  * [AWS CloudHSM の概要](https://docs.aws.amazon.com/ja_jp/cloudhsm/latest/userguide/introduction.html)
* Black Belt
  * [AWS CloudHSM](https://pages.awscloud.com/rs/112-TZM-766/images/20190723_AWS-Blackbelt-CloudHSM_A.pdf)
