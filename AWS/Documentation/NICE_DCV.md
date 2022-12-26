
NICE DCV は AWS が開発している製品ではないが、AWS 上で使用できるようになっており、ドキュメントも用意されている。

要点のみ、おさえていく。

[NICE DCV 管理者ガイド](https://docs.aws.amazon.com/ja_jp/dcv/latest/adminguide/what-is-dcv.html)


[インストール](https://docs.aws.amazon.com/ja_jp/dcv/latest/adminguide/setting-up-installing.html)

前提条件、インストール方法、インストール後のチェック方法がまとめられている。


[NICE DCV サーバのライセンス](https://docs.aws.amazon.com/ja_jp/dcv/latest/adminguide/setting-up-license.html)

定期的に S3 バケットに接続して、有効なライセンスがあるかどうかを判断している。
そのバケットに対する s3:GetObject の許可が必要。


[NICE DCV サーバの起動](https://docs.aws.amazon.com/ja_jp/dcv/latest/adminguide/manage-start.html)

dcvserver サービスを起動すれば良い。


[NICE DCV サーバーの TCP ポートの変更](https://docs.aws.amazon.com/ja_jp/dcv/latest/adminguide/manage-port.html)

デフォルトは 8443。1024 以上の番号を使用する必要がある。


[NICE DCV 認証の設定](https://docs.aws.amazon.com/ja_jp/dcv/latest/adminguide/security-authentication.html)

デフォルトは system。認証が OS に一任される。


[NICE DCV セッションの管理](https://docs.aws.amazon.com/ja_jp/dcv/latest/adminguide/managing-sessions.html)

コンソールとバーチャルの 2 種類ある。


[NICE DCV セッションの開始](https://docs.aws.amazon.com/ja_jp/dcv/latest/adminguide/managing-sessions-start.html)

次のようなコマンドでセッションを開始できる。

```
sudo dcv create-session --owner dcv-user --user dcv-user my-session
```


[NICE DCV サーバパラメータリファレンス](https://docs.aws.amazon.com/ja_jp/dcv/latest/adminguide/config-param-ref.html)

