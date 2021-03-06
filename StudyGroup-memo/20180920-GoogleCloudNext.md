**Day2**

# 基調講演

基調講演はこちらに公式の記事が出ています。  
https://cloud-ja.googleblog.com/2018/09/next-tokyo-2-announcement.html

以下は私がとったメモです。

##### Cloud Armor

SQL インジェクション攻撃についてルールを追加するだけで、160 を超えるルールを防いでくれる。

DDos 攻撃を受けた際は、NG の IP アドレスをブラックリストに追加する。
これは、BiqQuery ML によって自動的に検出している。

**Cloud Memory Store が東京リージョンで GA リリース**

##### コロプラ様

* GCP 移行。一年で 9 割が移行。
* ライブマイグレーションや Spanner といったテクノロジー、文化を取り入れることが選定動機。
* GKE。マイクロサービスに寄せていった。
* Spanner。こんなに簡単に負荷対策できるとは、と驚いた。
* SRE。当初は Dev, Ops の橋渡しだったが、今は、エンジニアとビジネスの共通言語という認識。ビジネスに目線が向いている。

##### 佐藤さん

* 80 % の企業がマルチクラウド
* 2005 年と比べると、サーバコストは 15 % 減だが、管理コストは８割ほど増えている。
* クラウドごとに仕様が異なり、定型作業もやり方が異なる問題がある。
* Kubernetes、コンテナの技術により違いを吸収する。
* 75 % の企業が Kubernetes を使用（アンケート結果より）
* サービスメッシュへの対応 → Istio。
* Cloud Service Platform(CSP): ワークロードがどこにあるかに関係なく、一貫したインタフェースで管理を行うことができる。GKE on-prem, istio, StackDriver の拡張などが含まれている。

##### yuta さん

サーバレスの話。全ての環境で同じようにサーバレスのワークロードを行えるようなプロダクトをローンチしている。

* 運用モデル：オープン、イベントドリブン、サービスベース。
* knative：オンプレミスで動作可能。

##### メロディーさん

DevOps は大事で、開発者の好きなツールを選べるようにしている。しかし、セキュリティとのバランスは大事で忘れてはならない。

skaffold がソースの変更を検知して、イメージをビルドして ローカルの Kubernetes クラスタにデプロイしてくれる。GitHub 上でプルリクをマージすると Cloud Build が走って、Container Registry にアップロードされる。脆弱性スキャンもされる。Spinnaker でカナリアリリースする。Stackdriver ではバージョン間の性能比較機能がある。

##### 菅野さん

経営者に確認したクラウド選定のポイント

* アジリティの重要性
* セキュリティへの信頼
* 経費削減

##### エイミーさん

**メールセキュリティ**

1 分間に 1000 万通のスパムメールをブロックしている。
情報保護モードでメールの送受信ができる。有効期限を設けたり、パスコードを携帯電話経由で受け取ったり。

##### NEC様

マルチクラウド対応を進めている。
イノベーションに対する拘り、熱意に共感するところがあった。Google はサービスの完成度が高い。Speech API に注目している。精度だけでなく、多言語に対応している。

NEC ネッツエスアイが GCP の各メニューを揃えている。
NEC マネジメントパートナーで GCP の教育メニューを提供予定。エンジニアを増やしていく。

# 企業として正しく使うための GCP セキュリティ

組織として適切に運用するためのセキュリティがこのセッションの注力ポイント。
2、3 年前は組織内の 1 つのプロジェクトで BigQuery を使う流れだったが、複数プロジェクトで採用される流れに変わってきた。なので、組織全体で考える必要がある。

##### 認証、アカウント

Google アカウントとは？

* 個人管理と企業管理の 2 種類がある
* 企業管理のものは企業ドメイン(Cloud Identity, Cloud Identity Premium, G Suite)

企業管理のアカウント

* アカウントのライフサイクル管理が可能（退職時など）
  * AD との連携が可能（Google Cloud Directory Sync で同期できる）
* 2 要素認証の必須化
* SAML 対応
* セットアップにドメイン認証が必要（DNS の TXT レコードの追加が必要）

##### 階層構造

**組織、フォルダ、プロジェクト、リソースの階層構造**

* フォルダ：複数のプロジェクトをまとめる。部門やグループ会社など、権限移譲しても良い相手に対して作成することを推奨。
* プロジェクト：組織なしのプロジェクトはオーナーが消えると消えてしまう点に注意。

**プロジェクトの払い出し方**

* 環境を分ける
  * アプリケーションごと、マイクロサービスごと
  * 環境種別ごと（本番、開発、ステージング、検証）
* 同一ユーザに対する権限を分けたい
  * プロジェクト内で権限を小分けするのではなく、プロジェクト単位で分けるのがオススメ。
* 管理者はプロジェクトの命名規則などを作っておく。

**組織ポリシー**

様々なポリシーを設定できる。

* サービスアカウントの作成を無効化
* API を制限

##### 権限管理

**IAM**

* グループに対して割り当てるのが推奨

**サービスアカウント**

* アプリケーションやサーバのための ID
* アンチパターンは Google アカウントをコード内に書くこと。代わりにサービスアカウントを使う。

**役割**

* 複数の権限をまとめたもの
* 基本の役割がオーナ、編集者、閲覧者。ただ、これからはなるべく使わないでほしい。プロジェクト全体に対して設定できるので、最小権限を実現できない。
* 事前定義された役割を使うのが良い
* それで対応できない場合はカスタムの役割
* resourcemanager.projects.setIamPolicy はかなり強い役割。

**IAM Conditions**

アクセスする際の属性で割り当てられる役割を変更。
設定可能な属性は次の通り。

* 時間
* 接続元 IP アドレス

##### オンプレミスと GCP の通信

**VPC と接続する方法**

* VPN
* Dedicated interconnect (指定拠点まで、自社の NW を伸ばす必要がある)
* Partner Interconnect

**VPC**

* 共有 VPC を使用して接続ポリシーを中央管理する。
* VPC Service Control を有効化していれば、ゾーンをまたいだ通信を禁止できる。

##### 監査対応

* Stackdriver Logging に蓄積するのがベストプラクティス
* 監査ログの種類
  * Admin Activity Log
  * Data Access Log
  * Access Transparency log（Google サポートやエンジニアのアクセス）
* 監査用プロジェクトに各プロジェクトのログを集約するのがオススメ
* ログが多すぎる時はフィルタで除外設定もできる

# ソフトバンクの GCP 利用事例 〜社員の膨大な活動データを活かす方法〜

* 社内で徹底活用したものをお客様に情報提供している

**GCP をどのように活用しているか？**

* カレンダーのデータを分析し、最適化に向けた施作を実施。
* サーバのサイジング予測
* 商品の需要予測（精度は人と同じくらいで工数を 1/10 になった）
* スタッフの勤務シフト予測

**開発プロセス**

データの探索、簡易なモデルでベンチマーク、モデル開発、デプロイ、予測

**意思決定への貢献を可能に**

* 施作の実行
* 調達、開発の決定
* 高速な効果測定

##### Peaple Analytics

営業の人 300 人を分析。BigQuery は非構造のデータを入れて可視化していくことができるので面白い。パッション、ビジネスセンス、巻き込み力、スピード、コーディネート力、伝達力を可視化して、ビジネスリソースの最適化をしたい。

通信の秘密はあるが、Mail の From, To をみる、アクセス記録を見るなどして、分析。どのワークログに価値があるのか、まだ誰も分かっていない。

プロフィール、ソーシャルスタイル、誰とコラボレーションしているか、働く時間、働く環境（温度、湿度、色、匂いなど）から最適な環境を探っていく。

社員の行動管理が目的ではない。さらなる成長、活躍の機械を与えることが目的。

（所感：スピーカーの方のパッションがすごかった）

# システム運用のプロがおくる実際の構成を基にした Stackdriver の効果的な活用術

Stackdriver アカウントを作る必要がある。
Stackdriver にはいくつかのサービスがある。

* Monitoring
* Logging
* Error Report
* Debugger
* Trace
* Profiler

Agent のインストールが必要

* Monitoring
  * stackdriver-agent
* Logging
  * google-fluentd


**Monitoring**

* ダッシュボード上に必要なメトリクスを設定し、可視化する。
* Apache, nginx, MySQL, postgres, Redis など様々なアプリに対応
* URL 応答監視を設定できる
* メール、Slack、PagerDuty(オンコールしてくれるサービス) など様々な方法で通知可能

**まとめ**

* 各サービスで使用しているリソース全てが見えるのでデフォルトだと見にくい。絞る必要がある。
* デフォルト実装されていないアプリの場合は、カスタムメトリクスの実装が必要で大変

# BigQuery を使用した分析基盤の運用を進めていく上で見えてきた課題、乗り越えてきた軌跡

BigQuery の導入でバッチ処理にかかる時間が劇的に減った。

**今までの課題**

* オンプレの Hadoop はスケールアウトが辛い
* ストレージのキャパシティ管理が辛い。オンプレとの通信が辛い。
* データマート作成に1日かかる

→ DataLake(S3)、BigQuery の導入

**運用負荷が高い状況の解消**

* 性能劣化しない基盤
* インフラ運用からの解放
* キャパシティ管理からの解放
* データ活用の民主化が進む基盤
* 構造を把握しやすい基盤

**BigQuery の利点**

* キャパシティ管理不要
* フルマネージド
* ロード処理、クエリが分かれている
* データの受け渡しが用意（他でも BigQuery を使っていると権限の付与だけで良い）
* 定額課金での使用も可能

##### 運用

* 定額なので、slots が上限となっている。プロジェクトごとに slot を割振って、大事なクエリが実行されるようにしている。
* 権限管理。google group で行っている
* ユーザ教育。教育動画や勉強会を開催。新機能はメルマガでお知らせ
* メタデータ管理。他の DB も一元的に見れるようにしている
* DataLake 構成
* 他の DWH とのデータ連携は bot でできるように

##### 工夫

データソースが多岐にわたり、大規模。
Google Analytics のデータは簡単に BigQuery に流せるので簡単。ただ、他の製品だと、そう簡単にいかない場合も。

設計思考はイベントドリブン、疎結合。

S3 バケット通知イベントのSQSメッセージを起点に、その後に続く各処理も SQS によるキュー管理で連携処理を実行。GKE 経由でファイルコピー等を行う。キューがたまってきたら、GKE をスケールさせる。

* プロジェクトは事業、用途単位で分割している。
* 権限管理は Google Group。
* スキーマの変換。変換モジュールを作った。
* エスケープ文字に非対応なので、クレンジング処理を入れた。
* テーブル作成の生産性向上。CREATE TABLE が公式にサポートされた。
* 無尽蔵にテーブルを作られるリスクに対しては、テーブルサイズの保存期間のデフォルト値を設定した。（ただし、サイズは設定できない。監視は可能）

# テクノロジーが活きる現場を作る：チェンジマネジメント対談

チェンジマネジメントとは？
いかにテクノロジーが活きる現場を作るか。についての対談。

G Suite を使って会社の中に変化を起こしていく。いかに現場、トップを巻き込んでやっていくか。

**ファミリーマート様**

6 月末に決まったものが 10 月に導入しないといけない。
推進リーダーを指名して、その人たちにレクチャーして、業務を止めずに移行した。現場の方達に当事者意識を持ってもらうのが大事。アンバサダーという形で発信してもらう立場の方を任命した。また、部長を集めて社内の理解を深めた。  
使い方を自分たちで考えてやってもらったのが、結果的には良かった。ただ、教えて欲しいという要望も多かったので、全社員教育を行う。今となっては講師も内製でできるので、外部から呼ぶ必要もなくなった。

全員参加型に持っていくには、社内のボトムアップ活動に持っていかなければならない。

**Jフロントリテイリング様**

まず社長にこの機能を使えば武器となるとアピールして、社長を巻き込んだ。社内報やビデオで社長にも出演していただいて、この武器を使って変革をもたらしてくれと発信してくれた。トップの人に良さを体感してもらう。  
拠点が多かったので、各拠点にエバンジェリストを立てて、先行利用してもらった。初めは嫌がっていた人がいたけど、役割をちゃんと伝えて、みなさんがやるんですよと伝えた。そうすることで自分ごととして捉えてくれるようになった。  
フェーズごとにわけた。1 フェーズ目はハングアウト、2 フェーズ目は Google Drive。3 フェーズ目は共同編集。うまく使っている部署の取り組みをピックアップして周りに伝えていった。新しいツールをポンと与えるだけでなく、現場の方がどういう課題を持っていて、Before & After を示しながら発信していくのが大事。もうちょっと、イベント感、お祭り感を出せれば、というところが課題と感じたところ。

社員みんなに伝えていくのが大事。そのためには仕組みと仕掛け。それを継続していくのも難しい。変化をすると大雑把に問いかけるのではなく、第一ステップとしてこれをやるという分かりやすいメッセージに落とし込む。いつまでに何をやるかを分かりやすく伝える。

# scikit-learn 開発者が送る、Google Cloud と機械学習の実力

### scikit-learn

XGBoost=勾配ブースティング決定木。Kaggle で勝利をおさめたエントリーで多く使用されている。

Tensorflow との違いはディープラーニング向けに最適化されていないということ。ただし、ビジネスの現場では scikit-learn の XGBoost が適しているケースが多い。なぜその結果が出たのかを理解しやすいメリットもある。一方で、非構造データの前処理は難しい。ここがニューラルネットワークの優れているところ。Google の各種機械学習 API は非構造データの分類に有用。

XGBoost は並列化が容易。並列化というと分散処理を考えがちだが、大きなインスタンスを使えば 1 台で対応可能。

ビッグデータという言葉は誤用されている。典型的なデータサイズはそこまで大きくない。サイジングを大げさに見積もる必要はない。クラスタの構築に時間もかかるし、アジリティの消失を過小評価すべきではない。また、早い段階でスキーマを決めると後で変えられないのでアンチパターン。

近年は、スケールアップが容易になってきた。4 TB の RAM の VM を利用可能。NVME SSD による高速な I/O が可能。シングルマシンなので、結果もすぐ帰ってくる。scikit-learn など慣れたツールを使うのも容易。CSV を pandas などの効率的なフォーマットに変換できる。また、近年 pandas のメモリ効率が上がっている。

**Deep Learning VM イメージ**  
機械学習のパッケージ類がプリインストールされている。

**Cloud ML Engine**  
scikit-learn をサポート。

**Kubeflow**  
どこでも誰でも使えることを目標に開発している。
