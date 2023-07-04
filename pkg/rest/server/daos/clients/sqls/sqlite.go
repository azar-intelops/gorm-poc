package sqls

import (
	"os"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var o sync.Once

const FileName = "sqlite.db"

type SQLiteClient struct {
	DB *gorm.DB
}

var err error
var sqliteClient *SQLiteClient

func InitSqliteDB() (*SQLiteClient, error) {
	o.Do(func() {
		if _, err = os.Stat(FileName); err == nil {
			err = os.Remove(FileName)
		}
		var db *gorm.DB
		db, err = gorm.Open(sqlite.Open(FileName), &gorm.Config{})
		sqliteClient = &SQLiteClient{
			DB: db,
		}
	})

	return sqliteClient, err
}
