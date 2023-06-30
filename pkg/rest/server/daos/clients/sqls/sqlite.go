package sqls

import (
	"os"
	"sync"

	"github.com/jinzhu/gorm"
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
		db, err = gorm.Open("sqlite3", FileName)
		sqliteClient = &SQLiteClient{
			DB: db,
		}
	})

	return sqliteClient, err
}
