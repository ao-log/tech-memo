
### Rails Tutorial より

```ruby
  test "should get about" do
    get about_path
    assert_response :success
    assert_select "title", "About | #{@base_title}"
  end
```

##### integration test

```ruby
  test "layout links" do
    get root_path
    assert_template 'static_pages/home'
    assert_select "a[href=?]", root_path, count: 2
    assert_select "a[href=?]", help_path
    assert_select "a[href=?]", about_path
    assert_select "a[href=?]", contact_path
  end
```

# 参考

[Rails:テスティングガイド](https://railsguides.jp/testing.html)
