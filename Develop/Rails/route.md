
### ルーティング

```
get  '/help', to: 'static_pages#help'
```

GETリクエストが /help に送信されたとき、StaticPages コントローラーの help アクションを呼び出す動作となる。


### Restful

```ruby
  resources :users
```

[Rails Tutorial](https://railstutorial.jp/chapters/sign_up?version=5.1#code-users_resource) より

| HTTPリクエスト | URL | アクション | 名前付きルート | 用途 |
|:--|:--|:--|:--|:--|
| GET | /users | index | users_path | すべてのユーザーを一覧するページ |
| GET | /users/1 | show | user_path(user) | 特定のユーザーを表示するページ |
| GET | /users/new | new | new_user_path | ユーザーを新規作成するページ (ユーザー登録) |
| POST | /users | create | users_path | ユーザーを作成するアクション |
| GET | /users/1/edit | edit | edit_user_path(user) | id=1のユーザーを編集するページ |
| PATCH | /users/1 | update | user_path(user) | ユーザーを更新するアクション |
| DELETE | /users/1 | destroy | user_path(user) | ユーザーを削除するアクション |
