package server

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/docs"
	categoryHandler "github.com/ch3nnn/webstack-go/internal/handler/category"
	configHandler "github.com/ch3nnn/webstack-go/internal/handler/config"
	dashboardHandler "github.com/ch3nnn/webstack-go/internal/handler/dashboard"
	indexHandler "github.com/ch3nnn/webstack-go/internal/handler/index"
	siteHandler "github.com/ch3nnn/webstack-go/internal/handler/site"
	userHandler "github.com/ch3nnn/webstack-go/internal/handler/user"
	"github.com/ch3nnn/webstack-go/internal/middleware"
	"github.com/ch3nnn/webstack-go/pkg/jwt"
	"github.com/ch3nnn/webstack-go/pkg/log"
	httpx "github.com/ch3nnn/webstack-go/pkg/server/http"
	assets "github.com/ch3nnn/webstack-go/web"
)

func NewHTTPServer(
	engine *gin.Engine,
	logger *log.Logger,
	conf *viper.Viper,
	jwt *jwt.JWT,
	dashboardHandler *dashboardHandler.Handler,
	indexHandler *indexHandler.Handler,
	userHandler *userHandler.Handler,
	siteHandler *siteHandler.Handler,
	categoryHandler *categoryHandler.Handler,
	configHandler *configHandler.Handler,
) *httpx.Server {
	gin.SetMode(gin.DebugMode)
	s := httpx.NewServer(
		engine,
		logger,
		httpx.WithServerHost(conf.GetString("http.host")),
		httpx.WithServerPort(conf.GetInt("http.port")),
	)
	// 上传大小限制
	s.MaxMultipartMemory = 32 << 20 // 32MB

	s.StaticFS("/assets", http.FS(assets.Static))
	s.SetHTMLTemplate(template.Must(template.New("").ParseFS(assets.Templates, "templates/**/*")))

	// Swagger Doc
	docs.SwaggerInfo.BasePath = "/v1"
	s.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
		ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.PersistAuthorization(true),
	))

	// Middleware
	s.Use(
		middleware.CORSMiddleware(),
		middleware.RequestLogMiddleware(logger),
		middleware.ResponseLogMiddleware(logger),
	)
	// 404
	s.NoRoute(v1.ErrHandler404)
	// Index HTML
	s.GET("/", indexHandler.Index)
	// About HTML
	s.GET("/about", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "about.html", nil)
	})
	// Login HTML
	s.GET("login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "admin_login.html", nil)
	})
	// Render HTML
	render := s.Group("/admin").Use(middleware.NoStrictAuth(jwt, logger))
	{
		render.GET("", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "admin_index.html", nil)
		})
		render.GET("dashboard", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "dashboard.html", nil)
		})
		render.GET("modify_password", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "admin_modify_password.html", nil)
		})
		render.GET("category", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "category_view.html", nil)
		})
		render.GET("site", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "site_list.html", nil)
		})
		render.GET("site/add", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "site_add.html", nil)
		})
		render.GET("config", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "conf_web.html", nil)
		})
	}

	v1 := s.Group("/api")
	{
		// No route group has permission
		noAuthRouter := v1.Group("")
		{
			noAuthRouter.POST("/login", userHandler.Login)
			noAuthRouter.GET("/about", indexHandler.About)
		}
		// Strict permission routing group
		strictAuthRouter := v1.Group("/admin").Use(middleware.StrictAuth(jwt, logger))
		{
			// Dashboard
			strictAuthRouter.GET("/dashboard", dashboardHandler.Dashboard) // SSE（Server-Sent Events）
			// User
			strictAuthRouter.GET("/info", userHandler.Info)
			strictAuthRouter.POST("/logout", userHandler.Logout)
			strictAuthRouter.PATCH("/modify_password", userHandler.UpdatePassword)
			// Category
			strictAuthRouter.GET("/category", categoryHandler.List)
			strictAuthRouter.POST("/category", categoryHandler.Create)
			strictAuthRouter.PUT("/category/update", categoryHandler.Update)
			strictAuthRouter.GET("/category/:id", categoryHandler.Detail)
			strictAuthRouter.DELETE("/category/:id", categoryHandler.Delete)
			// Site
			strictAuthRouter.GET("/site", siteHandler.List)
			strictAuthRouter.GET("/site/sync/:id", siteHandler.SyncSite)
			strictAuthRouter.POST("/site", siteHandler.Create)
			strictAuthRouter.DELETE("/site/:id", siteHandler.Delete)
			strictAuthRouter.PUT("/site/:id", siteHandler.Update)
			strictAuthRouter.GET("/site/export", siteHandler.Export)
			// Config
			strictAuthRouter.GET("/config", configHandler.Config)
			strictAuthRouter.PUT("/config", configHandler.Update)

		}
	}

	return s
}
