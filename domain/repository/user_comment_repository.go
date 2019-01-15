package repository

import (
	"github.com/morihara-y/go-api-sample/domain/model"
)

// UserCommentRepository is interface to fetch from user_comment.
type UserCommentRepository interface {
	FetchByID(id int) (comment model.Comment, err error)
}
