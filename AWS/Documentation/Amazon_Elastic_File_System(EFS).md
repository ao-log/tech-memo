
# Document

[Amazon EFS アクセスポイントを使用してマウントする](https://docs.aws.amazon.com/ja_jp/efs/latest/ug/efs-access-points.html)

* すべてのファイルシステム要求に対してユーザーアイデンティティ (ユーザーの POSIX グループなど) を適用できる
  * 新規ファイル、ディレクトリはアクセスポイントのユーザー ID、グループ ID に設定される
  * アクセスポイントのユーザー ID を root にする場合は `ClientRootAccess` の IAM アクセス許可が必要
* ファイルシステムに対して別のルートディレクトリを適用できる
  * 当該ディレクトリの読み書き権限がないユーザーは読み書きできない
* IAM を使用できる
  * ファイルシステムポリシーにより設定。アイデンティティポリシー、ファイルシステムポリシーのどちらかで設定されていればよい
  * デフォルトの EFS ファイルシステムポリシーは、認証に IAM を使用せずマウントターゲットを使用してファイルシステムに接続できる任意の匿名クライアントにフルアクセスを許可
  * [IAM を使用してファイルシステムのデータアクセスを制御する](https://docs.aws.amazon.com/ja_jp/efs/latest/ug/iam-access-control-nfs-efs.html) も参照するとよい
  * クライアント側では以下のアクションを設定可能
    * `elasticfilesystem:ClientMount`
    * `elasticfilesystem:ClientWrite`
    * `elasticfilesystem:ClientRootAccess`
* マウントする際は以下の例のように行う
```
mount -t efs -o tls,iam,accesspoint=fsap-abcdef0123456789a fs-abc0123def456789a: /localmountpoint
```


# re:Post

[IAM および EFS アクセスポイントを使用して、特定の EC2 インスタンスにディレクトリアクセスを許可するにはどうすればよいですか?](https://repost.aws/ja/knowledge-center/efs-access-points-directory-access)

* インスタンスのセキュリティグループでアウトバウンドの 2049/tcp の疎通性が必要
* マウントする場合は次のコマンド
```
sudo mount -t nfs4 -o nfsvers=4.1,rsize=1048576,wsize=1048576,hard,timeo=600,retrans=2,noresvport file-system-id.efs.ap-southeast-2.amazonaws.com:/ /efs
```
* アクセスポイント、IAM 認証には `amazon-efs-utils` が必要
```
sudo yum install -y amazon-efs-utils
```
* アクセスポイントの作成
  * OwnUid、OwnGID、および許可が指定されている場合にのみ、アクセスポイントのルートディレクトリが自動作成される。このディレクトリがない場合はマウントに失敗する
* IAM 認証を行う場合、インスタンスロールには次のようなポリシーをアタッチする
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "elasticfilesystem:ClientMount",
        "elasticfilesystem:ClientWrite"
      ],
      "Resource": "arn:aws:elasticfilesystem:ap-southeast-2:123456789012:file-system/fs-8ce001b4",
      "Condition": {
        "ForAnyValue:StringEquals": {
          "elasticfilesystem:AccessPointArn": "arn:aws:elasticfilesystem:ap-southeast-2:123456789012:access-point/fsap-0093c87d798ae5ccb"
        }
      }
    }
  ]
}
```
* EFS 側では次のようなポリシーを設定する。リソースベースのポリシーを設定している場合、インスタンスロール側で権限がない場合は Access Denied となりマウント失敗する
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": "*"
      },
      "Action": [
        "elasticfilesystem:ClientMount",
        "elasticfilesystem:ClientWrite"
      ],
      "Resource": "arn:aws:elasticfilesystem:ap-southeast-2:123456789012:file-system/fs-8ce001b4",
      "Condition": {
        "StringEquals": {
          "aws:PrincipalArn": "arn:aws:iam::123456789012:role/App_team_role",
          "elasticfilesystem:AccessPointArn": "arn:aws:elasticfilesystem:ap-southeast-2:123456789012:access-point/fsap-0093c87d798ae5ccb"
        }
      }
    }
  ]
}
```






