今回 6 回目の開催とのこと。

# Preferred Networksの機械学習クラスタを支える技術

https://www.slideshare.net/pfi/tech-inmlcluster20180729-julytechfesta2018

##### 事例

* ピッキングの自動化
* 自動運転（強化学習）
* 自然言語によるロボット制御

##### Chainer

**Define-by-Run**
実行しながらニューラルネットワークを構築する。

**ChainerMN**
NCCL, CUDA-Aware MPI などで高性能を実現。
重み情報を交換するために、ノードをまたぐプロセス間通信が多く発生する。
Infiniband によって、高速、低レイテンシを実現。

Tesla P100 が 1024 個ある。8 GPU x 16 Servers。
バイセクション bandwidth は 56 Gbps x 2。

Top500 のインダストリ領域で、国内一位。

##### なぜオンプレ？

大量の計算機で誰にも成し遂げられないことをしたい。クラウドは無限ではない。
日々、16 GPU, 32 GPU の学習を日常的に回したい。
Infiniband はクラウドで対応していない。

##### クラスタサービス概要

Slack 経由で Azure にリモートデスクトップ環境
コンテナは Docker、Mesos と Kubernetes の二つ。
JupyterHub もコンテナ上で動いている。

内製のジョブスケジューラ。
必要なリソースと環境を指定すると、Docker イメージがビルドされる。
データセットを指定すると、自動的にマウントされる。

##### 求められる要件

###### シームレスな実行

yaml で宣言的に指定。
Dockerfile を書かずに、どのライブラリが必要かを指定。
JupyterHub 上での作業も出来る。

モデルの作成が多い。良さそうなアルゴリズムができたら大規模実行。
トライアンドエラーを繰り返すので、トレーサビリティが重要。
後で追えるように。

###### 概念モデル

**Project**
ソースを配置。

**Job**
yaml で定義可能
Docker イメージ自動ビルド。
マルチノード実行可能。
障害時の自動リトライ

**Dataset**
バージョニング可能。Immutable。ACL 設定可能

**Result**
結果の返却。

###### 再現性の問題

ニューラルネット：初期重み。ドロップアウトなどのランダムノイズ。→乱数シードの固定→メタデータとしてユーザが保存
データセット：validation 用のsplit/shuffle。分散ワーカ用のsplit/shuffle。→乱数シードの固定→メタデータとしてユーザが保存
コード：バージョン。→Dockerイメージ→ツールで提供
CPU：マルチスレッド→制御不能
GPU：cuDNNの幾つかの演算は再現性を保証していない

###### GPU クラスタ

mesos containerizer → cgroup でアクセスできるデバイスを制限
kubernetes: nvidia-docker, nvidia/k8s-device-plugin → kubelet が device オプション付きで起動

Infiniband
kubernetes は ダミーの device をたくさん作ってコンテナに見せている

###### クラスタ利用率の向上

MostResourceRequest → リソースを詰め詰めにするスケジューリング。

ノードとジョブにラベルをつけることで配置制約を指定できるようにしている。

公平性。誰にも公平にリソースを割り当てられるように。
→あまりできていない。GPU 数に上限を設けたりとか。

**ギャングスケジューリング。**
複数のコンテナ(pod)を同時にノードにスケジューリングする。

**プリエンプションとリエントランシー（再開可能性）。**
ジョブに優先度を設定し、低優先度のものを止める。止まったものは後でリランする。
現状では、エポック終了後に保存されるモデルスナップショットを読み込んでリスタート。

**グランドチャレンジ。**
全系を活用した大規模計算。

###### 今後の展望

遊休資源への対応。→ Kubernetes への一本化。クラスタ利用率の向上。
Infiniband の isolation。
スケジューリングも機械学習で効率化。
ハードウェアトポロジを考慮した配置（ブロックダイアグラムを考慮）。
リエントランシーの追求。
Fairness on GPU。Dominant Resource Fair。
ハイブリッドクラウドの活用。
日常的なグランドチャレンジ。

効率的で柔軟な機械学習クラスタは今や世界中の会社にとっての課題。
OSS 活用で楽になることが多い。
OSS への貢献は大事。

Convergence of HPC and Cloud
ボトルネックは計算資源であってはならない。
最強の計算資源でクリエイティブティを最大化した。

##### Chainer-operator

Kubernetes 上でマルチノードができる。


# ハイパフォーマンスコンピューティング向けコンテナ技術 Singularity と NVIDIA GPU CLoud で作るハイブリッド機械学習環境の構築

https://speakerdeck.com/pottava/building-a-hybrid-environment-for-machine-learning-with-singularity-and-ngc-1

###### 要件

* 高性能GPU
* 数々のMLフレームワーク
* 社内認証との統合
* ファイル、権限

GPU だけではない。
CPU の AVX 命令や MPI のマルチノードの分散環境を採用するフレームワークも。

Horovod も mpirun で実行するらしい。

コンテナで対応する。
コンテナイメージを変えることで、各ソフトに対応できる。

dockerd → containerd → コンテナ用の隔離した環境を作成

###### Docker の悩み

* root レスでは動かない。
* 実行ユーザの扱いが難しい。共有ストレージへ適切に読み書きする難しさ。
* priviliged が必要なケースもある。
* マルチノードが大変。

###### Singularity

* Docker のいいところ + HPC のサポート
* 高性能ハードウェアや既存のジョブスケジューラがそのまま使える
* singularity run したユーザのプロセスとして動作する。

* Docker イメージからコマンド一つで Singularity 用に変換できる。
* 「mpirun -np 4 singularity run イメージ」のように使える
* Infiniband も問題なく使える

佐藤先生のスライドが詳しいらしい。

###### 構築、運用

* Dockerfile を CI でバージョン管理。
* プライベートな Docker レジストリである HARBOR など。
* サーバそのものは従来の構成管理

###### 実行環境

GPU を使う場合は、CUDA をベースイメージに。
CI サーバを使って、テスト、ビルド、プッシュ。
フレームワークごとに Dockerfile を準備。

**NGC**

NVIDIA 社の Docker レジストリ。
イメージは毎月更新される。CUDA、ライブラリはもちろん同梱。無料。

###### Kubernetes でハイブリッド環境

NVIDIA もサポートを宣言。

ジョブを定義した yaml を Apply
→　条件が合うノードにデプロイ

Federated Cluster
→　ハイブリッドなクラスタが組める

###### Kubernetes の課題

* サーバ間通信が DNS なので、MPI だと足を引っ張られる。
* Docker だとリソースをクラスタ最大にできない

###### ハイブリッド環境構築の現実

* Rancher 経由で認証認可や Docker レジストリと連携。ラッパーに任せる。
* ジョブ投入時にマニフェストを動的に生成（LDAP に合うようにしたりとか）

###### 機械学習環境の CI/CD

* サーバを k8s クラスタにジョイン
* イメージの管理
* ジョブ投入部分の作り込み

###### IaaS にありがちな重いインテグレーション

* アカウント連携、ファイル共有
* クラスタのライフサイクル管理（チューニング、廃棄）
* 財務管理（利用量の測定、予算管理）

###### ReScale

* ファイル転送は暗号化
* 隔離された環境の構築
* 予算管理可能（CIDR の制御、上限設定）

ReScale の Web 画面から Docker イメージをダウンロードできる。
端末上でイメージが動いて、Jupyter Notebiik が起動する。
Web 上でリソース量を選んでジョブ投入できる。

# ゼロから分かる Kubernetes Native なアプリケーションの作り方：Operator Framework の概要とデモを交えた適用例

イベント監視

```
$ kubectl get events -w
```

# 英語

動詞に注目。
中学レベルの英語文法はおさえると良い。
ちょこちょこ続けるのがコツ。
五感を使う。
過去形は覚える
見ないで言えるようになる

# イゼルローン要塞攻略

https://www.slideshare.net/irix_jp/the-strategy-from-the-iserlohn-fortress-at-jtf2018

###### 背景知識

銀河帝国と自由惑星同盟は惰性的な戦争状態。
経済面と水面下の謀略によって、漁夫の利を狙うフェザーン。

まず、世界観を決める。一貫した行動指針を作る。
目的のための手段を考える。

相手と会話するときも、どの階層かを考えながら話す。
「芸の絶対化」とならないように注意。

自分で前提条件を整えるというアプローチ。
* 状況作り
* 環境作り

自分自身の目指すべき状態をコントロールする思考が「戦略」
自分の目指す場所、そのためにやるべきことを整理する。
ルールは従うものではなく、変えるもの。

# 招待講演(メルカリさん)

https://speakerdeck.com/kazeburo/mercari-infrastructure-and-software

###### 最新の構成

DNS: Route53
CDN: Akamai, Fastly, ImageFlux
Infra: Sakura + GCP, US: AWS + GCP, UK: GCP

###### 歴史

2013年開始。さくらの VPS で 1 台に Web + DB 構成で開始。
2014年 US リリース。AWS 経験者が増加。インフラ専任者が少なく、マネージドサービスを利用して構築。
2015年 SREチーム発足。サービスの信頼性、スケーラビリティ向上に取り組む。
* memcached の有効活用。N+1対策、SQLチューニング、Master/Slave使い分け。
* HTTP2化、PGP upgrafe/最適化

マネージドサービスをサーバにリプレイス
ELB → DNS-RR, NGINX
Internal ELB → Internal DNS(Bind)、Consul
ElastiCache → memcached on EC2
RDS → MySQL on EC2

JP, US, UK で共通の構成とした。
JP の実績のある構成。少人数での運用。

###### Infrastructure を支える Software

**Consul**

* Service Discovery
  - DNS インタフェースを利用することで内部LBとして利用
  - 内部向けサービスのエンドポイント

**OpenResty**

* nginx をベースに作られた Web Platform
* nginx のスケーラビリティをそのままに様々な拡張が可能
* グローバル側 nginx と OpenResty の間は KeepAlive が有効。TCP コネクションの有効活用。
* 一つの商品に大量アクセスが来る対策として、並列度の制御。リクエストが来た時 memcached に add。売れるまでをこれで耐える。

###### Microservices

* Container/Docker
* Kubernetes
* Spinnaker
* Terraform

###### 課題

handshake を避ける
コネクションを持続するオープンソースを書いた(chocon/go製)

###### まとめ

ソフトウェアによって可用性、パフォーマンスを向上、インフラを定義。
多くのソフトが SRE から生まれている。

###### SRE としての矜持

サーバの能力を引き出し、アプリエンジニアが作成したコードを最高の形で動かすこと。
いかにして、お客様に「いつでも快適に安全に使える」信頼性を提供するか。
