
# Auto Scaling

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
* ステップスケーリング: 一つのメトリクスに対して、複数のスケーリング調整値を指定可能。
  * ウォームアップ周期: 新しいインスタンスがサービス開始できるまで何秒要するかを設定。これにより、ウォームアップ期間中に次のトリガーが発動したときでも、条件合致時の全台を起動するのではなく現在起動中の台数を加味して起動する。
* ターゲット追跡受けーリング: CPU 使用率を 40 % に維持するような指定方法。
* 予測スケーリング: 過去のメトリクスをもとに将来の需要を予測し、キャパシティの増減をスケジュールする。
* スケジュールスケーリング: 一度限り、あるいは定期的なスケーリングを設定可能。

#### その他機能

* ミックスインスタンスグループ: スポットインスタンスを混ぜることが可能。
* インスタンスの保護: 保護対象したインスタンスをスケールイン対象外とできる。
* スタンバイ: 一時的に Auto Scaling グループから外す。ELB の対象からも外れる。
* デタッチ: インスタンスを起動したまま、Auto Scaling グループから外す。
* ライフサイクルフック: インスタンスの起動、削除時にカスタムアクションを実行できる。


# 参考

* [[AWS Black Belt Online Seminar] Amazon EC2 Auto Scaling and AWS Auto Scaling 資料及び QA 公開](https://aws.amazon.com/jp/blogs/news/webinar-bb-amazon-ec2-auto-scaling-and-aws-auto-scaling-2019/)
* [Amazon EC2 Auto Scaling ドキュメント](https://docs.aws.amazon.com/ja_jp/autoscaling/ec2/userguide/what-is-amazon-ec2-auto-scaling.html)



