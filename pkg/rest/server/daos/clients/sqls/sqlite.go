package sqls

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var o sync.Once

const FileName = "sqlite.db"

type SQLiteClient struct {
	DB *gorm.DB
}

var err error
var sqliteClient *SQLiteClient

const (
	DBUserName     = "root"
	DBUserPassword = "notsosecretpassword"
	DBHost         = "0.0.0.0"
	DBPort         = "3306"
	DBName         = "gorm"
)

func InitSqliteDB() (*SQLiteClient, error) {
	o.Do(func() {
		// if _, err = os.Stat(FileName); err == nil {
		// 	err = os.Remove(FileName)
		// }
		var db *gorm.DB
		
		// db, err = gorm.Open("sqlite3", FileName)
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUserName, DBUserPassword, DBHost, DBPort, DBName)
		db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to the Database! \n", err.Error())
			os.Exit(1)
		}
		sqliteClient = &SQLiteClient{
			DB: db,
		}
	})

	return sqliteClient, err
}
