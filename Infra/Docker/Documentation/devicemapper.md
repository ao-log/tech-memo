
[Use the Device Mapper storage driver (deprecated)](https://docs.docker.com/storage/storagedriver/device-mapper-driver/)
[Device Mapper ストレージ・ドライバの使用](https://docs.docker.jp/v17.06/engine/userguide/storagedriver/device-mapper-driver.html)

* Docker v25.0 にて deprecated
* Device Mapper はカーネルベースのフレームワーク。devicemapper はストレージドライバとして、シンプロビジョニング、スナップショット機能をイメージとコンテナの管理のために使用する
* ファイルレベルではなくブロックレベルで動作
* /etc/docker/daemon.json でストレージドライバを指定できる
* `docker info` によりストレージドライバを確認できる
```
$ docker info
...
  Storage Driver: devicemapper
```


### loop-lvm モード

テスト用。ローカルディスク上のファイルを物理ディスクもしくはブロックデバイスのように読み書きできる loopback メカニズムを使用。loopback メカニズムと OS ファイルシステムレイヤーとのやり取りが発生するので、IO 性能が悪い


### direct-lvm モード

* 本番環境では devicemapper を direct-lvm モードで使用するべき。thin pool を作る仕組みで、loopback デバイスよりも高速
* 構成できるブロックデバイスは 1 つ。ただし手動設定を行うことで複数個設定できる
* 以下の例のように `daemon.json` ファイルを設定する
```json
{
  "storage-driver": "devicemapper",
  "storage-opts": [
    "dm.directlvm_device=/dev/xdf",
    "dm.thinp_percent=95",
    "dm.thinp_metapercent=1",
    "dm.thinp_autoextend_threshold=80",
    "dm.thinp_autoextend_percent=20",
    "dm.directlvm_device_force=false"
  ]
}
```
* 構築方法
  * LVM の論理グループ `thinpool`、`thinpoolmeta` を作成する
  * `lvconvert` コマンドによりボリュームを thin pool、メタデータの保存場所に変換
  * 自動拡張の設定を `/etc/lvm/profile/docker-thinpool.profile` にて行い、`lvchange` コマンドで適用
  * `/var/lib/docker/` を削除しておく。そうすると Docker はコンテナイメージ、コンテナの管理用に新しい LVM pool を使用する
  * `docker info` で以下のように表示されれば OK
  ```
   Pool Name: docker-thinpool ← pool 名
   Pool Blocksize: 524.3 kB
   Base Device Size: 10.74 GB
   Backing Filesystem: xfs
   Data file: ← 空
   Metadata file: ← 空
  ```


### devicemapper の動作

* `/var/lib/docker/devicemapper/metadata/`: メタデータ用
* `/var/lib/docker/devicemapper/mnt/`: イメージとコンテナレイヤーが格納される
* 読み取りはブロックレベルで行われる。コンテナ自体はブロックを持っていないが、親イメージへのポインタを持っておりそこからブロックを読み取る。その後はコンテナ上のメモリに乗る
* 書き込みはコンテナの書き込み可能レイヤーにブロックが割り当てられる。ファイルの上書き時は変更されたブロックのみを書き込む


### パフォーマンス

* 大量のブロックの書き込み時のパフォーマンスは悪い。書き込みが多いワークロードではストレージドライバをバイパスするデータボリュームを使用する必要がある
* ファイルのコピーをメモリ上に載せるため、メモリを多く消費する


