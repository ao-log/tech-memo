# diff

```
// インデックスとの差分（まだ、git add していないもの）
$ git diff

// パスを付与するとそのパスのファイルのみが対象となる
$ git diff PATH

// 最新コミットと git add したものの差分
$ git diff --cached

// 最新コミットとの差分
$ git diff HEAD
```

##### 便利なオプション

|オプション|説明|
|---|---|
|--name-only|ファイル名のみ|

##### コミット間の差分

```
// 指定したコミットとの差分
$ git diff COMMIT

// 最新コミットと、その一つ前のコミット間の差分
$ git diff HEAD HEAD^

// 2 つのコミット間の差分
$ git diff COMMIT1 COMMIT2
```

##### 参考

[git diff を徹底攻略！よく使う便利オプションまとめ、完全版。](http://www-creators.com/archives/755)
