
[AWS Fargateとは?](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/userguide/what-is-fargate.html)

[Fargate タスク定義の考慮事項](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/userguide/fargate-task-defs.html)

制約、設定可能なタスクサイズなどがまとめられている。ネットワークモードは awsvpc のみ。タスクストレージはプラットフォームバージョンごとに異なる。


[タスクでのデータボリュームの使用](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/userguide/using_data_volumes.html)


[Fargate タスクネットワーキング](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/userguide/fargate-task-networking.html)

* ENI が提供される。パブリック IP アドレスも付与可能。
* コンテナ間は localhost で通信可能。
* PV 1.4.0 ではタスク ENI が割り当てられる。1.3.0 では更に Fargate ENI が割り当てられ ECR ログイン情報、Secrets の取得は Fargate ENI が使用される。


[AWS Fargate タスクのメンテナンス](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/userguide/task-maintenance.html)

* サービスから起動されたタスクの場合、ホストの問題である場合は Health Dashboard には通知されない。


[Amazon ECSタスクメタデータエンドポイント](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/userguide/task-metadata-endpoint-fargate.html)

* タスクメタデータエンドポイントバージョン 4 はプラットフォームバージョン 1.4.0 以降を使用するタスクで利用可能。
* ```${ECS_CONTAINER_METADATA_URI_V4}/stats``` のようなパスに対してクエリできる。


