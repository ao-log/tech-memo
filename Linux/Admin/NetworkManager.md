
オプションは省略できるが、かぶった場合（s - status, show、d - down, delete）に
意図した通りのオプションが選ばれない場合があるので注意。

```
// デバイス一覧
# nmcli device

// デバイスの接続
# nmcli device connect デバイス名

// デバイスの切断
# nmcli device disconnect デバイス名

// コネクションの一覧
# nmcli connection show

// コネクションの詳細表示
# nmcli connection show コネクション名
// 詳細表示。ipv4 の出力のみ。
# nmcli -fields ipv4 c show コネクション名

// ifup
# nmcli connection up コネクション名

// ifdown
# nmcli connection down コネクション名
```

```
// コネクション名変更(この例だと System eth0 → eth0)
# nmcli con mod "System eth0" connection.id eth0

// IPv4 アドレス変更
# nmcli con mod コネクション名 ipv4.addresses アドレス

// アドレス設定方法。Static or DHCP。
# nmcli con mod コネクション名 ipv4.method [manual|auto]

// インタフェースを OS 起動時に有効化。
# nmcli con mod コネクション名 connection.autoconnect yes

// 設定を消す場合には - (マイナス)をつける。
# nmcli con mod コネクション名 -ipv4.gateway 192.168.0.254
```

# 参考

[RHEL:NETWORKMANAGER のコマンドラインツール NMCLI の使用](https://access.redhat.com/documentation/ja-jp/red_hat_enterprise_linux/7/html/networking_guide/sec-Using_the_NetworkManager_Command_Line_Tool_nmcli)

[めもめも:nmcliで仮想ブリッジ作成](http://enakai00.hatenablog.com/entry/20141121/1416551748)
