/**
 * @Author: chentong
 * @Date: 2025/02/11 13:25
 */

package dashboard

import (
	"flag"
	"os"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/ch3nnn/webstack-go/internal/dal/query"
	"github.com/ch3nnn/webstack-go/internal/dal/repository"
	s "github.com/ch3nnn/webstack-go/internal/service"
	"github.com/ch3nnn/webstack-go/pkg/config"
	"github.com/ch3nnn/webstack-go/pkg/jwt"
	"github.com/ch3nnn/webstack-go/pkg/log"
)

var (
	logger *log.Logger
	j      *jwt.JWT
)

func setupRepository(t *testing.T) (*repository.Repository, sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm connection: %v", err)
	}

	query.SetDefault(db)

	return repository.NewRepository(logger, db), mock
}

func TestMain(m *testing.M) {
	err := os.Setenv("APP_CONF", "../../../config/local.yml")
	if err != nil {
		panic(err)
	}

	envConf := flag.String("conf", "config/local.yml", "config path, eg: -conf ./config/local.yml")
	flag.Parse()
	conf := config.NewConfig(*envConf)

	logger = log.NewLog(conf)
	j = jwt.NewJwt(conf)

	code := m.Run()
	os.Exit(code)
}

func TestDashboardService_Dashboard(t *testing.T) {
	ctx := &gin.Context{}
	repo, _ := setupRepository(t)
	srv := s.NewService(logger, j, repo)

	dashboardService := NewService(srv)
	_, err := dashboardService.Dashboard(ctx)

	assert.NoError(t, err)
}

// 仪表盘是 SSE 首屏：采集一旦阻塞，前端就只能在骨架屏里干等。
// cpu.Percent(interval) 会阻塞整整一个 interval（实测 1001ms），这里给 500ms 上限守住。
func TestService_Dashboard_采集不能阻塞(t *testing.T) {
	ctx := &gin.Context{}

	repo, _ := setupRepository(t)
	srv := s.NewService(logger, j, repo)
	dashboardService := NewService(srv)

	start := time.Now()
	resp, err := dashboardService.Dashboard(ctx)
	elapsed := time.Since(start)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Less(t, elapsed.Milliseconds(), int64(500),
		"Dashboard 采集耗时 %dms，说明又阻塞在 CPU 采样上了", elapsed.Milliseconds())
}
