package gormgen

import (
	"github.com/ch3nnn/webstack-go/internal/repository/mysql"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

// Method 根据动态 sql 生成 orm
type Method interface {
	// GetParentIdsByGroupParentId
	// SELECT GROUP_CONCAT(a.parent_id) AS parent_ids from (SELECT parent_id  FROM category  GROUP BY parent_id)  as a
	GetParentIdsByGroupParentId() (gen.M, error)
}

func GenerateTable(tables []string) {

	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/repository/mysql/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// reuse your gorm db
	repo, _ := mysql.New()
	g.UseDB(repo.GetDbW())

	// 自定义字段的数据类型
	// 统一数字类型为int64,兼容protobuf (columnType gorm.ColumnType) (dataType string)
	dataMap := map[string]func(columnType gorm.ColumnType) (dataType string){
		"tinyint":   func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"smallint":  func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"mediumint": func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"bigint":    func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"int":       func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"integer":   func(columnType gorm.ColumnType) (dataType string) { return "int64" },
	}
	// 要先于`ApplyBasic`执行
	g.WithDataTypeMap(dataMap)

	// 将非默认字段名的字段定义为自动时间戳和软删除字段;
	// 自动时间戳默认字段名为:`updated_at`、`created_at, 表字段数据类型为: INT 或 DATETIME
	// 软删除默认字段名为:`deleted_at`, 表字段数据类型为: DATETIME
	//autoUpdateTimeField := gen.FieldGORMTag("update_time", func(tag field.GormTag) field.GormTag {return })
	//autoCreateTimeField := gen.FieldGORMTag("create_time", "column:create_time;type:int unsigned;autoCreateTime")

	// 模型自定义选项组
	//fieldOpts := []gen.ModelOpt{autoCreateTimeField, autoUpdateTimeField}

	// 分类
	category := g.GenerateModel("category")
	g.ApplyInterface(func(Method) {}, model.Category{})
	// 网址
	site := g.GenerateModel("site", gen.FieldRelate(field.BelongsTo, "Category", category, &field.RelateConfig{
		GORMTag: map[string]string{"foreignKey": "CategoryID"},
	}),
	)
	// 菜单
	menu := g.GenerateModel("menu")
	// 菜单动作
	menuAction := g.GenerateModel("menu_action")
	// 管理员菜单栏表
	adminMenu := g.GenerateModel("admin_menu")
	// 管理员表
	admin := g.GenerateModel("admin")
	// 定时任务表
	cronTask := g.GenerateModel("cron_task")
	// 已授权接口地址表
	authorizedApi := g.GenerateModel("authorized_api")
	// 已授权的调用方表
	authorized := g.GenerateModel("authorized")
	// 自定义表结构生成
	for _, table := range tables {
		g.ApplyBasic(g.GenerateModel(table))
	}

	// 创建全部模型文件, 并覆盖前面创建的同名模型
	//allModel := g.GenerateAllTable(fieldOpts...)

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(site, menu, menuAction, adminMenu, admin, cronTask, authorizedApi, authorized)
	//g.ApplyBasic(allModel)

	// Generate the code
	g.Execute()
}

//func main() {
//	GenerateTable(nil)
//}
