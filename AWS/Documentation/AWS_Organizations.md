# Document

[AWS Organizations のベストプラクティス](https://docs.aws.amazon.com/ja_jp/organizations/latest/userguide/orgs_best-practices.html)



# BlackBelt

[[AWS Black Belt Online Seminar] AWS Organizations](https://pages.awscloud.com/rs/112-TZM-766/images/20180214_AWS-Blackbelt-Organizations.pdf)

* 管理アカウントは組織全体の AWS 利用料の支払い、新規 AWS アカウントの作成、招待、組織ポリシーの適用などの役割を持つ
* OU(組織単位)
* SCP。組織ルート、OU、アカウントに対して割り当て可能。OU の下位階層に継承される。ルートアカウントも適用対象
* IAM ユーザは 1 アカウントに集約し、各アカウントに switch role するようにするのもプラクティスの一つ



# 参考

* Document
  * [AWS Organizations の概要](https://docs.aws.amazon.com/ja_jp/organizations/latest/userguide/orgs_introduction.html)
* Black Belt
  * [[AWS Black Belt Online Seminar] AWS Organizations](https://pages.awscloud.com/rs/112-TZM-766/images/20180214_AWS-Blackbelt-Organizations.pdf)


