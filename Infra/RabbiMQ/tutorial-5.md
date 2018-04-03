
## このチュートリアルの概要

このページは RabbitMQ のチュートリアルを元にしている。

https://www.rabbitmq.com/tutorials/tutorial-five-python.html

前章の方法では、複数の基準に基づいたルーティングができなかった。

この章では複数の基準に従ってメッセージをルーティングできるようにする。
例としては、syslogがseverity (info/warn/crit...) と facility (auth/cron/kern...)でログを振り分けるのと同じことを出来るようにする。
ここでは、topic exchange を導入する。

## この章で学べること

* topic exchange では、ルーティングキーはドットで区切られた複数ワードにする。例えば、"kern.info"

* binding key も同じ形式。2種類の特殊な記号を使うことができる。
  * (*) は 1 ワードを表している
  * (#) はゼロ個、もしくは複数個のワードを表している。

## ソース

###### emit_log_topic.py

```python
#!/usr/bin/env python
import pika
import sys

connection = pika.BlockingConnection(pika.ConnectionParameters(host='localhost'))
channel = connection.channel()

channel.exchange_declare(exchange='topic_logs',
                         exchange_type='topic')

routing_key = sys.argv[1] if len(sys.argv) > 2 else 'anonymous.info'
message = ' '.join(sys.argv[2:]) or 'Hello World!'
channel.basic_publish(exchange='topic_logs',
                      routing_key=routing_key,
                      body=message)
print(" [x] Sent %r:%r" % (routing_key, message))
connection.close()
```

###### receive_logs_topic.py

```python
#!/usr/bin/env python
import pika
import sys

connection = pika.BlockingConnection(pika.ConnectionParameters(host='localhost'))
channel = connection.channel()

channel.exchange_declare(exchange='topic_logs',
                         exchange_type='topic')

result = channel.queue_declare(exclusive=True)
queue_name = result.method.queue

binding_keys = sys.argv[1:]
if not binding_keys:
    sys.stderr.write("Usage: %s [binding_key]...\n" % sys.argv[0])
    sys.exit(1)

for binding_key in binding_keys:
    channel.queue_bind(exchange='topic_logs',
                       queue=queue_name,
                       routing_key=binding_key)

print(' [*] Waiting for logs. To exit press CTRL+C')

def callback(ch, method, properties, body):
    print(" [x] %r:%r" % (method.routing_key, body))

channel.basic_consume(callback,
                      queue=queue_name,
                      no_ack=True)

channel.start_consuming()
```


## 実行

##### topic_logs exchange の全てのメッセージを受け取る場合。

```
 $ python receive_logs_topic.py "#"
```

プロデューサを実行する。

```
$ python emit_log_topic.py "kern.critical" "A critical kernel error"
 [x] Sent 'kern.critical':'A critical kernel error'
$ python emit_log_topic.py "kern.info" "A info"
 [x] Sent 'kern.info':'A info'
```

コンシューマ側で全てのメッセージを受信する。

```
 $ python receive_logs_topic.py "#"
  [*] Waiting for logs. To exit press CTRL+C
  [x] 'kern.critical':b'A critical kernel error'
  [x] 'kern.info':b'A info'
```

##### kern.* のメッセージのみ受け取る場合。

```
 $ python receive_logs_topic.py "kern.*"
```

プロデューサを実行する。

```
$ python emit_log_topic.py "kern.critical" "A critical kernel error"
 [x] Sent 'kern.critical':'A critical kernel error'
$ python emit_log_topic.py "auth.critical" "A critical auth error"
 [x] Sent 'auth.critical':'A critical auth error'
```

コンシューマ側では kern.* のメッセージのみ受け取っている。

```
$ python receive_logs_topic.py "kern.*"
 [*] Waiting for logs. To exit press CTRL+C
 [x] 'kern.critical':b'A critical kernel error'
```

##### *.critical のメッセージのみ受け取る場合。

```
 $ python receive_logs_topic.py "*.critical"
```

プロデューサを実行する。

```
$ python emit_log_topic.py "kern.critical" "A critical kernel error"
 [x] Sent 'kern.critical':'A critical kernel error'
$ python emit_log_topic.py "auth.critical" "A critical auth error"
 [x] Sent 'auth.critical':'A critical auth error'
$ python emit_log_topic.py "auth.info" "A auth info"
 [x] Sent 'auth.info':'A auth info
```

コンシューマ側では *.critical のメッセージのみ受け取っている。

```
$ python receive_logs_topic.py "*.critical"
 [*] Waiting for logs. To exit press CTRL+C
 [x] 'kern.critical':b'A critical kernel error'
 [x] 'auth.critical':b'A critical auth error'
```

##### binding 状況

```
$ sudo rabbitmqctl list_bindings
...
topic_logs      exchange        amq.gen-IKWs5j5OWmF2BPYCS4PKeA  queue   #       []
```

