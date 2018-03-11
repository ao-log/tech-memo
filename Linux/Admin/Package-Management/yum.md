Yum について、よく使うコマンド、設定をまとめました。

# 想定環境

RHEL 7, CentOS 7

# yum コマンド関連

### パッケージアップデート

OS インストール後は、セキュリティパッチを適用するために、パッケージアップデートするのが鉄板です。

```
// アップデート可能なもの全て
# yum update

// セキュリティ更新を含むアップデートのみ(セキュリティに関するアップデートのみ。バグフィックスは適用しない)
# yum update-minimal --security

// セキュリティ更新を含むアップデートのみ(セキュリティに関するアップデートを行い、最新にする)
# yum update --security
```

### 使うリポジトリを指定したインストール

```
# yum install --disablerepo \* --enablerepo epel,base パッケージ名
```

### グループインストール

```
// 一覧を取得
# yum groups list

// インストール
# yum groups install グループ名
```

### パッケージを探す

```
// クエリ文字列にマッチするパッケージ一覧を取得
# yum search QUERY

// コマンド名などから、それを含むパッケージ名を取得
# yum provides FEATURE
```

### パッケージに含まれるファイル一覧を取得

```
// パッケージ未インストールの場合にも一覧取得可能
# repoquery --list パッケージ名
```

### パッケージの削除

```
# yum remove パッケージ名
```

### キャッシュの消去

古い不正な情報が残っていて、インストール失敗する場合に実行する場合に使用。

```
yum clean all
```

##### 使用履歴

```
# yum history

// undo もできる…が、今の時代なら、仮想マシンのスナップショットや、
// Immutable Infrastructure のようなサーバ管理をしたほうがよいかも。
# yum history undo 履歴番号
```

# リポジトリ関連

### リポジトリの一覧

```
// enableのみ
# yum repolist

// 全て
# yum repolist all
```

### リポジトリの有効化/無効化

```
// 有効化
# yum-config-manager --enable リポジトリ名
// 無効化
# yum-config-manager --disable リポジトリ名
```

# rpm コマンド

rpm コマンドは、個々のパッケージの情報を得たいときに。

### パッケージの一覧

```
$ rpm -qa
```

### パッケージの情報

```
$ rpm -qi パッケージ名
```

### パッケージが含むファイルの一覧

```
$ rpm -ql パッケージ名
```

# その他

### epel リポジトリの追加

```
# yum install epel-release
```

### proxy の設定

```/etc/yum.conf``` に記述する。

```
proxy=http://proxy.example.com:8080
```

### インストール用 ISO イメージをリポジトリ化

ISO ファイルをマウントする。

```
$ cp mount_dir/media.repo /etc/yum.repos.d/dvd.repo
```

```/etc/yum.repos.d/dvd.repo``` ファイルを編集。

```
baseurl=file:///mount_dir
```

### RPM ファイルをカレントに展開

```
# rpm2cpio RPMファイル | cpio -idv
```

# 参考

[システム管理者のガイド: 第8章 Yum](https://access.redhat.com/documentation/ja-jp/red_hat_enterprise_linux/7/html/system_administrators_guide/ch-yum)
