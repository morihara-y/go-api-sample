package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/morihara-y/go-api-sample/interfaces/api/server/httputil"

	"github.com/morihara-y/go-api-sample/application/usecase"

	"github.com/gin-gonic/gin"
)

type cmntHandler struct {
	apiUsecase usecase.APIUsecase
}

// NewCmntHandler returns CommentHandler
func NewCmntHandler(apiUsecase usecase.APIUsecase) CommentHandler {
	return &cmntHandler{
		apiUsecase: apiUsecase,
	}
}

func (h *cmntHandler) FetchComment(ctx *gin.Context) {
	idstr := ctx.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Sprintln(err)
		httputil.NewHTTPError(ctx, http.StatusBadRequest, err)
		return
	}
	result, err := h.apiUsecase.FetchComment(id)
	if err != nil {
		fmt.Sprintln(err)
		httputil.NewHTTPError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
