package usecase

import (
	"github.com/morihara-y/go-api-sample/domain/model"
	"github.com/morihara-y/go-api-sample/domain/service"
)

// APIUsecase is interface to run api
type APIUsecase interface {
	FetchComment(id int) (*model.CommentResponse, error)
}

type apiUsecase struct {
	commentService service.CommentService
}

// NewAPIUsecase returns APIUsecase
func NewAPIUsecase(commentService service.CommentService) APIUsecase {
	return &apiUsecase{
		commentService: commentService,
	}
}

func (u *apiUsecase) FetchComment(id int) (*model.CommentResponse, error) {
	comment, err := u.commentService.FetchByID(id)
	if err != nil {
		return nil, err
	}
	commentResponse := &model.CommentResponse{
		Data: comment,
	}
	return commentResponse, nil
}
