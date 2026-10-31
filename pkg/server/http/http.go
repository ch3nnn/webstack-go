package http

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ch3nnn/webstack-go/pkg/log"
)

func NewGinDefaultServer() *gin.Engine {
	gin.ForceConsoleColor()
	return gin.Default()
}

type Server struct {
	*gin.Engine
	httpSrv *http.Server
	host    string
	port    int
	logger  *log.Logger
	// 请求上下文的根：关闭时整体取消，让 SSE 这类长连接的 handler 退出
	baseCtx    context.Context
	cancelBase context.CancelFunc
}

type Option func(s *Server)

func NewServer(engine *gin.Engine, logger *log.Logger, opts ...Option) *Server {
	s := &Server{
		Engine: engine,
		logger: logger,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func WithServerHost(host string) Option {
	return func(s *Server) {
		s.host = host
	}
}

func WithServerPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

func (s *Server) Start(ctx context.Context) error {
	s.baseCtx, s.cancelBase = context.WithCancel(ctx)
	s.httpSrv = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.host, s.port),
		Handler: s,
		// 每个请求的 ctx 都由 baseCtx 派生，Stop 里取消它即可让长连接退出
		BaseContext: func(net.Listener) context.Context { return s.baseCtx },
	}

	if err := s.httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.logger.Sugar().Fatalf("listen: %s\n", err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.Sugar().Info("Shutting down server...")

	// 先取消请求上下文：SSE 这类长连接永远不会变成 idle，只靠 Shutdown 会一直
	// 等到超时（表现为 "Server forced to shutdown: context deadline exceeded"）。
	// 代价是在途的普通请求会被中断——服务本来也要停了。
	if s.cancelBase != nil {
		s.cancelBase()
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := s.httpSrv.Shutdown(ctx); err != nil {
		// 优雅关闭没走完不是致命错误，记日志并把错误交回调用方，不要把进程硬杀
		s.logger.Sugar().Error("Server shutdown incomplete: ", err)
		return err
	}

	s.logger.Sugar().Info("Server exiting")
	return nil
}
