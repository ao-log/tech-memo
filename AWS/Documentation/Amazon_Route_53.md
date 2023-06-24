
# Document

## 概要

[Amazon Route 53 とは](https://docs.aws.amazon.com/ja_jp/Route53/latest/DeveloperGuide/Welcome.html)

主に次の３つの機能がある。

* ドメイン名の登録
* 名前解決
* リソースの正常性チェック


[ドメイン登録の仕組み](https://docs.aws.amazon.com/ja_jp/Route53/latest/DeveloperGuide/welcome-domain-registration.html)

次の流れで対応。

1. ドメインと同じ名前の hosted zone を作成
1. レジストラに送信(Amazon Registrar, Inc. or AWS のレジストラ関連会社 Gandi )
1. レジストリに送信
1. パブリック WHOIS データベースに登録


[ウェブサイトやウェブアプリケーションへのインターネットトラフィックのルーティング](https://docs.aws.amazon.com/ja_jp/Route53/latest/DeveloperGuide/welcome-dns-service.html)

名前解決の流れについて書かれている。


[Amazon Route 53 がリソースの正常性をチェックする方法](https://docs.aws.amazon.com/ja_jp/Route53/latest/DeveloperGuide/welcome-health-checks.html)

以下の値を設定

* プロトコル
* リクエスト間隔
* 失敗と判定する連続回数のしきい値
* 異常がある場合の通知方法


[Amazon Route 53 の概念](https://docs.aws.amazon.com/ja_jp/Route53/latest/DeveloperGuide/route-53-concepts.html)

ドメイン登録

* ドメイン名
* ドメインレジストラ
* ドメインレジストリ
* ドメインリセラー

DNS

* 最上位ドメイン (TLD)
* 権威ネームサーバー
* 再利用可能な委任セット: 4 つの一連の権威ネームサーバー。再利用可能な委任セットを作成し、新しいホストゾーンに関連付けることが可能。
* DNS リゾルバー
* ホストゾーン
* ルーティングポリシー
* サブドメイン
* TTL

ヘルスチェック

* DNS フェイルオーバー


## 開始方法

[開始方法](https://docs.aws.amazon.com/ja_jp/Route53/latest/DeveloperGuide/getting-started.html)

Route 53 のドメインを登録し、S3 の静的 Web サイトホスティングに対してレコードを設定する流れ。


## ドメイン名の登録

[新しいドメインの登録](https://docs.aws.amazon.com/ja_jp/Route53/latest/DeveloperGuide/domain-register.html)

* 登録するドメイン名の入力
* ドメインの登録者、管理者、技術担当者の連絡先情報
* Privacy protection: WHOIS データベースに対して連絡先情報を非表示にするかどうか
* 有効期限日前にドメイン登録を自動更新するかどうか


[別のレジストラへの許可のない移管を防ぐためのドメインのロック](https://docs.aws.amazon.com/ja_jp/Route53/latest/DeveloperGuide/domain-lock.html)

多くのレジストリでドメインをロックし、他者がドメインを別のレジストラに移管することを防止するように設定することが可能。


[ドメインの移管](https://docs.aws.amazon.com/ja_jp/Route53/latest/DeveloperGuide/domain-transfer.html)

別のレジストラから Route 53、Route 53 から別のレジストラ、もしくは AWS アカウント間でドメイン登録を移管可能。


## DNS サービスとして Amazon Route 53 を設定

[Route 53 を使用中のドメインの DNS サービスにする](https://docs.aws.amazon.com/ja_jp/Route53/latest/DeveloperGuide/migrate-dns-domain-in-use.html)

別のサービスから Route 53 に移行するときの流れについて書かれている。


[ルーティングポリシーの選択](https://docs.aws.amazon.com/ja_jp/Route53/latest/DeveloperGuide/routing-policy.html)

クエリに応答する方法を設定できる。

* シンプルルーティングポリシー
* フェイルオーバールーティングポリシー
* 位置情報ルーティングポリシー
* 複数値回答ルーティングポリシー
* 加重ルーティングポリシー

など。


[エイリアスレコードと非エイリアスレコードの選択](https://docs.aws.amazon.com/ja_jp/Route53/latest/DeveloperGuide/resource-record-sets-choosing-alias-non-alias.html)

エイリアスレコードを作成することで Zone Apex にも CNAME のような別名リソースに転送するレコードを作成可能。


[Amazon Route 53 での DNSSEC 署名の設定](https://docs.aws.amazon.com/ja_jp/Route53/latest/DeveloperGuide/dns-configuring-dnssec.html)

* DNSSEC により DNS リゾルバーは DNS 応答が Route 53 から送信され改ざんされていないことを検証できる
* キー署名キー(KSK)とゾーン署名キー(ZSK) の 2 種類のキーがある
  * 署名対象が公開鍵(DNSKEY レコード)の鍵は KSK
  * 署名対象がゾーンに含まれる残りのレコードとなる鍵は ZSK
* KSK 作成には KMS のカスタマー管理キーが us-east-1 リージョンに必要


## Route 53 Resolver

[VPC とネットワークの間における DNS クエリの解決](https://docs.aws.amazon.com/ja_jp/Route53/latest/DeveloperGuide/resolver.html)

Route 53 Resolver。

* インバウンドエンドポイント: VPC 外から VPC 内の名前解決をできるようにする。
* アウトバウンドエンドポイント: VPC 内から VPC 外の名前解決をできるようにする。


## Amazon Route 53 ヘルスチェックの作成と DNS フェイルオーバーの設定

[Amazon Route 53 ヘルスチェックの種類](https://docs.aws.amazon.com/ja_jp/Route53/latest/DeveloperGuide/health-checks-types.html)

以下からヘルスチェック方法を選ぶことができる。

* エンドポイントをモニタリングするヘルスチェック
* 他のヘルスチェック (算出したヘルスチェック) を監視するヘルスチェック
* CloudWatch アラームをモニタリングするヘルスチェック


[Amazon Route 53 がヘルスチェックの正常性を判断する方法](https://docs.aws.amazon.com/ja_jp/Route53/latest/DeveloperGuide/dns-failover-determining-health-of-endpoints.html)

* エンドポイントの場合、18 %を超えるヘルスチェッカーがエンドポイントを正常であるとレポートした場合、Route 53 はそのエンドポイントを正常と見なす。


[DNS フェイルオーバーの設定](https://docs.aws.amazon.com/ja_jp/Route53/latest/DeveloperGuide/dns-failover-configuring.html)

次のドキュメントが分かりやすいかもしれない。アクティブ/パッシブ、アクティブ/アクティブ、組み合わせの 3 パターンについて設定例を案内している。

* [DNS フェイルオーバーで Route 53 のヘルスチェックを使用する方法を教えてください。](https://aws.amazon.com/jp/premiumsupport/knowledge-center/route-53-dns-health-checks/)


## Route 53 Resolver DNS Firewall

[Route 53 Resolver DNS Firewall](https://docs.aws.amazon.com/ja_jp/Route53/latest/DeveloperGuide/resolver-dns-firewall.html)

* VPC からのアウトバウンドトラフィックをフィルタリング、規制できる
* ホワイトリスト、ブロックリストの両指定が可能
* Firewall Manager により Organizations のアカウント全体の設定、管理を一元的に行うことができる
* DNS Firewall は Route 53 Resolver を通過するアウトバウンドの通信が対象。一方で Network Firewall はネットワーク層、アプリケーション層のトラフィックに対してフィルタリング



# BlackBelt

[Amazon Route 53 Resolver](https://pages.awscloud.com/rs/112-TZM-766/images/20191016_AWS_Blackbelt_Route53_Resolver.pdf)

* DNS
  * ネームサーバ: 権威サーバ。階層構造になっている
  * フルサービスリゾルバ: ルートから順に問い合わせて名前解決する。TTL が失効するまでキャッシュを保持
  * スタブリゾルバ: 一般には OS の DNS クライアント
  * フォワーダー


[Amazon Route 53 Hosted Zone](https://pages.awscloud.com/rs/112-TZM-766/images/20191105_AWS_Blackbelt_Route53_Hosted_Zone_A.pdf)

* ドメイン名の基本
  * P11: 名前解決の流れ
  * P13: レジストラ、レジストリ
* ネームサーバの基本
  * P21: 権限移譲元を「親ゾーン」、権限移譲先を「子ゾーン」と呼ぶ。
  * P22: 権限移譲。親ゾーンで子ゾーンの設定を行う。
    * NS レコードでネームサーバの FQDN を指定
    * A レコードでネームサーバの FQDN に対応する IP アドレスを指定
  * P27: NS レコードは親ゾーンに子ゾーンの設定を行う。また、子ゾーンにおいても自ゾーンの NS レコードを定義する。
  * P30: Zone Apex には CNAME を定義できない(CNAME を設定すると同一の名前で別のリソースレコードを設定できないため。Zone Apex には SOA, NS レコードが必要なため)。
  * P33: ネガティブキャッシュ(存在しない RRSet を問い合わせ)の TTL 値は SOA レコードで設定。
* Amazon Route 53 Hosted Zone  


# 参考

* Document
  * [Amazon Route 53 とは](https://docs.aws.amazon.com/ja_jp/Route53/latest/DeveloperGuide/Welcome.html)
  * [API Reference](https://docs.aws.amazon.com/ja_jp/Route53/latest/APIReference/Welcome.html)
* サービス紹介ページ
  * [Amazon Route 53](https://aws.amazon.com/jp/route53/)
  * [よくある質問](https://aws.amazon.com/jp/route53/faqs/)
* [ナレッジセンター](https://aws.amazon.com/jp/premiumsupport/knowledge-center/#Amazon_Route_53)
* Black Belt
  * [Amazon Route 53 Hosted Zone](https://pages.awscloud.com/rs/112-TZM-766/images/20191105_AWS_Blackbelt_Route53_Hosted_Zone_A.pdf)
  * [Amazon Route 53 Resolver](https://pages.awscloud.com/rs/112-TZM-766/images/20191016_AWS_Blackbelt_Route53_Resolver.pdf)


