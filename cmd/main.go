package main

import (
	"github.com/Entrio/subenv"
	"gorm-rest/internal/app/api"
	"gorm-rest/pkg/ftools/flog"
	"strings"
)

func main() {
	app := strings.ToLower(subenv.Env("APP_NAME", ""))
	switch app {
	case "api":
		api.Start()
	default:
		flog.Error().Msg("Please, set env app name. Example: api, file-server, worker")
	}
}
