
[Configure logging drivers](https://docs.docker.com/config/containers/logging/configure/)

* デフォルトでは `json-file` ロギングドライバが使用される。ログローテションされないため、ディスク容量消費に注意
* local ロギングドライバはデフォルトでログローテーションするのでおすすめ
* `docker run` の `--log-driver` オプションで任意のロギングドライバを指定可能
* ログの配送モードはデフォルトは直接コンテナからドライバに送信する。non-blocking を選択することもできる
* `max-buffer-size` で最大バッファサイズを調整可能


[Amazon CloudWatch Logs logging driver](https://docs.docker.com/config/containers/logging/awslogs/)

* `daemon.json` に以下のような設定が必要
```json
{
  "log-driver": "awslogs",
  "log-opts": {
    "awslogs-region": "us-east-1"
  }
}
```
* オプション
  * awslogs-region
  * awslogs-endpoint
  * awslogs-group
  * awslogs-stream
  * awslogs-create-group
  * awslogs-datetime-format
  * awslogs-multiline-pattern
  * tag
  * awslogs-force-flush-interval-seconds
  * awslogs-max-buffered-events


[Fluentd logging driver](https://docs.docker.com/config/containers/logging/fluentd/)


