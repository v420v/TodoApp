package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func ConnectDB() (*bun.DB, error) {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	sqldb, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(db)/%s?parseTime=true", user, pass, dbName))
	if err != nil {
		return nil, err
	}

	bunDB := bun.NewDB(sqldb, mysqldialect.New())

	bunDB.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	return bunDB, nil
}
