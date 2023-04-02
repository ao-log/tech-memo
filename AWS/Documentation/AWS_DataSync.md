# BlackBelt

[AWS DataSync](https://pages.awscloud.com/rs/112-TZM-766/images/20210316_AWSBlackBelt2021_AWS-DataSync-v2.pdf)

* オンプレミス - AWS ストレージサービス間もしくは AWS ストレージサービス間のデータの移動を行えるフルマネージドサービス
* 特徴
  * エージェント、サービス間で複数の TCP セッションを生成い並列に転送
  * エージェント、サービス間はデータを圧縮
  * 複数のエンドポイントを生成しロードバランス
* 構成
  * オンプレミスから転送する場合: DataSync Agent が必要。仮想アプライアンスとして提供されている。vSphere, Hyper-V, KVM に対応
  * 自アカウント内の AWS ストレージサービス間: Agent は不要
  * クロスアカウント: EC2 Agent を配置することで対応可能
* 用語
  * エージェント: データの転送に使用する VM
  * ロケーション: 送信元、送信先のストレージ
  * タスク: 送信元、送信先、データ転送に関する各種の設定
  * タスク実行: 実際に実行されたタスク



# 参考

* Document
  * [AWS DataSync の概要](https://docs.aws.amazon.com/ja_jp/datasync/latest/userguide/what-is-datasync.html)
* Black Belt
  * [AWS DataSync](https://pages.awscloud.com/rs/112-TZM-766/images/20210316_AWSBlackBelt2021_AWS-DataSync-v2.pdf)


