
### Rails アプリケーションの作成

```
$ rails new
```

### rails server

```
$ rails server -p PORT -e ENVIRONMENT
```

### rails generate

##### scaffolod

```
$ rails generate scaffold User name:string email:string
```

##### controler

```
書式: generate controller ControllerName action1 action2

$ rails generate controller StaticPages home help
```

##### integration test

```
$ rails generate integration_test site_layout
      invoke  test_unit
      create    test/integration/site_layout_test.rb
```

### db

```
$ rails db:migrate
```

### test

```
$ rails test

# model のみ
$ rails test:models

# integration test
$ rails test:integration
```

### console

```
$ rails console

# サンドボックスモード
$ rails console --sandbox
```

### ルーティング

```
rails routes
```

# 参考

[Railsのコマンドラインツール](https://railsguides.jp/command_line.html)
