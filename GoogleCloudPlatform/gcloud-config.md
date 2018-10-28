https://cloud.google.com/sdk/gcloud/reference/config/

### プロジェクト

##### 切り替え

```
$ gcloud config set project プロジェクト名
```


##### プロジェクトの一覧

```
$ gcloud projects list
```

##### プロジェクトIDをセット

```
export PROJECT_ID=$(gcloud config get-value project)
もしくは
export PROJECT_ID=$(gcloud config list project --format "value(core.project)")
```

### config 設定を丸ごと切り替え

```shell
# 作成
$ gcloud config configurations create コンフィグ名

# 切り替え
$ gcloud config configurations activate コンフィグ名

# 一覧
$ gcloud config configurations list
```

https://cloud.google.com/sdk/gcloud/reference/config/configurations/

### リージョン

##### 設定

```
$ gcloud config set compute/zone us-central1-f
```

##### 確認

```
$ gcloud config list compute/zone
```
