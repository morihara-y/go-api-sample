package service

import (
	"github.com/morihara-y/go-api-sample/domain/model"
	"github.com/morihara-y/go-api-sample/domain/repository"
)

// CommentService is interface to fetch comments from go files
type CommentService interface {
	FetchByID(id int) (model.Comment, error)
}

type commentService struct {
	userCommentRepo repository.UserCommentRepository
}

// NewCommentService returns CommentService
func NewCommentService(userCommentRepo repository.UserCommentRepository) CommentService {
	return &commentService{
		userCommentRepo: userCommentRepo,
	}
}

func (s *commentService) FetchByID(id int) (model.Comment, error) {
	return s.userCommentRepo.FetchByID(id)
}
