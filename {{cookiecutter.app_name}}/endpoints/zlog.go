package endpoints

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog/hlog"
)

// zlogFormatter 实现 middleware.LogFormatter 接口
type zlogFormatter struct{}

// zlogEntry 实现 middleware.LogEntry 接口
type zlogEntry http.Request

func (f zlogFormatter) NewLogEntry(r *http.Request) middleware.LogEntry {
	return (*zlogEntry)(r)
}

func (l *zlogEntry) Write(status, sz int, dur time.Duration) {
	r := (*http.Request)(l)
	hlog.FromRequest(r).Info().
		Str("method", r.Method).
		Str("url", r.URL.String()).
		Int("status", status).
		Int("sz", sz).
		Str("dur", dur.String()).
		Msg("")
}

func (l *zlogEntry) Panic(v interface{}, stack []byte) {
	r := (*http.Request)(l)
	hlog.FromRequest(r).Error().Str("panic", fmt.Sprint(v)).Bytes("stack", stack).Msg("")
}
