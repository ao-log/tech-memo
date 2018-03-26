
### 式展開

```ruby
> f_name = "Michael"
> "#{f_name} Hartl"

> a.each { |s| puts "a #{s}" }
a bar
a baz
a foo
```

### 配列

```ruby
> a[1..3]
=> ["baz", "foo", "neko"]

> %w[a b c].map { |char| char.upcase }
```

### 連想配列

```ruby
> user = { name: "user", email: "user@example.com" }
=> {:name=>"user", :email=>"user@example.com"}
> user[:name]
=> "user"

> user.each do |key, value|
    "Key = #{key}, Value = #{value}"
  end
```

### 論理値

!! を用いることで、論理値に変換することができる。

```
!!user.authenticate("foobar")
```

### メモ化

```ruby
User.find_by(id: session[:user_id])
```

インスタンス変数に格納するようにする。

```ruby
if @current_user.nil?
  @current_user = User.find_by(id: session[:user_id])
else
  @current_user
end
```

上記コードは以下のようにも書ける。Ruby でよく使うイディオムらしい。
利用シーンとしては、nil なら代入し、そうでないなら何もしない場合。

```
@current_user ||= User.find_by(id: session[:user_id])
```
