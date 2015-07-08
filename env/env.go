package env

import (
	log "github.com/Sirupsen/logrus"
	"github.com/anominet/anomi/cache"
	"github.com/anominet/anomi/env/internal"
	"os"
)

const (
	VERSION             = "0.1"
	DEFAULT_API_PORT    = "8080"
	DEFAULT_REDIS_HOST  = "127.0.0.1"
	DEFAULT_REDIS_PORT  = "6379"
	DEFAULT_SEPARATOR   = ":"
	DEFAULT_AUTH_HEADER = "X-USER-TOKEN"
	REDIS_HOST_ENV_VAR  = "REDIS_PORT_" + DEFAULT_REDIS_PORT + "_TCP_ADDR"
)

var DEFAULT_SERIALIZER = cache.JsonSerialzer{}

type Env struct {
	C          cache.Cache
	AuthHeader string
	Log        internal.Logger
}

func New(redis_host string, debug bool) *Env {
	e := Env{}
	e.Log = internal.Logger{log.New()}
	e.Log.Formatter = &log.TextFormatter{
		ForceColors:      true,
		DisableColors:    false,
		DisableTimestamp: false,
		FullTimestamp:    true,
		TimestampFormat:  "",
		DisableSorting:   false,
	}
	if debug {
		e.Log.Level = log.DebugLevel
	}

	if redis_host == DEFAULT_REDIS_HOST {
		if redis_host = os.Getenv(REDIS_HOST_ENV_VAR); redis_host == "" {
			redis_host = DEFAULT_REDIS_HOST
		}
	}
	e.Log.Debug("Using redis host: " + redis_host)

	e.C = &cache.RedisCache{}
	e.C.Dial(redis_host + ":" + DEFAULT_REDIS_PORT)
	e.C.SetSerializer(DEFAULT_SERIALIZER)
	e.C.SetSeparator(DEFAULT_SEPARATOR)
	e.C.SetLogger(e.Log)
	e.AuthHeader = DEFAULT_AUTH_HEADER
	return &e
}
