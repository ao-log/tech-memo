##### 仮想イメージ選定

次のサイトから用途に合うものを探します。

https://app.vagrantup.com/boxes/search

例えば、CentOS 7 の場合はこちら。

https://app.vagrantup.com/centos/boxes/7


```
$ mkdir centos7 && cd centos7
```

##### Vagrantfile を生成

```
$ vagrant init centos/7
```

##### 仮想ゲストの起動

```
$ vagrant up --provider=libvirt
```

##### 仮想マシンへの接続

```
$ vagrant ssh
```

##### 稼働状況確認

```
$ vagrant status
```

##### シャットダウン

```
$ vagrant halt
```

### Vagrantfile

```vagrant init``` コマンドで自動生成されるものを参考にする。

```
# ポートのマッピング
config.vm.network "forwarded_port", guest: 80, host: 8080

# IPアドレスの固定
config.vm.network "private_network", ip: "192.168.33.10"

# リソース
config.vm.provider "libvirt" do |vb|
  vb.memory = "1024"
end
```
