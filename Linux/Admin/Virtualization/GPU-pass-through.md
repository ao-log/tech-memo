[RHEL:17.7. GPU デバイスの割り当て](https://access.redhat.com/documentation/ja-jp/red_hat_enterprise_linux/7/html/virtualization_deployment_and_administration_guide/sect-device-gpu)
にほぼ全て書かれているが、自分用に簡易メモを残しておく。

### 確認した環境

ホストOS: RHEL 7.2
ゲストOS: Windows Server 2016

### 大雑把な流れ

1. ホストOS側で IOMMU サポートを有効化
2. ホストOSから GPU をデタッチ
3. 仮想ゲストに GPU をアタッチ
4. 仮想ゲスト上で GPU ドライバをインストール

### カーネルの IOMMU サポートを有効にする

・/etc/sysconfig/grub

```
# Intel VT-d システムの場合
GRUB_CMDLINE_LINUX="... intel_iommu=on iommu=pt"
```

ブートローダー設定の再生成

```
# grub2-mkconfig -o /etc/grub2.cfg
```

再起動し、有効化。

```
# reboot
```

### GPU をホストのバインドから除外

まず、PCI バスアドレス、デバイス ID を特定。

```
# lspci -Dnn | grep VGA
0000:02:00.0 VGA compatible controller [0300]: NVIDIA Corporation GK106GL [Quadro K4000] [10de:11fa] (rev a1)

→ PCI バスアドレス = 0000:02:00.0
　　　　デバイス ID = 10de:11fa
```

・/etc/sysconfig/grub

```
GRUB_CMDLINE_LINUX="... pci-stub.ids=10de:11fa"
```

ブートローダー設定の再生成

```
# grub2-mkconfig -o /etc/grub2.cfg
```

再起動し、有効化。

```
# reboot
```

### PCI デバイスを仮想ゲストにアタッチ

デバイスに関する情報を XML で出力する。

```
# virsh nodedev-dumpxml pci_0000_02_00_0
...
必要になるのはこの部分。

    <iommuGroup number='13'>
      <address domain='0x0000' bus='0x02' slot='0x00' function='0x0'/>
      <address domain='0x0000' bus='0x02' slot='0x00' function='0x1'/>
    </iommuGroup>
```

アタッチ用の XML ファイルを作成。

```
<hostdev mode='subsystem' type='pci' managed='yes'>
  <driver name='vfio'/>
  <source>
    <address domain='0x0000' bus='0x02' slot='0x00' function='0x0'/>
  </source>
</hostdev>
```

先ほど作成した XML ファイルを指定し、アタッチ。

```
# virsh attach-device ドメイン XMLファイル --persistent
```

### 仮想ゲスト側の作業

仮想ゲスト OS 上で GPU ドライバをインストールすれば、仮想ゲストで GPU が認識されるようになる。
