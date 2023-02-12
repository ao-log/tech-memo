# EC2(Windows)

[Amazon EC2 とは](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/WindowsGuide/concepts.html)


[Amazon EC2 で実行する Windows のベストプラクティス](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/WindowsGuide/ec2-best-practices.html)

* Windows ドライバーを更新(インスタンスタイプによっては、AWS PV、ENA、NVMe の各ドライバーを更新する必要がある)
* 最新の AMI を使用
* 移行前にシステム/アプリケーションパフォーマンスをテスト
* 起動エージェントを更新
* セキュリティ
  * 最小アクセス(アクセス許可元を絞る)
  * 最小権限
  * 設定管理(パッチ適用など)
  * 変更管理
  * 監査ログ
* バックアップと復旧


[Windows インスタンスの設定](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/WindowsGuide/ec2-windows-instances.html)

[EC2Launch v2 の概要](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/WindowsGuide/ec2launch-v2-overview.html)

様々なタスクを実行できる。
また、以下のようなステージ構成となっており、一部のタスクは特定のステージでのみ実行可能。

* Boot
* Network
* PreReady
* PostReady
* UserData

実行するタスクには以下のようなものがある。

* インスタンスに関する情報をレンダリングする新しい壁紙を設定
* ローカルマシンに作成される管理者アカウントの属性を設定
* 追加ボリュームのドライブ文字を設定
* コンピュータ名を設定
* ユーザーデータを実行
* メタデータサービスと KMS サーバーに到達するように永続的な静的ルートを設定
* 非ブートパーティションを MBR または GPT に設定
* Sysprep 後に Systems Manager (SSM) サービスを開始。
* ENA 設定を最適化
* ジャンボフレームを有効化
* EC2Launch v2 で実行するように Sysprep を設定
* Windows イベントログを発行
など。



[最新世代のインスタンスタイプへの移行](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/WindowsGuide/migrating-latest-types.html)

以下の作業について書かれているページ。

* パート 1: AWS PV ドライバーのインストールとアップグレード
* パート 2: ENA のインストールとアップグレード
* パート 3: AWS NVMe ドライバーのアップグレード
* パート 4: EC2Config および EC2Launch の更新
* パート 5: ベアメタルインスタンスのシリアルポートドライバーのインストール
* パート 6: 電源管理設定の更新
* パート 7: 新しいインスタンスタイプ用のインテルチップセットドライバーの更新


[Amazon EC2 キーペアと Windows インスタンス](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/WindowsGuide/ec2-key-pairs.html)

Windows インスタンスでは、プライベートキーを使用して管理者パスワードを取得する。



# BlackBelt

[AWS Black Belt Online Seminar Amazon EC2 - Windows](https://www.slideshare.net/AmazonWebServicesJapan/aws-black-belt-online-seminar-amazon-ec2-windows)

* P18: キーペアを使用して管理者パスワードを取得可能
* P22: いくつかのツールがインストールされている
  * EC2Config サービス
  * 準仮想化(PV)ドライバ
  * AWS Tools for Windows PowerShell
* P25: EC2Config の役割
  * 暗号化パスワードの設定
  * ユーザーデータの実行
  * Windows のアクティベーション
  * ボリュームのフォーマット、マウント
* P26: EC2Config のプロパティ
  * [Set Computer Name]
  * [User Data]
  * [Event Log]
  * [CloudWatch Logs]
  * [Wallpaper Information]
* P27: sysprep
  * パスワードに関して設定可能(保持するなど)


