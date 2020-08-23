# AWS ElasticBeanstalk

## 対応している言語
Go、Java、.NET、Node.js、PHP、Python、Ruby 

## 用語

* アプリケーション：　最上位に位置。アプリケーション内に環境、アプリケーションバージョンを作成。
* アプリケーションバージョン：　S3 上のオブジェクトとして配置してあるアプリケーションコード。
* 環境：　AWS リソースを管理する単位。
* 環境枠：　ウェブサーバー環境枠 or ワーカー環境枠
* プラットフォーム：　OS、プログラム言語ランタイム、ウェブサーバー、アプリケーションサーバーの組み合わせ。

## アクセス許可

[アクセス許可](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/concepts-roles.html)

* サービスロール； ElasticBeanstalk がヘルスステータスの情報収集に使用
* インスタンスプロファイル: EC2 インスタンスに付与。S3 へのアクションの許可など。
* ユーザーポリシー: ユーザに設定する権限。ElasticBeanstalk 用の管理ポリシーもある。

## プラットフォーム

[Elastic Beanstalk プラットフォームのサポートポリシー](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/platforms-support-policy.html)

プラットフォームブランチのコンポーネント (OS、ランタイム、アプリケーションサーバー、またはウェブサーバー) がサプライヤによって EOL (End of Life) とマークされた場合、Elastic Beanstalk はプラットフォームブランチを廃止としてマーク。


## プラットフォームの拡張

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
        |-- nginx.conf
        `-- conf.d/
            `-- custom.conf
    `-- hooks/  # Platform hooks。所定のタイミングでスクリプトを実行できる。
        |-- prebuild/
        |   |-- 01_set_secrets.sh
        |   `-- 12_update_permissions.sh
        |-- predeploy/
        |   `-- 01_some_service_stop.sh
        `-- postdeploy/
            |-- 01_set_tmp_file_permissions.sh
            |-- 50_run_something_after_deployment.sh
            `-- 99_some_service_start.sh
```

hook がどこで実行されるかは上記ドキュメントの「インスタンスデプロイワークフロー」を参照のこと。

## カスタムプラットフォーム

[Elastic Beanstalk カスタムプラットフォーム](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/custom-platforms.html)

プラットフォーム全体を開発できる。

## Docker の使用

[Docker コンテナからの Elastic Beanstalk アプリケーションのデプロイ](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/create_deploy_docker.html)

Docker コンテナからデプロイすることも可能。

## チュートリアル

各言語ごとにチュートリアルがある。
例えば flask のチュートリアルはこちら。ローカル環境で Python アプリを開発し、EB CLI でデプロイする流れ。

[Elastic Beanstalk への flask アプリケーションのデプロイ](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/create-deploy-python-flask.html)


## アプリケーション

#### アプリケーションバージョン

バージョン数はクォータがある。ライフサイクルポリシーにより、どのように削除するかのポリシーを設定可能。

[アプリケーションソースバンドルを作成する](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/applications-sourcebundle.html)

親ディレクトリを含めない状態にして ZIP ファイルに固めればよい。

## 環境

#### デプロイ

[Elastic Beanstalk 環境へのアプリケーションのデプロイ](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/using-features.deploy-existing-version.html)

次のデプロイポリシーから選ぶことができる。

* All at once (一度にすべて)
* Rolling (ローリング)
* Rolling with additional batch (追加バッチによるローリング)
* Immutable (変更不可)
* Traffic splitting (トラフィック分割)

#### 設定変更

[設定変更](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/environments-updating.html)

設定の多くはダウンタイムやリソースの置き換えなく対応可能。
起動設定や VPC など一部の設定は置き換え必要。置き換え方法を次の２つから選ぶことができる。

* ローリング更新
* イミュータブルな更新

#### プラットフォーム更新

[Elastic Beanstalk 環境のプラットフォームバージョンの更新](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/using-features.platform.upgrade.html)

[マネージドプラットフォーム更新](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/environment-platform-update-managed.html)

メンテナンスウィンドウ内にプラットフォームの更新を開始するように設定可能。2019 年 11 月 25 日以降に作成された環境では可能な限りデフォルトで有効になっている。

#### 環境のコンポーネント

[Elastic Beanstalk 環境の Auto Scaling グループ](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/using-features.managing.as.html)

デフォルトではアウトバウンドネットワークトラフィックに基づいたスケーリングになる。

[Elastic Beanstalk 環境のロードバランサー](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/using-features.managing.elb.html)

[Elastic Beanstalk で Amazon Virtual Private Cloud (Amazon VPC) を設定する](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/using-features.managing.vpc.html)

## 環境の設定（アドバンスト）


## 環境のモニタリング

[環境のモニタリング](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/environments-health.html)

[ベーシックヘルスレポート](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/using-features.healthstatus.html)

ヘルスステータスの色の判定方法などが書かれている。

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


# 参考


* [AWS Elastic Beanstalk とは](https://docs.aws.amazon.com/ja_jp/elasticbeanstalk/latest/dg/Welcome.html)
* [よくある質問](https://aws.amazon.com/jp/elasticbeanstalk/faqs/)
* [AWS Black Belt Online Seminar 2017 AWS Elastic Beanstalk](https://www.slideshare.net/AmazonWebServicesJapan/aws-black-belt-online-seminar-2017-aws-elastic-beanstalk)
* [AWS Black Belt Online Seminar 2017 AWS体験ハンズオン～Deploy with EB CLI編～](https://www.slideshare.net/AmazonWebServicesJapan/aws-black-belt-online-seminar-2017-awsdeploy-with-eb-cli)
