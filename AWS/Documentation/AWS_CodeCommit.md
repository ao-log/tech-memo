
# CodeCommit

[AWS CLI 認証情報ヘルパーを使用した、Linux、macOS、または UNIX での AWS CodeCommit リポジトリへの HTTPS 接続のセットアップ手順](https://docs.aws.amazon.com/ja_jp/codecommit/latest/userguide/setting-up-https-unixes.html)

次のコマンドを実行することで設定する。
```
git config --global credential.helper '!aws codecommit credential-helper $@'
git config --global credential.UseHttpPath true
```

~/.gitconfig に以下の設定を行われる。
```
[credential]    
    helper = !aws --profile CodeCommitProfile codecommit credential-helper $@
    UseHttpPath = true
```

この状態で IAM のクレデンシャル情報で git の各種操作ができる。


## リポジトリを操作する

[AWS CodeCommit リポジトリイベントの通知を設定する](https://docs.aws.amazon.com/ja_jp/codecommit/latest/userguide/how-to-repository-email.html)

プルリクエストなどのイベントを Amazon SNS トピックに送信するよう設定できる。


[ロールを使用して AWS CodeCommit リポジトリへのクロスアカウントアクセスを設定する](https://docs.aws.amazon.com/ja_jp/codecommit/latest/userguide/cross-account.html)

* アカウント A: リポジトリを所有するアカウント
  * IAM ロールを作成。信頼ポリシーでアカウント B を許可。CodeCommit へのアクセス許可を設定したポリシーをアタッチ
* アカウント B: リポジトリを利用するアカウント
  * アクション sts:AssumeRole、リソース アカウント A の IAM ロールを許可

アカウント B の利用者は以下の手順で対応

~/.aws/config を設定
```
[profile MyCrossAccountAccessProfile]
region = us-east-2
account = 111122223333
role_arn = arn:aws:iam::111122223333:role/MyCrossAccountRepositoryContributorRole
source_profile = default
output = json
```

リポジトリを clone
```
git clone codecommit://MyCrossAccountAccessProfile@MySharedDemoRepo
```



# BlackBelt

[AWS Black Belt Online Seminar AWS CodeCommit & AWS CodeArtifact](https://pages.awscloud.com/rs/112-TZM-766/images/20201020_BlackBelt_AWS_CodeCommit_AWS_CodeArtifact.pdf)

* CodeCommit
  * Git オブジェクトは S3、Git インデックスは DynamoDB、暗号化キーは KMS で管理されている
  * Code Guru Reviewer for Java とも連携可能
  * 接続方法
    * IAM の Git 認証情報で HTTPS。IAM でユーザー名、パスワードを生成する。git clone 時にユーザー名、パスワードを入力
    * IAM と SSH キーを紐付け
    * git-remote-codecommit
  * プルリクエスト
    * マージ時に満たさなければならない承認ルールを作成できる
    * IAM の Condition で開発チーム以外はプルリクを main ブランチにマージさせないような制御が可能
* CodeArtifact
  * ドメイン - リポジトリの構造
  * aws codeartifact login ... コマンドでログイン。以降、そのリポジトリを使用するが 12 時間で token 切れとなる



# 参考

* Document
  * [AWS CodeCommit とは](https://docs.aws.amazon.com/ja_jp/codecommit/latest/userguide/welcome.html)
* サービス紹介ページ
  * [AWS CodeCommit](https://aws.amazon.com/jp/codecommit/)
  * [よくある質問](https://aws.amazon.com/jp/codecommit/faqs/)
* [ナレッジセンター](https://aws.amazon.com/jp/premiumsupport/knowledge-center/#AWS_CodeCommit)
* Black Belt
  * [AWS Black Belt Online Seminar AWS CodeCommit & AWS CodeArtifact](https://pages.awscloud.com/rs/112-TZM-766/images/20201020_BlackBelt_AWS_CodeCommit_AWS_CodeArtifact.pdf)

