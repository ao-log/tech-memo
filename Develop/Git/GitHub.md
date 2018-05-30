# GitHub への push

### 前提

GitHub に鍵を登録しておく必要がある。

```
// リモートの指定
$ git remote add origin git@github.com:ao-log/tech-memo.git

// リモートが空でない場合は、pull しておく
$ git pull origin master

// push
$ git push -u origin master
```

# GitHub の Contribution が反映されない問題への対応

Author、メールアドレスが正しいかを確認する。
間違った情報でコミットしている場合、次の記事の方法で修正できる。

https://qiita.com/sea_mountain/items/d70216a5bc16a88ed932

# Wiki 内検索

Chrome 拡張が便利。
https://github.com/linyows/github-wiki-search
