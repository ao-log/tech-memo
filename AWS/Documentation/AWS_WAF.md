# BlackBelt

[AWS Managed Rules for AWS WAF の活⽤](https://pages.awscloud.com/rs/112-TZM-766/images/202203_AWS_Black_Belt_AWS_Managed_Rules_for_AWS_WAF.pdf)

* ルールグループにはいくつかの種類がある
  * ベースラインルールグループ: 一般的な対策ルールを提供
  * ユースケース別ルールグループ: SQL インジェクションなどユースケースごとのルールを提供
  * IP レピュテーションルールグループ: アクセス元の IP アドレスに基づいてリクエストをブロック
  * Bot コントロールルールグループ: ボットと判断されるトラフィックを検知し、それに基づいたアクションを行う
* 導入ステップ
  * 初めは Count にしておき、検知率や誤検出の有無を確認してから Block へと変更する


[AWS WAF でできる Bot 対策](https://pages.awscloud.com/rs/112-TZM-766/images/202206_AWS_Black_Belt_AWS_WAF_Bot_mitigation.pdf)



# 参考

* Document
  * [AWS WAF](https://docs.aws.amazon.com/ja_jp/waf/latest/developerguide/waf-chapter.html)
* Black Belt
  * [AWS Managed Rules for AWS WAF の活⽤](https://pages.awscloud.com/rs/112-TZM-766/images/202203_AWS_Black_Belt_AWS_Managed_Rules_for_AWS_WAF.pdf)
  * [AWS WAF でできる Bot 対策](https://pages.awscloud.com/rs/112-TZM-766/images/202206_AWS_Black_Belt_AWS_WAF_Bot_mitigation.pdf)

