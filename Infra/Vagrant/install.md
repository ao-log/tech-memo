## 前提

* CPU が仮想化支援機能に対応していること。
* BIOS で仮想化支援機能を有効化しておくこと。

## 当記事で検証した環境

* elementary OS 0.4.1 Loki (16.04.1-Ubuntu)

## 準備

Vagrant は仮想化ソフトとしてデフォルトで VirtualBox を使用します。
ただし、OS のパッケージマネージャでインストールしたものを使うと、仮想ゲスト起動時にホスト OS がフリーズしました。
VirtualBox のバージョンを上げる手もありますが、ここでは KVM を採用します。

#### KVM

関連パッケージのインストール。

```
$ sudo apt install qemu-kvm
$ sudo apt install libvirt-bin
$ sudo apt install libvirt-dev
```

Vagrant を利用するユーザを libvirtd グループに属させます。

```
$ sudo usermod -aG libvirtd ユーザ名
```


#### Vagrant

ubuntu のパッケージマネージャのものではなく、こちらで公開されているものをインストールして使います。

https://www.vagrantup.com/downloads.html

理由は vagrant plugin コマンドが失敗したため。

##### プラグイン追加

```shell
# libvirt 用。
$ vagrant plugin install vagrant-libvirt

# box を libvirt に変換する用途。
$ vagrant plugin install vagrant-mutate

# スナップショット用
$ vagrant plugin install sahara
```
