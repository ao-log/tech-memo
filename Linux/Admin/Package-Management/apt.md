
```shell
# 検索
$ apt search QUERY

# パッケージインストール
$ sudo apt install パッケージ

# パッケージの管理 DB を更新。パッケージの更新はしない。
$ sudo apt update

# インストールされているパッケージの更新
$ sudo apt upgrade

# パッケージの削除（設定ファイルも含めて）
$ sudo apt purge パッケージ

# パッケージ一覧
$ dpkg -l

# パッケージが含むファイル一覧
$ dpkg --listfiles パッケージ

# 鍵の一覧
$ apt-key list

# 鍵の削除
$ sudo apt-key del KEY
```

# 参考

[Qiita:[Ubuntu] apt-get まとめ](https://qiita.com/white_aspara25/items/723ae4ebf0bfefe2115c)
