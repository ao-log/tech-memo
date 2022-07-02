
[ConfigMap](https://kubernetes.io/ja/docs/concepts/configuration/configmap/)

機密性のない情報を格納するためのオブジェクト。

Pod から参照する環境変数のような使い方もあれば、ConfigMap に基づいて動作を調整するアドオンやオペレーターもある。


#### pod 内のコンテナに設定する方法

4 種類ある。

* コンテナ内のコマンドと引数
* 環境変数をコンテナに渡す
* 読み取り専用のボリューム内にファイルを追加し、アプリケーションがそのファイルを読み取る
* Kubernetes API を使用して ConfigMap を読み込むコードを書き、そのコードを Pod 内で実行する


#### リファレンス

[ConfigMap v1 core](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#configmap-v1-core)


#### サンプル

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: game-demo
data:
  # プロパティーに似たキー。各キーは単純な値にマッピングされている
  player_initial_lives: "3"
  ui_properties_file_name: "user-interface.properties"

  # ファイルに似たキー
  game.properties: |
    enemy.types=aliens,monsters
    player.maximum-lives=5    
  user-interface.properties: |
    color.good=purple
    color.bad=yellow
    allow.textmode=true    
```


#### Pod からの参照

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: configmap-demo-pod
spec:
  containers:
    - name: demo
      image: alpine
      command: ["sleep", "3600"]
      env:
        # 環境変数を定義します。
        - name: PLAYER_INITIAL_LIVES # ここではConfigMap内のキーの名前とは違い
                                     # 大文字が使われていることに着目してください。
          valueFrom:
            configMapKeyRef:
              name: game-demo           # この値を取得するConfigMap。
              key: player_initial_lives # 取得するキー。
        - name: UI_PROPERTIES_FILE_NAME
          valueFrom:
            configMapKeyRef:
              name: game-demo
              key: ui_properties_file_name
      volumeMounts:
      - name: config
        mountPath: "/config"
        readOnly: true
  volumes:
    # Podレベルでボリュームを設定し、Pod内のコンテナにマウントします。
    - name: config
      configMap:
        # マウントしたいConfigMapの名前を指定します。
        name: game-demo
        # ファイルとして作成するConfigMapのキーの配列
        items:
        - key: "game.properties"
          path: "game.properties"
        - key: "user-interface.properties"
          path: "user-interface.properties"
```

