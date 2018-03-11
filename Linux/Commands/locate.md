
パターンにマッチするファイルを探すコマンド。

```
$ locate conf
/bin/hciconfig
/boot/config-4.13.0-32-generic
/boot/config-4.13.0-36-generic
/boot/grub/x86_64-efi/configfile.mod
/etc/adduser.conf
...
```

正規表現も使える。

```
$ locate -r conf$
/etc/adduser.conf
/etc/apg.conf
/etc/appstream.conf
/etc/brltty.conf
/etc/ca-certificates.conf
...
```
