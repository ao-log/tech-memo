
# サービスアカウントに関するオペレーション

### 作成

my-sa-123 という名前で作成する例。

```
$gcloud iam service-accounts create my-sa-123 \
    --display-name "my service account"
```

### サービスアカウントに権限付与

```
gcloud projects add-iam-policy-binding $DEVSHELL_PROJECT_ID \
    --member serviceAccount:my-sa-123@$DEVSHELL_PROJECT_ID.iam.gserviceaccount.com \
    --role roles/editor
```

### 明示的にサービスアカウントを指定

```
$ export GOOGLE_APPLICATION_CREDENTIALS=JSONへのパス
```

アカウントのアクティベートが必要

```
$ gcloud auth activate-service-account \      
    storage-admin@aoao22pj.iam.gserviceaccount.com \
    --key-file=JSONへのパス
```

アカウントの確認。

```
$ gcloud auth list
```

### VM へのサービスアカウント設定

VM インスタンス作成時に、その VM のサービスアカウントを指定することもできる。
