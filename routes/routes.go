package routes

import (
	"github.com/cezarine/API-GO-GIN/controllers"
	"github.com/gin-gonic/gin"
)

const port = ":8000"

func HandRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("", controllers.ExibePaginaIndex)
	r.GET("/alunos", controllers.GetAlunos)
	r.GET("/alunos/cpf/:cpf", controllers.GetAlunoCPF)
	r.GET("/:nome", controllers.Saudacao)
	r.GET("/alunos/:id/:ativo", controllers.GetAluno)
	r.GET("/alunos/:id", controllers.GetAluno)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.DELETE("/alunos/deleta/:id", controllers.DeletaAluno)
	r.PUT("/alunos/:id/:ativo", controllers.AtivaInativaAluno)
	r.PUT("/alunos/:id", controllers.EditaAluno)
	r.Run(port)
}
