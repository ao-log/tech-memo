
### generate

##### コード一式生成

```
$ rails generate model User name:string email:string
```

##### migrate 用のファイル作成

```
$ rails generate migration add_index_to_users_email
```

下記の通り、記述する。

```ruby
class AddIndexToUsersEmail < ActiveRecord::Migration[5.1]
  def change
    add_index :users, :email, unique: true
  end
end
```

##### カラム追加

```
$ rails generate migration add_password_digest_to_users password_digest:string
```

下記内容が自動生成されている。

```ruby
class AddPasswordDigestToUsers < ActiveRecord::Migration[5.1]
  def change
    add_column :users, :password_digest, :string
  end
end
```


### db/migrate

##### migrate 関連のオペレーション

```
// DB 作成
$ rails db:create

// マイグレート
$ rails db:migrate

// ロールバック
$ rails db:rollback

// reset
$ rails db:migrate:reset

// マイグレート状況の確認
$ rake db:migrate:status
```

##### スキーマ

rails db:migrate 実行時に、```src/db/schema.rb``` に記録される。


##### fixtures

TODO:　```src/test/fixtures/users.yml```

### source

##### validates

```ruby
# メールアドレスのバリデーション。Rails Tutorial より。
  before_save { self.email = email.downcase }
  VALID_EMAIL_REGEX = /\A[\w+\-.]+@[a-z\d\-.]+\.[a-z]+\z/i
  validates :email, presence: true, length: { maximum: 255 },
                    format: { with: VALID_EMAIL_REGEX },
                    uniqueness: { case_sensitive: false }
```

model に password_digest 属性がある場合。ハッシュ化して保存。

```ruby
  has_secure_password
```

##### 1:N

```ruby
class User < ApplicationRecord
  has_many :microposts
end
```

##### reference

```ruby
class Micropost < ApplicationRecord
  belongs_to :user
end
```

### rails console で確認

```
$ rails console --sandbox
```

##### レコード作成

```ruby
> user = User.new(name: "hoge", email: "hoge@example.com")

# save も含めて実行する場合
> user = User.create(name: "hoge", email: "hoge@example.com")
```

```ruby
# 妥当性確認
> user.valid?
=> true
```

```ruby
# save
> user.save

# save 失敗した場合。user.save 実行前は空。実行後は以下の内容になる。
> user.errors.full_messages
=> ["Name can't be blank"]

> user.errors.any?         
=> true
```

```ruby
# 中身の確認
> user.name
=> "hoge"
> user.email
=> "hoge@example.com"

> puts user.attributes.to_yaml
---
id:
name: hoge
email: hoge@example.com
created_at:
updated_at:
password_digest:
=> nil

# user.attributes.to_yaml と同様
> y user.attributes

```

##### レコード更新

```ruby
$ user.update_attribute(:name, "fuga")
```

##### 検索

```ruby
# User から 1 件抽出
> User.first

# 全件抽出
> User.all

# レコード件数
> User.all.length

# 検索
> User.find_by(email: "hoge@example.com")
> User.find_by_email("hoge@example.com")
```

# 参考

[Active Record バリデーション](https://railsguides.jp/active_record_validations.html)

[Active Record の関連付け](https://railsguides.jp/association_basics.html#belongs-to%E9%96%A2%E9%80%A3%E4%BB%98%E3%81%91)

[Qiita:rails generate migrationコマンドまとめ](https://qiita.com/zaru/items/cde2c46b6126867a1a64)
