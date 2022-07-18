
[[待望のアプデ]EC2インスタンスメタデータサービスv2がリリースされてSSRF脆弱性等への攻撃に対するセキュリティが強化されました！](https://dev.classmethod.jp/articles/ec2-imdsv2-release/)

* IMDSv1 の場合は GET メソッドでメタデータの取得ができる。
```
curl http://169.254.169.254/latest/meta-data/security-credentials/
```

* IMDSv2 の場合は事前に PUT で Token を取得する必要がある。
* v2 のみに強制できる。
* メタデータサービスを無効化できる。
* メタデータレスポンスの TTL を短くし複数ホストを経由した取得を防止できる



[EBSボリューム作成時のべき等性を担保できるようになりました](https://dev.classmethod.jp/articles/ensuring-idempotency-for-ebs/)

AWS CLI で create-volume する際に client-token を設定しておくことで冪等性を担保できる。
もう一度同じコマンドを実行しても別のボリュームが作成されない。

```
aws ec2 create-volume \
    --volume-type gp3 \
    --size 10 \
    --availability-zone ap-northeast-1a \
    --client-token 550e8400-e29b-41d4-a716-446655440000
```

コマンドの応答がなかったけど実はリソースが作成されている場合もあり、そういった場合の対策にもなる。


