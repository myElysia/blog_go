// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"blogGo/src/model"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newUserGroup(db *gorm.DB, opts ...gen.DOOption) userGroup {
	_userGroup := userGroup{}

	_userGroup.userGroupDo.UseDB(db, opts...)
	_userGroup.userGroupDo.UseModel(&model.UserGroup{})

	tableName := _userGroup.userGroupDo.TableName()
	_userGroup.ALL = field.NewAsterisk(tableName)
	_userGroup.ID = field.NewUint(tableName, "id")
	_userGroup.CreatedAt = field.NewTime(tableName, "created_at")
	_userGroup.UpdatedAt = field.NewTime(tableName, "updated_at")
	_userGroup.DeletedAt = field.NewField(tableName, "deleted_at")
	_userGroup.Name = field.NewString(tableName, "name")
	_userGroup.Roles = userGroupManyToManyRoles{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Roles", "model.Role"),
		Permissions: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Roles.Permissions", "model.Permission"),
		},
	}

	_userGroup.fillFieldMap()

	return _userGroup
}

type userGroup struct {
	userGroupDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Name      field.String
	Roles     userGroupManyToManyRoles

	fieldMap map[string]field.Expr
}

func (u userGroup) Table(newTableName string) *userGroup {
	u.userGroupDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userGroup) As(alias string) *userGroup {
	u.userGroupDo.DO = *(u.userGroupDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userGroup) updateTableName(table string) *userGroup {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewUint(table, "id")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")
	u.DeletedAt = field.NewField(table, "deleted_at")
	u.Name = field.NewString(table, "name")

	u.fillFieldMap()

	return u
}

func (u *userGroup) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userGroup) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 6)
	u.fieldMap["id"] = u.ID
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
	u.fieldMap["name"] = u.Name

}

func (u userGroup) clone(db *gorm.DB) userGroup {
	u.userGroupDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userGroup) replaceDB(db *gorm.DB) userGroup {
	u.userGroupDo.ReplaceDB(db)
	return u
}

type userGroupManyToManyRoles struct {
	db *gorm.DB

	field.RelationField

	Permissions struct {
		field.RelationField
	}
}

func (a userGroupManyToManyRoles) Where(conds ...field.Expr) *userGroupManyToManyRoles {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a userGroupManyToManyRoles) WithContext(ctx context.Context) *userGroupManyToManyRoles {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a userGroupManyToManyRoles) Session(session *gorm.Session) *userGroupManyToManyRoles {
	a.db = a.db.Session(session)
	return &a
}

func (a userGroupManyToManyRoles) Model(m *model.UserGroup) *userGroupManyToManyRolesTx {
	return &userGroupManyToManyRolesTx{a.db.Model(m).Association(a.Name())}
}

type userGroupManyToManyRolesTx struct{ tx *gorm.Association }

func (a userGroupManyToManyRolesTx) Find() (result []*model.Role, err error) {
	return result, a.tx.Find(&result)
}

func (a userGroupManyToManyRolesTx) Append(values ...*model.Role) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a userGroupManyToManyRolesTx) Replace(values ...*model.Role) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a userGroupManyToManyRolesTx) Delete(values ...*model.Role) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a userGroupManyToManyRolesTx) Clear() error {
	return a.tx.Clear()
}

func (a userGroupManyToManyRolesTx) Count() int64 {
	return a.tx.Count()
}

type userGroupDo struct{ gen.DO }

type IUserGroupDo interface {
	gen.SubQuery
	Debug() IUserGroupDo
	WithContext(ctx context.Context) IUserGroupDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserGroupDo
	WriteDB() IUserGroupDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserGroupDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserGroupDo
	Not(conds ...gen.Condition) IUserGroupDo
	Or(conds ...gen.Condition) IUserGroupDo
	Select(conds ...field.Expr) IUserGroupDo
	Where(conds ...gen.Condition) IUserGroupDo
	Order(conds ...field.Expr) IUserGroupDo
	Distinct(cols ...field.Expr) IUserGroupDo
	Omit(cols ...field.Expr) IUserGroupDo
	Join(table schema.Tabler, on ...field.Expr) IUserGroupDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserGroupDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserGroupDo
	Group(cols ...field.Expr) IUserGroupDo
	Having(conds ...gen.Condition) IUserGroupDo
	Limit(limit int) IUserGroupDo
	Offset(offset int) IUserGroupDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserGroupDo
	Unscoped() IUserGroupDo
	Create(values ...*model.UserGroup) error
	CreateInBatches(values []*model.UserGroup, batchSize int) error
	Save(values ...*model.UserGroup) error
	First() (*model.UserGroup, error)
	Take() (*model.UserGroup, error)
	Last() (*model.UserGroup, error)
	Find() ([]*model.UserGroup, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserGroup, err error)
	FindInBatches(result *[]*model.UserGroup, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserGroup) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserGroupDo
	Assign(attrs ...field.AssignExpr) IUserGroupDo
	Joins(fields ...field.RelationField) IUserGroupDo
	Preload(fields ...field.RelationField) IUserGroupDo
	FirstOrInit() (*model.UserGroup, error)
	FirstOrCreate() (*model.UserGroup, error)
	FindByPage(offset int, limit int) (result []*model.UserGroup, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserGroupDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userGroupDo) Debug() IUserGroupDo {
	return u.withDO(u.DO.Debug())
}

func (u userGroupDo) WithContext(ctx context.Context) IUserGroupDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userGroupDo) ReadDB() IUserGroupDo {
	return u.Clauses(dbresolver.Read)
}

func (u userGroupDo) WriteDB() IUserGroupDo {
	return u.Clauses(dbresolver.Write)
}

func (u userGroupDo) Session(config *gorm.Session) IUserGroupDo {
	return u.withDO(u.DO.Session(config))
}

func (u userGroupDo) Clauses(conds ...clause.Expression) IUserGroupDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userGroupDo) Returning(value interface{}, columns ...string) IUserGroupDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userGroupDo) Not(conds ...gen.Condition) IUserGroupDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userGroupDo) Or(conds ...gen.Condition) IUserGroupDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userGroupDo) Select(conds ...field.Expr) IUserGroupDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userGroupDo) Where(conds ...gen.Condition) IUserGroupDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userGroupDo) Order(conds ...field.Expr) IUserGroupDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userGroupDo) Distinct(cols ...field.Expr) IUserGroupDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userGroupDo) Omit(cols ...field.Expr) IUserGroupDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userGroupDo) Join(table schema.Tabler, on ...field.Expr) IUserGroupDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userGroupDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserGroupDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userGroupDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserGroupDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userGroupDo) Group(cols ...field.Expr) IUserGroupDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userGroupDo) Having(conds ...gen.Condition) IUserGroupDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userGroupDo) Limit(limit int) IUserGroupDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userGroupDo) Offset(offset int) IUserGroupDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userGroupDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserGroupDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userGroupDo) Unscoped() IUserGroupDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userGroupDo) Create(values ...*model.UserGroup) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userGroupDo) CreateInBatches(values []*model.UserGroup, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userGroupDo) Save(values ...*model.UserGroup) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userGroupDo) First() (*model.UserGroup, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserGroup), nil
	}
}

func (u userGroupDo) Take() (*model.UserGroup, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserGroup), nil
	}
}

func (u userGroupDo) Last() (*model.UserGroup, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserGroup), nil
	}
}

func (u userGroupDo) Find() ([]*model.UserGroup, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserGroup), err
}

func (u userGroupDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserGroup, err error) {
	buf := make([]*model.UserGroup, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userGroupDo) FindInBatches(result *[]*model.UserGroup, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userGroupDo) Attrs(attrs ...field.AssignExpr) IUserGroupDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userGroupDo) Assign(attrs ...field.AssignExpr) IUserGroupDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userGroupDo) Joins(fields ...field.RelationField) IUserGroupDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userGroupDo) Preload(fields ...field.RelationField) IUserGroupDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userGroupDo) FirstOrInit() (*model.UserGroup, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserGroup), nil
	}
}

func (u userGroupDo) FirstOrCreate() (*model.UserGroup, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserGroup), nil
	}
}

func (u userGroupDo) FindByPage(offset int, limit int) (result []*model.UserGroup, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userGroupDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userGroupDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userGroupDo) Delete(models ...*model.UserGroup) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userGroupDo) withDO(do gen.Dao) *userGroupDo {
	u.DO = *do.(*gen.DO)
	return u
}
