# この記事について

[Google Cloud Platform Japan Blog](https://cloudplatform-jp.googleblog.com/) では日々様々な情報が発信されています。新サービスの発表、事例紹介だけでなく、ポイントを分かりやすく説明した資料、網羅的で重厚な資料、Google の謎テクノロジーについての紹介など、役立つ情報が豊富！　ただ、情報量がとてつもなく多いので、分野ごと、項目ごとに整理して、後で目当ての記事をすぐ見つけられるようにしたいと思い、この記事を書きました。

今年は Google Cloud Platform 10 周年と節目の年。改めて記事を掘り起こしてみると、新たな発見があるかも知れません。

また、時代ごとの変遷を感じ取れる資料として、記事後半にはサービスが公開された時期を年表にしたもの、Google が公開している論文リンク集も作成しました。

# 各分野の記事

## サービス全体を俯瞰する記事

### Google Cloud Platform の概要がわかる記事

Google Cloud では、多くの方に各サービスの勘所をわかりやすく紹介するために、オンラインの番組を毎週木曜に放映しています。生で見るだけでなく、後で見返すこともできます。

* [Cloud OnAir | Google Cloud 生放送番組](https://cloudplatformonline.com/onair-japan.html)

放映内容は後日 [Google Cloud Platform Japan Blog](https://cloudplatform-jp.googleblog.com/) にもアップされます。特に全体の概要がわかる放映回は次のものです。

* [Cloud OnAir 番組レポート : ニーズに合わせたクラウドのベストな使い方とは](https://cloudplatform-jp.googleblog.com/2018/03/Google-Cloud-OnAir-05-Live.html)
* [Cloud OnAir 第 1 回開催報告：徹底解剖 GCP のここがすごい](https://cloudplatform-jp.googleblog.com/2017/10/cloud-onair-japan-01-gcp-overview.html)

また、数多くのサービスの概要を PDF 2 ページでまとめた資料もあります。

* [『 1 行でわかる Google Cloud Platform』公開](https://cloudplatform-jp.googleblog.com/2018/05/Google-Cloud-Platform-PaperPrint.html)

### 東京リージョンの開設

リージョンの開設も記事になることがあります。東京リージョンの開始は2016年11月8日です！

* [東京 GCP リージョン の正式運用を開始しました](https://cloudplatform-jp.googleblog.com/2016/11/tokyo-region-now-open.html)

### GCP の特徴、強み

他社クラウドと比較した違いや、GCP の特徴・強みについて書かれた記事です。強みとしては、ライブマイグレーションやプレウォーミングの必要のないロードバランサ、インスタンス起動時間、グローバルな SDN ベースの高速なネットワークなどについて書かれています。

* [AWS に精通したユーザー向けの GCP ガイドを拡充](https://cloudplatform-jp.googleblog.com/2016/08/aws-gcp.html)
* [差別化と技術的競争優位の源泉となる Google Cloud Platform](https://cloudplatform-jp.googleblog.com/2015/10/google-cloud-platform_1.html)
* [オンプレミスを利用する IT 管理者のための Google Cloud Platform ガイド](https://cloudplatform-jp.googleblog.com/2017/02/Google-Cloud-Platform-for-data-center-professionals-what-you-need-to-know.html)

### マルチクラウド

一方で、単一のクラウドだけでなく複数のクラウドを使う時代にもなってきました。その時代での対応方針について書かれたのが次の記事です。（2016年の記事なので少し古いかもしれませんが・・）

* [マルチ クラウド : どこでも同じように実行できることの意義](https://cloudplatform-jp.googleblog.com/2016/08/blog-post.html)

### オンプレミスからの移行

移行にあたってはプロセスに従って、抜け・漏れなく手堅く進めたいです。そのためのガイドラインも提供されています。

* [ホワイトペーパー : GCP へのリフト＆シフト](https://cloudplatform-jp.googleblog.com/2018/01/whitepaper-lift-and-shift-to-Google-Cloud-Platform.html)
* [GCP への移行を成功に導く 5 つのフェーズ](https://cloudplatform-jp.googleblog.com/2016/06/gcp-5.html)
* [ソリューション ガイド : VM 移行のベスト プラクティス](https://cloudplatform-jp.googleblog.com/2017/05/Solution-guide-best-practices-for-migrating-Virtual-Machines.html)
* [CloudEndure と共同で GCP への VM 移行サービスを無償提供](https://cloudplatform-jp.googleblog.com/2017/03/no-cost-VM-migration-to-Google-Cloud-Platform-now-available-with-CloudEndure.html)

### クラウドの利点を活かす

単に VM を移行するだけだとインフラの運用コストが残ったままです。マネージドサービスを活用することで運用コストを低減しつつ、開発に集中でき、スケーラビリティも手に入れることができます。その紹介がこちら。

* [Cloud OnAir 番組レポート : クラウド移行後の最適化方法を伝授](https://cloudplatform-jp.googleblog.com/2018/02/Google-Cloud-OnAir-02-Live.html)

### 実際に試してみる

技術は実際に手を動かして試してみないと、手にしっかり馴染まないです。まず、試してみるためのガイドが用意されています。

* [簡単スタートアップ ガイドのお知らせ](https://cloudplatform-jp.googleblog.com/2015/08/blog-post_27.html)

### GCPUG

ディベロッパーコミュニティの活動も盛んです。私もお邪魔させていただき、多くのことを学ばせていただきました。最新情報や便利なオープンソースソフトウェア、組織内でどう回しているかなど、生きた知見を学べる場です。主催者の方たちは大変だと思うのですが、本当に感謝です。

* [デベロッパー コミュニティ紹介シリーズ : Google Cloud Platform User Group](https://cloudplatform-jp.googleblog.com/2018/02/google-cloud-platform-user-group-GCPUG.html)
* [デベロッパー コミュニティ紹介シリーズ : Kubernetes Meetup Tokyo](https://cloudplatform-jp.googleblog.com/2018/04/Kubernetes-Meetup-Tokyo.html)
* [デベロッパー コミュニティ紹介シリーズ：bq_sushi](https://cloudplatform-jp.googleblog.com/2018/04/developer-community-bqsushi.html)
* [デベロッパー コミュニティ紹介シリーズ : TensorFlow User Group](https://cloudplatform-jp.googleblog.com/2017/12/Dev-Community-01-TensorFlow-User-Group-Tokyo.html)

#### Podcast

Podcast もあります。英語を聞く練習にもなります。。

* [Google Cloud Platform の情報を Podcast で公開](https://cloudplatform-jp.googleblog.com/2016/02/google-cloud-platform-podcast.html)
* [GCP Podcast 100 回記念 : 人気エピソード ベスト 10](https://cloudplatform-jp.googleblog.com/2017/11/GCP-podcast-hits-100-episodes-here-are-the-10-most-popular.html)

### GCP のアイコンとアーキテクチャ図

プレゼンや設計図の作成に利用できるアイコンとアーキテクチャ図のサンプルも提供されています。

* [GCP のアイコンとアーキテクチャ図サンプルを公開](https://cloudplatform-jp.googleblog.com/2016/12/Google-Cloud-Platform-icons-and-sample-architectural-diagrams-for-your-designing-pleasure.html)

### 事例

事例も紹介されています。いくつかピックアップさせていただきました。

* [Pokémon GO の爆発的ヒットを支える Google Cloud](https://cloudplatform-jp.googleblog.com/2016/10/pokemon-go-google-cloud.html)
* [事例 : ベアメタルから GCP への移行で Sentry が学んだこと](https://cloudplatform-jp.googleblog.com/2017/11/looking-back-on-our-migration-from-bare-metal-to-GCP-Sentry.html)
* [株式会社メルカリの導入事例: 先端技術を手軽に活用できる Google Cloud Platform はベストな選択肢](https://cloudplatform-jp.googleblog.com/2016/10/google-cloud-platform_28.html)

## セキュリティ

### Google のセキュリティを俯瞰する記事

セキュリティに関する Google の取り組みや GCP で利用できるサービスを網羅的に俯瞰するには Cloud OnAir が便利です。

* [Cloud OnAir 番組レポート :GCP で構築するセキュアなサービス、基本と最新プロダクトのご紹介](https://cloudplatform-jp.googleblog.com/2018/11/Cloud-OnAir-20181101-gcp.html)
* [Cloud OnAir 番組レポート : Google Cloud Platform のセキュリティ、全てお話します](https://cloudplatform-jp.googleblog.com/2018/02/Google-Cloud-OnAir-03-Live.html)

### セキュリティはクラウドをためらう理由ではなく、導入する理由に

アジリティと経費削減だけではなく、今やセキュリティがクラウド選定の理由となる時代です。

* [MIT SMR 調査 : セキュリティはクラウドを躊躇する理由ではなく、導入する理由に](https://cloudplatform-jp.googleblog.com/2017/11/turns-out-security-drives-cloud-adoption-not-the-other-way-around.html)

### クラウドの安全性、Google が取得している認証

エンタープライズでの選定に耐えうるには、第三者機関による認証が欠かせません。主要な認証を取得していると書かれているのが次の記事です。

* [Google が提供するクラウドの安全性について](https://cloudplatform-jp.googleblog.com/2014/09/google.html)
* [Google Cloud Platform の新しいセキュリティと個人情報保護の ISO 認証を取得](https://cloudplatform-jp.googleblog.com/2016/05/google-cloud-platform-iso.html)

### Google 内部のセキュリティ設計

Google 内部でどのようなセキュリティ設計を行なっているかが書かれた記事です。階層的、網羅的に厳重な対策が行われています。

* [ホワイト ペーパー : Google インフラストラクチャのセキュリティ設計](https://cloudplatform-jp.googleblog.com/2017/01/how-we-secure-our-infrastructure.html)
* [Google データセンターにみるセキュリティとデータ保護のベスト プラクティス](https://cloudplatform-jp.googleblog.com/2016/03/google_31.html)

また、カスタムチップの Titan により、ハードウェアベースの “信頼の基点”（Root of Trust）が確立されています。その紹介記事がこちら。

* [GCP のセキュリティ : “信頼の基点” としての Titan](https://cloudplatform-jp.googleblog.com/2017/09/Titan-in-depth-security-in-plaintext.html)

### GCP 利用におけるセキュリティのガイドライン

GCP を利用してサービスを構築する際、自組織内のポリシーをどのように設計するか、またどのように GCP のサービスを組み合わせてセキュリティを確立するかのガイドラインがこちらです。

* [効果的なクラウド ガバナンスに向けて ―― 組織の要件に応じた GCP ポリシーの設計](https://cloudplatform-jp.googleblog.com/2018/02/designing-policies-for-GCP-customers-large-and-small.html)

### Cloud IAM

IAM により権限の管理ができます。全体的なことについて書かれた記事がこちら。

* [IAM ベスト プラクティス ガイドを公開](https://cloudplatform-jp.googleblog.com/2016/04/iam.html)
* [フローで学ぶ Cloud IAM](https://cloudplatform-jp.googleblog.com/2018/03/getting-to-know-Cloud-IAM.html)
* [AWS ユーザー向けの Google Cloud IAM ガイドを公開](https://cloudplatform-jp.googleblog.com/2017/04/Google-Cloud-IAM-for-AWS-users.html)

より限定的な内容についての、プラクティスが書かれた記事がこちらです。

* [Google Cloud のサービス アカウント キーを安全に管理する](https://cloudplatform-jp.googleblog.com/2017/08/help-keep-your-Google-Cloud-service-account-keys-safe.html)
* [組織を GCP のリソース階層にマッピングする](https://cloudplatform-jp.googleblog.com/2017/06/mapping-your-organization-with-the-Google-Cloud-Platform--resource-hierarchy.html)
* [すべての Google Cloud リソースを Cloud Resource Manager で一元管理](https://cloudplatform-jp.googleblog.com/2017/02/centrally-manage-all-your-Google-Cloud-resources-with-Cloud-Resource-Manager.html)

GCP では日々新機能がリリースされていますが、Compute Engine のより細かなリソースであったり日時でコントロールできる機能がリリースされています。

* [Cloud IAM の新機能で Compute Engine リソースをきめ細かく管理](https://cloudplatform-jp.googleblog.com/2018/11/get-more-control-over-your-compute-engine-resources-with-new-cloud-iam-features.html)

### Cloud Identity-Aware Proxy

App Engine などの HTTPS でアクセスするアプリケーションに対して、アクセス制御ポリシーを組み込めるサービスです。ウェブサービスの前段にプロキシを配置して、そこにアクセス制御のコントロールを一任させるものです。

* [アプリケーション アクセスを簡単かつ安全に管理する Cloud Identity-Aware Proxy](https://cloudplatform-jp.googleblog.com/2017/09/Cloud-Identity-Aware-Proxy-a-more-secure-way-to-move-internal-apps-to-GCP.html)
* [Cloud Identity-Aware Proxy におけるアクセス制御の仕組み](https://cloudplatform-jp.googleblog.com/2017/05/Getting-started-with-Cloud-Identity-Aware-Proxy.html)
* [Cloud Identity のアイデンティティ管理機能が GCP で利用可能に](https://cloudplatform-jp.googleblog.com/2017/07/enterprise-identity-made-easy-in-GCP-with-Cloud-Identity.html)
* [Cloud Identity-Aware Proxy : クラウド アプリケーションのアクセス保護](https://cloudplatform-jp.googleblog.com/2017/04/Cloud-Identity-Aware-Proxy-protect-application-access-on-the-cloud.html)

### Cloud Endpoints

API 呼び出しの認証、認可を行うサービスです。

* [ハイブリッド クラウドのセキュリティ : Cloud Endpoints による API 呼び出しの認証](https://cloudplatform-jp.googleblog.com/2017/10/more-secure-hybrid-cloud-deployments-with-Google-Cloud-Endpoints.html)

### Cloud Armor

Cloud Armor は DDoS 攻撃への防御や IP アドレスベースのフィルタリングを設定できるサービスです。SQL インジェクション対策もアルファリリースで提供されているので、将来的にはアプリケーションレイヤも防護できるようになる？

* [Cloud Armor : インターネットに接続されたサービスを DDoS 攻撃から防御](https://cloudplatform-jp.googleblog.com/2018/05/getting-to-know-Cloud-Armor-defense-at-scale-for-internet-facing-services.html)

### Cloud KMS

暗号化、復号処理を行う鍵の管理を行うサービスです。

* [Cloud KMS を使い暗号化キーの管理をクラウドで](https://cloudplatform-jp.googleblog.com/2017/01/managing-encryption-keys-in-the-cloud-introducing-Google-Cloud-Key-Management-Service.html)

### Cloud Audit Logging

監査ログを残すサービスです。

* [Cloud Audit Logging 活用のベスト プラクティス](https://cloudplatform-jp.googleblog.com/2018/04/best-practices-for-working-with-Google-Cloud-Audit-Logging.html)
* [Cloud Audit Logging が多くの GCP サービスで利用可能に](https://cloudplatform-jp.googleblog.com/2017/01/Google-Cloud-Audit-Logging-now-available-across-the-GCP-stack.html)

### Cloud Security Command Center

自分の GCP プロジェクトに対してセキュリティの診断をし、可視化してくれるサービスのようです。

* [Cloud Security Command Center で GCP 環境をモニタリング](https://cloudplatform-jp.googleblog.com/2018/05/monitor-your-GCP-environment-with-Cloud-Security-Command-Center.html)

### Cloud DLP API

リアルタイムに機密データの検出、分類、編集をしてくれるサービスのようです。定義済みの検出器があるほか、自分でカスタム検出器を定義することもできるようです。

* [Cloud DLP API による機密データの厳重な管理](https://cloudplatform-jp.googleblog.com/2018/04/take-charge-of-your-sensitive-data-with-the-Cloud-DLP-API.html)

### 技術的な対策の知見

クラウドの利用者としては知らなくてもいい知識も多いですが、勉強になる記事です。

* [PCIe のファジング : GPU 導入に向けたセキュリティ検証と対策](https://cloudplatform-jp.googleblog.com/2017/03/fuzzing-PCI-Express-security-in-plaintext.html)
* [GCP のセキュリティ : クラッシュの自動分析による脆弱性の検出](https://cloudplatform-jp.googleblog.com/2017/04/crash-exploitability-analysis-on-Google-Cloud-Platform-security-in-plaintext.html)
* [Google Cloud の KVM セキュリティを強化する 7 つの方法](https://cloudplatform-jp.googleblog.com/2017/02/7-ways-we-harden-our-KVM-hypervisor-at-Google-Cloud-security-in-plaintext.html)

## コンピューティング

### App Engine

#### 歴史

App Engine は最古から存在する GCP のサービスで誕生から 10 周年になります。エンタープライズ向けのサーバインフラからではなく、開発にフォーカスした PaaS からスタートしているのが GCP らしいですね。ブログでは時々歴史に関する記事もポストされます。

* [誕生から 10 年、App Engine の歩みを振り返って](https://cloudplatform-jp.googleblog.com/2018/06/reflecting-on-our-ten-year-App-Engine-journey.html)

#### 概要を学ぶための資料

* [Cloud OnAir 第 2 回開催報告：アプリランタイムについて学ぼう](https://cloudplatform-jp.googleblog.com/2017/10/cloud-onair-japan-02-gcp-app-runtime.html)

#### gVisor ベースの第二世代ランタイム

今年に入ってからのトピックとしては、gVisor ベースの第二世代ランタイムがサポートされるようになりました。

* [サーバーレス コンピューティングの実現に向けて](https://cloudplatform-jp.googleblog.com/2018/08/bringing-the-best-of-serverless-to-you.html)
* [Open-sourcing gVisor, a sandboxed container runtime](https://cloud.google.com/blog/products/gcp/open-sourcing-gvisor-a-sandboxed-container-runtime)

#### ランタイム

各言語のランタイムは新しいバージョンに対応していっています。

* [App Engine 向けの Go 1.11 ランタイムをベータ リリース](https://cloudplatform-jp.googleblog.com/2018/11/go-1-11-is-now-available-on-app-engine.html)
* [App Engine の Python サポートを強化](https://cloudplatform-jp.googleblog.com/2017/07/enhancing-the-Python-experience-on-App-Engine.html)
* [Java 8 サポートの App Engine スタンダード環境をベータ リリース](https://cloudplatform-jp.googleblog.com/2017/07/Google-App-Engine-standard-now-supports-Java-8.html)
* [Now, you can deploy your Node.js app to App Engine standard environment](https://cloud.google.com/blog/products/gcp/now-you-can-deploy-your-node-js-app-to-app-engine-standard-environment)

#### セキュリティ面の機能強化

* [App Engine ファイアウォールを正式リリース](https://cloudplatform-jp.googleblog.com/2017/10/App-Engine-firewall-now-generally-available.html)
* [App Engine アプリ用 SSL 証明書のマネージド サービスを無料提供](https://cloudplatform-jp.googleblog.com/2017/10/introducing-managed-SSL-for-Google-App-Engine.html)

#### 事例

* [任天堂株式会社の導入事例：ビッグタイトル『Super Mario Run』のバックエンドを支えた Google App Engine](https://cloudplatform-jp.googleblog.com/2017/06/nintendo-super-mario-run-google-cloud.html)

### Compute Engine

#### 概要をおさえる資料

* [Cloud OnAir 番組レポート :Google Compute Engine に Deep Dive](https://cloudplatform-jp.googleblog.com/2018/07/cloud-onair-google-compute-engine-deep-dive-2018-07-19.html)

#### マネージドインスタンスグループ

VM のデプロイ時にカナリアテストを行うなど、デプロイ方法の制御方法が豊富に用意されています。

* [Compute Engine のインスタンス グループを自在に更新できる Managed Instance Group Updater](https://cloudplatform-jp.googleblog.com/2017/09/meet-Compute-Engines-new-managed-instance-group-updater.html)

#### 管理に便利な機能

稼働中の VM インスタンスからイメージが作成できたり、削除からの保護といった有用な機能について書かれた記事です。

* [Compute Engine インスタンスの管理が容易に](https://cloudplatform-jp.googleblog.com/2018/03/managing-your-Compute-Engine-instances-just-got-easier.html)
* [Kubernetes と Container Engine でより良いノード管理を](https://cloudplatform-jp.googleblog.com/2017/04/toward-better-node-management-with-Kubernetes-and-Google-Container-Engine.html)

ディスクの高速データコピーには、スナップショットが便利という記事です。

* [ビッグデータを低コストで超高速転送](https://cloudplatform-jp.googleblog.com/2015/05/blog-post.html)

#### 起動高速化

プロファイリング方法や起動時間短縮方法について書かれた記事です。

* [Compute Engine インスタンスの起動を高速化する 3 つのステップ](https://cloudplatform-jp.googleblog.com/2017/08/three-steps-to-Compute-Engine-startup-time-bliss-Google-Cloud-Performance-Atlas.html)

#### ライブマイグレーション

謎の超技術の一つ。ゲスト VM に影響を与えず行われるマイグレーションについての記事です。

* [ライブ マイグレーションを使い、Google Compute Engine にダウンタイムのないサービス インフラストラクチャーを](https://cloudplatform-jp.googleblog.com/2015/03/google-compute-engine.html)

#### 大規模利用

巨大なサイズの VM インスタンスは複数ノードに分散させるよりは一台に全て乗せたいケースに有用です。Intel Distribution for Python の scikit-learn を使用したときのパフォーマンス向上についても性能比のグラフが載せられています。

* [96 vCPU の Compute Engine マシンタイプを正式リリース](https://cloudplatform-jp.googleblog.com/2018/03/96-vCPU-Compute-Engine-instances-are-now-generally-available.html)

22 万コアて。。すごいです（小並感）。

* [MIT の数学教授が 22 万コアの Compute Engine クラスタを構築](https://cloudplatform-jp.googleblog.com/2017/05/220000-cores-and-counting-MIT-math-professor-breaks-record-for-largest-ever-Compute-Engine-job.html)


#### Market Place

様々なオープンソースがプレインストールされた VM イメージが提供されています。なお、Cloud Launcher は現在は Market Place に改名されています。

* [Cloud Launcher で新たに 21 のオープンソース ソリューションが利用可能に](https://cloudplatform-jp.googleblog.com/2017/11/21-new-open-source-solutions-available-from-Google-Cloud-Launcher.html)
* [Cloud Launcher アップデート : 本番環境レベルのソリューションや簡単に使えるサービスが続々登場](https://cloudplatform-jp.googleblog.com/2017/04/stay-up-to-speed-with-Google-Cloud-Launcher-more-production-grade-solutions-same-easy-to-use-service.html)

### Kubernetes Engine

#### 歴史

Google が大きく関わっているプロダクトということもあり、歴史的な記事もいくつか掲載されています。

* [Google Container Engine が GA に](https://cloudplatform-jp.googleblog.com/2015/08/google-container-engine-ga99.html)
* [Google から世界へ : Kubernetes 誕生の物語](https://cloudplatform-jp.googleblog.com/2016/08/google-kubernetes.html)
* [オープンソースから始まり、持続可能な成功に至るまで ―― Kubernetes の “卒業” に寄せて](https://cloudplatform-jp.googleblog.com/2018/03/from-open-source-to-sustainable-success-the-Kubernetes-graduation-story.html)

#### 概要をおさえる資料

Cloud OnAir をはじめとして概要を知るための記事がこちら。

* [Cloud OnAir 番組レポート：Deep Dive Google Kubernetes Engine](https://cloudplatform-jp.googleblog.com/2018/08/Cloud-OnAir-20180802-deep-dive-google-kubernetes.html)
* [知りたいけど、今さら聞けない Kubernetes](https://cloudplatform-jp.googleblog.com/2015/02/gcp-kubernetes.html)
* [コンテナ クラスターを構成するものとは？](https://cloudplatform-jp.googleblog.com/2015/02/gcp.html)

#### ソリューションガイド

マイクロサービスから、継続的デリバリのパイプラインまで、ソリューション設計に関するガイドです。

* [ホワイトペーパー : モノリスからマイクロサービスへの移行](https://cloudplatform-jp.googleblog.com/2018/02/whitepaper-embark-on-journey-from-monoliths-to-microservices.html)
* [ソリューション ガイド : 本番運用に向けた Container Engine 環境の準備](https://cloudplatform-jp.googleblog.com/2017/06/Solutions-guide-preparing-Container-Engine-environments-for-production.html)
* [ソリューション ガイド : Spinnaker、Container Engine、Container Builder で信頼性の高いデプロイ体制を構築](https://cloudplatform-jp.googleblog.com/2017/10/building-reliable-deployments-with-Spinnaker-Container-Engine-and-Container-Builder.html)

#### コンテナのセキュリティ

上記のガイドライン以外にもセキュリティに関する記事も公開されています。

* [コンテナのセキュリティ（第 1 回）: 3 つの主要領域](https://cloudplatform-jp.googleblog.com/2018/05/exploring-container-security-an-overview.html)
* [コンテナのセキュリティ（第 2 回）: ノード イメージとコンテナ イメージ](https://cloudplatform-jp.googleblog.com/2018/05/exploring-container-security-Node-and-container-operating-systems.html)
* [コンテナのセキュリティ（第 3 回）: Grafeas で管理するコンテナ イメージのメタデータ](https://cloudplatform-jp.googleblog.com/2018/06/exploring-container-security-digging-into-Grafeas-container-image-metadata.html)
* [Kubernetes Engine 1.8 でのコンテナ セキュリティの強化](https://cloudplatform-jp.googleblog.com/2017/12/precious-cargo-securing-containers-with-Kubernetes-Engine-18.html)

#### Container Registry

Container Registry の脆弱性スキャン機能。イメージに脆弱性があるかどうかをスキャンし、結果をレポートしてくれるサービスです。

* [Container Registry の脆弱性スキャンで安全な CI/CD パイプラインを](https://cloudplatform-jp.googleblog.com/2018/09/guard-against-security-vulnerabilities-with-container-registry-vulnerability-scanning.html)

#### Spinneker

継続的デリバリと継続的デプロイの違いは、本番環境へのデプロイに人手を介するかどうかです。Spinneker は人による承認プロセスを挟むこともできます。継続的デリバリを実現するプロダクトです。

* [ネットフリックスの Spinnaker がGoogle Cloud Platform で利用可能に](https://cloudplatform-jp.googleblog.com/2015/11/spinnaker-google-cloud-platform.html)
* [Spinnaker 1.0 : マルチクラウド対応の継続的デリバリ プラットフォーム](https://cloudplatform-jp.googleblog.com/2017/06/spinnaker-10-continuous-delivery.html)
* [Kayenta : Google と Netflix が共同開発したオープンソースの自動カナリア分析ツール](https://cloudplatform-jp.googleblog.com/2018/06/introducing-Kayenta-an-open-automated-canary-analysis-tool-from-Google-and-Netflix.html)

#### GitLab、GitHub からのデプロイ

GitLab から GKE にデプロイする方法が紹介されている記事です。

* [GitLab から Kubernetes Engine へのデプロイが数クリックで可能に](https://cloudplatform-jp.googleblog.com/2018/05/now-you-can-deploy-to-Kubernetes-Engine-from-GitLab-with-a-few-clicks.html)

こちらの記事は [Google Cloud Platform Japan Blog](https://cloudplatform-jp.googleblog.com/) の記事ではないのですが、GitHub からも連携できるようになっています。

* [GitHubとGoogle Cloud Build の連携を試してみた！](https://medium.com/google-cloud-jp/try-github-cloudbuild-integration-5149175105fb)

#### Cloud Services Platform

今年の 7 月に発表された Managed Istio、GKE On-Prem、Knative などに関する記事です。

* [Cloud Services Platform 登場 : サービスの高度な運用基盤をクラウドとオンプレミスの両方で提供](https://cloudplatform-jp.googleblog.com/2018/07/cloud-services-platform-bringing-the-best-of-the-cloud-to-you.html)

#### マイクロサービスの開発例

API の公開にはマイクロサービスがぴったりです。この記事では、Kubernetes を使って構築した例が紹介されています。

* [マイクロサービスでスケーラブルな API を開発](https://cloudplatform-jp.googleblog.com/2016/07/api_27.html)

#### 事例

* [株式会社メルカリの導入事例：Kubernetes を駆使したマイクロサービス化でグローバルサービスの開発効率を劇的に向上](https://cloudplatform-jp.googleblog.com/2018/01/Google-Cloud-Platform-Mercari-kubernetes.html)
* [ゲスト投稿 : ML API、Cloud Pub/Sub、Cloud Functions でサーバーレス のデジタル アーカイブを構築した Incentro](https://cloudplatform-jp.googleblog.com/2018/02/how-we-built-a-serverless-digital-archive-with-machine-learning-APIs-Cloud-Pub-Sub-and-Cloud-Functions.html)
* [ゲスト投稿 : Spinnaker のマネージド パイプライン テンプレートと IaC で、本番システムに秩序をもたらした Waze](https://cloudplatform-jp.googleblog.com/2017/09/how-Waze-tames-production-chaos-using-Spinnaker-managed-pipeline-templates-and-infrastructure-as-code.html)

## サーバレス

サーバレスと一言で言っても Function as a Service にとどまる話ではありません。Cloud OnAir の放送では Cloud Function を中心に据えつつも、幅広い観点から紹介しています。

* [Cloud OnAir 番組レポート : 最新版 GCP ではじめる、サーバーレスアプリケーションの開発](https://cloudplatform-jp.googleblog.com/2018/11/20181108-Cloud-OnAir-serverless-application.html)

こちらは Google のサーバレスに対する目指す姿、方向性が書かれた記事です。

* [サーバーレス コンピューティングの実現に向けて](https://cloudplatform-jp.googleblog.com/2018/08/bringing-the-best-of-serverless-to-you.html)

## Deployment Manager

インフラをコードで定義し、自動的に構築できるサービスです。

* [Deployment Manager の GA リリースでデプロイが容易になります](https://cloudplatform-jp.googleblog.com/2015/08/deployment-manager-ga.html)
* [Cloud Deployment Manager の基本的な使い方](https://cloudplatform-jp.googleblog.com/2016/11/cloud-deployment-manager.html)
* [Cloud Deployment Manager でプロジェクトを自動作成](https://cloudplatform-jp.googleblog.com/2017/04/automating-project-creation-with-Google-Cloud-Deployment-Manager.html)
* [オープンソースへの取り組み : Google と HashiCorp のエンジニアが語る GCP インフラストラクチャ管理](https://cloudplatform-jp.googleblog.com/2017/03/partnering-on-open-source-Google-and-HashiCorp-engineers-on-managing-GCP-infrastructure.html)
* [ゲスト投稿 : GCP インフラストラクチャをコードとして管理する Terraform](https://cloudplatform-jp.googleblog.com/2017/05/guest-post-using-Terraform-to-manage-Google-Cloud-Platform-infrastructure-as-code.html)

## データ

Google では、ビッグをつけずに単にデータというらしいです。

### データ分析基盤

全体像については Cloud OnAir が分かりやすいです。

* [Cloud OnAir 番組レポート : GCP で構築するデータ分析基盤の最新情報をご紹介](https://cloudplatform-jp.googleblog.com/2018/11/Cloud-OnAir-20181115-GCP.html)
* [Cloud OnAir 番組レポート : クラウドを活用したデータ分析基盤の第一歩](https://cloudplatform-jp.googleblog.com/2018/04/Cloud-OnAir-report-0412.html)
* [Cloud OnAir 番組レポート : クラウドを活用したリアルタイムデータ分析](https://cloudplatform-jp.googleblog.com/2018/05/Cloud-OnAir-2018-05-10.html)
* [Cloud OnAir 番組レポート：第 3 回 No-ops で大量データ処理基盤を簡単に構築する](https://cloudplatform-jp.googleblog.com/2017/11/cloud-onair-japan-03-gcp-no-ops.html)
* [基本から学ぶ ビッグデータ / データ分析 / 機械学習 サービス群)](https://www.slideshare.net/GoogleCloudPlatformJP/ss-82505511)

### データベースセキュリティ

* [Google Cloud データベース セキュリティのベスト プラクティス](https://cloudplatform-jp.googleblog.com/2018/06/best-practices-for-securing-your-Google-Cloud-databases.html)

### Cloud Storage

ストレージクラスについての記事。長期アーカイブデータ用の Coldline でも瞬時にデータにアクセスできるのは驚きです。

* [Cloud Storage の新ストレージ クラスを発表 : 4 つのクラスで多様なワークロードに対応](https://cloudplatform-jp.googleblog.com/2016/10/cloud-storage-4.html)

パフォーマンス向上に関する Tips 集です。

* [Cloud Storage のパフォーマンスを最適化する](https://cloudplatform-jp.googleblog.com/2018/03/optimizing-your-Cloud-Storage-performance-Google-Cloud-Performance-Atlas.html)

データ保護のため、権限のチェック、機密データのチェック、対策のための行動、そしてこれらの行動を日常的に回すことについて、書かれた記事です。

* [Cloud Storage バケットのデータ保護に役立つ 4 つのステップ](https://cloudplatform-jp.googleblog.com/2017/10/4-steps-for-hardening-your-Cloud-Storage-buckets-taking-charge-of-your-security.html)

Pub/Sub と連携できます。

* [Cloud Storage のオブジェクトイベントが Cloud Pub/Sub サブスクリプションの対象に](https://cloudplatform-jp.googleblog.com/2017/04/Cloud-Storage-introduces-Cloud-Pub-Sub-notifications.html)

### Cloud SQL

今は第 2 世代が提供されています。競合サービスとの性能比較のデータも載っています。

* [Cloud SQL 第 2 世代のパフォーマンスと機能を詳説](https://cloudplatform-jp.googleblog.com/2016/08/cloud-sql-2.html)
* [モバイル ゲーム開発向けの Cloud SQL 第 2 世代ソリューション ガイドを公開](https://cloudplatform-jp.googleblog.com/2016/10/cloud-sql-2.html)

こちらは 2014 年の記事です。

* [Google Cloud SQL がどなたでも利用できるようになりました ー SLA、500 GB のデータベース、暗号化の新機能とともに](https://cloudplatform-jp.googleblog.com/2014/02/google-cloud-sql-sla500-gb.html)

PostgreSQL も正式リリースを迎えています。

* [Cloud SQL for PostgreSQL を正式リリース](https://cloudplatform-jp.googleblog.com/2018/07/Cloud-SQL-for-PostgreSQL-now-generally-available-and-ready-for-your-production-workloads.html)

### Spanner

#### サービスの概要

+ [Cloud Natural Language API の新機能と Cloud Spanner の正式リリース](https://cloudplatform-jp.googleblog.com/2017/05/Cloud-Natural-Language-API-enters-beta.html)
* [NoSQL から新しい SQL へ : グローバルなミッションクリティカル DB へと進化を遂げた Cloud Spanner](https://cloudplatform-jp.googleblog.com/2017/07/from-NoSQL-to-New-SQL-how-Spanner-became-a-global-mission-critical-database.html)
* [Cloud Spanner : ミッションクリティカルな業務に対応するグローバルなデータベース サービス](https://cloudplatform-jp.googleblog.com/2017/02/introducing-Cloud-Spanner-a-global-database-service-for-mission-critical-applications.html)

#### 理論面の解説

従来の常識を覆すサービスということもあり、理論面の解説記事もあります。

* [Cloud Spanner と CAP 定理](https://cloudplatform-jp.googleblog.com/2017/02/inside-Cloud-Spanner-and-the-CAP-Theorem.html)
* [Cloud Spanner が外部整合性を提供する理由](https://cloudplatform-jp.googleblog.com/2018/03/why-you-should-pick-strong-consistency-whenever-possible.html)

#### 事例

* [Cloud Spanner を使ってメール パーソナライズ システムを再構築したリクルートテクノロジーズ](https://cloudplatform-jp.googleblog.com/2018/04/how-we-used-Cloud-Spanner-to-build-our-email-personalization-system-from-Soup-to-nuts.html)
* [国際自動車株式会社の導入事例：何千ものタクシーと乗客を、スマホを振るだけでマッチング。リアルタイム性が求められるビッグデータ処理を Cloud Spanner で一挙可能に](https://cloudplatform-jp.googleblog.com/2018/03/google-cloud-platform-kokusaijidosya-cloudspanner.html)
* [『ドラゴンボール レジェンズ』の舞台裏を支える Google Cloud](https://cloudplatform-jp.googleblog.com/2018/06/Behind-the-scenes-with-the-Dragon-Ball-Legends-GCP-backend.html)

### Bigtable

#### 概要をおさえる資料

* [Cloud OnAir 番組レポート : Cloud Bigtable に迫る、基本機能も含めユースケースまで丸ごと紹介](https://cloudplatform-jp.googleblog.com/2018/09/Cloud-OnAir-cloud-bigtable-20180830.html)

#### サービスの公開アナウンス

Google 社内での歴史からユースケースまで書かれています。

* [Cloud Bigtable が GA リリース、ペタバイト スケールの NoSQL ワークロードに対応](https://cloudplatform-jp.googleblog.com/2016/08/cloud-bigtable-ga-nosql.html)

#### 事例

* [Fastly の導入事例 : 履歴統計 DB を MySQL から Cloud Bigtable にダウンタイムなしで移行](https://cloudplatform-jp.googleblog.com/2017/07/how-we-moved-our-historical-stats-from-mysql-bigtable-zero-downtime.html)

### Datastore

* [毎月 15 兆クエリを処理し、それ以上にも対応できる Cloud Datastore](https://cloudplatform-jp.googleblog.com/2016/08/15-cloud-datastore.html)

### BigQuery

#### 概要をおさえる資料

* [Cloud OnAir 番組レポート : BigQuery の仕組みからベストプラクティスまでのご紹介](https://cloudplatform-jp.googleblog.com/2018/09/Cloud-OnAir-BigQuery-20180906.html)
* [Cloud OnAir 番組レポート : Cloud Bigtable に迫る、基本機能も含めユースケースまで丸ごと紹介](https://cloudplatform-jp.googleblog.com/2018/09/Cloud-OnAir-cloud-bigtable-20180830.html)
* [“フルマネージド” の新たな基準を確立する BigQuery](https://cloudplatform-jp.googleblog.com/2016/09/bigquery.html)

#### 便利な機能や使い方

* [UDF を新たにサポートした Google BigQuery で、より “ディープ” なクラウド アナリティクスを！](https://cloudplatform-jp.googleblog.com/2015/08/udf-google-bigquery.html)
* [BigQuery Data Transfer Service を正式リリース](https://cloudplatform-jp.googleblog.com/2017/11/announcing-bigquery-data-transfer-service-general-availability.html)
* [BigQuery への課金データのエクスポートで GCP 利用コストを管理](https://cloudplatform-jp.googleblog.com/2017/12/monitor-and-manage-your-costs-with.html)

#### 理論面の解説

* [BigQuery によるインメモリ クエリの実行](https://cloudplatform-jp.googleblog.com/2016/08/bigquery.html)

#### 事例

* [ソニーネットワークコミュニケーションズ株式会社の導入事例動画： BigQuery の導入で煩雑な業務から解放、運用にパラダイムシフトを](https://cloudplatform-jp.googleblog.com/2017/09/Google-Cloud-Platform-BigQuery-sonynetwork.html)

### Data Studio

#### BigQuery のデータをダッシュボードで可視化

* [Data Studio と BigQuery を使った BI ダッシュボードの作り方](https://cloudplatform-jp.googleblog.com/2017/06/how-to-build-a-bi-dashboard-using-google-data-studio-and-bigquery.html)

### Dataflow

* [Cloud Dataflow パイプラインのモニタリングと高速化](https://cloudplatform-jp.googleblog.com/2016/06/cloud-dataflow_20.html)
* [Cloud Dataflow の新機能 : パイプラインの詳細を表示する](https://cloudplatform-jp.googleblog.com/2016/06/cloud-dataflow_28.html)

### Dataproc

* [Spark / Hadoop マネージド サービスの Google Cloud Dataproc が一般リリース](https://cloudplatform-jp.googleblog.com/2016/03/spark-hadoop-google-cloud-dataproc.html)

### Pub/Sub

* [Google Cloud Dataflow と Cloud Pub/Sub が正式リリース](https://cloudplatform-jp.googleblog.com/2015/08/google-cloud-dataflow-cloud-pubsub.html)
* [Google Cloud Pub/Sub が gRPC のサポートで大幅に高速化](https://cloudplatform-jp.googleblog.com/2016/03/google-cloud-pubsub-grpc.html)

## ネットワーク

#### 概要

[Cloud OnAir 番組レポート :Google Networking Deep Dive](https://cloudplatform-jp.googleblog.com/2018/08/Cloud-OnAir-20180809-Google-Networking-Deep-Dive.html)

#### Google が擁するインフラ

[LA - 香港間を 120 Tbps で結ぶ海底ケーブル システムを構築へ](https://cloudplatform-jp.googleblog.com/2016/10/la-120-tbps.html)

[アジアで Google へのアクセスがより速く : FASTER - 台湾間の新海底ケーブルが開通](https://cloudplatform-jp.googleblog.com/2016/09/google-faster.html)

[新リージョンと海底ケーブルの増設でグローバル インフラストラクチャを拡張](https://cloudplatform-jp.googleblog.com/2018/01/expanding-our-global-infrastructure-new-regions-and-subsea-cables.html)

[GCP のイントラゾーン レイテンシを 40 % 低減する Andromeda 2.1](https://cloudplatform-jp.googleblog.com/2017/11/Andromeda-2-1-reduces-GCPs-intra-zone-latency-by-40-percent.html)

[データセンター ネットワークに対する Google の取り組み - 最新世代の Jupiter を紹介](https://cloudplatform-jp.googleblog.com/2015/06/google-jupiter.html)

#### 設計ガイド

[一から学べる Google Cloud Networking リファレンス ガイド](https://cloudplatform-jp.googleblog.com/2016/10/google-cloud-networking_4.html)

[内部負荷分散によるスケーラブルなプライベート サービスの構築](https://cloudplatform-jp.googleblog.com/2016/12/blog-post.html)

[Google Cloud CDN が CDN Interconnect に加わり、ユーザーに選択肢を提供](https://cloudplatform-jp.googleblog.com/2017/03/Google-Cloud-CDN-joins-CDN-Interconnect-providers-delivering-choice-to-users.html)

#### VPC

[Cloud VPC のファイアウォール管理をサービス アカウントで簡素化](https://cloudplatform-jp.googleblog.com/2018/01/simplify-Cloud-VPC-firewall-management-with-service-accounts.html)

[ファイアウォール ルールを堅固にする 3 つの方法](https://cloudplatform-jp.googleblog.com/2018/01/three-ways-to-configure-robust-firewall-rules.html)

[VPC の VM インスタンスで複数のネットワーク インターフェースをサポート](https://cloudplatform-jp.googleblog.com/2017/10/with-multiple-network-interfaces-connect-third-party-devices-to-GCP-workloads.html)

[Shared VPC : 複数プロジェクトにまたがる仮想ネットワークを一元管理](https://cloudplatform-jp.googleblog.com/2017/06/getting-started-with-shared-VPC.html)

[Google VPC を特徴づける 4 つのキーワード](https://cloudplatform-jp.googleblog.com/2017/07/reimagining-virtual-private-clouds.html)

#### Cloud NAT

[Cloud NAT : ソフトウェア定義型の新しいネットワーク アドレス変換サービス](https://cloudplatform-jp.googleblog.com/2018/10/cloud-nat-deep-dive-into-our-new-network-address-translation-service.html)

#### Dedicated Interconnect

[Dedicated Interconnect : Google Cloud に高速でプライベート接続](https://cloudplatform-jp.googleblog.com/2017/10/announcing-dedicated-interconnect-your-fast-private-on-ramp-to-Google-Cloud.html)

[グローバル ルーティングや新ロケーションをサポートした Dedicated Interconnect を正式リリース](https://cloudplatform-jp.googleblog.com/2017/11/Google-Cloud-Dedicated-Interconnect-gets-global-routing-more-locations-and-is-GA.html)


[Network Service Tiers をアルファ リリース : クラウド ネットワークの選択が可能に](https://cloudplatform-jp.googleblog.com/2017/09/introducing-Network-Service-Tiers-your-cloud-network-your-way.html)


[GCP プライベート ネットワークの柔軟性を高める拡張可能なサブネットワーク](https://cloudplatform-jp.googleblog.com/2016/09/gcp.html)


[GCP を支えるロードバランサの設計を公開](https://cloudplatform-jp.googleblog.com/2016/03/gcp.html)



[HTTP/2 で Google Cloud Platform もより早く](https://cloudplatform-jp.googleblog.com/2015/10/http2-google-cloud-platform.html)

[CDN Interconnect でアカマイとの連携が実現する、企業の接続性とパフォーマンスの向上](https://cloudplatform-jp.googleblog.com/2015/11/cdn-interconnect.html)

[うるう秒への対応も万全 : 新しい パブリック NTP サーバーを公開](https://cloudplatform-jp.googleblog.com/2016/12/ntp.html)



## 機械学習

### 機械学習プロジェクトの進め方

要件定義からモデルの構築、運用までのライフサイクル全体に渡って、全体的な視点を提供してくれる資料です。

* [Cloud OnAir 番組レポート : 機械学習のプロジェクトの進め方](https://cloudplatform-jp.googleblog.com/2018/05/Cloud-OnAir-2018-05-24.html)

こちらは「機械学習とは何か、ディープラーニング、AI との違いは何か」のとっかかりの部分から理解するのに適した資料です。

* [Cloud OnAir 番組レポート：第 4 回 今話題の機械学習・GCPで何ができるのか？](https://cloudplatform-jp.googleblog.com/2017/11/cloud-onair-japan-04-ml-gcp.html)

他にも様々なガイドが提供されています。

* [機械学習用データの収集と準備](https://cloudplatform-jp.googleblog.com/2018/08/preparing-and-curating-your-data-for-machine-learning.html)
* [CIO 向けのデータ アナリティクス / 機械学習ガイドを公開](https://cloudplatform-jp.googleblog.com/2017/08/CIOs-guide-to-data-analytics-and-machine-learning.html)

### AutoML

#### 概要をおさえる資料

* [Cloud OnAir 番組レポート : AutoML Vision で学ぶ実践的機械学習](https://cloudplatform-jp.googleblog.com/2018/12/Cloud-OnAir-20181129-AutoML-Vision.html)

#### 技術記事

* [ラーメン二郎とブランド品で AutoML Vision の認識性能を試す](https://cloudplatform-jp.googleblog.com/2018/03/automl-vision-in-action-from-ramen-to-branded-goods.html)

#### 事例

* [株式会社LIFULL の導入事例：物件画像のカテゴリー分類を AutoML で自動化。数十秒かかっていた分類を自動化で 2 秒](https://cloudplatform-jp.googleblog.com/2018/07/AutoML-LIFULL.html)

### TensorFlow

GCP のサービスではないですが、Google の主要ソフトウェアということで・・

#### 技術記事

* [TensorFlow Playground でわかるニューラルネットワーク](https://cloudplatform-jp.googleblog.com/2016/07/tensorflow-playground.html)
* [夏休みの自由工作：TensorFlowでじゃんけんマシンを作る](https://cloudplatform-jp.googleblog.com/2017/10/my-summer-project-a-rock-paper-scissors-machine-built-on-tensorflow.html)

#### 事例

* [NTT ドコモ「AIタクシー」を支える TensorFlow と需要予測モデル](https://cloudplatform-jp.googleblog.com/2018/04/ntt-docomo-ai-taxi-tensorflow.html)
* [キュウリ農家とディープラーニングをつなぐ TensorFlow](https://cloudplatform-jp.googleblog.com/2016/08/tensorflow_5.html)

### TPU

専用のプロセッサを用いることで、特定用途の計算を低消費電力で、高速に解くことができます。コストに対する性能比も高くなります。なお、[Colaboratory](https://colab.research.google.com/notebooks/welcome.ipynb#recent=true) でも無料で TPU を使うことができます。

技術面を掘り下げた記事が多いです。

* [機械学習アプリを大幅に高速化する TPU カスタム チップを開発](https://cloudplatform-jp.googleblog.com/2016/05/tpu.html)
* [機械学習用チップの性能評価 : TPU の研究論文を公開](https://cloudplatform-jp.googleblog.com/2017/04/quantifying-the-performance-of-the-TPU-our-first-machine-learning-chip.html)
* [深層学習に特化したプロセッサ、Cloud TPU の設計](https://cloudplatform-jp.googleblog.com/2018/08)
* [Google の Tensor Processing Unit (TPU) で機械学習が 30 倍速くなるメカニズム](https://cloudplatform-jp.googleblog.com/2017/05/an-in-depth-look-at-googles-first-tensor-processing-unit-tpu.html)

Publickey の記事ですが、TPU 3.0 は液冷のようです。冷却効率が高く、かつ消費電力を抑えることができるので、電気代削減だけでなく、環境負荷にも考慮されたソリューションだと言えます。

* [Publickey: Google、機械学習専用の第三世代プロセッサ「TPU 3.0」を発表。Google初の液冷システム採用。Google I/O 2018](https://www.publickey1.jp/blog/18/googletpu_30googlegoogle_io_2018.html)

また、Edge TPU はエッジ用の TPU です。エッジ側では使える電力や設置スペースが限られるなどの制約がありますが、それらの制約を克服しつつ高速な推論処理を行うことができます。

* [Cloud OnAir 番組レポート : クラウドからエッジまで！進化する GCP の IoT サービス](https://cloudplatform-jp.googleblog.com/2018/11/Cloud-OnAir-20181122-IoT.html)

### GPU

機械学習や HPC をはじめとして、GPU のコンピュータクラスタは需要が高まっています。GPU も利用可能です。

* [Compute Engine と Cloud Machine Learning で GPU が利用可能に](https://cloudplatform-jp.googleblog.com/2017/03/GPUs-are-now-available-for-Google-Compute-Engine-and-Cloud-Machine-Learning.html)
* [クラウド GPU が 2017 年から利用可能に](https://cloudplatform-jp.googleblog.com/2016/11/gpu-2017.html)
* [通常の半額で使用できるプリエンプティブル GPU](https://cloudplatform-jp.googleblog.com/2018/01/introducing-preemptible-gpus-50-off.html)


### Dialogflow

* [Dialogflow Enterprise Edition で音声とテキストによる新しい会話アプリの構築方法を提供](https://cloudplatform-jp.googleblog.com/2017/11/introducing-Dialogflow-Enterprise-Edition-a-new-way-to-build-voice-and-text-conversational-apps.html)

#### 事例

* [株式会社ポケモンの導入事例：Dialogflow の活用で簡単かつ短期間に音声対応アプリを開発。多言語対応でグローバル展開が容易な GCP はポケモンに最適](https://cloudplatform-jp.googleblog.com/2018/08/the-pokemon-company-dialogflow-google-cloud.html)

### 学習済みモデルの API

特定用途のモデルを作成するには自前で訓練する必要がありますが、学習済みモデルで事足りることもあります。

* [Cloud Vision API による不適切コンテンツのフィルタリング](https://cloudplatform-jp.googleblog.com/2016/09/cloud-vision-api.html)
* [画像認識の常識を変える Google Cloud Vision API](https://cloudplatform-jp.googleblog.com/2015/12/google-cloud-vision-api.html)
* [新しい Cloud Speech-to-Text で通話や動画音声のテキスト変換精度が向上](https://cloudplatform-jp.googleblog.com/2018/05/toward-better-phone-call-and-video-transcription-with-new-Cloud-Speech-to-Text.html)
* [Cloud Natural Language API で非構造化テキストを構造化する](https://cloudplatform-jp.googleblog.com/2016/08/cloud-natural-language-api_29.html)
* [Cloud Video Intelligence API 登場と Cloud ML の最新アップデート](https://cloudplatform-jp.googleblog.com/2017/03/announcing-google-cloud-video-intelligence-api-and-more-cloud-machine-learning-updates.html)

### 事例

* [キユーピー株式会社の導入事例動画：キユーピー + ブレインパッド + Google の取り組みで次世代の AI 検査装置を実現](https://cloudplatform-jp.googleblog.com/2017/09/google-gcp-ai-kewpie.html)

### Datalab

* [Cloud Datalab 新ベータ : ローカルでも実行でき、TensorFlow も利用可能に](https://cloudplatform-jp.googleblog.com/2016/08/cloud-datalab-tensorflow.html)

### IoT

* [Cloud IoT Core を正式リリースしました](https://cloudplatform-jp.googleblog.com/2018/02/the-thing-is-Cloud-IoT-Core-is-now-generally-available27.html)
* [Cloud OnAir 番組レポート : クラウドからエッジまで！進化する GCP の IoT サービス](https://cloudplatform-jp.googleblog.com/2018/11/Cloud-OnAir-20181122-IoT.html)

## Develop

# Develop

### Cloud Shell

* [Cloud Shell が GA リリース、料金は無料に](https://cloudplatform-jp.googleblog.com/2016/08/cloud-shell-ga.html)

### Cloud Source Repositories

* [Cloud Source Repositories が正式リリース、5 ユーザーで 50 GB まで無料](https://cloudplatform-jp.googleblog.com/2017/07/Cloud-Source-Repositories-now-GA-and-free-for-up-to-five-users-and-50GB-of-storage.html)
* [Cloud Source Repositories と Container Builder によるサーバーレスでの自動デプロイ](https://cloudplatform-jp.googleblog.com/2018/04/automatic-serverless-deployments-with-Cloud-Source-Repositories-and-Container-Builder.html)

### API Design

* [Google における API のバージョニング](https://cloudplatform-jp.googleblog.com/2017/07/versioning-APIs-at-Google.html)
* [API デザイン : URL には名前と識別子のどちらを使うべきか](https://cloudplatform-jp.googleblog.com/2017/11/API-design-choosing-between-names-and-identifiers-in-URLs.html)

### gRPC

* [gRPC : オープンソースの RPC フレームワークがバージョン 1.0 に](https://cloudplatform-jp.googleblog.com/2016/09/grpc-rpc-10.html)
* [gRPC を使用した効率的なモバイル アプリの構築](https://cloudplatform-jp.googleblog.com/2015/07/grpc.html)
* [gRPC API を Cloud Endpoints で管理する](https://cloudplatform-jp.googleblog.com/2017/05/manage-your-gRPC-APIs-with-Google-Cloud-Endpoints.html)
* [gRPC の評価 : 複数言語のサポート具合をチェックする](https://cloudplatform-jp.googleblog.com/2017/05/putting-gRPC-multi-language-support-to-the-test.html)

### Stackdriver

* [Google Stackdriver が GA リリースに](https://cloudplatform-jp.googleblog.com/2016/10/google-stackdriver-ga.html)
* [Google Cloud Monitoring: Stackdriver で Google Cloud Platform をモニタリング](https://cloudplatform-jp.googleblog.com/2015/01/gcp-google-cloud-monitoring-stackdriver.html)
* [Stackdriver Debugger が GA リリース : 本番環境でのデバッグに最適](https://cloudplatform-jp.googleblog.com/2016/10/stackdriver-debugger-ga.html)
* [App Engine 向けの Stackdriver Trace が GA リリース](https://cloudplatform-jp.googleblog.com/2016/05/app-engine-stackdriver-trace-ga.html)
* [Stackdriver Logging でフィルタリングすべきログ メッセージの見分け方](https://cloudplatform-jp.googleblog.com/2017/09/preventing-log-waste-with-Stackdriver-Logging.html)
* [DevOps で役に立つ Stackdriver の 6 つのポイント](https://cloudplatform-jp.googleblog.com/2016/06/devops-stackdriver-6.html)

## SRE、CRE

SRE = Site Reliability Engineering（サイト信頼性エンジニアリング）、CRE = Customer Reliability Engineering（顧客信頼性エンジニアリング）です。実は、[Google Cloud Platform Japan Blog](https://cloudplatform-jp.googleblog.com/) にも多くの記事が投稿されています。また、書籍『[SRE サイトリライアビリティエンジニアリング
――Googleの信頼性を支えるエンジニアリングチーム](https://www.oreilly.co.jp/books/9784873117911/)」という分厚い本も出版されており、エラーバジェットをはじめとした様々な概念、そして信頼性を高めたりコントロールするためのプラクティスを学ぶことができます。

* [SRE への冒険の始まり : Google Mission Control にようこそ](https://cloudplatform-jp.googleblog.com/2016/07/sre-google-mission-control.html)
* [Google の新しい専門職 : CRE が必要な理由](https://cloudplatform-jp.googleblog.com/2016/10/google-cre.html)
* [SRE の教訓 : Google におけるインシデント管理とは](https://cloudplatform-jp.googleblog.com/2017/03/Incident-management-at-Google-adventures-in-SRE-land.html)
* [SLO、SLI、SLA について考える : CRE が現場で学んだこと](https://cloudplatform-jp.googleblog.com/2017/02/availability-part-deux-CRE-life-lessons.html)
* [優れた SLO を策定するには : CRE が現場で学んだこと](https://cloudplatform-jp.googleblog.com/2017/11/building-good-SLOs-CRE-life-lessons.html)
* [SLO のエスカレーション ポリシー : CRE が現場で学んだこと](https://cloudplatform-jp.googleblog.com/2018/02/an-example-escalation-policy-CRE-life-lessons.html)
* [エスカレーション ポリシーの適用 : CRE が現場で学んだこと](https://cloudplatform-jp.googleblog.com/2018/02/applying-the-escalation-policy-CRE-life-lessons.html)
* [事後分析を外部と共有することの意義 : CRE が現場で学んだこと](https://cloudplatform-jp.googleblog.com/2017/12/fearless-shared-postmortems-CRE-life-lessons.html)
* [SLO 違反への対処 : CRE が現場で学んだこと](https://cloudplatform-jp.googleblog.com/2018/01/consequences-of-SLO-violations-CRE-life-lessons.html)
* [信頼性の高いリリースとロールバック : CRE が現場で学んだこと](https://cloudplatform-jp.googleblog.com/2017/04/reliable-releases-and-rollbacks-CRE-life-lessons.html)
* [エラー バジェットの使い過ぎが意味するもの : CRE が現場で学んだこと](https://cloudplatform-jp.googleblog.com/2018/07/understanding-error-budget-overspend-cre-life-lessons.html)
* [エラー バジェットの使い過ぎを解消する : CRE が現場で学んだこと (後編)](https://cloudplatform-jp.googleblog.com/2018/07/cre-life-lessons-good-housekeeping-for-error-budgets.html)
* [クラウド時代のトラブルシューティング : 解決に役立つプロバイダーとのコミュニケーション（前編）](https://cloudplatform-jp.googleblog.com/2018/07/Troubleshooting-tips-How-to-talk-so-your-cloud-provider-will-listen-and-understand.html)
* [クラウド時代のトラブルシューティング : 解決に役立つプロバイダーとのコミュニケーション（後編）](https://cloudplatform-jp.googleblog.com/2018/07/Troubleshooting-tips-Help-your-cloud-provider-help-you.html)
* [SRE との “壁” を取り除くには : CRE が現場で学んだこと](https://cloudplatform-jp.googleblog.com/2017/07/making-the-most-of-an-SRE-service-takeover-CRE-life-lessons.html)
* [SRE のサポートを受けるべきアプリとは？ : CRE が現場で学んだこと](https://cloudplatform-jp.googleblog.com/2017/07/why-should-your-app-get-SRE-support-CRE-life-lessons.html)
* [SRE へのサポート移行で失敗しないために : CRE が現場で学んだこと](https://cloudplatform-jp.googleblog.com/2017/07/how-SREs-find-the-landmines-in-a-service-CRE-life-lessons.html)
* [カナリアのおかげで命拾い : CRE が現場で学んだこと](https://cloudplatform-jp.googleblog.com/2017/04/how-release-canaries-can-save-your-bacon-CRE-life-lessons.html)
* [ダーク ローンチとは何か : CRE が現場で学んだこと](https://cloudplatform-jp.googleblog.com/2017/08/CRE-life-lessons-what-is-a-dark-launch-and-what-does-it-do-for-me.html)
* [ダーク ローンチの実用性 : CRE が現場で学んだこと](https://cloudplatform-jp.googleblog.com/2017/08/cre-life-lessons-practicalities-of-dark-launches.html)
* [可用性とどう向き合うべきか、それが問題だ : CRE が現場で学んだこと](https://cloudplatform-jp.googleblog.com/2017/02/available-or-not-that-is-the-question-CRE-life-lessons.html)
* [ロード シェディングを利用して想定以上のトラフィックをさばく : CRE が現場で学んだこと](https://cloudplatform-jp.googleblog.com/2016/12/cre.html)
* [SRE の教訓 : App Engine で 1 日 1,000 億件以上のリクエストに対応する方法](https://cloudplatform-jp.googleblog.com/2016/04/sre-app-engine-1-1000.html)
* [自業自得の DDoS 攻撃から身を守るには : CRE が現場で学んだこと](https://cloudplatform-jp.googleblog.com/2016/11/ddos-cre.html)


## その他

### データセンター

* [IT インフラの標準化に向け、Open Compute Project に参加](https://cloudplatform-jp.googleblog.com/2016/03/it-open-compute-project.html)
* [データセンター向けの新しいディスクを探求](https://cloudplatform-jp.googleblog.com/2016/02/blog-post_27.html)
* [Google の Data Center で 360度ツアー](https://cloudplatform-jp.googleblog.com/2016/03/google-date-center-360.html)

### HPC

* [ゲスト投稿 : 創薬用の大規模仮想スクリーニングに GCP を活用するSilicon Therapeutics](https://cloudplatform-jp.googleblog.com/2017/08/guest-post-using-GCP-for-massive-drug-discovery-virtual-screening.html)

# リリース時期

GCP の各サービスのリリース時期を並べた表です。

| サービス名 | リリースレベル | 日付 | 記事 |
|---|---|---|---|
| Google App Engine | - | 2008/04/07 | [Introducing Google App Engine + our new blog](https://googleappengine.blogspot.com/2008/04/introducing-google-app-engine-our-new.html) |
| Cloud Storage | preview | 2010/05/19 | [Google Storage for Developers: A Preview](http://googlecode.blogspot.com/2010/05/google-storage-for-developers-preview.html) |
| BigQuery | preview | 2010/05/19 | [BigQuery and Prediction API: Get more from your data with Google](http://googlecode.blogspot.com/2010/05/bigquery-and-prediction-api-get-more.html) |
| | GA | 2011/11/14 | [Google BigQuery Service: Big data analytics at Google speed](https://cloudplatform.googleblog.com/2011/11/google-bigquery-service-big-data.html) |
| Cloud SQL | preview | 2011/10/06 | [Google Cloud SQL: your database in the cloud](http://googlecode.blogspot.com/2011/10/google-cloud-sql-your-database-in-cloud.html) |
| | GA | 2014/02/11 | [Google Cloud SQL now Generally Available with an SLA, 500GB databases, and encryption](https://cloudplatform.googleblog.com/2014/02/google-cloud-sql-now-generally-available.html) |
| | 第2世代 | 2016/08/17 | [Cloud SQL Second Generation performance and feature deep dive](https://cloud.google.com/blog/products/gcp/cloud-sql-second-generation-performance-and-feature-deep-dive) |
| Google Compute Engine | preview | 2012/06/28 | [Google Compute Engine launches, expanding Google’s cloud offerings](https://cloudplatform.googleblog.com/2012/06/google-compute-engine-launches.html) |
| | GA | 2013/12/2 | [Google Compute Engine is now Generally Available with expanded OS support, transparent maintenance, and lower prices](https://cloudplatform.googleblog.com/2013/12/google-compute-engine-is-now-generally-available.html) |
| Cloud Datastore | beta | 2013/05/15 | [Ushering in the next generation of computing at Google I/O](https://cloudplatform.googleblog.com/2013/05/ushering-in-next-generation-of.html) |
| | GA | 2016/08/16 | [Google Cloud Datastore serves over 15 trillion queries per month and is ready for more](https://cloud.google.com/blog/products/gcp/google-cloud-datastore-serves-over-15) |
| Google Container Engine(現: Google Kubernetes Engine) | alpha | 2014/11/04 | [Unleashing Containers and Kubernetes with Google Container Engine](https://cloudplatform.googleblog.com/2014/11/unleashing-containers-and-kubernetes-with-google-compute-engine.html) |
| | GA | 2015/08/26 | [Google Container Engine is Generally Available](https://cloudplatform.googleblog.com/2015/08/Google-Container-Engine-is-Generally-Available.html) |
| Cloud DNS | GA | 2015/04/13 | [Google’s network edge: presence, connectivity and choice for today’s enterprise](https://cloudplatform.googleblog.com/2015/04/Googles-network-edge-presence-connectivity-and-choice-for-todays-enterprise.html) |
| Cloud BigTable | beta | 2015/05/06 | [Announcing Google Cloud Bigtable: The same database that powers Google Search, Gmail and Analytics is now available on Google Cloud Platform](https://cloudplatform.googleblog.com/2015/05/introducing-Google-Cloud-Bigtable.html) |
| | GA | 2016/08/18 | [Google Cloud Bigtable is generally available for petabyte-scale NoSQL workloads](https://cloud.google.com/blog/products/gcp/google-cloud-bigtable-is-generally-available-for-petabyte-scale-nosql-workloads) |
| Cloud VPN  | GA | 2015/05/20 | - |
| Container Registry | GA | 2015/06/22 | [Container Engine & Container Registry Updates - New Features & Pricing](https://cloudplatform.googleblog.com/2015/06/Container-Engine-Container-Registry-Updates-New-Features-Pricing.html) |
| Deployment Manager | GA | 2015/07/22 | - |
| Cloud Dataflow | GA | 2015/08/12 | [Announcing General Availability of Google Cloud Dataflow and Cloud Pub/Sub](https://cloudplatform.googleblog.com/2015/08/Announcing-General-Availability-of-Google-Cloud-Dataflow-and-Cloud-Pub-Sub.html) |
| Cloud Pub/Sub | GA | 2015/08/12 | 同上 |
| HTTP(S) Load Balancing | GA | 2015/10/29 | [Bringing you more flexibility and better Cloud Networking performance, GA of HTTPS Load Balancing and Akamai joins CDN Interconnect](https://cloudplatform.googleblog.com/2015/11/bringing-you-more-flexibility-and-better-Cloud-Networking-performance-GA-of-HTTPS-Load-Balancing-and-Akamai-joins-CDN-Interconnect.html) |
| Cloud Dataproc | GA | 2016/02/23 | [Google Cloud Dataproc managed Spark and Hadoop service now GA](https://cloud.google.com/blog/products/gcp/google-cloud-dataproc-managed-spark-and-hadoop-service-now-ga) |
| Cloud IAM | beta | 2016/03/24 | [How Google is bringing cloud computing innovation to the enterprise](https://cloud.google.com/blog/products/gcp/how-google-is-bringing-cloud-computing-innovation-to-the-enterprise) |
| Cloud Shell | GA | 2016/08/03 | [Cloud Shell now GA, and still free](https://cloud.google.com/blog/products/gcp/cloud-shell-now-ga-and-still-free) |
| Stackdriver Logging | GA | 2016/10/20 | [Google Stackdriver is now generally available for hybrid cloud monitoring, logging and diagnostics](https://cloud.google.com/blog/products/gcp/google-stackdriver-generally-available) |
| Cloud Natural Language API | GA | 2016/11/15 | [Google Cloud Machine Learning family grows with new API, editions and pricing](https://cloud.google.com/blog/products/gcp/cloud-machine-learning-family-grows-with-new-api-editions-and-pricing) |
| Internal Load Balancing | GA | 2016/12/08 | [Building scalable private services with Internal Load Balancing](https://cloud.google.com/blog/products/gcp/building-scalable-private-services-with-internal-load-balancing) |
| Cloud Spanner | beta | 2017/02/14 | [Introducing Cloud Spanner: a global database service for mission-critical applications](https://cloud.google.com/blog/products/gcp/introducing-cloud-spanner-a-global-database-service-for-mission-critical-applications) |
| | GA | 2017/05/16 | [Cloud Spanner is now production-ready; let the migrations begin!](https://cloud.google.com/blog/products/gcp/cloud-spanner-is-now-production-ready-let-the-migrations-begin) |
| Cloud Container Builder(現: Cloud Build) | GA | 2017/03/06 | [Google Cloud Container Builder: a fast and flexible way to package your software](https://cloud.google.com/blog/products/gcp/google-cloud-container-builder-a-fast-and-flexible-way-to-package-your-software) |
| Cloud Functions | beta | 2017/03/09 | [Google Cloud Functions: a serverless environment to build and connect cloud services](https://cloud.google.com/blog/products/gcp/google-cloud-functions-a-serverless-environment-to-build-and-connect-cloud-services_13) |
| | GA | 2018/07/19 | [Bringing the best of serverless to you](https://cloudplatform.googleblog.com/2018/07/bringing-the-best-of-serverless-to-you.html) |
| Cloud Datalab | GA | 2017/03/10 | [100 announcements (!) from Google Cloud Next '17](https://www.blog.google/products/google-cloud/100-announcements-google-cloud-next-17/) |
| Cloud Machine Learning Engine | GA | 2017/03/08 | 同上 |
| Cloud KMS | GA | 2017/03/16 | [Cloud KMS GA, new partners expand encryption options](https://cloud.google.com/blog/products/gcp/cloud-kms-ga-new-partners-expand-encryption-options) |
| Translation API | GA | 2017/04/06 | - |
| Cloud Speech-to-Text API | GA | 2017/04/18 | [Cloud Speech API is now generally available](https://cloud.google.com/blog/products/gcp/cloud-speech-api-is-now-generally-available) |
| Cloud Vision API | GA | 2017/05/18 | - |
| Cloud Source Repositories | GA | 2017/05/25 | [Cloud Source Repositories: now GA and free for up to five users and 50GB of storage](https://cloud.google.com/blog/products/gcp/cloud-source-repositories-now-ga-and-free-for-up-to-five-users-and-50gb-of-storage) |
| Cloud Identity-Aware Proxy | GA | 2017/08/31 | [Cloud Identity-Aware Proxy: a simple and more secure way to manage application access](https://cloud.google.com/blog/products/gcp/cloud-identity-aware-proxy-a-more-secure-way-to-move-internal-apps-to-gcp) |
| Cloud Firestore | beta | 2017/10/03 | [Introducing Cloud Firestore: Our New Document Database for Apps](https://firebase.googleblog.com/2017/10/introducing-cloud-firestore.html)
| Cloud Interconnect | GA | 2017/10/23 | [Google Cloud Dedicated Interconnect gets global routing, more locations, and is GA](https://cloud.google.com/blog/products/gcp/google-cloud-dedicated-interconnect-gets-global-routing-more-locations-and-is-ga) |
| Video Intelligence API | GA | 2017/11/30 | - |
| Cloud AutoML | alpha | 2018/01/17 | [Cloud AutoML: Making AI accessible to every business](https://cloud.google.com/blog/topics/inside-google-cloud/cloud-automl-making-ai-accessible-every-business) |
| | beta | 2018/07/19 | - |
| Cloud TPU | beta | 2018/02/12 | [Cloud TPU machine learning accelerators now available in beta](https://cloud.google.com/blog/products/gcp/cloud-tpu-machine-learning-accelerators-now-available-in-beta) |
| | GA | 2018/06/27 | - |
|  | v2 GA<br>v3 alpha | 2018/08/31 | [What makes TPUs fine-tuned for deep learning?](https://cloud.google.com/blog/products/ai-machine-learning/what-makes-tpus-fine-tuned-for-deep-learning)
| | v3 beta | 2018/10/10 | - |
| Cloud IoT Core | GA | 2018/02/21 | [The thing is . . . Cloud IoT Core is now generally available](https://cloud.google.com/blog/products/gcp/the-thing-is-cloud-iot-core-is-now-generally-available) |
| Cloud Armor | beta | 2018/03/19 | [Introducing new ways to protect and control your GCP services and data](https://cloud.google.com/blog/products/gcp/introducing-new-ways-to-protect-and-control-your-gcp-services-and-data) |
| Cloud Data Loss Prevention (DLP) API | GA | 2018/03/21 | [Take charge of your sensitive data with the Cloud Data Loss Prevention (DLP) API](https://cloud.google.com/blog/products/gcp/take-charge-of-your-sensitive-data-with-the-cloud-dlp-api) |
| Dialogflow Enterprise Edition | GA | 2018/04/17 | [Dialogflow Enterprise Edition is now generally available](https://cloud.google.com/blog/products/gcp/dialogflow-enterprise-edition-is-now-generally-available) |
| Cloud Composer | GA | 2018/07/19 | [What a week! 105 announcements from Google Cloud Next '18](https://cloud.google.com/blog/topics/inside-google-cloud/what-week-105-announcements-google-cloud-next-18) |
| Managed Istio | alpha | 同上 | 同上 |
| Cloud Talent Solution | GA | 2018/08/03 | - |
| Cloud Text-to-Speech API | GA | 2018/08/28 | [Announcing updates to Cloud Speech-to-Text and the general availability of Cloud Text-to-Speech](https://cloud.google.com/blog/products/ai-machine-learning/announcing-updates-to-cloud-speech-to-text-and-general-availability-of-cloud-text-to-speech) |
| Cloud Memorystore for Redis | GA | 2018/09/19 | [Announcing general availability of Cloud Memorystore for Redis](https://cloud.google.com/blog/products/databases/announcing-general-availability-of-cloud-memorystore-for-redis) |
| Data Studio | GA | 2018/09/21 | [Unlock insights with ease: Data Studio and Cloud Dataprep are now generally available](https://cloud.google.com/blog/products/data-analytics/unlock-insights-with-ease-data-studio-and-cloud-dataprep-are-now-generally-available) |
| Cloud Dataprep | GA | 2018/09/21 | 同上 |
| Cloud Security Command Center | beta | 2018/12/05 | [Cloud Security Command Center is now in beta and ready to use](https://cloud.google.com/blog/products/identity-security/cloud-security-command-center-is-now-in-beta) |

# 論文

いくつかのプロダクトについては論文で詳細が公開されています。なお、論文のポイントはグーグル合同会社の中井さんが連載で紹介してくださっています。

* [グーグル合同会社 中井悦司氏によるグーグルクラウドに関連する技術コラム](https://www.school.ctc-g.co.jp/columns/nakai2/)

また、論文の一覧を表にしました。

| プロダクト名 | 発行年 | 論文 |
|---|---|---|
| Google File Sysytem | 2003 | [The Google File System](https://ai.google/research/pubs/pub51) |
| Bigtable | 2006 | [Bigtable: A Distributed Storage System for Structured Data](https://ai.google/research/pubs/pub27898)  |
| Chubby | 2006 | [The Chubby lock service for loosely-coupled distributed systems](https://ai.google/research/pubs/pub27897)  |
| Dremel | 2010 | [Dremel: Interactive Analysis of Web-Scale Datasets](https://ai.google/research/pubs/pub36632)  
| FlumeJava | 2010 | [FlumeJava: Easy, Efficient Data-Parallel Pipelines](https://ai.google/research/pubs/pub35650)  
| Megastore | 2011 | [Megastore: Providing Scalable, Highly Available Storage for Interactive Services)](https://ai.google/research/pubs/pub36971)  |
| Spanner| 2012 | [Spanner: Google's Globally-Distributed Database](https://ai.google/research/pubs/pub39966)|
| MilWheel | 2013 | [MillWheel: Fault-Tolerant Stream Processing at Internet Scale](https://ai.google/research/pubs/pub41378)  |
| Borg | 2015 | [Large-scale cluster management at Google with Borg](https://ai.google/research/pubs/pub43438)  |
| Borg | 2016 | [Borg, Omega, and Kubernetes](https://ai.google/research/pubs/pub44843)  |
| Maglev | 2016 | [Maglev: A Fast and Reliable Software Network Load Balancer](https://ai.google/research/pubs/pub44824)  |

# Google の方の資料

Google の歴史やテクノロジーの掘り下げには Google の方が作成された資料もあります。その一部を紹介させていただきます。

* [YAPC Asia 2015「Google Cloud Platformの謎テクノロジーを掘り下げる」のまとめ](https://qiita.com/kazunori279/items/3ce0ba40e83c8cc6e580)  
* [Googleのインフラ技術から考える理想のDevOps](https://www.slideshare.net/enakai/googledevops)  
* [GCP誕生から10年、その進化の歴史を振り返る](http://ascii.jp/elem/000/001/757/1757103/)

# 最後に

記事が膨大で力尽きつつあります。。

なお、筆者は Google の回し者ではありません。(というより、仕事でクラウドを使ったことがありません・・。使いたいです！)
