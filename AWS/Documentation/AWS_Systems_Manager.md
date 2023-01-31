# Document

[AWS Systems Manager とは?](https://docs.aws.amazon.com/ja_jp/systems-manager/latest/userguide/what-is-systems-manager.html)


[EC2 インスタンス用 AWS Systems Manager のセットアップ](https://docs.aws.amazon.com/ja_jp/systems-manager/latest/userguide/systems-manager-setting-up-ec2.html)


[AWS Systems Manager Parameter Store](https://docs.aws.amazon.com/ja_jp/systems-manager/latest/userguide/systems-manager-parameter-store.html)


[AWS Systems Manager Automation](https://docs.aws.amazon.com/ja_jp/systems-manager/latest/userguide/systems-manager-automation.html)


[AWS Systems Manager Session Manager](https://docs.aws.amazon.com/ja_jp/systems-manager/latest/userguide/session-manager.html)


[AWS Systems Manager Run Command](https://docs.aws.amazon.com/ja_jp/systems-manager/latest/userguide/execute-remote-commands.html)


[AWS Systems Manager ドキュメント](https://docs.aws.amazon.com/ja_jp/systems-manager/latest/userguide/sysman-ssm-docs.html)



# BlackBelt

[[AWS Black Belt Online Seminar] AWS Systems Manager](https://pages.awscloud.com/rs/112-TZM-766/images/20200212_AWSBlackBelt_SystemsManager_0214.pdf)

* EC2 インスタンスをマネージドインスタンスにするには SSM Agent を導入するとよい
  * オフィシャルイメージにはプリインストール済み
  * インターネットへのアウトバウンドの疎通性もしくは VPC Endpoint の用意が必要
  * IAM ロールに AmazonSSMManagedInstanceCore 管理ポリシーが必要
  * 一括実行の単位となるリソースグループを作成しておくと管理しやすいしておくとよい
* SSM インベントリ: アプリケーション一覧などの構成情報を記録、可視化
* SSM ドキュメント
* Automation: ワークフローを実行できる
* ステートマネージャ
* メンテナンスウィンドウ
* セッションマネージャ
* パラメータストア



# 参考

* Document
  * [AWS Systems Manager とは?](https://docs.aws.amazon.com/ja_jp/systems-manager/latest/userguide/what-is-systems-manager.html)
* Black Belt
  * [[AWS Black Belt Online Seminar] AWS Systems Manager](https://pages.awscloud.com/rs/112-TZM-766/images/20200212_AWSBlackBelt_SystemsManager_0214.pdf)


