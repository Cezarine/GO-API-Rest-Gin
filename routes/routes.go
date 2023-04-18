package routes

import (
	"github.com/cezarine/API-GO-GIN/controllers"
	"github.com/gin-gonic/gin"
)

const port = ":8000"

func HandRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.GetAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.Run(port)
}
