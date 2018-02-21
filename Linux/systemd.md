systemd についてのサーバ構築、運用時によく使うコマンド、TIPS をまとめました。

ゼロから体系的に学ぶための記事ではありません。それらはいい記事、スライドがあるので、紹介させていただきます。

* [RHEL:第9章 SYSTEMD によるサービス管理](https://access.redhat.com/documentation/ja-jp/red_hat_enterprise_linux/7/html/system_administrators_guide/chap-managing_services_with_systemd)
* [Linux女子部 systemd徹底入門](https://www.slideshare.net/enakai/linux-27872553)

では、よく使うコマンド、TIPS をまとめさせて頂きます。

# systemctl コマンド

### サービス一覧

```--type service``` をオプションにつけるとサービスだけ抽出してくれます。

```shell-session
// ロード済みサービスの一覧
# systemctl list-units --type service

// サービスの一覧
# systemctl list-unit-files --type service
```

### 自動起動

```shell-session
// 自動起動の有効化
# systemctl enable [サービス名]

// 自動起動の無効化
# systemctl disable [サービス名]

// 自動起動が有効になっているか確認
# systemctl is-enabled [サービス名]
```

### サービスの起動、停止

```shell-session
// サービスが実行中かどうかの確認
# systemctl is-active [サービス名]

// サービスの稼動状況確認
# systemctl status [サービス名]

// サービスの開始
# systemctl start [サービス名]

// サービスの停止
# systemctl stop [サービス名]

// サービスの再起動
# systemctl restart [サービス名]
```

### ログの確認

あるサービスのログだけを見たい時は、次のコマンドを実行します。

```
# journalctl -u [サービス名]
```

### 設定ファイルの確認

個々のサービスなどの設定を行うファイルを **Unitファイル** と呼びます。Unit ファイルを開かずとも、systemctl コマンドでファイルパスとその中身を確認できます。


```shell-session
# systemctl cat sshd
```

```ini
# /usr/lib/systemd/system/sshd.service
[Unit]
Description=OpenSSH server daemon
Documentation=man:sshd(8) man:sshd_config(5)
After=network.target sshd-keygen.service
(以下略)
```


# 自作サービスを作る

デーモン化したいスクリプトを ```/opt/sample.sh``` とすると…次の通りサービス起動に関する設定ファイルを書きます。
※ graphical.target の場合は、WantedBy の行を graphical.target に書きかえてください。

```ini:/etc/systemd/system/sample.service
[Unit]
Description = sample daemon

[Service]
ExecStart = /opt/sample.sh
Type = simple

[Install]
WantedBy = multi-user.target
```

次のコマンドでサービス起動します。

```shell-session
# systemctl start sample
```

OS 起動時に自動起動する場合は、enable にします。

```shell-session
# systemctl enable sample
```

Unit ファイルを編集した場合は、次のコマンドで反映します。

```shell-session
# systemctl daemon-reload
```

### プロセスが停止した時の動作を指定

unit ファイル中の Restart はプロセスが停止した時の動作を指定できます。```systemctl restart [サービス名]```を実行した時の挙動には関係のないパラメータです。デフォルトは「no」で何もしません。always にするとプロセス停止時に再起動を試みてくれます。

```
[Service]
Restart=always
```

Restart のその他の設定値の説明は [systemd.service の man ページ](https://www.freedesktop.org/software/systemd/man/systemd.service.html#) の「Restart=」の所をご参照ください。


# Unit ファイルの知識

### Unit タイプ

実は、サービス以外にもいろいろあります。

| Unit タイプ | ファイル拡張子 | 詳細 |
|:--|:--|:--|
| Service unit | .service | システムサービス |
| Target unit | .target | systemd unit のグループ |
| Automount unit | .automount | ファイルシステムの自動マウントポイント |
| Device unit | .device | カーネルが認識するデバイスファイル |
| Mount unit | .mount | ファイルシステムのマウントポイント |
| 他にもいくつかあります | ... | ... |

[RHEL:システム管理者のガイド - 表9.1 利用可能な systemd Unit タイプより。一部改変。](https://access.redhat.com/documentation/ja-jp/red_hat_enterprise_linux/7/html/system_administrators_guide/chap-Managing_Services_with_systemd#tabl-Managing_Services_with_systemd-Introduction-Units-Types)

といってもサーバ管理者が扱うのは、ほぼサービスになるかと思います。

### ユニットファイルの場所とファイル名

Unit ファイルの置き場所は次のように使い分けます。

|パス|説明|
|---|---|
|/usr/lib/systemd/system|パッケージで配布されたUnitファイルの配置場所|
|/etc/systemd/system|主にユーザ作成のUnit置き場。上記ディレクトリよりも優先されます|


Unit ファイルの起動順序は Unit ファイル中で「Before=」「After=」により指定します。例えば、sshd サービスの場合を見てみます。

```shell-session
# systemctl cat sshd
```

```ini
[Unit]
After=network.target sshd-keygen.service
```

あるサービスの前後の Unit をツリー表示することができます。--after, --before が直感と違うかもしれないのでご注意ください。

```shell-session
// 指定したサービスの先に起動する Unit
# systemctl list-dependencies [サービス名] --after

// 指定したサービスの後に起動する Unit
# systemctl list-dependencies [サービス名] --before
```




# ターゲット

ターゲットとは何か。いわば、ランレベルの事です。

|ランレベル|target|説明|
|---|---|---|
|0|poweroff.target|システム停止|
|1|rescue.target|シングルユーザーモード|
|2,3,4|multi-user.target|マルチユーザーモード|
|5|graphical.target|GUIモード|
|6|reboot.target|システム再起動|
|-|emergency.target|緊急モード|

### モード変更方法

システムのシャットダウンは systemctl を使って行うこともできます。
指定した target に該当する処理を呼び出しています。

```shell-session
// シャットダウンして電源を切る
# systemctl poweroff

// 再起動
# systemctl reboot
```

次のようなモード変更もできます。これらの操作は ssh ではなく、コンソール上で行いましょう。

```
// シングルユーザモードへ
# systemctl rescue

// 緊急モードへ
# systemctl emergency
```


### デフォルトターゲット

デフォルトターゲットとは、OS 起動時のランレベルです。multi-user ターゲットか、graphical ターゲットのどちらかになっていると思います。まず、現在何がデフォルトターゲットになっているかを確認します。

```shell-session
$ systemctl get-default
multi-user.target
```

multi-user ターゲットから graphical ターゲットに変更する場合は、次のコマンドを実行します。

```shell-session
# systemctl set-default graphical.target
Removed symlink /etc/systemd/system/default.target.
Created symlink from /etc/systemd/system/default.target to /usr/lib/systemd/system/graphical.target.
```

シンボリックリンクの差し替えを行っていると分かります。OS 起動時には ```default.target``` ファイルを確認し、Unit の起動処理を進めていきます。

では、例えば multi-user.target はどの Unit を処理していくか。ディレクトリを確認してみます。

```shell-session
$ ls -l /etc/systemd/system/multi-user.target.wants
合計 0
lrwxrwxrwx. 1 root root 46 10月 25 19:12 NetworkManager.service -> /usr/lib/systemd/system/NetworkManager.service
lrwxrwxrwx. 1 root root 37 10月 25 19:12 acpid.service -> /usr/lib/systemd/system/acpid.service
lrwxrwxrwx. 1 root root 38 10月 25 19:12 auditd.service -> /usr/lib/systemd/system/auditd.service
lrwxrwxrwx. 1 root root 39  2月 18 11:00 chronyd.service -> /usr/lib/systemd/system/chronyd.service
...
```

ここにある Unit ファイルを処理します。サービスの自動起動は ```systemctl [disable | enable]``` で切り替えていましたが、このコマンドでシンボリックリンクを貼ったり削除したりすることで、どのサービスを起動するかを制御しているわけです。

```shell-session
# systemctl disable chronyd
Removed symlink /etc/systemd/system/multi-user.target.wants/chronyd.service.

# systemctl enable chronyd
Created symlink from /etc/systemd/system/multi-user.target.wants/chronyd.service to ¥
/usr/lib/systemd/system/chronyd.service.
```



# 参考

[RHEL:第9章 SYSTEMD によるサービス管理](https://access.redhat.com/documentation/ja-jp/red_hat_enterprise_linux/7/html/system_administrators_guide/chap-managing_services_with_systemd)

[RHEL:9.6. システムのユニットファイルの作成および変更](https://access.redhat.com/documentation/ja-jp/red_hat_enterprise_linux/7/html/system_administrators_guide/sect-managing_services_with_systemd-unit_files)

[めもめも：Systemd入門(4) - serviceタイプUnitの設定ファイル](http://enakai00.hatenablog.com/entry/20130917/1379374797)

[Linux女子部 systemd徹底入門](https://www.slideshare.net/enakai/linux-27872553)

[Qiita:Systemd メモ書き](https://qiita.com/a_yasui/items/f2d8b57aa616e523ede4)

[Qiita:Systemdを使ってさくっと自作コマンドをサービス化してみる](https://qiita.com/DQNEO/items/0b5d0bc5d3cf407cb7ff)

[書籍:プロのためのLinuxシステム構築・運用技術](https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%81%AE%E3%81%9F%E3%82%81%E3%81%AE-Linux%E3%82%B7%E3%82%B9%E3%83%86%E3%83%A0%E6%A7%8B%E7%AF%89%E3%83%BB%E9%81%8B%E7%94%A8%E6%8A%80%E8%A1%93-Software-Design-plus/dp/4774145017)

[Wikipedia:systemd](https://ja.wikipedia.org/wiki/Systemd)
