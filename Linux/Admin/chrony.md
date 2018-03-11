##### /etc/chrony.conf

```
// サーバの指定
# These servers were defined in the installation:
server metadata.google.internal iburst

// NTP クライアントアクセスの許可対象
# Allow NTP client access from local network.
#allow 192.168/16
```

# 稼動状況確認

##### ソース

```
$ chronyc sources
210 Number of sources = 1
MS Name/IP address         Stratum Poll Reach LastRx Last sample               
===============================================================================
^* metadata.google.internal      2   8   377   113    -55us[  -58us] +/-  819us
```

##### トラッキング

```
$ chronyc tracking
Reference ID    : A9FEA9FE (metadata.google.internal)
Stratum         : 3
Ref time (UTC)  : Sat Jan 13 06:25:25 2018
System time     : 0.000002059 seconds fast of NTP time
Last offset     : -0.000002856 seconds
RMS offset      : 0.000155379 seconds
Frequency       : 80.839 ppm slow
Residual freq   : -0.000 ppm
Skew            : 0.046 ppm
Root delay      : 0.000759258 seconds
Root dispersion : 0.000387361 seconds
Update interval : 128.2 seconds
Leap status     : Normal
```
