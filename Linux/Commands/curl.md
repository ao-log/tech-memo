### ヘッダ確認

```
curl --verbose https://example.com > /dev/null
```


### JSON を POST

```
curl http://example.com/ \
  -X POST  \
  -H 'Content-Type:application/json'  \
  -d '{"name":"uname","email":"uname@example.com"}'
```
