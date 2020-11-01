
# Auto Scaling

#### Auto Scaling の整理

* EC2 Auto Scaling: EC2 インスタンス。
* Application Auto Scaling: ECSクラスタ、EMRクラスタ、Auroraレプリカなど。

#### バランシング

[アベイラビリティーゾーンの追加](https://docs.aws.amazon.com/ja_jp/autoscaling/ec2/userguide/as-add-availability-zone.html)

**1 つの AZ に問題が発生すると、影響を受けていない AZ で新しいインスタンスを起動する。**AZ が正常な状態に戻ると AZ に渡ってインスタンスを移動的に再分散する。

update-auto-scaling-group コマンドを使用して、Auto Scaling グループにサブネットを追加できる。
set-subnets コマンドを使用して、Application Load Balancer で新しいサブネットを有効にする。

[自動スケーリングのメリット](https://docs.aws.amazon.com/ja_jp/autoscaling/ec2/userguide/auto-scaling-benefits.html)

Auto Scaling グループは可用性ゾーン間で不均衡になる可能性がある。
再バランシングして補正されるようになっており、その際は新しいインスタンスを起動してから古いインスタンスを停止する動作となる。


#### ライフサイクル

[ライフサイクル](https://docs.aws.amazon.com/ja_jp/autoscaling/ec2/userguide/AutoScalingGroupLifecycle.html)

起動するまでの間は Pending。ライフサイクルフックを設定できる。

正常に起動したインスタンスは InService の状態。 

ヘルスチェックの失敗やスケールインによって、Terminating → Terminated へと遷移する。ライフサイクルフックを設定できる。

デタッチすることもできる。

スタンバイ状態にして一時的に外し、再び組み込むこともできる。



## Auto Scaling Group

[複数のインスタンスタイプと購入オプションを使用する Auto Scaling グループ](https://docs.aws.amazon.com/ja_jp/autoscaling/ec2/userguide/asg-purchase-options.html)

オンデマンドインスタンスとスポットインスタンスの組み合わせからなるフリートを使用可能。
配分戦略で、オンデマンドとスポットの容量をどのように満たすかを設定可能。

[起動テンプレートを使用した Auto Scaling グループの作成](https://docs.aws.amazon.com/ja_jp/autoscaling/ec2/userguide/create-asg-launch-template.html)

[起動設定を使用した Auto Scaling グループの作成](https://docs.aws.amazon.com/ja_jp/autoscaling/ec2/userguide/create-asg.html)

起動テンプレートもしくは起動設定から Auto Scaling Group を作成可能。ドキュメントでは、EC2 の最新機能を使用できるように**起動テンプレートが推奨されている。**

[Elastic Load Balancing および Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/ja_jp/autoscaling/ec2/userguide/autoscaling-load-balancer.html)

ヘルスチェックにおいては、全てのロードバランサからのチェックに合格する必要がある。**一つでも異常と判定されると置き換えられる動作となる。**

1 つのアベイラビリティーゾーンが異常ありまたは使用不可になると、Amazon EC2 Auto Scaling は、影響を受けていないアベイラビリティーゾーンで新しいインスタンスを起動する。異常のあるアベイラビリティーゾーンが正常な状態に戻ると、Amazon EC2 Auto Scaling は Auto Scaling グループのアベイラビリティーゾーンにわたって均等にインスタンスを自動的に再分散する。

[最大インスタンス有効期限](https://docs.aws.amazon.com/ja_jp/autoscaling/ec2/userguide/asg-max-instance-lifetime.html)

最大インスタンス有効期間の機能を設定しておくことで、稼働時間が最大許容時間に達したインスタンスが置き換えられる。

[インスタンスの更新](https://docs.aws.amazon.com/ja_jp/autoscaling/ec2/userguide/asg-instance-refresh.html)

インスタンスの更新を行うことにより、ASG のインスタンスを置き換えることができる。一気に更新するのではなく、ローリングアップデートされる動作となる。



## スケーリング

[スケーリング](https://docs.aws.amazon.com/ja_jp/autoscaling/ec2/userguide/scaling_plan.html)

スケーリング方法

* 希望台数の維持
* 希望台数の変更によるスケール
* スケジュールに基づくスケーリング
* 需要に基づくスケーリング（リソース使用率を指定値に維持）
* 予測スケーリング

#### 動的なスケーリング

[動的スケーリング](https://docs.aws.amazon.com/ja_jp/autoscaling/ec2/userguide/as-scale-based-on-demand.html)

CloudWatch アラームを使用してスケーリングする仕組みとなっている。

[ターゲット追跡スケーリング](https://docs.aws.amazon.com/ja_jp/autoscaling/ec2/userguide/as-scaling-target-tracking.html)

例としては CPU 使用率を 40 % に維持するような指定方法。ステップスケーリングで細かく指定するのが面倒な場合は、こちらを推奨。

事前定義されたリクエストを使用可能。

* ASGAverageCPUUtilization
* ASGAverageNetworkIn
* ASGAverageNetworkOut
* ALBRequestCountPerTarget


[ステップスケーリング](https://docs.aws.amazon.com/ja_jp/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html)

一つのメトリクスに対して、複数のスケーリング条件を設定可能。例えば CPU 使用率 70 % で +1 台、80 % で + 2 台といった具合。条件を一つにすれば簡易スケーリングと同じことができる。
  * ウォームアップ周期: 新しいインスタンスがサービス開始できるまで何秒要するかを設定。これにより、ウォームアップ期間中に次のトリガーが発動したときでも、条件合致時の全台を起動するのではなく現在起動中台数の差分台数だけ起動する。

簡易スケーリング: 現在は非推奨。あるメトリクスが X % 以上になったら、といった条件でスケーリング。

[スケジュールされたスケーリング](https://docs.aws.amazon.com/ja_jp/autoscaling/ec2/userguide/schedule_time.html)

一度限り、あるいは cron のような定期的なスケーリングを設定可能。

予測スケーリング

過去のメトリクスをもとに将来の需要を予測し、キャパシティの増減をスケジュールする。

[終了ポリシー、スケールインからの保護](https://docs.aws.amazon.com/ja_jp/autoscaling/ec2/userguide/as-instance-termination.html)

デフォルトの終了ポリシーではインスタンスが各 AZ に均等に配置されるようになっている。終了ポリシーは複数あり、変更することも可能。
**スケールインから保護する設定**もある。インスタンスごとに個別に設定することも可能。

[ライフサイクルフック](https://docs.aws.amazon.com/ja_jp/autoscaling/ec2/userguide/lifecycle-hooks.html)

**インスタンスの起動時、削除時にカスタムアクションを実行可能。**タイムアウト期間（デフォルトは 1 時間）が経過するまで待機状態となる。（Pending:Wait、Terminating:Wait）。

Amazon EventBridge、Amazon SNS、Amazon SQS を使用して通知を設定可能。

なお、ライフサイクルフックを追加した場合、ヘルスチェックの猶予期間が始まるのは、ライフサイクルフックアクションが完了してインスタンスが InService 状態になってから。

[インスタンスの一時的な削除](https://docs.aws.amazon.com/ja_jp/autoscaling/ec2/userguide/as-enter-exit-standby.html)

スタンバイ状態にすることができる。
スタンバイ状態にすると対象インスタンスはロードバランサから登録解除され、希望する容量の値も減少する。

[スケーリングの中断](https://docs.aws.amazon.com/ja_jp/autoscaling/ec2/userguide/as-suspend-resume-processes.html)

管理上の中断が発生する場合がある。インスタンスの起動を 24 時間以上試みている場合が該当。

**スケーリングプロセスを Suspend することができる。もとに戻すときは Resume。**



## Monitoring

[ヘルスチェック](https://docs.aws.amazon.com/ja_jp/autoscaling/ec2/userguide/healthcheck.html)

EC2 インスタンスのチェックは、ステータスチェックにより行う。stopped, stopping, terminated, terminating の状態もチェックしている。

**ELB はデフォルトでは有効になっていない。**

ヘルスチェックの猶予期間を設定することで、インスタンスの起動後、指定した時間の間ヘルスチェックを実施しない。

[グループメトリクス](https://docs.aws.amazon.com/ja_jp/autoscaling/ec2/userguide/as-instance-monitoring.html)

グループメトリクスを有効化することで、ASG に関する CloudWatch メトリクスが採取できるようになる。



## 設定項目

#### 起動テンプレートの作成

起動設定だと作成後の修正ができない。AMI などを更新した上でバージョン管理できる起動テンプレートを推奨。

[Auto Scaling グループの起動テンプレートを作成する](https://docs.aws.amazon.com/ja_jp/autoscaling/ec2/userguide/create-launch-template.html)

```
// ドキュメントからのコピペ
$ aws ec2 create-launch-template \
    --launch-template-name <テンプレート名>
    --version-description <バージョン> \
    ---launch-template-data '{"NetworkInterfaces":[{"DeviceIndex":0,"AssociatePublicIpAddress":true,"Groups":["sg-7c227019"],"DeleteOnTermination":true}],"ImageId":"ami-01e24be29428c15b2","InstanceType":"t2.micro","TagSpecifications": [{"ResourceType":"instance","Tags":[{"Key":"purpose","Value":"webserver"}]}]}'
```

含まれる情報

* AMI の ID
* インスタンスタイプ
* キーペア
* セキュリティグループ
* その他 EC2 インスタンスを起動するために使用するパラメータ

#### Auto Scaling グループの作成

[起動テンプレートを使用して Auto Scaling グループを作成する](https://docs.aws.amazon.com/ja_jp/autoscaling/ec2/userguide/create-asg-launch-template.html)


[create-auto-scaling-group](https://docs.aws.amazon.com/cli/latest/reference/autoscaling/create-auto-scaling-group.html)

```
// 主要オプションを抜粋
$ aws ec2 create-auto-scaling-group \
    --auto-scaling-group-name <value> \
    [--launch-template <value>] \
    [--instance-id <value>] \
    --min-size <value> \
    --max-size <value> \
    [--desired-capacity <value>] \
    [--default-cooldown <value>] \
    [--availability-zones <value>] \
    [--load-balancer-names <value>] \
    [--target-group-arns <value>] \
    [--health-check-type <value>] \
    [--health-check-grace-period <value>] \
    [--vpc-zone-identifier <value>] \
    [--tags <value>] \
```


# 参考

* [[AWS Black Belt Online Seminar] Amazon EC2 Auto Scaling and AWS Auto Scaling 資料及び QA 公開](https://aws.amazon.com/jp/blogs/news/webinar-bb-amazon-ec2-auto-scaling-and-aws-auto-scaling-2019/)
* [Amazon EC2 Auto Scaling ドキュメント](https://docs.aws.amazon.com/ja_jp/autoscaling/ec2/userguide/what-is-amazon-ec2-auto-scaling.html)



