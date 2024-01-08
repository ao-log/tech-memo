
[Networking overview](https://docs.docker.com/network/)


[Understand container communication](https://gdevillele.github.io/engine/userguide/networking/default_network/container-communication/)

[コンテナ通信の理解](https://docs.docker.jp/v17.06/engine/userguide/networking/default_network/container-communication.html)

* `net.ipv4.conf.all.forwarding` が 1 の場合は、IP フォワーディングが有効になる。この場合のみコンテナ間通信が可能
* コンテナ間の通信
  * 各コンテナは docker0 ブリッジに接続される
  * `--link=コンテナ名_または_ID:エイリアス` により iptables の ACCEPT ルールのペアが作成される


[Docker コンテナ・ネットワークの理解](https://docs.docker.jp/engine/userguide/networking/dockernetworks.html)


