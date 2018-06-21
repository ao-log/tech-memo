# この記事について

Dockerfile Best Practices を独学用にまとめたものです。理解を深めるために、順序を入れ替えたり、元の記事にない記述を足したり、逆に削ったりしています。まとめていて感じたのは、レイヤがどのようにできるか、どのような条件でキャッシュを使うかの理解が重要だと感じました。

https://docs.docker.com/develop/develop-images/dockerfile_best-practices/

### レイヤの考え方について

Docker イメージは read オンリーのレイヤーから構成されています。個々のレイヤは Dockerfile の各行に該当します。レイヤは前のレイヤからの差分としてスタックされ積み重なっていきます。例えば、次の Dockerfile を考えます。

```Dockerfile
FROM ubuntu:15.04
COPY . /app
RUN make /app
CMD python /app/app.py
```

各行のコマンドは一つのレイヤを作ります。

イメージからコンテナを起動すると、書き込み可能な新しいレイヤが作られます。稼働中のコンテナへのすべての変更はこのレイヤに対してなされます。

レイヤについての詳細は [storage drivers](https://docs.docker.com/storage/storagedriver/) のページをご参照ください。

# キャッシュの有効活用

イメージは複数のレイヤを積み重ねたものになっており、次のコマンドでレイヤの積みかさなりを確認できます。

```shell
$ docker history イメージ名
```

レイヤのキャッシュを有効活用することで、ビルド時間の短縮や、ディスク消費量の抑制につながります。逆にキャッシュを使っていることを知らないと思わぬ処理がスキップされたりして、意図せぬ動作の原因となります。

キャッシュを使うかどうかを、どのように判定しているかを説明します。

* イメージを作る時、Docker は Dockerfile の記述に従って上から順に処理していきます。すでに存在するレイヤを探してキャッシュに存在すれば、それを再利用します。もし、キャッシュを使いたくないのであれば、ビルド時に ```--no-cache=true``` オプションを指定します。

* 親イメージがキャッシュにあればそれを使います。次にそのイメージの子イメージを確認します。手順が同じ行まではキャッシュを使いますが、手順が異なる行以降はキャッシュを使わず新しいレイヤを作成します。

* ADD, COPY の場合は、各ファイルについてチェックサムを計算します。最終修正時刻、最終アクセス時刻はチェックサムに影響しません。すでに存在するイメージのチェックサムと比較し、異なる場合はキャッシュは使われません。

* ```RUN apt-get -y update``` の場合は、コマンドの文字列だけを比較します。よって、後述するように、過去に ```apt-get -y update``` を実行して作成したイメージがある場合はそのキャッシュを使うので、以降作成するイメージは ```apt-get -y update``` がスキップされます。これを避けるテクニックは RUN の所で後述します。

# マルチステージビルド

マルチステージビルドは Docker 17.05 以降で使える機能です。

例えば、バイナリを作る場合は、バイナリを作るために様々なファイルやパッケージが必要になりますが、バイナリ以外のものが Docker イメージに含まれるのは Dirty だと言えます。ここで、一回目のビルドでバイナリを作り、二回目のビルドで先ほど作成したバイナリをコピーするだけにします。すると不要なファイルが含まれない状態にでき、Docker イメージを小さくできるほかレイヤ数も抑えることができます。

以下、要点を抜粋した Dockerfie です。FROM 行が 2 個あること、COPY で前のビルドを指定し、バイナリをコピーしていることがポイントです。

```Dockerfile
FROM golang:1.9.2-alpine3.6 AS build
...(略)
RUN go build -o /bin/project

FROM scratch
COPY --from=build /bin/project /bin/project
ENTRYPOINT ["/bin/project"]
CMD ["--help"]
```


# その他のテェックリスト

* ビルドに含めたくないファイルは ```.dockerignore``` ファイルに記述します。
* 不要なパッケージをインストールしないこと
* 一つのコンテナは一つのことだけに関心を持つようにすること。言い換えると、複数アプリを稼働させないこと。
* パッケージインストールなど複数行にわたって書くものはなるべくアルファベット順にしたほうが良いです。理由は、重複を防ぐなどリストのメンテを楽にするメリットがあるため。


# Dockerfile の各コマンド

以下は各コマンドの使い方についての Tips です。

### FROM

* なるべく公式リポジトリのイメージを使うこと。

* [Alpine イメージ](https://hub.docker.com/_/alpine/) がオススメ。最小構成になるようにコントロールされており、サイズが 5 MB 以下の小さなイメージです。

### RUN

##### apt-get

update と install を別々の行にせずに、次のように一行で書くのが良いです。  

**理由:**
Dockerfile のビルド時は過去のキャッシュを使うため。もし、過去に apt-get update を実行したイメージがあるとそのキャッシュを使うので、apt-get update が実行されません。結果として、最新のパッケージをインストールしてくれません。よって、キャッシュを使わないように工夫する必要があるのです。

```Dockerfile
RUN apt-get update && apt-get install -y \
    aufs-tools \
    automake \
    build-essential \
    curl \
    dpkg-sig \
    libcap-dev \
    libsqlite3-dev \
    mercurial \
    reprepro \
    ruby1.9.1 \
    ruby1.9.1-dev \
    s3cmd=1.1.* \
 && rm -rf /var/lib/apt/lists/*
```

##### パイプで渡すコマンドを実行するとき

set -o pipefail から始めると良いです。  

**理由:**
コマンドの成功判定に exit code を用いるため。set -o pipefail がある場合は途中のコマンドが失敗すると RUN も失敗とみなされますが、ない場合は最後のコマンドさえ通れば exit code = 0 となり成功とみなされるため。

```Dockerfile
RUN set -o pipefail && wget -O - https://some.site \
    | wc -l > /number
```

### CMD

次のような形式で使うのが良いです。

```Dockerfile
CMD [“executable”, “param1”, “param2”…]
```

なお、ENTRYPOINT の使用に慣れていない場合は、ENTRYPOINT と一緒に CMD [“パラメータ”, “パラメータ”] の形式で使うのは避けた方が良いです。ENTRYPOINT や CMD との併用についての詳細な情報は [リファレンス](https://docs.docker.com/engine/reference/builder/#exec-form-entrypoint-example) もご参照ください。

### EXPOSE

伝統的なポートを使うのが良いです。
例えば、http であれば 80、MongoDB であれば 27017。

### ENV

プログラム中の定数のような使い方ができます。

```Dockerfile
ENV PG_MAJOR 9.3
ENV PG_VERSION 9.3.4
RUN curl -SL http://example.com/postgres-$PG_VERSION.tar.xz | tar -xJC /usr/src/postgress && …
ENV PATH /usr/local/postgres-$PG_MAJOR/bin:$PATH
```

イメージのビルド時に一時的に使う環境変数であれば、ENV を使わず RUN 内で対応すると良いです。理由は、ENV で指定した変数はイメージに組み込まれてしまい、コンテナにアクセスすると ENV で設定した変数がセットされてしまうため。

```Dockerfile
FROM alpine
RUN export ADMIN_USER="mark" \
    && echo $ADMIN_USER > ./mark \
    && unset ADMIN_USER
CMD sh
```

### ADD or COPY

ADD よりも COPY の方が望ましいです。これは、COPY にはローカルファイルをイメージ内にコピーする機能しかできないためで、単にファイルをコピーしたい時は COPY を使うと意図が明確になります。

ADD は tar の配置などに利用できますが、以下の書き方はアンチパターンです。ADD の行でレイヤが追加されてしまうためです。

```Dockerfile
ADD http://example.com/big.tar.xz /usr/src/things/
RUN tar -xJf /usr/src/things/big.tar.xz -C /usr/src/things
RUN make -C /usr/src/things all
And instead, do something like:
```

代わりに以下のように書くと良いです。レイヤ数を最小化しつつ、tar 展開後の不要なファイルを残さない利点もあります。

```Dockerfile
RUN mkdir -p /usr/src/things \
    && curl -SL http://example.com/big.tar.xz \
    | tar -xJC /usr/src/things \
    && make -C /usr/src/things all
```

### ENTRYPOINT

イメージの主たるコマンドを ENTRYPOINT に設定するのが有用な使い方の一つです。例えば、次のように設定します。デフォルトの引数は CMD で設定します。

```Dockerfile
ENTRYPOINT ["s3cmd"]
CMD ["--help"]
```

すると、コンテナ起動時に次のようにコマンドを書けます。もし、ENTRYPOINT を指定していない場合は、イメージ名とコマンド名が同じなので、s3cmd s3cmd と同じ言葉を続けて書かなければなりません。

```shell
$ docker run s3cmd ls s3://mybucket
```

ENTRYPOINT にヘルパースクリプトを設定することもできます。以下は、postgres 公式イメージの例です。

```shell
#!/bin/bash
set -e

if [ "$1" = 'postgres' ]; then
    chown -R postgres "$PGDATA"

    if [ -z "$(ls -A "$PGDATA")" ]; then
        gosu postgres initdb
    fi

    exec gosu postgres "$@"
fi

exec "$@"
```

Dockerfile では、次のように ENTRYPOINT にヘルパースクリプトを渡します。

```Dockerfile
COPY ./docker-entrypoint.sh /
ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["postgres"]
```

### USER

特権が不要な場合は、USER を使うことで非 root ユーザでサービスを実行します。

sudo の使用やインストールは、TTY やシグナル送信の予期せぬ動作につながるため、避けてください。デーモンの初期化を root で行う必要があるものの、非 root ユーザで動作させる場合、[gosu](https://github.com/tianon/gosu) の使用を検討してください。

### WORKDIR

ディレクトリの移動は、WORKDIR に絶対パスを指定して行うと良いです。RUN コマンド内で cd をするのは可読性が悪くなります。
