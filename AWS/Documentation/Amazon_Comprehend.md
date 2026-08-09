
[Amazon Comprehend とは](https://docs.aws.amazon.com/ja_jp/comprehend/latest/dg/what-is.html)

* ドキュメント内のエンティティ、キーフレーズ、言語、感情、その他の共通要素を認識することでインサイトを作成


[Insights](https://docs.aws.amazon.com/ja_jp/comprehend/latest/dg/concepts-insights.html)

* 以下のようなインサイトを得られる

  * エンティティ: 人、位置所、場所など
  * Events: 特定の種類のイベントと関連する詳細
  * キーフレーズ: ドキュメントに含まれるキーフレーズを抽出
  * 個人を特定できる情報 : 住所や銀行口座番号、電話番号など
  * 主要言語: ドキュメント内の主要言語を特定
  * センチメント : ポジティブのこともあれば、中立やネガティブ、あるいはそれらが混在することもある
  * ターゲットセンチメント: ドキュメントに記載されている特定のエンティティのセンチメント
  * 構文分析: ドキュメント内の各単語を解析し、その単語の品詞を特定


[API を使用したリアルタイムの分析](https://docs.aws.amazon.com/ja_jp/comprehend/latest/dg/using-api-sync.html)

```
aws comprehend detect-entities \
    --region region \
    --language-code "en" \
    --text "It is raining today in Seattle."
```

応答
```json
{
    "Entities": [
        {
            "Text": "today",
            "Score": 0.97,
            "Type": "DATE",
            "BeginOffset": 14,
            "EndOffset": 19
        },
        {
            "Text": "Seattle",
            "Score": 0.95,
            "Type": "LOCATION",
            "BeginOffset": 23,
            "EndOffset": 30
        }
    ],
    "LanguageCode": "en"
}
```


[信頼と安全性](https://docs.aws.amazon.com/ja_jp/comprehend/latest/dg/trust-safety.html)

* Amazon Comprehend Trust and Safety が対応できること
  * Toxicity detection: 有害、攻撃的、または不適切な可能性のあるコンテンツを検出
  * Intent classification: 明示的または暗示的な悪意のある意図を持つコンテンツを検出
  * Privacy protection: 個人識別情報を検出して編集することができる


[個人を特定できる情報 (PII)](https://docs.aws.amazon.com/ja_jp/comprehend/latest/dg/pii.html)

* 非同期ジョブでは PII の位置のオフセットだけでなく、該当箇所をマスクするなどの編集ができる


[カスタム分類](https://docs.aws.amazon.com/ja_jp/comprehend/latest/dg/how-document-classification.html)

[マルチクラスモード](https://docs.aws.amazon.com/ja_jp/comprehend/latest/dg/prep-classifier-data-multi-class.html)

* 以下の例のような CSV を用意して、トレーニングさせる
```
CLASS,Text of document 1
```


[カスタム分類のリアルタイム分析 (コンソール)](https://docs.aws.amazon.com/ja_jp/comprehend/latest/dg/custom-sync.html)

* エンドポイントに、カスタムモデルをアタッチすることで呼び出せるようになる

