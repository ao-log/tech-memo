
[AppArmor security profiles for Docker](https://docs.docker.com/engine/security/apparmor/)

* Docker は `docker-default` プロファイルを自動生成する。
* `docker run` 時に、無指定時は `docker-default` プロファイルが使用される。
* ポリシーは以下のように指定する。
```
$ docker run --rm -it --security-opt apparmor=docker-default hello-world
```

#### プロファイルのロード、アンロード

```
// ロード
$ apparmor_parser -r -W /path/to/your_profile
$ docker run --rm -it --security-opt apparmor=your_profile hello-world

// アンロード
$ apparmor_parser -R /path/to/profile
```

#### AppArmor のデバッグ

```
// ロードされているプロファイルの一覧
$ aa-status
```



