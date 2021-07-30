package app

import (
	"github.com/gin-gonic/gin"
	"github.com/msd79/bookstrore_outh-api/src/domain/access_token"
	"github.com/msd79/bookstrore_outh-api/src/http"
	"github.com/msd79/bookstrore_outh-api/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atService := access_token.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token:id", atHandler.GetByID)

	router.Run(":8080")
}
