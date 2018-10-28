https://cloud.google.com/solutions/continuous-delivery-spinnaker-kubernetes-engine


##### ゾーの設定

```
$ gcloud config set compute/zone us-central1-f
```

##### コンテナ用のクラスタ作成

```
$ gcloud container clusters create spinnaker-tutorial \
    --machine-type=n1-standard-2
```

### サービスアカウント設定

##### サービスアカウントの作成

```
$ gcloud iam service-accounts create  spinnaker-storage-account \
    --display-name spinnaker-storage-account
```

##### 環境変数設定

サービスアカウントのメールアドレスと現在のプロジェクトIDを環境変数に格納。

```
$ export SA_EMAIL=$(gcloud iam service-accounts list \
    --filter="displayName:spinnaker-storage-account" \
    --format='value(email)')

$ export PROJECT=$(gcloud info --format='value(config.project)')
```

##### サービスアカウントに権限付与

storage.admin ロールをサービスアカウントにバインド。

```
$ gcloud projects add-iam-policy-binding \
    $PROJECT --role roles/storage.admin --member serviceAccount:$SA_EMAIL
```

##### サービスアカウントキーのダウンロード

カレントディレクトリに、「spinnaker-sa.json」がダウンロードされる。

```
$ gcloud iam service-accounts keys create spinnaker-sa.json --iam-account $SA_EMAIL
```


[参考]
http://bufferings.hatenablog.com/entry/2018/03/01/225028

```
$ kubectl create clusterrolebinding user-admin-binding \     
    --clusterrole=cluster-admin \
    --user=$(gcloud config get-value account)

$ kubectl create serviceaccount tiller --namespace kube-system

$ kubectl create clusterrolebinding tiller-admin-binding \
    --clusterrole=cluster-admin \
    --serviceaccount=kube-system:tiller

$ kubectl create clusterrolebinding \
    --clusterrole=cluster-admin \
    --serviceaccount=default:default spinnaker-admin
```



## Helm

helm バイナリをカレントディレクトリに配置。

```
$ wget https://storage.googleapis.com/kubernetes-helm/helm-v2.9.0-linux-amd64.tar.gz
$ tar zxvf helm-v2.9.0-linux-amd64.tar.gz
$ cp linux-amd64/helm .
```

```shell
# 初期設定。
$ ./helm init --service-account=tiller

# レポジトリのアップデート。
$ ./helm repo update

# helm server, client のバージョン確認
$ ./helm version
```

## Spinnaker

Cloud Storage のバケット作成。

```
$ export BUCKET=$PROJECT-spinnaker-config
$ gsutil mb -c regional -l us-central1 gs://$BUCKET
```

##### 環境変数設定

```
$ export SA_JSON=$(cat spinnaker-sa.json)
```

「spinnaker-config.yaml」を以下の内容で作成する。

```yaml
storageBucket: $BUCKET
gcs:
  enabled: true
  project: $PROJECT
  jsonKey: '$SA_JSON'

# Disable minio the default
minio:
  enabled: false


# Configure your Docker registries here
accounts:
- name: gcr
  address: https://gcr.io
  username: _json_key
  password: '$SA_JSON'
  email: 1234@5678.com
EOF
```


```
$ ./helm install stable/spinnaker -f spinnaker-config.yaml --timeout 600

$ export DECK_POD=$(kubectl get pods --namespace default -l "component=deck" \
    -o jsonpath="{.items[0].metadata.name}")

$ kubectl port-forward --namespace default $DECK_POD 8080:9000 >> /dev/null &
```

## Docker イメージ

wget https://gke-spinnaker.storage.googleapis.com/sample-app.tgz

tar xzfv sample-app.tgz

cd sample-app

git init
git add .
git commit -m "Initial commit"

gcloud source repos create sample-app
git config credential.helper gcloud.sh

git remote add origin https://source.developers.google.com/p/$PROJECT/r/sample-app
git push origin master

ソースコードが表示されることを確認。
https://console.cloud.google.com/code/develop/browse/sample-app/master


## タグ付け

git tag v1.0.0
git push --tags


kubectl apply -f k8s/services

sed s/PROJECT/$PROJECT/g spinnaker/pipeline-deploy.json | curl -d@- -X \
    POST --header "Content-Type: application/json" --header \
    "Accept: /" http://localhost:8080/gate/pipelines
