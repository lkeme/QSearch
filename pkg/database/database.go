package database

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB
var SQLDB *sql.DB

// Select ...
func Select(driver, dsn string) (dc gorm.Dialector, err error) {
	if driver == "" || dsn == "" {
		return nil, errors.New("database field cannot be empty")
	}
	switch driver {
	case "sqlite":
		dc = sqlite.Open(dsn)
	case "mysql":
		dc = mysql.Open(dsn)
	default:
		return nil, errors.New("database connection not supported")
	}
	return
}

// Connect ...
func Connect(dc gorm.Dialector, logMode string) (err error) {
	// init
	DB, err = gorm.Open(dc, logOpts(logMode))
	// err
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// sqlDb
	SQLDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

// Opts ...
func Opts() (err error) {
	// 空闲连接池数
	SQLDB.SetMaxIdleConns(100)
	// 最大连接数
	SQLDB.SetMaxOpenConns(25)
	// 链接超时
	SQLDB.SetConnMaxLifetime(time.Second * time.Duration(5*60))

	return
}

// CurrentDatabase ...
func CurrentDatabase() (dbname string) {
	dbname = DB.Migrator().CurrentDatabase()
	return
}

func logOpts(logMode string) *gorm.Config {
	// 日志模式 	Error Warn Info Silent
	var level gormlogger.LogLevel
	switch logMode {
	case "silent", "Silent":
		level = gormlogger.Silent
	case "error", "Error":
		level = gormlogger.Error
	case "warn", "Warn":
		level = gormlogger.Warn
	case "info", "Info":
		level = gormlogger.Info
	default:
		level = gormlogger.Info
	}

	// 初始化GORM日志配置
	newLogger := gormlogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		gormlogger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  level,       // Log level(这里记得根据需求改一下)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	return &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   newLogger,
	}
}
