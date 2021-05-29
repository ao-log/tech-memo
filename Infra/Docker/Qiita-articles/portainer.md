以下に投稿した記事。

https://qiita.com/ao_log/items/d8ef847c826746f9e84b

# 記事

コンテナを管理する方法は色々あると思うのですが、 自宅で技術検証をするときは portainer が便利そうでした。
公式サイトはこちらです。軽量な管理 UI で、簡単に Docker ホスト、Swarm クラスタを管理できると記載されています。
https://portainer.io/

# よいと感じた点

* GUI で稼働中のコンテナやイメージ、ネットワーク、ストレージを見れます。マウス操作だけでいいので、モニタリングが手軽です。
* 画面もおしゃれだと思いました(人によって好みが違うとは思いますが・・)。
* コンソール接続もできます。```docker exec -ti コンテナ名 bash``` をしなくても接続できるので手軽です。
* 著名なアプリのコンテナテンプレートが用意されていて、GUI 上でポチポチしているだけでコンテナを作れます。今まで触ったことのない技術に触りたくなりました。

# 導入方法

コンテナイメージ「portainer/portainer」をプルして起動します。

脱線しますが、個人的には OS は ubuntu を使うのが楽です。最近のソフトも結構パッケージ化されています。（仕事で Red Hat Enterprise Linux に機械学習の Caffe をインストールしたときは大変すぎました。。）

また、素の Docker を使うよりは Docker Compose を使うのがオススメです。コンテナの起動方法を yaml で書けるので、docker コマンドのオプションを覚える必要がなく、毎回長いオプションを打つ必要もなくなります。

docker compose は次のようにインストールします(Debian 系の場合)。

```
$ sudo apt install docker-compose
$ sudo usermod -aG docker ユーザ名
```

docker-compose.yaml は次の内容にします。

```
$ cat docker-compose.yaml
version: '2'

services:
  portainer:
    image: portainer/portainer
    ports:
      - 9000:9000
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
```

では、コンテナを起動します。

```
$ docker-compose up
```

# 管理画面

http://localhost:9000 にアクセスするとログイン画面になります。初回はパスワードを設定します。

![2018-03-07 21.54.18 からのスクリーンショット.png](https://qiita-image-store.s3.amazonaws.com/0/48133/885c0437-f108-c9d9-5d0d-4a2dc5efbc03.png)

Local を選択します。

![2018-03-07 21.55.24 からのスクリーンショット.png](https://qiita-image-store.s3.amazonaws.com/0/48133/7020ad1f-1a4b-4d4a-9511-e05607334d04.png)

### ダッシュボード

![2018-03-07 21.57.26 からのスクリーンショット.png](https://qiita-image-store.s3.amazonaws.com/0/48133/6e86d8c8-9d08-8427-7a6d-77a1b5434684.png)

### コンテナ一覧

![2018-03-07 21.58.18 からのスクリーンショット.png](https://qiita-image-store.s3.amazonaws.com/0/48133/57de9dc9-5a5c-cc29-f3d5-1a9ec0e7c374.png)

### コンテナ詳細

コンテナに対して起動、停止、削除のアクションを発行できます。また、リソース使用状況やログの表示もできます。

![2018-03-07 22.50.27 からのスクリーンショット.png](https://qiita-image-store.s3.amazonaws.com/0/48133/f27ecc95-164d-74ad-590d-c7485461cbee.png)

### コンソール接続

[Containers] → [コンテナ名] → [Console] と遷移し、Connect をクリックするとコンソール接続できます。ブラウザ上で作業できます。
```docker exec -ti コンテナ名 bash``` をしなくても接続できるので、ちょっとコンテナの中を確認したい時に便利です。

![2018-03-07 23.06.05 からのスクリーンショット.png](https://qiita-image-store.s3.amazonaws.com/0/48133/8a33b639-6c11-8f21-8d81-9ff8116f1231.png)

### イメージ一覧

![2018-03-07 22.00.10 からのスクリーンショット.png](https://qiita-image-store.s3.amazonaws.com/0/48133/4d358b58-8976-7085-8ef2-19c9eb733b40.png)

### ネットワーク一覧

私の場合は、docker compose がよしなに作ってくれたサブネットがいくつか。裏でツールがしてくれている処理も、こうなっているんだなと実感できます。

![2018-03-07 22.00.27 からのスクリーンショット.png](https://qiita-image-store.s3.amazonaws.com/0/48133/61999446-1ed9-a2e6-84f2-942145b8e90b.png)

### ストレージ一覧

![2018-03-07 22.00.45 からのスクリーンショット.png](https://qiita-image-store.s3.amazonaws.com/0/48133/e1edce22-f2ad-cb6e-72e3-f95d7699eeca.png)

### アプリのテンプレート

著名なアプリのコンテナテンプレートです。これだけ並んでいると、今まで触ったことのない技術のコンテナを触ってみようという気になりました。皆様はいかがでしょう？

![2018-03-07 22.02.59 からのスクリーンショット.png](https://qiita-image-store.s3.amazonaws.com/0/48133/54a39da8-f8f3-ca55-2db2-26946710dfaa.png)
