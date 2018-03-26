# 導入

インストーラを使う。ダウンロードはこちらから。
https://www.anaconda.com/download/

# conda コマンド

##### パッケージを探す

```
$ conda search クエリ
```

文字列を含むパッケージ全てを抽出するので、 python だけを探したい時は、正規表現を使う。

```
$ conda search ^python$
```

Anaconda 自身もパッケージとして扱うことができる。

```
$ conda search ^anaconda$
```

##### 仮想環境の作成

```
$ conda create -n 仮想環境名
```

python 2.7.14 を入れたい場合は、次のように後方でパッケージを指定する。
この場合は、その環境には指定したパッケージとその他最低限のパッケージしかインストールされない。
インストールもすぐに完了する。

```
$ conda create -n py2.7.14 python=2.7.14
```

##### パッケージの追加

```
$ conda install パッケージ名
```

指定した仮想環境にのみインストールすることもできる。

```
$ conda install -n py2.7.14 numpy
```

##### パッケージの削除

```
$ conda remove パッケージ名
```

仮想環境を指定してアンインストールすることもできる。

```
$ conda remove -n py2.7.14 numpy
```

##### パッケージの一覧

```
$ conda list
```

指定した仮想環境のパッケージ一覧を見ることもできる。

```
$ conda list -n py2.7.14
```

##### 仮想環境の一覧

```
* が付いているのが今いる環境
$ conda env list
```

##### 仮想環境の削除

```
$ conda remove -n 仮想環境名 --all
```

##### 仮想環境の利用

```
// 仮想環境に入る
$ source activate 仮想環境名

// 仮想環境内で利用できるパッケージの一覧
$ pip list

// 仮想環境内でパッケージを追加できる。このパッケージは、他の環境には入らない。
// ただし、conda でインストールしたものとバッティングするので pip はあまり使わず、conda install で統一した方が良いと思う。

$ pip install tensorflow

// 仮想環境から抜ける
$ source deactivate
```

##### パッケージリストの管理

export

```
// base の場合は、仮想環境名に base を指定
$ conda env export -n 仮想環境名 | tee conda-env.yaml
```

環境作成

```
$ conda env create --file ./conda-env.yaml
```

## Anaconda Navigator

|アプリ|説明|
|---|---|
|jupyterlab|jupyter notebook の後継|
|jupyter notebook|ノートブック|
|qtconsole|GUIのiPython|
|spyder|開発環境|
