
[Kubernetes アプリケーションの公開 Part1: Service と Ingress リソース](https://aws.amazon.com/jp/blogs/news/exposing-kubernetes-applications-part-1-service-and-ingress-resources/)

* Service
  * ClusterIP: クラスター内部でのみアクセス可能な仮想 IP アドレスを割り当て
  * NodePort: 各クラスターノード上の静的ポートを介して上記の ClusterIP を公開
  * LoadBalancer: 自動的に ClusterIP を作成し、NodePort を設定し、クラスターのインフラストラクチャ環境 (クラウドプロバイダーなど) に Service のバックエンドの Pod を公開するための負荷分散コンポーネントを作成することを指示
  * kube-proxy により iptables により設定される
  * 各 Service の ClusterIP には、<service-name>.<namespace-name>.svc.cluster.local 形式のクラスター内部でのみアクセス可能な DNS 名が提供される
  * Service コントローラーは、新しい Service リソースが作成されるのを監視し、spec.type が LoadBalancer の場合、コントローラーはクラウドプロバイダーの API を使用してロードバランサーをプロビジョニング
  * Service ごとにロードバランサーを作るのは非効率。Ingress で統合できる。例えば ALB の場合、URL パスごとに転送先の Service を指定できる



[AWS Load Balancer Controller を使った Blue/Green デプロイメント、カナリアデプロイメント、A/B テスト](https://aws.amazon.com/jp/blogs/news/using-aws-load-balancer-controller-for-blue-green-deployment-canary-deployment-and-a-b-testing/)

以下のようなマニフェストにより Blue/Green デプロイができる。weight を調整することで徐々に重みづけを変えていくことも可能。

```json
alb.ingress.kubernetes.io/actions.blue-green: |
  {
     "type":"forward",
     "forwardConfig":{
       "targetGroups":[
         {
           "serviceName":"hello-kubernetes-v1",
           "servicePort":"80",
           "weight":0
         },
         {
           "serviceName":"hello-kubernetes-v2",
           "servicePort":"80",
           "weight":100
         }
       ]
     }
   }
```


[Amazon VPC CNI による Kubernetes NetworkPolicy のサポート](https://aws.amazon.com/jp/blogs/news/amazon-vpc-cni-now-supports-kubernetes-network-policies/)

* Amazon VPC CNI により NetworkPolicy がネイティブサポートされるようになった
* 従来 iptables が広く採用されていたが、管理やルール数の増加によるパフォーマンスへの影響などの課題があった
* eBPF によるアプローチを採用
* Amazon VPC CNI の最新バージョンでは、クラスター内の全てのノードに CNI バイナリと ipamd プラグインと共に、Node Agent もインストールされる。aws-node の DaemonSet のコンテナとして稼働



[Amazon EKS での Kubernetes アップグレードの計画](https://aws.amazon.com/jp/blogs/news/planning-kubernetes-upgrades-with-amazon-eks/)

[Amazon EKS が Kubernetes 1.22 のサポートを開始](https://aws.amazon.com/jp/blogs/news/amazon-eks-now-supports-kubernetes-1-22/)

[Amazon EKS アドオンで Amazon EBS CSI ドライバーが一般利用可能になりました](https://aws.amazon.com/jp/blogs/news/amazon-ebs-csi-driver-is-now-generally-available-in-amazon-eks-add-ons/)

[【レポート】Amazon EKS と Datadog によるマイクロサービスの可観測性（パートナーセッション） #PAR-03 #AWSSummit](https://dev.classmethod.jp/articles/aws_summit_japan_2021_datadog/)


[Amazon EKS で GitOps パイプラインを構築する](https://aws.amazon.com/jp/blogs/news/building-a-gitops-pipeline-with-amazon-eks/)



