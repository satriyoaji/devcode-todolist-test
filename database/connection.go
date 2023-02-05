package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/satriyoaji/todolist-challenge-go/config"
	"github.com/satriyoaji/todolist-challenge-go/helpers"
)

var DB *sql.DB

func Init() {

	connectionString := config.GoDotEnvVariable("MYSQL_USER") +
		":" + config.GoDotEnvVariable("MYSQL_PASSWORD") +
		"@tcp(" + config.GoDotEnvVariable("MYSQL_HOST") +
		":" + config.GoDotEnvVariable("MYSQL_PORT") +
		")/" + config.GoDotEnvVariable("MYSQL_DBNAME") +
		"?parseTime=True&loc=Local"
	connection, err := sql.Open("mysql", connectionString)
	helpers.OutputPanicError(err)

	DB = connection

	//connection.AutoMigrate('...')
}

func GetConnection() *sql.DB {
	return DB
}
