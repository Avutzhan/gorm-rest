package fmigrate

import (
	"gorm-rest/pkg/ftools/fdb"
	"gorm-rest/pkg/ftools/flog"
)

func Run(model ...interface{}) {
MigrateLoop:
	for _, m := range model {
		if result := fdb.Client.AutoMigrate(m); result != nil {
			flog.Error().Err(result).Msg("error migration model")

			break MigrateLoop
		}
	}
}
