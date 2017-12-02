# フェリーハッカソン のアレ

https://github.com/mizunekoko/mu_sun のサーバーサイド側

# 起動方法

sql/init.sqlにあるDB作成とテーブル作成を行った後，プロジェクトのルートディレクトリにて以下のコマンドを実行。
```bash
go run ./lib/main.go
```
go 1.9.2でやってます。

# web API

8080番で上がっています。

## GET /user/{user_id}/rank

ユーザーのランキングを取得します。{user_id}にはユーザーIDが入ります。Intです。

例
```bash
$ curl hostname:8080/user/1/rank
{
  "rank": 6,
  "user": {
    "id": 1,
    "name": "hoge",
    "score": 300
  }
}
```

rankがランキングです。
userがユーザーのデータです。idがユーザーID，nameがユーザー名，scoreがスコアです。


## POST /register

ユーザーを新規登録します。
新規登録するユーザーの情報のJSONを送ります。

```bash

$ curl hostname:8080/register -X POST --data '{"name": "hoge", "score": 300}'
{
  "rank": 7,
  "user": {
    "id": 9,
    "name": "hoge",
    "score": 300
  }
}
```

帰ってくるものはランキング取得時と同じです。ユーザーIDは自動的に振られます。


## GET /pass/{user_id}/

指定されたユーザーが最も最近に通過した地点のIDを得ます。

```bash
$ curl hostname:8080/pass/2/
{
  "point_id": "hoge",
  "user_id": "2"
}
```

これもJSONが返され，point_idが地点のID，user_idがユーザーIDです。

## POST /pass/{user_id}/{point_id}/

最近の通過情報を登録します。

```bash
$ curl hostname:8080/pass/1/4/ -X POST
{
  "point_id": "4",
  "user_id": "1"
}
```
