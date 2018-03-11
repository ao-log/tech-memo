# 自作サービスを作る

デーモン化したいスクリプトを ```/opt/sample.sh``` とすると…次の通りサービス起動に関する設定ファイルを書きます。
※ graphical.target の場合は、WantedBy の行を graphical.target に書きかえてください。

・/etc/systemd/system/sample.service

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
