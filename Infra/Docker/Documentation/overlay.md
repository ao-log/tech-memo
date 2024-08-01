
[Use the OverlayFS storage driver](https://docs.docker.com/storage/storagedriver/overlayfs-driver/)

* カーネルドライバーは `OverlayFS`。Docker ストレージドライバーは `overlay2`
* xfs ファイルシステムでサポートされる。`d_type=true` を設定する必要がある
* /etc/docker/daemon.json でストレージドライバを指定できる
* `docker info` によりストレージドライバを確認できる
```
$ docker info
...
Storage Driver: overlay2
 Backing Filesystem: xfs
 Supports d_type: true
```
* lowerdir, upperdir の階層構造となっており、merged view として見せている。上位レイヤーの内容で上書きされる仕組み
* コンテナにてファイルを上書きする場合は `copy_up` オペレーションによりファイルレベルでコピーされる。そのため、大きなファイルだとパフォーマンスに影響が出る。`copy_up` は当該ファイルを初めて上書きするときのみ実行される
* ファイル削除時には whiteout ファイルが作成される仕組み

