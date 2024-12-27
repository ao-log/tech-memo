
[Introducing NLB TCP configurable idle timeout](https://aws.amazon.com/jp/blogs/networking-and-content-delivery/introducing-nlb-tcp-configurable-idle-timeout/)

* NLB のアイドルタイムアウトはデフォルトで TCP 350 秒、UDP 120 秒
* アイドルタイムアウト後にトラフィック送信を試みた場合は NLB はクライアントに RST を応答する
* アイドルタイムアウトを防ぐには TCP keepalive を設定する
* NLB のアイドルタイムアウト時間 > アプリケーションのアイドルタイムアウト時間、とすることが望ましい
* NLB のタイムアウトを長くした場合は、フローテーブルが溢れるリスクを高める

