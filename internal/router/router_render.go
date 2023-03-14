package router

import (
	"github.com/ch3nnn/webstack-go/internal/api/index"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/render/admin"
	"github.com/ch3nnn/webstack-go/internal/render/authorized"
	"github.com/ch3nnn/webstack-go/internal/render/config"
	"github.com/ch3nnn/webstack-go/internal/render/cron"
	"github.com/ch3nnn/webstack-go/internal/render/dashboard"
	"github.com/ch3nnn/webstack-go/internal/render/generator"
	adminIdex "github.com/ch3nnn/webstack-go/internal/render/index"
	"github.com/ch3nnn/webstack-go/internal/render/install"
	"github.com/ch3nnn/webstack-go/internal/render/tool"
	"github.com/ch3nnn/webstack-go/internal/render/upgrade"
)

func setRenderRouter(r *resource) {

	renderInstall := install.New(r.logger)
	renderIndex := index.New(r.logger, r.db, r.cache)
	renderAdminIndex := adminIdex.New(r.logger, r.db, r.cache)
	renderDashboard := dashboard.New(r.logger, r.db, r.cache)
	renderGenerator := generator_handler.New(r.logger, r.db, r.cache)
	renderConfig := config.New(r.logger, r.db, r.cache)
	renderAuthorized := authorized.New(r.logger, r.db, r.cache)
	renderTool := tool.New(r.logger, r.db, r.cache)
	renderAdmin := admin.New(r.logger, r.db, r.cache)
	renderUpgrade := upgrade.New(r.logger, r.db, r.cache)
	renderCron := cron.New(r.logger, r.db, r.cache)

	// 无需记录日志，无需 RBAC 权限验证
	notRBAC := r.mux.Group("", core.DisableTraceLog, core.DisableRecordMetrics)
	{
		// 首页
		notRBAC.GET("", renderIndex.Index())

		// admin 首页
		notRBAC.GET("/admin", renderAdminIndex.Index())

		// 仪表盘
		notRBAC.GET("/dashboard", renderDashboard.View())

		// 安装
		notRBAC.GET("/install", renderInstall.View())
		notRBAC.POST("/install/execute", renderInstall.Execute())

		// 管理员
		notRBAC.GET("/login", renderAdmin.Login())
		notRBAC.GET("/admin/modify_password", renderAdmin.ModifyPassword())
		notRBAC.GET("/admin/modify_info", renderAdmin.ModifyInfo())
	}

	// 无需记录日志，需要 RBAC 权限验证
	render := r.mux.Group("", core.DisableTraceLog, core.DisableRecordMetrics)
	{
		// 配置信息
		render.GET("/config/email", renderConfig.Email())
		render.GET("/config/code", renderConfig.Code())

		// 代码生成器
		render.GET("/generator/gorm", renderGenerator.GormView())
		render.POST("/generator/gorm/execute", renderGenerator.GormExecute())

		render.GET("/generator/handler", renderGenerator.HandlerView())
		render.POST("/generator/handler/execute", renderGenerator.HandlerExecute())

		// 调用方
		render.GET("/authorized/list", renderAuthorized.List())
		render.GET("/authorized/add", renderAuthorized.Add())
		render.GET("/authorized/api/:id", renderAuthorized.Api())
		render.GET("/authorized/demo", renderAuthorized.Demo())

		// 管理员
		render.GET("/admin/list", renderAdmin.List())
		render.GET("/admin/add", renderAdmin.Add())
		render.GET("/admin/site", renderAdmin.Site())
		render.GET("/admin/site/add", renderAdmin.AddSite())
		render.GET("/admin/category", renderAdmin.Category())
		render.GET("/admin/menu", renderAdmin.Menu())
		render.GET("/admin/menu_action/:id", renderAdmin.MenuAction())
		render.GET("/admin/action/:id", renderAdmin.AdminMenu())

		// 升级
		render.GET("/upgrade", renderUpgrade.UpgradeView())
		render.POST("/upgrade/execute", renderUpgrade.UpgradeExecute())

		// 工具箱
		render.GET("/tool/hashids", renderTool.HashIds())
		render.GET("/tool/logs", renderTool.Log())
		render.GET("/tool/cache", renderTool.Cache())
		render.GET("/tool/data", renderTool.Data())
		render.GET("/tool/websocket", renderTool.Websocket())

		// 后台任务
		render.GET("/cron/list", renderCron.List())
		render.GET("/cron/add", renderCron.Add())
		render.GET("/cron/edit/:id", renderCron.Edit())
	}
}
