package zlog

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rs/zerolog"
)

// ZHTTPLogger 使用 zerolog 为 HTTP 请求/响应写日志
type ZHTTPLogger struct{}

// bufReader 会将从源 Reader 读取出的数据 buffer 起来，之后可以通过 Bytes 方法获得
type bufReader struct {
	src    io.ReadCloser
	reader io.Reader
	buf    bytes.Buffer
}

// mockReader 用于仿制一个源 Reader，包括最后返回的错误
type mockReader struct {
	reader io.Reader
	data   []byte
	err    error
}

var (
	_ io.ReadCloser = (*bufReader)(nil)
	_ io.ReadCloser = (*mockReader)(nil)
)

func (l ZHTTPLogger) LogRequest(req *http.Request) {
	req.Body = newBufReader(req.Body)
}

func (l ZHTTPLogger) LogResponse(req *http.Request, resp *http.Response, err error, dur time.Duration) {
	logger := zerolog.Ctx(req.Context()).With().
		Str("comp", "http").
		Str("req.method", req.Method).
		Str("req.url", req.URL.String()).
		Bytes("req.body", req.Body.(*bufReader).Data()).
		Str("dur", dur.String()).Logger()

	if resp != nil {
		r := newMockReader(resp.Body)
		resp.Body = r
		logger = logger.With().
			Int("resp.code", resp.StatusCode).
			Bytes("resp.body", r.Data()).Logger()

		err = r.Err()
	}

	if err != nil {
		logger.Error().Err(err).Msg("")
	} else {
		logger.Info().Msg("")
	}
}

func newBufReader(src io.ReadCloser) *bufReader {
	ret := &bufReader{
		src: src,
	}
	ret.reader = io.TeeReader(src, &ret.buf)
	return ret
}

func (r *bufReader) Read(p []byte) (n int, err error) {
	return r.reader.Read(p)
}

func (r *bufReader) Close() error {
	return r.src.Close()
}

func (r *bufReader) Data() []byte {
	return r.buf.Bytes()
}

func newMockReader(src io.ReadCloser) *mockReader {
	// NOTE: ioutil.ReadAll 不会返回 io.EOF
	data, err := ioutil.ReadAll(src)
	src.Close()
	return &mockReader{
		reader: bytes.NewReader(data),
		data:   data,
		err:    err,
	}
}

func (r *mockReader) Read(p []byte) (n int, err error) {
	n, err = r.reader.Read(p)
	if err != nil {
		// NOTE: bytes.Reader 返回的错误只有 EOF
		if err != io.EOF {
			panic(fmt.Errorf("bytes.Reader returns non io.EOF error: %s", err))
		}
		// 若源 Reader 有出错，则复原
		if r.err != nil {
			err = r.err
		}
	}
	return
}

func (r *mockReader) Close() error {
	return nil
}

func (r *mockReader) Data() []byte {
	return r.data
}

func (r *mockReader) Err() error {
	return r.err
}
