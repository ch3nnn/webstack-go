package http

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"path/filepath"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/ch3nnn/webstack-go/pkg/log"
)

func newTestLogger(t *testing.T) *log.Logger {
	t.Helper()
	conf := viper.New()
	conf.Set("log.log_level", "error")
	conf.Set("log.log_file_name", filepath.Join(t.TempDir(), "test.log"))
	conf.Set("log.encoding", "console")
	return log.NewLog(conf)
}

func freePort(t *testing.T) int {
	t.Helper()
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("取空闲端口失败: %v", err)
	}
	defer func() { _ = l.Close() }()
	return l.Addr().(*net.TCPAddr).Port
}

// 回归：SSE 这类长连接永远不会变成 idle，Stop 必须能让它退出。
// 修复前 Shutdown 会一直等到 5s 超时并返回 context deadline exceeded。
func TestStopEndsLongLivedConnection(t *testing.T) {
	gin.SetMode(gin.TestMode)
	engine := gin.New()

	handlerDone := make(chan struct{})
	connected := make(chan struct{})

	engine.GET("/sse", func(c *gin.Context) {
		stream := make(chan any)
		done := c.Request.Context().Done()
		go func() {
			defer close(handlerDone)
			defer close(stream)
			close(connected)
			select {
			case <-done:
			case <-time.After(30 * time.Second):
			}
		}()
		c.Stream(func(w io.Writer) bool {
			_, ok := <-stream
			return ok
		})
	})

	port := freePort(t)
	srv := NewServer(engine, newTestLogger(t), WithServerHost("127.0.0.1"), WithServerPort(port))
	go func() {
		if err := srv.Start(context.Background()); err != nil {
			t.Logf("server start 返回: %v", err)
		}
	}()

	client := &http.Client{}
	go func() {
		resp, err := client.Get(fmt.Sprintf("http://127.0.0.1:%d/sse", port))
		if err == nil {
			defer func() { _ = resp.Body.Close() }()
			_, _ = io.Copy(io.Discard, resp.Body)
		}
	}()

	select {
	case <-connected:
	case <-time.After(3 * time.Second):
		t.Fatal("SSE 连接未建立")
	}

	start := time.Now()
	if err := srv.Stop(context.Background()); err != nil {
		t.Fatalf("Stop 应正常返回，实际: %v（耗时 %v）", err, time.Since(start))
	}
	if elapsed := time.Since(start); elapsed > 2*time.Second {
		t.Fatalf("Stop 被长连接拖住: %v", elapsed)
	}

	select {
	case <-handlerDone:
	case <-time.After(time.Second):
		t.Fatal("SSE handler 没有随关闭退出")
	}
}
