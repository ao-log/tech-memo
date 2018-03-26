epel からインストールします。

```
# yum install epel-release
# yum install python34
```

### venv を作れない場合

次のエラーが出る場合

```
$ python3 -m venv flask
Error: Command '['/home/vagrant/python/flask/bin/python3', '-Im', 'ensurepip', '--upgrade', '--default-pip']' returned non-zero exit status 1
```

locale を次のように設定するとうまく行くようになる。

```
$ export LC_ALL=en_US.UTF-8
$ python3 -m venv flask
```

なお、locale は設定前は次のようになっていた。

```
$ locale
locale: Cannot set LC_CTYPE to default locale: No such file or directory
locale: Cannot set LC_MESSAGES to default locale: No such file or directory
locale: Cannot set LC_ALL to default locale: No such file or directory
LANG=ja_JP.UTF-8
LC_CTYPE="ja_JP.UTF-8"
LC_NUMERIC=ja_JP.UTF-8
LC_TIME=ja_JP.UTF-8
LC_COLLATE="ja_JP.UTF-8"
LC_MONETARY=ja_JP.UTF-8
LC_MESSAGES="ja_JP.UTF-8"
LC_PAPER=ja_JP.UTF-8
LC_NAME=ja_JP.UTF-8
LC_ADDRESS=ja_JP.UTF-8
LC_TELEPHONE=ja_JP.UTF-8
LC_MEASUREMENT=ja_JP.UTF-8
LC_IDENTIFICATION=ja_JP.UTF-8
LC_ALL=
```

```python3 -m venv --without-pip 環境名``` でもいいようだ。

[参考:stackoverflow](https://stackoverflow.com/questions/24123150/pyvenv-3-4-returned-non-zero-exit-status-1)
