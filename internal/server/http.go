package server

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"strings"

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
	"github.com/ch3nnn/webstack-go/pkg/tools"
	assets "github.com/ch3nnn/webstack-go/web"
)

// templateFuncs 是模板里可用的函数集合。
//
// imgSrc：图标可能是外链、SVG 的 data URI，或位图的纯 base64，统一拼成可直接用的 src。
// 返回 template.URL 是必须的：html/template 的 URL 过滤器会把 data: 协议整个替换成
// #ZgotmplZ，图片就成空白了；template.URL 只是跳过协议过滤，属性里的引号仍会被转义。
func templateFuncs() template.FuncMap {
	return template.FuncMap{
		"imgSrc": func(icon string) template.URL {
			return template.URL(tools.ImgSrc(icon))
		},
	}
}

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
	s.SetHTMLTemplate(template.Must(template.New("").
		Funcs(templateFuncs()).
		ParseFS(assets.Templates, "templates/**/*")))

	// 管理端 SPA 产物：静态资源走 /ui，index.html 读入内存后由 NoRoute 回退返回
	spaFS, err := fs.Sub(assets.SPA, "admin/dist")
	if err != nil {
		panic(err)
	}
	spaIndex, err := fs.ReadFile(spaFS, "index.html")
	if err != nil {
		panic(err)
	}
	s.StaticFS("/ui", http.FS(spaFS))

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
	// 公共站保持服务端渲染（SEO 依赖）
	s.GET("/", indexHandler.Index)
	s.GET("/about", indexHandler.AboutPage)

	// 管理端为 SPA：/login 与 /admin/** 一律返回 dist/index.html，鉴权由前端路由守卫接管
	serveSPA := func(ctx *gin.Context) {
		ctx.Header("Cache-Control", "no-cache")
		ctx.Data(http.StatusOK, "text/html; charset=utf-8", spaIndex)
	}
	// 404
	s.NoRoute(func(ctx *gin.Context) {
		if p := ctx.Request.URL.Path; p == "/login" || p == "/admin" || strings.HasPrefix(p, "/admin/") {
			serveSPA(ctx)
			return
		}
		v1.ErrHandler404(ctx)
	})

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
