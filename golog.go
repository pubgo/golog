package golog

import (
	"github.com/pubgo/assert"
	"github.com/rs/zerolog"
	"os"
	"runtime"
)

func (t *Cfg) appendExtra(ctx zerolog.Context) zerolog.Context {
	h, err := os.Hostname()
	assert.ErrWrap(err, "get Hostname error")

	return ctx.Str("service", cfg.Service, ).
		Str("hostname", h).
		Str("version", cfg.Version).
		Int("cpus", runtime.NumCPU()).
		Bool("debug", cfg.Debug)
}

func (t *Cfg) initLevel() {
	l, err := zerolog.ParseLevel(cfg.LogLevel)
	assert.ErrWrap(err, "parse log level error")
	zerolog.SetGlobalLevel(l)
}
