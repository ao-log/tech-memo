
# CloudFormation

## 用語

* テンプレート: AWS リソースの構成を定義するもの。YAML or JSON 形式のファイル。
* スタック: テンプレートから作られるもの。スタックが管理の単位。テンプレートから新規作成することができ、出来上がったスタックに対して、更新、削除を行うこともできる。
* 変更セット: リソース変更前に変更点を確認可能。



## 基礎

[テンプレートの基礎についての学習](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/gettingstarted.templatebasics.html)

Type でリソースの種類を設定。
Properties で各項目を設定。

```yaml
Resources:
  HelloBucket:
    Type: AWS::S3::Bucket
    Properties:
      AccessControl: PublicRead
      WebsiteConfiguration:
        IndexDocument: index.html
        ErrorDocument: error.html
```

**Ref** 関数を使用することにより、戻り値を取得可能。
また、Parameters を設定することで、スタック作成時に当該項目をユーザに指定させることができる。

```yaml
Parameters:
  KeyName:
    Description: The EC2 Key Pair to allow SSH access to the instance
    Type: 'AWS::EC2::KeyPair::KeyName'
Resources:
  Ec2Instance:
    Type: 'AWS::EC2::Instance'
    Properties:
      SecurityGroups:
        - !Ref InstanceSecurityGroup
        - MyExistingSecurityGroup
      KeyName: !Ref KeyName
      ImageId: ami-7a11e213
  InstanceSecurityGroup:
    Type: 'AWS::EC2::SecurityGroup'
    Properties:
      GroupDescription: Enable SSH access via port 22
      SecurityGroupIngress:
        - IpProtocol: tcp
          FromPort: '22'
          ToPort: '22'
          CidrIp: 0.0.0.0/0
```

**Fn::GetAtt** により他のリソースの属性値を取得可能。利用可能な属性値は [AWS リソースおよびプロパティタイプのリファレンス](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html) より参照可能。

```yaml
Resources:
  myBucket:
    Type: 'AWS::S3::Bucket'
  myDistribution:
    Type: 'AWS::CloudFront::Distribution'
    Properties:
      DistributionConfig:
        Origins:
          - DomainName: !GetAtt 
              - myBucket
              - DomainName
            Id: myS3Origin
            S3OriginConfig: {}
        Enabled: 'true'
        DefaultCacheBehavior:
          TargetOriginId: myS3Origin
          ForwardedValues:
            QueryString: 'false'
          ViewerProtocolPolicy: allow-all
```

Mapping を使用することによりキーと値のペアを定義可能。

```yaml
Parameters:
  KeyName:
    Description: Name of an existing EC2 KeyPair to enable SSH access to the instance
    Type: String
Mappings:
  RegionMap:
    us-east-1:
      AMI: ami-76f0061f
    us-west-1:
      AMI: ami-655a0a20
    eu-west-1:
      AMI: ami-7fd4e10b
    ap-southeast-1:
      AMI: ami-72621c20
    ap-northeast-1:
      AMI: ami-8e08a38f
Resources:
  Ec2Instance:
    Type: 'AWS::EC2::Instance'
    Properties:
      KeyName: !Ref KeyName
      ImageId: !FindInMap 
        - RegionMap
        - !Ref 'AWS::Region'
        - AMI
      UserData: !Base64 '80'
```

**Outputs** によりスタックの出力値を設定可能。

```yaml
Outputs:
  InstallURL:
    Value: !Join 
      - ''
      - - 'http://'
        - !GetAtt 
          - ElasticLoadBalancer
          - DNSName
        - /wp-admin/install.php
    Description: Installation URL of the WordPress website
  WebsiteURL:
    Value: !Join 
      - ''
      - - 'http://'
        - !GetAtt 
          - ElasticLoadBalancer
          - DNSName
```



## ドキュメント

[開始する](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/GettingStarted.Walkthrough.html)

次のステップで作業する。

1. テンプレートの作成
1. テンプレートで使用する依存リソースがすでに存在していることを確認(キーペアなど作成済みのものを使用する場合)
1. スタックの作成
1. スタックの作成状況を確認。イベントを表示してリソース作成状況を確認可能。



## ベストプラクティス

[AWS CloudFormation ベストプラクティス](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/best-practices.html)

* スタックの分割は、例えばチームの責務の範囲といった分け方をする。
* 共有リソースはエクスポートし、利用するスタック側からはクロススタック参照を行う。
* 本番、開発環境も一つのテンプレートで対応できる。パラメータ、マッピング、条件セクションを利用する。
* ネストされたスタックを使用することで、共通リソースの更新時に他のスタックも追随できるようにする。
* テンプレートには認証情報を埋め込まない。
* パラメータには制約を設定できるので、無効な値の入力を防ぐことができる。
* AWS::CloudFormation::Init を使用することにより、EC2 インスタンスの起動時に任意の設定を行うことができる。
* aws cloudformation validate-template により検証を行うことができる。
* **CloudFormation で作成したリソースは、その他の方法で変更しないようにする。**
* スタック更新時には変更セットを使用する。
* 重要なリソースがあるスタックは、スタックポリシーで保護する。
* テンプレートのソースをレポジトリで管理し、リビジョンごとの差分をトレースし、戻せるようにする。



## 継続的デリバリー

[CodePipeline を使用した継続的デリバリー](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/continuous-delivery-codepipeline.html)

CodePipeline を使用して、レポジトリへのテンプレートのコミット → テストスタックの作成 → 本番スタックの作成のようなワークフローのパイプラインを作成可能。



## スタックの操作

### コンソールの使用

[AWS CloudFormation コンソールでのスタックの作成](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/cfn-console-create-stack.html)

以下の項目を設定可能。
* 基本的な項目
  * テンプレート
  * パラメータ
* その他オプション
  * タグ
  * 使用するサービスロール
  * スタックポリシー
  * ロールバック設定（指定したいずれかのアラームのしきい値を超えた場合にロールバック可能）
  * 通知オプション
  * 失敗時のロールバックの有無。デフォルトは有効
  * タイムアウト値
  * スタックの削除保護
  
まず変更セットを使用して、そこから新規にスタックを作成することも可能。


[AWS マネジメントコンソール での AWS CloudFormation スタックデータとリソースの表示](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/cfn-console-view-stack-data-resources.html)

スタックの画面では、以下のタブから各情報を参照可能。

* スタック情報
* イベント
* リソース
* 出力
* パラメータ
* テンプレート

ステータスコード。以下は一部の例。

* CREATE_COMPLETE
* CREATE_IN_PROGRESS
* CREATE_FAILED
* DELETE_COMPLETE
* REVIEW_IN_PROGRESS
* ROLLBACK_COMPLETE
* UPDATE_COMPLETE
* IMPORT_COMPLETE



### AWS CLI

[AWS コマンドラインインターフェイスの使用](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/cfn-using-cli.html)

```shell
# スタックの作成
$ aws cloudformation create-stack --stack-name myteststack \
    --template-body file:///home/testuser/mytemplate.json \
    --parameters ParameterKey=Parm1,ParameterValue=test1 ParameterKey=Parm2,ParameterValue=test2

# スタックの一覧
$ aws cloudformation list-stacks --stack-status-filter CREATE_COMPLETE

# スタックの詳細
$ aws cloudformation describe-stacks --stack-name myteststack

# スタックのイベント履歴
$ aws cloudformation describe-stack-events --stack-name myteststack

# リソースのリスト表示
$ aws cloudformation list-stack-resources --stack-name myteststack

# テンプレートの取得
$ aws cloudformation get-template --stack-name myteststack

# テンプレートの検証
$ aws cloudformation validate-template \
    --template-url https://s3.amazonaws.com/cloudformation-templates-us-east-1/S3_Bucket.template

# 変更セットの作成と実行
$ aws cloudformation deploy --template /path_to_template/my-template.json \
    --stack-name my-new-stack \
    --parameter-overrides Key1=Value1 Key2=Value2

# スタックの削除
$ aws cloudformation delete-stack --stack-name myteststack
```


### スタックの更新

[AWS CloudFormation スタックの更新](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks.html)

２つの方式がある。

* 直接更新
* 変更セットを用いた更新


[スタックのリソースの更新動作](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html)

更新動作は次の３つのいずれか。

* 中断を伴わない更新
* 一時的な中断を伴う更新(AWS::EC2::Instance など)
* 置換(物理 ID も違うものになる。AWS::RDS::DBInstance リソースタイプの Engine プロパティなど)


[スタックのリソースが更新されないようにする](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/protect-stack-resources.html)

**スタックポリシー** により、スタックの更新、削除から保護できるようにポリシーを設定可能。
デフォルトではすべてのリソースが保護される。


[更新のロールバックを続ける](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-continueupdaterollback.html)

ほとんどの場合、スタックのロールバックを続けるには、更新のロールバックが失敗する原因となるエラーを修正する必要がある。



### ドリフト

[スタックとリソースに対するアンマネージド型設定変更の検出](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/using-cfn-stack-drift.html)

ドリフトの検出により、リソースの現在の設定とテンプレートとの差分を検出可能。
なお、ドリフトの検出は一部のリソースでのみサポートされている。



### リソースのインポート

[既存リソースの CloudFormation 管理への取り込み](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/resource-import.html)

既存のリソースを AWS CloudFormation 管理に取り込むことが可能。
なお、リソースのインポートは一部のリソースでのみサポートされている。

マネジメントコンソールでは [Stack Actions] → [Import resources into stack] で対応可能。
テンプレート側では "DeletionPolicy": "Retain" と指定しておく。


[スタック間でのリソースの移動](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/refactor-stacks.html)

次の流れでリソース移動が可能。

1. 削除対象のリソースに「"DeletionPolicy": "Retain"」を設定しスタックを更新
1. 削除対象のリソースに関する記述を削除
1. 移動先スタックでインポート



### スタックの出力値のエクスポート

[スタックの出力値のエクスポート](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/using-cfn-stack-exports.html)

スタックの出力値をエクスポートすることにより、他のスタックからインポートできるようになる。
テンプレートの Output セクションの Export フィールドを使用する。
インポートするには Fn::ImportValue 関数を使用する。


[エクスポートされた出力値をインポートするスタックのリスト](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/using-cfn-stack-imports.html)

エクスポートされた値を削除するには、インポートされていない状態を作る必要がある。
ListImports の API により、どのスタックでエクスポートされた値を使用しているかを一覧表示できる。



### スタックのネスト

[ネストされたスタックの操作](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/using-cfn-nested-stacks.html)

* AWS::CloudFormation::Stack により子のスタックを作成。
* スタックの更新時はルートスタックから起動する必要がある。



## テンプレート

### テンプレートの書式

[テンプレートの分析](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/template-anatomy.html)

以下の形式となっている。

```yaml
AWSTemplateFormatVersion: "version date"

Description:
  String

Metadata:
  template metadata

Parameters:
  set of parameters

Mappings:
  set of mappings

Conditions:
  set of conditions

Transform:
  set of transforms

Resources:
  set of resources

Outputs:
  set of outputs
```


#### Metadata

[メタデータ](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html)

[AWS::CloudFormation::Authentication](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/aws-resource-authentication.html)

AWS::CloudFormation::Init リソースで指定したファイルまたはソースの認証情報を指定。

[AWS::CloudFormation::Init](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/aws-resource-init.html)

cfn-init ヘルパースクリプト用のメタデータを Amazon EC2 インスタンスに取り込む。
要は EC2 インスタンスの構成を定義できる。
config は複数作成することも可能。Configset を用いることで、どの順番で config を呼び出すかを指定できる。

```yaml
Resources: 
  MyInstance: 
    Type: AWS::EC2::Instance
    Metadata: 
      AWS::CloudFormation::Init: 
        config: 
          packages: 
            :
          groups: 
            :
          users: 
            :
          sources: 
            :
          files: 
            :
          commands: 
            :
          services: 
            :
    Properties: 
      :
```

[AWS::CloudFormation::Interface](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-interface.html)

パラメータは通常アルファベット順となる。
これを設定することで、パラメーターのグループ化とソートを独自に定義することが可能。


#### Parameters

テンプレートで記述した内容に従って、パラメータをスタック作成時に設定できる。
設定したパラメータは Ref 関数により取得可能。

パラメータ例。

```yaml
Parameters: 
  InstanceTypeParameter: 
    Type: String
    Default: t2.micro
    AllowedValues: 
      - t2.micro
      - m1.small
      - m1.large
    Description: Enter t2.micro, m1.small, or m1.large. Default is t2.micro.
```

上記例にないもの。

* AllowedPattern: String 型の場合に使用できるパターン。
* ConstraintDescription: 制約違反の場合に表示するメッセージ。
* MaxLength, MinLength: String 型の文字数。
* MaxValue, MinValue: Number 型の値の範囲。
* NoEcho: 値をマスク。

Type は以下の種類がある。

* String
* Number
* List<Number>
* CommaDelimitedList
* [AWS 固有のパラメータ型](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/parameters-section-structure.html#aws-specific-parameter-types)。(AWS::EC2::VPC::Id、AWS::EC2::Subnet::Id、AWS::EC2::SecurityGroup::GroupName、など)
* SSM パラメータタイプ


[動的な参照を使用してテンプレート値を指定する](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/dynamic-references.html)

以下をサポート。

* AWS Systems Manager パラメータストア。プレーンテキスト値の ssm
* AWS Systems Manager パラメータストア。安全な文字列の ssm-secure
* AWS Secrets Manager に保存されている値。

例。
```yaml
  MyRDSInstance:
    Type: 'AWS::RDS::DBInstance'
    Properties:
      DBName: MyRDSInstance
      AllocatedStorage: '20'
      DBInstanceClass: db.t2.micro
      Engine: mysql
      MasterUsername: '{{resolve:secretsmanager:MyRDSSecret:SecretString:username}}'
      MasterUserPassword: '{{resolve:secretsmanager:MyRDSSecret:SecretString:password}}'
```


#### Mappings

[Mappings](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/mappings-section-structure.html)

値の取得には Fn::FindInMap を使用。

例。

```yaml
AWSTemplateFormatVersion: "2010-09-09"
Mappings: 
  RegionMap: 
    us-east-1:
      HVM64: ami-0ff8a91507f77f867
      HVMG2: ami-0a584ac55a7631c0c
    ap-northeast-1:
      HVM64: ami-06cd52961ce9f0d85
      HVMG2: ami-053cdd503598e4a9d
Resources: 
  myEC2Instance: 
    Type: "AWS::EC2::Instance"
    Properties: 
      ImageId: !FindInMap [RegionMap, !Ref "AWS::Region", HVM64]
      InstanceType: m1.small
```


#### Conditions

条件が true の場合のみ、リソースを作成できるようにできる。

例。
```yaml
AWSTemplateFormatVersion: "2010-09-09"

Parameters: 
  EnvType: 
    ...
    AllowedValues: 
      - prod
      - test

Conditions: 
  CreateProdResources: !Equals [ !Ref EnvType, prod ]

Resources: 
  NewVolume: 
    Type: "AWS::EC2::Volume"
    Condition: CreateProdResources
    ...
```


#### Transform

マクロを指定できる。


#### Resources

必須。作成するリソースを指定する。

```yaml
Resources:
  MyEC2Instance:
    Type: "AWS::EC2::Instance"
    Properties:
      ImageId: "ami-0ff8a91507f77f867"
```


#### Outputs

次の用途で使用可能。

* 値のエクスポート
* 応答として返す
* コンソールで値を表示する

例

```yaml
Outputs:
  InstanceID:
    Description: The Instance ID
    Value: !Ref EC2Instance
```

エクスポート例

```yaml
Outputs:
  StackVPC:
    Description: The ID of the VPC
    Value: !Ref MyVPC
    Export:
      Name: !Sub "${AWS::StackName}-VPCID"
```



### チュートリアル

[チュートリアル](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/CHAP_Using.html)



### テンプレートスニペット

[テンプレートスニペット](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/CHAP_TemplateQuickRef.html)

多数のサンプルシナリオがある。



### カスタムリソース

[カスタムリソース](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/template-custom-resources.html)

[AWS Lambda-backed カスタムリソース](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/template-custom-resources-lambda.html)

カスタムリソースと Lambda 関数を関連付けると、Lambda 関数はカスタムリソースが作成、更新、削除されるたびに実行される。
AMIM ID の動的な検索など、Lambda 関数を使用して様々なシナリオを実現できる。



### マクロ

[AWS CloudFormation マクロを使用したテンプレートのカスタム処理の実行](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/template-macros.html)

* AWS::Include transform: テンプレートスニペットをテンプレートに挿入可能。テンプレートの共通化用途で使用できる。
* AWS::Serverless: SAM で記述されたテンプレートを変換。

マクロを自作することもできる。



## StackSets

[AWS CloudFormation StackSets の操作](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/what-is-cfnstacksets.html)

複数のアカウント、リージョンのリソースを作成可能。



## テンプレートリファレンス

[AWS リソースおよびプロパティタイプのリファレンス](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)

全 AWS リソースについての Resources のリファレンス。


[仕様の形式](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/cfn-resource-specification-format.html)

Property の仕様。

```json
"Property name": {
  "Documentation": "Link to the relevant documentation",
  "DuplicatesAllowed": "true or false",
  "ItemType": "Type of list or map (non-primitive)",
  "PrimitiveItemType": "Type of list or map (primitive)",
  "PrimitiveType": "Type of value (primitive)",
  "Required": "true or false",
  "Type": "Type of value (non-primitive)",
  "UpdateType": "Mutable, Immutable, or Conditional",
}
```



### リソース属性

[リソース属性リファレンス](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/aws-product-attribute-reference.html)


#### CreationPolicy

[CreationPolicy 属性](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/aws-attribute-creationpolicy.html)

指定数の成功シグナルを受信するかまたはタイムアウト期間が超過するまでは、ステータスが作成完了にならないようにする。リソースにシグナルを送信するには、cfn-signal ヘルパースクリプトまたは SignalResource API を使用。

対応しているリソースは以下のもののみ。

* AWS::AutoScaling::AutoScalingGroup
* AWS::EC2::Instance
* AWS::CloudFormation::WaitCondition 


#### DeletionPolicy

[DeletionPolicy 属性](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html)

リソース削除時の挙動を設定。
Delete, Retain, Snapshot の３つのオプションがある。

デフォルトでは Delete。例外は AWS::RDS::DBCluster リソース、DBClusterIdentifier プロパティを指定しない AWS::RDS::DBInstance リソースでデフォルトポリシーは Snapshot。


#### DependsOn

[DependsOn 属性](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html)

リソース作成順の依存関係を設定できる。
無指定時は AWS CloudFormation は、可能な限り並行してリソースを作成、更新、削除する動作となるが、それでは困る場合に設定。


#### UpdatePolicy

[UpdatePolicy 属性](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html)

以下のリソースについて更新する方法を設定できる。

* AWS::AutoScaling::AutoScalingGroup
* AWS::ElastiCache::ReplicationGroup
* AWS::Elasticsearch::Domain
* AWS::Lambda::Alias

例えば AutoScaling グループの場合は以下のような設定項目がある。

```yaml
UpdatePolicy:
  AutoScalingRollingUpdate:
    # 同時に更新するインスタンスの最大数。
    MaxBatchSize: Integer
    # 最低限、維持する台数。
    MinInstancesInService: Integer
    # 成功判定に使用する、成功のシグナルを送信するインスタンスの割合。
    MinSuccessfulInstancesPercent: Integer
    # CloudFormation が一時停止する時間の長さ。この時間を超過すると更新は失敗する。
    PauseTime: String
    # 指定したプロセスについて Auto Scaling グループのスケールを停止する。
    SuspendProcesses:
      - List of processes
    # 新しいインスタンスからのシグナルを待機するかを設定。
    WaitOnResourceSignals: Boolean
```    


#### UpdateReplacePolicy

[UpdateReplacePolicy 属性](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/aws-attribute-updatereplacepolicy.html)

スタックの更新により置換が発生する場合、スナップショットを残す、元のリソースを保持する、など設定できる。


### 組み込み関数

[組み込み関数リファレンス](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference.html)

以下のような組み込み関数がある。

* Fn::Base64 - Base64 エンコーディングする。ユーザーデータの指定時などに使用。
* Fn::Cidr
* 条件関数 - Fn::If など。
* Fn::FindInMap - Mapping のキーに対応する値を返す。
* Fn::GetAtt - 指定した属性値を返す。
* Fn::GetAZs
* Fn::ImportValue - 別のスタックによってエクスポートされた値を返す。
* Fn::Join - 一連の値を特定の区切り文字で区切って 1 つの値に連結。
* Fn::Select - リストから指定したインデックスの値を取得。
* Fn::Split - 区切り記号を指定し、文字列を文字列値のリストに分割。
* Fn::Sub - 変数値の展開
* Fn::Transform
* Ref - 指定したパラメータまたはリソースの値を返す。


### ヘルパースクリプト

[CloudFormation ヘルパースクリプトリファレンス](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/cfn-helper-scripts-reference.html)

* cfn-init: リソースメタデータの取得と解釈、パッケージのインストール、ファイルの作成、およびサービスの開始で使用。
* cfn-signal: CreationPolicy または WaitCondition でシグナルを送信するために使用。
* cfn-get-metadata: メタデータの取得に使用
* cfn-hup: 変更が検出されたときにカスタムフックを実行するために使用。

ヘルパースクリプトは /opt/aws/bin にインストールされている。

以下は [cfn-signal](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/cfn-signal.html) に記載されている例。

Metadata でどのような構成にするかを指定し、UserData で cfn-init, cfn-signal を順に実行している。

```yaml
AWSTemplateFormatVersion: 2010-09-09
Description: Simple EC2 instance
Resources:
  MyInstance:
    Type: AWS::EC2::Instance
    Metadata:
      'AWS::CloudFormation::Init':
        config:
          files:
            /tmp/test.txt:
              content: Hello world!
              mode: '000755'
              owner: root
              group: root
    Properties:
      ImageId: ami-a4c7edb2
      InstanceType: t2.micro
      UserData: !Base64
        'Fn::Join':
          - ''
          - - |
              #!/bin/bash -x
            - |
              # Install the files and packages from the metadata
            - '/opt/aws/bin/cfn-init -v '
            - '         --stack '
            - !Ref 'AWS::StackName'
            - '         --resource MyInstance '
            - '         --region '
            - !Ref 'AWS::Region'
            - |+

            - |
              # Signal the status from cfn-init
            - '/opt/aws/bin/cfn-signal -e $? '
            - '         --stack '
            - !Ref 'AWS::StackName'
            - '         --resource MyInstance '
            - '         --region '
            - !Ref 'AWS::Region'
            - |+

    CreationPolicy:
      ResourceSignal:
        Timeout: PT5M
```        



## サンプルテンプレート

[サンプルテンプレート](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/sample-templates-services-ap-northeast-1.html)



## Black Belt

[20200826 AWS Black Belt Online Seminar AWS CloudFormation](https://www.slideshare.net/AmazonWebServicesJapan/20200826-aws-black-belt-online-seminar-aws-cloudformation-238501102)

* テンプレート
  * P26: テンプレートの要素
  * P36: Resources
  * P39: 組み込み関数、疑似パラメータ
* 開発
  * P51: cfn-lint。入力した値の整合性をチェックしてくれるツール。
* テスト
* デプロイ
  * P65: CodePipeline によるデプロイが可能。
  * P66: Change Set で変更内容を事前確認。
  * P67: StackSets による複数アカウント、リージョンへのデプロイ。
* 運用
  * P73: スタック更新の流れ。ドリフトの確認。Change Set による変更内容の確認。
  * P84: ライフサイクルに応じてテンプレートを分割。
  * P93: ヘルパースクリプト。現時点では cfn-init より State Manager の利用が推奨されるらしい。
  * P95: Dynamic Reference。SSM、Secrets Manager の値を参照したいときに使用。


[20201006 AWS Black Belt Online Seminar AWS CloudFormation deep dive](https://www.slideshare.net/AmazonWebServicesJapan/20201006-aws-black-belt-online-seminar-aws-cloudformation-deep-dive)

* P9: StackSets
* P20: リソースのインポート
* P24: テンプレートのリファクタリング
* P27: Ansible と組み合わせたサーバの構成管理
* P32: 一つのテンプレートで複数の環境に対応
* P36: カスタムリソース
* P38: マクロ
* P40: スタックの作成権限とリソースの保護
* P42: CDK と CLoudFormation の使い分け
* P44: CodePipeline からスタックを作成
* P47: CloudFormation による Blue/Green Deployment
* P49: リソースプロバイダを使用した独自リソースの作成



# 参考

* Document
  * [AWS CloudFormation とは](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/Welcome.html)
* サービス紹介ページ
  * [AWS CloudFormation](https://aws.amazon.com/jp/cloudformation/)
  * [よくある質問](https://aws.amazon.com/jp/cloudformation/faqs/)
* Black Belt
  * [20200826 AWS Black Belt Online Seminar AWS CloudFormation](https://www.slideshare.net/AmazonWebServicesJapan/20200826-aws-black-belt-online-seminar-aws-cloudformation-238501102)
  * [20201006 AWS Black Belt Online Seminar AWS CloudFormation deep dive](https://www.slideshare.net/AmazonWebServicesJapan/20201006-aws-black-belt-online-seminar-aws-cloudformation-deep-dive)

