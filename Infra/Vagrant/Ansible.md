

# Ansible でプロビジョニング

プロビジョニングで zsh をインストールする例。

```
├── Vagrantfile
└── provisioning
    ├── hosts
    └── site.yml
```

```
$ cat Vagrantfile
# -*- mode: ruby -*-

Vagrant.configure("2") do |config|
  config.vm.box = "centos/7"

  config.vm.network "private_network", ip: "192.168.33.30"

  config.vm.provision "ansible" do |ansible|
    ansible.playbook = "provisioning/site.yml"
    ansible.inventory_path = "provisioning/hosts"
    ansible.limit = 'all'
  end

  config.vm.provider "libvirt" do |vb|
    vb.memory = "1024"
  end
end

```

```
$ cat provisioning/hosts
[vagrants]
192.168.33.30
```

```
$ cat provisioning/site.yml
- hosts: vagrants
  sudo: true
  user: vagrant
  tasks:
    - name: install packages zsh
      yum: name=zsh update_cache=yes
```


# 参考

[AnsibleとVagrantで開発環境を構築する](https://knowledge.sakura.ad.jp/2882/)
