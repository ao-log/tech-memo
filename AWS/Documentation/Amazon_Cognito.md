
[Amazon Cognito とは](https://docs.aws.amazon.com/ja_jp/cognito/latest/developerguide/what-is-amazon-cognito.html)

* ユーザープール: ユーザーディレクトリ
* ID プール: AWS サービスにアクセスするための一時認証情報を取得できる


## チュートリアル

[ブラウザからの Amazon S3 バケット内の写真の表示](https://xn--docs-u53c.aws.amazon.com/ja_jp/sdk-for-javascript/v2/developer-guide/s3-example-photos-view.html)

以下のコードで認証情報を取得。
```js
AWS.config.credentials = new AWS.CognitoIdentityCredentials({
    IdentityPoolId: 'IDENTITY_POOL_ID',
});
```



# BlackBelt

[Amazon Cognito](https://pages.awscloud.com/rs/112-TZM-766/images/20200630_AWS_BlackBelt_Amazon_Cognito_ver2.pdf)

* ユーザプール
  * ユーザ
    * 管理者がユーザを作成、もしくはユーザ自身にサインアップを許可することも可能
    * サインアップ後は管理者が確認 もしくは E メール、SMS による確認ができれば CONFIRMED になる
    * グループに属することができる



# 参考

* Document
  * [Amazon Cognito とは](https://docs.aws.amazon.com/ja_jp/cognito/latest/developerguide/what-is-amazon-cognito.html)
* Black Belt
  * [Amazon Cognito](https://pages.awscloud.com/rs/112-TZM-766/images/20200630_AWS_BlackBelt_Amazon_Cognito_ver2.pdf)


