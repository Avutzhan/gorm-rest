package fconf

import (
	"fmt"
	"github.com/Entrio/subenv"
	"time"
)

type conf struct {
	server   server
	database database
	setting  setting
}

type server struct {
	ip   string
	port int
}

type auth struct {
	username string
	password string
}

type database struct {
	server    server
	auth      auth
	name      string
	migration bool
	debug     bool
	ping      int
}

type setting struct {
	debug    bool
	loglevel string
	needCors bool
}

var client *conf

func init() {
	client = &conf{
		server: server{
			ip:   subenv.Env("APP_SERVER_IP", ""),
			port: subenv.EnvI("APP_SERVER_PORT", 8080),
		},
		database: database{
			server: server{
				ip:   subenv.Env("APP_DATABASE_SERVER_IP", ""),
				port: subenv.EnvI("APP_DATABASE_SERVER_PORT", 3306),
			},
			auth: auth{
				username: subenv.Env("APP_DATABASE_AUTH_USERNAME", "root"),
				password: subenv.Env("APP_DATABASE_AUTH_PASSWORD", "root"),
			},
			name:      subenv.Env("APP_DATABASE_NAME", "dms_medical"),
			migration: subenv.EnvB("APP_DATABASE_MIGRATION", true),
			debug:     subenv.EnvB("APP_DATABASE_DEBUG", true),
			ping:      subenv.EnvI("APP_DATABASE_PING", 15),
		},
		setting: setting{
			debug:    subenv.EnvB("APP_SETTING_DEBUG", true),
			loglevel: subenv.Env("APP_SETTING_LOG_LEVEL", "debug"),
			needCors: subenv.EnvB("APP_SETTING_CORS", false),
		},
	}
}

func DatabaseConnect() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		client.database.auth.username,
		client.database.auth.password,
		client.database.server.ip,
		client.database.server.port,
		client.database.name,
	)
}

func DatabaseMigration() bool {

	return client.database.migration
}

func DatabaseDebug() bool {

	return client.database.debug
}

func DatabasePing() time.Duration {
	return time.Duration(client.database.ping) * time.Second
}

func ServerUP() string {
	return fmt.Sprintf("%s:%d", client.server.ip, client.server.port)
}

func LogLevel() string {
	switch client.setting.loglevel {
	case "disable", "fatal", "error", "warn", "info", "debug":
		return client.setting.loglevel
	default:
		return "info"
	}
}

func UseCors() bool {
	return client.setting.needCors
}
