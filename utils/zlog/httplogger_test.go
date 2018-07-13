package zlog

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"
	"testing/iotest"

	//"gopkg.in/jarcoal/httpmock.v1"
	"github.com/stretchr/testify/assert"
)

type NoopReadCloser struct {
	io.Reader
	closed bool
}

func (r *NoopReadCloser) Close() error {
	r.closed = true
	return nil
}

func TestBufReader(t *testing.T) {

	assert := assert.New(t)

	// nil bufReader
	{
		r := (*bufReader)(nil)
		p := []byte{}
		n, err := r.Read(p)
		assert.Equal(0, n)
		assert.Equal(io.EOF, err)

		assert.Nil(r.Close())

		assert.Nil(r.Data())
	}

	// 正常状态 bufReader
	{
		src := &NoopReadCloser{
			Reader: bytes.NewReader([]byte("hello")),
		}
		r := newBufReader(src)
		assert.False(src.closed)

		data, err := ioutil.ReadAll(r)
		assert.Equal([]byte("hello"), data)
		assert.NoError(err)
		assert.Equal([]byte("hello"), r.Data())

		r.Close()
		assert.True(src.closed)

	}

}

func TestReplayReader(t *testing.T) {

	assert := assert.New(t)

	// 正常情况
	{
		src := &NoopReadCloser{
			Reader: bytes.NewReader([]byte("world")),
		}
		assert.False(src.closed)

		r := newReplayReader(src)
		assert.True(src.closed) // 源应该已经关闭
		data, err := ioutil.ReadAll(r)
		assert.Equal([]byte("world"), data)
		assert.NoError(err)
	}

	// 异常情况
	{
		src := &NoopReadCloser{
			Reader: iotest.TimeoutReader(iotest.OneByteReader(bytes.NewReader([]byte("world")))),
		}
		assert.False(src.closed)

		r := newReplayReader(src)
		assert.True(src.closed) // 源应该已经关闭
		data, err := ioutil.ReadAll(r)
		assert.Equal([]byte("w"), data)
		assert.Equal(iotest.ErrTimeout, err)

		assert.Equal([]byte("w"), r.Data())
		assert.Equal(iotest.ErrTimeout, r.Err())

	}

}
