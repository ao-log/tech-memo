
(2018年2月25日 作成記事)
https://qiita.com/ao_log/items/383fd9b4c039e670ceb4

Apache で 安全性の高い SSL/TLSサーバを構築する際に確認するポイントをまとめました。セキュリティ面に配慮しながらの構築が求められる領域となります。

# 参照すべき情報

IPA などの機関が出しているガイドラインへは最低限のこととして目を通します。『SSL/TLS暗号設定ガイドライン～安全なウェブサイトのために（暗号設定対策編）～』は熟読したほうがいいです。
https://www.ipa.go.jp/security/vuln/ssl_crypt_config.html

# セキュリティランクの診断

こちらのサイトで診断できます。
https://www.ssllabs.com/ssltest/index.html

# TLS/SSL の有効/無効化

設定ファイルは ```/etc/httpd/conf.d/ssl.conf``` (ディストリビューションによって多少異なるかもしれません)になります。SSLv3 、SSLv2 は無効化します。次の記事によると、「遅くとも 2018年6月30日までには、全てのSSL/TLS 1.0を無効化する必要がある。」とあるので、TLSv1 も無効化しておくのがよいです。
http://www.intellilink.co.jp/article/pcidss/18.html

```
# 高セキュリティ型 (TLSv1.2 のみ有効)
SSLProtocol TLSv1.2
# TLS1.1, 1.2 のみ有効
SSLProtocol All -SSLv2 -SSLv3 -TLSv1
```

サーバ側での有効/無効化の状況は次のコマンドで確認できます。無効化している場合は、「handshake failure」となります。

```shell-session
// TLSv1.2
$ openssl s_client -connect www.example.com:443 -tls1_2 -brief
// TLSv1.1
$ openssl s_client -connect www.example.com:443 -tls1_1 -brief
// TLSv1
$ openssl s_client -connect www.example.com:443 -tls1 -brief
// SSL3
$ openssl s_client -connect www.example.com:443 -ssl3 -brief
```


# 暗号化方式の決定順序

暗号化方式決定の決定をサーバ側で行います。

```
SSLHonorCipherOrder on
```

# 暗号スイートの設定

暗号スイートは「鍵交換＿署名＿暗号化＿ハッシュ関数」の組によって構成されます。RC4 のような安全性の低いものは使わないように設定します。

具体的にどう設定するかは、[『SSL/TLS暗号設定ガイドライン～安全なウェブサイトのために（暗号設定対策編）～』](https://www.ipa.go.jp/files/000045645.pdf) の 「C.2.2. OpenSSL 系での暗号スイートの設定例 (86ページ)」に書かれています。

### Forward Secrecy の有効化について

設定によっては、Forward Secrecy が有効になりません。Forward Secrecy が有効になっていない場合はどういうリスクがあるか。過去の通信内容が攻撃者によって蓄積されていると、サーバの秘密鍵が漏洩した場合にセッション鍵が復号され、過去の通信内容も復号されます。「6.3.1 秘密鍵漏えい時の影響範囲を狭める手法の採用（Perfect Forward Secrecy の重要性） (36 ページ)」に詳しく書かれています。

では、どのように Forward Secrecy を有効化するかは、こちらに詳しく書かれています。暗号スイートに高速なECDHEスイートを可能な限り使用する設定となります。
https://rms-digicert.ne.jp/howto/basis/Forward_Secrecy_Apache_Ngix.html#apache_forward_secrecy

# HTTP Strict Transport Security（HSTS）の設定有効化

サーバからブラウザ側に HTTPS で接続するように指示する設定です。というのも HTTP から HTTPS にリダイレクトする構成の場合、攻撃者の用意したサイトへ誘導し中間車攻撃を行うこともできるようです。それを防ぐために、HSTS が有効な場合は、ブラウザ側で HTTP ではなく HTTPS 側で接続する動作をしてくれます。ただ、HSTS 有効化をサーバからブラウザに伝える初回、そして有効期限切れとなった場合には HTTP でのアクセスを許す点には注意が必要です。

HSTS の有効期限を一年に設定するには、次のように書きます。

```
Header always set Strict-Transport-Security "max-age=31536000; includeSubDomains"
```

なお、HTTP へのアクセスを強制的に HTTPS へリダイレクトさせるには、次のように書きます。

```
<VirtualHost *:80>
  ServerAlias *
  RewriteEngine On
  RewriteRule ^(.*)$ https://%{HTTP_HOST}$1 [redirect=301]
</VirtualHost>
```


# 鍵長

RSA の場合は 2048 ビット以上とします。サーバ証明書の申請段階で 2048 ビット以上が要件になっている場合があり、この場合は既に実現できていることになります。なお、2048 ビットの RSA 秘密鍵は次のコマンドで作成します。

```
$ openssl genrsa 2048 > private.key
```





# openssl のパッチが当たっているかどうかの確認

ディストリビュータによって、RPM で確認できるパッケージのバージョンのつけ方がまちまちです。よって、バージョンを頼りにパッチが当たっているかどうかを判断するのは危ういです。changelog を見ることを推奨します。

```shell-session
$ rpm -q --changelog openssl | less

* 水  5月 17 2017 Tomáš Mráz <tmraz@redhat.com> 1.0.2k-8
- fix regression in openssl req -x509 command (#1450015)

* 木  4月 13 2017 Tomáš Mráz <tmraz@redhat.com> 1.0.2k-7
- handle incorrect size gracefully in aes_p8_cbc_encrypt()

...

// CVE-2014-3566 = poodle です。
* 水 10月 15 2014 Tomáš Mráz <tmraz@redhat.com> 1.0.1e-39
- fix CVE-2014-3567 - memory leak when handling session tickets
- fix CVE-2014-3513 - memory leak in srtp support　
- add support for fallback SCSV to partially mitigate CVE-2014-3566
  (padding attack on SSL3)

...

// CVE-2014-0160 = Heartbleed です。
* 火  4月 08 2014 Tomáš Mráz <tmraz@redhat.com> 1.0.1e-34
- fix CVE-2014-0160 - information disclosure in TLS heartbeat extension

...

```

# サーバ証明書の有効期限

上述した IPA の資料に次の記述がありました。証明書の更新忘れを防ぐために有効期限一年にし、定例業務化するという工夫になります。

> サーバ管理者は、1 年程度の有効期間を持つサーバ証明書を選択し、サーバ証明書の更新作業を、年次の定型業務と位置付けることが望ましい。

また、万一秘密鍵が漏洩した場合でも、復号できるのは過去最長一年までです（Forward Secrecy を有効化していない場合は、通信の記録を採取されていると、通信の暗号化に使うセッション鍵も復号され、通信内容も復号されるので）。有効期間は長期間にせず、業務プロセスを考慮しつつほどほどの短さにするのがよさそうです。
