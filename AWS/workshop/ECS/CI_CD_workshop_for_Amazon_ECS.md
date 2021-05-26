
[CI/CD WORKSHOP FOR AMAZON ECS](https://cicd-for-ecs.workshop.aws/en/)


[TOOLS AND RESOURCES](https://cicd-for-ecs.workshop.aws/en/3-setup/01-tools_resources.html)

タスク実行ロールの作成。サービスロール「AmazonECSTaskExecutionRolePolicy」をアタッチ。

Container Insights の有効化。

```
aws ecs put-account-setting-default --name containerInsights --value enabled
```



## [LAB 1 : ROLLING UPDATE](https://cicd-for-ecs.workshop.aws/en/4-basic/lab1-rolling.html)

[CREATE REPOS](https://cicd-for-ecs.workshop.aws/en/4-basic/lab1-rolling/11-repos.html)

* ECR リポジトリの作成
* CodeCommit リポジトリの作成
* CodeCommit リポジトリをクローンし、ローカル PC 上でソース、Dockerfile、buildspec.yml を配備


[CREATE ECS RESOURCES](https://cicd-for-ecs.workshop.aws/en/4-basic/lab1-rolling/12-ecsresources.html)

* CloudWatch Logs のロググループ作成
* タスク定義の登録

```
$ aws ecs register-task-definition --cli-input-json file://task-definition.json
```

```json
{
  "family": "hello-dude",
  "executionRoleArn": "${TASK_EXEC_ROLE_ARN}",
  "networkMode": "bridge",
  "cpu": "256",
  "memory": "512",
  "requiresCompatibilities": [
    "EC2"
  ],
  "containerDefinitions": [
    {
      "name": "hello-dude",
      "image": "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/hello-dude",
      "logConfiguration": {
        "logDriver": "awslogs",
        "options": {
          "awslogs-group": "/ecs/hello-dude",
          "awslogs-region": "${AWS_REGION}",
          "awslogs-stream-prefix": "ecs"
        }
      },
      "portMappings": [
        {
          "containerPort": 80,
          "protocol": "tcp"
        }
      ],
      "essential": true
    }
  ]
}
```

* ALB の作成
* ECS サービスの作成

```
$ aws ecs create-service --cli-input-json file://staging_service.json
```

```json
{
    "serviceName": "hello-dude", 
    "launchType": "EC2", 
    "loadBalancers": [
        {
            "targetGroupArn": "${STAGING_TARGET_GROUP_ARN}", 
            "containerName": "hello-dude",
            "containerPort": 80
        }
    ], 
    "desiredCount": 2, 
    "cluster": "arn:aws:ecs:${AWS_REGION}:${AWS_ACCOUNT_ID}:cluster/staging", 
    "serviceName": "hello-dude", 
    "deploymentConfiguration": {
        "maximumPercent": 200, 
        "minimumHealthyPercent": 50
    }, 
    "healthCheckGracePeriodSeconds": 2, 
    "schedulingStrategy": "REPLICA", 
    "taskDefinition": "arn:aws:ecs:${AWS_REGION}:${AWS_ACCOUNT_ID}:task-definition/hello-dude:1"
}
```


[CREATE THE PIPELINE](https://cicd-for-ecs.workshop.aws/en/4-basic/lab1-rolling/13-pipeline.html)

* CodePipeline でパイプライン作成
  * Source ステージは CodeCommit
  * Build ステージをここで作成。
  * Deploy ステージ


CodeBuild で使用する buildspec.yaml は以下内容

```yaml
version: 0.2

phases:
  install:
    runtime-versions:
      docker: 18
  pre_build:
    commands:
      - echo Logging in to Amazon ECR...
      - aws --version
      - $(aws ecr get-login --region $AWS_DEFAULT_REGION --no-include-email)
      - REPOSITORY_URI=<IMAGE_REPO_URI>
      - COMMIT_HASH=$(echo $CODEBUILD_RESOLVED_SOURCE_VERSION | cut -c 1-7)
      - IMAGE_TAG=${COMMIT_HASH:=latest}
  build:
    commands:
      - echo Build started on `date`
      - echo Building the Docker image...
      - docker build -t $REPOSITORY_URI:latest .
      - docker tag $REPOSITORY_URI:latest $REPOSITORY_URI:$IMAGE_TAG
  post_build:
    commands:
      - echo Build completed on `date`
      - echo Pushing the Docker images...
      - docker push $REPOSITORY_URI:latest
      - docker push $REPOSITORY_URI:$IMAGE_TAG
      - echo Writing image definitions file...
      - printf '[{"name":"hello-dude","imageUri":"%s"}]' $REPOSITORY_URI:$IMAGE_TAG > imagedefinitions.json
artifacts:
    files: imagedefinitions.json
```


[DEPLOY A CHANGE](https://cicd-for-ecs.workshop.aws/en/4-basic/lab1-rolling/15-deploy-change.html)

ソースをプッシュすると自動的にパイプラインが開始される。



## [LAB 2 : BLUE GREEN DEPLOYMENT](https://cicd-for-ecs.workshop.aws/en/4-basic/lab2-bluegreen.html)

[CREATE ECS SERVICE](https://cicd-for-ecs.workshop.aws/en/4-basic/lab2-bluegreen/11-ecsresources.html)

* Blue、Green 用の二つのターゲットグループを作成
* CloudWatch Logs のロググループを作成
* タスク定義を登録
* サービスを作成


[CREATE CODEDEPLOY APPLICATION](https://cicd-for-ecs.workshop.aws/en/4-basic/lab2-bluegreen/12-codedeployapp.html)

* タスク定義と appspec.yaml をソースに含める。

appspec.yaml は以下内容とする。<TASK_DEFINITION> はプレースホルダーで CodePipeline によって上書きされる。

```json
version: 0.0
Resources:
  - TargetService:
      Type: AWS::ECS::Service
      Properties:
        TaskDefinition: <TASK_DEFINITION>
        LoadBalancerInfo:
          ContainerName: "hello-dude"
          ContainerPort: 80
```

* CodeDeploy のアプリケーション、デプロイグループを作成。


[EXTEND PIPELINE](https://cicd-for-ecs.workshop.aws/en/4-basic/lab2-bluegreen/13-pipeline.html)

buildspec.yaml では以下のように imageDetail.json を出力し、出力アーティファクトに含める必要がある。

```
phases:
  post_build:
    ...
    - printf '{"ImageURI":"%s"}' $REPOSITORY_URI:$IMAGE_TAG > imageDetail.json
...
artifacts:
  files: 
#    - imagedefinitions.json
    - imageDetail.json
    - appspec.yaml
    - taskdef-prod.json
```

* CodePipeline の設定
  * Rolling Update 時に作成した Deploy ステージの削除
  * 承認ステージの追加
  * Deploy ステージの作成



## [LAB 3 : EXPLORING CODEDEPLOY](https://cicd-for-ecs.workshop.aws/en/4-basic/lab3-codedeploy.html)

以下について書かれている。

* using a test port for task replacement set validation,
* rolling back to previous version,
* using canary and linear deployments.


[SET UP TEST PORT](https://cicd-for-ecs.workshop.aws/en/4-basic/lab3-codedeploy/1-testport.html)

* テストポートを設定している。
* Specify when to reroute traffic. に 1 時間を設定している。


[PUSH CHANGE, TEST, AND ROLLBACK](https://cicd-for-ecs.workshop.aws/en/4-basic/lab3-codedeploy/2-change.html)

新しいタスクセットが起動し、「Reroute traffic」を押すまでの間、テストポートが新しいタスクセットに向いている。
「curl $PROD_ALB_URL:8080/hello/bob」で動作を確認できる。


[CANARY DEPLOYMENT](https://cicd-for-ecs.workshop.aws/en/4-basic/lab3-codedeploy/3-canary.html)

デプロイ設定で「CodeDeployDefault.ECSCanary10Percent5Minutes.」を選ぶ。
「Reroute traffic」を押すことでカナリアデプロイが実行され、トラフィックの 10 % が Green 側を向くようになる。
5 分間経過すると Green 側に 100 % 向くようになる。それまでに「Stop and roll back」でロールバックさせることもできる。



## [LAB 4: GITOPS](https://cicd-for-ecs.workshop.aws/en/5-advanced/lab4-gitops.html)

* front, middle, back の 3 段構成。
* front は ALB 下に。
* front から middle はサービスディスカバリにより http://middle.lab4/middle でアクセス。
* middle から back はサービスディスカバリにより http://back.lab4/back でアクセス。

[CREATE APPLICATION PIPELINE](https://cicd-for-ecs.workshop.aws/en/5-advanced/lab4-gitops/01-app-pipleline.html)

CloudFormation で Source, Build 構成のパイプラインを作成している。
以下のリソースが作られる。

* Type: AWS::CodeCommit::Repository
* Type: AWS::IAM::Role ← CodeCommit 更新時にパイプラインを発行する際に使用
* Type: AWS::Events::Rule
* Type: AWS::ECR::Repository
* Type: AWS::IAM::Role ← CodeBuild 用
* Type: AWS::IAM::Role ← CodePipeline 用
* Type: AWS::S3::Bucket
* Type: AWS::CodeBuild::Project
* Type: AWS::CodePipeline::Pipeline


[2. Create private DNS namespace in AWS CloudMap](https://cicd-for-ecs.workshop.aws/en/5-advanced/lab4-gitops/02-namespace.html)

namespace を作成する。

```
aws servicediscovery create-private-dns-namespace --name lab4 --vpc $PROD_VPC
```


[3. Create service pipelines](https://cicd-for-ecs.workshop.aws/en/5-advanced/lab4-gitops/03-service-piplelines.html)

CloudFormation で Front, Middle, Back それぞれの Source, Deploy(CFn) 構成のパイプラインを作成している。
それぞれ以下のリソースが作られる。

* Type: AWS::CodeCommit::Repository
* Type: AWS::IAM::Role
* Type: AWS::Events::Rule
* Type: AWS::IAM::Role
* Type: AWS::IAM::Role
* Type: AWS::S3::Bucket
* Type: AWS::CodePipeline::Pipeline

パイプラインでデプロイする CFn では以下のリソースが作られる。front は更に ALB 関連リソースを作成。

* Type: "AWS::Logs::LogGroup"
* Type: AWS::ServiceDiscovery::Service
* Type: AWS::ECS::Service
* Type: AWS::ECS::TaskDefinition

