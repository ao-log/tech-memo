
[TypeScript for JavaScript Programmers](https://www.typescriptlang.org/docs/handbook/typescript-in-5-minutes.html)


## Defining Types

型推論を行うような記述が可能。

```ts
const user = {
  name: "Hayes",
  id: 0,
};
```

interface によりオブジェクトの形状を記述できる。

```ts
interface User {
  name: string;
  id: number;
}
```

interface に準拠していることを示す書き方。「: Type名」の記法。

```ts
const user: User = {
  name: "Hayes",
  id: 0,
};
```

class で interface を使用できる。

```ts
interface User {
  name: string;
  id: number;
}

class UserAccount {
  name: string;
  id: number;

  constructor(name: string, id: number) {
    this.name = name;
    this.id = id;
  }
}

const user: User = new UserAccount("Murphy", 1);
```

戻り値、引数の定義にもインタフェースを使用できる。

```ts
function getAdminUser(): User {
  //...
}

function deleteUser(user: User) {
  // ...
}
```


## Composing Types

型を構成する方法は Unions, Generics が主流な方法。

#### Unions

```ts
type MyBool = true | false;
```

次のように関数の引数、戻り値をかける。
```ts
function getLength(obj: string | string[]) {
  return obj.length;
}
```

#### Generics

Generics を使用することで、実際に利用されるまで型が確定しないシーンに対応することができる。

```ts
interface Backpack<Type> {
  add: (obj: Type) => void;
  get: () => Type;
}

// ここでは Type を string として扱う。
declare const backpack: Backpack<string>;

// この場合、object は string となる。
const object = backpack.get();

// backpack は string なので数値を入れることはできない。
backpack.add(23);
```


## Structural Type System

```rs
interface Point {
  x: number;
  y: number;
}

function logPoint(p: Point) {
  console.log(`${p.x}, ${p.y}`);
}

// logs "12, 26"
const point = { x: 12, y: 26 };
logPoint(point);
```

インタフェース Point と point は同じ構造を持っている。
この場合は同じタイプをみなされる。
ダックタイピングとも呼ばれる。

