少し前に Python 3.7 が GAE の standard environment で使えるようになったとアナウンスされました（2018/10/14 時点ではベータ版）。また、second generation runtime が使えるようになりました。

[Introducing App Engine Second Generation runtimes and Python 3.7](https://cloud.google.com/blog/products/gcp/introducing-app-engine-second-generation-runtimes-and-python-3-7)

本件に関して、既に Qiita やその他のブログで記事を執筆されている方もいらっしゃいます。それらを拝見して、私が昔書いたアプリ （内部で　pandas を使用しています） を second generation 上で簡単に動かせるのでは？ と思って試してみたのがこの記事を書こうと思ったきっかけです。

second generation runtime になって、制約が緩くなったり新しくできることが増えています。特にサードパーティーの Python パッケージを追加できるのが大きいです。今まで動かせなかったようなアプリも動かせるようになっているのでは、と思います。

# second generation runtime について

まず、ざっくりと second generation runtime の概要を説明します。

second generation のランタイムでは、gVisor ベースでコンテナのアイソレーションが実現されています。そのため、first generation と比べて制限が緩くなっています。

* 任意の Python パッケージを追加可能
* /tmp 上へのファイル読み書きが可能

この他にも様々な制限の緩和、非推奨となった機能などがあります。詳しくは公式ドキュメントの方にまとまってますので、そちらをご参照ください。

[Understanding differences between Python 2 and Python 3 on the App Engine standard environment](https://cloud.google.com/appengine/docs/standard/python3/python-differences)



# Web アプリケーションについて

今回、GAE にデプロイするのはExcel もしくは CSV ファイルの表をマークダウン形式に変換する Web アプリケーションです。

![overview.png](https://qiita-image-store.s3.amazonaws.com/0/48133/bb24ae97-dce5-abdd-ef4e-83884552d6e3.png)

second generation runtime の機能のうち役立つのは次の 2 点。

###### サードパーティーのパッケージをインストール可能

今回は以下のパッケージを使用します。pandas などのパッケージも含まれています。

* flask
* flask-bootstrap
* pandas
* xlrd

###### /tmp へのファイル読み書き

アップロードしたファイルは一時的に /tmp 領域に書き出しています。
それを pandas の read_csv メソッドで読み出しています。

### second generation runtime 設定のポイント

では実際にどのように second generation runtime の設定を行うのか。要点だけ書き出します。

##### ランタイムの指定

[app.yaml](https://cloud.google.com/appengine/docs/standard/python3/config/appref) にてランタイムの指定をします。python37 にします。

```yaml:app.yaml
runtime: python37
```

##### パッケージ

requirements.txt に必要なパッケージを書き出します。以下は公式ドキュメントの例です（私のアプリはバージョンを固定していないので・・本来は pip freeze で取得したバージョンにするべきです）。

```
Flask==0.10.1
python-memcached==1.54
```

https://cloud.google.com/appengine/docs/standard/python3/specifying-dependencies




### デプロイ

今回試す Web アプリのソースは Github 上に公開しています。Cloud Shell 上で次の通りコマンドを実行してください。

```Shell
# ソースの取得
$ git clone https://github.com/ao-log/gae-second-generation-python-demo
$ cd gae-second-generation-python-demo

# デプロイ
$ gcloud app deploy .
```

https://[プロジェクト名].appspot.com/ にアクセスして、デプロイされていることをご確認ください。

# 参考

* [Introducing App Engine Second Generation runtimes and Python 3.7](https://cloud.google.com/blog/products/gcp/introducing-app-engine-second-generation-runtimes-and-python-3-7)

* [App Engine standard environment runtimes](https://cloud.google.com/appengine/docs/standard/appengine-generation)
* [Understanding differences between Python 2 and Python 3 on the App Engine standard environment](https://cloud.google.com/appengine/docs/standard/python3/python-differences)
* [GAE スタンダード環境で scikit-learn を使う](https://medium.com/google-cloud-jp/gae-standard-with-scikit-learn-dff5ad1a0ea1)
