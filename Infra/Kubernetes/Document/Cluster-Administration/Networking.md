
[クラスターのネットワーク](https://kubernetes.io/ja/docs/concepts/cluster-administration/networking/)

すべての Pod は独自のIPアドレスを持つ。
Pod 間通信に関して重要な性質。

#### AWS VPC CNI for Kubernetes

* VPC ネットワークと同じ IP アドレスを Pod に割り当てることができる。
* ENI をノードに割り当て、Pod に ENI のセカンダリ IP アドレス範囲から割り当てる。


