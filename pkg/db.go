package pkg

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
)

func InitMysql() (*sqlx.DB, error) {
	//import versi 1
	// user := os.Getenv("DB_USER")
	// pass := os.Getenv("DB_PASS")
	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// dbname := os.Getenv("DB_NAME")


	//import versi 2 distructuring
	var options []any
	options = append(options, os.Getenv("DB_USER"))
	options = append(options, os.Getenv("DB_PASS"))
	options = append(options, os.Getenv("DB_HOST"))
	options = append(options, os.Getenv("DB_PORT"))
	options = append(options, os.Getenv("DB_NAME"))

	//username:password@tcp(host:port)/dbname
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", options...)
	return sqlx.Connect("mysql", connStr)

}
