package index

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/ch3nnn/webstack-go/internal/dal/model"
	"github.com/ch3nnn/webstack-go/internal/dal/repository"
	s "github.com/ch3nnn/webstack-go/internal/service"
)

func TestService_RecordVisit(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo, _ := setupRepository(t)
	srv := s.NewService(logger, j, repo)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	visitDao := repository.NewMockISysVisitDao(ctrl)
	visitFunc := repository.NewMockiCustomGenSysVisitFunc(ctrl)

	// IP 必须哈希后入库，且 day 是当天
	visitDao.EXPECT().WithContext(gomock.Any()).Return(visitFunc)
	visitFunc.EXPECT().Record(gomock.Any()).DoAndReturn(func(visit *model.SysVisit) error {
		assert.Len(t, visit.IP, 64, "IP 应该存 sha256（64 位十六进制）")
		assert.NotEqual(t, "203.0.113.9", visit.IP, "IP 不能以明文落库")
		assert.Len(t, visit.Day, len("2026-11-01"))
		return nil
	})

	svc := NewService(srv, nil, nil, nil, visitDao)

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = httptest.NewRequest("GET", "/", nil)
	ctx.Request.RemoteAddr = "203.0.113.9:5000"

	svc.RecordVisit(ctx)
}

func TestService_RecordVisit_统计失败不影响首页(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo, _ := setupRepository(t)
	srv := s.NewService(logger, j, repo)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	visitDao := repository.NewMockISysVisitDao(ctrl)
	visitFunc := repository.NewMockiCustomGenSysVisitFunc(ctrl)
	visitDao.EXPECT().WithContext(gomock.Any()).Return(visitFunc)
	visitFunc.EXPECT().Record(gomock.Any()).Return(assert.AnError)

	svc := NewService(srv, nil, nil, nil, visitDao)

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = httptest.NewRequest("GET", "/", nil)

	// 落库失败只记日志，不能把错误抛回请求链路
	assert.NotPanics(t, func() { svc.RecordVisit(ctx) })
}
