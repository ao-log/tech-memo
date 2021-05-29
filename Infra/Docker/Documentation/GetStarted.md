https://docs.docker.com/get-started/overview/


#### Docker ボリューム

```
// ボリュームの作成
$ docker volume create todo-db

// ボリュームをマウントしてコンテナ起動
$ docker run -dp 3000:3000 -v todo-db:/etc/todos getting-started

// ボリュームの確認
$ docker volume inspect todo-db
[
    {
        "CreatedAt": "2019-09-26T02:18:36Z",
        "Driver": "local",
        "Labels": {},
        "Mountpoint": "/var/lib/docker/volumes/todo-db/_data",
        "Name": "todo-db",
        "Options": {},
        "Scope": "local"
    }
]
```


#### バインドマウント

バインドマウントはローカルホストの一部を共用するようなマウントのこと。

```
$ docker run -dp 3000:3000 \
     -w /app -v "$(pwd):/app" \
     node:12-alpine \
     sh -c "yarn install && yarn run dev"
```


#### Multi container apps

* 一つのコンテナでは一つのことをすべき（複数を詰め込むべきではない）
  * スケールを別々に行えるため
  * バージョン管理の容易さ
  * 複数プロセスの管理は煩雑なため

異なるコンテナとの通信はネットワーキングにより解決する。

```
// ネットワークの作成
$ docker network create todo-app

// ネットワークを指定してコンテナを起動
$ docker run -d \
     --network todo-app --network-alias mysql \
     -v todo-mysql-data:/var/lib/mysql \
     -e MYSQL_ROOT_PASSWORD=secret \
     -e MYSQL_DATABASE=todos \
     mysql:5.7

// 同一ネットワーク内に別コンテナを起動
$ docker run -it --network todo-app nicolaka/netshoot

// エイリアスで設定した名前で名前解決可能
$ dig mysql
```




#### Image-building best practices

**Security scanning**

```docker scan イメージ``` によりイメージのスキャンが可能。

**Image layering**

```docker image history イメージ``` によりレイヤーの確認が可能。

**Layer caching**

Dockerfile の各行の実行結果はレイヤーとしてキャッシュされている。次回以降は前回のキャッシュを使いつつ、差分のある箇所からビルドすることが可能。

**Multi-stage builds**

次の例のように複数ステージのビルドが可能。
最後のステージでビルドの成果物を COPY で配置することで、ビルド時のみ必要なパッケージ類などを含めないようにすることができる。

```Dockerfile
# syntax=docker/dockerfile:1
FROM maven AS build
WORKDIR /app
COPY . .
RUN mvn package

FROM tomcat
COPY --from=build /app/target/file.war /usr/local/tomcat/webapps 
```


