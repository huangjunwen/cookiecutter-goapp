package zlog

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog/hlog"
)

// ZLogFormatter 使用 zerolog 为 chi.Router 写 request 日志，例如：
//
//     router.Use(hlog.NewHandler(someZeroLogger))
//     router.Use(middleware.RequestLogger(ZLogFormatter{}))
//     router.Use(middleware.Recoverer)
//
type ZLogFormatter struct{}

// zLogEntry 实现 middleware.LogEntry 接口
type zLogEntry http.Request

var (
	_ middleware.LogFormatter = ZLogFormatter{}
	_ middleware.LogEntry     = (*zLogEntry)(nil)
)

// NewLogEntry 实现 middleware.LogFormatter 接口
func (f ZLogFormatter) NewLogEntry(r *http.Request) middleware.LogEntry {
	return (*zLogEntry)(r)
}

func (l *zLogEntry) Write(status, sz int, dur time.Duration) {
	r := (*http.Request)(l)
	hlog.FromRequest(r).Info().
		Str("method", r.Method).
		Str("url", r.URL.String()).
		Int("status", status).
		Int("sz", sz).
		Str("dur", dur.String()).
		Msg("")
}

func (l *zLogEntry) Panic(v interface{}, stack []byte) {
	r := (*http.Request)(l)
	hlog.FromRequest(r).Error().Str("panic", fmt.Sprint(v)).Bytes("stack", stack).Msg("")
}
