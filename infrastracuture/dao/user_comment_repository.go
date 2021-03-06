package dao

import (
	"database/sql"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"

	"github.com/morihara-y/go-api-sample/domain/model"
	"github.com/morihara-y/go-api-sample/domain/repository"
)

type userCommentRepository struct{}

// NewUserCommentRepository returns repository.UserCommentRepository
func NewUserCommentRepository() repository.UserCommentRepository {
	return &userCommentRepository{}
}

func (r *userCommentRepository) FetchByID(id int) (comment model.Comment, err error) {
	db, err := sql.Open("mysql", "comment:password@localhost:3306/comment")
	if err != nil {
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
