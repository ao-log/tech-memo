
# AWS IAM

## 開始方法

[最初の IAM 管理者のユーザーおよびグループの作成](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/getting-started_create-admin-group.html)

```shell
# グループの作成
aws iam create-group --group-name Admins

# グループの確認
aws iam list-groups

# ポリシーをグループにアタッチ
aws iam attach-group-policy --group-name Admins --policy-arn arn:aws:iam::aws:policy/AdministratorAccess

# アタッチされているポリシーを確認
aws iam list-attached-group-policies --group-name Admins
```



## チュートリアル

[IAM チュートリアル: タグに基づいて AWS リソースにアクセスするためのアクセス許可を定義する](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/tutorial_attribute-based-access-control.html)

属性ベースのアクセスコントロール (ABAC) は、属性に基づいてアクセス許可を定義する認証戦略のこと。
AWS ではタグによるアクセスコントロールが可能。

ポリシーの Condition 句の例。
```json
            "Condition": {
                "StringEquals": {
                    "iam:ResourceTag/access-project": "${aws:PrincipalTag/access-project}",
                    "iam:ResourceTag/access-team": "${aws:PrincipalTag/access-team}",
                    "iam:ResourceTag/cost-center": "${aws:PrincipalTag/cost-center}"
                }
            }
```



## ID

### ユーザー

[ユーザー](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/id_users.html)

IAM ユーザーとしてアクセスする方法は以下の通り。

* コンソールパスワード
* アクセスキー
* CodeCommit で使用するアクセスキー
* サーバ証明書


[パスワードの管理](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/id_credentials_passwords.html)

* パスワードポリシーを設定可能
* ユーザーが自分のパスワードを変更できるようにするには「iam:ChangePassword」アクションの許可が必要。


[IAM ユーザーのアクセスキーの管理](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/id_credentials_access-keys.html)

* シークレットアクセスキーはアクセスキーの発行時のみ表示可能
* **アクセスキーは定期的に更新することを推奨。元のアクセスキーはいきなり削除せず無効化しておくことを推奨。**


[AWS での多要素認証 (MFA) の使用](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/id_credentials_mfa.html)



### グループ

[IAM グループ](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/id_groups.html)



### ロール

[IAM ロール](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/id_roles.html)

* ユーザーは 1 人の特定の人に一意に関連付ける。一方でロールはそれを必要とする任意の人が引き受けるようになっている。
* ロールには標準の長期認証情報 (パスワードやアクセスキーなど) が関連付けられない。代わりに、ロールを引き受けると、ロールセッション用の一時的なセキュリティ認証情報が提供される。
* アクセス権の委任。例えば別の AWS アカウントに対してアクセスの許可が可能。


[用語と概念](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/id_roles_terms-and-concepts.html)

* AWS サービスロール: AWS サービスが使用するロール
* AWS サービスにリンクされたロール: AWS サービスに直接リンクされた一意のタイプのサービスロール
* 委任: 例えば別の AWS アカウントに対してアクセスの許可が可能。アクセス許可ポリシー（どのアクションを許可するか）、信頼ポリシーを設定する（誰に割り当てるか）。


[一般的なシナリオ](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/id_roles_common-scenarios.html)

* ロールを切り替えるアクセス許可を IAM ユーザーに付与する事が可能
* IAM ユーザーはコンソール上からロールを切替可能
* ロールを切り替える際は AssumeRole の API を実行している


[アクセス権を委任するポリシーの例](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/id_roles_create_policy-examples.html)

チュートリアルがある。[IAM チュートリアル: AWS アカウント間の IAM ロールを使用したアクセスの委任](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/tutorial_cross-account-with-roles.html)

開発用アカウントから本番アカウントの S3 バケットへのアクセスをできるようにしたい場合。
本番アカウントで IAM ロールを作成する。ロール作成時に別の AWS アカウントを選択しておく(**信頼ポリシーで当ロールを誰が引き受けられるかを設定**)。
また、開発アカウントでは以下のように **AssumeRole のポリシーを使用したいユーザー、グループにアタッチする。**
```json
{
  "Version": "2012-10-17",
  "Statement": {
    "Effect": "Allow",
    "Action": "sts:AssumeRole",
    "Resource": "arn:aws:iam::PRODUCTION-ACCOUNT-ID:role/UpdateApp"
  }
}
```

[BlackBelt の資料](https://www.slideshare.net/AmazonWebServicesJapan/20190130-aws-black-belt-online-seminar-aws-identity-and-access-management-aws-iam-part2/20) が分かりやすい。


[Amazon EC2 インスタンスで実行されるアプリケーションに IAM ロールを使用してアクセス許可を付与する](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/id_roles_use_switch-role-ec2.html)

* インスタンスプロファイルは、インスタンスで実行されるアプリケーションにロールの一時的な認証情報を提供可能。
* アプリケーションは実行時に Amazon EC2 インスタンスメタデータからセキュリティ認証情報を取得。これらは、ロールを表す一時的セキュリティ認証情報で制限された期間の間有効。
* IAM ユーザーが EC2 インスタンスにロールを割り当てるとき、**iam:PassRole を設定しておくと許可したロールのみしか割当できなくなる。**（例えば、IAMロールの権限が大きい場合にそのロールの割当が許可されていると、実質的にはユーザにはロールの権限の操作が可能になってしまうため）



### 一時的な認証情報

[IAM の一時的なセキュリティ認証情報](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/id_credentials_temp.html)

* 一時的セキュリティ認証情報は、ユーザーのリクエストに応じて動的に生成され提供される。
* エンドポイントは https://sts.amazonaws.com。

用途

* ID フェデレーション: AWS 外部の ID に対して AWS リソースへのアクセス権を付与可能。
* クロスアカウントでの委任が可能。
* EC2 インスタンスのロール。


[AWS リソースを使用した一時的な認証情報の使用](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/id_credentials_temp_use-resources.html)

```shell
# アクセスキー ID、シークレットアクセスキー、セッショントークンを抽出
aws sts assume-role --role-arn arn:aws:iam::123456789012:role/role-name --role-session-name "RoleSession1" --profile IAM-user-name > assume-role-output.txt

# ロールの権限でコマンドを実行
export AWS_ACCESS_KEY_ID=AKIAI44QH8DHBEXAMPLE
export AWS_SECRET_ACCESS_KEY=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
export AWS_SESSION_TOKEN=AQoDYXdzEJr...<remainder of security token>
aws ec2 describe-instances --region us-west-1
```



## アクセス管理

[AWS リソースのアクセス管理](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/access.html)

**アクセスの許可、拒否は以下のルールに従う。**

* デフォルトでは、すべてのリクエストが明示的に拒否
* 明示的な許可が含まれている場合、デフォルト設定は上書きされる
* 明示的な拒否がある場合、そちらが優先される。
* アクセス許可の境界、Organizations SCP、セッションポリシーがある場合、明示的な拒否で上書きされる場合がある。


[IAM でのポリシーとアクセス許可](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/access_policies.html)

6 つのポリシータイプがある。

* **アイデンティティベースのポリシー**: IAM エンティティにアタッチされるポリシー(AWS 管理ポリシー、カスタマー管理ポリシー、インライン管理ポリシー)
* **リソースベースのポリシー**: バケットポリシー、IAM ロールの信頼ポリシー
* アクセス許可の境界: エンティティに付与できるアクセス許可の上限を定義したもの
* Organizations SCP: 組織または組織単位(OU) のメンバーアカウントのアクセス許可の上限を定義したもの
* ACL: S3 などで使用されるアクセスコントロールポリシー
* セッションポリシー: AssumeRole などでロールセッションを作成し、プログラムに渡すもの



### ポリシーとアクセス許可

[管理ポリシーとインラインポリシー](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/access_policies_managed-vs-inline.html)

ポリシーはバージョニングされている。最大 5 つのバージョンが保存されるようになっている。


[IAM エンティティのアクセス許可の境界](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/access_policies_boundaries.html)

ユーザーやロールのアクセス許可の上限を設定するもの。アクセスの許可を設定するものではない。


[アイデンティティベースのポリシーおよびリソースベースのポリシー](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/access_policies_identity-vs-resource.html)

同一の AWS アカウント内の場合、アイデンティティベース、リソースベースどちらか片方のポリシーで Allow が設定されている場合、そのリクエストは許可される。ただし、明示的な拒否がある場合は拒否が優先される。

一方で、クロスアカウントの場合は、両方のポリシーで Allow しないと許可されない。


[ポリシーの例](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/access_policies_examples.html)

以下のような制御が可能。

* 特定の日付範囲内のみ許可
* 特定リージョンに対する操作のみ許可
* 特定の送信元 IP アドレス以外は拒否
* 特定のリソースタグが付与されたリソースに対する操作のみ許可
* MFA 認証している場合のみ操作を許可
* 特定のタグが付与されたロールのみ引き受けることが可能



### IAM ポリシーの管理

[IAM ポリシーの作成](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/access_policies_create.html)

* カスタマー管理ポリシーを作成可能。AWS 管理ポリシーは編集不可。
* インラインポリシーの使用は推奨されない。
* JSON 形式で直打ちするだけでなく、ビジュアルエディターによる編集も可能。


[IAM Policy Simulator を使用した IAM ポリシーのテスト](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/access_policies_testing-policies.html)

IAM Policy Simulator により IAM ポリシーのテストが可能。


[IAM ポリシーのバージョニング](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/access_policies_managed-versioning.html)

カスタマー管理ポリシーでは、最大 5 つのバージョンが保存される。
複数バージョンがある場合、実際に使用されるのはデフォルトバージョンとして指定したバージョンとなる。



## セキュリティ

[IAM でのセキュリティのベストプラクティス](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/best-practices.html)

上記 URL より一部を印象する。

* AWS アカウントのルートユーザーのアクセスキーをロックする
* 個々の IAM ユーザーを作成する
* IAM グループの使用
* 最小限の特権
* インラインポリシーではなくカスタマー管理ポリシーを使用する
* 強度の高いパスワードポリシーを設定
* MFA の有効化
* Amazon EC2 インスタンスで実行するアプリケーションに対し、ロールを使用する
* ロールを使用してアクセス許可を委任する
* アクセスキーを共有しない
* 認証情報を定期的にローテーションする。
* 不要な認証情報の削除
* AWS アカウントのアクティビティの監視



## Access Analyzer

[AWS IAM Access Analyzer を使用する](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/what-is-access-analyzer.html)

対象リソースのポリシーを分析して信頼ゾーン(分析する範囲。自アカウントもしくは Organization)の外部からのアクセスが可能になっているかどうかを分析できるもの。
（一方でアクセスアドバイザーは過去のアクセス履歴を取得するもの）


[Access Analyzer リソースタイプ](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/access-analyzer-resources.html)

サポートされているリソースタイプは以下のとおり。

* Amazon Simple Storage Service バケット
* AWS Identity and Access Management ロール
* AWS Key Management Service キー
* AWS Lambda の関数とレイヤー
* Amazon Simple Queue Service キュー



## トラブルシューティング

[IAM のトラブルシューティング](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/troubleshoot.html)

[一般的な IAM の問題のトラブルシューティング](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/troubleshoot_general.html)

* アクセスキーを紛失した場合はリセット可能
* IAM はグローバルリソースで結果整合性なので、変更がすぐ反映されない場合がある


[IAM ロールのトラブルシューティング](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/troubleshoot_roles.html)

* ロールを引き受けることができない場合
  * ロール使用側: sts:AssumeRole が許可されており、引き受けるロールがリソースに設定されているか
  * ロールの提供元: 信頼ポリシーで許可されているかどうか
* iam:PassRole の権限がない場合
  * リンクされたサービスロールを作成する場合、サービスにロールを渡す権限が必要になる。iam:PassRole を許可していない場合はエラーとなる。


[IAM および Amazon EC2 のトラブルシューティング](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/troubleshoot_iam-ec2.html)

* 一時的な認証情報を取得できない場合、IAM ロールが設定されているかどうかを確認する。また、インスタンスメタデータサービス(IMDS) にアクセス可能化を確認する



## リファレンス

[AWS Identity and Access Management のリファレンス情報](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/reference.html)


[IAM 識別子](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/reference_identifiers.html)

* フレンドリ名: IAM ユーザー作成時などに付与する名前。
* ARN: 「arn:partition:service:region:account:resource」の形式。
* 一意の識別子: 「AIDAJQABLZS4A3QDU576Q」のような形式。一位に識別可能。同じ IAM ユーザーのフレンドリ名を再作成した場合でも識別子は別のものになる。IAM ユーザーは「AIDA」、IAM ロールは「AROA」で始まる。


[IAM JSON ポリシーのリファレンス](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/reference_policies.html)

JSON ポリシーのリファレンス。以下の要素がある。

* Version
* Id
* Statement
* Sid: 
* Effect: Allow or Deny
* Principal: ロール用の信頼ポリシー、リソースベースのポリシーで使用可能。アイデンティティベースのポリシーでは使用できない。
* NotPrincipal
* Action
* NotAction
* Resource
* NotResource
* Condition


[IAM JSON ポリシーの要素: 条件演算子](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/reference_policies_elements_condition_operators.html)

Condition で使用可能な条件演算子。StringEquals、StringLike など。


[複数のキーまたは値による条件の作成](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/reference_policies_multi-value-conditions.html)

AND と OR は以下の通り。

* 複数の条件演算子がある場合: AND
* 一つの条件演算子内に複数のキーがある場合: AND
* 一つのキーに複数の値がある場合: OR

[IAM ポリシーの要素: 変数とタグ](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/reference_policies_variables.html)

次の例のようにポリシー変数を使用することが可能。

```
      "Resource": ["arn:aws:s3:::mybucket/${aws:username}/*"]
```

aws:SourceIp、aws:SecureTransport、aws:userid、aws:username などがある。


[ポリシーの評価論理](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/reference_policies_evaluation-logic.html)

ポリシーの評価順。

* デフォルトでは、すべてのリクエストが明示的に拒否
* 明示的な許可が含まれている場合、デフォルト設定は上書きされる
* 明示的な拒否がある場合、そちらが優先される。
* アクセス許可の境界、Organizations SCP、セッションポリシーがある場合、明示的な拒否で上書きされる場合がある。

アイデンティティベース、リソースベースのポリシーがある場合。

* 同一の AWS アカウント内の場合、アイデンティティベース、リソースベースどちらか片方のポリシーで Allow が設定されている場合、そのリクエストは許可される。ただし、明示的な拒否がある場合は拒否が優先される。
* 一方で、クロスアカウントの場合は、両方のポリシーで Allow しないと許可されない。


[AWS グローバル条件コンテキストキー](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/reference_policies_condition-keys.html)

aws:SourceIp、aws:SourceVpc、aws:SourceVpce、aws:userid などのキーを Condition で使用可能。


[AWS のサービスのアクション、リソース、および条件キー](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/reference_policies_actions-resources-contextkeys.html)

ここで各サービスごとのアクションで指定できるリソースタイプ、条件キーを確認できる。

* リソースタイプ列: 列が空の場合、リソースレベルのアクセスを許可していない
* 条件キー列: Condition で指定可能な条件


# 参考

* [IAM とは](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/introduction.html)
* [AWS Identity and Access Management (IAM)](https://aws.amazon.com/jp/iam/)
* [よくある質問](https://aws.amazon.com/jp/iam/faqs/)
* [ナレッジセンター](https://aws.amazon.com/jp/premiumsupport/knowledge-center/#AWS_Identity_and_Access_Management)
* Black Belt
  * [20190129 AWS Black Belt Online Seminar AWS Identity and Access Management (AWS IAM) Part1](https://www.slideshare.net/AmazonWebServicesJapan/20190129-aws-black-belt-online-seminar-aws-identity-and-access-management-iam-part1)
  * [20190130 AWS Black Belt Online Seminar AWS Identity and Access Management (AWS IAM) Part2](https://www.slideshare.net/AmazonWebServicesJapan/20190130-aws-black-belt-online-seminar-aws-identity-and-access-management-aws-iam-part2)
* [AWSマルチアカウントにおけるIAMユーザー設計戦略を考えてみる](https://iselegant.hatenablog.com/entry/2020/05/24/215808)
* [IAM ロールの PassRole と AssumeRole をもう二度と忘れないために絵を描いてみた](https://dev.classmethod.jp/articles/iam-role-passrole-assumerole/)
