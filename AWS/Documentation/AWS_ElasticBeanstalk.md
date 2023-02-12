# AWS ElasticBeanstalk

[AWS Elastic Beanstalk とは](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/Welcome.html)

対応している言語

* Go
* Java
* .NET
* Node.js
* PHP
* Python
* Ruby 


[Elastic Beanstalk を使用して開始する](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/GettingStarted.html)

ここでは次の流れで対応している。

1. アプリケーション、環境の作成
1. 作成した環境の確認
1. 新しいバージョンのデプロイ
1. 環境の設定を変更
1. クリーンアップ


[Elastic Beanstalk の概念](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/concepts.html)

用語。

* アプリケーション： 最上位に位置。アプリケーション内に環境、アプリケーションバージョンを作成。
* アプリケーションバージョン： S3 上のオブジェクトとして配置してあるアプリケーションコード。
* 環境： AWS リソースを管理する単位。
* 環境枠： ウェブサーバー環境枠 or ワーカー環境枠
* プラットフォーム： OS、プログラム言語ランタイム、ウェブサーバー、アプリケーションサーバーの組み合わせ。
* 保存された設定: 環境設定を記述したテンプレート



## アクセス許可

[アクセス許可](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/concepts-roles.html)

* サービスロール； ELB、Auto Scaling グループなど BeansTalk のサービスが他のサービスを呼び出すときに使用するロール。
* インスタンスプロファイル: EC2 インスタンスに付与。S3 へのアクションの許可など。
* ユーザーポリシー: ユーザに設定する権限。ElasticBeanstalk 用の管理ポリシーもある。



## プラットフォーム

[Elastic Beanstalk プラットフォームの用語集](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/platforms-glossary.html)

プラットフォームのバージョンには、X.Y.Z 形式。X はメジャーバージョン、Y はマイナーバージョン、Z はパッチバージョン。


[Elastic Beanstalk でサポートされているプラットフォーム](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/concepts.platforms.html)

プラットフォームブランチのコンポーネント (OS、ランタイム、アプリケーションサーバー、またはウェブサーバー) がサプライヤによって EOL (End of Life) とマークされた場合、Elastic Beanstalk はプラットフォームブランチを廃止としてマークする。


[Elastic Beanstalk Linux プラットフォームの拡張](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/platforms-linux-extend.html)

色々と拡張できる。上記ドキュメントからの引用になるが、以下のように目的に従ってファイルを配置する。

```
~/my-app/
|-- web.jar  # アプリケーション
|-- Procfile  # 起動コマンドのカスタマイズ
|-- readme.md
|-- .ebextensions/  # 環境のカスタマイズ
|   |-- options.config        # Option settings
|   `-- cloudwatch.config     # Other .ebextensions sections, for example files and container commands
`-- .platform/
    `-- nginx/                # Proxy configuration
    |   |-- nginx.conf
    |   `-- conf.d/
    |       `-- custom.conf
    |-- hooks/  # Platform hooks。所定のタイミングでスクリプトを実行できる。
    |   |-- prebuild/
    |   |   |-- 01_set_secrets.sh
    |   |   `-- 12_update_permissions.sh
    |   |-- predeploy/
    |   |   `-- 01_some_service_stop.sh
    |   `-- postdeploy/
    |       |-- 01_set_tmp_file_permissions.sh
    |       |-- 50_run_something_after_app_deployment.sh
    |       `-- 99_some_service_start.sh
    `-- confighooks/          # Configuration deployment hooks
        |-- prebuild/
        |   `-- 01_set_secrets.sh
        |-- predeploy/
        |   `-- 01_some_service_stop.sh
        `-- postdeploy/
            |-- 01_run_something_after_config_deployment.sh
            `-- 99_some_service_start.sh
```

hook がどこで実行されるかは上記ドキュメントの「インスタンスデプロイワークフロー」を参照のこと。


[Elastic Beanstalk カスタムプラットフォーム](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/custom-platforms.html)

プラットフォーム全体を開発できる。


[Docker コンテナからの Elastic Beanstalk アプリケーションのデプロイ](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/create_deploy_docker.html)

Docker コンテナからデプロイすることも可能。



## チュートリアル

各言語ごとにチュートリアルがある。
zip ファイルをダウンロードし展開することで、EB CLI などを使用しデプロイできる。

例えば flask のチュートリアルはこちら。ローカル環境で Python アプリを開発し、EB CLI でデプロイする流れ。

[Elastic Beanstalk への flask アプリケーションのデプロイ](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/create-deploy-python-flask.html)



## アプリケーション

[アプリケーションバージョンの管理](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/applications-versions.html)

アプリケーションをデプロイしたときにアプリケーションバージョンが作成される。
デプロイせず、アプリケーションバンドルをアップロードすることも可能。


[アプリケーションバージョンライフサイクルの設定](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/applications-lifecycle.html)

バージョン数にはクォータがある。
アプリケーションバージョンライフサイクルポリシーを適用することで、どのように削除するかのポリシーを設定可能。


[アプリケーションソースバンドルを作成する](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/applications-sourcebundle.html)

バンドルには親フォルダを含めないように ZIP (or WAR)に固める必要がある。



## 環境の管理

[Elastic Beanstalk 環境マネジメントコンソールを使用する](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/environments-console.html)

* 設定はアプリケーションに保存できる。また同環境、もしくは別の環境でロードできる。
* CNAME スワップを行うことができる。よって、新環境を作成しテストしたあとにスワップすることで、Blue/Green Deployment を行うことができる。
* 環境のクローンを作成できる。
* 環境の再構築を行うことができる(DB インスタンスがある場合は、既存のものは削除されてしまう)。
* 環境の終了(DB インスタンスがある場合は、削除されてしまう)。
* 環境の復元。1 時間以内であれば可能。1 時間を過ぎた場合はアプリケーションの概要ページから復元可能。


#### デプロイ

[Elastic Beanstalk 環境へのアプリケーションのデプロイ](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/using-features.deploy-existing-version.html)

デプロイ時に次のデプロイポリシーから選ぶことができる。

* All at once (一度にすべて)
* Rolling (ローリング)
* Rolling with additional batch (追加バッチによるローリング)
* Immutable (新しく Auto Scaling グループを作成し、そちらにデプロイ)
* Traffic splitting (Canary テストを行いながらデプロイ)

また、aws:elasticbeanstalk:command 名前空間のオプションでも設定可能。
.ebextensions の場合の例。

```yaml
option_settings:
  aws:elasticbeanstalk:command:
    DeploymentPolicy: Rolling
    BatchSizeType: Percentage
    BatchSize: 25
```


[Elastic Beanstalk を使用したブルー/グリーンデプロイ](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/using-features.CNAMESwap.html)

新しい環境をデプロイしてから CNAME をスワップする方法。


#### 設定変更

[設定変更](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/environments-updating.html)

設定の多くはダウンタイムやリソースの置き換えなく対応可能。
起動設定や VPC など一部の設定は置き換え必要。置き換え方法を次の２つから選ぶことができる。

* ローリング更新
* イミュータブルな更新(新しい Auto Scaling Group にデプロイ)


#### プラットフォーム更新

[Elastic Beanstalk 環境のプラットフォームバージョンの更新](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/using-features.platform.upgrade.html)

[マネージドプラットフォーム更新](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/environment-platform-update-managed.html)

* メンテナンスウィンドウ内にプラットフォームの更新を開始するように設定可能。
* メジャーバージョンは対象外。マイナーおよびパッチ、パッチのみのどちらかの設定が可能。
* 2019 年 11 月 25 日以降に作成された環境では可能な限りデフォルトで有効になっている。


#### 環境タイプ

[環境タイプ](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/using-features-managing-env-types.html)

* シングルインスタンス環境: 一つの EC2 インスタンス。Elastic IP が付与されている。

途中から環境タイプを変更することも可能。


#### ワーカー環境

[Elastic Beanstalk ワーカー環境](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/using-features-managing-env-tiers.html)



## 環境の設定

#### 環境のコンポーネント

[Elastic Beanstalk 環境の Auto Scaling グループ](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/using-features.managing.as.html)

デフォルトではアウトバウンドネットワークトラフィックに基づいたスケーリングになる。



## 環境の設定（アドバンスト）

デフォルトの設定値は次の方法で上書き可能。

* 設定ファイル
* 保存済み設定
* コマンドラインオプション(AWS CLI, EB CLI)
* Elastic Beanstalk API


[設定オプション](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/command-options.html)

* 設定オプションは、aws:autoscaling:asg のような名前空間で構成されている。
* 複数箇所で同じ項目を設定されている場合は [優先順位](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/command-options.html#configuration-options-precedence) に基づいて適用。


[環境を作成する前](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/environment-configuration-methods-before.html)

各設定方法について書かれている。

**.ebextensions**

YAML or JSON 形式で記述可能。

```yaml
option_settings:
  - namespace:  aws:elasticbeanstalk:application
    option_name:  Application Healthcheck URL
    value:  /health
```

**保存済み設定**

aws elasticbeanstalk create-configuration-template などの方法で設定を保存可能。

**EB CLI**

eb config コマンドによる設定。


[すべての環境に対する汎用オプション](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/command-options-general.html)

汎用オプションの一覧。


#### .ebextensions

[オプション設定](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/ebextensions-optionsettings.html)

option_settings キーにより環境の設定を変更可能。


[Linux サーバーでのソフトウェアのカスタマイズ](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/customize-containers-ec2.html)

以下の設定が可能。

* パッケージ
* グループ
* ユーザー
* ソース
* ファイル
* コマンド
* サービス
* コンテナコマンド


[Elastic Beanstalk の保存された設定を使用する](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/environment-configuration-savedconfig.html)

環境の設定を S3 のオブジェクトとして YAML 形式で保存可能。


[Elastic Beanstalk 環境の HTTPS の設定](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/configuring-https.html)

.ebextensions で ACM の ARN を指定できる。

```yaml
option_settings:
  aws:elb:listener:443:
    SSLCertificateId: arn:aws:acm:us-east-2:1234567890123:certificate/xxxx
    ListenerProtocol: HTTPS
    InstancePort: 80
```


[秘密キーを Amazon S3 に安全に保存する](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/https-storingprivatekeys.html)

秘密鍵は S3 に格納しておき、デプロイ時にダウンロード可能。



## 環境のモニタリング

[ベーシックヘルスレポート](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/using-features.healthstatus.html)

ヘルスステータスの色

* Grey: 環境の更新中。
* Green: 正常
* Yellow: 1 つ以上のヘルスチェックで失敗
* Red: 3 つ以上のヘルスチェックで失敗。リクエストが一貫して失敗

ELB を使用している環境では 10 秒価格でヘルスチェックのリクエストを送信。タイムアウトは 5 秒。5 回連続して失敗した場合、対象のインスタンスをアービスから除外する。

基本的なヘルスレポートでは CloudWatch にメトリクスを公開しない。リソースから発行されたメトリクスのみ使用可能。


[拡張状態ヘルスレポートおよびモニタリング](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/health-enhanced.html)

* ヘルスエージェントをインスタンス上で稼働させる必要がある。
* ベーシックヘルスレポートでは 5 分間隔だが、拡張ヘルスレポートの場合は 10 秒間隔。環境全体の状態は 60 秒間隔で CloudWatch に公開。
* ウェブサーバー環境の場合、環境内またはバッチ内の各インスタンスが 2 分間にわたって連続して行われる 12 のヘルスチェックに合格する必要がある。


[環境の Amazon CloudWatch カスタムメトリクスの発行](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/health-enhanced-cloudwatch.html)

EnvironmentHealth: 拡張ヘルスレポートシステムからパブリッシュされる CloudWatch メトリクス。ステータスに対応する数値がメトリクスとして報告される。

* 0 – OK
* 1 – Info
* 5 – Unknown
* 10 – No data
* 15 – Warning
* 20 – Degraded
* 25 – Severe


[拡張ヘルスログ形式](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/health-enhanced-serverlogs.html)

所定のパスに所定の書式でログ出力することで、ヘルスチェック結果の判定材料として使用できる。


[インスタンスログの表示](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/using-features.logging.html)

* バンドルログの採取方法
* CloudWatch Logs へのログ送信



## EB CLI

[EB CLI による Elastic Beanstalk 環境の管理](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/eb-cli3-getting-started.html)

```shell
eb create    # 環境作成
eb status    # 環境のステータス確認
eb health    # インスタンスに関するヘルス情報
eb events    # 環境のイベント確認
eb logs      # 環境のログ
eb oprn      # Web ページを開く
eb deploy    # デプロイ
eb config    # 設定オプションの確認
eb terminate # 環境の終了
```



## AWS サービスの統合

[Amazon CloudWatch Logs で Elastic Beanstalk を使用する](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/AWSHowTo.cloudwatchlogs.html)

指定したパスのログファイルを CloudWatch Logs に送信するように設定できる。



# 参考

* Document
  * [AWS Elastic Beanstalk とは](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/Welcome.html)
* サービス紹介ページ
  * [AWS Elastic Beanstalk](https://aws.amazon.com/jp/elasticbeanstalk/)
  * [よくある質問](https://aws.amazon.com/jp/elasticbeanstalk/faqs/)
* Black Belt
  * [AWS Black Belt Online Seminar 2017 AWS Elastic Beanstalk](https://d1.awsstatic.com/webinars/jp/pdf/services/20170111_AWS-Blackbelt-Elastic-Beanstalk.pdf)
  * [AWS Black Belt Online Seminar 2017 AWS体験ハンズオン～Deploy with EB CLI編～](https://www.slideshare.net/AmazonWebServicesJapan/aws-black-belt-online-seminar-2017-awsdeploy-with-eb-cli)

