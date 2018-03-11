# インストール

```
$ sudo yum install dstat
```

# コマンド使用方法

##### 利用状況の推移確認

```
$ dstat
You did not select any stats, using -cdngy by default.
----total-cpu-usage---- -dsk/total- -net/total- ---paging-- ---system--
usr sys idl wai hiq siq| read  writ| recv  send|  in   out | int   csw
  1   1  97   1   0   0| 548k  156k|   0     0 |   0     0 |  96   133
  0   0 100   0   0   0|   0     0 | 198B  158B|   0     0 |  20    21
  0   0 100   0   0   0|   0     0 |  66B  830B|   0     0 |  15    16
  0   0 100   0   0   0|   0     0 | 144B  700B|   0     0 |  14    12
  0   1  99   0   0   0|   0     0 |  66B  350B|   0     0 |  32    31
```

##### CPU を多く使っているプロセス

```
$ dstat -ta --top-cpu
----system---- ----total-cpu-usage---- -dsk/total- -net/total- ---paging-- ---system-- -most-expensive-
     time     |usr sys idl wai hiq siq| read  writ| recv  send|  in   out | int   csw |  cpu process   
13-01 06:09:06|  1   1  98   1   0   0| 390k  110k|   0     0 |   0     0 |  76   102 |systemd      0.1
13-01 06:09:07|  1   2  97   0   0   0|   0     0 | 132B  102B|   0     0 |  41    27 |                
13-01 06:09:08|  0   0 100   0   0   0|   0     0 |  66B 1046B|   0     0 |  30    26 |                
13-01 06:09:09|  2   1  97   0   0   0|   0     0 |  66B  406B|   0     0 |  66    36 |node_exporter2.0
13-01 06:09:10|  0   1  99   0   0   0|   0     0 |  66B  406B|   0     0 |  28    18 |
```

##### I/O の多いプロセス

```
$ dstat -ta --top-io-adv --top-bio-adv
Terminal width too small, trimming output.
----system---- ----total-cpu-usage---- -dsk/total- -net/total- ---paging-- ---system-- -------most-expensive-i/o-process------->
     time     |usr sys idl wai hiq siq| read  writ| recv  send|  in   out | int   csw |process              pid  read write cpu>
13-01 06:19:40|  1   0  99   0   0   0| 208k   60k|   0     0 |   0     0 |  52    66 |bash                  971   183k 985B0.0%>
13-01 06:19:41|  0   1  99   0   0   0|   0     0 | 132B  102B|   0     0 |  19    16 |                                        >
13-01 06:19:42|  0   0 100   0   0   0|   0     0 |  66B 1230B|   0     0 |  33    26 |                                        >
13-01 06:19:43|  1   0  99   0   0   0|   0     0 |  66B  438B|   0     0 |  28    19 |                                        >
13-01 06:19:44|  0   0 100   0   0   0|   0     0 |1142B 1666B|   0     0 |  55    64 |prometheus            342   757B 240B  0%>
13-01 06:19:45|  1   0  99   0   0   0|   0     0 |  66B  140B|   0     0 |  28    25 |                                        >
```

# オプション

|オプション|説明|
|---|---|
|-t|時刻を表示|
|-c|CPU使用状況。-C でコア指定(例:-C 0,1,2,3,total)|
|-m|メモリ使用状況|
|-d|IO状況|
|-n|ネットワーク使用状況(B/sec)。-N でインタフェース指定|

# 参考

[Qiita:dstatの便利なオプションまとめ](https://qiita.com/harukasan/items/b18e484662943d834901)
