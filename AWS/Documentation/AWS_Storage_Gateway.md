# BlackBelt

[AWS Storage Gateway](https://pages.awscloud.com/rs/112-TZM-766/images/AWS-Black-Belt_2023_AWS-Storage-Gateway_0131_v1.pdf)

* オンプレミス環境から標準的なプロトコル(MB, NFS, iSCSI) を用いて AWS のストレージを利用できる
* Storage Gateway はオンプレミスの仮想マシンや EC2 に構成できる
* 4 種類ある
  * S3 Gateway: SMB, NFS でマウント
  * Volume Gateway: iSCSI でマウント
  * Tape Gateway: iSCSI VTL で接続
  * FSx File Gateway: SMB でマウント
* 導入方法
  * Storage Gateway の VM イメージを使用
  * Storaga Gateway が導入済みのハードウェアアプライアンスをラックマウントして使用
  * Storage Gateway の AMI を使用して EC2 を起動
* S3 File Gateway
  * クライアントからは SMB, NFS で利用可能
  * S3 File Gateway にキャッシュされるのでレイテンシーを低減できる



# 参考

* Document
  * [AWS Storage Gateway のドキュメント](https://docs.aws.amazon.com/ja_jp/storagegateway/index.html)
* Black Belt
  * [AWS Storage Gateway](https://pages.awscloud.com/rs/112-TZM-766/images/AWS-Black-Belt_2023_AWS-Storage-Gateway_0131_v1.pdf)


