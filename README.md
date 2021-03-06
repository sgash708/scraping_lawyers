# スクレイピング * DDD
並列処理でスクレイピングを行います。
実行したデータは、踏み台サーバを経由してMySQLに格納します。

# 技術
* golang ver1.16
* MySQL 5.7
* ChromeDriver

# DDDについて
## domain層
* ビジネスロジックを記述
* DBアクセスなどの技術的な実装はさせない
* どの層にも依存させない

## infra層
* DBアクセス(CRUD)を記述
* domain層に依存する
* domain層のrepository定義したインターフェースを実装する
* DIP(依存逆転原則)
    * 抽象的に依存させることで依存性を逆転させる
    * 抽象とは、<code>名前|引数|返り値 だけを持つ関数</code>

## application/usecase層
* interfaces層から情報を受け取る
    * domain層で定義している関数を用いて任意のビジネスロジックを実行する

### 使用例
* validationやユーザID生成などのビジネスロジックを記述
* infra層で実装したDBアクセスに関する処理をdomain層を介して間接的に呼び出す

## interfaces層
* 外部データの差異を吸収してusecaseに渡し、結果を返却する役割を担う役割
* HTTPリクエストを受け取り、UseCaseを使って処理実行
    * クライアントに返す、サーバログ出力
