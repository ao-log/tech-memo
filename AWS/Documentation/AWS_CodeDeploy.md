
# CodeDeploy

## はじめに

[CodeDeploy とは](https://docs.aws.amazon.com/ja_jp/codedeploy/latest/userguide/welcome.html)

CodeDeploy はアプリケーションのデプロイを自動化するサービス。デプロイ先は次の 3 つをサポート。

* EC2 (オンプレミスも可能)
* ECS
* Lambda 関数


#### 用語集

* アプリケーション：　最上位に位置。デプロイグループ、リビジョンの管理単位。
* デプロイグループ：　デプロイに関する設定を行う。
* デプロイ：　個々のデプロイに対する設定を行う。
* リビジョン：　AppSpec ファイルとアプリケーションの組み合わせ。


#### デプロイタイプ

インプレースデプロイと Blue/Green Deployment の 2 種類。


##### インプレースプロイ

EC2 or オンプレミスでのみ使用可能。
デプロイ対象は指定タグが付いているインスタンスもしくは Auto Scaling Group。
デプロイ中は ELB から登録解除され、デプロイ後に登録される。

**処理の流れ**

1. S3 または GitHub にアーカイブしたファイル（アプリケーションリビジョンと呼ぶ）を配置。
1. CodeDeploy に必要な情報（デプロイグループ = 対象の EC2 or Auto Scaling Group、アプリケーションリビジョンの情報など）。
1. デプロイ先の CodeDeploy Agent が CodeDeploy をポーリング。
1. CodeDeploy Agent がアプリケーションリビジョンを取得し、AppSpec file の内容に従ってデプロイ。


#### Blue/Green デプロイ

**EC2 /オンプレミス**

置き換え先のインスタンスを対象にデプロイする。
置き換え先のインスタンスが ELB に登録され、トラフィックはそれらに再ルーティングされる。
元のインスタンスは ELB から登録解除される。インスタンスは削除することもできるし残すこともできる。

**ECS**

新しいタスクセットにトラフィックが移行される。
トラフィックの移行は All-at-once、線形または Canary から選択できる。

**Lambda**

新しいバージョンの Lambda 関数にトラフィックを移行する。


[デプロイ - AWS Lambda コンピューティングプラットフォームを選択](https://docs.aws.amazon.com/ja_jp/codedeploy/latest/userguide/deployment-steps-lambda.html)

Lambda 関数のデプロイまでの流れが書かれているページ。


[デプロイ - Amazon ECS コンピューティングプラットフォームを選択](https://docs.aws.amazon.com/ja_jp/codedeploy/latest/userguide/deployment-steps-ecs.html)

ECS のデプロイまでの流れが書かれているページ。

以下の内容のリソースを予め作っておく必要がある（デプロイグループで指定）

* Amazon ECS クラスター
* Amazon ECS サービス
* Application Load Balancer または Network Load Balancer
* 本稼働リスナー
* テストリスナー (オプション)
* 2 つのターゲットグループ

デプロイ時は次のステップで各イベントが処理される。

* BeforeInstall
* Install: Blue のタスクセットにデプロイ。
* AfterInstall
* AllowTestTraffic: テストリスナーからターゲットグループ 2 にトラフィックをルーティング。
* AfterAllowTestTraffic: フックによりテストリスナーに対しテストできる。失敗した場合、ロールバックが有効になっているとロールバックされる。
* BeforeAllowTraffic:
* AllowTraffic:	本稼働リスナーからターゲットグループ 2 にトラフィックをルーティング。
* 全てのイベントが完了したら、デプロイステータスは Succeeded になり元のタスクセットは削除される。

デプロイが失敗した場合などにロールバックするよう設定できる。新しいデプロイ ID が割り当てられ、元のタスクセットにルーティングする処理を行う。


[デプロイ - EC2/オンプレミス コンピューティングプラットフォームを選択](https://docs.aws.amazon.com/ja_jp/codedeploy/latest/userguide/deployment-steps-server.html)

EC2 インスタンスへのデプロイまでの流れが書かれているページ。

以下の内容のリソースを予め作っておく必要がある。

* インスタンスに CodeDeploy エージェントをインストール
* タグを使用してデプロイメントグループ内のインスタンスを識別する場合は、タグを有効にする。
* IAM インスタンスプロファイル
* サービスロール

デプロイ時は次の流れで処理される。

* インプレースデプロイ: 最新のアプリケーションリビジョンでインスタンスを更新
* Blue/Green デプロイ: ロードバランサーでデプロイグループ用の代替セットを登録し、元のインスタンスを登録解除

デプロイが失敗した場合などにロールバックするよう設定できる。新しいデプロイ ID が割り当てられ、以前にデプロイされたリビジョンを再デプロイする動作となる。



## 他のサービスとの統合

[CodeDeploy と Amazon EC2 Auto Scaling の統合](https://docs.aws.amazon.com/ja_jp/codedeploy/latest/userguide/integrations-aws-auto-scaling.html)

デプロイ中にスケールアウトが発生すると新規起動したインスタンスには前回のリビジョンが反映される。この動作を防ぐにはデプロイ中に ASG のサスペンドが必要。


[CodeDeploy と Elastic Load Balancing の統合](https://docs.aws.amazon.com/ja_jp/codedeploy/latest/userguide/integrations-aws-elastic-load-balancing.html)


**Blue/Green Deployment の場合**

置き換え先環境のインスタンスが登録されると、置き換え前環境のインスタンスが登録解除される。

**インプレースデプロイメントの場合**

インプレースデプロイ中は、デプロイ先のインスタンスに対するインターネットトラフィックのルーティングがブロックされ、デプロイ完了時点でルーティングが再開される。



## チュートリアル

[チュートリアル: CodeDeploy を使用して Amazon EC2 Auto Scaling グループにアプリケーションをデプロイする](https://docs.aws.amazon.com/ja_jp/codedeploy/latest/userguide/tutorials-auto-scaling-group.html)

以下流れで作業。

Auto Scaling Group の作成、設定。

```shell
# 起動設定の作成
aws autoscaling create-launch-configuration \
  --launch-configuration-name CodeDeployDemo-AS-Configuration \
  --image-id image-id \
  --key-name key-name \
  --iam-instance-profile CodeDeployDemo-EC2-Instance-Profile \
  --instance-type t1.micro

# Auto Scaling Group の作成
aws autoscaling create-auto-scaling-group \
  --auto-scaling-group-name CodeDeployDemo-AS-Group \
  --launch-configuration-name CodeDeployDemo-AS-Configuration \
  --min-size 1 \
  --max-size 1 \
  --desired-capacity 1 \
  --availability-zones availability-zone \
  --tags Key=Name,Value=CodeDeployDemo,PropagateAtLaunch=true

# CodeDeploy エージェントのインストール
aws ssm create-association \
  --name AWS-ConfigureAWSPackage \
  --targets Key=tag:Name,Values=CodeDeployDemo \
  --parameters action=Install, name=AWSCodeDeployAgent \
  --schedule-expression "cron(0 2 ? * SUN *)" 
```

アプリケーションをデプロイ

```shell
# アプリケーションの作成
aws deploy create-application --application-name SimpleDemoApp

# デプロイグループの作成
aws deploy create-deployment-group \
  --application-name SimpleDemoApp \
  --auto-scaling-groups CodeDeployDemo-AS-Group \
  --deployment-group-name SimpleDemoDG \
  --deployment-config-name CodeDeployDefault.OneAtATime \
  --service-role-arn service-role-arn

# デプロイ実行
aws deploy create-deployment \
  --application-name SimpleDemoApp \
  --deployment-config-name CodeDeployDefault.OneAtATime \
  --deployment-group-name SimpleDemoDG \
  --s3-location bucket=bucket-name,bundleType=zip,key=samples/latest/SampleApp_Linux.zip
```

なお、Auto Scaling グループ内にインスタンスが新規作成されたときにも自動的にデプロイ(インプレースデプロイ)される。


[チュートリアル Amazon ECS サービスをデプロイする](https://docs.aws.amazon.com/ja_jp/codedeploy/latest/userguide/tutorial-ecs-deployment.html)

デプロイ時は次のような appspec ファイルを使用する。

```yaml
version: 0.0
Resources:
  - TargetService:
      Type: AWS::ECS::Service
      Properties:
        TaskDefinition: "arn:aws:ecs:aws-region-id:aws-account-id::task-definition/ecs-demo-task-definition:revision-number"
        LoadBalancerInfo:
          ContainerName: "your-container-name"
          ContainerPort: your-container-port
```          

デプロイグループでは以下の設定を行う。

* ECS クラスター名
* ECS サービス名
* ELB 名
* 本番用リスナーポート
* テストリスナーポート(オプショナル)
* ターゲットグループ 1
* ターゲットグループ 2
* Traffic rerouting (すぐにトラフィックを転送するか、指定した時間待機するか)


[チュートリアル: 検証テストを使用して Amazon ECS サービスをデプロイする](https://docs.aws.amazon.com/ja_jp/codedeploy/latest/userguide/tutorial-ecs-deployment-with-hooks.html)

このチュートリアルでは AfterAllowTestTraffic 中にテストを実行する。

Lambda 関数は [ステップ 3: ライフサイクルフックLambda関数を作成する](https://docs.aws.amazon.com/ja_jp/codedeploy/latest/userguide/tutorial-ecs-with-hooks-create-hooks.html) で作成している。

appspec ファイルでは次のようにライフサイクルフックを指定する。

```yaml
version: 0.0
Resources:
  - TargetService:
      Type: AWS::ECS::Service
      Properties:
        TaskDefinition: "arn:aws:ecs:aws-region-id:aws-account-id::task-definition/ecs-demo-task-definition:revision-number"
        LoadBalancerInfo:
          ContainerName: "sample-website"
          ContainerPort: 80
Hooks:
  - AfterAllowTestTraffic: "arn:aws:lambda:aws-region-id:aws-account-id:function:AfterAllowTestTraffic"
```



## CodeDeploy エージェント

[CodeDeploy エージェントの使用](https://docs.aws.amazon.com/ja_jp/codedeploy/latest/userguide/codedeploy-agent.html)




## インスタンス

[Tagging instances for deployment groups in CodeDeploy](https://docs.aws.amazon.com/ja_jp/codedeploy/latest/userguide/instances-tagging.html)

デプロイグループの設定で、デプロイ対象とする EC2 インスタンスをタグで指定可能。



## デプロイ設定

[CodeDeploy でのデプロイ設定の使用](https://docs.aws.amazon.com/ja_jp/codedeploy/latest/userguide/deployment-configurations.html)

* EC2
  * CodeDeployDefault.AllAtOnce: 全台失敗した場合のみデプロイ失敗と判定
  * CodeDeployDefault.OneAtATime: 1 度に 1 台ずつデプロイ。インスタンスへのデプロイ失敗した場合はデプロイ全体が失敗判定となる。例外は最後のインスタンスの場合でデプロイ失敗しても成功判定となる。
* ECS
  * トラフィックの移行方法は 3 通りある。
    * Canary
    * 線形
    * All-at-once



## アプリケーション

[CodeDeploy のアプリケーションでの作業](https://docs.aws.amazon.com/ja_jp/codedeploy/latest/userguide/applications.html)



## デプロイグループ

[CodeDeploy でのデプロイグループの使用](https://docs.aws.amazon.com/ja_jp/codedeploy/latest/userguide/deployment-groups.html)

* 複数のデプロイグループを CodeDeploy のアプリケーションに関連付けることが可能
* 例えば、デプロイグループごとに環境名(prod, staging など)のタグが付与された EC2 インスタンスを対象にする使い方が考えられる


[インプレースデプロイ用のデプロイグループを作成する (コンソール)](https://docs.aws.amazon.com/ja_jp/codedeploy/latest/userguide/deployment-groups-create-in-place.html)

* デプロイ対象として Auto Scaling グループ名もしくは EC2 インスタンスのタグ名を設定可能。
* 各インスタンスは、デプロイ中に ELB から登録解除されて、トラフィックがルーティングされなくなる。デプロイ完了時に再登録される。


[EC2/オンプレミス Blue/Green デプロイ用のデプロイグループを作成する (コンソール)](https://docs.aws.amazon.com/ja_jp/codedeploy/latest/userguide/deployment-groups-create-blue-green.html)

* Traffic rerouting の設定は 2 つから選ぶ。
  * [Reroute traffic immediately]: 置き換え先環境のインスタンスがプロビジョニングされ、最新のアプリケーションリビジョンがインストールされるとすぐに、ロードバランサーに自動的に登録され、トラフィックがそれらに再ルーティングされる。元の環境内のインスタンスは、登録解除される。
  * [I will choose whether to reroute traffic]: 置き換え先環境のインスタンスは、手動でトラフィックを再ルーティングしないかぎり、ELB に登録されない。指定した待機時間中にトラフィックが再ルーティングされない場合、デプロイステータスは停止に変更される。


[デプロイ用のAmazon ECSデプロイグループを作成する (コンソール)](https://docs.aws.amazon.com/ja_jp/codedeploy/latest/userguide/deployment-groups-create-ecs.html)

* テストリスナーを設定した場合、AfterAllowTestTraffic フック時に実行する 1 つ以上の Lambda 関数を AppSpec file で指定可能。この関数では検証テストを実行でき、検証テストに失敗するとデプロイのロールバックがトリガーされる。検証テストに成功すると、デプロイのライフサイクルの次のフック BeforeAllowTraffic がトリガーされる。テストリスナーポートが指定されていない場合は、AfterAllowTestTraffic フック時に何も行わない。
* Traffic rerouting の設定は 2 つから選ぶ。
  * [すぐにトラフィックを再ルーティング]: トラフィックは、置き換えタスクセットがプロビジョニングされた後でデプロイにより自動的に再ルーティングされる。
  * [トラフィックを再ルーティングするタイミングを指定します]: 置き換えタスクセットが正常にプロビジョニングされた後に待機する日数や時間数、および分数を選択。この待機時間中に、AppSpec file で指定された Lambda 関数の検証テストが実行される。トラフィックが再ルーティングされる前に待機時間が経過すると、デプロイステータスは停止となる。



## アプリケーションリビジョン

[CodeDeploy のアプリケーション リビジョンでの作業](https://docs.aws.amazon.com/ja_jp/codedeploy/latest/userguide/application-revisions.html)


[CodeDeploy のリビジョンの計画](https://docs.aws.amazon.com/ja_jp/codedeploy/latest/userguide/application-revisions-plan.html)

* ECS, Lambda はリビジョン = appspec ファイル。
* EC2 の場合は、リビジョン = デプロイ対象のファイル。


[appspec ファイルを CodeDeploy のリビジョンに追加する](https://docs.aws.amazon.com/ja_jp/codedeploy/latest/userguide/application-revisions-appspec-file.html)

**ECS の場合の書式**

```yaml
# This is an appspec.yml template file for use with an Amazon ECS deployment in CodeDeploy.
# The lines in this template that start with the hashtag are 
#   comments that can be safely left in the file or 
#   ignored.
# For help completing this file, see the "AppSpec File Reference" in the  
#   "CodeDeploy User Guide" at
#   https://docs.aws.amazon.com/codedeploy/latest/userguide/app-spec-ref.html
version: 0.0
# In the Resources section, you must specify the following: the Amazon ECS service, task definition name, 
# and the name and port of the your load balancer to route traffic,
# target version, and (optional) the current version of your AWS Lambda function. 
Resources:
  - TargetService:
      Type: AWS::ECS::Service
      Properties:
        TaskDefinition: "" # Specify the ARN of your task definition (arn:aws:ecs:region:account-id:task-definition/task-definition-family-name:task-definition-revision-number)
        LoadBalancerInfo: 
          ContainerName: "" # Specify the name of your Amazon ECS application's container
          ContainerPort: "" # Specify the port for your container where traffic reroutes 
# Optional properties
        PlatformVersion: "" # Specify the version of your Amazon ECS Service
        NetworkConfiguration:
          AwsvpcConfiguration:
            Subnets: ["",""] # Specify one or more comma-separated subnets in your Amazon ECS service
            SecurityGroups: ["",""] # Specify one or more comma-separated security groups in your Amazon ECS service
            AssignPublicIp: "" # Specify "ENABLED" or "DISABLED"             
# (Optional) In the Hooks section, specify a validation Lambda function to run during 
# a lifecycle event. 
Hooks:
# Hooks for Amazon ECS deployments are:
    - BeforeInstall: "" # Specify a Lambda function name or ARN
    - AfterInstall: "" # Specify a Lambda function name or ARN
    - AfterAllowTestTraffic: "" # Specify a Lambda function name or ARN
    - BeforeAllowTraffic: "" # Specify a Lambda function name or ARN
    - AfterAllowTraffic: "" # Specify a Lambda function name or ARN
```

**EC2 の場合の書式**

```yaml
# This is an appspec.yml template file for use with an EC2/オンプレミス deployment in CodeDeploy.
# The lines in this template starting with the hashtag symbol are 
#   instructional comments and can be safely left in the file or 
#   ignored.
# For help completing this file, see the "AppSpec File Reference" in the  
#   "CodeDeploy User Guide" at
#   https://docs.aws.amazon.com/codedeploy/latest/userguide/app-spec-ref.html
version: 0.0
# Specify "os: linux" if this revision targets Amazon Linux, 
#   Red Hat Enterprise Linux (RHEL), or Ubuntu Server  
#   instances.
# Specify "os: windows" if this revision targets Windows Server instances.
# (You cannot specify both "os: linux" and "os: windows".)
os: linux 
# os: windows
# During the Install deployment lifecycle event (which occurs between the 
#   BeforeInstall and AfterInstall events), copy the specified files 
#   in "source" starting from the root of the revision's file bundle 
#   to "destination" on the Amazon EC2 instance.
# Specify multiple "source" and "destination" pairs if you want to copy 
#   from multiple sources or to multiple destinations.
# If you are not copying any files to the Amazon EC2 instance, then remove the
#   "files" section altogether. A blank or incomplete "files" section
#   may cause associated deployments to fail.
files:
  - source: 
    destination:
  - source:
    destination:
# For deployments to Amazon Linux, Ubuntu Server, or RHEL instances,
#   you can specify a "permissions" 
#   section here that describes special permissions to apply to the files 
#   in the "files" section as they are being copied over to 
#   the Amazon EC2 instance.
#   For more information, see the documentation.
# If you are deploying to Windows Server instances,
#   then remove the 
#   "permissions" section altogether. A blank or incomplete "permissions"
#   section may cause associated deployments to fail.
permissions:
  - object:
    pattern:
    except:
    owner:
    group:
    mode: 
    acls:
      -
    context:
      user:
      type:
      range:
    type:
      -
# If you are not running any commands on the Amazon EC2 instance, then remove 
#   the "hooks" section altogether. A blank or incomplete "hooks" section
#   may cause associated deployments to fail.
hooks:
# For each deployment lifecycle event, specify multiple "location" entries 
#   if you want to run multiple scripts during that event.
# You can specify "timeout" as the number of seconds to wait until failing the deployment 
#   if the specified scripts do not run within the specified time limit for the 
#   specified event. For example, 900 seconds is 15 minutes. If not specified, 
#   the default is 1800 seconds (30 minutes).
#   Note that the maximum amount of time that all scripts must finish executing 
#   for each individual deployment lifecycle event is 3600 seconds (1 hour). 
#   Otherwise, the deployment will stop and CodeDeploy will consider the deployment
#   to have failed to the Amazon EC2 instance. Make sure that the total number of seconds 
#   that are specified in "timeout" for all scripts in each individual deployment 
#   lifecycle event does not exceed a combined 3600 seconds (1 hour).
# For deployments to Amazon Linux, Ubuntu Server, or RHEL instances,
#   you can specify "runas" in an event to
#   run as the specified user. For more information, see the documentation.
#   If you are deploying to Windows Server instances,
#   remove "runas" altogether.
# If you do not want to run any commands during a particular deployment
#   lifecycle event, remove that event declaration altogether. Blank or 
#   incomplete event declarations may cause associated deployments to fail.
# During the ApplicationStop deployment lifecycle event, run the commands 
#   in the script specified in "location" starting from the root of the 
#   revision's file bundle.
  ApplicationStop:
    - location: 
      timeout:
      runas:
    - location: 
      timeout:
      runas: 
# During the BeforeInstall deployment lifecycle event, run the commands 
#   in the script specified in "location".
  BeforeInstall:
    - location: 
      timeout:
      runas: 
    - location: 
      timeout:
      runas:
# During the AfterInstall deployment lifecycle event, run the commands 
#   in the script specified in "location".
  AfterInstall:
    - location:     
      timeout: 
      runas:
    - location: 
      timeout:
      runas:
# During the ApplicationStart deployment lifecycle event, run the commands 
#   in the script specified in "location".
  ApplicationStart:
    - location:     
      timeout: 
      runas:
    - location: 
      timeout:
      runas:
# During the ValidateService deployment lifecycle event, run the commands 
#   in the script specified in "location".
  ValidateService:
    - location:     
      timeout: 
      runas:
    - location: 
      timeout:
      runas:
```


[CodeDeploy リポジトリタイプの選択](https://docs.aws.amazon.com/ja_jp/codedeploy/latest/userguide/application-revisions-repository-type.html)

* ECS, Lambda の場合は appspec のみ。以下の方法で指定。
  * S3 バケット上のオブジェクト
  * コンソールの AppSpec エディタ
  * AWS CLI 使用時はファイルパス指定可能
* EC2 の場合はリポジトリとして以下の箇所を使用可能。
  * S3 バケット
  * GitHub
  * Bitbucket



## デプロイ

[CodeDeploy を使用してデプロイを再デプロイおよびロールバックする](https://docs.aws.amazon.com/ja_jp/codedeploy/latest/userguide/deployments-rollback-and-redeploy.html)

* ロールバックは、以前にデプロイされたリビジョンを新しいリビジョンとして再デプロイすることによって行う。
* 自動ロールバックはデプロイグループもしくはデプロイにて設定する。設定した場合、デプロイが失敗した場合、または指定した監視しきい値に達した場合に自動ロールバックし、アプリケーションリビジョンの最後の既知の正常なバージョンがデプロイされる。



## リファレンス

[CodeDeploy AppSpec File リファレンス](https://docs.aws.amazon.com/ja_jp/codedeploy/latest/userguide/reference-appspec-file.html)



# BlackBelt

[20210126 AWS Black Belt Online Seminar AWS CodeDeploy](https://pages.awscloud.com/rs/112-TZM-766/images/20210126_BlackBelt_CodeDeploy.pdf)

* P18: デプロイ設定 - EC2
  * 各設定ごとに v2 をどのようにデプロイしているかを図解している。
* P27: デプロイ設定 - ECS, Lambda
  * Linear, Canary, All-at-once の v2 へのトラフィックのルーティングのされ方を図解している。
* P47: リビジョンの構成 - EC2
* P77: Blue/Green Deployment トラフィックの切り替え - ECS



# 参考

* Document
  * [CodeDeploy とは](https://docs.aws.amazon.com/ja_jp/codedeploy/latest/userguide/welcome.html)
* サービス紹介ページ
  * [AWS CodeDeploy](https://aws.amazon.com/jp/codedeploy/features/)
  * [よくある質問](https://aws.amazon.com/jp/codedeploy/faqs//)
* [ナレッジセンター](https://aws.amazon.com/jp/premiumsupport/knowledge-center/#AWS_CodeDeploy)
* Black Belt
  * [20210126 AWS Black Belt Online Seminar AWS CodeDeploy](https://pages.awscloud.com/rs/112-TZM-766/images/20210126_BlackBelt_CodeDeploy.pdf)

