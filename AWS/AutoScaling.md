
# Auto Scaling

## どのようにスケーリングするか

#### 希望台数

* 台数維持(台数減少を検知して機能台数まで回復させる)
* 手動スケーリング(希望台数を手動で変更することでスケーリング)
* 自動スケーリング(ポリシーに応じて自動スケーリング)

#### Auto Scaling の整理

* EC2 Auto Scaling: EC2 インスタンス。
* Application Auto Scaling: ECSクラスタ、EMRクラスタ、Auroraレプリカなど。
* AWS Auto Scaling: スケーリングプラン

#### スケーリング方法

* 簡易スケーリング: 現在は非推奨。あるメトリクスが X % 以上になったら、といった条件でスケーリング。
  * クールダウンピリオド    
* [ステップスケーリング](https://docs.aws.amazon.com/ja_jp/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html): 一つのメトリクスに対して、複数のスケーリング条件を設定可能。例えば CPU 使用率 70 % で +1 台、80 % で + 2 台といった具合。条件を一つにすれば簡易スケーリングと同じことができる。
  * ウォームアップ周期: 新しいインスタンスがサービス開始できるまで何秒要するかを設定。これにより、ウォームアップ期間中に次のトリガーが発動したときでも、条件合致時の全台を起動するのではなく現在起動中台数の差分台数だけ起動する。
* [ターゲット追跡受けーリング](https://docs.aws.amazon.com/ja_jp/autoscaling/application/userguide/application-auto-scaling-target-tracking.html): CPU 使用率を 40 % に維持するような指定方法。ステップスケーリングで細かく指定するのが面倒な場合は、こちらを推奨。
* [スケジュールスケーリング](https://docs.aws.amazon.com/ja_jp/autoscaling/application/userguide/application-auto-scaling-scheduled-scaling.html): 一度限り、あるいは定期的なスケーリングを設定可能。
* 予測スケーリング: 過去のメトリクスをもとに将来の需要を予測し、キャパシティの増減をスケジュールする。

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

#### ヘルスチェック

デフォルトは EC2 のみ。ELB を選択すると ELB と EC2 の両方の条件でヘルスチェックする。

[Auto Scaling グループへの Elastic Load Balancing ヘルスチェックの追加](https://docs.aws.amazon.com/ja_jp/autoscaling/ec2/userguide/as-add-elb-healthcheck.html)

#### [猶予期間](https://docs.aws.amazon.com/ja_jp/autoscaling/ec2/userguide/healthcheck.html#health-check-grace-period)

ヘルスチェックを猶予する期間。この期間の間は unhealthy の判定を行わない。そのため、unhealthy と判定されて無駄にインスタンスを起動することがなくなる。

## その他機能

* ミックスインスタンスグループ: スポットインスタンスを混ぜることが可能。
* インスタンスの保護: 保護対象したインスタンスをスケールイン対象外とできる。
* スタンバイ: 一時的に Auto Scaling グループから外す。ELB の対象からも外れる。
* デタッチ: インスタンスを起動したまま、Auto Scaling グループから外す。
* ライフサイクルフック: インスタンスの起動、削除時にカスタムアクションを実行できる。


# 参考

* [[AWS Black Belt Online Seminar] Amazon EC2 Auto Scaling and AWS Auto Scaling 資料及び QA 公開](https://aws.amazon.com/jp/blogs/news/webinar-bb-amazon-ec2-auto-scaling-and-aws-auto-scaling-2019/)
* [Amazon EC2 Auto Scaling ドキュメント](https://docs.aws.amazon.com/ja_jp/autoscaling/ec2/userguide/what-is-amazon-ec2-auto-scaling.html)



