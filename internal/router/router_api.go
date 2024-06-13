package router

import (
	"github.com/ch3nnn/webstack-go/internal/api/admin"
	"github.com/ch3nnn/webstack-go/internal/api/authorized"
	"github.com/ch3nnn/webstack-go/internal/api/category"
	"github.com/ch3nnn/webstack-go/internal/api/config"
	"github.com/ch3nnn/webstack-go/internal/api/cron"
	"github.com/ch3nnn/webstack-go/internal/api/helper"
	"github.com/ch3nnn/webstack-go/internal/api/menu"
	"github.com/ch3nnn/webstack-go/internal/api/site"
	"github.com/ch3nnn/webstack-go/internal/api/tool"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
)

func setApiRouter(r *resource) {
	// helper
	helperHandler := helper.New(r.logger, r.db, r.cache)

	helpers := r.mux.Group("/helper")
	{
		helpers.GET("/md5/:str", helperHandler.Md5())
		helpers.POST("/sign", helperHandler.Sign())
	}

	// admin
	adminHandler := admin.New(r.logger, r.db, r.cache)

	// 需要签名验证，无需登录验证，无需 RBAC 权限验证
	login := r.mux.Group("/api", r.interceptors.CheckSignature())
	{
		login.POST("/login", adminHandler.Login())
	}

	// 需要签名验证、登录验证，无需 RBAC 权限验证
	notRBAC := r.mux.Group("/api", core.WrapAuthHandler(r.interceptors.CheckLogin), r.interceptors.CheckSignature())
	{
		notRBAC.POST("/admin/logout", adminHandler.Logout())
		notRBAC.PATCH("/admin/modify_password", adminHandler.ModifyPassword())
		notRBAC.GET("/admin/info", adminHandler.Detail())
		notRBAC.PATCH("/admin/modify_personal_info", adminHandler.ModifyPersonalInfo())
	}

	// 需要签名验证、登录验证、RBAC 权限验证 TODO
	//api := r.mux.Group("/api", core.WrapAuthHandler(r.interceptors.CheckLogin), r.interceptors.CheckSignature(), r.interceptors.CheckRBAC())
	api := r.mux.Group("/api")
	{
		// authorized
		authorizedHandler := authorized.New(r.logger, r.db, r.cache)
		api.POST("/authorized", authorizedHandler.Create())
		api.GET("/authorized", authorizedHandler.List())
		api.PATCH("/authorized/used", authorizedHandler.UpdateUsed())
		api.DELETE("/authorized/:id", core.AliasForRecordMetrics("/api/authorized/info"), authorizedHandler.Delete())

		api.POST("/authorized_api", authorizedHandler.CreateAPI())
		api.GET("/authorized_api", authorizedHandler.ListAPI())
		api.DELETE("/authorized_api/:id", core.AliasForRecordMetrics("/api/authorized_api/info"), authorizedHandler.DeleteAPI())

		api.POST("/admin", adminHandler.Create())
		api.GET("/admin", adminHandler.List())
		api.PATCH("/admin/used", adminHandler.UpdateUsed())
		api.PATCH("/admin/offline", adminHandler.Offline())
		api.PATCH("/admin/reset_password/:id", core.AliasForRecordMetrics("/api/admin/reset_password"), adminHandler.ResetPassword())
		api.DELETE("/admin/:id", core.AliasForRecordMetrics("/api/admin"), adminHandler.Delete())

		api.POST("/admin/menu", adminHandler.CreateAdminMenu())
		api.GET("/admin/menu/:id", core.AliasForRecordMetrics("/api/admin/menu"), adminHandler.ListAdminMenu())

		// menu
		menuHandler := menu.New(r.logger, r.db, r.cache)
		api.POST("/menu", menuHandler.Create())
		api.GET("/menu", menuHandler.List())
		api.GET("/menu/:id", core.AliasForRecordMetrics("/api/menu"), menuHandler.Detail())
		api.PATCH("/menu/used", menuHandler.UpdateUsed())
		api.PATCH("/menu/sort", menuHandler.UpdateSort())
		api.DELETE("/menu/:id", core.AliasForRecordMetrics("/api/menu"), menuHandler.Delete())
		api.POST("/menu_action", menuHandler.CreateAction())
		api.GET("/menu_action", menuHandler.ListAction())
		api.DELETE("/menu_action/:id", core.AliasForRecordMetrics("/api/menu_action"), menuHandler.DeleteAction())

		// category
		categoryHandler := category.New(r.logger, r.db, r.cache)
		api.GET("/category", categoryHandler.List())
		api.GET("/category/:id", categoryHandler.Detail())
		api.POST("/category", categoryHandler.Create())
		api.DELETE("/category/:id", categoryHandler.Delete())
		api.PATCH("/category/used", categoryHandler.UpdateUsed())
		api.PATCH("/category/sort", categoryHandler.UpdateSort())

		// site
		siteHandler := site.New(r.logger, r.db, r.cache)
		api.GET("/site", siteHandler.List())
		api.PATCH("/site/used", siteHandler.UpdateUsed())
		api.DELETE("/site/:id", siteHandler.Delete())
		api.POST("/site", siteHandler.Create())
		api.PUT("/site", siteHandler.UpdateSite())
		api.GET("/site/category", siteHandler.CategoryList())
		api.GET("/site/sync/:id", siteHandler.SyncSite())

		// tool
		toolHandler := tool.New(r.logger, r.db, r.cache)
		api.GET("/tool/hashids/encode/:id", core.AliasForRecordMetrics("/api/tool/hashids/encode"), toolHandler.HashIdsEncode())
		api.GET("/tool/hashids/decode/:id", core.AliasForRecordMetrics("/api/tool/hashids/decode"), toolHandler.HashIdsDecode())
		api.POST("/tool/cache/search", toolHandler.SearchCache())
		api.PATCH("/tool/cache/clear", toolHandler.ClearCache())
		api.GET("/tool/data/dbs", toolHandler.Dbs())
		api.POST("/tool/data/tables", toolHandler.Tables())
		api.POST("/tool/data/mysql", toolHandler.SearchMySQL())
		api.POST("/tool/send_message", toolHandler.SendMessage())

		// config
		configHandler := config.New(r.logger, r.db, r.cache)
		api.PATCH("/config/email", configHandler.Email())

		// cron
		cronHandler := cron.New(r.logger, r.db, r.cache, r.cronServer)
		api.POST("/cron", cronHandler.Create())
		api.GET("/cron", cronHandler.List())
		api.GET("/cron/:id", core.AliasForRecordMetrics("/api/cron/detail"), cronHandler.Detail())
		api.POST("/cron/:id", core.AliasForRecordMetrics("/api/cron/modify"), cronHandler.Modify())
		api.PATCH("/cron/used", cronHandler.UpdateUsed())
		api.PATCH("/cron/exec/:id", core.AliasForRecordMetrics("/api/cron/exec"), cronHandler.Execute())

	}
}
