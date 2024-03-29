
## 記事

[AWS ParallelCluster 3.7.0 で Ubuntu 22.04 や専用のログインノードの追加など新たにサポートされました](https://dev.classmethod.jp/articles/aws-parallelcluster-v370-released/)

* Ubuntu 22.04 LTS を新規サポート
* 新機能ログインノードをサポート
* EBS のルートボリュームの必要最小サイズが 35GB から 40 GB へ引き上げ
* IMDSv2 がデフォルト。インスタンスメタデータを参照するようなカスタムブートストラップアクション、ジョブ利用時に注意


[AWS ParallelCluster のメジャーアップデート v3.0.0 がリリースされました](https://dev.classmethod.jp/articles/aws-parallel-cluster-v300-released/)



[AWS ParallelCluster でもSSH接続はやめて、セッションマネージャーで接続できます](https://dev.classmethod.jp/articles/connect-to-aws-parallelcluster-with-session-manager/)



[AWS ParallelCluster コンピュートノードが終了するまでの時間設定](https://dev.classmethod.jp/articles/aws-parallelcluster-autoscaling/)

* scaledown_idletime によりインスタンス終了までの時間を設定


[AWS ParallelCluster 複数のS3バケットへのアクセス設定](https://dev.classmethod.jp/articles/aws-parallelcluster-s3bucket-access-settings/)

* s3_read_resourceと、s3_read_write_resource で指定するS3バケットは各々複数指定できない
* そのため、IAM ポリシー側での制御が必要

