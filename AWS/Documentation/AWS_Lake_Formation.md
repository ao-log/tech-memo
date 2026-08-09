

# Black Belt

[Lake Formation 基礎編](https://pages.awscloud.com/rs/112-TZM-766/images/AWS-Black-Belt_2023_AWS-Lake-Formation_1010_v1.pdf)

* AWS Glue Data Catalog に対するメタデータのアクセス制御と Amazon S3 ロケーション内のデータへのアクセス制御
* AWS Lake Formation によるアクセスの制御は IAM による制御の後に⾏われる。そのため Lake Formation による細かな制御を⾏うには、まず IAM による制御を荒くする必要がある
* LF タグが推奨。利用者、テーブルそれぞれにタグを設定できる。ただし、列レベルまでしか制御できない
