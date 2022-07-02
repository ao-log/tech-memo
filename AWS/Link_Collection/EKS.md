
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



[Amazon EKS での Kubernetes アップグレードの計画](https://aws.amazon.com/jp/blogs/news/planning-kubernetes-upgrades-with-amazon-eks/)

[Amazon EKS が Kubernetes 1.22 のサポートを開始](https://aws.amazon.com/jp/blogs/news/amazon-eks-now-supports-kubernetes-1-22/)

[Amazon EKS アドオンで Amazon EBS CSI ドライバーが一般利用可能になりました](https://aws.amazon.com/jp/blogs/news/amazon-ebs-csi-driver-is-now-generally-available-in-amazon-eks-add-ons/)

[【レポート】Amazon EKS と Datadog によるマイクロサービスの可観測性（パートナーセッション） #PAR-03 #AWSSummit](https://dev.classmethod.jp/articles/aws_summit_japan_2021_datadog/)


[Amazon EKS で GitOps パイプラインを構築する](https://aws.amazon.com/jp/blogs/news/building-a-gitops-pipeline-with-amazon-eks/)



