
* [永続ボリューム](https://kubernetes.io/ja/docs/concepts/storage/persistent-volumes/)
* [ボリュームの動的プロビジョニング(Dynamic Volume Provisioning)](https://kubernetes.io/ja/docs/concepts/storage/dynamic-provisioning/)

* PV は PVC 経由で使用。
* PVC の Dynamic Provisioning を使用することで、PVC 作成タイミングで動的に PV が作成される。事前に Storage クラスの作成が必要。


#### EKS の場合の例

* [Amazon EKS で永続的ストレージを使用するにはどうすればよいですか?](https://aws.amazon.com/jp/premiumsupport/knowledge-center/eks-persistent-storage/)
* [Amazon EBS CSI ドライバー](https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/ebs-csi.html)
* [テスト用マニフェスト](https://github.com/kubernetes-sigs/aws-ebs-csi-driver/tree/master/examples/kubernetes/dynamic-provisioning/manifests)

サービスアカウントの作成。
```
eksctl create iamserviceaccount \
    --name ebs-csi-controller-sa \
    --namespace kube-system \
    --cluster my-cluster \
    --attach-policy-arn arn:aws:iam::111122223333:policy/AmazonEKS_EBS_CSI_Driver_Policy \
    --approve \
    --role-only
```

EKS EBS CSI アドオンの追加。
```
eksctl create addon --name aws-ebs-csi-driver \
    --cluster my-cluster \
    --service-account-role-arn  \
    arn:aws:iam::111122223333:role/role-name \
    --force
```

・StorageClass
```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: ebs-claim
spec:
  accessModes:
    - ReadWriteOnce
  storageClassName: ebs-sc
  resources:
    requests:
      storage: 4Gi
```

・PVC
```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: ebs-claim
spec:
  accessModes:
    - ReadWriteOnce
  storageClassName: ebs-sc
  resources:
    requests:
      storage: 4Gi
```

・Pod
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: app
spec:
  containers:
  - name: app
    image: centos
    command: ["/bin/sh"]
    args: ["-c", "while true; do echo $(date -u) >> /data/out.txt; sleep 5; done"]
    volumeMounts:
    - name: persistent-storage
      mountPath: /data
  volumes:
  - name: persistent-storage
    persistentVolumeClaim:
      claimName: ebs-claim
```
