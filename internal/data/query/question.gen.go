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

	"wuzigoweb/internal/data/model"
)

func newQuestion(db *gorm.DB, opts ...gen.DOOption) question {
	_question := question{}

	_question.questionDo.UseDB(db, opts...)
	_question.questionDo.UseModel(&model.Question{})

	tableName := _question.questionDo.TableName()
	_question.ALL = field.NewAsterisk(tableName)
	_question.ID = field.NewInt64(tableName, "id")
	_question.QuestionContent = field.NewString(tableName, "question_content")
	_question.AppID = field.NewInt64(tableName, "app_id")
	_question.UserID = field.NewInt64(tableName, "user_id")
	_question.CreateTime = field.NewTime(tableName, "create_time")
	_question.UpdateTime = field.NewTime(tableName, "update_time")
	_question.DeleteAt = field.NewInt32(tableName, "delete_at")

	_question.fillFieldMap()

	return _question
}

// question 题目
type question struct {
	questionDo questionDo

	ALL             field.Asterisk
	ID              field.Int64  // id
	QuestionContent field.String // 题目内容（json格式）
	AppID           field.Int64  // 应用 id
	UserID          field.Int64  // 创建用户 id
	CreateTime      field.Time   // 创建时间
	UpdateTime      field.Time   // 更新时间
	DeleteAt        field.Int32  // 是否删除

	fieldMap map[string]field.Expr
}

func (q question) Table(newTableName string) *question {
	q.questionDo.UseTable(newTableName)
	return q.updateTableName(newTableName)
}

func (q question) As(alias string) *question {
	q.questionDo.DO = *(q.questionDo.As(alias).(*gen.DO))
	return q.updateTableName(alias)
}

func (q *question) updateTableName(table string) *question {
	q.ALL = field.NewAsterisk(table)
	q.ID = field.NewInt64(table, "id")
	q.QuestionContent = field.NewString(table, "question_content")
	q.AppID = field.NewInt64(table, "app_id")
	q.UserID = field.NewInt64(table, "user_id")
	q.CreateTime = field.NewTime(table, "create_time")
	q.UpdateTime = field.NewTime(table, "update_time")
	q.DeleteAt = field.NewInt32(table, "delete_at")

	q.fillFieldMap()

	return q
}

func (q *question) WithContext(ctx context.Context) IQuestionDo { return q.questionDo.WithContext(ctx) }

func (q question) TableName() string { return q.questionDo.TableName() }

func (q question) Alias() string { return q.questionDo.Alias() }

func (q question) Columns(cols ...field.Expr) gen.Columns { return q.questionDo.Columns(cols...) }

func (q *question) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := q.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (q *question) fillFieldMap() {
	q.fieldMap = make(map[string]field.Expr, 7)
	q.fieldMap["id"] = q.ID
	q.fieldMap["question_content"] = q.QuestionContent
	q.fieldMap["app_id"] = q.AppID
	q.fieldMap["user_id"] = q.UserID
	q.fieldMap["create_time"] = q.CreateTime
	q.fieldMap["update_time"] = q.UpdateTime
	q.fieldMap["delete_at"] = q.DeleteAt
}

func (q question) clone(db *gorm.DB) question {
	q.questionDo.ReplaceConnPool(db.Statement.ConnPool)
	return q
}

func (q question) replaceDB(db *gorm.DB) question {
	q.questionDo.ReplaceDB(db)
	return q
}

type questionDo struct{ gen.DO }

type IQuestionDo interface {
	gen.SubQuery
	Debug() IQuestionDo
	WithContext(ctx context.Context) IQuestionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IQuestionDo
	WriteDB() IQuestionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IQuestionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IQuestionDo
	Not(conds ...gen.Condition) IQuestionDo
	Or(conds ...gen.Condition) IQuestionDo
	Select(conds ...field.Expr) IQuestionDo
	Where(conds ...gen.Condition) IQuestionDo
	Order(conds ...field.Expr) IQuestionDo
	Distinct(cols ...field.Expr) IQuestionDo
	Omit(cols ...field.Expr) IQuestionDo
	Join(table schema.Tabler, on ...field.Expr) IQuestionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IQuestionDo
	RightJoin(table schema.Tabler, on ...field.Expr) IQuestionDo
	Group(cols ...field.Expr) IQuestionDo
	Having(conds ...gen.Condition) IQuestionDo
	Limit(limit int) IQuestionDo
	Offset(offset int) IQuestionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IQuestionDo
	Unscoped() IQuestionDo
	Create(values ...*model.Question) error
	CreateInBatches(values []*model.Question, batchSize int) error
	Save(values ...*model.Question) error
	First() (*model.Question, error)
	Take() (*model.Question, error)
	Last() (*model.Question, error)
	Find() ([]*model.Question, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Question, err error)
	FindInBatches(result *[]*model.Question, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Question) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IQuestionDo
	Assign(attrs ...field.AssignExpr) IQuestionDo
	Joins(fields ...field.RelationField) IQuestionDo
	Preload(fields ...field.RelationField) IQuestionDo
	FirstOrInit() (*model.Question, error)
	FirstOrCreate() (*model.Question, error)
	FindByPage(offset int, limit int) (result []*model.Question, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IQuestionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (q questionDo) Debug() IQuestionDo {
	return q.withDO(q.DO.Debug())
}

func (q questionDo) WithContext(ctx context.Context) IQuestionDo {
	return q.withDO(q.DO.WithContext(ctx))
}

func (q questionDo) ReadDB() IQuestionDo {
	return q.Clauses(dbresolver.Read)
}

func (q questionDo) WriteDB() IQuestionDo {
	return q.Clauses(dbresolver.Write)
}

func (q questionDo) Session(config *gorm.Session) IQuestionDo {
	return q.withDO(q.DO.Session(config))
}

func (q questionDo) Clauses(conds ...clause.Expression) IQuestionDo {
	return q.withDO(q.DO.Clauses(conds...))
}

func (q questionDo) Returning(value interface{}, columns ...string) IQuestionDo {
	return q.withDO(q.DO.Returning(value, columns...))
}

func (q questionDo) Not(conds ...gen.Condition) IQuestionDo {
	return q.withDO(q.DO.Not(conds...))
}

func (q questionDo) Or(conds ...gen.Condition) IQuestionDo {
	return q.withDO(q.DO.Or(conds...))
}

func (q questionDo) Select(conds ...field.Expr) IQuestionDo {
	return q.withDO(q.DO.Select(conds...))
}

func (q questionDo) Where(conds ...gen.Condition) IQuestionDo {
	return q.withDO(q.DO.Where(conds...))
}

func (q questionDo) Order(conds ...field.Expr) IQuestionDo {
	return q.withDO(q.DO.Order(conds...))
}

func (q questionDo) Distinct(cols ...field.Expr) IQuestionDo {
	return q.withDO(q.DO.Distinct(cols...))
}

func (q questionDo) Omit(cols ...field.Expr) IQuestionDo {
	return q.withDO(q.DO.Omit(cols...))
}

func (q questionDo) Join(table schema.Tabler, on ...field.Expr) IQuestionDo {
	return q.withDO(q.DO.Join(table, on...))
}

func (q questionDo) LeftJoin(table schema.Tabler, on ...field.Expr) IQuestionDo {
	return q.withDO(q.DO.LeftJoin(table, on...))
}

func (q questionDo) RightJoin(table schema.Tabler, on ...field.Expr) IQuestionDo {
	return q.withDO(q.DO.RightJoin(table, on...))
}

func (q questionDo) Group(cols ...field.Expr) IQuestionDo {
	return q.withDO(q.DO.Group(cols...))
}

func (q questionDo) Having(conds ...gen.Condition) IQuestionDo {
	return q.withDO(q.DO.Having(conds...))
}

func (q questionDo) Limit(limit int) IQuestionDo {
	return q.withDO(q.DO.Limit(limit))
}

func (q questionDo) Offset(offset int) IQuestionDo {
	return q.withDO(q.DO.Offset(offset))
}

func (q questionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IQuestionDo {
	return q.withDO(q.DO.Scopes(funcs...))
}

func (q questionDo) Unscoped() IQuestionDo {
	return q.withDO(q.DO.Unscoped())
}

func (q questionDo) Create(values ...*model.Question) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Create(values)
}

func (q questionDo) CreateInBatches(values []*model.Question, batchSize int) error {
	return q.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (q questionDo) Save(values ...*model.Question) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Save(values)
}

func (q questionDo) First() (*model.Question, error) {
	if result, err := q.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Question), nil
	}
}

func (q questionDo) Take() (*model.Question, error) {
	if result, err := q.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Question), nil
	}
}

func (q questionDo) Last() (*model.Question, error) {
	if result, err := q.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Question), nil
	}
}

func (q questionDo) Find() ([]*model.Question, error) {
	result, err := q.DO.Find()
	return result.([]*model.Question), err
}

func (q questionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Question, err error) {
	buf := make([]*model.Question, 0, batchSize)
	err = q.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (q questionDo) FindInBatches(result *[]*model.Question, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return q.DO.FindInBatches(result, batchSize, fc)
}

func (q questionDo) Attrs(attrs ...field.AssignExpr) IQuestionDo {
	return q.withDO(q.DO.Attrs(attrs...))
}

func (q questionDo) Assign(attrs ...field.AssignExpr) IQuestionDo {
	return q.withDO(q.DO.Assign(attrs...))
}

func (q questionDo) Joins(fields ...field.RelationField) IQuestionDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Joins(_f))
	}
	return &q
}

func (q questionDo) Preload(fields ...field.RelationField) IQuestionDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Preload(_f))
	}
	return &q
}

func (q questionDo) FirstOrInit() (*model.Question, error) {
	if result, err := q.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Question), nil
	}
}

func (q questionDo) FirstOrCreate() (*model.Question, error) {
	if result, err := q.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Question), nil
	}
}

func (q questionDo) FindByPage(offset int, limit int) (result []*model.Question, count int64, err error) {
	result, err = q.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = q.Offset(-1).Limit(-1).Count()
	return
}

func (q questionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = q.Count()
	if err != nil {
		return
	}

	err = q.Offset(offset).Limit(limit).Scan(result)
	return
}

func (q questionDo) Scan(result interface{}) (err error) {
	return q.DO.Scan(result)
}

func (q questionDo) Delete(models ...*model.Question) (result gen.ResultInfo, err error) {
	return q.DO.Delete(models)
}

func (q *questionDo) withDO(do gen.Dao) *questionDo {
	q.DO = *do.(*gen.DO)
	return q
}