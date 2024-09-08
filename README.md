
<div align="center">
<h1>Todo App</h1>
<img width="600px" alt="Screenshot 2024-08-19 at 3 10 17" src="https://github.com/user-attachments/assets/3d8e1ff8-8d2f-4d9a-a935-9848774dd01e">
</div>

## ENV
```env file
DB_USER=root
DB_PASS=secret
DB_NAME=gotodo
```

## 実行
Hot-reloadingのために <a href="https://github.com/air-verse/air">Air</a> を使用しており、ソースコード(Go REST API)に変更が検出されると自動的にアプリケーションを再コンパイルし再起動します。

```sh
docker-compose up -d
```

<a href="http://127.0.0.1">http://127.0.0.1</a>

```sh
docker-compose down --rmi all
```

## スキーマ管理
<a href="https://github.com/sqldef/sqldef">sqldef</a> を使うことで、データベーススキーマとコードベースのSQLスキーマファイルとの間の一貫性を保つことができ、スキーマのズレが生じるリスクを低減できます。簡単なTodoアプリなので <a href="https://github.com/sqldef/sqldef">sqldef</a> に決めました。
```sh
$ docker compose exec api bash
$ mysqldef -h db -P 3306 -u root -p secret gotodo < ./internal/db/createTable.sql
```

## API ドキュメント
mainブランチにプッシュすると、自動的にSwaggerを用いて `/docs` フォルダにAPIドキュメントを生成するGitHubActionを設定しました。
<a href="https://v420v.github.io/TodoApp/swagger/">GitHub Pages</a>でいつでも見れるようにしました。

## ER図
ER図は <a href="https://github.com/k1LoW/tbls">tbls</a> で管理し、`createTable.sql`に変更がある度にgithub actionで更新しています。ファイルは `/internal/db/schema` に配置してます。

## CSRF対策 (WIP)
<a href="https://github.com/gorilla/csrf">gorilla/csrf</a> を使い対策を行いました。

## LICENSE
MIT


