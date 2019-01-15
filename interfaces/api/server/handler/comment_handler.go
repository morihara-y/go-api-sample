package handler

import (
	"github.com/gin-gonic/gin"
)

// CommentHandler is interface of handler to search comments
type CommentHandler interface {
	FetchComment(ctx *gin.Context)
}
