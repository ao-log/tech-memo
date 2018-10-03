
### 機能

* 権限の設定は「service.resource.verb」の形式。
* 役割
  * 基本の役割: オーナー、編集者、閲覧者の三つ
  * 事前定義された役割
  * カスタムの役割
* 階層構造は、上から順に組織、フォルダ、プロジェクト、リソースとなっている。

### gcloud

##### IAM ポリシーの一覧を取得

```
gcloud projects get-iam-policy PROJECT-ID
# json 形式
gcloud projects get-iam-policy PROJECT-ID --format json
```

##### IAM ポリシーを設定

```
gcloud projects set-iam-policy PROJECT-ID iam.json
```

##### ロールを設定

```
gcloud projects add-iam-policy-binding PROJECT-ID \
      --member user:email3@gmail.com --role roles/editor
```

### 参考

[概要](https://cloud.google.com/iam/docs/overview?hl=ja)
