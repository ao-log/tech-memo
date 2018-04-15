
# リスト

リストについてのメモ。

# 参考

以下の URL、書籍を参考にしている。

##### Python チュートリアル

https://docs.python.jp/3/tutorial/

##### 書籍『Effective Python』

『Effective Python』のサンプルコードは次の URL で公開されている。
https://github.com/bslatkin/effectivepython/tree/master/example_code

### リストの長さ

```python
>>> s = 'supercalifragilisticexpialidocious'
>>> len(s)
34
```

### リストのコピー

```python
>>> a = list(range(5))

# コピーされない（参照のみ渡している）
>>> b = a
# コピーされる
>>> c = a[:]

# リスト a の内容を更新。
>>> a[1] = 'update'

# リスト a の内容になっている。
>>> print('b=%s' % b)
b=[0, 'update', 2, 3, 4]

# こちらは更新されていない。
>>> print('c=%s' % c)
c=[0, 1, 2, 3, 4]
```

### range

リストではない。イテラブルなオブジェクト。

```python
>>> print(range(5))
range(0, 5)

# リスト化するには、関数　list() を使う。
>>> print(list(range(5)))
[0, 1, 2, 3, 4]

# イテラブルなので、指定回数　 for 文を回す場合はリスト化しなくて良い。
>>> for i in range(5):
... print(i)
0
1
2
3
4
```

### リスト操作

```python
>>> fruits = ['orange', 'apple', 'pear', 'banana', 'kiwi', 'apple', 'banana']

# 要素があるか確認
>>> 'orange' in fruits
True

# リストの末尾に追加
>>> fruits.append('grape')

# リストの末尾の要素を取り出し
>>> fruits.pop()
'grape'

# ソート(非破壊的。ソートされたリストを返す)
>>> sorted(fruits)
['apple', 'apple', 'banana', 'banana', 'kiwi', 'orange', 'pear']

# ソート（破壊的。元のリストをソートする）
>>> fruits.sort()
>>> fruits
['apple', 'apple', 'banana', 'banana', 'kiwi', 'orange', 'pear']
```

ソートは、[ソート HOW TO](https://docs.python.jp/3/howto/sorting.html) も参照のこと。

### キュー

```python
>>> from collections import deque

>>> queue = deque(['one', 'two', 'three'])

# 末尾に追加
>>> queue.append('four')

# 先頭の要素を取り出し
>>> queue.popleft()
'one'

# 末尾の要素を取り出し
>>> queue.pop()
'four'
```

### リスト内包表記

```python
>>> print([x**2 for x in range(10)])
[0, 1, 4, 9, 16, 25, 36, 49, 64, 81]

>>> print([x**2 for x in range(10) if x %2 == 0])
[0, 4, 16, 36, 64]

>>> from math import pi
>>> [str(round(pi, i)) for i in range(1, 6)]
['3.1', '3.14', '3.142', '3.1416', '3.14159']
```

# タプル

格納した要素の値を更新することはできない。

# 集合型

```python
# 中括弧で囲む
>>> basket = {'apple', 'orange', 'apple', 'pear', 'orange', 'banana'}

>>> 'orange' in basket    
True

# 和集合
>>> basket_2 = {'lemon', 'orange'}
>>> print(basket & basket_2)
{'orange'}
```

# 辞書型

```python
>>> user = {'name': 'hoge', 'email': 'hoge@example.com'}

# 中身の取り出し
>>> for k, v in user.items():
...    print(k, v)
name hoge
email hoge@example.com

# 値の更新
>>> user['name'] = 'fuga'
>>> user
{'email': 'hoge@example.com', 'name': 'fuga'}
```

# その他

### enumerate

```python
>>> basket = {'apple', 'orange', 'apple', 'pear', 'orange', 'banana'}

>>> for i, v in enumerate(basket):
... print(i, v)
0 pear
1 banana
2 orange
3 apple

```

### defaultdict

https://docs.python.jp/3/library/collections.html#collections.defaultdict

初期化出来る辞書。値が格納されていない場合に初期化する処理を if 文で書くのは面倒だが、その手間を省けるしすっきり書ける。

```python
>>> from collections import defaultdict

>>> s = [('yellow', 1), ('blue', 2), ('yellow', 3), ('blue', 4), ('red', 1)]
>>> d = defaultdict(list)
>>> for k, v in s:
...    d[k].append(v)

>>> print(sorted(d.items()))
[('blue', [2, 4]), ('red', [1]), ('yellow', [1, 3])]
```

なお、辞書の初期化は次のようにも書ける。（が、Python のドキュメントによると defaultdict を使ったほうが速いらしい）

```python
>>> d = {}
>>> for k, v in s:
...     d.setdefault(k, []).append(v)
...
>>> sorted(d.items())
[('blue', [2, 4]), ('red', [1]), ('yellow', [1, 3])]
```

### OrderedDict

https://docs.python.jp/3/library/collections.html#collections.OrderedDict

OrderedDict を用いることで、辞書に格納した順番が保持される。

普通の辞書の場合。

```python
>>> a = {}
>>> a['foo'] = 1
>>> a['bar'] = 2
>>> a['hoge'] = 3
// 実行するたびに異なる順番になる。
>>> print(a)
{'bar': 2, 'hoge': 3, 'foo': 1}
```

OrderedDict の場合。

```python
>>> from collections import OrderedDict
>>> od = OrderedDict()
>>> od['foo'] = 1
>>> od['bar'] = 2
>>> od['hoge'] = 3
>>> print(od)
OrderedDict([('foo', 1), ('bar', 2), ('hoge', 3)])
```

### bisect

https://docs.python.jp/3/library/bisect.html

ソート済みのリストを 2 分探索で探索。

普通のリストの場合。

```
x = list(range(10**7))
i = x.index(9991234)
```

2 分探索。

```
x = list(range(10**7))
i = bisect_left(x, 9991234)
```

### namedtuple

https://docs.python.jp/3/library/collections.html#collections.namedtuple

フィールドに名前を使ってアクセスできるようになる。
csv や SQL などスキーマの決まっているデータの取得時に便利。

・employees.csv

```
tom,29,leader
jeff,35,manager
```


```python
>>> EmployeeRecord = namedtuple('EmployeeRecord', 'name, age, grade')
>>>for emp in map(EmployeeRecord._make, csv.reader(open("employees.csv", "r"))):
...    print(emp.name, emp.grade)
tom leader
jeff manager
```
