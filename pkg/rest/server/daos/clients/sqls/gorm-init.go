package sqls

import (
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var o sync.Once

// only when sqlite is choosed
const FileName = "sqlite.db"

type SQLClient struct {
	DB *gorm.DB
}

var db *gorm.DB
var err error
var sqlClient *SQLClient

// only when mysql is choosed
const (
	DBUserName     = "root"
	DBUserPassword = "notsosecretpassword"
	DBHost         = "0.0.0.0"
	DBPort         = "3306"
	DBName         = "gorm"
)

func InitSqlDB() (*SQLClient, error) {

	o.Do(func() {
		// only when sqlite is choosed
		if _, err = os.Stat(FileName); err == nil {
			err = os.Remove(FileName)
		}

		db, err = gorm.Open(sqlite.Open(FileName), &gorm.Config{})

		// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUserName, DBUserPassword, DBHost, DBPort, DBName)
		// mysqlConfigs := mysql.Config{
		// 	DSN: dsn, // data source name
		// 	DefaultStringSize: 256, // default size for string fields
		// 	DisableDatetimePrecision: true, // disable datetime precision, which not supported before MySQL 5.6
		// 	DontSupportRenameIndex: true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		// 	DontSupportRenameColumn: true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		// 	SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
		// }
		// db, _ = gorm.Open(mysql.New(mysqlConfigs), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to the Database! \n", err.Error())
			os.Exit(1)
		}
		sqlClient = &SQLClient{
			DB: db,
		}
	})
	return sqlClient, err
}
