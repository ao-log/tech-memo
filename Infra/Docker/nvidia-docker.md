
[NVIDIA Container Toolkit](https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/overview.html)

[Architecture Overview](https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/arch-overview.html)

```nvidia-container-runtime-hook``` により runc の prestart に hook をかけている。

OCI のランタイム標準は [こちら](https://github.com/opencontainers/runtime-spec/blob/main/config.md)。


[prestart](https://github.com/NVIDIA/nvidia-container-toolkit/blob/main/cmd/nvidia-container-runtime-hook/main.go#L68)

prestart にて nvidia-container-cli を実行している。

コードを読めずよく分からないが [ここ](https://github.com/NVIDIA/libnvidia-container/blob/main/src/cgroup.c#L166) でコンテナに対して GPU を割り当てている？
ccあたりも確認が必要か？


[User Guide](https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/user-guide.html)

dockerd の起動オプションにランタイムを追加する。
```
sudo tee /etc/systemd/system/docker.service.d/override.conf <<EOF
[Service]
ExecStart=
ExecStart=/usr/bin/dockerd --host=fd:// --add-runtime=nvidia=/usr/bin/nvidia-container-runtime
EOF
```

docker コマンドでは ```--gpus``` により GPU を指定可能。以下のような指定方法ができる。

* 0,1,2, or GPU-fef8089b
* all: デフォルト
* none
* void

GPU UUID は ```nvidia-smi``` コマンドによってクエリ可能。
```
nvidia-smi -i 3 --query-gpu=uuid --format=csv
```


[NVIDIA Container Toolkit (NVIDIA Docker) は何をしてくれるか](https://qiita.com/tkusumi/items/f275f0737fb5b261a868)

コンテナネームスペースに NVIDIA ライブラリ、デバイスをマウントしてくれる。
mount コマンドで確認可能。



