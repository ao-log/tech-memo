# Document

[Amazon GuardDuty とは](https://docs.aws.amazon.com/ja_jp/guardduty/latest/ug/what-is-guardduty.html)

* EKS の監査ログ等もモニタリング対象とすることができる
* Malware Protection により EBS ボリュームを分析できる


[GuardDuty の開始方法](https://docs.aws.amazon.com/ja_jp/guardduty/latest/ug/guardduty_settingup.html)

* Amazon GuardDuty は、VPC Flow Logs など独立したストリームを取得するため、別途設定を行う必要はない


[検出結果タイプ](https://docs.aws.amazon.com/ja_jp/guardduty/latest/ug/guardduty_finding-types-active.html)

* 各 AWS サービスごとの検出結果の種別はこちらにまとまっている


[検出結果の管理](https://docs.aws.amazon.com/ja_jp/guardduty/latest/ug/findings_management.html)

* コンソール上から検出結果をフィルタリングして確認できる
* S3 にエクスポートできる
* EventBridge にイベント発行できる



# BlackBelt

[【AWS Black Belt Online Seminar】Amazon GuardDuty](https://pages.awscloud.com/rs/112-TZM-766/images/20180509_AWS-BlackBelt_Amazon-GuardDuty.pdf)

* 脅威的リスクを検知するサービス。以下のログを使用
  * VPC Flow Logs
  * AWS CloudTrail Event Logs
  * DNS Logs
* 検出時は EventBridge → Lambda のように連携できる



# 参考

* Document
  * [Amazon GuardDuty とは](https://docs.aws.amazon.com/ja_jp/guardduty/latest/ug/what-is-guardduty.html)
* Black Belt
  * [【AWS Black Belt Online Seminar】Amazon GuardDuty](https://pages.awscloud.com/rs/112-TZM-766/images/20180509_AWS-BlackBelt_Amazon-GuardDuty.pdf)


