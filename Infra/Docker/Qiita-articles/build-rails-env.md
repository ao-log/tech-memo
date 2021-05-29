以下に投稿した記事。

https://qiita.com/ao_log/items/13b923b5af43fc3dafbc

# 記事

Rails の開発をするとき、自分の PC をなるべく汚したくないと思っていました。そんなときに役立つのが Docker です！　各ソフトは全部コンテナ側に押し込んでしまいましょう。この記事では Rails の開発環境を Docker で作成する方法を紹介します。

Docker は Linux VM 上で動かします。コンテナは次の 4 つを作成します。

* Rails Server: rails を動作させるサーバ
* Postgres: 今回採用するデータベース
* pgweb: Postgres の内容を Web で表示
* cloud9 V3 SDK: IDE

cloud9 はブラウザで開発できる IDE です。Web サービスで提供されているものだけでなく、自分の PC 上に環境を構築することができます。

### 注意点

以下、環境構築方法を紹介していきます。macOS での手順です。ただ、他の OS でも同じような考え方で構築できると思います。


# Docker 環境の構築

Linux VM を構築しその上で Docker を動かすわけですが、簡単に構築できるツールがあります。Docker Tools です。その方法は、こちらの記事で書かせていただきました。

[Docker Toolbox を使って macOS上に Docker 環境を構築する](https://qiita.com/ao_log/items/b18cf793f4f8cf8d53f4)

概要をざっくり説明すると、Docker Tools とは Docker の管理に便利なツールを寄せ集めたものです。インストーラの画面の指示に従うだけで、簡単に Docker 環境を構築できます。今回使うのは次の 3 つのツールです。

* Docker Machine: Docker Engine が稼働する VM を作成したり、管理するツール
* Docker Compose: 複数コンテナを使用するDockerアプリを定義、実行するツール
* Docker QuickStart shell: macOS上のターミナルから VM 上の Docker Engine を操作するための環境整備

Docker Tools インストール後に、ランチャーパッドから「Docker QuickStart shell」のアイコンをクリックすると自動的に Linux VM が構築されます。また、その上の Docker 環境を利用するための環境変数をセットした状態で、ターミナルが起動してくれます。以降はこのターミナル上で、コンテナ構築を行っていきます。

# コンテナ作成の準備

今回作成するコンテナを改めて書くと次の 4 つになります。

* Rails Server: rails を動作させるサーバ
* Postgres: 今回採用するデータベース
* pgweb: Postgres の内容を Web で表示
* cloud9 V3 SDK: IDE

このうち、Postgres, pgweb は公開されているイメージをそのまま使います。Rails Server, cloud9 は Dockerfile を作成しビルドします。

### ディレクトリ、ファイル構成

以下の構成にします。

```
|- docker-compose.yml    Docker Compose の設定ファイル
|- Dockerfike-web        Rails Server 用の Dockerfile
|- Dockerfile-cloud9     cloud9 V3 SDK 用の Dockerfile
|- src/                  Rails のソース類を配置するディレクトリ
    |- Gemfile           Rails Server の bundle 用
    |- Gemfile.lock      Rails Server の bundle 用
```

Rails のソースは src に配置し、ここをコンテナ側からも参照できるようにします。

### Dockerfile の内容

##### Rails Server

Rails Server 用の Dockerfile です。今回は、複数の Dockerfile を作成するため、ファイル名にサフィックスをつけて用途が分かるようにしておきます。

・Dockerfile-web

```dockerfile:Dockerfile-web
FROM ruby
RUN ln -sf /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
RUN apt-get update -qq && apt-get install -y build-essential libpq-dev nodejs
ENV BUNDLE_JOBS=4 \
    APP_DIR=/app
RUN mkdir $APP_DIR
WORKDIR $APP_DIR
COPY src/Gemfile $APP_DIR
COPY src/Gemfile.lock $APP_DIR
RUN bundle install
```

##### cloud9

cloud9 用の Dockerfile です。[Cloud9 Core](https://github.com/c9/core) のインストール手順に習っています。

・Dockerfile-cloud9

```dockerfile:Dockerfile-cloud9
FROM node
RUN ln -sf /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
RUN git clone https://github.com/c9/core.git /cloud9
WORKDIR /cloud9
RUN ./scripts/install-sdk.sh
WORKDIR /workspace
```

> [memo] 私の環境だと、4 行目の install-sdk.sh 内の処理で npm install に失敗しました。registry.npmjs.org にアクセスできない旨のエラーメッセージ。macOS 側の DNS にルータのアドレスが設定されていたのですが、8.8.8.8 にすることで改善しました。

### docker-compose.yml

以下の内容にします。

```yaml
version: '3'
services:
  postgres:
    image: postgres
    volumes:
      - postgres-data:/var/lib/postgresql/data

  postgres-gui:
    image: donnex/pgweb
    command: -s --bind=0.0.0.0 --listen=8080 --url postgresql://postgres@postgres/myapp_development?sslmode=disable
    ports:
      - "7000:8080"
    depends_on:
      - postgres

  web:
    build:
      context: .
      dockerfile: Dockerfile-web
    command: bundle exec rails s -p 3000 -b '0.0.0.0'
    volumes:
      - ./src:/app
      - bundle-data:/usr/local/bundle
    ports:
      - "3000:3000"
    depends_on:
      - postgres

  cloud9:
    build:
      context: .
      dockerfile: Dockerfile-cloud9
    command: node /cloud9/server.js --port 80 -w /workspace -l 0.0.0.0 --auth user:password
    ports:
      - 9000:80
    volumes:
      - ./src:/workspace
      - cloud9-data:/cloud9

volumes:
  postgres-data:
  bundle-data:
  cloud9-data:
```

ポイントは次の通りです。

* 永続化するデータは volumes で指定します。コンテナを削除しても残っています。
* 起動順序は depends_on で指定します。ただ、この方法だと、コンテナ上のサービス起動が完了していることが保障されません。保障する方法はこの記事では対象外とします。

### Gemfile

Rails Server の bundle 用です。下記内容にします。

・src/Gemfile

```:src/Gemfile
source 'https://rubygems.org'
gem 'rails', '5.1.4'
```

・src/Gemfile.lock

```:src/Gemfile.lock
空にする。
```

# コンテナのビルド

上記ファイルの準備ができたら、ビルドを実行します。数十分程度かかります。

```shell
$ docker-compose build
```

このコマンドを実行すると、次の内容の処理が実行されます。

* Dockerfile の内容に従ってイメージのビルド。

> 2018年2月3日時点ではここまで書いた手順でうまくいっています。なお、イメージのビルドは、PC 環境の違いやバージョン違いなどの理由ですんなり通らないことも多いので、トライアンドエラーで頑張らないといけない場合も結構あります。

# Rails アプリケーションの作成

では、早速できたイメージからコンテナを起動し、rails new で Rails アプリケーションを作成します。

```shell
$ docker-compose run --rm web rails new . --force --database=postgresql
```

処理が終わると rails new で作られたファイル群が src ディレクトリの下に出来ます。

```shell-session
$ ls src/
Gemfile      README.md    app          config       db           log          public       tmp
Gemfile.lock Rakefile     bin          config.ru    lib          package.json test         vendor
```

DB との連携のため、rails の設定ファイルを編集します。

・src/config/database.yml

```yaml:src/config/database.yml
default: &default
  adapter: postgresql
  encoding: unicode
  host: postgres
  username: postgres
  password:
  pool: 5

development:
  <<: *default
  database: myapp_development

test:
  <<: *default
  database: myapp_test
```

DB を作成しておきます。

```
$ docker-compose start postgres
$ docker-compose run --rm web rails db:create
```

# コンテナを起動

コンテナを起動します。

```
$ docker-compose up

# バックグラウンドで起動する場合は -d オプションをつけます。
$ docker-compose up -d
```

では、表示確認をしてみましょう。

### Rails Server

http://192.168.99.100:3000 にアクセスすると、Rails のページを表示できます。

![スクリーンショット 2017-12-30 22.34.46.png](https://qiita-image-store.s3.amazonaws.com/0/48133/f0c6d919-85e0-3eaf-131f-e068c74f1816.png)

### pgweb

http://192.168.99.100:7000 にアクセスすると、pgweb のページを表示できます。

![スクリーンショット 2018-02-03 18.25.05.png](https://qiita-image-store.s3.amazonaws.com/0/48133/30f285b5-4c6d-6267-389d-192f6a971a2c.png)

### cloud9 v3 SDK

http://192.168.99.100:9000 にアクセスすると、cloud9 のページを表示できます。ユーザ名は「user」、パスワードは「password」です。

![スクリーンショット 2018-02-03 18.27.07.png](https://qiita-image-store.s3.amazonaws.com/0/48133/873f1af1-3944-7dfb-5e9f-17d8f8772d73.png)

# よく使う Docker Compose コマンド

### Rails 関連

rails コマンドはコンテナを起動して実行します。

```shell-session
// db:migrate
$ docker-compose run --rm web rails db:migrate

// テスト
$ docker-compose run --rm web rails test
```


### コンテナ稼動に関するもの

##### 稼動状況の確認

```shell-session
$ docker-compose ps
        Name                      Command               State           Ports         
--------------------------------------------------------------------------------------
rails_cloud9_1         node /cloud9/server.js --p ...   Up      0.0.0.0:9000->80/tcp  
rails_postgres-gui_1   /app/pgweb_linux_amd64 -s  ...   Up      0.0.0.0:7000->8080/tcp
rails_postgres_1       docker-entrypoint.sh postgres    Up      5432/tcp              
rails_web_1            bundle exec rails s -p 300 ...   Up      0.0.0.0:3000->3000/tcp
```

##### コンテナ上で動作しているプロセスの表示

```shell-session
$ docker-compose top
```

##### 一時停止

```shell-session
$ docker-compose stop
```

##### 停止したコンテナの削除（volume で永続化したデータは残ります）

```shell-session
$ docker-compose rm
Going to remove rails_web_1, rails_postgres-gui_1, rails_postgres_1, rails_cloud9_1
Are you sure? [yN] y ← y を入力し [Enter]
```

##### ボリュームの削除

```shell-session
// 一覧の表示
$ docker volume ls

// 削除
$ docker volume rm ボリューム名
```


# 参考
[Quickstart: Compose and Rails](https://docs.docker.com/compose/rails/)

[Cloud9 Core](https://github.com/c9/core)
