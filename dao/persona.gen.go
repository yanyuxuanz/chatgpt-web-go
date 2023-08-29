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

func newPersona(db *gorm.DB, opts ...gen.DOOption) persona {
	_persona := persona{}

	_persona.personaDo.UseDB(db, opts...)
	_persona.personaDo.UseModel(&model.Persona{})

	tableName := _persona.personaDo.TableName()
	_persona.ALL = field.NewAsterisk(tableName)
	_persona.ID = field.NewInt64(tableName, "id")
	_persona.UserID = field.NewInt64(tableName, "user_id")
	_persona.Title = field.NewString(tableName, "title")
	_persona.Avatar = field.NewString(tableName, "avatar")
	_persona.Description = field.NewString(tableName, "description")
	_persona.Context = field.NewString(tableName, "context")
	_persona.Status = field.NewInt32(tableName, "status")
	_persona.System = field.NewInt32(tableName, "system")
	_persona.CreateTime = field.NewTime(tableName, "create_time")
	_persona.UpdateTime = field.NewTime(tableName, "update_time")

	_persona.fillFieldMap()

	return _persona
}

type persona struct {
	personaDo

	ALL         field.Asterisk
	ID          field.Int64
	UserID      field.Int64
	Title       field.String
	Avatar      field.String
	Description field.String
	Context     field.String
	Status      field.Int32 // 1 正常 0异常 4审核中
	System      field.Int32 // 系统级别的角色
	CreateTime  field.Time
	UpdateTime  field.Time

	fieldMap map[string]field.Expr
}

func (p persona) Table(newTableName string) *persona {
	p.personaDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p persona) As(alias string) *persona {
	p.personaDo.DO = *(p.personaDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *persona) updateTableName(table string) *persona {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.UserID = field.NewInt64(table, "user_id")
	p.Title = field.NewString(table, "title")
	p.Avatar = field.NewString(table, "avatar")
	p.Description = field.NewString(table, "description")
	p.Context = field.NewString(table, "context")
	p.Status = field.NewInt32(table, "status")
	p.System = field.NewInt32(table, "system")
	p.CreateTime = field.NewTime(table, "create_time")
	p.UpdateTime = field.NewTime(table, "update_time")

	p.fillFieldMap()

	return p
}

func (p *persona) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *persona) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 10)
	p.fieldMap["id"] = p.ID
	p.fieldMap["user_id"] = p.UserID
	p.fieldMap["title"] = p.Title
	p.fieldMap["avatar"] = p.Avatar
	p.fieldMap["description"] = p.Description
	p.fieldMap["context"] = p.Context
	p.fieldMap["status"] = p.Status
	p.fieldMap["system"] = p.System
	p.fieldMap["create_time"] = p.CreateTime
	p.fieldMap["update_time"] = p.UpdateTime
}

func (p persona) clone(db *gorm.DB) persona {
	p.personaDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p persona) replaceDB(db *gorm.DB) persona {
	p.personaDo.ReplaceDB(db)
	return p
}

type personaDo struct{ gen.DO }

type IPersonaDo interface {
	gen.SubQuery
	Debug() IPersonaDo
	WithContext(ctx context.Context) IPersonaDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPersonaDo
	WriteDB() IPersonaDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPersonaDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPersonaDo
	Not(conds ...gen.Condition) IPersonaDo
	Or(conds ...gen.Condition) IPersonaDo
	Select(conds ...field.Expr) IPersonaDo
	Where(conds ...gen.Condition) IPersonaDo
	Order(conds ...field.Expr) IPersonaDo
	Distinct(cols ...field.Expr) IPersonaDo
	Omit(cols ...field.Expr) IPersonaDo
	Join(table schema.Tabler, on ...field.Expr) IPersonaDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPersonaDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPersonaDo
	Group(cols ...field.Expr) IPersonaDo
	Having(conds ...gen.Condition) IPersonaDo
	Limit(limit int) IPersonaDo
	Offset(offset int) IPersonaDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPersonaDo
	Unscoped() IPersonaDo
	Create(values ...*model.Persona) error
	CreateInBatches(values []*model.Persona, batchSize int) error
	Save(values ...*model.Persona) error
	First() (*model.Persona, error)
	Take() (*model.Persona, error)
	Last() (*model.Persona, error)
	Find() ([]*model.Persona, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Persona, err error)
	FindInBatches(result *[]*model.Persona, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Persona) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPersonaDo
	Assign(attrs ...field.AssignExpr) IPersonaDo
	Joins(fields ...field.RelationField) IPersonaDo
	Preload(fields ...field.RelationField) IPersonaDo
	FirstOrInit() (*model.Persona, error)
	FirstOrCreate() (*model.Persona, error)
	FindByPage(offset int, limit int) (result []*model.Persona, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPersonaDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	FilterWithNameAndRole(name string, role string) (result []model.Persona, err error)
}

// FilterWithNameAndRole SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
func (p personaDo) FilterWithNameAndRole(name string, role string) (result []model.Persona, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, name)
	generateSQL.WriteString("SELECT * FROM persona WHERE name = ? ")
	if role != "" {
		params = append(params, role)
		generateSQL.WriteString("AND role = ? ")
	}

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (p personaDo) Debug() IPersonaDo {
	return p.withDO(p.DO.Debug())
}

func (p personaDo) WithContext(ctx context.Context) IPersonaDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p personaDo) ReadDB() IPersonaDo {
	return p.Clauses(dbresolver.Read)
}

func (p personaDo) WriteDB() IPersonaDo {
	return p.Clauses(dbresolver.Write)
}

func (p personaDo) Session(config *gorm.Session) IPersonaDo {
	return p.withDO(p.DO.Session(config))
}

func (p personaDo) Clauses(conds ...clause.Expression) IPersonaDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p personaDo) Returning(value interface{}, columns ...string) IPersonaDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p personaDo) Not(conds ...gen.Condition) IPersonaDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p personaDo) Or(conds ...gen.Condition) IPersonaDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p personaDo) Select(conds ...field.Expr) IPersonaDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p personaDo) Where(conds ...gen.Condition) IPersonaDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p personaDo) Order(conds ...field.Expr) IPersonaDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p personaDo) Distinct(cols ...field.Expr) IPersonaDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p personaDo) Omit(cols ...field.Expr) IPersonaDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p personaDo) Join(table schema.Tabler, on ...field.Expr) IPersonaDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p personaDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPersonaDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p personaDo) RightJoin(table schema.Tabler, on ...field.Expr) IPersonaDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p personaDo) Group(cols ...field.Expr) IPersonaDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p personaDo) Having(conds ...gen.Condition) IPersonaDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p personaDo) Limit(limit int) IPersonaDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p personaDo) Offset(offset int) IPersonaDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p personaDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPersonaDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p personaDo) Unscoped() IPersonaDo {
	return p.withDO(p.DO.Unscoped())
}

func (p personaDo) Create(values ...*model.Persona) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p personaDo) CreateInBatches(values []*model.Persona, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p personaDo) Save(values ...*model.Persona) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p personaDo) First() (*model.Persona, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Persona), nil
	}
}

func (p personaDo) Take() (*model.Persona, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Persona), nil
	}
}

func (p personaDo) Last() (*model.Persona, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Persona), nil
	}
}

func (p personaDo) Find() ([]*model.Persona, error) {
	result, err := p.DO.Find()
	return result.([]*model.Persona), err
}

func (p personaDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Persona, err error) {
	buf := make([]*model.Persona, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p personaDo) FindInBatches(result *[]*model.Persona, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p personaDo) Attrs(attrs ...field.AssignExpr) IPersonaDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p personaDo) Assign(attrs ...field.AssignExpr) IPersonaDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p personaDo) Joins(fields ...field.RelationField) IPersonaDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p personaDo) Preload(fields ...field.RelationField) IPersonaDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p personaDo) FirstOrInit() (*model.Persona, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Persona), nil
	}
}

func (p personaDo) FirstOrCreate() (*model.Persona, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Persona), nil
	}
}

func (p personaDo) FindByPage(offset int, limit int) (result []*model.Persona, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p personaDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p personaDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p personaDo) Delete(models ...*model.Persona) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *personaDo) withDO(do gen.Dao) *personaDo {
	p.DO = *do.(*gen.DO)
	return p
}
