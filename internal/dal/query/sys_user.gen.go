// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/ch3nnn/webstack-go/internal/dal/model"
)

func newSysUser(db *gorm.DB, opts ...gen.DOOption) sysUser {
	_sysUser := sysUser{}

	_sysUser.sysUserDo.UseDB(db, opts...)
	_sysUser.sysUserDo.UseModel(&model.SysUser{})

	tableName := _sysUser.sysUserDo.TableName()
	_sysUser.ALL = field.NewAsterisk(tableName)
	_sysUser.ID = field.NewInt(tableName, "id")
	_sysUser.Username = field.NewString(tableName, "username")
	_sysUser.Password = field.NewString(tableName, "password")
	_sysUser.CreatedAt = field.NewTime(tableName, "created_at")
	_sysUser.UpdatedAt = field.NewTime(tableName, "updated_at")
	_sysUser.DeletedAt = field.NewTime(tableName, "deleted_at")

	_sysUser.fillFieldMap()

	return _sysUser
}

type sysUser struct {
	sysUserDo sysUserDo

	ALL       field.Asterisk
	ID        field.Int
	Username  field.String
	Password  field.String
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Time

	fieldMap map[string]field.Expr
}

func (s sysUser) Table(newTableName string) *sysUser {
	s.sysUserDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysUser) As(alias string) *sysUser {
	s.sysUserDo.DO = *(s.sysUserDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysUser) updateTableName(table string) *sysUser {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt(table, "id")
	s.Username = field.NewString(table, "username")
	s.Password = field.NewString(table, "password")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewTime(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *sysUser) WithContext(ctx context.Context) ISysUserDo { return s.sysUserDo.WithContext(ctx) }

func (s sysUser) TableName() string { return s.sysUserDo.TableName() }

func (s sysUser) Alias() string { return s.sysUserDo.Alias() }

func (s sysUser) Columns(cols ...field.Expr) gen.Columns { return s.sysUserDo.Columns(cols...) }

func (s *sysUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysUser) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 6)
	s.fieldMap["id"] = s.ID
	s.fieldMap["username"] = s.Username
	s.fieldMap["password"] = s.Password
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s sysUser) clone(db *gorm.DB) sysUser {
	s.sysUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysUser) replaceDB(db *gorm.DB) sysUser {
	s.sysUserDo.ReplaceDB(db)
	return s
}

type sysUserDo struct{ gen.DO }

type ISysUserDo interface {
	gen.SubQuery
	Debug() ISysUserDo
	WithContext(ctx context.Context) ISysUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysUserDo
	WriteDB() ISysUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysUserDo
	Not(conds ...gen.Condition) ISysUserDo
	Or(conds ...gen.Condition) ISysUserDo
	Select(conds ...field.Expr) ISysUserDo
	Where(conds ...gen.Condition) ISysUserDo
	Order(conds ...field.Expr) ISysUserDo
	Distinct(cols ...field.Expr) ISysUserDo
	Omit(cols ...field.Expr) ISysUserDo
	Join(table schema.Tabler, on ...field.Expr) ISysUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysUserDo
	Group(cols ...field.Expr) ISysUserDo
	Having(conds ...gen.Condition) ISysUserDo
	Limit(limit int) ISysUserDo
	Offset(offset int) ISysUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysUserDo
	Unscoped() ISysUserDo
	Create(values ...*model.SysUser) error
	CreateInBatches(values []*model.SysUser, batchSize int) error
	Save(values ...*model.SysUser) error
	First() (*model.SysUser, error)
	Take() (*model.SysUser, error)
	Last() (*model.SysUser, error)
	Find() ([]*model.SysUser, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysUser, err error)
	FindInBatches(result *[]*model.SysUser, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysUser) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysUserDo
	Assign(attrs ...field.AssignExpr) ISysUserDo
	Joins(fields ...field.RelationField) ISysUserDo
	Preload(fields ...field.RelationField) ISysUserDo
	FirstOrInit() (*model.SysUser, error)
	FirstOrCreate() (*model.SysUser, error)
	FindByPage(offset int, limit int) (result []*model.SysUser, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysUserDo) Debug() ISysUserDo {
	return s.withDO(s.DO.Debug())
}

func (s sysUserDo) WithContext(ctx context.Context) ISysUserDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysUserDo) ReadDB() ISysUserDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysUserDo) WriteDB() ISysUserDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysUserDo) Session(config *gorm.Session) ISysUserDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysUserDo) Clauses(conds ...clause.Expression) ISysUserDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysUserDo) Returning(value interface{}, columns ...string) ISysUserDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysUserDo) Not(conds ...gen.Condition) ISysUserDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysUserDo) Or(conds ...gen.Condition) ISysUserDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysUserDo) Select(conds ...field.Expr) ISysUserDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysUserDo) Where(conds ...gen.Condition) ISysUserDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysUserDo) Order(conds ...field.Expr) ISysUserDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysUserDo) Distinct(cols ...field.Expr) ISysUserDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysUserDo) Omit(cols ...field.Expr) ISysUserDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysUserDo) Join(table schema.Tabler, on ...field.Expr) ISysUserDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysUserDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysUserDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysUserDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysUserDo) Group(cols ...field.Expr) ISysUserDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysUserDo) Having(conds ...gen.Condition) ISysUserDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysUserDo) Limit(limit int) ISysUserDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysUserDo) Offset(offset int) ISysUserDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysUserDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysUserDo) Unscoped() ISysUserDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysUserDo) Create(values ...*model.SysUser) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysUserDo) CreateInBatches(values []*model.SysUser, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysUserDo) Save(values ...*model.SysUser) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysUserDo) First() (*model.SysUser, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUser), nil
	}
}

func (s sysUserDo) Take() (*model.SysUser, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUser), nil
	}
}

func (s sysUserDo) Last() (*model.SysUser, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUser), nil
	}
}

func (s sysUserDo) Find() ([]*model.SysUser, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysUser), err
}

func (s sysUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysUser, err error) {
	buf := make([]*model.SysUser, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysUserDo) FindInBatches(result *[]*model.SysUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysUserDo) Attrs(attrs ...field.AssignExpr) ISysUserDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysUserDo) Assign(attrs ...field.AssignExpr) ISysUserDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysUserDo) Joins(fields ...field.RelationField) ISysUserDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysUserDo) Preload(fields ...field.RelationField) ISysUserDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysUserDo) FirstOrInit() (*model.SysUser, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUser), nil
	}
}

func (s sysUserDo) FirstOrCreate() (*model.SysUser, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUser), nil
	}
}

func (s sysUserDo) FindByPage(offset int, limit int) (result []*model.SysUser, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysUserDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysUserDo) Delete(models ...*model.SysUser) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysUserDo) withDO(do gen.Dao) *sysUserDo {
	s.DO = *do.(*gen.DO)
	return s
}
