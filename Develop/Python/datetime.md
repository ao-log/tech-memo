### 現在時刻

```python
>>> import datetime
>>> now = datetime.datetime.now()

>>> print(now)
2018-05-30 13:27:50.588271

>>> datetime.datetime.utcnow().isoformat()
'2018-07-28T04:43:34.589898'
```

### 時刻の差分

```python
>>> print(now - datetime.timedelta(weeks=10))
2018-03-21 13:27:50.588271

>>> print(now - datetime.timedelta(days=10))
2018-05-20 13:27:50.588271

>>> print(now - datetime.timedelta(hours=10))
2018-05-30 03:27:50.588271

>>> print(now - datetime.timedelta(minutes=10))
2018-05-30 13:17:50.588271

>>> print(now - datetime.timedelta(seconds=10))
2018-05-30 13:27:40.588271
```

### 文字列から datetime 型へ

```python
>>> string_date = '2018/05/29 12:20:30'

>>> dt = dt = datetime.datetime.strptime(string_date, '%Y/%m/%d %H:%M:%S')

>>> dt
datetime.datetime(2018, 5, 29, 12, 20, 30)
```
