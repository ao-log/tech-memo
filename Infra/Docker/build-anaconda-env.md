私は、Anaconda の JupyterLab を Docker で使っています。この記事では、その理由と、環境構築方法を示します。

# Docker で使う理由

### パス問題の解消

Anaconda は便利なソフトなのですが、デメリットもあります。その一つが、OS 主要コマンドの一部パスを Anaconda 内蔵のコマンドに通してしまう問題。```[Anaconda インストールディレクトリ]/bin``` に PATH が通されるのですが、この中に curl, openssl などのコマンドもあります。

Anaconda を Docker コンテナとして動作させることで、パスの問題を解消できます。

### Docker を使うデメリット

普通に OS に Anaconda をインストールするのとは使用感が変わってきます。

* Python コマンドは docker コマンド経由での実行になる。
* Spyder などの GUI ツールを使うのは一手間かかる（この記事では、これらを使えるようにする方法は書きません）

また、Docker 環境構築が必要な分だけ、環境構築に一手間かかります。

# 環境構築手順

macOS での手順を示します。他の OS でも同じような考え方でできると思いますが、この記事では扱いません。

### Docker 環境の構築

まず DockerTools をインストールします。これは、Docker 環境構築に便利なソフトを寄せ集めたものです。
[Docker Toolbox を使って macOS上に Docker 環境を構築する](https://qiita.com/ao_log/items/b18cf793f4f8cf8d53f4)

DockerTools はインストーラの指示に従うだけでインストールできます。インストール後に、ランチャーパッドから「Docker QuickStart shell」のアイコンをクリックします。すると自動的に Docker を稼動させるための Linux VM が構築されます。また、VM 上の Docker 環境を利用するための環境変数をセットした状態で、ターミナルを起動してくれます。以降はこのターミナル上で、コンテナ構築を行っていきます。

### Anaconda 用コンテナの作成

今回は、Docker Compose を使って構築します。ファイル構成は次の通りにします。

```
|- docker-compose.yaml
|- Dockerfile
|- workspace/  (Pythonソースやノートブックを置くディレクトリ。コンテナの /workspace にマッピングします)
```

docker-compose.yaml, Dockerfile は次の内容にします。

・docker-compose.yml

```yaml
version: '3'
services:
    anaconda:
        build: .
        volumes:
            - ./workspace:/workspace
        ports:
            - "8888:8888"
```

・Dockerfile

```Dockerfile
FROM continuumio/anaconda3
WORKDIR /workspace
CMD jupyter-lab --no-browser \
  --port=8888 --ip=0.0.0.0 --allow-root
```

上記の準備ができたらコンテナイメージをビルドします。

```
$ docker-compose build
```

ビルドが終わったらコンテナを起動します。

```
$ docker-compose up

// ターミナルに token が出力されます。これを JupyterLab へのログインに使います。
http://0.0.0.0:8888/?token=...
```

ブラウザで、 http://192.168.99.100:8888 にアクセスすると、JupyterLab の画面にアクセスできます。コンテナ起動時に出力された token 文字列を入力するとログインでき、使えるようになります。

![スクリーンショット 2018-02-10 18.28.15.png](https://qiita-image-store.s3.amazonaws.com/0/48133/6f010ddf-6948-7555-3abc-41e9d325b02d.png)


# どのように運用するか

### 複数の仮想環境を作りたい場合

普通の Anaconda の使い方だと、```conda create``` で仮想環境を作り、activate して仮想環境に入って作業、という流れになるかと思います。Docker の場合は、環境ごとに Docker イメージを作るのが一つの手と考えています。それぞれで JupyterLab が必要な場合は、ポート番号を変更して対応します。

```
PJ-A/
  |- docker-compose.yaml
  |- Dockerfile
  |- workspace/

PJ-B/
  |- docker-compose.yaml
  |- Dockerfile
  |- workspace/
```

### Anaconda のパッケージを揃えたい場合

現在、作業している Anaconda のパッケージ情報は次のようにして yaml で出力できます。

```
# export
$ docker-compose run --rm anaconda conda env export -n root > workspace/env.yaml
```

この yaml を読み込むことで Anaconda のパッケージを揃えます。

まず、workspace ディレクトリに env.yaml を配置します。
次に Dockerfile を編集します。

```Dockerfile
FROM continuumio/anaconda3
WORKDIR /workspace
COPY workspace/env.yaml /workspace
RUN conda env update --file /workspace/env.yaml
CMD jupyter-lab --no-browser \
  --port=8888 --ip=0.0.0.0 --allow-root
```

次のコマンドを実行することで、イメージをビルドし直します。

```
$ docker-compose build
```

### Python スクリプトを実行する場合

```workspace/test.py``` を実行するには次のコマンドを実行します。

```shell-session
$ docker-compose run --rm anaconda python test.py
```

コマンドが長いので、エイリアスを使うのが楽です。

```shell-session
$ alias dpython="docker-compose run --rm anaconda python"
$ dpython test.py
```
