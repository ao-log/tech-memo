## コマンドの使い方

```
// リポジトリの初期化
git init

// リポジトリの状態確認
git status

// ステージ領域へファイル追加
git add ファイル名

// コミット
git commit -m "メッセージ"
// 直前のコミットメッセージを修正
git commit --amend

// コミットログの確認
git log
// グラフィカルにログ表示
git log --graph

// リポジトリのログの確認
git reflog

// コミットの統合。2 つめのコミットの pick を fixup に書き換える。
git rebase -i HEAD~2


// 差分確認
git diff
// 最新コミットとの差分
git diff HEAD
// 前のコミットとの差分
git diff HEAD^

// 指定したコミット ID の状態にする
git reset --hard コミットID
```

## ブランチ

```
// ブランチの表示
git branch
// リモートリポジトリを含めて表示
git branch -a

// ブランチ切り替え
git checkout ブランチ名

// ブランチを作成しつつ切り替え
git checkout -b ブランチ名

// マージ
git merge ブランチ名
```

## リモートリポジトリ

```
// リモートリポジトリを origin という名前で登録
git remote add origin リポジトリパス

// push。-u により upstream は master であることを設定。以降、git pull する際に origin の master から pull される。
git push -u origin master
```

