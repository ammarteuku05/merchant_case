package app

import (
	"fmt"
	"merchant-service/handler"
	"merchant-service/infra"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	router.Use(handler.CORSMiddleware())
	infra.RegisterApi(router)

	port := os.Getenv("APP_PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
