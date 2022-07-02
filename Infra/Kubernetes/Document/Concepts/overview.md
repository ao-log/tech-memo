
[Overview](https://kubernetes.io/docs/concepts/overview/)

[Kubernetesのコンポーネント](https://kubernetes.io/ja/docs/concepts/overview/components/)

コントロールプレーン
* kube-apiserver
* etcd
* kube-scheduler
* kube-controller-manager
* cloud-controller-manager

ノード
* kubelet
* kube-proxy
  * ノード上のネットワークルールを管理。iptables を操作している。
* container runtime


[ノード](https://kubernetes.io/ja/docs/concepts/architecture/nodes/)

ノードは Kubelet によってコントロールプレーンに登録される。手動で登録することも可能。

ノードのステータスは以下の 5 種類。

* Ready
* DiskPressure
* MemoryPressure
* PIDPressure
* NetworkUnavailable

Ready condition が pod-eviction-timeout に設定された時間を超えて Unknown, False になっている場合、ノードコントローラによって削除がスケジュールされる。eviction が実施され eviction のタイムアウトは 5 分。

また、ノードコントローラは問題が発生したノードに対して状態を表す taint を付与する。

kubelet は .status および Lease オブジェクトの作成、更新を行う。.status の更新間隔は 5 分間。

ノードコントローラはハートビートの受信を停止した場合などに NodeStatus の Ready コンディションを Unknown に変更する。Unknown に設定する際のハートビートの未受信時間は node-monitor-grace-period で設定されておりデフォルトで 40 秒。


