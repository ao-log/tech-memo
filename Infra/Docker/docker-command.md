
# 典型的な流れ。

イメージの取得から、コンテナ起動、コンテナ停止、削除までの典型的な流れは下記の通り。

```shell-session
# イメージの取得
$ docker pull centos

# 起動　　(-d: バックグラウンド実行、-t: tty 接続をアロケート、-i: コンテナの標準入力を開く)
$ docker run -dti --name centos centos bash

# コンテナに接続
$ docker exec -ti centos bash
もしくは
$ docker attach コンテナID

# 稼動状況確認 (停止したコンテナも見たい場合は、ps -a とする)
$ docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED              STATUS              PORTS               NAMES
9864b6ec5c25        centos              "bash"              About a minute ago   Up 5 seconds                            centos

# コンテナの停止
$ docker kill centos
もしくは、コンテナIDでも良い。
$ docker kill 9864

# コンテナの削除
$ docker rm centos
```

# イメージ関連

```shell-session

# イメージの一覧
$ docker images

# イメージを探す
$ docker search TERM

# イメージの取得
$ docker pull NAME[:TAG]

# イメージの詳細
$ docker inspect NAME[:TAG]

# イメージの削除
$ docker rmi IMAGE
```

# タグ付け

```shell-session
# :バージョンがない場合は、latest に。
$ docker tag centos:latest mycent:1.0
```

# ボリューム

```shell-session
# ローカルの ~/share とコンテナの /share をマッピング
$ docker run -dti -v ~/share:/share centos /bin/bash
```

# ポート

```shell-session
# コンテナの 80 番ポートを、ホストの 8080 番ポートに晒す
$ docker run -dti -p 8080:80 centos
```

# コンテナの情報確認

```shell-session
$ docker inspect ID

# format で特定の情報を出力
$ docker inspect --format="{{.NetworkSettings.IPAddress}}" ID
```

# その他

### プロセス

コンテナ内のプロセス確認。

```shell-session
$ docker top ID
```

### イベント

```shell-session
# コンテナのイベント（起動、削除など）をリアルタイムに表示
$ docker events
```

### ネットワーク

```shell-session
# ネットワークの一覧
$ docker network ls
```

### ボリューム

```shell-session
# ボリュームの一覧
$ docker volume ls
```
