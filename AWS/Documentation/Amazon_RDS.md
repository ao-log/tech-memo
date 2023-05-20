# Document

## 概要

[Amazon RDS DB インスタンス](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/Overview.DBInstance.html)

* DB インスタンスの FQDN は db1.123456789012.us-east-1.rds.amazonaws.com のような形式。
* 作成時にマスターユーザーアカウントを作成する。DB インスタンス作成時にパスワードを設定する。


[DB インスタンスクラス](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html)

3 種類のインスタンスクラス

* スタンダード
* メモリ最適化
* バーストパフォーマンス

サポートされているデータベースエンジン。

* MariaDB
* Microsoft SQL Server
* MySQL
* Oracle
* PostgreSQL


[Amazon RDS DB インスタンスストレージ](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/CHAP_Storage.html)

データベースおよびログのストレージに EBS ボリュームを使用している。

モニタリングする際に有用な指標。

* IOPS
* レイテンシー
* スループット
* キューの深度(サービスされるのを待つキュー内の I/O リクエスト数)

以下の作業中は I/O キャパシティーを消費するためパフォーマンスが低下する可能性がある。

* マルチ AZ スタンバイの作成
* リードレプリカの作成
* ストレージタイプを変更する

システムリソースを使い切っていないにも関わらず、キャパシティーを増強してもトランザクションレートが増加しない場合、その他のボトルネック要因がある可能性がある。例えば行ロックとインデックスページロックの競合などの要因。


[Amazon RDS での高可用性 (マルチ AZ)](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/Concepts.MultiAZ.html)

* 異なる AZ 上でスタンバイレプリカが稼働する。
* シングル AZ 配置よりも書き込みとコミットのレイテンシーが増加する場合がある。
* フェイルオーバー時間は通常 60～120 秒。
* フェイルオーバー時に、DB インスタンスの DNS レコードが自動的に変更される。JVM はデフォルト設定では JVM を再起動するまで反映されないので、TTL の設定値を変える対応が必要。



## ベストプラクティス

[Amazon RDS のベストプラクティス](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/CHAP_BestPractices.html)

* メモリ、CPU、ストレージの使用状況をモニタリング。
* 最大ストレージ容量に近づいたら、DB インスタンスをスケールアップ。
* 自動バックアップを有効にする。1 日のうちで書き込み IOPS が低くなる時間帯にバックアップウィンドウを設定。
* クライアントアプリケーションが DB インスタンスの DNS データをキャッシュしている場合、有効期限 (TTL) の値を 30 秒未満に設定。
* 拡張モニタリングの有効化。
* クエリのチューニング。



## DB インスタンスの設定

[Amazon RDS Proxy による接続の管理](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/rds-proxy.html)

* プロキシは現在のライターインスタンスを自動的に判別。
* プロキシごとにターゲットグループがある。
* プロキシは、関連付けられた RDS または Aurora データベースのライターインスタンスに対して接続プーリングを実行する。
* データベース認証情報は、AWS Secrets Manager に保存する。


[オプショングループを使用する](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/USER_WorkingWithOptionGroups.html)

* DB エンジンの追加の機能を有効にして設定するためのもの。


[DB パラメータグループを使用する](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/USER_WorkingWithParamGroups.html)

* DB エンジンの設定値を管理するもの。
* 無指定時はデフォルトのパラメータグループが使用される。
* 一部パラメータは再起動不要だが、再起動が必要なパラメータもある。



## DB インスタンスの管理

[一時的に Amazon RDS DB インスタンスを停止する](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/USER_StopInstance.html)

* 停止していた場合でも、7 日経過すると自動的に起動する。


[Amazon RDS DB インスタンスを変更する](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/Overview.DBInstance.Modifying.html)

* テストインスタンスで変更をテストすることを推奨する。
* すぐに適用することもできるし、次回のメンテナンスウィンドウ時に行うようにすることもできる。


[DB インスタンスのメンテナンス](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Maintenance.html)

* マルチ AZ の場合、まずスタンバイ側にメンテナンスを実行したあとプライマリに昇格し、旧プライマリ側でメンテナンスする動作となる。
* メンテナンスはメンテナンスウィンドウ時に適用される。
* メンテナンスウィンドウ時に適用されず手動で対応が必要なものもある。


[DB インスタンスのエンジンバージョンのアップグレード](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Upgrading.html)

* メジャーバージョンのアップグレード、マイナーバージョンのアップグレードの 2 種類がある。
* メジャーバージョンのアップグレードは互換性のない変更が導入される場合がある。
* マイナーバージョン自動アップグレードの設定があり有効化すると自動でアップグレードされる。


[リードレプリカの使用](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/USER_ReadRepl.html)

* Amazon RDS は、ソースインスタンスのスナップショットを作成し、このスナップショットから読み取り専用インスタンスを作成する。次に、プライマリ DB インスタンスが変更されるたびに、DB エンジンの非同期レプリケーションを使用してリードレプリカを更新する。


[DB インスタンスの削除](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/USER_DeleteInstance.html)

* 削除保護を設定することが可能。
* 削除時に最終 DB スナップショットを作成可能。また、自動バックアップを保持するかどうか選択可能。



## Amazon RDS DB インスタンスのバックアップと復元

[バックアップの使用](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html)

* DB インスタンス全体をバックアップしている。
* 最初のスナップショットは、フル DB インスタンスのデータ。以降は差分のみ。
* スナップショットはコピー、共有が可能。
* 自動バックアップはバックアップウィンドウ中に行われる。バックアッププロセスの開始時にストレージ I/O が一時中断することがある(通常は数秒間)。
* デフォルトのバックアップ保持期間は 7 日。0 〜 35 日で設定可能。0 の場合は無効。


[DB スナップショットの作成](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/USER_CreateSnapshot.html)

* Single-AZ DB インスタンスでこの DB スナップショットを作成すると、I/O が短時間中断。
* Multi-AZ の場合はバックアップはスタンバイから取得されるため I/O の中断は発生しない。ただし、SQL Server の場合、一時中断する。


[DB のスナップショットからの復元](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/USER_RestoreFromSnapshot.html)

* 復元する場合は新規 DB インスタンスが作成される動作。


[ポイントインタイムリカバリ](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/USER_PIT.html)

* DB インスタンスを特定の時点に復元し、新しい DB インスタンスを作成することが可能。



## DB インスタンスのモニタリング

[Amazon RDS ​のモニタリングの概要](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/MonitoringOverview.html)

以下のサービス、ツールが使用可能。

* Amazon RDS レポート作成ツール
  * Amazon RDS イベント
  * データベースのログファイル 
  * Amazon RDS 拡張モニタリング 
  * Amazon RDS Performance Insights
  * Amazon RDS 推奨事項
* AWS サービス
  * Amazon CloudWatch
  * AWS CloudTrail 
  * Trusted Advisor 


[Amazon RDS 推奨事項を使用する](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/USER_Recommendations.html)

* 推奨事項が移動で表示されるようになっている。


[拡張モニタリング](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html)

* OS のリアルタイムのメトリクス。
* デフォルトでは拡張モニタリングメトリクスは 30 日間 CloudWatch Logs に保存される。


[Amazon RDSパフォーマンスインサイトの使用](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/USER_PerfInsights.html)

[Amazon RDS イベント通知の使用](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/USER_Events.html)

[Amazon RDS データベースログファイル](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/USER_LogAccess.html)



## セキュリティ

[Amazon RDS でのデータベース認証](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/database-authentication.html)

次の３つの認証方法がある。

* パスワード認証
* IAM データベース認証
* Kerberos 認証


[MariaDB、MySQL、および PostgreSQL の IAM データベース認証](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html)

* DB インスタンス側の設定で IAM データベース認証を有効化しておく必要がある
* 「rds-db:connect」のアクションに対する許可が必要
* MariaDB、MySQL では、認AWSAuthenticationPlugin によって処理される
```
CREATE USER jane_doe IDENTIFIED WITH AWSAuthenticationPlugin AS 'RDS'; 
```
* PostgreSQL ではユーザーに rds_iam ロールを付与する
```
CREATE USER db_userx; 
GRANT rds_iam TO db_userx;
```
* 認証トークンを使用して接続する。認証トークンの有効期限は 15 分だが、認証に成功した後の確立後のセッションには影響しない
* 認証トークンは AWS CLI の場合は次のように取得できる
```
aws rds generate-db-auth-token \
   --hostname rdsmysql.123456789012.us-west-2.rds.amazonaws.com \
   --port 3306 \
   --region us-west-2 \
   --username jane_doe
```
* 接続時は以下のように password で認証トークンを指定する。SSL 証明書は事前にダウンロードしておく必要がある
```
mysql --host=hostName \
    --port=portNumber \
    --ssl-ca=full_path_to_ssl_certificate \
    --enable-cleartext-plugin \
    --user=userName \
    --password=authToken
```



# BlackBelt

[20180425 AWS Black Belt Online Seminar Amazon Relational Database Service (Amazon RDS)](https://www.slideshare.net/AmazonWebServicesJapan/20180425-aws-black-belt-online-seminar-amazon-relational-database-service-amazon-rds-96509889)

* P20: マルチ AZ
  * フェイルオーバ時もエンドポイントは変わらない。
  * スタンバイ側はアクセス不可
* P23: リードレプリカ
* P24: スケールアップ
* P31: バックアップ
  * 自動でのスナップショット取得(35日分まで)
  * 手動スナップショットは期限なし
  * リストア(スナップショットを元にする)
  * ポイントインタイムリカバリ
* P35: リネーム: エンドポイント名を変更
* P38: パラメータグループ、オプショングループ
* P39: メンテナンスウィンドウ
* P41: 拡張モニタリング
* P44: イベント通知
* P45: ログアクセス(エラー、スロークエリなど) 
* P48: VPC、サブネットグループ
* P49: セキュリティグループ
* P50: DB インスタンスの暗号化



# 参考

* Document
  * [Amazon Relational Database Service (Amazon RDS) とは](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/Welcome.html)
  * [API Reference](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/APIReference/Welcome.html)
* サービス紹介ページ
  * [Amazon Relational Database Service (RDS)](https://aws.amazon.com/jp/rds/)
  * [よくある質問](https://aws.amazon.com/jp/rds/faqs/)
* Black Belt
  * [20180425 AWS Black Belt Online Seminar Amazon Relational Database Service (Amazon RDS)](https://www.slideshare.net/AmazonWebServicesJapan/20180425-aws-black-belt-online-seminar-amazon-relational-database-service-amazon-rds-96509889)

