
### ポートフォワード

以下のように接続元サーバから中継サーバを通して転送先サーバへポートフォワードする場合を考える。
接続元サーバからは中継サーバの yy 番ポートに接続すれば、転送先サーバのポート xx に転送されるようにする。

```
(転送先サーバのポートxx)
  |
publicゾーンに属するインタフェース
(中継サーバ)
internalゾーンに属するインタフェース
  |
(接続元サーバ)
```

```
# firewall-cmd --zone=internal --add-forward-port=port=yy:proto=tcp:toport=xx:toaddr=x.x.x.x --parmanent
# firewall-cmd --zone=internal --add-masquerade --permanent
# firewall-cmd --zone=public --add-masquerade --permanent
# firewall-cmd --reload
```
