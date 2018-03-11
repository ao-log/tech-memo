
## /etc/named.conf

##### リゾルバ設定

Listen するアドレスを適切に記載すること。リゾルバ機能を 192.168.0/24 に提供する場合は、自身のアドレス 192.168.0.x をここに書き、そのアドレスで Listen するようにすること。

```
options{
    listen-on port 53 { 127.0.0.1; x.x.x.x; };
}
```

DNS 問い合わせの転送先

```
options{
    forwarders{ 他のリゾルバのアドレス; };
```

##### オープンリゾルバとして機能しなくする

攻撃者に利用されることを防ぐため。オープンリゾルバになっていると、送信元を細工したクエリを投げることで攻撃対象に大量にパケットを送って負荷をかけたり、悪用される可能性がある。

下記設定は [■設定ガイド：オープンリゾルバー機能を停止するには【BIND編】](https://jprs.jp/tech/notice/2013-04-18-fixing-bind-openresolver.html) より引用

```
        acl "TRUSTSRC" {            // TRUSTSRCというacl作成の設定を追加します
          192.0.2.0/24;             // クエリ元のネットワークを記載します
          2001:db8:1::/48;          // クエリ元のネットワークを記載します
          localhost;                // 標準で入る設定を残します
          localnets;                // 標準で入る設定を残します
        };

        options {
          ...
          recursion yes;                   // リゾルバーとして動作します
          allow-query { TRUSTSRC; };       // TRUSTSRCからのみクエリを許可します
          allow-recursion { TRUSTSRC; };   // TRUSTSRCからのみリゾルバーとして動作します
          allow-query-cache { TRUSTSRC; }; // TRUSTSRCからのみキャッシュの内容を返します
          ...
        };
```

## オペレーション

##### named.conf のチェック

```
# named-checkconf
```

##### 設定の反映(シリアル番号が増えているゾーン情報のみ)

```
# rndc reload
```
