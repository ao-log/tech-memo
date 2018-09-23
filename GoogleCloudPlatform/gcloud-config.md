https://cloud.google.com/sdk/gcloud/reference/config/

### プロジェクト

##### 切り替え

```
$ gcloud config set project プロジェクト名
```

##### プロジェクトIDをセット

```
export PROJECT_ID=$(gcloud config get-value project)
もしくは
export PROJECT_ID=$(gcloud config list project --format "value(core.project)")
```

### リージョン

##### 設定

```
$ gcloud config set compute/zone us-central1-f
```

##### 確認

```
$ gcloud config list compute/zone
```
