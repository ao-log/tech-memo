##### いつ rotate されたかのチェック

```
$ cat /var/lib/logrotate/logrotate.status
logrotate state -- version 2
"/var/log/yum.log" 2018-1-6-17:0:0
"/var/log/boot.log" 2018-1-7-13:51:17
"/var/log/chrony/*.log" 2018-1-6-17:0:0
"/var/log/wtmp" 2018-1-6-17:0:0
"/var/log/spooler" 2018-1-7-13:51:17
"/var/log/btmp" 2018-1-6-17:0:0
"/var/log/maillog" 2018-1-7-13:51:17
"/var/log/wpa_supplicant.log" 2018-1-6-17:0:0
"/var/log/secure" 2018-1-7-13:51:17
"/var/log/ppp/connect-errors" 2018-1-6-17:0:0
"/var/log/messages" 2018-1-7-13:51:17
"/var/log/cron" 2018-1-7-13:51:17
```

##### 設定ファイル

|項目|説明|
|---|---|
|rotate|世代数|
|daily/weekly/monthly/yearly|ローテーションの間隔|
|dateext|ログファイルの拡張子を日付表記（デフォルトでYYYYMMDD）にする。 |
|size|指定されたサイズよりも大きい場合のみローテーションする。 |
|maxage|指定された日数以前のファイルを削除する。ローテーションされたタイミングでチェックと削除処理を実行する。 |
|missingok|ログファイルが存在しなくてもエラーを出さない|
|notifempty|ログファイルが空の場合、ローテーションしない|
|compress|gzipで圧縮|
|delaycompress|1世代目は圧縮しない|
|sharedscripts|複数ログに対して、prerotate, postrotate で実行されたコマンドを一度だけ実行する。|
|postrotate [COMMAND] endscript|ログローテーション後に記述されたコマンドを実行|
