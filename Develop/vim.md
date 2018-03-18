# vim

### カーソル移動

|キー|説明|
|--|--|
|h|左|
|j|下|
|k|上|
|l|右|
|^|行頭|
|$|行末尾|
|b|前の単語|
|w|後ろの単語|
|B|前の単語(記号も含む)|
|W|後ろの単語(記号も含む)|
|Ctrl+f|ページ送り|
|Ctrl+b|ページ送り（back）|
|gg|ファイルの先頭|
|G|ファイルの末尾|

# vimrc

```
" ハイライト
syntax enable

" 非 vi 互換モードへ
" 方向キーでABCDが入力される問題への対処。
" ただ、/.vimrc が存在する時点で、 set noconpatible になっているらしい。書かなくてもいいかも。
set nocompatible
 
" autoindent: 改行時にインデントを継続
" tabstop: タブ時に置き換えられる空白文字の個数
" expandtab: タブ入力を複数の空白入力に置き換え
set autoindent tabstop=2 expandtab 

" マウスの有効化
set mouse=a
```

# vimrc のコマンドがエラーになる場合

エラーの内容。

```
Sorry, the command is not available in this version: syntax enable
```

パッケージを確認する。tiny 版になっている。

```
$ dpkg -l | grep vim
ii  vim-common                                        2:7.4.1689-3ubuntu1.2                                     amd64        Vi IMproved - Common files
ii  vim-tiny                                          2:7.4.1689-3ubuntu1.2                                     amd64        Vi IMproved - enhanced vi editor - compact version
```

次のコマンドで vim をインストール。

```
$ sudo apt-get install vim
```

パッケージが追加されている。これで、上記エラーが取れるはず。

```
$ dpkg -l | grep vim
ii  vim                                               2:7.4.1689-3ubuntu1.2                                     amd64        Vi IMproved - enhanced vi editor
ii  vim-common                                        2:7.4.1689-3ubuntu1.2                                     amd64        Vi IMproved - Common files
ii  vim-runtime                                       2:7.4.1689-3ubuntu1.2                                     all          Vi IMproved - Runtime files
ii  vim-tiny                                          2:7.4.1689-3ubuntu1.2                                     amd64        Vi IMproved - enhanced vi editor - compact version
```


