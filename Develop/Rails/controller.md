
# Controller の役割と関連

* action (new, show など)を実装している
* アクションを呼び出すと、対応するビューを表示する
* /users への POST リクエストは create アクションに送られる
* redirect_to @user は redirect_to user_url(@user) と等価。つまり show へのリダイレクト。
* helper で実装した関数を使用可能
* ユーザ登録失敗時は new を再描画。ただし、@user の持つエラー情報をもとに、エラー時用のパーシャルをページ上部に表示。

※_pathと_urlの違いは、_urlは完全なURLの文字列を返す。上記で redirect_to で _url になっているのは、HTTPの標準としてはリダイレクトのときに完全なURLが要求されるため。

[Rails Tutorial](https://railstutorial.jp/chapters/sign_up?version=5.1#code-create_action_strong_parameters) より

```ruby
class UsersController < ApplicationController
  def show
    @user = User.find(params[:id])
  end

  def new
    @user = User.new
  end

  def create
    @user = User.new(user_params)
    if @user.save
      flash[:success] = "Welcome to the Sample App!"
      redirect_to @user
    else
      render 'new'
    end
  end

  private

    def user_params
      params.require(:user).permit(:name, :email, :password,
                                   :password_confirmation)
    end
end
```

# Strong Parameters

[Rails Tutorial](https://railstutorial.jp/chapters/sign_up?version=5.1#code-create_action_strong_parameters) より

```params[:user]``` をそのまま使うと攻撃者が細工したリクエストをそのまま受け付けてしまう。
対策として、Strong Parameters を使うことで、必須のパラメータと許可されたパラメータを指定する。
user_param　メソッドはプライベートにする。

```ruby
  def create
    @user = User.new(user_params)
    if @user.save
      # 保存の成功をここで扱う。
    ...
  end

  private

    def user_params
      params.require(:user).permit(:name, :email, :password,
                                   :password_confirmation)
    end
```

# 参考

[Rails ドキュメント:コントローラ](http://railsdoc.com/controller)

[Rails:Railsでparamsを使ってデータを取得する](https://qiita.com/To_BB/items/fe9cada1a0bcfe5e3efb)
