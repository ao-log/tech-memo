
```
$ python36 -m venv test

$ source test/bin/activate

$ pip install testinfra
$ pip install paramiko
```


```python
def test_passwd_file(host):
    passwd = host.file("/etc/passwd")
    assert passwd.contains("root")
    assert passwd.user == "root"
    assert passwd.group == "root"
    assert passwd.mode == 0o644
```

```
$ py.test -v test_myinfra.py
======================== test session starts ========================
platform linux -- Python 3.6.6, pytest-3.9.3, py-1.7.0, pluggy-0.8.0 -- /home/k-aono/test/bin/python36
cachedir: .pytest_cache
rootdir: /home/k-aono/testinfra, inifile:
plugins: testinfra-1.17.0
collected 1 item                                                    

test_myinfra.py::test_passwd_file[local] PASSED               [100%]

===================== 1 passed in 0.05 seconds ======================
```

エラーを含むケースの場合

```
$ py.test -v test_myinfra.py
======================== test session starts ========================
platform linux -- Python 3.6.6, pytest-3.9.3, py-1.7.0, pluggy-0.8.0 -- /home/k-aono/test/bin/python36
cachedir: .pytest_cache
rootdir: /home/k-aono/testinfra, inifile:
plugins: testinfra-1.17.0
collected 3 items                                                   

test_myinfra.py::test_passwd_file[local] PASSED               [ 33%]
test_myinfra.py::test_nginx_is_installed[local] FAILED        [ 66%]
test_myinfra.py::test_nginx_running_and_enabled[local] FAILED [100%]

============================= FAILURES ==============================
__________________ test_nginx_is_installed[local] ___________________

host = <testinfra.host.Host object at 0x7f19a54bca90>

    def test_nginx_is_installed(host):
        nginx = host.package("nginx")
>       assert nginx.is_installed
E       assert False
E        +  where False = <package nginx>.is_installed

test_myinfra.py:10: AssertionError
_______________ test_nginx_running_and_enabled[local] _______________

host = <testinfra.host.Host object at 0x7f19a54bca90>

    def test_nginx_running_and_enabled(host):
        nginx = host.service("nginx")
>       assert nginx.is_running
E       assert False
E        +  where False = <service nginx>.is_running

test_myinfra.py:15: AssertionError
================ 2 failed, 1 passed in 0.14 seconds =================

(test) [k-aono@instance-1 testinfra]$ echo $?
1
```
