
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
# systemctl list-dependencies サービス名 --after

// 指定したサービスの後に起動する Unit
# systemctl list-dependencies サービス名 --before
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
