

### インスタンスの作成

##### イメージとマシンタイプの確認

利用できるイメージをマシンタイプを確認します。

```shell-session

$ gcloud compute images list
NAME                                                  PROJECT            FAMILY                            DEPRECATED  STATUS
centos-6-v20180104                                    centos-cloud       centos-6                                      READY
centos-7-v20180104                                    centos-cloud       centos-7                                      READY
...
```

```shell-session
$ gcloud compute machine-types list
NAME            ZONE                    CPUS  MEMORY_GB  DEPRECATED
f1-micro        us-central1-f           1     0.60
g1-small        us-central1-f           1     1.70
n1-highcpu-16   us-central1-f           16    14.40
n1-highcpu-2    us-central1-f           2     1.80
n1-highcpu-32   us-central1-f           32    28.80
```

##### インスタンスの作成

次のコマンドでインスタンスを作成します。イメージは centos-7 、マシンタイプは最小の f1-micro を指定しています。ディスクサイズ無指定時は10GBとなります。

```
$ gcloud compute instances create インスタンス名 ¥
  --image-family centos-7 ¥
  --image-project centos-cloud ¥
  --machine-type f1-micro
```

※ --image エイリアスも使えるのですが、将来削除予定とワーニングが出力されます。

```
WARNING: Image aliases are deprecated and will be removed in a future version. Please use --image-family=centos-7 and --image-project=centos-cloud instead.
```

ちなみに作成にかかる時間を time コマンドで計測すると 16 秒ほどでした。

```
real	0m16.389s
user	0m0.620s
sys	0m0.131s
```

##### インスタンス作成時に有用なオプション

インスタンス作成の際には、他にも有用なオプションがあります。

|オプション|内容|
|---|---|
|--deletion-protection|インスタンスの削除をプロテクト|
|--no-address|グローバルアドレスを付与しない|
|--preemptible|プリエンプティブインスタンスにする|

### インスタンスへのログイン

インスタンスができたのでログインしてみましょう。次のコマンドでログインできます。

```
$ gcloud compute ssh インスタンス名
```

### インスタンスへのオペレーション

インスタンスの利用状況確認、起動、停止、削除などのオペレーションも、gcloud コマンドを通してできます。

##### 利用状況の確認

```shell-session
# インスタンスの一覧
$ gcloud compute instances list
NAME           ZONE           MACHINE_TYPE  PREEMPTIBLE  INTERNAL_IP  EXTERNAL_IP    STATUS
test-instance  us-central1-a  f1-micro                   10.128.0.2   35.x.x.x       RUNNING
...

# リージョンごとの利用状況
$ gcloud compute regions list
NAME                  CPUS  DISKS_GB  ADDRESSES  RESERVED_ADDRESSES  STATUS  TURNDOWN_DATE
asia-east1            0/8   0/2048    0/8        0/1                 UP
asia-northeast1       0/8   0/2048    0/8        0/1                 UP
...
```

##### インスタンスの起動停止

```shell-session
# インスタンスの起動
$ gcloud compute instances start インスタンス名

# インスタンスの停止
$ gcloud compute instances stop インスタンス名
```

##### インスタンスの削除

```
$ gcloud compute instances delete インスタンス名
```

### ファイルの送受信

gcloud compute scp を使います。

```
$ gcloud compute scp ~/Downloads/prometheus-2.0.0.linux-amd64.tar.gz instance-1:~/
```


##### ネットワークに関連する設定の情報取得

```
# ファイアウォールルール
$ gcloud compute firewall-rules list

# http(tcp:80ポート)を許可
$ gcloud compute firewall-rules create allow-http --description "Incoming http allowed." --allow tcp:80

# サブネット一覧
$ gcloud compute networks subnets list
```

# 参考

[ネットワークとファイアウォールの使用](https://cloud.google.com/compute/docs/networking?hl=ja)

[ファイアウォールルールの使用](https://cloud.google.com/compute/docs/vpc/using-firewalls?hl=ja)

[gcloudリファレンス](https://cloud.google.com/sdk/gcloud/reference/)
