# Go言語でAPIを開発するためのサンプル  
Go言語でAPIを開発するためのサンプルファイルです。  

## 機能一覧  
・configの設定値を取得するAPI（/api/v1/config）  
・usersテーブルの全レコードを取得するAPI（/api/v1/users）  

## 使用技術  
Go "1.20.2"  
Gin  
Nginx  
Docker  
docker-compose  
Air  
MySQL  

## 注意点  
このアプリの起動には以下が必要です。  
・docker環境  
・「.env」（開発環境やテスト環境用の環境変数情報）  

## 使い方  
①ビルド用のコマンド  
```
$ docker compose build --no-cache
```  

<br/>

②起動用のコマンド  
```
$ docker compose up -d
```  

## 環境変数  
ENV=  
TZ=Asia/Tokyo  
MYSQL_DATABASE=  
MYSQL_USER=  
MYSQL_PASSWORD=  
MYSQL_ROOT_PASSWORD=  

## 参考記事
技術ブログも作成していますので、興味がある方は下記の記事を参考にしてみて下さい。  
[・GolangでAPIを開発する方法まとめ](https://tomoyuki65.com/how-to-develop-api-in-golang/)  
