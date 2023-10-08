
[Docker run reference](https://docs.docker.com/engine/reference/run/)


-d を付与するとバックグラウンド。
```
$ docker run -d -p 80:80 my_image nginx -g 'daemon off;'
```

-d なしだとフォアグラウンド。
```
$ docker run -a stdin -a stdout -i -t ubuntu /bin/bash
```

#### Network Settings

```
--network="bridge" : Connect a container to a network
                      'bridge': create a network stack on the default Docker bridge
                      'none': no networking
                      'container:<name|id>': reuse another container's network stack
                      'host': use the Docker host network stack
```

**Network: bridge**
```
With the network set to bridge a container will use docker’s default networking setup. A bridge is setup on the host, commonly named docker0, and a pair of veth interfaces will be created for the container. One side of the veth pair will remain on the host attached to the bridge while the other side of the pair will be placed inside the container’s namespaces in addition to the loopback interface. An IP address will be allocated for containers on the bridge’s network and traffic will be routed though this bridge to the container.

Containers can communicate via their IP addresses by default. To communicate by name, they must be linked.
```
* docker0 ブリッジインタフェースを使用する。
* veth のペアを作成し片方はブリッジに、もう片方はコンテナのネームスペース内に配置する。

**Network: host**
bridge と比べるとパフォーマンス面で優位。bridge は仮想化のレイヤがあるため。ネットワークパフォーマンスがクリティカルな要件で採用が推奨される。

#### Memory

300 M までのメモリを使用できる。また、300 M のスワップメモリも使用できる。
```
$ docker run -it -m 300M ubuntu:14.04 /bin/bash
```

```--memory-reservation``` はソフトリミット。超えて使用することも可能。
```
$ docker run -it -m 500M --memory-reservation 200M ubuntu:14.04 /bin/bash
```

```--oom-kill-disable``` を指定することにより、カーネルの OOM Killer 対象外となる。
```
$ docker run -it -m 100M --oom-kill-disable ubuntu:14.04 /bin/bash
```

#### CPU

