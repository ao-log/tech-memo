# Document

[AWS Config とは?](https://docs.aws.amazon.com/ja_jp/config/latest/developerguide/WhatIsConfig.html)


[Amazon S3 バケットへの設定スナップショットの配信](https://docs.aws.amazon.com/ja_jp/config/latest/developerguide/deliver-snapshot-cli.html)

* AWS Config が記録しているリソースの設定と、これらのリソース間の関係を設定スナップショットとして S3 バケットに保存できる


[AWS Config から Amazon SNS トピックに送信される通知](https://docs.aws.amazon.com/ja_jp/config/latest/developerguide/notifications-for-AWS-Config.html)

* リソース設定の変更などを通知できる


[コンフォーマンスパック](https://docs.aws.amazon.com/ja_jp/config/latest/developerguide/conformance-packs.html)

* コンフォーマンスパックは AWS Config マネージドルールまたはカスタムルールおよび修復アクションのリストを含む YAML テンプレートから作成。様々なテンプレートが用意されている



## Knowledge Center

[AWS リソースが準拠していない場合に AWS Config を使用して通知を受けるにはどうすればよいですか?](https://aws.amazon.com/jp/premiumsupport/knowledge-center/config-resource-non-compliant/)



# BlackBelt

[[AWS Black Belt Online Seminar] AWS Config](https://pages.awscloud.com/rs/112-TZM-766/images/20190618_AWS-Blackbelt_Config.pdf)

* AWS リソースの構成情報の変更をトレースできる。どの時点でどの設定であったかを確認できる
* 構成情報の保存期間は 30 日から 7 年間で設定可能。スナップショットとして S3 に保存することも可能
* 設定タイムラインでは時系列で構成情報の変更点を確認可能。構成情報だけでなく、どの IAM エンティティから変更されたかの CloudTrail の記録も紐づけられている
* リソースのクエリ。例えば、特定のセキュリティグループをアタッチしているリソース一覧をクエリできる
* Config Rules によるルール準拠状況の評価。マネージドルールのほか、カスタムルールも作成可能
* ルール評価のタイミングはリソース設定が変更されたとき。もしくは定期的
* ルールに準拠していない場合、修復アクションを実行することも可能。SSM ドキュメントによる実行、もしくは当該イベントに対し Lambda 関数を実行するように EventBridge ルールを作成する
* マルチアカウント環境のデータを集約可能。集約ビューを作成する流れ


[[AWS Black Belt Online Seminar] AWS Config Update](https://pages.awscloud.com/rs/112-TZM-766/images/20201208_AWSBlackBelt_ConfigUpdate_A.pdf)

* 適合パック: 複数の Config Rules と修復パッケージを用途に応じてパッケージ化したもの



# 参考

* Document
  * [AWS Config とは?](https://docs.aws.amazon.com/ja_jp/config/latest/developerguide/WhatIsConfig.html)
* Black Belt
  * [[AWS Black Belt Online Seminar] AWS Config](https://pages.awscloud.com/rs/112-TZM-766/images/20190618_AWS-Blackbelt_Config.pdf)
  * [[AWS Black Belt Online Seminar] AWS Config Update](https://pages.awscloud.com/rs/112-TZM-766/images/20201208_AWSBlackBelt_ConfigUpdate_A.pdf)


