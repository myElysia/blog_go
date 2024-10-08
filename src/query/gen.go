// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q          = new(Query)
	Article    *article
	Category   *category
	Comment    *comment
	Permission *permission
	Role       *role
	Tag        *tag
	Todo       *todo
	User       *user
	UserGroup  *userGroup
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Article = &Q.Article
	Category = &Q.Category
	Comment = &Q.Comment
	Permission = &Q.Permission
	Role = &Q.Role
	Tag = &Q.Tag
	Todo = &Q.Todo
	User = &Q.User
	UserGroup = &Q.UserGroup
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:         db,
		Article:    newArticle(db, opts...),
		Category:   newCategory(db, opts...),
		Comment:    newComment(db, opts...),
		Permission: newPermission(db, opts...),
		Role:       newRole(db, opts...),
		Tag:        newTag(db, opts...),
		Todo:       newTodo(db, opts...),
		User:       newUser(db, opts...),
		UserGroup:  newUserGroup(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Article    article
	Category   category
	Comment    comment
	Permission permission
	Role       role
	Tag        tag
	Todo       todo
	User       user
	UserGroup  userGroup
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:         db,
		Article:    q.Article.clone(db),
		Category:   q.Category.clone(db),
		Comment:    q.Comment.clone(db),
		Permission: q.Permission.clone(db),
		Role:       q.Role.clone(db),
		Tag:        q.Tag.clone(db),
		Todo:       q.Todo.clone(db),
		User:       q.User.clone(db),
		UserGroup:  q.UserGroup.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:         db,
		Article:    q.Article.replaceDB(db),
		Category:   q.Category.replaceDB(db),
		Comment:    q.Comment.replaceDB(db),
		Permission: q.Permission.replaceDB(db),
		Role:       q.Role.replaceDB(db),
		Tag:        q.Tag.replaceDB(db),
		Todo:       q.Todo.replaceDB(db),
		User:       q.User.replaceDB(db),
		UserGroup:  q.UserGroup.replaceDB(db),
	}
}

type queryCtx struct {
	Article    IArticleDo
	Category   ICategoryDo
	Comment    ICommentDo
	Permission IPermissionDo
	Role       IRoleDo
	Tag        ITagDo
	Todo       ITodoDo
	User       IUserDo
	UserGroup  IUserGroupDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Article:    q.Article.WithContext(ctx),
		Category:   q.Category.WithContext(ctx),
		Comment:    q.Comment.WithContext(ctx),
		Permission: q.Permission.WithContext(ctx),
		Role:       q.Role.WithContext(ctx),
		Tag:        q.Tag.WithContext(ctx),
		Todo:       q.Todo.WithContext(ctx),
		User:       q.User.WithContext(ctx),
		UserGroup:  q.UserGroup.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
