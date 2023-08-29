// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"chatgpt-web-new-go/model"
)

func newInstalledPlugin(db *gorm.DB, opts ...gen.DOOption) installedPlugin {
	_installedPlugin := installedPlugin{}

	_installedPlugin.installedPluginDo.UseDB(db, opts...)
	_installedPlugin.installedPluginDo.UseModel(&model.InstalledPlugin{})

	tableName := _installedPlugin.installedPluginDo.TableName()
	_installedPlugin.ALL = field.NewAsterisk(tableName)
	_installedPlugin.ID = field.NewInt64(tableName, "id")
	_installedPlugin.UserID = field.NewInt64(tableName, "user_id")
	_installedPlugin.PluginID = field.NewInt64(tableName, "plugin_id")
	_installedPlugin.Status = field.NewInt32(tableName, "status")
	_installedPlugin.CreateTime = field.NewTime(tableName, "create_time")
	_installedPlugin.UpdateTime = field.NewTime(tableName, "update_time")

	_installedPlugin.fillFieldMap()

	return _installedPlugin
}

type installedPlugin struct {
	installedPluginDo

	ALL        field.Asterisk
	ID         field.Int64
	UserID     field.Int64
	PluginID   field.Int64
	Status     field.Int32
	CreateTime field.Time
	UpdateTime field.Time

	fieldMap map[string]field.Expr
}

func (i installedPlugin) Table(newTableName string) *installedPlugin {
	i.installedPluginDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i installedPlugin) As(alias string) *installedPlugin {
	i.installedPluginDo.DO = *(i.installedPluginDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *installedPlugin) updateTableName(table string) *installedPlugin {
	i.ALL = field.NewAsterisk(table)
	i.ID = field.NewInt64(table, "id")
	i.UserID = field.NewInt64(table, "user_id")
	i.PluginID = field.NewInt64(table, "plugin_id")
	i.Status = field.NewInt32(table, "status")
	i.CreateTime = field.NewTime(table, "create_time")
	i.UpdateTime = field.NewTime(table, "update_time")

	i.fillFieldMap()

	return i
}

func (i *installedPlugin) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *installedPlugin) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 6)
	i.fieldMap["id"] = i.ID
	i.fieldMap["user_id"] = i.UserID
	i.fieldMap["plugin_id"] = i.PluginID
	i.fieldMap["status"] = i.Status
	i.fieldMap["create_time"] = i.CreateTime
	i.fieldMap["update_time"] = i.UpdateTime
}

func (i installedPlugin) clone(db *gorm.DB) installedPlugin {
	i.installedPluginDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i installedPlugin) replaceDB(db *gorm.DB) installedPlugin {
	i.installedPluginDo.ReplaceDB(db)
	return i
}

type installedPluginDo struct{ gen.DO }

type IInstalledPluginDo interface {
	gen.SubQuery
	Debug() IInstalledPluginDo
	WithContext(ctx context.Context) IInstalledPluginDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IInstalledPluginDo
	WriteDB() IInstalledPluginDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IInstalledPluginDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IInstalledPluginDo
	Not(conds ...gen.Condition) IInstalledPluginDo
	Or(conds ...gen.Condition) IInstalledPluginDo
	Select(conds ...field.Expr) IInstalledPluginDo
	Where(conds ...gen.Condition) IInstalledPluginDo
	Order(conds ...field.Expr) IInstalledPluginDo
	Distinct(cols ...field.Expr) IInstalledPluginDo
	Omit(cols ...field.Expr) IInstalledPluginDo
	Join(table schema.Tabler, on ...field.Expr) IInstalledPluginDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IInstalledPluginDo
	RightJoin(table schema.Tabler, on ...field.Expr) IInstalledPluginDo
	Group(cols ...field.Expr) IInstalledPluginDo
	Having(conds ...gen.Condition) IInstalledPluginDo
	Limit(limit int) IInstalledPluginDo
	Offset(offset int) IInstalledPluginDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IInstalledPluginDo
	Unscoped() IInstalledPluginDo
	Create(values ...*model.InstalledPlugin) error
	CreateInBatches(values []*model.InstalledPlugin, batchSize int) error
	Save(values ...*model.InstalledPlugin) error
	First() (*model.InstalledPlugin, error)
	Take() (*model.InstalledPlugin, error)
	Last() (*model.InstalledPlugin, error)
	Find() ([]*model.InstalledPlugin, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.InstalledPlugin, err error)
	FindInBatches(result *[]*model.InstalledPlugin, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.InstalledPlugin) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IInstalledPluginDo
	Assign(attrs ...field.AssignExpr) IInstalledPluginDo
	Joins(fields ...field.RelationField) IInstalledPluginDo
	Preload(fields ...field.RelationField) IInstalledPluginDo
	FirstOrInit() (*model.InstalledPlugin, error)
	FirstOrCreate() (*model.InstalledPlugin, error)
	FindByPage(offset int, limit int) (result []*model.InstalledPlugin, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IInstalledPluginDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	FilterWithNameAndRole(name string, role string) (result []model.InstalledPlugin, err error)
}

// FilterWithNameAndRole SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
func (i installedPluginDo) FilterWithNameAndRole(name string, role string) (result []model.InstalledPlugin, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, name)
	generateSQL.WriteString("SELECT * FROM installed_plugin WHERE name = ? ")
	if role != "" {
		params = append(params, role)
		generateSQL.WriteString("AND role = ? ")
	}

	var executeSQL *gorm.DB
	executeSQL = i.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (i installedPluginDo) Debug() IInstalledPluginDo {
	return i.withDO(i.DO.Debug())
}

func (i installedPluginDo) WithContext(ctx context.Context) IInstalledPluginDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i installedPluginDo) ReadDB() IInstalledPluginDo {
	return i.Clauses(dbresolver.Read)
}

func (i installedPluginDo) WriteDB() IInstalledPluginDo {
	return i.Clauses(dbresolver.Write)
}

func (i installedPluginDo) Session(config *gorm.Session) IInstalledPluginDo {
	return i.withDO(i.DO.Session(config))
}

func (i installedPluginDo) Clauses(conds ...clause.Expression) IInstalledPluginDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i installedPluginDo) Returning(value interface{}, columns ...string) IInstalledPluginDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i installedPluginDo) Not(conds ...gen.Condition) IInstalledPluginDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i installedPluginDo) Or(conds ...gen.Condition) IInstalledPluginDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i installedPluginDo) Select(conds ...field.Expr) IInstalledPluginDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i installedPluginDo) Where(conds ...gen.Condition) IInstalledPluginDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i installedPluginDo) Order(conds ...field.Expr) IInstalledPluginDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i installedPluginDo) Distinct(cols ...field.Expr) IInstalledPluginDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i installedPluginDo) Omit(cols ...field.Expr) IInstalledPluginDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i installedPluginDo) Join(table schema.Tabler, on ...field.Expr) IInstalledPluginDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i installedPluginDo) LeftJoin(table schema.Tabler, on ...field.Expr) IInstalledPluginDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i installedPluginDo) RightJoin(table schema.Tabler, on ...field.Expr) IInstalledPluginDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i installedPluginDo) Group(cols ...field.Expr) IInstalledPluginDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i installedPluginDo) Having(conds ...gen.Condition) IInstalledPluginDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i installedPluginDo) Limit(limit int) IInstalledPluginDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i installedPluginDo) Offset(offset int) IInstalledPluginDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i installedPluginDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IInstalledPluginDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i installedPluginDo) Unscoped() IInstalledPluginDo {
	return i.withDO(i.DO.Unscoped())
}

func (i installedPluginDo) Create(values ...*model.InstalledPlugin) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i installedPluginDo) CreateInBatches(values []*model.InstalledPlugin, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i installedPluginDo) Save(values ...*model.InstalledPlugin) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i installedPluginDo) First() (*model.InstalledPlugin, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstalledPlugin), nil
	}
}

func (i installedPluginDo) Take() (*model.InstalledPlugin, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstalledPlugin), nil
	}
}

func (i installedPluginDo) Last() (*model.InstalledPlugin, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstalledPlugin), nil
	}
}

func (i installedPluginDo) Find() ([]*model.InstalledPlugin, error) {
	result, err := i.DO.Find()
	return result.([]*model.InstalledPlugin), err
}

func (i installedPluginDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.InstalledPlugin, err error) {
	buf := make([]*model.InstalledPlugin, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i installedPluginDo) FindInBatches(result *[]*model.InstalledPlugin, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i installedPluginDo) Attrs(attrs ...field.AssignExpr) IInstalledPluginDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i installedPluginDo) Assign(attrs ...field.AssignExpr) IInstalledPluginDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i installedPluginDo) Joins(fields ...field.RelationField) IInstalledPluginDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i installedPluginDo) Preload(fields ...field.RelationField) IInstalledPluginDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i installedPluginDo) FirstOrInit() (*model.InstalledPlugin, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstalledPlugin), nil
	}
}

func (i installedPluginDo) FirstOrCreate() (*model.InstalledPlugin, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstalledPlugin), nil
	}
}

func (i installedPluginDo) FindByPage(offset int, limit int) (result []*model.InstalledPlugin, count int64, err error) {
	result, err = i.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = i.Offset(-1).Limit(-1).Count()
	return
}

func (i installedPluginDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i installedPluginDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i installedPluginDo) Delete(models ...*model.InstalledPlugin) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *installedPluginDo) withDO(do gen.Dao) *installedPluginDo {
	i.DO = *do.(*gen.DO)
	return i
}
