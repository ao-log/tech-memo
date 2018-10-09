

# Virtual Private Cloud（VPC）ネットワークの概要

https://cloud.google.com/vpc/docs/vpc?hl=ja

制限

* 1 つのプロジェクトに作成できるネットワークは 5 個まで

インスタンス作成時

* インスタンスを作成する場合は、ゾーン、ネットワーク、サブネットを選択します。選択可能なサブネットは、選択したリージョンのサブネットに限定されています。

VPC ネットワークには、自動モードとカスタムモードの 2 種類がある。

* 自動モード: ネットワークが作成されると、各リージョンから 1 つのサブネットがネットワーク内に自動的に作成されます。

###### ルート

* ルートは、インスタンスから出るパケット（下りトラフィック）のパスを定義。
インターネットに出ていくための default-internet-gateway など 4 つのタイプがある。

###### ファイアウォール

* すべての VPC ネットワークには、2 つの暗黙のファイアウォール ルールがあります。
  * 1 つの暗黙ルールはすべての下りトラフィックを許可
  * もう 1 つはすべての上りトラフィックを拒否

送信元

* デフォルトの送信元は、すべての受信トラフィック（0.0.0.0/0）です。送信元を絞り込むには、送信元フィルタを使用します。
  * IP アドレスの範囲（GCP の内部または外部）
  * サービス アカウントまたはネットワーク タグによって識別されるインスタンス


webserver でタグ付けされたインスタンスへのすべての上り TCP トラフィックを拒否、下りは 80/tcp のみ許可するルールを作成。

https://cloud.google.com/vpc/docs/using-firewalls?hl=ja

```
  gcloud compute firewall-rules create deny-subnet1-webserver-access \
      --network my-network \
      --action deny \
      --direction ingress \
      --rules tcp \
      --source-ranges 0.0.0.0/0 \
      --priority 1000 \
      --target-tags webserver

  gcloud compute firewall-rules create vm1-allow-ingress-tcp-port80-from-subnet1 \
      --network my-network \
      --action allow \
      --direction ingress \
      --rules tcp:80 \
      --priority 50 \
      --target-tags webserver
```

###### 共有 VPC

https://cloud.google.com/vpc/docs/shared-vpc?hl=ja

複数のプロジェクトから共通の VPC ネットワークにリソースを接続できる

### 参考

* [Google Cloud DNS のドキュメント](https://cloud.google.com/dns/docs/?hl=ja)
* [Google Cloud Load Balancing Documentation](https://cloud.google.com/load-balancing/docs/?hl=ja)
* [HTTP(S) Load Balancing Concepts](https://cloud.google.com/load-balancing/docs/https/?hl=ja)
* [Virtual Private Cloud のドキュメント](https://cloud.google.com/vpc/docs/?hl=ja)

* [[Cloud OnAir] Google Networking Deep Dive ! その技術と設計の紹介 2018年8月9日 放送](https://www.slideshare.net/GoogleCloudPlatformJP/cloud-onair-google-networking-deep-dive-201889)
* [Google Cloud のネットワークとロードバランサ](https://www.slideshare.net/GoogleCloudPlatformJP/google-cloud-72158836)
