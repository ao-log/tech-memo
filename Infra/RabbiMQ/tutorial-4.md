
## このチュートリアルの概要

このページは RabbitMQ のチュートリアルを元にしている。

https://www.rabbitmq.com/tutorials/tutorial-four-python.html

この章ではメッセージの内容に応じて、別々のキューにメッセージを送る方法に焦点をあてる。
例えばログの深刻度に応じて、送付先のキューを変える用途となる。

## この章で学べること

##### プロデューサ

* メッセージの種別によって宛先キューを変えるには、exchange type の direct が適している。binding_key にマッチしたキューにのみ送信する。

```
channel.exchange_declare(exchange='direct_logs',
                         exchange_type='direct')
```

* ルーティング先は、routing_key (binding_key のこと)によって与える。

```
channel.basic_publish(exchange='direct_logs',
                      routing_key=severity,
                      body=message)
```

##### コンシューマ

* コンシューマ側では、どの exchange, routing_key のメッセージを受信するかを指定し、キュー名と関連付ける。

```
result = channel.queue_declare(exclusive=True)
queue_name = result.method.queue

for severity in severities:
    channel.queue_bind(exchange='direct_logs',
                       queue=queue_name,
                       routing_key=severity)
```

## ソース

###### emit_log_direct.py

```python
#!/usr/bin/env python
import pika
import sys

connection = pika.BlockingConnection(pika.ConnectionParameters(host='localhost'))
channel = connection.channel()

channel.exchange_declare(exchange='direct_logs',
                         exchange_type='direct')

severity = sys.argv[1] if len(sys.argv) > 2 else 'info'
message = ' '.join(sys.argv[2:]) or 'Hello World!'
channel.basic_publish(exchange='direct_logs',
                      routing_key=severity,
                      body=message)
print(" [x] Sent %r:%r" % (severity, message))
connection.close()
```

###### receive_logs_direct.py

```python
#!/usr/bin/env python
import pika
import sys

connection = pika.BlockingConnection(pika.ConnectionParameters(host='localhost'))
channel = connection.channel()

channel.exchange_declare(exchange='direct_logs',
                         exchange_type='direct')

result = channel.queue_declare(exclusive=True)
queue_name = result.method.queue

severities = sys.argv[1:]
if not severities:
    sys.stderr.write("Usage: %s [info] [warning] [error]\n" % sys.argv[0])
    sys.exit(1)

for severity in severities:
    channel.queue_bind(exchange='direct_logs',
                       queue=queue_name,
                       routing_key=severity)

print(' [*] Waiting for logs. To exit press CTRL+C')

def callback(ch, method, properties, body):
    print(" [x] %r:%r" % (method.routing_key, body))

channel.basic_consume(callback,
                      queue=queue_name,
                      no_ack=True)

channel.start_consuming()
```


## 実行

2 つのウィンドウを用意し、それぞれでコンシューマを起動させておく。片方は error メッセージ用、もう片方は warning, info メッセージ用である。

```
// ウィンドウ 1 （error メッセージ用）
$ python receive_logs_direct.py error
 [*] Waiting for logs. To exit press CTRL+C

// ウィンドウ 2 （warning, info メッセージ用）
$ python receive_logs_direct.py warning info
 [*] Waiting for logs. To exit press CTRL+C
```

プロデューサを実行する。

```
$ python emit_log_direct.py error "error message"
 [x] Sent 'error':'error message'
$ python emit_log_direct.py warning "warning message"
 [x] Sent 'warning':'warning message'
$ python emit_log_direct.py info "info message"
 [x] Sent 'info':'info message'
```

すると、それぞれにマッチしたメッセージが送られる。

```
// ウィンドウ 1 （error メッセージ用）
$ python receive_logs_direct.py error
 [*] Waiting for logs. To exit press CTRL+C
 [x] 'error':b'error message'

// ウィンドウ 2 （warning, info メッセージ用）
$ python receive_logs_direct.py warning info
 [*] Waiting for logs. To exit press CTRL+C
 [x] 'warning':b'warning message'
 [x] 'info':b'info message'
```

binding 状況を確認する。

```
$ sudo rabbitmqctl list_bindings
...
direct_logs     exchange        amq.gen-fOZ0e3n5MpGcg16lCKNdyg  queue   error   []
direct_logs     exchange        amq.gen-QAQg3jYx2vJJw1hiQqR9jw  queue   info    []
direct_logs     exchange        amq.gen-fOZ0e3n5MpGcg16lCKNdyg  queue   warning []
```
