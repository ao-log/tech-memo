プレイブックの書き方のメモ。

### vars

・web.yml

```yaml
- hosts: web
  become: yes
  vars_files:
    - vars/production.yml
  roles:
    - common
```

・ vars/production.yml
```yaml
tenant:
  user: example_user
  uid: 1010
```

・roles/common/tasks.main.yml

```yaml
- name: Add user
  user:
    name: "{{ tenant.user }}"
    uid: "{{ tenant.uid}}"
```

### with_items

```yaml
- name: Locate files
  copy:
    src={{ item.src }}
    dest={{ item.dest }}
    owner=root group=root mode=0644
  with_items:
    - { src: example1.conf, dest: /etc/example1.conf }
    - { src: example2.conf, dest: /etc/example2.conf }
```

### ファイルに追記

```yaml
- name: Add env settings to .bash_profile
  become: yes
  become_user: "{{ tenant.user }}"
  lineinfile:
    dest=~{{ tenant.user }}/.bash_profile
    line={{ item }}
  with_items:
    - "export PATH=/opt/myapp/bin:$PATH"
    - "export LD_LIBRARY_PATH=/opt/myapp/lib:$LD_LIBRARY_PATH"
```

### ディレクトリ作成

```yaml
- name: Make directories
  file:
    path={{ item }}
    state=directory
    owner={{ tenant.user }}
    group={{ tenant.user }}
    mode=0755
  with_items:
    - "{{ dirs }}"
```
