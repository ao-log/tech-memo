M
Google Cloud Platform を使用する際、Web ブラウザ上からだけではなく、コマンドラインから出来ると何かと便利。Cloud SDK を使用すれば実現できます。ここでは、Compute Engine のインスタンスを作成するまでの手順を示します。

# Cloud SDK のインストール

まず、Cloud SDK を PC にインストールします。手順は Document に詳しく記載されているので、そちらをご参照ください。

[参考：クイックスタート](https://cloud.google.com/sdk/docs/quickstarts?hl=ja)

例えば、macOS の場合の流れを要約すると、下記の通りとなります。
1. SDK のダウンロード
2. tarボールを任意の場所に展開
3. 次のコマンドを実行。
```$ ./google-cloud-sdk/install.sh```
4. 次のコマンドを設定し、初期設定を実施。
```$ gloud init```

# 利用できるコンポーネント

次のツールがデフォルトで用意されています。gcloud も、もちろん入っています。
* gcloud
* bq
* gsutil

[参考：Google Cloud SDK](https://cloud.google.com/sdk/docs/overview?hl=ja)

現在のコンポーネントインストール状況を確認するには、次のコマンドを実行します。

```$ gcloud components list```

脱線しますが、例えば kubectl をインストールするには、次のコマンドを実行します。

```$ gcloud components install kubectl```


# 設定

前提としてプロジェクトを作成し、課金を有効にしておく必要があります。
gcloud コマンドのコンフィグを確認するには次のコマンドを実行します。

```
$ gcloud config list
```

##### プロジェクトの設定

もし、プロジェクトが別のものになっていた場合は、次のコマンドでセットします。

```
$ gcloud config set project プロジェクト名
```

##### ゾーンの設定

次のコマンドで利用できるゾーンの一覧を確認します。

```shell-session
$ gcloud compute zones list
NAME                    REGION                STATUS  NEXT_MAINTENANCE  TURNDOWN_DATE
asia-east1-b            asia-east1            UP
asia-east1-a            asia-east1            UP
...
```

次のコマンドでゾーンを設定します。ゾーン名は「us-central1-a」など。

```
$ gcloud config set compute/zone ゾーン名
```

# その他の雑多なメモ

##### フォーマット変更

```
$ gcloud compute instances list --format=json
```

##### 認証周り

```
# アクティブなアカウントの表示
$ gcloud auth list

# アクセストークンの表示
$ gcloud auth application-default print-access-token

# プロジェクト一覧の表示
$ gcloud projects list
```
