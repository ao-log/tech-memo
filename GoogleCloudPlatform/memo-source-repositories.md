
## 料金

https://cloud.google.com/source-repositories/pricing?hl=en

## GCE 上で Source Repositories を使えるようにする。

### 事前準備

#### リポジトリ作成

リポジトリ作成。

```
$ gcloud source repos create gce-dev
```

#### サービスアカウントの作成

プロジェクトIDをセットしておく。

```
$ export PROJECT_ID=$(gcloud config get-value project)
```

サービスアカウントを作成。

```
$ gcloud iam service-accounts \
    create gce-service-account \
    --display-name "gce-service-account"
```

サービスアカウントに source.writer 権限を付与。

```
$ gcloud projects add-iam-policy-binding ${PROJECT_ID} \
    --member serviceAccount:gce-service-account@${PROJECT_ID}.iam.gserviceaccount.com \
    --role roles/source.writer
```

[参考]

* [サービスアカウントへの役割の付与](https://cloud.google.com/iam/docs/granting-roles-to-service-accounts?hl=ja)
* [役割について](https://cloud.google.com/iam/docs/understanding-roles?hl=ja)

#### VMインスタンスの作成

```
$ gcloud compute instances create my-centos \
  --image-family centos-7 \
  --image-project centos-cloud \
  --machine-type f1-micro \
  --preemptible \
  --service-account gce-service-account@${PROJECT_ID}.iam.gserviceaccount.com \
  --scopes cloud-source-repos
```

[参考]

* [インスタンスのサービス アカウントの作成と有効化](https://cloud.google.com/compute/docs/access/create-enable-service-accounts-for-instances?hl=ja)


#### VM インスタンスに接続

```
$ gcloud compute ssh my-centos
```

#### Git インストール

```
$ sudo yum install -y git
```

#### git config

```
$ git config --global credential.'https://source.developers.google.com'.helper gcloud.sh
```

[参考:リポジトリをリモートとして追加する](https://cloud.google.com/source-repositories/docs/adding-repositories-as-remotes?hl=ja)


#### リポジトリのクローン

```
$ gcloud source repos clone gce-dev
```

#### インスタンスの削除

終了後は OS シャットダウン後に、インスタンスを削除します。

```
$ gcloud compute instances delete my-centos
```