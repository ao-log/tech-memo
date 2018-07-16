Docker コンテナ内でユーザ権限でプログラムを動作させるのに **gosu** を使っている方もいらっしゃると思います。Alpine イメージを使う時には他の選択肢もあるよ、ということで **su-exec** を紹介させていただきます。

# そもそも gosu とは？

指定したユーザ、グループでプログラムを動作させることができます。

[リポジトリはこちら]
https://github.com/tianon/gosu

### gosu が作られた動機

gosu の README に作られた動機が書かれています。su, sudo が予測不可能な TTY、シグナル送信の振る舞いをする問題を持っているためです。同等のことが、[Best practices for writing Dockerfiles](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/#use-multi-stage-builds) にも記載されており、gosu を検討してみてはと書かれています。

> Avoid installing or using sudo as it has unpredictable TTY and signal-forwarding behavior that can cause problems. If you absolutely need functionality similar to sudo, such as initializing the daemon as root but running it as non-root), consider using “gosu”.

### gosu と su のプロセス生成の違い

プロセス生成の例を見てみましょう。gosu は直接、指定した権限でプロセスを実行しているのに対し、su は子プロセスとして実行しています。

* gosu

```shell-session
# gosu nobody:root ps f -ef
PID   USER     TIME  COMMAND
    1 root      0:00 sh
   66 nobody    0:00 ps f -ef
```

* su コマンド

```shell-session
# su - testuser -c "ps f -ef"
No directory, logging in with HOME=/
UID        PID  PPID  C STIME TTY      STAT   TIME CMD
root         1     0  0 11:54 pts/0    Ss     0:00 bash
root        48     1  0 12:03 pts/0    S+     0:00 su - testuser -c ps f -ef
testuser    49    48  0 12:03 ?        Ss     0:00  \_ -su -c ps f -ef
testuser    52    49  0 12:03 ?        R      0:00      \_ ps f -ef
```


# su-exec の利点とは？

では、ここから su-exec について取り上げていきます。

[リポジトリはこちら]
https://github.com/ncopa/su-exec

gosu と同等のことが実現できるソフトです。su-exec の利点は次の通りです。

* gosu よりもサイズが小さい
* Alpine の main リポジトリに入っている。よって、インストールが簡単

### サイズの小ささ

次の三つのイメージのサイズを比較します。

* Alpine イメージ
* gosu をインストールしたイメージ (ベースイメージは Alpine)
* su-exec をインストールしたイメージ (ベースイメージは Alpine)

表にすると以下の通りとなりました。gosu も決してサイズが大きすぎるわけではないのですが、su-exec の小ささが際立ちますね。

| イメージ | サイズ |
| --- | --- |
|Alpine|4.41MB|
|gosu|6.99MB|
|su-exec|4.44MB|

### インストールの容易さ

su-exec をインストールする Dockerfile は以下の通りです。

```Dockerfile
FROM alpine
RUN set -ex; \
    apk add --no-cache su-exec;
```

一方で、gosu は Alpine の main リポジトリに入っていないです(2018年7月16日現在)。[こちらにインストール方法が記載されています](https://github.com/tianon/gosu/blob/master/INSTALL.md)。簡単ではあるものの、Dockerfile が煩雑になります。

サイズ、インストールの容易さ以外にも実績など様々な要素を考慮する必要があるかとは思いますが、su-exec も候補の一つとしていかがでしょう？

# [コラム] postgres、Redis の Alpine ベースイメージでも su-exec が採用されています

postgres、Redis の Alpine ベースイメージでも su-exec が採用されていました。ご参考までに、postgres の Docker Official イメージのパッケージングのページはこちら。
https://github.com/docker-library/postgres

Debian ベースと Alpine ベースの二つがあります。Alpine ベースの方は gosu ではなく su-exec が採用されています。

その時のやりとりがこちらに残っています。
https://github.com/docker-library/postgres/pull/119
その元となった議論が Redis の方でされています。
https://github.com/docker-library/redis/pull/40#issuecomment-169916910

要約すると、su-exec の作者の ncopa さんが、gosu ってサイズが大きいよねという問題提起とともに su-exec を書きました。su-exec はテストも十分行っており、Alpine もサイズを小さくすることを目標にしたイメージです。そのため、Alpine ベースの方は gosu ではなく su-exec を採用する流れとなりました。
