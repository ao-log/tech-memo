
## 記事

[AWS ParallelCluster でもSSH接続はやめて、セッションマネージャーで接続できます](https://dev.classmethod.jp/articles/connect-to-aws-parallelcluster-with-session-manager/)



[AWS ParallelCluster コンピュートノードが終了するまでの時間設定](https://dev.classmethod.jp/articles/aws-parallelcluster-autoscaling/)

* scaledown_idletime によりインスタンス終了までの時間を設定


[AWS ParallelCluster 複数のS3バケットへのアクセス設定](https://dev.classmethod.jp/articles/aws-parallelcluster-s3bucket-access-settings/)

* s3_read_resourceと、s3_read_write_resource で指定するS3バケットは各々複数指定できない
* そのため、IAM ポリシー側での制御が必要

