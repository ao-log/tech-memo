
[containerd/nerdctl](https://github.com/containerd/nerdctl)

nerdctl: Docker-compatible CLI for containerd


## 背景

[Dockerからcontainerdへの移行 (NTT Tech Conference 2022発表レポート)](https://medium.com/nttlabs/docker-to-containerd-4f3a56e6f2b6)

* Kubernetes 1.24 から dockershim が削除される
* 従来は Docker を介して containerd を使用していたが、近年は Kubernetes から直接 Containerd を使用
* nerdctl は containerd の一部で、このコマンドで Docker ライクな操作ができる
* containerd, nerdctl は色々な機能が追加されている
  * 高速イメージプル
  * 暗号化イメージ
  * イメージへの署名
  * 高速 rootless


## その他参考記事

[コマンドから見るnerdctlとdockerの違い](https://zenn.dev/yatoum/articles/12be13dacd049a)

