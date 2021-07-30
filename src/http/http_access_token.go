package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/msd79/bookstrore_outh-api/src/domain/access_token"
)

type AccessTokenHandler interface {
	GetByID(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetByID(c *gin.Context) {
	c.JSON(http.StatusNotImplemented)
}
