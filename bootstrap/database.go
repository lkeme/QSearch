package bootstrap

import (
	"github.com/lkeme/QSearch/pkg/config"
	"github.com/lkeme/QSearch/pkg/database"
	"gorm.io/gorm"
)

// SetupDatabase init db
func SetupDatabase() {
	var (
		logMode   = config.Get("database.log_mode")
		driver    = config.Get("database.driver")
		dsn       = config.Get("database.dsn")
		dialector gorm.Dialector
		err       error
	)

	// err
	if dialector, err = database.Select(driver, dsn); err != nil {
		panic(err)
	}
	// err
	if err = database.Connect(dialector, logMode); err != nil {
		panic(err)
	}
	// err
	if err = database.Opts(); err != nil {
		panic(err)
	}

}
