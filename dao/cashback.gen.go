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

func newCashback(db *gorm.DB, opts ...gen.DOOption) cashback {
	_cashback := cashback{}

	_cashback.cashbackDo.UseDB(db, opts...)
	_cashback.cashbackDo.UseModel(&model.Cashback{})

	tableName := _cashback.cashbackDo.TableName()
	_cashback.ALL = field.NewAsterisk(tableName)
	_cashback.ID = field.NewInt64(tableName, "id")
	_cashback.UserID = field.NewInt64(tableName, "user_id")
	_cashback.BenefitID = field.NewInt64(tableName, "benefit_id")
	_cashback.PayAmount = field.NewString(tableName, "pay_amount")
	_cashback.CommissionRate = field.NewString(tableName, "commission_rate")
	_cashback.CommissionAmount = field.NewString(tableName, "commission_amount")
	_cashback.Remarks = field.NewString(tableName, "remarks")
	_cashback.OrderID = field.NewInt64(tableName, "order_id")
	_cashback.Status = field.NewInt32(tableName, "status")
	_cashback.CreateTime = field.NewTime(tableName, "create_time")
	_cashback.UpdateTime = field.NewTime(tableName, "update_time")
	_cashback.IsDelete = field.NewInt32(tableName, "is_delete")

	_cashback.fillFieldMap()

	return _cashback
}

type cashback struct {
	cashbackDo

	ALL              field.Asterisk
	ID               field.Int64
	UserID           field.Int64
	BenefitID        field.Int64  // 受益者
	PayAmount        field.String // 支付金额（分）
	CommissionRate   field.String // 提成比例（1 - 10000）
	CommissionAmount field.String // 提成金额（分）
	Remarks          field.String // 评论
	OrderID          field.Int64
	Status           field.Int32 // 0异常 1正常 3审核中 6等待下发
	CreateTime       field.Time
	UpdateTime       field.Time
	IsDelete         field.Int32

	fieldMap map[string]field.Expr
}

func (c cashback) Table(newTableName string) *cashback {
	c.cashbackDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cashback) As(alias string) *cashback {
	c.cashbackDo.DO = *(c.cashbackDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cashback) updateTableName(table string) *cashback {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.UserID = field.NewInt64(table, "user_id")
	c.BenefitID = field.NewInt64(table, "benefit_id")
	c.PayAmount = field.NewString(table, "pay_amount")
	c.CommissionRate = field.NewString(table, "commission_rate")
	c.CommissionAmount = field.NewString(table, "commission_amount")
	c.Remarks = field.NewString(table, "remarks")
	c.OrderID = field.NewInt64(table, "order_id")
	c.Status = field.NewInt32(table, "status")
	c.CreateTime = field.NewTime(table, "create_time")
	c.UpdateTime = field.NewTime(table, "update_time")
	c.IsDelete = field.NewInt32(table, "is_delete")

	c.fillFieldMap()

	return c
}

func (c *cashback) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cashback) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 12)
	c.fieldMap["id"] = c.ID
	c.fieldMap["user_id"] = c.UserID
	c.fieldMap["benefit_id"] = c.BenefitID
	c.fieldMap["pay_amount"] = c.PayAmount
	c.fieldMap["commission_rate"] = c.CommissionRate
	c.fieldMap["commission_amount"] = c.CommissionAmount
	c.fieldMap["remarks"] = c.Remarks
	c.fieldMap["order_id"] = c.OrderID
	c.fieldMap["status"] = c.Status
	c.fieldMap["create_time"] = c.CreateTime
	c.fieldMap["update_time"] = c.UpdateTime
	c.fieldMap["is_delete"] = c.IsDelete
}

func (c cashback) clone(db *gorm.DB) cashback {
	c.cashbackDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cashback) replaceDB(db *gorm.DB) cashback {
	c.cashbackDo.ReplaceDB(db)
	return c
}

type cashbackDo struct{ gen.DO }

type ICashbackDo interface {
	gen.SubQuery
	Debug() ICashbackDo
	WithContext(ctx context.Context) ICashbackDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICashbackDo
	WriteDB() ICashbackDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICashbackDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICashbackDo
	Not(conds ...gen.Condition) ICashbackDo
	Or(conds ...gen.Condition) ICashbackDo
	Select(conds ...field.Expr) ICashbackDo
	Where(conds ...gen.Condition) ICashbackDo
	Order(conds ...field.Expr) ICashbackDo
	Distinct(cols ...field.Expr) ICashbackDo
	Omit(cols ...field.Expr) ICashbackDo
	Join(table schema.Tabler, on ...field.Expr) ICashbackDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICashbackDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICashbackDo
	Group(cols ...field.Expr) ICashbackDo
	Having(conds ...gen.Condition) ICashbackDo
	Limit(limit int) ICashbackDo
	Offset(offset int) ICashbackDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICashbackDo
	Unscoped() ICashbackDo
	Create(values ...*model.Cashback) error
	CreateInBatches(values []*model.Cashback, batchSize int) error
	Save(values ...*model.Cashback) error
	First() (*model.Cashback, error)
	Take() (*model.Cashback, error)
	Last() (*model.Cashback, error)
	Find() ([]*model.Cashback, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Cashback, err error)
	FindInBatches(result *[]*model.Cashback, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Cashback) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICashbackDo
	Assign(attrs ...field.AssignExpr) ICashbackDo
	Joins(fields ...field.RelationField) ICashbackDo
	Preload(fields ...field.RelationField) ICashbackDo
	FirstOrInit() (*model.Cashback, error)
	FirstOrCreate() (*model.Cashback, error)
	FindByPage(offset int, limit int) (result []*model.Cashback, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICashbackDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	FilterWithNameAndRole(name string, role string) (result []model.Cashback, err error)
}

// FilterWithNameAndRole SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
func (c cashbackDo) FilterWithNameAndRole(name string, role string) (result []model.Cashback, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, name)
	generateSQL.WriteString("SELECT * FROM cashback WHERE name = ? ")
	if role != "" {
		params = append(params, role)
		generateSQL.WriteString("AND role = ? ")
	}

	var executeSQL *gorm.DB
	executeSQL = c.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (c cashbackDo) Debug() ICashbackDo {
	return c.withDO(c.DO.Debug())
}

func (c cashbackDo) WithContext(ctx context.Context) ICashbackDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cashbackDo) ReadDB() ICashbackDo {
	return c.Clauses(dbresolver.Read)
}

func (c cashbackDo) WriteDB() ICashbackDo {
	return c.Clauses(dbresolver.Write)
}

func (c cashbackDo) Session(config *gorm.Session) ICashbackDo {
	return c.withDO(c.DO.Session(config))
}

func (c cashbackDo) Clauses(conds ...clause.Expression) ICashbackDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cashbackDo) Returning(value interface{}, columns ...string) ICashbackDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cashbackDo) Not(conds ...gen.Condition) ICashbackDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cashbackDo) Or(conds ...gen.Condition) ICashbackDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cashbackDo) Select(conds ...field.Expr) ICashbackDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cashbackDo) Where(conds ...gen.Condition) ICashbackDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cashbackDo) Order(conds ...field.Expr) ICashbackDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cashbackDo) Distinct(cols ...field.Expr) ICashbackDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cashbackDo) Omit(cols ...field.Expr) ICashbackDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cashbackDo) Join(table schema.Tabler, on ...field.Expr) ICashbackDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cashbackDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICashbackDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cashbackDo) RightJoin(table schema.Tabler, on ...field.Expr) ICashbackDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cashbackDo) Group(cols ...field.Expr) ICashbackDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cashbackDo) Having(conds ...gen.Condition) ICashbackDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cashbackDo) Limit(limit int) ICashbackDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cashbackDo) Offset(offset int) ICashbackDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cashbackDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICashbackDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cashbackDo) Unscoped() ICashbackDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cashbackDo) Create(values ...*model.Cashback) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cashbackDo) CreateInBatches(values []*model.Cashback, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cashbackDo) Save(values ...*model.Cashback) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cashbackDo) First() (*model.Cashback, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Cashback), nil
	}
}

func (c cashbackDo) Take() (*model.Cashback, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Cashback), nil
	}
}

func (c cashbackDo) Last() (*model.Cashback, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Cashback), nil
	}
}

func (c cashbackDo) Find() ([]*model.Cashback, error) {
	result, err := c.DO.Find()
	return result.([]*model.Cashback), err
}

func (c cashbackDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Cashback, err error) {
	buf := make([]*model.Cashback, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cashbackDo) FindInBatches(result *[]*model.Cashback, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cashbackDo) Attrs(attrs ...field.AssignExpr) ICashbackDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cashbackDo) Assign(attrs ...field.AssignExpr) ICashbackDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cashbackDo) Joins(fields ...field.RelationField) ICashbackDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cashbackDo) Preload(fields ...field.RelationField) ICashbackDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cashbackDo) FirstOrInit() (*model.Cashback, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Cashback), nil
	}
}

func (c cashbackDo) FirstOrCreate() (*model.Cashback, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Cashback), nil
	}
}

func (c cashbackDo) FindByPage(offset int, limit int) (result []*model.Cashback, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c cashbackDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cashbackDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cashbackDo) Delete(models ...*model.Cashback) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cashbackDo) withDO(do gen.Dao) *cashbackDo {
	c.DO = *do.(*gen.DO)
	return c
}
