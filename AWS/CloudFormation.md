
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

Ref 関数を使用することにより、戻り値を取得可能。
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

Fn::GetAtt により他のリソースの属性値を取得可能。利用可能な属性値は [AWS リソースおよびプロパティタイプのリファレンス](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html) より参照可能。

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

Outputs によりスタックの出力値を設定可能。
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

## ベストプラクティス

[AWS CloudFormation ベストプラクティス](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/best-practices.html)

* スタックの分割は、例えばチームの責務の範囲といった分け方をする。
* 共有リソースはエクスポートし、利用するスタック側からはクロススタック参照を行う。
* 本番、開発環境も一つのテンプレートで対応できる。パラメータ、マッピング、条件セクションを利用する。
* ネストされたスタックを使用することで、共通リソースの更新時に他のスタックも追随できるようにする。
* テンプレートには認証情報を埋め込まない。
* パラメータには成約を設定できるので、無効な値の入力を防ぐことができる。
* AWS::CloudFormation::Init を使用することにより、EC2 インスタンスの起動時に任意の設定を行うことができる。
* aws cloudformation validate-template により検証を行うことができる。
* CloudFormation で作成したリソースは、その他の方法で変更しないようにする。
* スタック更新時には変更セットを使用する。
* 重要なリソースがあるスタックは、スタックポリシーで保護する。
* テンプレートのソースをレポジトリで管理し、リビジョンごとの差分をトレースし、戻せるようにする。

## 継続的デリバリー

[CodePipeline を使用した継続的デリバリー](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/continuous-delivery-codepipeline.html)

レポジトリへのテンプレートのコミット → テストスタックの作成 → 本番スタックの作成のようなパイプラインを作成可能。

## スタックの作成

[AWS CloudFormation コンソールでのスタックの作成](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/cfn-console-create-stack.html)

以下の項目を設定可能。
* 使用するサービスロール
* スタックポリシー
* ロールバック設定（指定したいずれかのアラームのしきい値を超えた場合にロールバック可能）
* 失敗時のロールバックの有無。デフォルトは有効。
* タイムアウト値。
* スタックの削除からの保護。

変更セットを使用して、新規にスタックを作成することも可能。

## スタックの表示

[AWS マネジメントコンソール での AWS CloudFormation スタックデータとリソースの表示](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/cfn-console-view-stack-data-resources.html)

スタックの画面では、以下のタブから各情報を参照可能。

* スタック情報
* イベント
* リソース
* 出力
* パラメータ
* テンプレート

## AWS CLI

[AWS コマンドラインインターフェイスの使用](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/cfn-using-cli.html)

```shell
# スタックの作製
$ aws cloudformation create-stack --stack-name myteststack --template-body file:///home/testuser/mytemplate.json --parameters ParameterKey=Parm1,ParameterValue=test1 ParameterKey=Parm2,ParameterValue=test2

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
$ aws cloudformation validate-template --template-url https://s3.amazonaws.com/cloudformation-templates-us-east-1/S3_Bucket.template

# 変更が含まれるスタックを作成
$ aws cloudformation deploy --template /path_to_template/my-template.json --stack-name my-new-stack --parameter-overrides Key1=Value1 Key2=Value2

# スタックの削除
$ aws cloudformation delete-stack --stack-name myteststack
```

## スタックの更新

[AWS CloudFormation スタックの更新](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks.html)

直接更新と変更セットを用いた更新の２つの方式がある。

[スタックのリソースの更新動作](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html)

更新動作は次の３つのいずれか。
* 中断を伴わない更新
* 一時的な中断を伴う更新
* 置換

## スタックポリシー

[スタックのリソースが更新されないようにする](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/protect-stack-resources.html)

スタックの更新、削除から保護できるようにポリシーを設定可能。

## ドリフト

[スタックとリソースに対するアンマネージド型設定変更の検出](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/using-cfn-stack-drift.html)

ドリフトの検出により、リソースの現在の設定とテンプレートとの差分を検出可能。
なお、ドリフトの検出は一部のリソースでのみサポートされている。

## エクスポート

[スタックの出力値のエクスポート](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/using-cfn-stack-exports.html)

スタックの出力値をエクスポートすることにより、他のスタックからインポートできるようになる。

## スタックのネスト

[ネストされたスタックの操作](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/using-cfn-nested-stacks.html)

スタックの更新時はルートスタックから起動する必要がある。

## 既存リソースの CloudFormation 管理への取り込み

[既存リソースの CloudFormation 管理への取り込み](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/resource-import.html)

既存のリソースを AWS CloudFormation 管理に取り込むことが可能。リソースは削除してスタックの一部として再作成しなくても、作成された場所に関係なく AWS CloudFormation を使用して管理できる。

## テンプレート

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

### パラメータ

動的な参照は次のパターンに従う。

'{{resolve:service-name:reference-key}}'

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

### Condition

条件が true の場合のみ、リソースを作成できるようにできる。

例。
```yaml
AWSTemplateFormatVersion: "2010-09-09"
...
Parameters: 
  EnvType: 
    Description: Environment type.
    Default: test
    Type: String
    AllowedValues: 
      - prod
      - test
    ConstraintDescription: must specify prod or test.
Conditions: 
  CreateProdResources: !Equals [ !Ref EnvType, prod ]
Resources: 
  EC2Instance: 
    Type: "AWS::EC2::Instance"
    Properties: 
      ImageId: !FindInMap [RegionMap, !Ref "AWS::Region", AMI]
  MountPoint: 
    Type: "AWS::EC2::VolumeAttachment"
    Condition: CreateProdResources
    Properties: 
      InstanceId: 
        !Ref EC2Instance
      VolumeId: 
        !Ref NewVolume
      Device: /dev/sdh
  NewVolume: 
    Type: "AWS::EC2::Volume"
    Condition: CreateProdResources
    Properties: 
      Size: 100
      AvailabilityZone: 
        !GetAtt EC2Instance.AvailabilityZone
Outputs: 
  VolumeId: 
    Condition: CreateProdResources
    Value: 
      !Ref NewVolume
```

### Transform

マクロを指定できる。

## テンプレートスニペット

[テンプレートスニペット](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/CHAP_TemplateQuickRef.html)

多数のサンプルシナリオがある。

## カスタムリソース

サポートされていないリソースについても作成できる。

## StackSets

[AWS CloudFormation StackSets の操作](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/what-is-cfnstacksets.html)

複数のアカウント、リージョンのリソースを作成可能。

## テンプレートリファレンス

[テンプレートリファレンス](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/template-reference.html)


[DeletionPolicy 属性](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html)

スタックを削除したときの動作（削除、保持、スナップショット）を設定可能。

[DependsOn 属性](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html)

リソースの依存関係を設定可能。依存関係があり、特定の順序で作成が必要な場合などに使用。

[組み込み関数リファレンス](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference.html)

以下のような組み込み関数がある。

* Fn::Base64
* Fn::Cidr
* 条件関数
* Fn::FindInMap - Mapping のキーに対応する値を返す。
* Fn::GetAtt - 指定した属性値を返す。
* Fn::GetAZs
* Fn::ImportValue - 別のスタックによってエクスポートされた値を返す。
* Fn::Join
* Fn::Select
* Fn::Split
* Fn::Sub
* Fn::Transform
* Ref - 指定したパラメータまたはリソースの値を返す。

## サンプルテンプレート

[サンプルテンプレート](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/sample-templates-services-ap-northeast-1.html)


# 参考

* [AWS CloudFormation とは](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/Welcome.html)
* [AWS Black Belt Online Seminar AWS CloudFormation アップデート](https://www.slideshare.net/AmazonWebServicesJapan/aws-black-belt-online-seminar-aws-cloudformation)


