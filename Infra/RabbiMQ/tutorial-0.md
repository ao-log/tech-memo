チュートリアル従って進めていく。
ここでは、用語説明、チュートリアルを進めるにあたって準備を行う。

https://www.rabbitmq.com/tutorials/tutorial-one-python.html

### 事前準備

* RabbitMQ を起動しておく
* pip で pika をインストールする。

##### RabbtMQ

RabbitMQ は、ここではコンテナで起動することとする。
ポート 5672 が RabbtMQ の稼働ポート。15672 は管理画面の稼働ポート(id: guest, password: guest)である。

・docker-compose.yaml

```yaml
version: '2'

services:
  rabbitmq:
    image: rabbitmq:3-management
    ports:
      - 5672:5672
      - 15672:15672
    networks:
      - default
```

コンテナを起動

```
$ docker-compose up
```

##### Python 環境

Python の仮想環境に pika をインストールする。

```
$ python -m venv RabbitMQ
$ source RabbitMQ/bin/activate
$ pip install pika
```
