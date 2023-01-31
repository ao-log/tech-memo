
# Document

## Getting started

```
// Install the AWS CDK
npm install -g aws-cdk

// bootstraping
cdk bootstrap aws://ACCOUNT-NUMBER/REGION
```

## Your first AWS CDK app

TypeScript の場合
```
cdk init app --language typescript
```

lib/hello-cdk-stack.ts
```ts
import * as cdk from 'aws-cdk-lib';
import { aws_s3 as s3 } from 'aws-cdk-lib';

export class HelloCdkStack extends cdk.Stack {
  constructor(scope: cdk.App, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    new s3.Bucket(this, 'MyFirstBucket', {
      versioned: true
    });
  }
}
```

```
// テンプレート生成
cdk synth

// Deploy
cdk deploy
```

```
// 変更内容の確認
cdk diff
```


# 参考

* Document
  * [What is the AWS CDK?](https://docs.aws.amazon.com/cdk/v2/guide/home.html)
  * [API Reference](https://docs.aws.amazon.com/cdk/api/v1/docs/aws-construct-library.html)
* サービス紹介ページ
  * [AWS クラウド開発キット](https://aws.amazon.com/jp/cdk/)
  * [よくある質問](https://aws.amazon.com/jp/cdk/faqs/)
* Black Belt
  * [20200303 AWS Black Belt Online Seminar AWS Cloud Development Kit (CDK)](https://pages.awscloud.com/rs/112-TZM-766/images/20200303_BlackBelt_CDK.pdf)

