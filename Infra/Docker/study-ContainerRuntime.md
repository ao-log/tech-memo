
## 書籍 『イラストでわかる Docker と Kubernetes』

* イメージはレイヤ構造。コンテナ実行時にはレイヤを重ねた上に更に一枚読み書き可能なレイヤを作成する。
* レジストリもレイヤ単位でアップロード、ダウンロードする API を持っている。
* Docker のストレージドライバでレイヤを重ね合わせてコンテナのルートファイルシステムとして利用できるようにする。
  * overlay2 ではレイヤの tar ファイルを展開し、overlay ファイルシステムで重ね合わせている

docker save コマンドでイメージを tar 形式で出力できる。

```
$ tree
.
├── 2bcb04bdb83f7c5dc30f0edaca1609a716bda1c7d2244d4f5fbbdfef33da366c.json ← 実行コマンドや環境変数などの情報
├── 97f043d241d89a3fc54ba1f0dfb14fe804914ed7eadbef9dc7ceafafd7d75b72
│   ├── VERSION
│   ├── json
│   └── layer.tar
├── 988f79af9c1796b4dc7c4fafd6a36da69a8581bde5290d778ac1f4784fe45c77
│   ├── VERSION
│   ├── json
│   └── layer.tar
├── d9445f143133d24dd1e85ba979298e2b80465fe47dad90b469d2a07656b85d48
│   ├── VERSION
│   ├── json
│   └── layer.tar
├── manifest.json ← イメージの構成などの情報
└── repositories ← イメージの構成などの情報
```

* dockerd は HTTP API 経由で指示を受け付ける。dockerd から低レベルランタイム(バイナリ)を実行しコンテナに対する各種操作を行う。
* 高レベルランタイム(CRI) は kubelet やユーザの指示を受けて、イメージのプル、ルートファイルシステムの用意や、CNI プラグインによるネットワークの管理、霊レベルランタイムに対するコンテナ管理を行う。
* 低レベルランタイム(OCI) は高レベルランタイムから指示を受けて隔離された実行環境をコンテナとして作り出す。
* 低レベルランタイムを呼び出すときは shim をはさむ場合がある。アダプタ的な役割。
* Filesystem bundle
  * OCI Runtime Specification でコンテナのライフサイクルやコンテナのもととなる Filesystem bundle を仕様として定義している。
  * ルートファイルシステム、実行環境の設定ファイルが必要。
* OCI Image Specification
  * コンテナイメージの標準仕様
  * マニフェストにレイヤや構成などを設定。
  * 各レイヤ
  * Configuration。コンテナの各種情報（エントリポイント、環境変数、レイヤのリストなど）。Filesystem bundle を作成する際に使用。



## ウェブ調査

[コンテナユーザなら誰もが使っているランタイム「runc」を俯瞰する[Container Runtime Meetup #1発表レポート]](https://medium.com/nttlabs/runc-overview-263b83164c98)

以下のような階層構造

* kubelet などの API を呼び出す主体
* 高レベルランタイム(CRI ランタイム)。docker(docker shim 経由)、containerd, cri-o など。
* 低レベルランタイム(OCI ランタイム)。runcm gVisor など。

#### OCI ランタイム

* その仕様が OCI により標準として定められている。

以下の標準仕様を定めている。

* Image Specification：コンテナイメージの標準仕様。
* Runtime Specification：コンテナランタイムの標準仕様。
* Distribution Specification：コンテナレジストリの標準仕様。

#### OCI Runtime Specification

コンテナに対して行うことができる操作(create, start, kill, delete) とライフサイクルが定義されている。

コンテナのもととなる Filesystem bundle が必要。更に環境設定ファイルも必要。

runc run は create と start を実行できるコマンド。

runc run 内で runc init が呼ばれている。runc init は隔離環境を作成し、runc run から支持されたタイミングでエントリポイントのプログラムを実行する。



[コンテナランタイムの仕組みと、Firecracker、gVisor、Unikernelが注目されている理由。 Container Runtime Meetup #2](https://www.publickey1.jp/blog/20/firecrackergvisorunikernel_container_runtime_meetup_2.html)

* kubelet が CRI をつかって containerd とおしゃべりをして、その後ろにいる runc と OCI でやりとりをしてコンテナを生成



[Kubernetes、Dockerに依存しないKubernetes用の軽量コンテナランタイム「cri-o」正式版1.0リリース](https://www.publickey1.jp/blog/17/kubernetesdockerkubernetescri-o10.html)



[Deep Dive into Runtime Shim](https://speakerdeck.com/moricho/deep-dive-into-runtime-shim?slide=10)

* Runtime Shim とはコンテナプロセスと高レベルランタイムの間のコミュニケーションを取り持つ API





## ソース情報

## Open Container Initiative

[Open Container Initiative](https://opencontainers.org/)



#### Image Specification

[Open Container Initiative](https://github.com/opencontainers/image-spec/blob/master/spec.md)

* This specification defines an OCI Image, consisting of a manifest, an image index (optional), a set of filesystem layers, and a configuration.


[Open Container Initiative](https://ja.wikipedia.org/wiki/Open_Container_Initiative)



#### Runtime Specification

[Open Container Initiative Runtime Specification](https://github.com/opencontainers/runtime-spec/blob/master/spec.md)

[Filesystem Bundle](https://github.com/opencontainers/runtime-spec/blob/master/bundle.md)

次の 2 点が必要。

* config.json
* ルートファイルシステム

[Runtime and Lifecycle](https://github.com/opencontainers/runtime-spec/blob/master/runtime.md)

コンテナが持っているべき状態が state に記述されている。

**ライフサイクル**

1. create コマンドを実行
1. config.json に基づいてコンテナランタイム環境を作成。
1. prestart フックを実行。
1. createRuntime フックを実行。
1. reateContainer フックを実行。
1. startContainer フックを実行。
1. process で指定されたユーザ指定のプログラムを実行。
1. posttart フックを実行。
1. コンテナプロセスが終了
1. ランタイムの delete コマンドが実行される。
1. poststop フックを実行。

以下のオペレーションをサポートする必要がある。

* start <container-id>: コンテナの状態を返す。
* create <container-id> <path-to-bundle>: config.json の process 以外のプロパティが適用される。
* start <container-id>: process で指定されたユーザ指定のプログラムを実行。
* kill <container-id> <signal>: コンテナにシグナルを送る。
* delete <container-id>: コンテナの削除。 


[Configuration](https://github.com/opencontainers/runtime-spec/blob/master/config.md)

config.json の形式。以下は設定項目の一部。

* ociVersion
* root: ルートファイルシステム
* mounts
* process: start 時に必要。
  * terminal
  * cwd: working directory
  * env
  * args
  * commandLine
  * user
* POSIX process
  * rlimits
* Linux Process
  * apparmorProfile
  * capabilities
  * noNewPrivileges
  * oomScoreAdj
  * selinuxLabel
* Hostname
* hooks
* annotations
* linux
  * devices
  * uidMappings
  * gidMappings
  * sysctl
  * cgroupsPath
  * resources
    * network
    * pids
    * hugepageLimits
    * memory
    * cpu
    * devices
    * blockIO
  * namespaces

[Linux Container Configuration](https://github.com/opencontainers/runtime-spec/blob/master/config-linux.md)

Linux の設定項目がまとめられている。



#### runc

[runc](https://github.com/opencontainers/runc)

OCI バンドルの形式で用意が必要。

```shell
# create the top most bundle directory
mkdir /mycontainer
cd /mycontainer

# create the rootfs directory
mkdir rootfs

# export busybox via Docker into the rootfs directory
docker export $(docker create busybox) | tar -C rootfs -xvf -
```

```runc spec``` により config.spec のテンプレートを作成可能。

```runc run``` によりコンテナを実行。




#### namespace

[namespaces(7) — Linux manual page](https://man7.org/linux/man-pages/man7/namespaces.7.html)

リソースを分離する技術。以下のリソースタイプに対応。

* Cgroup
* IPC
* Network
* Mount
* PID
* Time
* User
* UTS

The namespaces API

* clone: 新しい子プロセスを作成。fork との違いは親プロセスのコンテキストの一部を共有できる。
* setns: プロセスを存在する namespace に属させる。
* unshare: プロセスを新しい namespace に移動する。
* ioctl: 


[44.7. 名前空間とは](https://access.redhat.com/documentation/ja-jp/red_hat_enterprise_linux/8/html/system_design_guide/what-namespaces-are_setting-limits-for-applications)


* プロセスがどの名前空間に所属するかを確認するには、/proc/<PID>/ns/ ディレクトリーのシンボリックリンクを確認


#### cgroup

[第1章 コントロールグループの概要 (CGROUP)](https://access.redhat.com/documentation/ja-jp/red_hat_enterprise_linux/7/html/resource_management_guide/chap-introduction_to_control_groups)

* cgroup を使用することにより、システム管理者は、システムリソースの割り当て、優先順位付け、拒否、管理および監視におけるより詳細なレベルの制御を行うことができる
* ハードウェアリソースはアプリケーションおよびユーザー間でスマートに分割することができる。

/sys/fs/cgroup/ に制御可能なリソースの一覧がある。



[第4章 カーネル機能](https://access.redhat.com/documentation/ja-jp/red_hat_enterprise_linux/7/html/kernel_administration_guide/kernel_features)

* コントロールグループ は、名前空間 が有効にしたリソースの使用量を制限




#### unchare コマンド

[unshare(1) — Linux manual page](https://man7.org/linux/man-pages/man1/unshare.1.html)

リソースの隔離に便利なコマンド。以下の namespace を作成できる。

* mount namespace
* UTS namespace
* IPC namespace
* network namespace
* PID namespace
* cgroup namespace
* user namespace
* time namespace



#### docker

[docker export](https://docs.docker.com/engine/reference/commandline/export/)

Export a container’s filesystem as a tar archive.

[docker save](https://docs.docker.com/engine/reference/commandline/save/)

Save one or more images to a tar archive (streamed to STDOUT by default)

export は単にルートファイルシステムとして出力。
save はメタ情報等含めて出力。



## その他参考資料

[Demystifying Containers - Part I: Kernel Space](https://medium.com/@saschagrunert/demystifying-containers-part-i-kernel-space-2c53d6979504)


