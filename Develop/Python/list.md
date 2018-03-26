
# リスト

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

# TODO

defaultdict,
OrderedDict,
bisect,
namedtuple
