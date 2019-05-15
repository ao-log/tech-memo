先日の [Google Cloud Next ’19](https://cloud.withgoogle.com/next/sf) で [Cloud Run](https://cloud.google.com/run/) というサービスが発表されました。フルマネージドのサーバーレス実行環境とのこと。Slack と連携するサービスを作って便利さを実感してみましたので、紹介させていただきます。

# 今回作るサービスの概要

技術が進歩しているのに何故かせっせと働く毎日。ちょっとした癒しが欲しくなるときがあります。slack を開いて猫ちゃんの画像が表示されたら、それはささやかな幸せだと思います。

Slack には Slash Command といってスラッシュ付きのコマンドを実行できます。自作のコマンドを作ることもでき、今回は「/neko」コマンドを実行するとランダムに猫ちゃんの画像が表示されるようにします。

Slash Command を実行。

![slack_neko.png](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/48133/110cdf5b-0c14-8f38-bfa4-a5bb7725ee29.png)

猫ちゃんの画像がランダムに表示されます。癒されます！

![slash neko_reply.png](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/48133/2fe676fe-a1fb-af28-6a93-b6cd518c55ca.png)


# 処理の流れ

処理の流れを図にしてみました。

![Flow.png](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/48133/9f7a2f60-e4e4-d2d7-9e09-77b9ac9995d3.png)

猫ちゃんの画像は Cloud Storage の非公開バケットに保存しておきます。
Slack 上で「/neko」コマンドを実行すると、Cloud Run の URL にアクセスします。
Cloud Run 上ではコンテナで Flask が動作しており、Cloud SDK で Cloud Storage 上の画像をダウンロードします。
そして、Slack にアップロードします。



# Cloud Run にデプロイするまでの作業

では、ここから実際の作業について述べていこうと思います。大雑把に書くと以下の流れです。

1. Cloud Storage 上に画像をアップロード
2. Slack 上にアプリ作成
3. ローカル PC 上で動作確認
4. Cloud Run にデプロイ

順に説明していきます。

## Cloud Storage 上に画像をアップロード

前提は Google Cloud Platform のアカウントを持っていることです。

ここの説明は大雑把に済ませようと思いますが、[Cloud Storage のページ](https://console.cloud.google.com/storage/) 上でバケットを作成し、そちらに猫ちゃんの画像をアップロードします。

![CloudStorage.png](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/48133/74096fb4-fee5-ca0e-a205-e39487f677fe.png)


## Slack 上にアプリ作成

Slack 上のアプリ作成は以下の URL で行います。

https://api.slack.com/apps

やることは次の通りです。

* アプリの作成
* Slash コマンドの作成
* ファイルアップロードに必要な情報を取得

### アプリの作成

まずは、アプリの作成です。「Create an App」のアイコンをクリックします。

![create_an_app.png](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/48133/4dc36e9f-0afc-0ac5-857c-985f4678bf04.png)

次にアプリ名とワークスペースを設定します。

![create_a_slack_app.png](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/48133/b08efe22-c5a6-94e4-65cc-bf9e64bba570.png)

### Slash Command の設定

アプリの作成ができたら、さらに様々な設定や情報の取得ができるようになります。Slash Command のリンクは赤で囲んだ箇所にあります。(2019年4月15日時点)

![your_apps.png](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/48133/93e9ff16-e4d4-9e27-b5ed-d88924e1e5be.png)

あとは、画面上の指示に従って作成します。Edit Command の画面でコマンド名「/neko」を設定しつつ、Request URL を設定します。最終的には Cloud Run のエンドポイントにするのですが、今は設定を保留しておきます。

![slash_command_edit.png](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/48133/66e0c69d-5696-4dc7-e2b3-0208ddfaf6fd.png)

### OAuth & Permissions

先ほど Slash Command の節で示した赤色で囲んだリンクのすぐ下に「OAuth & Permissions」のリンクがあります。このページにある「OAuth Access Token」が必要な情報です。ファイルのアップロードに使用します。

### チャンネル

ファイルのアップロード時にチャンネルを指定します。その時の文字列はチャンネルにアクセスした時の URL の {CHANNEL_NAME} の部分の文字列になります。

```
https://xxxx.slack.com/messages/{CHANNEL_NAME}/
```

### Signing Secret

これは、Slash Command の受け手となる Web サーバが、リクエストを検証するためのものです。不特定の誰かに URL を叩かれると困るので、Slack からのリクエストのみに反応するようにしたいです。その時に Signing Secret を使用してリクエストの正当性を検証します。詳しくは [Verifying requests from Slack](https://api.slack.com/docs/verifying-requests-from-slack) のページに書かれています。


## ローカル PC 上で動作検証

ここからローカル PC 上で開発していきます。
前提は Docker が使用できること、Cloud SDK がインストールされていることです。
Cloud Storage からファイルをダウンロードするためにサービスアカウントを作成します。その秘密鍵を JSON 形式で作成し PC にダウンロードします。また、IAM でサービスアカウントに対し、ストレージオブジェクト閲覧者の権限を付与しておきます。

### ファイルの準備

まず、ファイル類を準備します。

#### Dockerfile

```Dockerfile:Dockerfile
FROM python:3.7-alpine
RUN pip install Flask \
    google-cloud-storage
COPY src /src
WORKDIR /src
CMD ["python", "slack-neko-uploader.py", "--host=0.0.0.0", "--port=8080"]
```

#### Flask 用の Python コード

src ディレクトリを作成し、Python コードを配置します。トークン類は直書きせず環境変数で渡すようにします。

```python:src/slack-neko-uploader.py 
import os
import random
import requests
import hashlib
import hmac
from google.cloud import storage
from flask import Flask, request

BUCKET_NAME = os.environ['BUCKET_NAME']
SLACK_SIGNING_SECRET = os.environ['SLACK_SIGNING_SECRET']
SLACK_OAUTH_ACCESS_TOKEN = os.environ['SLACK_OAUTH_ACCESS_TOKEN']
SLACK_CHANNEL = os.environ['SLACK_CHANNEL']
SLACK_UPLOAD_URL = 'https://slack.com/api/files.upload'

app = Flask(__name__)

def verify(request):
    slack_secret = bytes(SLACK_SIGNING_SECRET, 'utf-8')
    timestamp = request.headers['X-Slack-Request-Timestamp']
    request_data = request.get_data().decode('utf-8')
    base_string = f"v0:{timestamp}:{request_data}".encode('utf-8')

    my_signature = 'v0=' + hmac.new(slack_secret, base_string, hashlib.sha256).hexdigest()    

    if hmac.compare_digest(my_signature, request.headers['X-Slack-Signature']):
        return True
    else:
        return False

def download_cat_image():
    storage_client = storage.Client()
    bucket = storage_client.get_bucket(BUCKET_NAME)
    blobs = [blob for blob in bucket.list_blobs()]
    source_blob_name = blobs[int(len(blobs) * random.random())].name

    bucket.blob(source_blob_name).download_to_filename('/tmp/cat.png')

def upload_file_to_slack():
    files = {'file': open('/tmp/cat.png', 'rb')}
    param = {
        'token': SLACK_OAUTH_ACCESS_TOKEN,
        'channels': SLACK_CHANNEL,
        'filename': 'cat.png',
        'title': 'cat'
    }

    requests.post(url=SLACK_UPLOAD_URL, params=param, files=files)
    os.remove('/tmp/cat.png')

@app.route('/slack/events', methods=['POST'])
def main():
    if verify(request):
        download_cat_image()
        upload_file_to_slack()
        return 'ok'
    else:
        return ('Request failed verification.', 401)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port='8080')
```

### Docker build

ここまでできたら Docker イメージをビルドします。

```
$ docker build -t slack-neko-bot:v1.0 .
```

### 動作確認

#### 環境変数

docker run でコンテナを起動する際に渡す環境変数は一つのファイルに書いておきます。

```envfile.txt
GOOGLE_APPLICATION_CREDENTIALS=サービスアカウントの秘密鍵のJSONファイルパス
BUCKET_NAME=バケット名
SLACK_SIGNING_SECRET=Signing Secret
SLACK_OAUTH_ACCESS_TOKEN=OAuth Access Token
SLACK_CHANNEL=チャンネル
```

docker run の引数で秘密鍵ファイルのパスを使うので設定しておきます。

```
$ export GOOGLE_APPLICATION_CREDENTIALS=サービスアカウントの秘密鍵のJSONファイルパス
```

#### docker run

環境変数の準備ができたら、docker run でコンテナを起動します。

```shell
$ docker run \
    -p 8080:8080 \
    --rm \
    --env-file envfile.txt \
    -e GOOGLE_APPLICATION_CREDENTIALS=/tmp/keys/credentials.json \
    -v `pwd`/$GOOGLE_APPLICATION_CREDENTIALS:/tmp/keys/credentials.json:ro \
    slack-neko-bot:v1.0
```

動作確認には [ngrok](https://ngrok.com/) というサービスが便利です。外部からのアクセスをローカル PC の指定ポート宛に転送できます。生成された URL の文字列に「/slack/events」を付与して Slack の Slash コマンド管理画面で設定すれば、ローカル PC に転送されて動作確認できます。Slack のチャンネル上で「/neko」コマンドを実行し、猫ちゃんの画像が表示されたら成功です！

## Cloud Run にデプロイ

Cloud SDK のコンポーネントを追加するなど準備が必要です。次の URL に記載されている作業を行なってください。

https://cloud.google.com/run/docs/setup?hl=ja

今後の作業のためにプロジェクトIDを環境変数に格納しておきます。

```
$ export PROJECT_ID=$(gcloud config get-value project)
```

次のコマンドで Cloud Build によりイメージをビルドしつつ、Container Registry に登録します。

```
$ gcloud builds submit --tag gcr.io/$PROJECT_ID/slack-neko-bot:v1.0
```

Cloud Run で渡す環境変数を設定しておきます。

```
export BUCKET_NAME=バケット名
export SLACK_SIGNING_SECRET=Signing Secret
export SLACK_OAUTH_ACCESS_TOKEN=OAuth Access Token
export SLACK_CHANNEL=チャンネル
```

次のコマンドで Cloud Run にデプロイします。環境変数もここで渡します。

**追記**　ここではシークレット情報を環境変数で渡していますが、**セキュリティ上好ましくありません**。Cloud KMS に格納し、復号して取り出すのが良いです。Cloud Run の管理コンソールからは --update-env-vars で渡した環境変数の文字列が丸見えになるなど、安全上の問題があります。

```
$ gcloud beta run deploy slack-neko-bot \
    --image gcr.io/$PROJECT_ID/slack-neko-bot:v1.0 \
    --update-env-vars BUCKET_NAME=$BUCKET_NAME,SLACK_SIGNING_SECRET=$SLACK_SIGNING_SECRET,SLACK_OAUTH_ACCESS_TOKEN=$SLACK_OAUTH_ACCESS_TOKEN,SLACK_CHANNEL=$SLACK_CHANNEL
```

うまくいけば HTTPS の URL が生成されます。

```
Service [slack-neko-bot] revision [slack-neko-bot-xxxxxxxx] has been deployed
and is serving traffic at https://slack-neko-bot-xxxxxxxx.a.run.app
```

これだけだとコンテナから GCP 上のサービスへのアクセス権がないので Cloud Storage からファイルをダウンロードできません。最後の仕上げとして、権限を付与します。Cloud Console からサービスのチェックボックスを選択した状態で右上の「情報パネルを表示」をクリックすると権限の編集画面になります。そこで、Cloud Run サービス エージェントのサービスアカウントに閲覧者の権限を付与します。ただ、権限が強すぎるのでもっと絞りたかったのですが、なぜかストレージ閲覧者を選択できなかったので、ひとまず閲覧者で断念しています・・

![add_role.png](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/48133/9674ae0f-2880-25de-9d67-2a3cf72828a3.png)


# 最後に

Docker での開発に慣れている人であれば、ローカル PC で検証した後にイメージをデプロイすれば良いだけなので楽だと感じました。コンテナなので、可搬性の高さを享受できるのが良い点だと思いました。

また、Kubernetes のようにマニフェストファイルを書く手間もかからないので、ステートレスなサービスや API Gateway のようなものを動かす環境として、とても楽だと実感できました。

# 参考

##### Cloud Run 関連

* [Google Cloud Next ’19 が始まりました](https://cloudplatform-jp.googleblog.com/2019/04/next19-recap-day1.html)
* [Google Cloud Next '19で発表された新機能を紹介します！　(Cloud Run, BigQuery Storage API, Cloud Data Fusion)](https://techblog.zozo.com/entry/google_cloud_next_19)
* [Qiita: Cloud Runで動かすGo + Echo Framework](https://qiita.com/nkumag/items/90a24c87330bfb503e29)

##### Cloud Run Documentation

* [Cloud Run documentation](https://cloud.google.com/run/docs/)
* [gcloud beta run](https://cloud.google.com/sdk/gcloud/reference/beta/run/?hl=ja)

##### Slack 関連

* [Slash Commands](https://api.slack.com/slash-commands)
* [files.upload](https://api.slack.com/methods/files.upload)
* [Verifying requests from Slack](https://api.slack.com/docs/verifying-requests-from-slack)
* [Qiita: プログラムからSlackに画像投稿する方法まとめ](https://qiita.com/stkdev/items/992921572eefc7de4ad8)
* [Verifying Slack Slash Commands in Google Cloud Functions](https://ratil.life/verifying-slack-slash-commands-in-google-cloud-functions/)

##### 開発ツール

* [ngrok](https://ngrok.com/)
* [Qiita: ngrokが便利すぎる](https://qiita.com/mininobu/items/b45dbc70faedf30f484e)
