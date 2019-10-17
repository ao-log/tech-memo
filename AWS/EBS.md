
## 特徴

* AZ ごとに独立している
* ボリュームのデータは複数の HW にレプリケートされている。
* EBS 最適化を有効化すると、他のトラフィックから分離し、EBS 専用の帯域を確保できる。
* アタッチ中でも、サイズ、IOPS を変更可能。拡張後は、OS からファイルシステム拡張が必要。IOPS 変更は徐々に反映される。変更後は 6 時間は変更不可。
* 暗号化可能。

## Snapshot

Snapshot は S3 上に保存される。採取するときはアンマウントするなど、書き込みがない状態を作ることが望ましい。
増分で作成される。古いものを削除した場合は、隣のバックアップにマージされる。

CloudWatchEvents と連携することでリージョン完コピーを自動化。
AWS ata Lifecycle Manager により定期的に Snapshot を取得可能。AWS Backup では EBS だけでなく他のサービスも含めてバックアップを管理可能。


## 種類

* SSD: 汎用 SSD、プロビジョンド IOPS
* HDD: スループット最適化 HDD(シーケンシャルアクセス、例えば Hadoop に向いている)、コールド HDD(ログやバックアップのアーカイブ先に向いている)

gp2 はデフォルトのボリュームタイプ。IOPS をバーストすることができる。バーストバケットモデルを採用し、クレジットが残っているときにバーストする。

## 料金

容量とリクエストに課金される。そのため、容量は小さいけれど大量のファイルがある場合は注意。

# 参考

* [AWS Black Belt Online Seminar Amazon EBS 資料及び QA 公開](https://aws.amazon.com/jp/blogs/news/webinar-bb-amazon-ebs-2019/)
* [EBS ドキュメント](https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/AmazonEBS.html)
