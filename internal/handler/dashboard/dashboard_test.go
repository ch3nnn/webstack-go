package dashboard

import (
	"bufio"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/ch3nnn/webstack-go/internal/handler"
	s "github.com/ch3nnn/webstack-go/internal/service"
	"github.com/ch3nnn/webstack-go/internal/service/dashboard"
	"github.com/ch3nnn/webstack-go/pkg/log"
)

// 首帧必须立刻到：仪表盘页面在收到第一条 SSE 事件之前只有一个骨架屏。
func TestHandler_Dashboard_首帧立即返回(t *testing.T) {
	gin.SetMode(gin.TestMode)

	h := NewHandler(
		handler.NewHandler(&log.Logger{Logger: zap.NewNop()}),
		dashboard.NewService(&s.Service{}),
	)

	router := gin.New()
	router.GET("/dashboard", h.Dashboard)

	srv := httptest.NewServer(router)
	defer srv.Close()

	start := time.Now()
	resp, err := http.Get(srv.URL + "/dashboard")
	assert.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	// 读到第一条 data 行就算首帧到达
	firstFrame := make(chan time.Duration, 1)
	go func() {
		reader := bufio.NewReader(resp.Body)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				firstFrame <- -1
				return
			}
			if strings.HasPrefix(line, "data:") {
				firstFrame <- time.Since(start)
				return
			}
		}
	}()

	select {
	case elapsed := <-firstFrame:
		assert.GreaterOrEqual(t, elapsed, time.Duration(0), "连接在读首帧时断开")
		assert.Less(t, elapsed.Milliseconds(), int64(500),
			"首帧耗时 %dms：SSE 没有在连接后立即推送，前端只能一直显示骨架屏", elapsed.Milliseconds())
	case <-time.After(5 * time.Second):
		t.Fatal("5s 内没有收到任何 SSE 帧")
	}
}
