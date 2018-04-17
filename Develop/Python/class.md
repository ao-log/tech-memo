### クラスの使い方

# 参考

以下の URL、書籍を参考にしている。

##### Python チュートリアル(クラスのページ)

https://docs.python.jp/3/tutorial/classes.html#class-objects)

##### 書籍『Effective Python』

『Effective Python』のサンプルコードは次の URL で公開されている。
https://github.com/bslatkin/effectivepython/tree/master/example_code

### はじめの例題クラス


```python
class MyClass:
    """A simple example class"""

    def __init__(self):
        self.my_string = "sample value"

    def function(self):
        return 'hello world'
```

```python
# インスタンス化。
>>> x = MyClass()
```

##### 属性

**属性**という言葉は、ドットに続く名前全てに対して使う。
以下のように属性用にセッター、ゲッターを使わず、パブリックな属性とした方が属性を直感的に扱えるし、Python 風である。ゲッター、セッターのような書き方や、値の取得、セット時に追加の処理を行いたい場合は、後述する property を使うのが効果的。

```python
# ゲッターを使わずに、普通にパブリックな属性へアクセスする。
>>> print(x.my_string)
sample value

# 例えば、属性の更新はセッターを介さずにこうするだけで良い。
>>> x.my_string = "updated"
>>> print(x.my_string)
updated

# 属性を追加することもできる。
>>> x.added_string = "added"
>>> print(x.added_string)
```

##### 関数

```python
# 関数の呼び出し。
>>> print(x.function())
hello world

# この呼び出し方も等価。
>>>print(MyClass.function(x))
hello world
```

##### その他

```python
# 属性の情報は __dict__ に保存されている。
>>> print(x.__dict__)
{'my_string': 'updated', 'added_string': 'added'}

# docstring は __doc__ に格納されている。
>>> print(x.__doc__)
A simple example class
```

### クラス変数、インスタンス変数

```python
class MyList:
    # クラス変数
    class_list = []       

    def __init__(self):        
        # インスタンス変数
        self.instance_list = []

    def append_class_list(self, x):
        self.class_list.append(x)

    def append_instance_list(self, x):
        self.instance_list.append(x)        
```

```python
>>> x = MyList()
>>> x.append_class_list("apple")
>>> x.append_instance_list("apple")

>>> y = MyList()
>>> y.append_class_list("lemon")
>>> y.append_instance_list("lemon")

# クラス変数は各インスタンスに共有されている点に注意。
>>> print(x.class_list)
['apple', 'lemon']

>>> print(y.class_list)
['apple', 'lemon']

# インスタンス変数だとインスタンスごとに格納されている。
>>> print(x.instance_list)
['apple']

>>> print(y.instance_list)
['lemon']
```

### property

```python
class PositiveValue:
    def __init__(self, value):
        self.value = value

    @property
    def value(self):
        return self._value

    @value.setter
    def value(self, value):
        if value <= 0:
            raise ValueError
        self._value = value
```

```python
# インスタンス作成。ただし、-1 なので ValueError を送出
>>> a = PositiveValue(-1)

# インスタンス作成。
>>> a = PositiveValue(10)

# 負の数値を代入すると、ValueError を送出。
>>> a.value = -10
```

### classmethod, staticmethod

##### インスタンスメソッド

インスタンス化して初めて使えるメソッド

##### クラスメソッド

インスタンス化せずに使えるメソッド

```python
class Cat:
    @classmethod
    def meow(cls):
        print('nyaan')
```

```python
>>> Cat.meow()
nyaan
```

##### クラスメソッドを活用したデザイン。

『Effective Python』の項目24 を参考にしつつ、クラス設計のデザインを考える。
https://github.com/bslatkin/effectivepython/blob/master/example_code/item_24.py

例えばプレゼンテーション資料を作成するためのクラスを実装したとして、利用者が自分の好みのプレゼンテンプレートを使用したい場合。

まず、様々なテンプレートのクラスを実装する（基底クラスを実装しないのはダックタイピングということで・・）。クラスメソッドで同名のメソッドを実装する。

```Python
class SimpleTemplate():
    @classmethod
    def load_template(self):
       ...
```

```Python
class DarkFantasyTemplate():
    @classmethod
    def load_template(self):
       ...
```

プレゼンテーション資料を作成するためのメソッドを実装する。その中の処理でテンプレートのロードを行う。

```Python
def create_presentation(template_cls):
    template_cls.load_template()
    ...
```

利用者は次のように、好みのクラスを指定すれば、テンプレートをロードできる。

```Python
create_presentation(SimpleTemplate)
```



### __call__

```python
class Counter(object):
    def __init__(self):
        self.counter = 0

    def __call__(self):
        self.counter += 1
        return 0

c = Counter()
for i in range(5):
    c()

print(c.counter)
```


# 参考

[Python言語リファレンス:データモデル](https://docs.python.jp/3/reference/datamodel.html#special-method-names)

[Python言語リファレンス:関数定義](https://docs.python.jp/3/reference/compound_stmts.html#function-definitions)

[Python標準ライブラリ:組み込み関数](https://docs.python.jp/3/library/functions.html)

[Dive Into Python 3 日本語版](http://diveintopython3-ja.rdy.jp/special-method-names.html)
