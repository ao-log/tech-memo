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
# systemctl enable サービス名

// 自動起動の無効化
# systemctl disable サービス名

// 自動起動が有効になっているか確認
# systemctl is-enabled サービス名
```

### サービスの起動、停止

```shell-session
// サービスが実行中かどうかの確認
# systemctl is-active サービス名

// サービスの稼動状況確認
# systemctl status サービス名

// サービスの開始
# systemctl start サービス名

// サービスの停止
# systemctl stop サービス名

// サービスの再起動
# systemctl restart サービス名
```

### ログの確認

あるサービスのログだけを見たい時は、次のコマンドを実行します。

```
# journalctl -u サービス名
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
