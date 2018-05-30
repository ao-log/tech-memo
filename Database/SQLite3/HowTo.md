### データベース作成

```
$ sqlite3 user.sqlite
```

### テーブル

##### テーブル一覧

```
> .tables
```

##### テーブル作成

テーブル作成の SQK を書いたファイルを作成。

・create.sql

```
create table users
(
    name text,
    hobby text
);
```

読み込み。

```
> .read create.sql
```

##### スキーマ確認

```
> .schema users
```

### import

```
> .import CSVファイル テーブル
```

### 操作終了

```
> .quit
```
