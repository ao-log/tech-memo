### f-strings

```python
name = 'myname'
email = 'myname@example.com'
print(f'name: {name}, email: {email}')
```

実行結果

```
name: myname, email: myname@example.com
```

### フォーマット

```python
item = ['tomato', 'tofu', 'moyashi']
fee = [100, 40, 25]

print(f'ITEM                FEE')
print('-' * 23)
for i, v in zip(item, value):
    print(f'{i:<10} {v:>10}円')
```    

実行結果

```
ITEM                FEE
-----------------------
tomato            100円
tofu               40円
moyashi            25円
```

### 参考

[PEP 498](https://www.python.org/dev/peps/pep-0498/)
