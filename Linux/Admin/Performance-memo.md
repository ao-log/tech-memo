
##### ps コマンド

```
// 実メモリ使用量でランキング
$ ps aux --sort -rss
```

##### キャッシュの解放

```
$ cat /proc/vm/sync ; echo 3 > /proc/sys/vm/drop_caches
```
