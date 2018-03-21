
## このチュートリアルの概要

このページは RabbitMQ のチュートリアルを元にしている。

https://www.rabbitmq.com/tutorials/tutorial-two-python.html

この章では、複数のコンシューマに対して、順にメッセージを送ることにフォーカスする。

ユースケースとしては、重い複数の処理があって一台のサーバでさばき切れない場合。
そのような場合は、複数サーバ上でワーカーを起動し、RabbitMQ が各ワーカーにメッセージを順に送信する。ワーカーはそのメッセージを受け取って処理を行う…という流れを実現できる。

## この章で学べること

* RabbitMQ はメッセージが正しくコンシューマに届いたかどうかをコンシューマからの Ack によって確認する。
Ack が返ってこない場合メッセージをリキューする。Ack 機能はデフォルトでは有効になっている（前の章では no_ack=True にして無効化していた）。

* RabbitMQ を停止するとデフォルトではキューが消える。残すには durable を有効化する。

```
channel.queue_declare(queue='task_queue', durable=True)
```

* 既に存在するキューの設定を変えることはできない。

* メッセージを永続化するには「delivery_mode」を 2 にする。
この設定によりメッセージをディスクに書き込むようになる。
ただ、メッセージ到着のたびに即座に書き込まない点には注意が必要。キャッシュにためてからディスクに書き込む動作となる。

```
channel.basic_publish(exchange='',
                      routing_key="task_queue",
                      body=message,
                      properties=pika.BasicProperties(
                         delivery_mode = 2, # make message persistent
                      ))
```

* QoS を設定することが出来る。prefetch=1 にすることで、一度に 1 つより多くのメッセージをワーカーに送らないようにする。
言い換えると、前のメッセージの Ack が帰ってくるまでは次のメッセージをワーカーに送らない。代わりにビジーではない次のワーカーにメッセージを送る。

```
channel.basic_qos(prefetch_count=1)
```


## ソース

###### new_task.py

```python
#!/usr/bin/env python
import pika
import sys

connection = pika.BlockingConnection(pika.ConnectionParameters(host='localhost'))
channel = connection.channel()

channel.queue_declare(queue='task_queue', durable=True)

message = ' '.join(sys.argv[1:]) or "Hello World!"
channel.basic_publish(exchange='',
                      routing_key='task_queue',
                      body=message,
                      properties=pika.BasicProperties(
                         delivery_mode = 2, # make message persistent
                      ))
print(" [x] Sent %r" % message)
connection.close()
```

###### worker.py

```python
#!/usr/bin/env python
import pika
import time

connection = pika.BlockingConnection(pika.ConnectionParameters(host='localhost'))
channel = connection.channel()

channel.queue_declare(queue='task_queue', durable=True)
print(' [*] Waiting for messages. To exit press CTRL+C')

def callback(ch, method, properties, body):
    print(" [x] Received %r" % body)
    time.sleep(body.count(b'.'))
    print(" [x] Done")
    ch.basic_ack(delivery_tag = method.delivery_tag)

channel.basic_qos(prefetch_count=1)
channel.basic_consume(callback,
                      queue='task_queue')

channel.start_consuming()
```


## 実行

2 つのウィンドウを用意し、それぞれでワーカーを起動させておく。
ワーカーは time.sleep(body.count(b'.'))　の行を修正し、ウィンドウ 1 では 2 秒の sleep、ウィンドウ 2 は 5 秒の sleep を実行する。
sleep 秒数の異なる ２ つのワーカーが次のメッセージを受信する様子を観察する。

```
// ウィンドウ 1　（2 秒間隔）
$ python worker-2sec.py 
 [*] Waiting for messages. To exit press CTRL+C

// ウィンドウ 2　（5 秒間隔）
$ python worker-5sec.py 
 [*] Waiting for messages. To exit press CTRL+C
```

プロデューサを実行する。

```
$ for i in `seq 10`; do python new_task.py $i; done
 [x] Sent '1'
 [x] Sent '2'
 [x] Sent '3'
 [x] Sent '4'
 [x] Sent '5'
 [x] Sent '6'
 [x] Sent '7'
 [x] Sent '8'
 [x] Sent '9'
 [x] Sent '10'
```

すると、それぞれのワーカーにメッセージが送られる。


```
// ウィンドウ 1　（2 秒間隔）
$ python worker-2sec.py 
 [*] Waiting for messages. To exit press CTRL+C
 [x] Received b'2'
 [x] Done
 [x] Received b'3'
 [x] Done
 [x] Received b'4'
 [x] Done
 [x] Received b'6'
 [x] Done
 [x] Received b'7'
 [x] Done
 [x] Received b'9'
 [x] Done
 [x] Received b'10'
 [x] Done


// ウィンドウ 2　（5 秒間隔）
$ python worker-5sec.py 
 [*] Waiting for messages. To exit press CTRL+C
 [x] Received b'1'
 [x] Done
 [x] Received b'5'
 [x] Done
 [x] Received b'8'
 [x] Done
```
