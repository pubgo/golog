package golog

import (
	"github.com/pubgo/golog/internal/hooks"
	zwriter "github.com/pubgo/golog/internal/writer"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"sync"
)

func New() *Cfg {
	return GetCfg()
}

type Cfg struct {
	Debug    bool     `json:"debug,omitempty" yml:"debug" toml:"debug"`
	Service  string   `json:"service,omitempty" yml:"service" toml:"service"`
	Version  string   `json:"version,omitempty" yml:"version" toml:"version"`
	LogLevel string   `json:"loglevel,omitempty" yml:"loglevel" toml:"loglevel"`
	Hooks    []string `json:"hooks,omitempty" yml:"hooks" toml:"hooks"`
	Writer   []map[string]interface{}
}

func (t *Cfg) Init() {
	t.initLevel()

	w := zwriter.NewZWriter()
	w.Append(zerolog.ConsoleWriter{})
	log.Logger = t.appendExtra(log.Output(w).With().Caller()).Logger()
	log.Logger = log.Logger.Hook(hooks.HH)
}

var cfg *Cfg
var once sync.Once

func GetCfg() *Cfg {
	once.Do(func() {
		cfg = new(Cfg)
		cfg.Debug = true
		cfg.Version = "v0.0.1"
		cfg.Service = "_test"
		cfg.LogLevel = "debug"
	})
	return cfg
}
