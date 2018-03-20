## firewall-cmd

### 設定確認

##### デフォルトゾーンのみ

```
# firewall-cmd --list-all
trusted (active)
  target: ACCEPT
  icmp-block-inversion: no
  interfaces: eth0
  sources:
  services:
  ports:
  protocols:
  masquerade: no
  forward-ports:
  source-ports:
  icmp-blocks:
  rich rules:
```

##### 全てのゾーン

```
# firewall-cmd --list-all-zones
```

### 設定変更

##### trusted ゾーンに変更

アクセスを全許可する場合は、trusted ゾーンへ。内部セグメントのインタフェース向けの設定。

なお、外部公開しているインタフェースを trusted にしてはいけない。ただし、クラウドサービスなど上流でフィルタリングしている場合は trusted でもよい。

```
# firewall-cmd --permanent --zone=trusted --change-interface=インタフェース名
```

##### サービスへの接続許可

```
# firewall-cmd --permanent --zone=public --add-service=http
```

なお、サービスを定義したファイルは、```/usr/lib/firewalld/services/``` にある。例えば、http の場合は以下の通り。

```
$ cat /usr/lib/firewalld/services/http.xml
<?xml version="1.0" encoding="utf-8"?>
<service>
  <short>WWW (HTTP)</short>
  <description>HTTP is the protocol used to serve Web pages. If you plan to make your Web server publicly available, enable this option. This option is not required for viewing pages locally or developing Web pages.</description>
  <port protocol="tcp" port="80"/>
</service>
```

##### リッチルールの追加

(例) あるアドレスから udp 500 番ポートの通信を許可する場合

```
# firewall-cmd --permanent --zone=public --add-rich-rule="rule family=ipv4 source address=[SRC ADDRESS] port port=500 protocol=udp accept"
```

##### 設定の反映

```
# firewall-cmd --reload
```


# 参考

[ファイアウォールの使用](https://access.redhat.com/documentation/ja-jp/red_hat_enterprise_linux/7/html/security_guide/sec-using_firewalls)
