# production と staging に対応した Ansible ソースの書き方

## ここで考察すること

production、staging の２つの環境があるとする。

やりたいことは、Ansible の一つのソースで、両方に対応できるようにする。
つまり、次のようにコマンドで環境を使い分けて実行すること。

```
# production を対象にする
$ ansible-playbook -i production_hosts site.yml

# staging を対象にする
$ ansible-playbook -i staging_hosts site.yml
```

次の記事を大いに参考にしている。

[Qiita: Ansibleのインベントリファイルでステージを切り替える](https://qiita.com/NewGyu/items/5de31d76d2488ab27ed6)


### 今回作成する Ansible ソースのポイント

ポイントを簡単に箇条書きにする。

* 環境ごとにファイルを用意する（staging_hosts, production_hosts）。このファイル中で、変数 stage に環境名を設定する。
* web, db という粒度でプレイブックを作成する(web.yml, db.yml)。これらのファイルから、環境に応じたロール、変数ファイルを呼び出す。


### ディレクトリ構成

```
|-- site.yml
|-- production_hosts
|-- staging_hosts
|-- web.yml
|-- db.yml
|-- roles
|   |-- staging_pretask/
|   |-- common/ 
|   |-- db/
|   `-- web/
`-- vars
　　　　|-- staging.yml
　　　　|-- production.yml
　　　　|-- web.yml
　　　　`-- db.yml
```

### site.yml

各ロールに該当するロールを import するだけの内容である。（余談だが、include はバージョン 2.8 で無くなるとのこと。）

```
- import_playbook: web.yml
- import_playbook: db.yml
```

### inventory

ステージごとにホストを記載する。
特徴的なのは、ここでステージの環境変数を設定していること。

##### production_hosts

```
[web]
web.prd.example.com

[db]
db.prd.example.com

[all:vars]
stage=production
```

##### staging_hosts

```
[web]
web.stg.example.com

[db]
db.stg.example.com

[all:vars]
stage=staging
```


### 上位プレイブックで環境の制御

web, db という粒度でプレイブックを作成する。
この粒度のプレイブックでは以下のことを行う。

* 自身を組み立てるのに必要なロールの呼び出し(roles)
* 変数の設定(vars_files)

ポイントは、vars_files でインベントリで指定した環境の変数を読み込んでいること。


##### Tips

* tags をロールごとに設定しておくと、プレイブック実行時に指定ロールのみの処理を実行可能。
* 特定ステージでのみ実行したいロールがある場合は、when を使ってステージを指定する。

##### web.yml

```
- hosts: web
  become: yes
  vars_files:
    - vars/{{ stage }}.yml
    - vars/web.yml
  roles:
    - { role: staging_pretask, tags: staging_pretask, when: stage == 'staging' }
    - { role: common,      tags: common }
    - { role: web,         tags: web }
```

db.yml も同じような考え方で作成する。

### タスク単位での、環境ごとに動作を変えたい場合

片方の環境でのみ実行したい処理は、when 句によってステージを指定する。

#####  roles/common/tasks/main.yml 

```
- name: install packages
  yum: name={{ item }} 
  with_items:
   - "{{ yum_pkg }}"
  when:
   - stage == "production"
```


### プレイブック実行


##### Tips

* --check オプションを付与することでテストランを実行することが出来る。（よく知られているとは思うが）

##### コマンド

-i でステージング、プロダクションを切り替える。

```
# production を対象とする。
$ ansible-playbook -i production_hosts site.yml

# staging を対象とする。
$ ansible-playbook -i staging_hosts site.yml

# staging を対象とする。web サーバのみ。
$ ansible-playbook -i staging_hosts web.yml

# staging を対象とする。web サーバを対象に、common ロールのみ。
$ ansible-playbook -i staging_hosts web.yml --tags common
```

# 参考

[Ansible公式: Best Practices](http://docs.ansible.com/ansible/latest/playbooks_best_practices.html)

[Ansible inventoryパターン](https://dev.classmethod.jp/server-side/ansible/ansible-inventory-pattern/)

[Qiita:Ansibleのインベントリファイルでステージを切り替える](https://qiita.com/NewGyu/items/5de31d76d2488ab27ed6)

