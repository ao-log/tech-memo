
[aws/amazon-ecs-agent](https://github.com/aws/amazon-ecs-agent)

#### proposals ディレクトリ下

https://github.com/aws/amazon-ecs-agent/blob/master/proposals/awsvpc-task-metadata.md

* pause container の起動時に ecs-ipam plugin によって割り当てられる IPv4 アドレスをキャッシュする。
* タスク内のコンテナは GET 169.254.170.2/v2/metadata API を発行する。
* このリクエストは ecs bridge のポート 51679 を介して ECS Agent に届く。
* ECS Agent はキャッシュ内からコンテナに割り当てるアドレスを探す。
* ECS Agent はコンテナにメタデータを提供する。

https://github.com/aws/amazon-ecs-agent/blob/master/proposals/eni.md

* bridge ネットワーキングモードの問題点
  * 各コンテナが docker0 ブリッジインタフェースをシェアするのでパフォーマンス面が課題となる
  * プライマリネットワークインタフェースのセキュリティグループがタスクに使用される

ECS Agent は ENI を dedicated な namespace に割り当てる。
ECS Agent は namespace 内のコンテナ用に credential endpoint へのルートを作成する。

プライマリネットワークインタフェースは docker0 に使用される。

セカンダリネットワークインタフェースの eth1 は ENI-1 ネットワークネームスペースに移される。タスク内のコンテナはこの namespace を共有する。
タスクの IP アドレスはセカンダリネットワークインタフェースによって割り当てられる。

ecs-eni-br ブリッジはコンテナによって使用される。eth0 と通信し credential、メタデータにアクセスするため。
ENI-1 ネームスペースと ecs-eni-br ブリッジ間の veth ペアはこの目的で使用される。

ECS Agent は pause コンテナを上記構成のために起動する。このコンテナは pause() POSIX API を実行するだけのものである。
ECS Agent は CNI Plugin を使用して ENI をこのコンテナのネームスペースに移動する。

**CNI plugin の動作**

1. ENI の MAC アドレスを取得
1. ENI デバイス名を取得
1. network gateway のマスクを取得
1. ENI のプライマリ IP アドレスを取得
1. ENI をコンテナの namespace に移動
1. インタフェースにプライマリ IP アドレスを付与
1. gateway 経由のインターネットへのルートを作成
1. ルートテーブルからのエントリーの削除
1. dhcp クライアントの開始

**credential へのルート確立**

1. 使用可能な IP アドレスを 169.254.0.0/16 から決定
1. ecs-eni ブリッジを作成
* veth ペアを作成
1. veth インタフェースを ecs-eni ブリッジに割り当て
1. veth インタフェースに IP アドレスを割り当て、コンテナのネームスペースに移動
1. 169.254.170.2 のルートをコンテナネームスペース内に作成


#### Amazon ECS CNI Plugins 

https://github.com/aws/amazon-ecs-cni-plugins/tree/55b2ae77ee0bf22321b14f2d4ebbcc04f77322e1

