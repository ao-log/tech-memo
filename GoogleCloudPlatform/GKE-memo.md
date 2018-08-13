

GKE チュートリアルを読んでいて気になった箇所のメモです。



[コンテナ化されたウェブ アプリケーションのデプロイ](https://cloud.google.com/kubernetes-engine/docs/tutorials/hello-app?hl=ja)


##### タグをつけつつイメージのビルド

タグ名は下記のような命名にする必要があります。

```
$ docker build -t gcr.io/${PROJECT_ID}/hello-app:v1 .
```

##### イメージをコンテナレジストリにアップロード

```
$ gcloud docker -- push gcr.io/${PROJECT_ID}/hello-app:v1
```

##### コンテナクラスタの作成

```
$ gcloud container clusters create hello-cluster --num-nodes=3
```

**memo**

クラスタ認証情報の取得は以下のコマンドで行います。これは、上記コマンドでコンテナクラスタを作成している場合には不要です。
```
$ gcloud container clusters get-credentials hello-cluster
```

##### アプリケーションのデプロイ

この方法だと、Pod は一つのみ。

```
$ kubectl run hello-web --image=gcr.io/${PROJECT_ID}/hello-app:v1 --port 8080
```

##### Service の作成

--type で LoadBalancer にしてますが、これは Cloud Load Balancer に対応。--port はロードバランサ用のポート番号。Pod へのトラフィックはラベル名　hello-web の Pod に対して行い、--target-port で指定したポート番号にトラフィックを転送します。この設定により、ワールドワイドに公開されます。

```
$ kubectl expose deployment hello-web --type=LoadBalancer --port 80 --target-port 8080
```

##### スケールアウト

```
$ kubectl scale deployment hello-web --replicas=3
```

##### ローリングアップデート

```
$ docker build -t gcr.io/${PROJECT_ID}/hello-app:v2 .
$ gcloud docker -- push gcr.io/${PROJECT_ID}/hello-app:v2
```

イメージの更新。

```
$ kubectl set image deployment/hello-web hello-web=gcr.io/${PROJECT_ID}/hello-app:v2
```

##### 掃除

```shell
# サービスの停止
$ kubectl delete service hello-web

# クラスタの停止
$ gcloud container clusters delete hello-cluster
```

# リファレンス

[gcloud container](https://cloud.google.com/sdk/gcloud/reference/container/?hl=ja)
