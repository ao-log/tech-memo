
[What Is Amazon Elastic Container Registry Public?](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/public/what-is-ecr.html)



[Quick start: Publishing to Amazon ECR Public using the AWS CLI](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/public/getting-started-cli.html)

```
aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws

// repositorycatalogdata.json
{
    "description": "This is a test repo for an Amazon ECR tutorial.",
    "architectures": [
        "x86"
    ],
    "operatingSystems": [
        "Linux"
    ],
    "logoImageBlob": "$(cat myrepoimage.png |base64 -w 0)",
    "aboutText": "This repository is used as a tutorial only.",
    "usageText": "This repository is not for public use."
}

aws ecr-public create-repository \
     --repository-name ecr-tutorial \
     --catalog-data file://repositorycatalogdata.json \
     --region us-east-1

docker push public.ecr.aws/registry_alias/ecr-tutorial
```



## Public Garalley

[Using the Amazon ECR Public Gallery](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/public/public-gallery.html)

* https://gallery.ecr.aws がパブリックギャラリーの URL となっている



## Public registries

[Amazon ECR public registries](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/public/public-registries.html)

* Public Garalley では 次の URL でアクセスできる
  * https://gallery.ecr.aws/registry_alias/repository_name
* イメージの URI は次の通り
  * public.ecr.aws/registry_alias/repository_name:image_tag
* プルは誰でもできる
* `GetAuthorizationToken` により Base64 エンコードされたトークンを取得できる。us-east-1 リージョンでの実施が必要。期限は 12 時間


[Updating registry settings](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/public/public-registry-settings.html)

* 通常はデフォルトのエイリアスが設定されるが、カスタムエイリアスをリクエストすることが可能。承認されるとデフォルトエイリアス、カスタムエイリアスの双方を使用可能となる



## Public repositories

[Public repository policies](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/public/public-repository-policies.html)

* 他の AWS アカウントや IAM エンティティにプッシュする権限を与えるような用途で使用できる



## Public images

[Pushing an image to a public repository](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/public/docker-push-ecr-image.html)

* Push に必要な権限
  * ecr-public:GetAuthorizationToken
  * sts:GetServiceBearerToken


[Pulling an image from the Amazon ECR Public Gallery](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/public/docker-pull-ecr-image.html)

* 認証していない場合でもプルできる。ただし、認証時と API 呼び出し数のリミットが異なる



## Security

[Troubleshooting Amazon ECR identity and access](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/public/security_iam_troubleshoot.html)



## Logging

[Logging Amazon ECR Public actions with AWS CloudTrail](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/public/logging-using-cloudtrail.html)

* API は us-east-1 の CloudTrail に記録される



## Service Quotas

[Amazon ECR Public service quotas](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/public/public-service-quotas.html)



## Trouble Shooting

[Amazon ECR Public troubleshooting](https://docs.aws.amazon.com/ja_jp/AmazonECR/latest/public/public-troubleshooting.html)

* docker pull が Access Denied となる場合はトークンの期限切れになっている場合がある。docker logout するとよい


