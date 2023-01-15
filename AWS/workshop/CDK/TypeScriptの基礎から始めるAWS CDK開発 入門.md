
[TypeScriptの基礎から始めるAWS CDK開発 入門](https://catalog.workshops.aws/typescript-and-cdk-for-beginner/ja-JP)


## TypeScript の基礎

[TypeScriptの概要](https://catalog.workshops.aws/typescript-and-cdk-for-beginner/ja-JP/typescript-basic/overview)

* TypeScript は、トランスパイル時に型検査を行う。トランスパイラによって javaScript に変換される。


[Cloud9のTypeScript実行環境](https://catalog.workshops.aws/typescript-and-cdk-for-beginner/ja-JP/typescript-basic/typescript-version)

* tsc コマンドは TypeScript Compiler
* Node.js は JavaScript のランタイム


[プロジェクトの作成](https://catalog.workshops.aws/typescript-and-cdk-for-beginner/ja-JP/typescript-basic/initial-project)

* ```npm init -y``` により ```package.json``` が生成される
* ```npm install @types/node``` によりカレントディレクトリ下に ```node_modules/@types/node``` ディレクトリが生成される
* ```tsc --init``` により ```tsconfig.json``` が生成される。このファイルにトランスパイルされたファイルの生成先を記述する


[変数](https://catalog.workshops.aws/typescript-and-cdk-for-beginner/ja-JP/typescript-basic/variables)

```ts
# let <変数名>: <型> = value; の形式で定義
let smallNumber: number = 1 ;
let smallNumber = 1 ; //型推論もできる

# 定数は const
const red: string = "red"

# var による指定も可能だが極力使わないようにする
```


[TypeScriptと型エイリアス](https://catalog.workshops.aws/typescript-and-cdk-for-beginner/ja-JP/typescript-basic/data-types/types)

以下のようにデータ型を定義できる。

```ts
type Staff = {
  firstName: string;
  lastName: string;
};

const staff: Staff = { firstName: "AWS", lastName: "太郎" };
```


[関数](https://catalog.workshops.aws/typescript-and-cdk-for-beginner/ja-JP/typescript-basic/functions)

コード例。

```ts
function saySomething(text: string) {
  console.log(text);
}

saySomething("Hello!");
```


[クラス](https://catalog.workshops.aws/typescript-and-cdk-for-beginner/ja-JP/typescript-basic/class)

コード例。

```ts
class Person {
  //?を付け加えてエラーを回避
  firstName?: string;
  lastName?: string;
  // greet メソッドを追加
  greet() {
    console.log(`私は${this.firstName} ${this.lastName}です。`);
  }
}

const taro = new Person();

taro.firstName = "AWS";
taro.lastName = "太郎";

// greetメソッドを使用
taro.greet();
```


[コンストラクタ](https://catalog.workshops.aws/typescript-and-cdk-for-beginner/ja-JP/typescript-basic/class/constructor)

コンストラクタをしようすることで、クラスからオブジェクトを作成する際にプロパティを初期化できる。

コード例。

```ts
class Person {
  //インスタンス化される際にコンストラクタによって値が代入されるため、オプションであることを示す?記号は外すことができます。
  firstName: string;
  lastName: string;
  // コンストラクタを定義
  constructor(firstName: string, lastName: string) {
    this.firstName = firstName;
    this.lastName = lastName;
  }
  greet() {
    console.log(`私は${this.firstName} ${this.lastName}です。`);
  }
}

// クラスからオブジェクトを生成する際、引数を渡してプロパティを初期化しています。
const taro = new Person("AWS", "太郎");
//次郎も生成してみましょう。
const jiro = new Person("Amazon","次郎");

taro.greet();
jiro.greet();
```


[インタフェース](https://catalog.workshops.aws/typescript-and-cdk-for-beginner/ja-JP/typescript-basic/class/interface)

肩エイリアスと似ている。インタフェースは継承ができる。

コード例。

```ts
// 型エイリアスではなく、interfaceを使って型を定義しています。
interface Human {
  firstName: string;
  lastName: string;
}

const aws: Human = { firstName: "AWS", lastName: "太郎" };
const amazon: Human = { firstName: "Amazon", lastName: "次郎" };

console.log(`${aws.firstName} ${aws.lastName}の上司は${amazon.firstName} ${amazon.lastName}です。`);
```


## AWS CDK 入門

作業の流れ。

```
// 初回のみ必要
cdk bootstrap

// CDK プロジェクトを作成
cdk init sample-app --language typescript

// CloudFormation テンプレートを生成
cdk synth

// CloudFormation スタックを作成
cdk deploy
```


・bin/cdk-workshop.ts
```ts
#!/usr/bin/env node
import * as cdk from 'aws-cdk-lib';
import { CdkWorkshopStack } from '../lib/cdk-workshop-stack';

const app = new cdk.App();
new CdkWorkshopStack(app, 'CdkWorkshopStack');
```
* lib/cdk-workshop-stack.ts の CdkWorkshopStack クラスをロードしてインスタンス化している。

・lib/cdk-workshop-stack.ts 
```ts
import { Duration, Stack, StackProps } from 'aws-cdk-lib';
import * as sns from 'aws-cdk-lib/aws-sns';
import * as subs from 'aws-cdk-lib/aws-sns-subscriptions';
import * as sqs from 'aws-cdk-lib/aws-sqs';
import { Construct } from 'constructs';

export class CdkWorkshopStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);

    const queue = new sqs.Queue(this, 'CdkWorkshopQueue', {
      visibilityTimeout: Duration.seconds(300)
    });

    const topic = new sns.Topic(this, 'CdkWorkshopTopic');

    topic.addSubscription(new subs.SqsSubscription(queue));
  }
}
```
* SQS のキューとトピックを作成している。
* export することで bin/cdk-workshop.ts から import できるようになる。


[WordPressのサイトをCDKを使って公開する](https://catalog.workshops.aws/typescript-and-cdk-for-beginner/ja-JP/cdk-introduction/wordpress)

[API Reference](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-construct-library.html) を見つつコードを書く。Overview のページには実装例があるので参考になる。




