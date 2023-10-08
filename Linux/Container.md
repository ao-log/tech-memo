
[コンテナの仕組みとLinuxカーネルのコンテナ機能［1］名前空間とは？](https://gihyo.jp/admin/serial/01/linux_containers/0002)

* Namespace
  * mount, UTS, PID, IPC, User, Network の名前空間がある
  * デフォルトでは全てのプロセスはデフォルトの名前空間に属する


[Linuxカーネルのコンテナ機能［2］ ─cgroupとは？（その1）](https://gihyo.jp/admin/serial/01/linux_containers/0003)

* cgroup
  * CPU, メモリなどのリソースについてグループごとに制限をかけられる
  * cgroupfs を通して操作する
  * `/sys/fs/cgroup/cpu` の下に `test01` グループを作成するような使用方法。`/sys/fs/cgroup/cpu/test01/tasks` に PID を記述することで、当該 PID が `test01` グループに属する


[Linuxカーネルのコンテナ機能［5］ ─ネットワーク](https://gihyo.jp/admin/serial/01/linux_containers/0006)

* コンテナではネットワーク名前空間を作成し、ホスト上のネットワークインタフェースをネットワーク名前空間に割り当てる使用方法が多い
* veth は仮想的なネットワークインタフェース。2 つのインタフェースがペアで作成されインタフェース間で通信を行うことができる。片方をホスト側のブリッジに接続し、もう片方をコンテナのネットワーク名前空間に割り当てる。コンテナ側のネットワークインタフェースは eth0 のようなインタフェースとして見える。veth は互いに異なる名前空間にないと通信できない


[Linuxカーネルのコンテナ機能 ― cgroupの改良版cgroup v2［1］](https://gihyo.jp/admin/serial/01/linux_containers/0037)

* cgroup v1 の課題がまとめられている


[Linuxカーネルのコンテナ機能 ― cgroupの改良版cgroup v2 ［2］](https://gihyo.jp/admin/serial/01/linux_containers/0038)

* `cgroup.procs` ファイルでプロセスを登録する


