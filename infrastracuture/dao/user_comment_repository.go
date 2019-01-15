package dao

import (
	"database/sql"

	"github.com/go-gorp/gorp"

	"github.com/morihara-y/go-api-sample/domain/model"
	"github.com/morihara-y/go-api-sample/domain/repository"
)

type userCommentRepository struct{}

// NewUserCommentRepository returns repository.UserCommentRepository
func NewUserCommentRepository() repository.UserCommentRepository {
	return &userCommentRepository{}
}

func (r *userCommentRepository) FetchByID(id int) (comment model.Comment, err error) {
	db, err := sql.Open("mysql", "comment:password@/comment")
	if err != nil {
		db.Close()
		return comment, err
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	defer dbmap.Db.Close()

	err = dbmap.SelectOne(&comment, "select * from user_comment where comment_id=?", id)
	if err != nil {
		return comment, err
	}

	return comment, nil
}
