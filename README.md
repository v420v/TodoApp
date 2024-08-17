
# Todo App

## ENV
```env file
DB_USER=root
DB_PASS=secret
DB_NAME=gotodo
```

## Hot-reloading
Hot-reloadingのために <a href="https://github.com/air-verse/air">Air</a> を使用しており、ソースコードに変更が検出されると自動的にアプリケーションを再コンパイルし再起動することができます。

```sh
docker-compose up -d
```

<a href="http://127.0.0.1">http://127.0.0.1</a>

```sh
docker-compose down --rmi all
```

## マイグレーション
<a href="https://github.com/sqldef/sqldef">sqldef</a> を使うことで、データベーススキーマとコードベースのSQLスキーマファイルとの間の一貫性を保つことができます。これにより、開発環境、テスト環境、本番環境でスキーマのズレが生じるリスクを低減できます。簡単なTodoアプリなので <a href="https://github.com/sqldef/sqldef">sqldef</a> に決めました。
```sh
docker exec -it go-server-app-1 bash
mysqldef -h db -P 3306 -u root -p secret gotodo < ./internal/db/createTable.sql
```

## API ドキュメント
GitHub Actionで、mainブランチにプッシュすると、自動的にSwaggerを用いて `/docs` フォルダにAPIドキュメントが生成されるようになっております。
<a href="https://v420v.github.io/TodoApp/swagger/">GitHub Pages</a>でいつでも見れるようにしました。

## ER図
ER図は <a href="https://github.com/k1LoW/tbls">tbls</a> で管理し、`createTable.sql`に変更がある度にgithub actionで更新しています。ファイルは `/internal/db/schema` に配置してます。

