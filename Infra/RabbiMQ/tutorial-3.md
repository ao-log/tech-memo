
## このチュートリアルの概要

このページは RabbitMQ のチュートリアルを元にしている。

https://www.rabbitmq.com/tutorials/tutorial-three-python.html

この章では複数のコンシューマに同じメッセージを送り届けることに焦点をあてる。
「Publish/Subscribe」 として知られるパターンである。

## この章で学べること

##### exchange

* プロデューサはキューに直接メッセージを送付しているわけではない。exchange に送っている。

* exchange はメッセージの処理方法を正確に指定する必要がある。キューに追加するか、複数のキューに追加するか、それともメッセージを破棄するか。

* exchange にはいくつかのタイプがある。利用可能なタイプは「direct」「topic」「headers」「fanout」である。

* この章では「fanout」 にフォーカスする。「fanout」exchange は、受信したメッセージを全てのキューに送る。


```
channel.exchange_declare(exchange='logs',
                         exchange_type='fanout')
```

* exchange のリストアップは次のコマンドで行う。

```
$ sudo rabbitmqctl list_exchanges
```


##### キュー

* キュー名を指定しなかった場合は、ランダムなキュー名がつけられる。また、exclusive=True とすることで、コンシューマの接続が切れるとキューも削除される動作となる。

```
result = channel.queue_declare(exclusive=True)
```

##### binding

* exchange とキューとの関連付けのことを「binding」とよぶ。コンシュー側でどの exchange 宛の通信をどのキューと関連付けるかを指定する。

```
channel.queue_bind(exchange='logs',
                   queue=result.method.queue)
```


* binding 状況の確認は次のコマンドを使用する。

```
$ sudo rabbitmqctl list_bindings
```


## ソース

###### emit_log.py

```python
#!/usr/bin/env python
import pika
import sys

connection = pika.BlockingConnection(pika.ConnectionParameters(host='localhost'))
channel = connection.channel()

channel.exchange_declare(exchange='logs',
                         exchange_type='fanout')

message = ' '.join(sys.argv[1:]) or "info: Hello World!"
channel.basic_publish(exchange='logs',
                      routing_key='',
                      body=message)
print(" [x] Sent %r" % message)
connection.close()
```

###### receive_logs.py

```python
#!/usr/bin/env python
import pika

connection = pika.BlockingConnection(pika.ConnectionParameters(host='localhost'))
channel = connection.channel()

channel.exchange_declare(exchange='logs',
                         exchange_type='fanout')

result = channel.queue_declare(exclusive=True)
queue_name = result.method.queue

channel.queue_bind(exchange='logs',
                   queue=queue_name)

print(' [*] Waiting for logs. To exit press CTRL+C')

def callback(ch, method, properties, body):
    print(" [x] %r" % body)

channel.basic_consume(callback,
                      queue=queue_name,
                      no_ack=True)

channel.start_consuming()
```


## 実行

2 つのウィンドウを用意し、それぞれでコンシューマを起動させておく。片方はログのファイル出力用、もう片方は画面表示用である。

```
// ウィンドウ 1 （ファイル出力用）
$ python receive_logs.py > logs_from_rabbit.log
 [*] Waiting for messages. To exit press CTRL+C

// ウィンドウ 2 （画面表示用）
$ python receive_logs.py
 [*] Waiting for messages. To exit press CTRL+C
```

プロデューサを実行する。

```
$ python emit_log.py
 [x] Sent 'info: Hello World!'
$ python emit_log.py
 [x] Sent 'info: Hello World!'
$ python emit_log.py
 [x] Sent 'info: Hello World!'
```

すると、それぞれのコンシューマに同じメッセージが送られる。

```
// ウィンドウ 1 （ファイル出力用）
$ cat logs_from_rabbit.log
 [*] Waiting for logs. To exit press CTRL+C
 [x] b'info: Hello World!'
 [x] b'info: Hello World!'
 [x] b'info: Hello World!'

// ウィンドウ 2 （画面表示用）
$ python receive_logs.py
 [*] Waiting for logs. To exit press CTRL+C
 [x] b'info: Hello World!'
 [x] b'info: Hello World!'
 [x] b'info: Hello World!'
```

binding 状況を確認する。logs exchange が 2 つのキューに binding されている。

```
$ sudo rabbitmqctl list_bindings
Listing bindings for vhost /...
logs    exchange        amq.gen-CjabB0vJaGaCy46iUbaMIQ  queue   amq.gen-CjabB0vJaGaCy46iUbaMIQ  []
logs    exchange        amq.gen-lvtNlJ-RqLh-Ps5MIKLywg  queue   amq.gen-lvtNlJ-RqLh-Ps5MIKLywg  []
```
