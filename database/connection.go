package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/satriyoaji/todolist-challenge-go/config"
	"github.com/satriyoaji/todolist-challenge-go/helpers"
	"github.com/satriyoaji/todolist-challenge-go/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	connectionStringGorm := config.GoDotEnvVariable("MYSQL_USER") +
		":" + config.GoDotEnvVariable("MYSQL_PASSWORD") +
		"@tcp(" + config.GoDotEnvVariable("MYSQL_HOST") +
		":" + config.GoDotEnvVariable("MYSQL_PORT") +
		")/" + config.GoDotEnvVariable("MYSQL_DBNAME") +
		"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connectionStringGorm), &gorm.Config{})
	db.AutoMigrate(&models.Todo{}, &models.Activity{})

	DB = connection

	//connection.AutoMigrate('...')
}

func GetConnection() *sql.DB {
	return DB
}
