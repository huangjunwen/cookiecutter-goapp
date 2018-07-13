package zlog

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ernesto-jimenez/httplogger"
	"github.com/rs/zerolog"
)

// ZHTTPLogger 使用 zerolog 为 HTTP 请求/响应写日志
//
// NOTE: 它会完整地读取请求/响应的 body 到内存并打印到日志中，故只适合于请求/响应都比较简短的场合使用（如 API 调用）
type ZHTTPLogger struct{}

// bufReader 会将从源 Reader 读取出的数据 buffer 起来，之后可以通过 Data 方法获得
type bufReader struct {
	src    io.ReadCloser
	reader io.Reader
	buf    bytes.Buffer
}

// replayReader 用于重播一个源 Reader，包括最后返回的错误
type replayReader struct {
	reader io.Reader
	data   []byte
	err    error
}

var (
	_ io.ReadCloser = (*bufReader)(nil)
	_ io.ReadCloser = (*replayReader)(nil)
)

// NewZLoggedTransport 封装一个 RoundTripper，使用 zerolog 打印请求/响应：包括 body
func NewZLoggedTransport(rt http.RoundTripper) http.RoundTripper {
	return httplogger.NewLoggedTransport(rt, ZHTTPLogger{})
}

func (l ZHTTPLogger) LogRequest(req *http.Request) {
	// 替换 request 的 body
	req.Body = newBufReader(req.Body)
}

func (l ZHTTPLogger) LogResponse(req *http.Request, resp *http.Response, err error, dur time.Duration) {

	var (
		respReader *replayReader
	)
	if resp != nil {
		// NOTE: RoundTrip must return err == nil if it obtained
		// a response, regardless of the response's HTTP status code.
		//
		// 所以当 resp 非空时，err 一定为空

		// 使用 replayReader 完整读取 response 的 body 并替换之
		respReader = newReplayReader(resp.Body)
		resp.Body = respReader
		// 若读取 response 的过程中出现了错误，则也算错误
		err = respReader.Err()
	}

	var (
		ev *zerolog.Event
	)
	logger := zerolog.Ctx(req.Context())
	if err != nil {
		ev = logger.Error().Err(err)
	} else {
		ev = logger.Info()
	}

	ev = ev.Str("comp", "http.client").Str("dur", dur.String()).
		Str("req.method", req.Method).Str("req.url", req.URL.String()).Bytes("req.body", req.Body.(*bufReader).Data())
	if respReader != nil {
		ev = ev.Int("resp.code", resp.StatusCode).Bytes("resp.body", respReader.Data())
	}
	ev.Msg("")

}

// newBufReader 新建一个 bufReader，src 可以为 nil
func newBufReader(src io.ReadCloser) *bufReader {
	if src == nil {
		return nil
	}
	ret := &bufReader{
		src: src,
	}
	// 从 src 读多少，就写多少到 buf 中
	ret.reader = io.TeeReader(src, &ret.buf)
	return ret
}

func (r *bufReader) Read(p []byte) (n int, err error) {
	if r == nil {
		return 0, io.EOF
	}
	return r.reader.Read(p)
}

func (r *bufReader) Close() error {
	if r == nil {
		return nil
	}
	return r.src.Close()
}

func (r *bufReader) Data() []byte {
	if r == nil {
		return nil
	}
	return r.buf.Bytes()
}

// newReplayReader 新建一个 replayReader，src 不可为 nil
func newReplayReader(src io.ReadCloser) *replayReader {
	// NOTE: ioutil.ReadAll 不会返回 io.EOF
	data, err := ioutil.ReadAll(src)
	src.Close()
	return &replayReader{
		reader: bytes.NewReader(data),
		data:   data,
		err:    err,
	}
}

func (r *replayReader) Read(p []byte) (n int, err error) {
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

func (r *replayReader) Close() error {
	return nil
}

func (r *replayReader) Data() []byte {
	return r.data
}

func (r *replayReader) Err() error {
	return r.err
}
