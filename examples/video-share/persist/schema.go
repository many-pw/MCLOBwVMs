package persist

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() *sqlx.DB {
	url := fmt.Sprintf("%s:%s@(%s:%s)/%s", os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	db, err := sqlx.Connect("mysql", url)
	if err != nil {
		fmt.Println(err)
	}

	return db
}
