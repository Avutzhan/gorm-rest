package fdb

import (
	"database/sql"
	"gorm-rest/pkg/ftools/fconf"
	"gorm-rest/pkg/ftools/flog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var Client *gorm.DB
var err error
var sqlDB *sql.DB

func Do() error {
	Client, err = gorm.Open(mysql.Open(fconf.DatabaseConnect()), &gorm.Config{DisableAutomaticPing: false, Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {

		return err
	}
	sqlDB, err = Client.DB()
	if err != nil {
		return err
	}

	if fconf.DatabaseDebug() {
		Client = Client.Debug()
	}
	go ping()
	return nil
}
func ping() {
	p := fconf.DatabasePing()
	if p == 0 {
		return
	}
PingLoop:
	for {
		time.Sleep(p)
		if result := sqlDB.Ping(); result != nil {
			if result := Do(); result != nil {
				flog.Warn().Err(result).Msg("error reconnect to db")
			} else {
				break PingLoop
			}
		}
	}
}
func Close() {
	if result := sqlDB.Close(); result != nil {
		flog.Warn().Err(result).Msg("error close db")
	}
}
