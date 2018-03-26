
# layout

コントローラに対応するレイアウトファイルがない場合は ```app/views/layouts/application.html.erb``` を使用。

```html
<!DOCTYPE html>
<html>
  <head>
    <title><%= full_title(yield(:title)) %></title>
    <%= render 'layouts/rails_default' %>
    <%= render 'layouts/shim' %>
  </head>
  <body>
    <%= render 'layouts/header' %>
    <div class="container">
      <%= yield %>
      <%= render 'layouts/footer' %>
    </div>
  </body>
</html>
```

```<%= yield(:title) %>``` に値を渡すには、view 側のファイルで provide を用いる。``` app/views/static_pages/home.html.erb ```

```
<% provide(:title, "Home") %>
```

# パーシャル

```
<%= render 'layouts/shim' %>
```

```app/views/layouts/_shim.html.erb``` というファイルを探してその内容を評価し、結果をビューに挿入
パーシャルはファイル名の先頭に「_」をつける。

# リンク

```
<%= link_to "About", about_path %>
```

# フォーム

```html
    <%= form_for(@user) do |f| %>
      <%= f.label :name %>
      <%= f.text_field :name %>

      <%= f.label :email %>
      <%= f.email_field :email %>

      <%= f.label :password %>
      <%= f.password_field :password %>

      <%= f.label :password_confirmation, "Confirmation" %>
      <%= f.password_field :password_confirmation %>

      <%= f.submit "Create my account", class: "btn btn-primary" %>
    <% end %>
```

### pluralize

複数形に対応してくれる。

```
>> helper.pluralize(1, "error")
=> "1 error"
>> helper.pluralize(5, "error")
=> "5 errors"
```

# ファイル置き場

|パス|説明|
|---|---|
|app/assets/stylesheets/|スタイルシート|
|app/assets/javascripts/|JavaScript|
|app/assets/images/|イメージファイル|

# asset pipeline とは？

スタイルシートを1つのCSSファイル (application.css) にまとめ、
JavaScriptファイルを1つのJSファイル (javascripts.js) にまとめる。
実際にアセットをまとめる処理を行うのはSprocketsというgem。

# デバッグ

```
<%= debug(params) if Rails.env.development? %>
```

# Bootstrap

##### Gemfile

```
gem 'bootstrap-sass'
gem 'sass-rails'
```

上記を記述し、bundle install

##### app/assets/stylesheets/custom.scss

```
@import "bootstrap-sprockets";
@import "bootstrap";
```

# 参考

[Rails:レイアウトとレンダリング](https://railsguides.jp/layouts_and_rendering.html)
