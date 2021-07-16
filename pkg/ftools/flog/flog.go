package flog

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"os"
)

var (
	info  *zerolog.Logger
	debug *zerolog.Logger
	warn  *zerolog.Logger
	error *zerolog.Logger
)

func Info() *zerolog.Event {
	return info.Info()
}
func Debug() *zerolog.Event {
	return info.Debug()
}
func Warn() *zerolog.Event {
	return info.Warn()
}
func Error() *zerolog.Event {
	return info.Error()
}
func InfoCtx(c *fiber.Ctx) *zerolog.Event {
	return info.Info().Bytes("id", c.Request().Header.Peek("X-Request-ID")).Bytes("path", c.Request().Header.RequestURI()).Bytes("method", c.Request().Header.Method()).Int("status", c.Response().StatusCode())
}
func DebugCtx(c *fiber.Ctx) *zerolog.Event {
	return debug.Debug().Bytes("id", c.Request().Header.Peek("X-Request-ID")).Bytes("path", c.Request().Header.RequestURI()).Bytes("method", c.Request().Header.Method()).Int("status", c.Response().StatusCode())
}

func ErrorCtx(c *fiber.Ctx) *zerolog.Event {
	return error.Error().Bytes("id", c.Request().Header.Peek("X-Request-ID")).Bytes("path", c.Request().Header.RequestURI()).Bytes("method", c.Request().Header.Method()).Int("status", c.Response().StatusCode())
}
func WarnCtx(c *fiber.Ctx) *zerolog.Event {
	return warn.Warn().Bytes("id", c.Request().Header.Peek("X-Request-ID")).Bytes("path", c.Request().Header.RequestURI()).Bytes("method", c.Request().Header.Method()).Int("status", c.Response().StatusCode())
}

func init() {
	fmt.Println(">> init log tools")
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"}
	multi := zerolog.MultiLevelWriter(consoleWriter)
	logger := zerolog.New(multi).With().Timestamp().Logger()

	info = &logger
	debug = &logger
	warn = &logger
	error = &logger
}
