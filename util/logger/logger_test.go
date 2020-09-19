// +build !binary_log

package logger_test

import (
	"errors"
	"fabiangonz98/go-cloudnative-practice/util/logger"
	"flag"
	"github.com/rs/zerolog"
	"time"
)

func setup() *logger.Logger {
	zerolog.TimeFieldFormat = ""
	zerolog.TimestampFunc = func() time.Time {
		return time.Date(2008, 1, 8, 17, 5, 05, 0, time.UTC)
	}

	return logger.New(true, true)
}

func ExampleLogger_Print() {
	setup().Print("hello world")

	// Output: {"level":"debug","time":1199811905,"message":"hello world"}
}

func ExampleLogger_Printf() {
	setup().Printf("hello %s", "world")

	// Output: {"level":"debug","time":1199811905,"message":"hello world"}
}

func ExampleLogger_Log() {
	setup().Log().Msg("hello world")

	// Output: {"time":1199811905,"message":"hello world"}
}

func ExampleLogger_Debug() {
	setup().Debug().Msg("debug log")

	// Output: {"level":"debug","time":1199811905,"message":"debug log"}
}

func ExampleLogger_Warn() {
	setup().Warn().Msg("debug log")

	// Output: {"level":"warn","time":1199811905,"message":"debug log"}
}

func ExampleLogger_Info() {
	setup().Info().Msg("debug log")

	// Output: {"level":"info","time":1199811905,"message":"debug log"}
}

func ExampleLogger_Error() {
	setup().Error().Msg("debug log")

	// Output: {"level":"error","time":1199811905,"message":"debug log"}
}

// TODO Check why there is an `unexpected exception` in this test case...
func ExampleLogger_Fatal() {
	err := errors.New("new error")
	service := "myservice"

	setup().
		Fatal().
		Err(err).
		Str("service", service).
		Msgf("Cannot start %s", service)

	// Output: {"level":"fatal","time":1199811905,"error":"new error","service":"myservice","message":"Cannot start myservice"}
}

func Example() {
	l := setup()
	debug := flag.Bool("debug", false, "sets log level to debug")

	flag.Parse()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	l.Debug().Msg("This message appears only when log level is Debug")
	l.Info().Msg("This message appears when log level is Debug or Info")

	if e := l.Debug(); e.Enabled() {
		// Computes when level is enabled
		e.Str("isDebug", "true").Msg("debug msg.")
	}

	// Output: {"level":"info","time":1199811905,"message":"This message appears when log level is Debug or Info"}
}
