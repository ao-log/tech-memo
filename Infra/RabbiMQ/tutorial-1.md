
## このチュートリアルの概要

このページは RabbitMQ のチュートリアルを元にしている。

https://www.rabbitmq.com/tutorials/tutorial-one-python.html

キュー「hello」対し、メッセージを send, receive する方法を学ぶ。

## この章で学べること

* メッセージを送信するプログラムのことを **プロデューサー** 呼ぶ。

* メッセージは **キュー** ためられていく。

* **consumer** メッセージを待ち受けるプログラムである。

* RabbitMQ はキューに直接メッセージを送信せず、exchange 経由でとなる。
空で指定された際はデフォルトの exchange を使う。この exchange は特別で、どのキューにメッセージを送信するかを指定することが出来る。
キューは「routing_key」で指定する。

```
channel.basic_publish(exchange='',
                      routing_key='hello',
                      body='Hello World!')
```

* メッセージ受信時は Pika ライブラリから callback メソッドが実行される動作となる。この例では、メッセージ受信時にメッセージを出力させている。

```
def callback(ch, method, properties, body):
    print(" [x] Received %r" % body)
```

キュー「hello」からのメッセージ受信設定は次のように行う。

```
channel.basic_consume(callback,
                      queue='hello',
                      no_ack=True)
```

次の関数を実行することで、内部でループ処理を実行する。メッセージを受信することでcallback メソッドを実行する動作となる。終了するには CTRL+C を実行する。

```
channel.start_consuming()
```

## ソース

###### send.py

```python
#!/usr/bin/env python
import pika

connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
channel = connection.channel()

channel.queue_declare(queue='hello')

channel.basic_publish(exchange='',
                      routing_key='hello',
                      body='Hello World!')
print(" [x] Sent 'Hello World!'")

connection.close()
```

###### recieve.py

```python
#!/usr/bin/env python
import pika

connection = pika.BlockingConnection(pika.ConnectionParameters(host='localhost'))
channel = connection.channel()

channel.queue_declare(queue='hello')

def callback(ch, method, properties, body):
    print(" [x] Received %r" % body)

channel.basic_consume(callback,
                      queue='hello',
                      no_ack=True)

print(' [*] Waiting for messages. To exit press CTRL+C')
channel.start_consuming()
```

## 実行

コンシューマを起動させておく。

```
$ python receive.py
 [*] Waiting for messages. To exit press CTRL+C
```

プロデューサを実行する。

```
$ python send.py
 [x] Sent 'Hello World!'
```

すると、コンシューマ側にもメッセージが送られる。

```
 [x] Received b'Hello World!'
```
