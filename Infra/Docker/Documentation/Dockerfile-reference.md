
[Dockerfile reference](https://docs.docker.com/engine/reference/builder/)

* FROM
* ARG
  * Dockerfile 内でのみ使用できる変数
* RUN
* CMD
  * `docker run` 時に上書きされる
* LABEL
  * メタデータ
* EXPOSE
* ENV
* ADD
  * `ADD rootfs.tar.xz /` のような使い方ができる
* COPY
  * ファイルのコピーのみ。コピー用途ではこちらが望ましいとされる
* ENTRYPOINT
  * `docker run` 時に上書きできない
  * 以下のような指定も可能。
  ```Dockerfile
  ENTRYPOINT ["ping"]
  CMD ["127.0.0.1", "-c", "50"]
  ```
* VOLUME
* USER
* WORKDIR
* ARG
* ONBUILD
* STOPSIGNAL
* HEALTHCHECK
* SHELL


[Dockerfile を書くベスト・プラクティス](https://docs.docker.jp/engine/userguide/eng-image/dockerfile_best-practice.html#)

* `.dockerignore` で除外対象のファイル、ディレクトリを指定
* 不要なパッケージをインストールしない
* コンテナごとに 1 プロセス
* レイヤを最小数にする
* `RUN apt-get update && apt-get install -y` のように実行することで、キャッシュを使用せず `apt-get update` が実行される。パッケージのバージョン指定は、何をキャッシュしているか気にせずに特定バージョンを取得した上での構築が強制される


