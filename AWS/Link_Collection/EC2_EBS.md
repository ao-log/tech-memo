

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


